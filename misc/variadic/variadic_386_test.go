// Copyright 2013 Mikkel Krautz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package variadic

import (
	"testing"
	"unsafe"
	"reflect"
	"math"
)

func TestIntArg0(t *testing.T) {
	fc := new(FunctionCall)
	fc.addr = arg0fn()
	fc.Words[0] = 0xcafebabe
	fc.NumArgs = 1
	ret := fc.Call()
	if ret != fc.Words[0] {
		t.Fatalf("got 0x%x; expected 0x%x", ret, fc.Words[0])
	}
}

func TestArgN(t *testing.T) {
	for n := 1; n < 10; n++ {
		fc := new(FunctionCall)
		fc.addr = argNfn()
		for i := 0; i < len(fc.Words); i++ {
			fc.Words[i] = 0
		}
		fc.Words[0] = uintptr(n)
		fc.Words[n] = uintptr(0xdead + n)
		fc.NumArgs = n+1
		ret := fc.Call()
		if ret != fc.Words[n] {
			t.Fatalf("got 0x%x; expected 0x%x", ret, fc.Words[n])
		}
	}
}

func TestFloat32Ret(t *testing.T) {
	fc := new(FunctionCall)
	fc.addr = floatret()
	fc.Words[0] = 0xf00
	fc.NumArgs = 1
	ret := fc.CallFloat32()
	expected := float32(3.141592)
	if ret != expected {
		t.Fatalf("got %v; expected %v", ret, expected)
	}
}

func TestFloat64Ret(t *testing.T) {
	fc := new(FunctionCall)
	fc.addr = doubleret()
	fc.Words[0] = 0xf00
	fc.NumArgs = 1
	ret := fc.CallFloat64()
	expected := float64(3.141592)
	if ret != expected {
		t.Fatalf("got %v; expected %v", ret, expected)
	}
}

func TestPrintf(t *testing.T) {
	buf := make([]byte, 256)
	fmt := "%i 0x%x %u\n\x00"
	fc := NewFunctionCall("sprintf")
	slicehdr := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	strhdr := (*reflect.StringHeader)(unsafe.Pointer(&fmt))
	fc.Words[0] = slicehdr.Data
	fc.Words[1] = strhdr.Data
	fc.Words[2] = 1
	fc.Words[3] = 0xcafebabe
	fc.Words[4] = 0xff
	fc.NumArgs = 5
	n := fc.Call()
	if n > 256 {
		t.Fatal("bad return value from sprintf")
	}
	val := string(buf[0:int(n)])
	expected := "1 0xcafebabe 255\n"
	if val != expected {
		t.Fatalf("got %v; expected %v", val, expected)
	}
}

func TestFloatPrintf(t *testing.T) {
	buf := make([]byte, 256)
	fmt := "%.6f %i\n\x00"
	fc := NewFunctionCall("sprintf")
	slicehdr := (*reflect.SliceHeader)(unsafe.Pointer(&buf))
	strhdr := (*reflect.StringHeader)(unsafe.Pointer(&fmt))
	fc.Words[0] = slicehdr.Data
	fc.Words[1] = strhdr.Data
	u64 := math.Float64bits(3.141592)
	fc.Words[2] = uintptr(u64 & 0xffffffff)
	fc.Words[3] = uintptr((u64 >> 32) & 0xffffffff)
	fc.Words[4] = 2
	fc.NumArgs = 5
	n := fc.Call()
	if n > 256 {
		t.Fatal("bad return value from sprintf")
	}
	val := string(buf[0:int(n)])
	expected := "3.141592 2\n"
	if val != expected {
		t.Fatalf("got %v; expected %v", val, expected)
	}
}