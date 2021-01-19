package resource

import (
	"reflect"

	"github.com/progrium/macdriver/objc"
	"github.com/rs/xid"
)

type Handle string

func NewHandle() *Handle {
	handle := Handle(xid.New().String())
	return &handle
}

func GetHandle(v interface{}) *Handle {
	return reflect.Indirect(reflect.ValueOf(v)).Field(0).Interface().(*Handle)
}

func SetHandle(v interface{}, h string) {
	hh := GetHandle(v)
	if hh != nil {
		*hh = Handle(h)
	}
}

type Applier interface {
	Apply(objc.Object) error
}

type Discarder interface {
	Discard(objc.Object) error
}
