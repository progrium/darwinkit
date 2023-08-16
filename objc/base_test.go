// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"testing"
	"unsafe"
)

func TestSetAssociatedObject(t *testing.T) {
	o1 := NewObject()
	o2 := NewObject()
	var i = 0
	key := unsafe.Pointer(&i)
	SetAssociatedObject(o1, key, o2, ASSOCIATION_RETAIN)
	if o2.RetainCount() != 2 {
		t.Fail()
	}
	o1.Release()
	if o2.RetainCount() != 1 {
		t.Fail()
	}
	o2.Release()

	o1 = NewObject()
	o2 = NewObject()
	o3 := NewObject()
	SetAssociatedObject(o1, key, o2, ASSOCIATION_RETAIN)
	if o2.RetainCount() != 2 {
		t.Fail()
	}
	SetAssociatedObject(o1, key, o3, ASSOCIATION_RETAIN)
	if o2.RetainCount() != 1 {
		t.Fail()
	}
	o1.Release()
	if o2.RetainCount() != 1 || o3.RetainCount() != 1 {
		t.Fail()
	}
	o2.Release()
	o3.Release()
}
