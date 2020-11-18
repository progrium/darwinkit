package core

import (
	"fmt"
	"reflect"
	"sync"

	"github.com/progrium/macdriver/pkg/objc"
)

var (
	callbacks sync.Map
	cbCounter uint64
)

type callbackEntry struct {
	cb       interface{}
	class    objc.Class
	instance objc.Object
}

func Callback(fn interface{}) (objc.Object, objc.Selector) {
	rfn := reflect.ValueOf(fn)
	cb, ok := callbacks.Load(rfn.Type().String())
	if !ok {
		cbCounter++
		delegateName := fmt.Sprintf("CallbackDelegate%d", cbCounter)
		c := objc.NewClass(delegateName, "NSObject")
		c.AddMethod("callback:", fn)
		objc.RegisterClass(c)
		cb = callbackEntry{
			cb:       fn,
			class:    c,
			instance: objc.Get(delegateName).Alloc().Init(),
		}
		callbacks.Store(rfn.Type().String(), cb)
	}
	return cb.(callbackEntry).instance, objc.Sel("callback:")
}
