// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package appkit

import (
	"github.com/progrium/macdriver/pkg/cocoa"
	"github.com/progrium/macdriver/pkg/objc"
)

type NSMenuItem struct {
	objc.Object
}

func NSMenuItem_Init(itemName string, action objc.Selector, keyEquivalent string) NSMenuItem {
	return NSMenuItem{objc.GetClass("NSMenuItem").Alloc().SendMsg("initWithTitle:action:keyEquivalent:",
		cocoa.String(itemName), action, cocoa.String(keyEquivalent))}
}

func NSMenuItem_New() NSMenuItem {
	return NSMenuItem{objc.GetClass("NSMenuItem").Alloc().Init()}
}

func (item NSMenuItem) SetSubmenu(submenu NSMenu) {
	item.SendMsg("setSubmenu:", submenu)
}
