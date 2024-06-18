// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"testing"

	"github.com/progrium/darwinkit/internal/assert"
)

func TestGetClass(t *testing.T) {
	cls := GetClass("NSObject")

	version := cls.Version()
	if version != 0 {
		t.Failed()
	}

	if cls.Name() != "NSObject" {
		t.Failed()
	}
}

func TestGetMethod(t *testing.T) {
	cls := GetClass("NSObject")

	m := cls.ClassMethod(RegisterSelectorName("alloc"))
	name := m.Name().Name()
	if name != "alloc" {
		t.Failed()
	}
	te := m.TypeEncoding()
	if te != "@16@0:8" {
		t.Failed()
	}
}

func TestClass_CopyMethodList(t *testing.T) {
	cls := GetClass("NSObject")
	ms := cls.CopyMethodList()
	for _, m := range ms {
		if m.ptr == nil {
			t.Fail()
		}
	}
}

func TestClass_CopyPropertyList(t *testing.T) {
	oc := GetClass("NSObject")
	ps := oc.CopyPropertyList()
	for _, p := range ps {
		_ = p
		// fmt.Println(p.GetAttributes())
		// fmt.Println(p.CopyAttributeList())
	}
}

func Test_CallMethod(t *testing.T) {
	// call method
	var o = NewObject()
	var count = Call[uint](o, Sel("retainCount"))
	assert.Equal(t, uint(1), count)

}

func Test_AddMethod(t *testing.T) {
	class := AllocateClass(GetClass("NSObject"), "MyClass1", 0)
	var o = class.CreateInstance(0)
	sel := Sel("plus:and:")
	ok := AddMethod(class, sel, func(o Object, v1 int, v2 int) int {
		return v1 + v2
	})
	assert.True(t, ok)
	ret := Call[int](o, sel, 1, 2)
	assert.Equal(t, 3, ret)

	//replace method
	ReplaceMethod(o.Class(), sel, func(o Object, v1 int, v2 int) int {
		return v1 + v2 + 10
	})
	ret = Call[int](o, sel, 1, 2)
	assert.Equal(t, 13, ret)
}

func Test_AddClassMethod(t *testing.T) {
	class := AllocateClass(GetClass("NSObject"), "MyClass2", 0)
	//TODO: why need call class.GetMethodImplementation to make meta class exists?
	o := class.CreateInstance(0)
	o.RetainCount()
	sel := Sel("plus:and:")
	ok := AddClassMethod(class, sel, func(c IClass, v1 int, v2 int) int {
		return v1 + v2
	})
	assert.True(t, ok)
	ret := Call[int](class, sel, 4, 5)
	assert.Equal(t, 9, ret)

	// replace method
	ReplaceClassMethod(class, sel, func(o Object, v1 int, v2 int) int {
		return v1 + v2 + 10
	})
	ret = Call[int](class, sel, 4, 5)
	assert.Equal(t, 19, ret)
}
