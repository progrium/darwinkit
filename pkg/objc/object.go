// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package objc implements access to the Objective-C runtime from Go
package objc

import (
	"C"
	"unsafe"
)

// An Object represents an Objective-C object, along with
// some convenience methods only found on NSObjects.
type Object interface {
	// SendMsg sends an arbitrary message to the method on the
	// object that is idenfieid by selectorName.
	SendMsg(selectorName string, args ...interface{}) Object

	// SendSuperMsg is like SendMsg, but sends to the object's
	// super class instead.
	SendSuperMsg(selectorName string, args ...interface{}) Object

	// Alloc sends the  "alloc" message to the object.
	Alloc() Object

	// Init sends the "init" message to the object.
	Init() Object

	// Retain sends the "retain" message to the object.
	Retain() Object

	// Release sends the "release" message to the object.
	Release() Object

	// AutoRelease sends the "autorelease" message to the object.
	AutoRelease() Object

	// Copy sends the "copy" message to the object.
	Copy() Object

	// String returns a string-representation of the object.
	// This is equivalent to sending the "description"
	// message to the object, except that this method
	// returns a Go string.
	String() string

	// Pointer returns the in-memory address of the object.
	Pointer() uintptr

	// Uint returns the value of the object as an uint64.
	Uint() uint64

	// Int returns the value of the object as an int64.
	Int() int64

	// Float returns the value of the object as a float64.
	Float() float64

	// Bool returns the value of the object as a bool.
	Bool() bool
}

// Type object is the package's internal representation of an Object.
// Besides implementing the Objct interface, object also implements
// the Class interface.
type object struct {
	ptr uintptr
}

// Return the Object as a uintptr.
//
// Using package unsafe, this uintptr can further
// be converted to an unsafe.Pointer.
func (obj object) Pointer() uintptr {
	return obj.ptr
}

func (obj object) Alloc() Object {
	return obj.SendMsg("alloc")
}

func (obj object) Init() Object {
	return obj.SendMsg("init")
}

func (obj object) Retain() Object {
	return obj.SendMsg("retain")
}

func (obj object) Release() Object {
	return obj.SendMsg("release")
}

func (obj object) AutoRelease() Object {
	return obj.SendMsg("autorelease")
}

func (obj object) Copy() Object {
	return obj.SendMsg("copy")
}

func (obj object) String() string {
	pool := GetClass("NSAutoreleasePool").Alloc().Init()
	defer pool.Release()

	descString := obj.SendMsg("description")
	utf8Bytes := descString.SendMsg("UTF8String")
	if utf8Bytes.Pointer() != 0 {
		return C.GoString((*C.char)(unsafe.Pointer(utf8Bytes.Pointer())))
	}

	return "(nil)"
}
