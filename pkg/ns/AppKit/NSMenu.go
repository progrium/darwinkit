// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package appkit

import (
	. "github.com/progrium/macdriver/pkg/ns/Foundation"
	"github.com/progrium/macdriver/pkg/objc"
)

type NSMenu struct {
	objc.Object
}

func NewNSMenu() NSMenu {
	return NSMenu{objc.GetClass("NSMenu").SendMsg("alloc").SendMsg("init")}
}

func NewNSMenuWithTitle(title string) NSMenu {
	return NSMenu{objc.GetClass("NSMenu").SendMsg("alloc").SendMsg("initWithTitle:", NSStringFromString(title))}
}

func (menu NSMenu) SetTitle(title string) {
	menu.SendMsg("setTitle:", NSStringFromString(title))
}

func (menu NSMenu) Title() string {
	return menu.SendMsg("title").String()
}

func (menu NSMenu) AddItem(item NSMenuItem) {
	menu.SendMsg("addItem:", item)
}

func (menu NSMenu) RemoveItem(item NSMenuItem) {
	menu.SendMsg("removeItem:", item)
}
