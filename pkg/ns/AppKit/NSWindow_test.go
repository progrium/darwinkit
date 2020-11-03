// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package appkit

import (
	"testing"

	. "github.com/progrium/macdriver/pkg/ns/Foundation"
)

func TestTitle(t *testing.T) {
	pool := NewNSAutoreleasePool()
	defer pool.Release()

	NSSharedApplication()

	window := NewNSWindow(NSRectMake(0, 0, 500, 500), 0, NSBackingStoreBuffered, false)

	title := "hey"

	window.SetTitle(title)
	if window.Title() != title {
		t.Errorf("bad title")
	}
}
