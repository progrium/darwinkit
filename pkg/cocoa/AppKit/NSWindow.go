// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package appkit

import (
	"github.com/progrium/macdriver/pkg/cocoa"
	. "github.com/progrium/macdriver/pkg/cocoa/Foundation"
	"github.com/progrium/macdriver/pkg/objc"
)

const (
	NSBorderlessWindowMask         NSUInteger = 0
	NSTitledWindowMask             NSUInteger = 1 << 0
	NSClosableWindowMask           NSUInteger = 1 << 1
	NSMiniaturizableWindowMask     NSUInteger = 1 << 2
	NSResizableWindowMask          NSUInteger = 1 << 3
	NSTexturedBackgroundWindowMask NSUInteger = 1 << 8
)

const (
	NSBackingStoreRetained    = 0
	NSBackingStoreNonretained = 1
	NSBackingStoreBuffered    = 2
)

type NSBackingStoreType NSUInteger

type NSWindow struct {
	objc.Object
}

func NSWindow_New() NSWindow {
	return NSWindow{objc.GetClass("NSWindow").Alloc().Init()}
}

func NSWindow_Init(rect NSRect, windowStyle NSUInteger, bufferingType NSBackingStoreType, deferCreation bool) NSWindow {
	obj := objc.GetClass("NSWindow").Alloc().
		SendMsg("initWithContentRect:styleMask:backing:defer:",
			rect, windowStyle, bufferingType, deferCreation)
	return NSWindow{obj}
}

func (win NSWindow) Display() {
	win.SendMsg("display")
}

func (win NSWindow) MakeKeyAndOrderFront(sender objc.Object) {
	win.SendMsg("makeKeyAndOrderFront:", sender)
}

func (win NSWindow) SetTitle(title string) {
	win.SendMsg("setTitle:", cocoa.String(title))
}

func (win NSWindow) Title() string {
	return win.SendMsg("title").String()
}

func (win NSWindow) SetContentView(view objc.Object) {
	win.SendMsg("setContentView:", view)
}

func (win NSWindow) ContentView() objc.Object {
	return win.SendMsg("contentView")
}
