// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSFont struct {
	objc.Object
}

func NSFont_Init(fontName string, size float64) NSFont {
	return NSFont{objc.Get("NSFont").Send("fontWithName:size:", core.String(fontName), size)}
}
