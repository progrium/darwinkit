// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"testing"
)

func TestSetDeallocListener(t *testing.T) {
	var notified = false
	o := NewObject()
	SetDeallocListener(o, func() {
		notified = true
	})
	o.Release()
	if !notified {
		t.Fail()
	}
}

func TestWithAutoreleasePool(t *testing.T) {

	var o = NewObject()
	WithAutoreleasePool(func() {
		o.Autorelease()
		if o.RetainCount() != 1 {
			t.Failed()
		}
	})
}
