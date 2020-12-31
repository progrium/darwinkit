// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cocoa

import (
	"github.com/progrium/macdriver/objc"
)

type NSColor struct {
	objc.Object
}

func NSColor_Init(r, g, b, a float64) NSColor {
	return NSColor{objc.Get("NSColor").Send("colorWithRed:green:blue:alpha:", r, g, b, a)}
}

func NSColor_Clear() NSColor {
	return NSColor{objc.Get("NSColor").Get("clearColor")}
}
