// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"reflect"
	"testing"

	"github.com/progrium/darwinkit/internal/assert"
)

func Test_getBlockTypeEncoding(t *testing.T) {
	var f = func(index uint, value Object) {

	}
	encoding := getBlockTypeEncoding(reflect.TypeOf(f))
	if encoding != "v@?Q@" {
		t.Fail()
	}
}

func Test_CallBlock(t *testing.T) {
	b := testBlock()
	v := CallBlock[int32](b, 10)
	assert.Equal(t, int32(5), v)
}

func Test_CreateGlobalBlock(t *testing.T) {
	b := CreateGlobalBlock(func(v int) int {
		return v / 2
	})
	v := CallBlock[int](b, 10)
	assert.Equal(t, 5, v)

	b2 := b.Copy()
	assert.Equal(t, b2, b)
	b.Release()
	b.Release()
	b.Release()
	v = CallBlock[int](b, 10)
	assert.Equal(t, 5, v)
}

func Test_CreateMallocBlock(t *testing.T) {
	b := CreateMallocBlock(func(v int) int {
		return v / 2
	})
	v := CallBlock[int](b, 10)
	assert.Equal(t, 5, v)

	b.Release()
	assert.Panics(t, func() {
		v = CallBlock[int](b, 10)
	})
}
