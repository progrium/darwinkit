package core

import (
	"fmt"
	"sync"

	"github.com/progrium/macdriver/objc"
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
	// TODO: better way to identify the same function thats already been registered
	cbCounter++
	delegateName := fmt.Sprintf("CallbackDelegate%d", cbCounter)
	c := objc.NewClass(delegateName, "NSObject")
	c.AddMethod("callback:", fn)
	objc.RegisterClass(c)
	cb := callbackEntry{
		cb:       fn,
		class:    c,
		instance: objc.Get(delegateName).Alloc().InitObject(),
	}
	//fmt.Fprintf(os.Stderr, "%#v %d\n", delegateName, &fn)
	//callbacks.Store(rfn.Type().String(), cb)
	return cb.instance, objc.Sel("callback:")
}
