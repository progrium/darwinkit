// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package ffi

import (
	"runtime"
	"testing"
	"time"
	"unsafe"

	"github.com/progrium/darwinkit/internal/assert"
)

func Test_closure(t *testing.T) {
	cif, status := PrepCIF(TypeSint64, []*Type{TypeSint64, TypeSint64})
	assert.Equal(t, OK, status)
	fn, handle, status := CreateClosure(cif, func(cif *CIF, ret unsafe.Pointer, args []unsafe.Pointer) {
		rp := (*int)(ret)
		*rp = 10
	})
	assert.Equal(t, OK, status)

	var ret int
	var arg1 int
	var arg2 int
	Call(cif, fn, unsafe.Pointer(&ret), []unsafe.Pointer{unsafe.Pointer(&arg1), unsafe.Pointer(&arg2)})
	assert.Equal(t, 10, ret)

	handle.Delete()
	runtime.GC()
	time.Sleep(time.Millisecond * 10)
	runtime.GC()
}
