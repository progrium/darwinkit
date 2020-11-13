// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

/*
#cgo LDFLAGS: -framework Foundation -framework CoreFoundation -framework WebKit
*/
import "C"
import "github.com/progrium/macdriver/pkg/objc"

var (
	True  objc.Object
	False objc.Object
)

func init() {
	True = NSNumber_WithBool(true)
	False = NSNumber_WithBool(false)
}

func String(str string) NSString {
	return NSString_FromString(str)
}

func Point(x float64, y float64) NSPoint {
	return NSPoint{X: x, Y: y}
}

func Size(width float64, height float64) NSSize {
	return NSSize{Width: width, Height: height}
}

func Rect(x, y, w, h float64) NSRect {
	return NSMakeRect(x, y, w, h)
}

func URL(url string) NSURL {
	return NSURL_Init(url)
}
