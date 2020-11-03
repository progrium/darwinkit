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

func (obj object64) SendMsg(selector string, args ...interface{}) Object {
	return sendMsg(obj, "objc_msgSend", selector, args...)
}

func (obj object64) SendSuperMsg(selector string, args ...interface{}) Object {
	return sendMsg(obj, "objc_msgSendSuper", selector, args...)
}

func (obj object64) Pointer() uintptr {
	return uintptr(obj.big & 0xffffffff)
}

func (obj object64) Alloc() Object {
	return obj.SendMsg("alloc")
}

func (obj object64) Init() Object {
	return obj.SendMsg("init")
}

func (obj object64) Retain() Object {
	return obj.SendMsg("retain")
}

func (obj object64) Release() Object {
	return obj.SendMsg("release")
}

func (obj object64) AutoRelease() Object {
	return obj.SendMsg("autorelease")
}

func (obj object64) Copy() Object {
	return obj.SendMsg("copy")
}

func (obj object64) String() string {
	pool := GetClass("NSAutoreleasePool").Alloc().Init()
	defer pool.Release()

	descString := obj.SendMsg("description")
	utf8Bytes := descString.SendMsg("UTF8String")
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
