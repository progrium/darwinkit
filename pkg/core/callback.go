package core

import (
	"fmt"
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
	cb, ok := callbacks.Load(fn)
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
		callbacks.Store(fn, cb)
	}
	return cb.(callbackEntry).instance, objc.Sel("callback:")
}
