// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"reflect"
	"testing"
	"unsafe"
)

type SomeGoObject struct {
	Object `objc:"SomeGoObject : NSObject"`
}

func (sgo *SomeGoObject) Dealloc() {
}

func (sgo *SomeGoObject) FiftyFive() int64 {
	return 55
}

func (sgo *SomeGoObject) ThirtyTwo() int64 {
	return 32
}

func (sgo *SomeGoObject) Sum() int64 {
	return sgo.Send("thirtyTwo").Int() + sgo.Send("fiftyFive").Int()
}

func (sgo *SomeGoObject) GoSum() int64 {
	return sgo.ThirtyTwo() + sgo.FiftyFive()
}

func TestGoObjectCallObjC(t *testing.T) {
	c := NewClassFromStruct(SomeGoObject{})
	c.AddMethod("fiftyFive", (*SomeGoObject).FiftyFive)
	c.AddMethod("thirtyTwo", (*SomeGoObject).ThirtyTwo)
	c.AddMethod("goSum", (*SomeGoObject).GoSum)
	c.AddMethod("sum", (*SomeGoObject).Sum)
	RegisterClass(c)

	sgo := Get("SomeGoObject").Send("alloc").Send("init")
	if sgo.Send("goSum").Int() != sgo.Send("sum").Int() {
		t.Errorf("calculated sums do not match")
	}
}

func TestObjectClass(t *testing.T) {
	// using core.String() would introduce an import cycle
	gostr := "hello world"
	hdrp := (*reflect.StringHeader)(unsafe.Pointer(&gostr))
	obj := Get("NSString").Alloc().Send("initWithBytes:length:encoding:", hdrp.Data, hdrp.Len, NSUTF8StringEncoding)

	cls := obj.Class()
	// NSString class method chosen at random
	r := cls.Send("defaultCStringEncoding").Int()
	if r == 0 {
		t.Errorf("unexpected default string encoding")
	}

	if !cls.Class().Equals(cls) {
		// the metaclass is returned by calling object_Get() on the class
		// object, not by sending -[x class]
		t.Errorf("sending -[x class] to a class object should return itself")
	}
}
