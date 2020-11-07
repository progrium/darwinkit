// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cocoa

import (
	"github.com/progrium/macdriver/pkg/core"
	"github.com/progrium/macdriver/pkg/objc"
)

type NSMenuItem struct {
	objc.Object
}

var NSMenuItem_ = objc.Get("NSMenuItem")

func NSMenuItem_Init(itemName string, action objc.Selector, keyEquivalent string) NSMenuItem {
	return NSMenuItem{NSMenuItem_.Alloc().Send("initWithTitle:action:keyEquivalent:",
		core.String(itemName), action, core.String(keyEquivalent))}
}

func NSMenuItem_New() NSMenuItem {
	return NSMenuItem{NSMenuItem_.Alloc().Init()}
}

func (item NSMenuItem) SetSubmenu(submenu NSMenu) {
	item.Send("setSubmenu:", submenu)
}
