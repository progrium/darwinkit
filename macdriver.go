package macdriver

import (
	"fmt"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"github.com/mitchellh/mapstructure"
	"github.com/progrium/macdriver/pkg/cocoa"
	"github.com/progrium/macdriver/pkg/core"
	"github.com/progrium/macdriver/pkg/objc"
	"github.com/rs/xid"
)

var (
	mutations chan mutation
)

func init() {
	mutations = make(chan mutation)
}

type mutation struct {
	fn  func(r Resource) (interface{}, error)
	ret chan interface{}
	r   Resource
}

type AppDelegate struct {
	objc.Object `objc:"AppDelegate : NSObject"`
}

func (delegate *AppDelegate) ApplicationDidFinishLaunching(notification objc.Object) {
	for m := range mutations {
		v, err := m.fn(m.r)
		m.ret <- v
		m.ret <- err
	}
}

func MutateInMainThread(fn func(r Resource) (interface{}, error), r Resource) (interface{}, error) {
	ret := make(chan interface{}, 2)
	mutations <- mutation{
		fn:  fn,
		ret: ret,
		r:   r,
	}
	v := <-ret
	e := <-ret
	var err error
	if e != nil {
		err = e.(error)
	}
	return v, err
}

type Handle string

func (h Handle) Type() string {
	parts := strings.Split(string(h), ":")
	return parts[0]
}

type Resource interface {
	Handle() Handle
	Apply(old, new reflect.Value, target objc.Object) error
	Init(v reflect.Value) (objc.Object, error)
}

type resource struct {
	handle Handle
}

func (r *resource) Handle() Handle {
	return r.handle
}

type Menu struct {
	resource

	Icon    string
	Title   string
	Tooltip string
	Items   []*MenuItem
}

func (m *Menu) Apply(old, new reflect.Value, target objc.Object) error {
	return nil
}

func (m *Menu) Init(v reflect.Value) (objc.Object, error) {
	return cocoa.NSMenu_New(), nil
}

type MenuItem struct {
	resource

	Title   string
	Icon    string
	Tooltip string
	Enabled bool
	Checked bool
}

func (m *MenuItem) Apply(old, new reflect.Value, target objc.Object) error {
	return nil
}

func (m *MenuItem) Init(v reflect.Value) (objc.Object, error) {
	return cocoa.NSMenuItem_New(), nil
}

type StatusItem struct {
	resource

	Icon string
	Text string
	Menu *Menu
}

func (s *StatusItem) Apply(old, new reflect.Value, target objc.Object) error {
	return nil
}

func (s *StatusItem) Init(v reflect.Value) (objc.Object, error) {
	return cocoa.NSStatusBar_System().StatusItemWithLength(cocoa.NSVariableStatusItemLength), nil
}

type Window struct {
	resource

	Title    string
	Position Point
	Size     Point
}

func (w *Window) Apply(old, new reflect.Value, target objc.Object) error {
	fmt.Fprintln(os.Stderr, "APPLY", w, target)
	obj := cocoa.NSWindow{Object: target}
	obj.SetTitle(w.Title)
	obj.SetFrameDisplay(core.Rect(w.Position.X, w.Position.Y, w.Size.X, w.Size.Y), true)
	return nil
}

func (w *Window) Init(v reflect.Value) (objc.Object, error) {
	return cocoa.NSWindow_Init(core.Rect(0, 0, 0, 0), cocoa.NSTitledWindowMask, cocoa.NSBackingStoreBuffered, false), nil
}

type Point struct {
	X float64
	Y float64
}

func NewHandle(type_ string) Handle {
	return Handle(fmt.Sprintf("%s:%s", type_, xid.New().String()))
}

type State struct {
	Menus       []*Menu
	StatusItems []*StatusItem
	Windows     []*Window

	objects    map[Handle]objc.Object
	lastValues map[Handle]reflect.Value
}

func NewState() *State {
	return &State{
		objects:    make(map[Handle]objc.Object),
		lastValues: make(map[Handle]reflect.Value),
	}
}

func (s *State) Update(handle Handle, data interface{}) error {
	v, err := s.Lookup(handle)
	if err != nil {
		return err
	}
	if err := mapstructure.Decode(data, v.Interface()); err != nil {
		return err
	}
	return nil
}

