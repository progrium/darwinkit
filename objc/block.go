// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

//#import <stdlib.h>
//#import <stdint.h>
//
// void* block_copy(void *ptr);
// void block_release(void *ptr);
// void* block_get_invoke(void* block);
// void* block_create_global(const char* signature, void* callable);
// void* block_create_malloc(const char* signature, void* callable, uintptr_t handle);
// void block_free(void* _block);
// void* testBlock();
import "C"
import (
	"reflect"
	"runtime"
	"runtime/cgo"
	"strings"
	"unsafe"

	"github.com/progrium/darwinkit/objc/ffi"
)

type Block struct {
	ptr unsafe.Pointer
}

func BlockFrom(ptr unsafe.Pointer) Block {
	return Block{ptr: ptr}
}

func (b Block) Ptr() unsafe.Pointer {
	return b.ptr
}

func (b Block) Copy() Block {
	return Block{C.block_copy(b.Ptr())}
}

func (b Block) Release() {
	C.block_release(b.Ptr())
}

func wrapBlockInGoFunc(bp unsafe.Pointer, funcType reflect.Type) reflect.Value {
	if bp == nil {
		panic("block is nil")
	}
	if funcType.NumOut() > 1 {
		panic("too many return values")
	}

	b := BlockFrom(bp)
	b = b.Copy()
	var sentinel = new(int)
	fv := reflect.MakeFunc(funcType, func(args []reflect.Value) (results []reflect.Value) {
		*sentinel = 1
		if funcType.NumOut() == 0 {
			callBlock(b, args, voidType)
			return nil
		} else {
			rv := callBlock(b, args, funcType.Out(0))
			return []reflect.Value{rv}
		}
	})
	runtime.SetFinalizer(sentinel, func(v *int) {
		b.Release()
	})
	return fv
}

func CallBlock[T any](b Block, params ...any) T {
	paramsReflect := make([]reflect.Value, len(params))
	for i := 0; i < len(params); i++ {
		paramsReflect[i] = reflect.ValueOf(params[i])
	}
	var ret T
	rt := reflect.TypeOf(ret)
	v := callBlock(b, paramsReflect, rt)
	return v.Interface().(T)
}

func callBlock(b Block, params []reflect.Value, rt reflect.Type) reflect.Value {
	ptr := b.Ptr()
	if ptr == nil {
		panic("object is nil")
	}

	argc := len(params)
	var args = make([]unsafe.Pointer, argc+1)
	var argTypes = make([]*ffi.Type, argc+1)
	args[0] = unsafe.Pointer(&ptr)
	argTypes[0] = ffi.TypePointer
	for i := 0; i < argc; i++ {
		args[i+1] = convertToObjcValue(params[i])
		if !params[i].IsValid() {
			argTypes[i+1] = ffi.TypePointer
		} else {
			argTypes[i+1] = toFFIType(params[i].Type())
		}
	}

	var retPtr unsafe.Pointer
	if rt.Kind() == reflect.Struct {
		size := rt.Size()
		if size == 0 {
			var i int64
			retPtr = unsafe.Pointer(&i)
		} else {
			buffer := make([]byte, rt.Size())
			retPtr = unsafe.Pointer(&buffer[0])
		}
	} else {
		var i int64
		retPtr = unsafe.Pointer(&i)
	}
	retType := toFFIType(rt)

	fn := C.block_get_invoke(ptr)
	if fn == nil {
		panic("block imp is nil")
	}

	cif, status := ffi.PrepCIF(retType, argTypes)
	if status != ffi.OK {
		panic("ffi prep cif status not ok")
	}
	ffi.Call(cif, fn, retPtr, args)

	return convertToGoValue(retPtr, rt)
}

// CreateMallocBlock wraps a Go function as an autoreleased Block.
func CreateMallocBlock(f any) Block {
	rf := reflect.ValueOf(f)
	ft := rf.Type()
	if ft.Kind() != reflect.Func {
		panic("not func type")
	}

	typeEncoding := getBlockTypeEncoding(ft)
	cte := C.CString(typeEncoding)
	imp, handle := wrapGoFuncAsBlockIMP(rf)
	bp := C.block_create_malloc(cte, imp.ptr, C.uintptr_t(handle))
	b := BlockFrom(bp)
	ObjectFrom(bp).Autorelease()
	return b
}

// CreateMallocBlock wraps a Go function as a global Block.
func CreateGlobalBlock(f any) Block {
	rf := reflect.ValueOf(f)
	ft := rf.Type()
	if ft.Kind() != reflect.Func {
		panic("not func type")
	}

	typeEncoding := getBlockTypeEncoding(ft)
	cte := C.CString(typeEncoding) // // cte will be freed by oc_dispose_helper
	imp, _ := wrapGoFuncAsBlockIMP(rf)
	bp := C.block_create_global(cte, imp.ptr)
	b := BlockFrom(bp)
	return b
}

func wrapGoFuncAsBlockIMP(rf reflect.Value) (imp IMP, handle cgo.Handle) {
	if rf.Kind() != reflect.Func {
		panic("f should be a func")
	}
	rt := rf.Type()
	if rt.NumOut() > 1 {
		panic("too many return value")
	}

	goArgc := rf.Type().NumIn()
	var objcArgTypes = make([]*ffi.Type, goArgc+1)
	objcArgTypes[0] = ffi.TypePointer // block pointer
	for i := 0; i < goArgc; i++ {
		objcArgTypes[i+1] = toFFIType(rt.In(i))
	}

	var retType *ffi.Type
	if rt.NumOut() == 0 {
		retType = ffi.TypeVoid
	} else {
		retType = toFFIType(rt.Out(0))
	}

	cif, status := ffi.PrepCIF(retType, objcArgTypes)
	if status != ffi.OK {
		panic("ffi prep cif status not ok")
	}

	fn, handle, status := ffi.CreateClosure(cif, func(cif *ffi.CIF, ret unsafe.Pointer, objcArgs []unsafe.Pointer) {
		var goArgs = make([]reflect.Value, len(objcArgs)-1)
		for i := 0; i < len(goArgs); i++ {
			goArgs[i] = convertToGoValue(objcArgs[i+1], rt.In(i))
		}
		results := rf.Call(goArgs)
		if len(results) == 1 {
			setGoValueToObjcPointer(results[0], ret)
		}
	})
	if status != ffi.OK {
		panic("ffi prep closure status not ok")
	}

	return IMPFrom(fn), handle
}

func getBlockTypeEncoding(ft reflect.Type) string {
	if ft.Kind() != reflect.Func {
		panic("not func type")
	}
	if ft.NumOut() > 1 {
		panic("to many return values")
	}
	var sb strings.Builder
	if ft.NumOut() == 0 {
		sb.WriteByte('v')
	} else {
		sb.WriteString(getTypeEncoding(ft.Out(0)))
	}
	sb.WriteString("@?") // block self as first parameter
	for i := 0; i < ft.NumIn(); i++ {
		sb.WriteString(getTypeEncoding(ft.In(i)))
	}
	return sb.String()
}

func testBlock() Block {
	return BlockFrom(C.testBlock())
}

func toUIntptr[T any](ptr *T) C.uintptr_t {
	return C.uintptr_t(uintptr(unsafe.Pointer(ptr)))
}
