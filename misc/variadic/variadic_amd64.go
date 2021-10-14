// Copyright 2012 Mikkel Krautz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package variadic

/*
#include <stdlib.h>
#include <dlfcn.h>

void *VariadicCall(void *ctx);
float VariadicCallFloat(void *ctx);
double VariadicCallDouble(void *ctx);

void *LookupSymAddr(char *str) {
	return dlsym(RTLD_DEFAULT, str);
}
*/
import "C"

import (
	"unsafe"
)

const (
	RDI = iota
	RDX
	RCX
	R8
	R9
	XMM0
	XMM1
	XMM2
	XMM3
	XMM4
	XMM5
	XMM6
	XMM7
)

type FunctionCall struct {
	Words     [14]uintptr
	NumFloat  int64
	NumMemory int64
	Memory    uintptr
	addr      unsafe.Pointer
}

// NewFunctionCall creates a new FunctionCall than can be
// used to call the C function named by the name parameter.
func NewFunctionCall(name string) *FunctionCall {
	fc := new(FunctionCall)
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	fc.addr = C.LookupSymAddr(cname)
	return fc
}

// NewFunctionCallAddr creates a new FunctionCall that can be
// used to cll the C function at the address given by the addr
// parameter.
func NewFunctionCallAddr(addr unsafe.Pointer) *FunctionCall {
	fc := new(FunctionCall)
	fc.addr = addr
	return fc
}

// Call calls the FunctionCall's underlying function, returning
// its return value as an uintptr.
func (f *FunctionCall) Call() uintptr {
	return uintptr(C.VariadicCall(unsafe.Pointer(f)))
}

// CallFloat32 calls the FunctionCall's underlying function, returning
// its return value as a float32.
func (f *FunctionCall) CallFloat32() float32 {
	return float32(C.VariadicCallFloat(unsafe.Pointer(f)))
}

// CallFloat64 calls the FunctionCall's underlying function, returning
// its return value as float64.
func (f *FunctionCall) CallFloat64() float64 {
	return float64(C.VariadicCallDouble(unsafe.Pointer(f)))
}