func (s *State) Init(type_ string, data interface{}) (string, error) {
	t, found := map[string]reflect.Type{
		"Window": reflect.TypeOf(Window{}),
	}[type_]
	if !found {
		return "", fmt.Errorf("type not found")
	}
	rv := reflect.New(t)
	v := rv.Interface()
	if err := mapstructure.Decode(data, v); err != nil {
		return "", err
	}
	r, ok := v.(*Window)
	if !ok {
		panic("not a resource")
	}
	r.resource.handle = NewHandle(type_)
	s.Windows = append(s.Windows, v.(*Window))
	return string(r.resource.handle), nil
}

func (s *State) Lookup(handle Handle) (found reflect.Value, err error) {
	// TODO: stop return value for Walk so it doesn't continue crawling once found
	err = Walk(s, func(v reflect.Value, parent reflect.Value, path []string) error {
		r, ok := v.Interface().(Resource)
		if ok {
			if r.Handle() == handle {
				found = v
			}
		}
		return nil
	})
	return found, err
}

func (s *State) Reconcile() error {
	return Walk(s, func(v reflect.Value, parent reflect.Value, path []string) error {
		r, ok := v.Interface().(Resource)
		if ok {
			var err error
			old := s.lastValues[r.Handle()]
			target, exists := s.objects[r.Handle()]
			if !exists {
				var target_ interface{}
				if target_, err = MutateInMainThread(func(r Resource) (interface{}, error) {
					return r.Init(v)
				}, r); err != nil {
					return err
				}
				s.objects[r.Handle()] = target_.(objc.Object)
			}
			_, err = MutateInMainThread(func(r Resource) (interface{}, error) {
				if err := r.Apply(old, v, target); err != nil {
					return nil, err
				}
				return nil, nil
			}, r)
			if err != nil {
				return err
			}
		}
		return nil
	})
}

func walk(v reflect.Value, path []string, visitor func(v reflect.Value, parent reflect.Value, path []string) error) error {
	for _, k := range keys(v) {
		subpath := append(path, k)
		vv := prop(v, k)
		if err := visitor(vv, v, subpath); err != nil {
			return err
		}
		if err := walk(vv, subpath, visitor); err != nil {
			return err
		}
	}
	return nil
}

func Walk(v interface{}, visitor func(v reflect.Value, parent reflect.Value, path []string) error) error {
	return walk(reflect.ValueOf(v), []string{}, visitor)
}

func prop(robj reflect.Value, key string) reflect.Value {
	rtyp := robj.Type()
	switch rtyp.Kind() {
	case reflect.Slice, reflect.Array:
		idx, err := strconv.Atoi(key)
		if err != nil {
			panic("non-numeric index given for slice")
		}
		rval := robj.Index(idx)
		if rval.IsValid() {
			return rval
		}
	case reflect.Ptr:
		return prop(robj.Elem(), key)
	case reflect.Map:
		rval := robj.MapIndex(reflect.ValueOf(key))
		if rval.IsValid() {
			return rval
		}
	case reflect.Struct:
		rval := robj.FieldByName(key)
		if rval.IsValid() {
			return rval
		}
		for i := 0; i < rtyp.NumField(); i++ {
			field := rtyp.Field(i)
			tag := strings.Split(field.Tag.Get("json"), ",")
			if tag[0] == key || field.Name == key {
				return robj.FieldByName(field.Name)
			}
		}
		panic("struct field not found: " + key)
	}
	//spew.Dump(robj, key)
	panic("unexpected kind: " + rtyp.Kind().String())
}

func keys(v reflect.Value) []string {
	switch v.Type().Kind() {
	case reflect.Map:
		var keys []string
		for _, key := range v.MapKeys() {
			k, ok := key.Interface().(string)
			if !ok {
				continue
			}
			keys = append(keys, k)
		}
		sort.Sort(sort.StringSlice(keys))
		return keys
	case reflect.Struct:
		t := v.Type()
		var f []string
		for i := 0; i < t.NumField(); i++ {
			name := t.Field(i).Name
			// first letter capitalized means exported
			if name[0] == strings.ToUpper(name)[0] {
				f = append(f, name)
			}
		}
		return f
	case reflect.Slice, reflect.Array:
		var k []string
		for n := 0; n < v.Len(); n++ {
			k = append(k, strconv.Itoa(n))
		}
		return k
	case reflect.Ptr:
		return keys(v.Elem())
	default:
		return []string{}
	}
}
