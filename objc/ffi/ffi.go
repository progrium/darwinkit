// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ffi

// #cgo CFLAGS: -x objective-c -Wno-unguarded-availability-new
// #cgo LDFLAGS: -l ffi
// #import <ffi/ffi.h>
// #import <stdint.h>
import "C"
import (
	"runtime"
	"runtime/cgo"
	"unsafe"

	"github.com/ebitengine/purego"
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

var (
	FFITypeVoid uintptr
	FFITypeUint8 uintptr
	FFITypeSint8 uintptr
	FFITypeUint16 uintptr
	FFITypeSint16 uintptr
	FFITypeUint32 uintptr
	FFITypeSint32 uintptr
	FFITypeUint64 uintptr
	FFITypeSint64 uintptr
	FFITypeFloat uintptr
	FFITypeDouble uintptr
	FFITypePointer uintptr
)

// TODO size_t depends on architecture
type FFIType struct {
	size 	 int
	alignment uint16
	// **FFIType array
	elements unsafe.Pointer
}

type FFICif struct {
	abi  FFIABI
	nargs uint32
	argTypes []*FFIType
	rtype *FFIType
	bytes uint32
	flags uint32
}

// TODO intel won't work
// TODO naming
// Default alignment is 8, nothing needed here for
// __attribute__((aligned(8))) on ffi.h
type FFIClosure struct {
	trampoline_table uintptr
	trampoline_table_entry uintptr
	cif *FFICif
	fun func(FFICif, unsafe.Pointer, unsafe.Pointer, unsafe.Pointer)
	user_data uintptr
}

// TODO const naming
type FFIStatus uint32
const (
	FFIStatusOK = iota
	FFIStatusBadTypedef
	FFIStatusBadABI
	FFIStatusBadArgType
)

type FFIABI uint32
const (
	FFIFirstABI = iota
	FFISysV
	FFIWin64
	FFILastABI
	FFIDefaultABI = FFISysV
)

var libffi uintptr

// TODO passing array pointer to C
// TODO usage comments not correct
var (
	// FFICall: FFI call function
	// @param cif: CIF pointer
	// @param fn: function pointer
	// @param rvalue: return value
	// @param avalues: arguments
	// @return: void
	FFICall func(cif unsafe.Pointer, fn unsafe.Pointer, rvalue unsafe.Pointer, avalues []unsafe.Pointer)

	// TODO atypes argument
	// FFIPrepCIF: FFI prepare CIF function
	// @param cif: CIF pointer
	// @param abi: ABI
	// @param nargs: number of arguments
	// @param rtype: return type
	// @param atypes: argument types
	// @return: status
	FFIPrepCIF func(cif unsafe.Pointer, abi FFIABI, nargs uint32, rtype unsafe.Pointer, atypes unsafe.Pointer) FFIStatus

	// FFIClosureAlloc: FFI closure alloc function
	// @param size: size
	// @param code: code pointer
	// @return: closure pointer
	FFIClosureAlloc func(size uint32, code unsafe.Pointer) unsafe.Pointer

	// TODO update typed pointers and function pointers
	// FFIPrepClosureLoc: FFI prepare closure loc function
	// @param closure: closure pointer
	// @param cif: CIF pointer
	// @param fun: function pointer
	// @param user_data: user data
	// @param codeloc: code location
	// @return: status
	FFIPrepClosureLoc func(closure unsafe.Pointer, cif unsafe.Pointer, fun uintptr, user_data unsafe.Pointer, codeloc unsafe.Pointer) FFIStatus

	// FFIClosureFree: FFI closure free function
	// @param closure: closure pointer
	// @return: void
	FFIClosureFree func(closure unsafe.Pointer)
)

func LoadFFI() {
	libffi, err := purego.Dlopen("libffi.dylib", purego.RTLD_NOW | purego.RTLD_GLOBAL)
	if err != nil {
		panic(err)
	}

	purego.RegisterLibFunc(&FFICall, libffi, "ffi_call")
	purego.RegisterLibFunc(&FFIPrepCIF, libffi, "ffi_prep_cif")
	purego.RegisterLibFunc(&FFIClosureAlloc, libffi, "ffi_closure_alloc")
	purego.RegisterLibFunc(&FFIPrepClosureLoc, libffi, "ffi_prep_closure_loc")
	purego.RegisterLibFunc(&FFIClosureFree, libffi, "ffi_closure_free")

	loadExternSymbol := func (symbol string) uintptr {
		ptr, err := purego.Dlsym(libffi, symbol)
		if err != nil {
			panic(err)
		}
		return ptr
	}

	FFITypeVoid = loadExternSymbol("ffi_type_void")
	FFITypeUint8 = loadExternSymbol("ffi_type_uint8")
	FFITypeSint8 = loadExternSymbol("ffi_type_sint8")
	FFITypeUint16 = loadExternSymbol("ffi_type_uint16")
	FFITypeSint16 = loadExternSymbol("ffi_type_sint16")
	FFITypeUint32 = loadExternSymbol("ffi_type_uint32")
	FFITypeSint32 = loadExternSymbol("ffi_type_sint32")
	FFITypeUint64 = loadExternSymbol("ffi_type_uint64")
	FFITypeSint64 = loadExternSymbol("ffi_type_sint64")
	FFITypeFloat = loadExternSymbol("ffi_type_float")
	FFITypeDouble = loadExternSymbol("ffi_type_double")
	FFITypePointer = loadExternSymbol("ffi_type_pointer")
}

// TODO move this to an appropriate place
func init() {
	LoadFFI()
}

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

func PrepCIF(rtype *Type, argtypes []*Type) (*CIF, FFIStatus) {
	var cif CIF
	s := FFIPrepCIF(unsafe.Pointer(&cif), FFIDefaultABI, uint32(len(argtypes)), unsafe.Pointer(rtype), unsafe.Pointer(&argtypes[0]))
	runtime.KeepAlive(rtype)
	runtime.KeepAlive(argtypes)
	return &cif, s
}

func Call(cif *CIF, fn unsafe.Pointer, rvalue unsafe.Pointer, avalues []unsafe.Pointer) {
	FFICall(unsafe.Pointer(cif), fn, rvalue, avalues)
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

func CreateClosure(cif *CIF, f ClosureHandle) (codeloc unsafe.Pointer, udHandle cgo.Handle, status FFIStatus) {
	closure := FFIClosureAlloc(uint32(unsafe.Sizeof(FFIClosure{})), unsafe.Pointer(&codeloc))
	guard := new(int)
	userData := UserData{
		cif:    cif,
		handle: f,
		guard:  guard,
	}
	runtime.KeepAlive(userData)
	runtime.SetFinalizer(guard, func(v *int) {
		FFIClosureFree(closure)
	})
	// keep this for now to not break the api
	// but this is no longer needed
	udHandle = cgo.NewHandle(userData)
	status = FFIPrepClosureLoc(closure, unsafe.Pointer(cif), purego.NewCallback(handleClosure), unsafe.Pointer(&userData), codeloc)
	return
}

func handleClosure(cif *CIF, ret unsafe.Pointer, args unsafe.Pointer, userData unsafe.Pointer) {
	userDataVal := *(*UserData)(userData)
	argsNum := int(cif.nargs)
	argS := unsafe.Slice((*unsafe.Pointer)(args), argsNum)
	userDataVal.handle(cif, ret, argS)
}
