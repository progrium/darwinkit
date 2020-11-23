package macdriver

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/manifold/qtalk/golang/rpc"
	"github.com/progrium/macdriver/pkg/objc"
	"github.com/rs/xid"
)

var initTypes = map[string]reflect.Type{
	"Window":     reflect.TypeOf(Window{}),
	"StatusItem": reflect.TypeOf(StatusItem{}),
}

type Resource interface {
	Handle() Handle
	Apply(old, new reflect.Value, target objc.Object) (objc.Object, error)
}

type resourcer interface {
	res() *resource
}

func SetHandle(v interface{}, h string) {
	rr, ok := v.(resourcer)
	if ok {
		r := rr.res()
		r.handle = Handle(h)
	}
}

func Sync(p *rpc.Peer, v interface{}) error {
	rr, ok := v.(resourcer)
	if !ok {
		return fmt.Errorf("value to sync is not a resource")
	}
	r := rr.res()
	handle := string(r.handle)
	if handle == "" {
		handle = reflect.ValueOf(v).Type().Elem().Name()
	}
	_, err := p.Call("Apply", []interface{}{handle, v}, &handle)
	r.handle = Handle(handle)
	return err
}

func Release(p *rpc.Peer, v interface{}) error {
	rr, ok := v.(resourcer)
	if !ok {
		return fmt.Errorf("value to sync is not a resource")
	}
	r := rr.res()
	handle := string(r.handle)
	if handle == "" {
		return fmt.Errorf("unable to release an uninitialized resource")
	}
	_, err := p.Call("Release", handle, nil)
	if err == nil {
		r.handle = Handle("")
	}
	return err
}

type resource struct {
	handle Handle
}

func (r *resource) Handle() Handle {
	return r.handle
}

func (r *resource) res() *resource {
	return r
}

type Handle string

func NewHandle(type_ string) Handle {
	return Handle(fmt.Sprintf("%s:%s", type_, xid.New().String()))
}

func (h Handle) Init() (reflect.Value, error) {
	t, found := initTypes[h.Type()]
	if !found {
		return reflect.Value{}, fmt.Errorf("type not found")
	}
	return reflect.New(t), nil
}

func (h Handle) Type() string {
	parts := strings.Split(string(h), ":")
	return parts[0]
}

func (h Handle) ID() string {
	parts := strings.Split(string(h), ":")
	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}

func (h Handle) String() string {
	return string(h)
}
