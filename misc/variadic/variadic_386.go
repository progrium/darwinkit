// Copyright 2013 Mikkel Krautz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package variadic

/*
#cgo LDFLAGS: -ldl

#include <dlfcn.h>
#include <stdlib.h>

void *VariadicCall(void *ctx);
float VariadicCallFloat(void *ctx);
double VariadicCallDouble(void *ctx);

void *LookupSymAddr(char *str) {
#ifdef RTLD_DEFAULT
	return dlsym(RTLD_DEFAULT, str);
#else
	return dlsym(NULL, str);
#endif
}
*/
import "C"

import "unsafe"

type FunctionCall struct {
	Words      [12]uintptr
	NumArgs    int
	addr       unsafe.Pointer
}

// NewFunctionCall creates a new FunctionCall than can be
// used to call the C function named by the name parameter.
func NewFunctionCall(name string) *FunctionCall {
	fc := new(FunctionCall)
	fc.addr = C.LookupSymAddr(C.CString(name))
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
	// Ensure stack is 16 byte aligned.
	rem := (f.NumArgs % 4)
	if rem > 0 {
		f.NumArgs = f.NumArgs + (4 - rem)
	}

	if f.NumArgs > len(f.Words) {
		panic("bad NumArgs")
	}
	
	if f.addr == nil {
		panic("variadic: Call called with nil function addr")
	}
	return uintptr(C.VariadicCall(unsafe.Pointer(f)))
}

// CallFloat32 calls the FunctionCall's underlying function, returning
// its return value as a float32.
func (f *FunctionCall) CallFloat32() float32 {
	if f.NumArgs > len(f.Words) {
		panic("bad NumArgs")
	}
	if f.addr == nil {
		panic("variadic: CallFloat32 called with nil function addr")
	}
	return float32(C.VariadicCallFloat(unsafe.Pointer(f)))
}

// CallFloat64 calls the FunctionCall's underlying function, returning
// its return value as float64.
func (f *FunctionCall) CallFloat64() float64 {
	if f.NumArgs > len(f.Words) {
		panic("bad NumArgs")
	}
	if f.addr == nil {
		panic("variadic: CallFloat64 called with nil function addr")
	}
	return float64(C.VariadicCallDouble(unsafe.Pointer(f)))
}
