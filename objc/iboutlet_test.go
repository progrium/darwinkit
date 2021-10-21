// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"reflect"
	"strings"
	"sync"
	"testing"
	"unsafe"
)

var iboutletOnce sync.Once

func registerIBOutletTestClass() {
	iboutletOnce.Do(func() {
		c := NewClassFromStruct(IBOutletTester{})
		c.AddMethod("myselfIsNil", (*IBOutletTester).MyselfIsNil)
		c.AddMethod("myselfIsMyself", (*IBOutletTester).MyselfIsMyself)
		RegisterClass(c)
	})
}

type IBOutletTester struct {
	Object `objc:"IBOutletTester : NSObject"`
	Myself Object `objc:"IBOutlet"`
}

func (ibo *IBOutletTester) MyselfIsNil() bool {
	return ibo.Myself == nil
}

func (ibo *IBOutletTester) MyselfIsMyself() bool {
	return ibo.Myself.Pointer() == ibo.Object.Pointer()
}

const NSUTF8StringEncoding = 4

func NSStringFromString(str string) Object {
	hdrp := (*reflect.StringHeader)(unsafe.Pointer(&str))
	return GetClass("NSString").Send("alloc").Send("initWithBytes:length:encoding:", hdrp.Data, hdrp.Len, NSUTF8StringEncoding)
}

func TestIBOutletKeyValueCodingImpl(t *testing.T) {
	registerIBOutletTestClass()

	Autorelease(func() {
		ibo := GetClass("IBOutletTester").Send("alloc").Send("init")
		ibo.Send("setValue:forKey:", ibo, NSStringFromString("Myself").Autorelease())

		if ibo.Send("myselfIsNil").Bool() {
			t.Fatal("nil iboutlet, value not properly set for key")
		}

		if !ibo.Send("myselfIsMyself").Bool() {
			t.Fatal("value not set, or incorrectly set.")
		}
	})
}

func TestIBOutletSetter(t *testing.T) {
	registerIBOutletTestClass()

	Autorelease(func() {
		ibo := GetClass("IBOutletTester").Send("alloc").Send("init")
		ibo.Send("setMyself:", ibo)

		if ibo.Send("myselfIsNil").Bool() {
			t.Fatal("nil iboutlet, value not properly set for key")
		}

		if !ibo.Send("myselfIsMyself").Bool() {
			t.Fatal("value not set, or incorrectly set.")
		}
	})
}

type Shadow struct {
	Object `objc:"Shadow : NSObject"`
	Number Object `objc:"IBOutlet"`
}

func (s *Shadow) SetNumber(magicNumber Object) {
	s.Number = magicNumber.Retain()
}

func TestIBOutletShadowPanic(t *testing.T) {
	defer func() {
		str := recover()
		if str == nil {
			t.Fatal("expected panic")
		}
		msg, ok := str.(string)
		if !ok {
			t.Fatal("expectd string return from recover()")
		}
		if !strings.Contains(msg, "would shadow IBOutlet") {
			t.Fatal("expected shadow panic")
		}
	}()

	c := NewClassFromStruct(Shadow{})
	c.AddMethod("setNumber:", (*Shadow).SetNumber)
	RegisterClass(c)
}
