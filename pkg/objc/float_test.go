// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"reflect"
	"testing"
	"unsafe"
)

func TestFloatArgsImplicit(t *testing.T) {
	expected := 50.0
	number := GetClass("NSNumber").Send("alloc").Send("initWithFloat:", expected)
	str := number.String()
	if str != "50" {
		t.Errorf("expected %v, got %v", expected, str)
	}
}

func TestDoubleArgsImplicit(t *testing.T) {
	expected := 51.0
	number := GetClass("NSNumber").Send("alloc").Send("initWithDouble:", expected)
	str := number.String()
	if str != "51" {
		t.Errorf("expected %v, got %v", expected, str)
	}
}

func TestFloatArgsExplicit(t *testing.T) {
	expected := float32(52.0)
	number := GetClass("NSNumber").Send("alloc").Send("initWithFloat:", expected)
	str := number.String()
	if str != "52" {
		t.Errorf("expected %v, got %v", expected, str)
	}
}

func TestDoubleArgsExplicit(t *testing.T) {
	expected := float64(53.0)
	number := GetClass("NSNumber").Send("alloc").Send("initWithDouble:", expected)
	str := number.String()
	if str != "53" {
		t.Errorf("expected %v, got %v", expected, str)
	}
}

func TestDoubleReturnValue(t *testing.T) {
	in := float64(54.0)
	out := GetClass("NSNumber").Send("alloc").Send("initWithDouble:", in).Send("doubleValue")
	if out.Float() != in {
		t.Errorf("expected %v, got %v", in, out.Float())
	}
}

func TestFloatReturnValue(t *testing.T) {
	in := float64(55.0)
	out := GetClass("NSNumber").Send("alloc").Send("initWithDouble:", in).Send("floatValue")
	if out.Float() != in {
		t.Errorf("expected %v, got %v", in, out.Float())
	}
}

type FloatTester struct {
	Object `objc:"FloatTester : NSObject"`
}

func (ft *FloatTester) Float64Returner() float64 {
	return 42.0
}

func (ft *FloatTester) Float32Returner() float32 {
	return 42.0
}

func TestFloat64RetGoObject(t *testing.T) {
	c := NewClassFromStruct(FloatTester{})
	c.AddMethod("float64Returner", (*FloatTester).Float64Returner)
	c.AddMethod("float32Returner", (*FloatTester).Float32Returner)
	RegisterClass(c)

	ft := GetClass("FloatTester").Send("alloc").Send("init")
	goFt := reflect.NewAt(reflect.TypeOf(FloatTester{}), unsafe.Pointer(object{ptr: ft.Pointer()}.internalPointer())).Interface().(*FloatTester)

	goAnswer64 := goFt.Float64Returner()
	objcAnswer64 := ft.Send("float64Returner").Float()
	if goAnswer64 != objcAnswer64 {
		t.Errorf("float64: expected %v, got %v", goAnswer64, objcAnswer64)
	}

	goAnswer32 := goFt.Float32Returner()
	objcAnswer32 := float32(ft.Send("float32Returner").Float())
	if goAnswer32 != objcAnswer32 {
		t.Errorf("float32: expected %v, got %v", goAnswer32, objcAnswer32)
	}
}
