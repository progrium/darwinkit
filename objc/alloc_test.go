// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"testing"
)

type GoStruct struct {
	Object `objc:"GoStruct : NSObject"`
	value  uint64
}

func (gs *GoStruct) SetValue() {
	gs.value = 0xbadc0c0a
}

func (gs *GoStruct) HasValue() bool {
	return gs.value == 0xbadc0c0a
}

func TestAlloc(t *testing.T) {
	c := NewClassFromStruct(GoStruct{})
	c.AddMethod("setValue", (*GoStruct).SetValue)
	c.AddMethod("hasValue", (*GoStruct).HasValue)
	RegisterClass(c)

	o := GetClass("GoStruct").Send("alloc").Send("init")
	o.Send("setValue")
	if !o.Send("hasValue").Bool() {
		t.Errorf("value was not set!")
	}
}
