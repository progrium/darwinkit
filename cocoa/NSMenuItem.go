// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
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

func NSMenuItem_Separator() NSMenuItem {
	return NSMenuItem{NSMenuItem_.Get("separatorItem")}
}

func (i NSMenuItem) SetSubmenu(submenu NSMenu) {
	i.Set("submenu:", submenu)
}

func (i NSMenuItem) Submenu() NSMenu {
	return NSMenu{i.Get("submenu")}
}

func (i NSMenuItem) Hidden() bool {
	return i.Get("hidden").Bool()
}

func (i NSMenuItem) SetHidden(b bool) {
	i.Set("hidden:", b)
}

func (i NSMenuItem) Enabled() bool {
	return i.Get("enabled").Bool()
}

func (i NSMenuItem) SetEnabled(b bool) {
	i.Set("enabled:", b)
}

func (i NSMenuItem) Title() string {
	return i.Get("title").String()
}

func (i NSMenuItem) SetTitle(s string) {
	i.Set("title:", core.String(s))
}

func (i NSMenuItem) Image() NSImage {
	return NSImage{i.Get("image")}
}

func (i NSMenuItem) SetImage(obj NSImage) {
	i.Set("image:", obj)
}

func (i NSMenuItem) ToolTip() string {
	return i.Get("toolTip").String()
}

func (i NSMenuItem) SetToolTip(s string) {
	i.Set("toolTip:", core.String(s))
}

func (i NSMenuItem) Target() objc.Object {
	return i.Get("target")
}

func (i NSMenuItem) SetTarget(obj objc.Object) {
	i.Set("target:", obj)
}

func (i NSMenuItem) Action() objc.Selector {
	return objc.Sel(i.Get("action").String())
}

func (i NSMenuItem) SetAction(sel objc.Selector) {
	i.Set("action:", sel)
}

func (i NSMenuItem) SetState(state int) {
	i.Set("state:", state)
}

func (i NSMenuItem) State() int64 {
	return i.Get("state").Int()
}
