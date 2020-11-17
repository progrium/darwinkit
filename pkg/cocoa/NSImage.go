// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cocoa

import (
	"github.com/progrium/macdriver/pkg/core"
	"github.com/progrium/macdriver/pkg/objc"
)

type NSImage struct {
	objc.Object
}

func NSImage_InitWithData(data core.NSData) NSImage {
	return NSImage{objc.Get("NSImage").Alloc().Send("initWithData:", data)}
}

func NSImage_ImageNamed(name string) NSImage {
	return NSImage{objc.Get("NSImage").Send("imageNamed:", core.String(name))}
}

func (i NSImage) Size() (size core.NSSize) {
	i.Send("size", &size)
	return size
}

func (i NSImage) SetSize(size core.NSSize) {
	i.Send("setSize:", size)
}
