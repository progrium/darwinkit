// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

// #import <stdint.h>
// void* C_NewDeallocListener(uintptr_t id);
// void Run_WithAutoreleasePool(uintptr_t ptr);
import "C"
import (
	"runtime"
	"runtime/cgo"
)

//export deleteHandle
func deleteHandle(hp C.uintptr_t) {
	h := cgo.Handle(hp)
	h.Delete()
}

// SetDeallocListener set a listener to be invoked when object ref count is decreased to zero(so the object is dealloced).
// Call dealloc method directly, but ref count is still large than 0 will not trigger the listener.
// Call multi times will remove previouse listener.
func SetDeallocListener(o Object, listener func()) {
	h := cgo.NewHandle(listener)
	lo := ObjectFrom(C.C_NewDeallocListener(C.uintptr_t(h)))
	defer lo.Release()
	SetAssociatedObject(o, AssociationKey("deallocListener"), lo, ASSOCIATION_RETAIN)
}

// WithAutoreleasePool runs code in a new AutoreleasePool.
func WithAutoreleasePool(task func()) {
	id := cgo.NewHandle(task)
	C.Run_WithAutoreleasePool(C.uintptr_t(id))
}

// Retain will retain the object and set a finalizer for the Go GC to release.
// This is the preferred way to retain objects. Note the object
// must be passed by reference.
func Retain(o IObject) {
	o.Retain()
	runtime.SetFinalizer(o, func(o IObject) {
		o.Release()
	})
}

//export runTaskAndDeleteHandle
func runTaskAndDeleteHandle(p C.uintptr_t) {
	h := cgo.Handle(p)
	task := h.Value().(func())
	task()
	h.Delete()
}
