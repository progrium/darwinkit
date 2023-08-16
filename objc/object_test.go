// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"testing"
)

func TestNSObject_GetClass(t *testing.T) {
	o := NewObject()
	if o.Class().Name() != "NSObject" {
		t.Fail()
	}
}

func TestNSObject_RetainCount(t *testing.T) {
	o := NewObject()
	if o.RetainCount() != 1 {
		t.Fail()
	}

	o.Retain()
	if o.RetainCount() != 2 {
		t.Fail()
	}

	o.Release()
	o.Release()
}
