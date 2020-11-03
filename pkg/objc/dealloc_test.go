// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import "testing"

type CustomDealloc struct {
	Object `objc:"CustomDealloc : NSObject"`
}

func (cd *CustomDealloc) Dealloc() {
	cd.SendSuperMsg("dealloc")
}

func TestGoObjectCustomDealloc(t *testing.T) {
	c := NewClass(CustomDealloc{})
	c.AddMethod("dealloc", (*CustomDealloc).Dealloc)
	RegisterClass(c)

	dc := GetClass("CustomDealloc").SendMsg("alloc").SendMsg("init")
	dc.SendMsg("dealloc")
}

type StdDealloc struct {
	Object `objc:"StdDealloc : NSObject"`
}

func TestGoObjectStandardDealloc(t *testing.T) {
	c := NewClass(StdDealloc{})
	RegisterClass(c)

	dc := GetClass("StdDealloc").SendMsg("alloc").SendMsg("init")
	dc.SendMsg("dealloc")
}
