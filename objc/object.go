// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

// Package objc implements access to the Objective-C runtime from Go
package objc

import "C"

import (
	"fmt"
	"strings"
	"unsafe"
)

// An Object represents an Objective-C object, along with
// some convenience methods only found on NSObjects.
type Object interface {
	// SendMsg sends an arbitrary message to the method on the
	// object that is identified by selectorName.
	Send(selector string, args ...interface{}) Object

	// SendSuperMsg is like SendMsg, but sends to the object's
	// super class instead.
	SendSuper(selector string, args ...interface{}) Object

	// Class returns the the special class object corresponding
	// to this object.
	Class() Class

	// Alloc sends the  "alloc" message to the object.
	Alloc() Object

	// Init sends the "init" message to the object.
	Init() Object

	// Retain sends the "retain" message to the object.
	Retain() Object

	// Release sends the "release" message to the object.
	Release() Object

	// Autorelease sends the "autorelease" message to the object.
	Autorelease() Object

	// Copy sends the "copy" message to the object.
	Copy() Object

	Equals(o Object) bool

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

	// CString returns the value of the object as a C string.
	CString() string

	Set(setter string, args ...interface{})
	Get(getter string) Object
	GetSt(getter string, ret interface{})
}

// Type object is the package's internal representation of an Object.
// Besides implementing the Object interface, object also implements
// the Class interface.
type object struct {
	ptr uintptr
}

func ObjectPtr(ptr uintptr) Object {
	return object{ptr: ptr}
}

// Pointer returns the object as a uintptr.
//
// Using package unsafe, this uintptr can further
// be converted to an unsafe.Pointer.
func (obj object) Pointer() uintptr {
	return obj.ptr
}

func (obj object) Class() Class {
	cls := obj.Send("class").(Class)
	lazilyRegisterClassInMap(cls.(object).className())
	return cls
}

func (obj object) Equals(o Object) bool {
	return obj.Pointer() == o.Pointer()
}

func (obj object) Alloc() Object {
	return obj.Send("alloc")
}

func (obj object) Init() Object {
	return obj.Send("init")
}

func (obj object) Retain() Object {
	return obj.Send("retain")
}

func (obj object) Release() Object {
	return obj.Send("release")
}

func (obj object) Autorelease() Object {
	return obj.Send("autorelease")
}

func (obj object) Copy() Object {
	return obj.Send("copy")
}

func (obj object) Get(getter string) Object {
	return obj.Send(getter)
}

func (obj object) GetSt(getter string, ret interface{}) {
	obj.Send(getter, ret)
}

func (obj object) Set(setter string, args ...interface{}) {
	obj.Send(fmt.Sprintf("set%s", strings.Title(setter)), args...)
}

func (obj object) CString() string {
	if obj.Pointer() == 0 {
		return ""
	}

	return C.GoString((*C.char)(unsafe.Pointer(obj.Pointer())))
}

func (obj object) String() string {
	// TODO: some kind of recover to catch when this doesnt work

	pool := GetClass("NSAutoreleasePool").Alloc().Init()
	defer pool.Release()

	bytes := obj.Send("description").Send("UTF8String")
	if bytes.Pointer() == 0 {
		return "(nil)"
	}

	return bytes.CString()
}
