package bridge

import (
	"context"
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strconv"
	"strings"
	"sync"

	"github.com/manifold/qtalk/golang/mux"
	"github.com/manifold/qtalk/golang/rpc"
	"github.com/mitchellh/mapstructure"
	"github.com/progrium/macdriver/pkg/cocoa"
	"github.com/progrium/macdriver/pkg/core"
	"github.com/progrium/macdriver/pkg/objc"
)

var (
	state *State
)

func Run() {
	state = NewState()
	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		session := mux.NewSession(context.Background(), struct {
			io.ReadCloser
			io.Writer
		}{os.Stdin, os.Stdout})
		peer := rpc.NewPeer(session, rpc.JSONCodec{})
		peer.Bind("debug", debug)
		peer.Bind("Apply", state.Apply)
		peer.Bind("Release", state.Release)
		//peer.Bind("Invoke", state.Invoke)
		go peer.Respond()
	})
	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}

func debug(handle string, data interface{}) string {
	return fmt.Sprintf("%#v", state.StatusItems[0].Menu)
}

type State struct {
	Menus       []*Menu
	StatusItems []*StatusItem
	Windows     []*Window

	objects    map[Handle]objc.Object
	released   map[Handle]bool
	lastValues map[Handle]reflect.Value

	caller rpc.Caller

	sync.Mutex
}

func Invoke(ptr string) error {
	// todo args
	ef, ok := exportedFuncs[ptr]
	// reflect call
	if ok {
		ef.fn.(func())()
	}
	return nil
}

func NewState() *State {
	return &State{
		objects:    make(map[Handle]objc.Object),
		lastValues: make(map[Handle]reflect.Value),
		released:   make(map[Handle]bool),
	}
}

func (s *State) Release(h string) (err error) {
	s.Lock()
	handle := Handle(h)
	s.released[handle] = true
	core.Dispatch(func() {
		//fmt.Fprintf(os.Stderr, "RECONCILING %v\n", handle)
		if err := s.Reconcile(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		s.Unlock()
	})
	// TODO: remove from State slices
	return nil
}

func (s *State) Apply(h string, patch interface{}, call *rpc.Call) (handle Handle, err error) {
	s.Lock()
	handle = Handle(h)

	Walk(patch, func(v, p reflect.Value, path []string) error {
		if path[len(path)-1] == "$fnptr" {
			p.SetMapIndex(reflect.ValueOf("Caller"), reflect.ValueOf(call.Caller))
		}
		return nil
	})

	if handle.ID() != "" {
		v, err := s.Lookup(handle)
		if err != nil {
			return handle, err
		}
		if v.IsValid() {
			if err := mapstructure.Decode(patch, v.Interface()); err != nil {
				return handle, err
			}
		}
	} else {

		v, err := handle.Init()
		if err != nil {
			return handle, err
		}
		if err := mapstructure.Decode(patch, v.Interface()); err != nil {
			return handle, err
		}
		handle = NewHandle(handle.Type())
		switch handle.Type() {
		case "Window":
			r := v.Interface().(*Window)
			r.handle = handle
			s.Windows = append(s.Windows, r)
		case "StatusItem":
			r := v.Interface().(*StatusItem)
			r.handle = handle
			s.StatusItems = append(s.StatusItems, r)
		default:
			panic("unknown type: " + handle.Type())
		}
	}
	core.Dispatch(func() {
		//fmt.Fprintf(os.Stderr, "RECONCILING %v\n", handle)
		if err := s.Reconcile(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
		s.Unlock()
	})
	return handle, err
}

func (s *State) Lookup(handle Handle) (found reflect.Value, err error) {
	// TODO: stop return value for Walk so it doesn't continue crawling once found
	err = Walk(s, func(v reflect.Value, parent reflect.Value, path []string) error {
		r, ok := v.Interface().(Resource)
		if ok && !v.IsNil() {
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
		if ok && !v.IsNil() {
			var err error
			old := s.lastValues[r.Handle()]
			if s.released[r.Handle()] {
				// if in released but not in objects,
				// its stale state that should have been cleaned up
				// so we will ignore it here
				if _, ok := s.objects[r.Handle()]; !ok {
					return nil
				}
				old = reflect.Value{}
				v = reflect.Value{}
				defer delete(s.objects, r.Handle())
			}
			target, exists := s.objects[r.Handle()]
			if exists {
				_, err := r.Apply(old, v, target)
				return err
			}
			obj, err := r.Apply(old, v, nil)
			if err != nil {
				return err
			}
			s.objects[r.Handle()] = obj
		}
		return nil
	})
}

func walk(v reflect.Value, path []string, visitor func(v reflect.Value, parent reflect.Value, path []string) error) error {
	for _, k := range keys(v) {
		subpath := append(path, k)
		vv := prop(v, k)
		if !vv.IsValid() {
			continue
		}
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
			return reflect.ValueOf(rval.Interface())
		}
	case reflect.Ptr:
		return prop(robj.Elem(), key)
	case reflect.Map:
		rval := robj.MapIndex(reflect.ValueOf(key))
		if rval.IsValid() {
			return reflect.ValueOf(rval.Interface())
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
		if !v.IsNil() {
			return keys(v.Elem())
		}
		return []string{}
	case reflect.String, reflect.Bool, reflect.Float64, reflect.Float32, reflect.Interface:
		return []string{}
	default:
		fmt.Fprintf(os.Stderr, "unexpected type: %s\n", v.Type().Kind())
		return []string{}
	}
}
