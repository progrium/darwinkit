package resource

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/progrium/macdriver/objc"
	"github.com/rs/xid"
)

type Handle string

func NewHandle(t string) *Handle {
	handle := Handle(fmt.Sprintf("%s:%s", t, Handle(xid.New().String())))
	return &handle
}

func HasHandle(v interface{}) bool {
	rv := reflect.Indirect(reflect.ValueOf(v))
	if rv.Kind() == reflect.Struct && rv.Type().NumField() > 0 && rv.Type().Field(0).Name == "Handle" {
		return true
	}
	return false
}

func GetHandle(v interface{}) *Handle {
	if !HasHandle(v) {
		return nil
	}
	rv := reflect.Indirect(reflect.ValueOf(v))
	h := rv.Field(0).Interface()
	hh, ok := h.(*Handle)
	if !ok {
		return nil
	}
	return hh
}

func SetHandle(v interface{}, h string) {
	if !HasHandle(v) {
		return
	}
	handle := Handle(h)
	reflect.Indirect(reflect.ValueOf(v)).Field(0).Set(reflect.ValueOf(&handle))
}

func (h *Handle) Type() string {
	parts := strings.Split(string(*h), ":")
	return parts[0]
}

func (h *Handle) ID() string {
	parts := strings.Split(string(*h), ":")
	if len(parts) > 1 {
		return parts[1]
	}
	return ""
}

func (h *Handle) Handle() string {
	if h == nil {
		return ""
	}
	return string(*h)
}

type Applier interface {
	Apply(objc.Object) (objc.Object, error)
}

type Discarder interface {
	Discard(objc.Object) error
}
