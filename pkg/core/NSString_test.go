// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"testing"
)

func TestString(t *testing.T) {
	hi := "hello, world!"
	str := String(hi)
	if str.String() != hi {
		t.Error("mismatch")
	}
}
