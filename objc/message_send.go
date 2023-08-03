// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

// #import <objc/message.h>
// void* msg_send_func = objc_msgSend;
// void* msgSendSuper_func = objc_msgSendSuper;
// #ifndef __arm64__
// void* msgSend_stret_func = objc_msgSend_stret;
// void* msgSendSuper_stret_func = objc_msgSendSuper_stret;
// #else
// void* msgSend_stret_func = objc_msgSend;
// void* msgSendSuper_stret_func = objc_msgSendSuper;
// #endif
import "C"
import (
	"reflect"
	"unsafe"

	"github.com/progrium/macdriver/objc/ffi"
)

// MsgSend wraps objc_msgSend
// type param T: the return value type. using Void for non-return value method.
func MsgSend[T any](receiver Pointer, selector Selector, params ...any) T {
	return callMsgSend[T](receiver, selector, C.msg_send_func, params...)
}

// objc_msgSendSuper
func MsgSendSuper[T any](receiver Pointer, selector Selector, params ...any) T {
	return callMsgSend[T](receiver, selector, C.msgSendSuper_func, params...)
}

// objc_msgSend_stret
// Note: arm64 do not need to use this for method that return a struct type.
func MsgSendStret[T any](receiver Pointer, selector Selector, params ...any) T {
	return callMsgSend[T](receiver, selector, C.msgSend_stret_func, params...)
}

// objc_msgSendSuper_stret
// Note: arm64 do not need to use this for method that return a struct type.
func MsgSendSuperStret[T any](receiver Pointer, selector Selector, params ...any) T {
	return callMsgSend[T](receiver, selector, C.msgSendSuper_stret_func, params...)
}

func callMsgSend[T any](receiver Pointer, selector Selector, fn unsafe.Pointer, params ...any) T {
	ptr := receiver.Ptr()
	if ptr == nil {
		panic("object is nil")
	}
	argc := len(params)

	var args = make([]unsafe.Pointer, argc+2)
	var argTypes = make([]*ffi.Type, argc+2)
	args[0] = unsafe.Pointer(&ptr)
	argTypes[0] = ffi.TypePointer
	args[1] = unsafe.Pointer(&selector.ptr)
	argTypes[1] = ffi.TypePointer
	for i := 0; i < argc; i++ {
		v := reflect.ValueOf(params[i])
		args[i+2] = convertToObjcValue(v)
		argTypes[i+2] = toFFIType(v.Type())
	}

	var ret T
	var retPtr unsafe.Pointer
	rt := reflect.TypeOf(ret)
	if directPointer(rt) {
		retPtr = unsafe.Pointer(&ret)
	} else {
		retPtr = convertToObjcValue(reflect.ValueOf(ret))
	}
	retType := toFFIType(reflect.TypeOf(ret))

	var imp IMP = MakeIMP(fn)
	cif, status := ffi.PrepCIF(retType, argTypes)
	if status != ffi.OK {
		panic("ffi prep cif status not ok")
	}
	ffi.Call(cif, imp.ptr, retPtr, args)
	if directPointer(rt) {
		return *(*T)(retPtr)
	}

	return convertToGoValue(retPtr, reflect.TypeOf(ret)).Interface().(T)
}
