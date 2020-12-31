// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

/*
#cgo LDFLAGS: -lobjc -framework Foundation
#define __OBJC2__ 1
#include <objc/runtime.h>
#include <objc/message.h>
#include <stdlib.h>

void *GoObjc_GetObjectSuperClassStruct(void *obj) {
	struct objc_super *s = malloc(sizeof(struct objc_super));
	s->receiver = obj;
	s->super_class = class_getSuperclass(object_getClass(obj));
	return s;
}
*/
import "C"
import (
	"math"
	"reflect"
	"unsafe"

	"github.com/progrium/macdriver/misc/variadic"
)

func unpackStruct(val reflect.Value) []uintptr {
	memArgs := []uintptr{}
	for i := 0; i < val.NumField(); i++ {
		v := val.Field(i)
		kind := v.Kind()
		switch kind {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32:
			memArgs = append(memArgs, uintptr(v.Int()))
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uintptr:
			memArgs = append(memArgs, uintptr(v.Uint()))
		case reflect.Int64, reflect.Uint64:
			i64 := v.Uint()
			memArgs = append(memArgs, uintptr(i64&0xffffffff))
			memArgs = append(memArgs, uintptr((i64>>32)&0xffffffff))
		case reflect.Float32:
			u32 := math.Float32bits(float32(v.Float()))
			memArgs = append(memArgs, uintptr(u32))
		case reflect.Float64:
			u64 := math.Float64bits(v.Float())
			memArgs = append(memArgs, uintptr(u64&0xffffffff))
			memArgs = append(memArgs, uintptr((u64>>32)&0xffffffff))
		case reflect.Ptr:
			memArgs = append(memArgs, val.Pointer())
		case reflect.Struct:
			args := unpackStruct(v)
			memArgs = append(memArgs, args...)
		}
	}
	return memArgs
}

func sendMsg(obj Object, sendFuncName string, selector string, restArgs ...interface{}) Object {
	// Keep ObjC semantics: messages can be sent to nil objects,
	// but the response is nil.
	if obj.Pointer() == 0 {
		return obj
	}

	sel := selectorWithName(selector)
	if sel == nil {
		return nil
	}

	args := []uintptr{}
	typeInfo := simpleTypeInfoForMethod(obj, selector)

	for i, arg := range restArgs {
		switch t := arg.(type) {
		case Object:
			args = append(args, t.Pointer())
		case Selector:
			args = append(args, uintptr(selectorWithName(t.Selector())))
		case uintptr:
			args = append(args, t)
		case int:
			args = append(args, uintptr(t))
		case uint:
			args = append(args, uintptr(t))
		case int8:
			args = append(args, uintptr(t))
		case uint8:
			args = append(args, uintptr(t))
		case int16:
			args = append(args, uintptr(t))
		case uint16:
			args = append(args, uintptr(t))
		case int32:
			args = append(args, uintptr(t))
		case uint32:
			args = append(args, uintptr(t))
		case int64:
			args = append(args, uintptr(t&0xffffffff))
			args = append(args, uintptr((t>>32)&0xffffffff))
		case uint64:
			args = append(args, uintptr(t&0xffffffff))
			args = append(args, uintptr((t>>32)&0xffffffff))
		case bool:
			if t {
				args = append(args, uintptr(1))
			} else {
				args = append(args, uintptr(0))
			}
		case float32:
			u32 := math.Float32bits(t)
			args = append(args, uintptr(u32))
		case float64:
			kind := string(typeInfo[3+i])
			if kind == encFloat {
				u32 := math.Float32bits(float32(t))
				args = append(args, uintptr(u32))
			} else if kind == encDouble {
				u64 := math.Float64bits(t)
				args = append(args, uintptr(u64&0xffffffff))
				args = append(args, uintptr((u64>>32)&0xffffffff))
			}
		default:
			val := reflect.ValueOf(arg)
			switch val.Kind() {
			case reflect.Ptr:
				args = append(args, val.Pointer())
			case reflect.Uintptr:
				args = append(args, uintptr(val.Uint()))
			case reflect.Struct:
				structArgs := unpackStruct(val)
				args = append(args, structArgs...)
			default:
				panic("unhandled kind")
			}
		}
	}

	fc := variadic.NewFunctionCall(sendFuncName)
	if sendFuncName == "objc_msgSend" {
		fc.Words[0] = obj.Pointer()
	} else if sendFuncName == "objc_msgSendSuper" {
		superPtr := C.GoObjc_GetObjectSuperClassStruct(unsafe.Pointer(obj.Pointer()))
		defer C.free(superPtr)
		fc.Words[0] = uintptr(superPtr)
	} else {
		panic("objc: unknown object.sendMsg sendFuncName")
	}
	fc.Words[1] = uintptr(sel)

	for i, v := range args {
		fc.Words[i+2] = v
	}
	fc.NumArgs = 2 + len(args)

	if len(typeInfo) > 0 {
		retEnc := string(typeInfo[0])
		if retEnc == encFloat {
			f64 := float64(fc.CallFloat32())
			return object64{big: math.Float64bits(f64)}
		} else if retEnc == encDouble {
			f64 := fc.CallFloat64()
			return object64{big: math.Float64bits(f64)}
		}
	}

	return object{ptr: fc.Call()}
}

func (obj object) Send(selector string, args ...interface{}) Object {
	return sendMsg(obj, "objc_msgSend", selector, args...)
}

func (obj object) SendSuper(selector string, args ...interface{}) Object {
	return sendMsg(obj, "objc_msgSendSuper", selector, args...)
}
