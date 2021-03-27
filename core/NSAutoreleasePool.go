package core

import "github.com/progrium/macdriver/objc"

type NSAutoreleasePool struct {
	objc.Object
}

var NSAutoreleasePool_ = objc.Get("NSAutoreleasePool")

func NSAutoreleasePool_New() NSAutoreleasePool {
	return NSAutoreleasePool{NSAutoreleasePool_.Alloc().Init()}
}
