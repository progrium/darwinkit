// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ffi

// #cgo CFLAGS: -x objective-c -Wno-unguarded-availability-new
// #cgo LDFLAGS: -l ffi
// #import <ffi/ffi.h>
// #import <stdint.h>
// ffi_status ffi_prep_cif0(uintptr_t cif, ffi_abi abi, unsigned int nargs, uintptr_t rtype, uintptr_t atypes);
// void ffi_call0(uintptr_t cif, void* fn, uintptr_t rvalue, uintptr_t avalue);
// void *ffi_closure_alloc0(uintptr_t code);
// ffi_status ffi_prep_closure_loc0(void* closure, uintptr_t cif, void* fn, uintptr_t user_data, void *codeloc);
//
// void forward_to_go(ffi_cif *cif, void *ret, void* args[], void * user_data);
//
import "C"
import (
	"runtime"
	"runtime/cgo"
	"unsafe"
)

type Type = C.ffi_type
type CIF = C.ffi_cif
type Arg = C.ffi_arg
type Status = C.ffi_status
type ABI = C.ffi_abi
type Closure C.ffi_closure

var OK Status = C.FFI_OK
var BAD_TYPEDEF Status = C.FFI_BAD_TYPEDEF
var BAD_ABI Status = C.FFI_BAD_ABI

var DEFAULT_ABI ABI = C.FFI_DEFAULT_ABI

var TypeVoid *Type = &C.ffi_type_void
var TypeUint8 *Type = &C.ffi_type_uint8
var TypeSint8 *Type = &C.ffi_type_sint8
var TypeUint16 *Type = &C.ffi_type_uint16
var TypeSint16 *Type = &C.ffi_type_sint16
var TypeUint32 *Type = &C.ffi_type_uint32
var TypeSint32 *Type = &C.ffi_type_sint32
var TypeUint64 *Type = &C.ffi_type_uint64
var TypeSint64 *Type = &C.ffi_type_sint64
var TypeFloat *Type = &C.ffi_type_float
var TypeDouble *Type = &C.ffi_type_double
var TypePointer *Type = &C.ffi_type_pointer

func IsStruct(t *Type) bool {
	return t._type == C.FFI_TYPE_STRUCT
}

func MakeStructType(types []*Type) *Type {
	nullTerminated := make([]*Type, len(types)+1)
	copy(nullTerminated, types)
	nullTerminated[len(types)] = nil

	return &Type{
		_type:    C.FFI_TYPE_STRUCT,
		elements: &nullTerminated[0],
	}
}

func PrepCIF(rtype *Type, argtypes []*Type) (*CIF, Status) {
	var cif CIF
	s := C.ffi_prep_cif0(toUintptrT(&cif), DEFAULT_ABI, C.uint(len(argtypes)), toUintptrT(rtype), toUintptrT(&argtypes[0]))
	runtime.KeepAlive(rtype)
	runtime.KeepAlive(argtypes)
	return &cif, s
}

func Call(cif *CIF, fn unsafe.Pointer, rvalue unsafe.Pointer, avalues []unsafe.Pointer) {
	C.ffi_call0(toUintptrT(cif), fn, C.uintptr_t(uintptr(rvalue)), toUintptrT(&avalues[0]))
	runtime.KeepAlive(cif)
	runtime.KeepAlive(avalues)
	runtime.KeepAlive(rvalue)
}

func toUintptrT[T any](p *T) C.uintptr_t {
	return C.uintptr_t(uintptr(unsafe.Pointer(p)))
}

type ClosureHandle func(cif *CIF, ret unsafe.Pointer, args []unsafe.Pointer)

type UserData struct {
	cif    *CIF          // keep CIF reference, must be kept alive until the closure itself is freed.
	handle ClosureHandle // the closure handle func
	guard  *int          // used to free resource when is gced
}

func CreateClosure(cif *CIF, f ClosureHandle) (codeloc unsafe.Pointer, udHandle cgo.Handle, status Status) {
	closure := C.ffi_closure_alloc0(toUintptrT(&codeloc))
	guard := new(int)
	userData := UserData{
		cif:    cif,
		handle: f,
		guard:  guard,
	}
	runtime.SetFinalizer(guard, func(v *int) {
		C.ffi_closure_free(closure)
	})
	udHandle = cgo.NewHandle(userData)
	status = C.ffi_prep_closure_loc0(closure, toUintptrT(cif), C.forward_to_go, C.uintptr_t(udHandle), codeloc)
	return
}

//export handleClosure
func handleClosure(cif *CIF, ret unsafe.Pointer, args unsafe.Pointer, userData unsafe.Pointer) {
	goUserData := cgo.Handle(userData).Value().(UserData)
	argsNum := int(cif.nargs)
	argS := unsafe.Slice((*unsafe.Pointer)(args), argsNum)
	goUserData.handle(cif, ret, argS)
}
