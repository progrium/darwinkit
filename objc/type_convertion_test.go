// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"reflect"
	"testing"

	"unsafe"

	"github.com/progrium/macdriver/internal/assert"
)

// for more array/dict tests, see foundation/foundation_custom_test.go

type TestStruct struct {
	A int64
	B int64
}

type TestDirectStruct struct {
	Ptr unsafe.Pointer
}

func Test_convertStruct(t *testing.T) {
	s := TestStruct{
		A: 10,
		B: 20,
	}
	p := convertToObjcValue(reflect.ValueOf(s))
	ns := *(*TestStruct)(p)
	assert.Equal(t, s, ns)

	ns2 := convertToGoValue(p, reflect.TypeOf(TestStruct{})).Interface().(TestStruct)
	assert.Equal(t, s, ns2)

	ds := TestDirectStruct{
		Ptr: p,
	}
	dp := convertToObjcValue(reflect.ValueOf(ds))
	nds := *(*TestDirectStruct)(dp)
	assert.Equal(t, ds, nds)

	nds2 := convertToGoValue(dp, reflect.TypeOf(TestDirectStruct{})).Interface().(TestDirectStruct)
	assert.Equal(t, ds, nds2)
}

func Test_setGoValueToObjcPointer(t *testing.T) {
	s := TestStruct{
		A: 10,
		B: 20,
	}
	rv := reflect.ValueOf(s)
	var ns TestStruct
	setGoValueToObjcPointer(rv, unsafe.Pointer(&ns))
	assert.Equal(t, s, ns)

	ds := TestDirectStruct{
		Ptr: unsafe.Pointer(&ns),
	}
	drv := reflect.ValueOf(ds)
	var nds TestDirectStruct
	setGoValueToObjcPointer(drv, unsafe.Pointer(&nds))
	assert.Equal(t, ds, nds)
}
