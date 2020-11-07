// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cocoa

import (
	"github.com/progrium/macdriver/pkg/core"
	"github.com/progrium/macdriver/pkg/objc"
)

const (
	NSTextAlignmentLeft      = 0
	NSTextAlignmentRight     = 1
	NSTextAlignmentCenter    = 2
	NSTextAlignmentJustified = 3
	NSTextAlignmentNatural   = 4
)

type NSTextView struct {
	objc.Object
}

func NSTextView_Init(frame core.NSRect) NSTextView {
	return NSTextView{objc.Get("NSTextView").Alloc().Send("initWithFrame:", frame)}
}
