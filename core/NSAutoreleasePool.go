package core

import "github.com/progrium/macdriver/objc"

type NSAutoreleasePool struct {
	objc.Object
}

var nsAutoreleasePool = objc.Get("NSAutoreleasePool")

func NSAutoreleasePool_New() NSAutoreleasePool {
	return NSAutoreleasePool{nsAutoreleasePool.Alloc().Init()}
}
