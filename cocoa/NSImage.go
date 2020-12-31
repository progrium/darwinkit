// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cocoa

/*
#cgo CFLAGS: -x objective-c
#cgo LDFLAGS: -framework Foundation
#import <Foundation/Foundation.h>
#include <Cocoa/Cocoa.h>
#include <CoreGraphics/CoreGraphics.h>
#include <objc/objc-runtime.h>
void NSImage_SetSize(void *id, double x, double y) {
    NSImage *image = id;
    [image setSize:NSMakeSize(x, y)];
}
*/
import "C"

import (
	"unsafe"

	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
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

// broken
func (i NSImage) Size() (size core.NSSize) {
	i.Send("size", &size)
	return size
}

func (i NSImage) SetSize(size core.NSSize) {
	C.NSImage_SetSize(unsafe.Pointer(i.Object.Pointer()), C.double(size.Width), C.double(size.Height))
}

func (i NSImage) SetTemplate(b bool) {
	i.Send("setTemplate:", b)
}

func (i NSImage) IsTemplate() bool {
	return i.Get("isTemplate").Bool()
}

func (i NSImage) SetValueForKey(value, key interface{}) {
	i.Send("setValue:forKey:", value, key)
}
