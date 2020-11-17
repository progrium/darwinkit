package macdriver

import (
	"fmt"
	"reflect"
	"strings"

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

type resource struct {
	handle Handle
}

func (r *resource) Handle() Handle {
	return r.handle
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
