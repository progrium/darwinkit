// Copyright 2022 Mikkel Krautz. All rights reserved.
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
	X0 = iota
	X1
	X2
	X3
	X4
	X5
	X6
	X7
	Q0
	Q1
	Q2
	Q3
	Q4
	Q5
	Q6
	Q7
)

type uint128 struct {
	Upper uintptr
	Lower uintptr
}

type FunctionCall struct {
	Words     [8]uintptr
	Simd      [8]uint128
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
	C.VariadicCall(unsafe.Pointer(f))
	return f.Words[0]
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
