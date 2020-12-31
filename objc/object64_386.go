// Copyright (c) 2013 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import "C"

import (
	"math"
	"unsafe"
)

// object64 is an object that holds
// a 64-bit value on a 32-bit architecture.
type object64 struct {
	big uint64
}

func (obj object64) Send(selector string, args ...interface{}) Object {
	return sendMsg(obj, "objc_msgSend", selector, args...)
}

func (obj object64) SendSuper(selector string, args ...interface{}) Object {
	return sendMsg(obj, "objc_msgSendSuper", selector, args...)
}

func (obj object64) Pointer() uintptr {
	return uintptr(obj.big & 0xffffffff)
}

func (obj object64) Alloc() Object {
	return obj.Send("alloc")
}

func (obj object64) Init() Object {
	return obj.Send("init")
}

func (obj object64) Retain() Object {
	return obj.Send("retain")
}

func (obj object64) Release() Object {
	return obj.Send("release")
}

func (obj object64) AutoRelease() Object {
	return obj.Send("autorelease")
}

func (obj object64) Copy() Object {
	return obj.Send("copy")
}

func (obj object64) String() string {
	pool := GetClass("NSAutoreleasePool").Alloc().Init()
	defer pool.Release()

	descString := obj.Send("description")
	utf8Bytes := descString.Send("UTF8String")
	if utf8Bytes.Pointer() != 0 {
		return C.GoString((*C.char)(unsafe.Pointer(utf8Bytes.Pointer())))
	}

	return "(nil)"
}

func (obj object64) Uint() uint64 {
	return uint64(obj.big)
}

func (obj object64) Int() int64 {
	return int64(obj.big)
}

func (obj object64) Bool() bool {
	return obj.big == 1
}

func (obj object64) Float() float64 {
	return math.Float64frombits(obj.big)
}
