// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cocoa

/*
#cgo LDFLAGS: -framework AppKit
*/
import "C"

func Color(r, g, b, a float64) NSColor {
	return NSColor_Init(r, g, b, a)
}

func Font(fontName string, size float64) NSFont {
	return NSFont_Init(fontName, size)
}
