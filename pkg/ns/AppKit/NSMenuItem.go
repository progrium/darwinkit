// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package appkit

import (
	. "github.com/progrium/macdriver/pkg/ns/Foundation"
	"github.com/progrium/macdriver/pkg/objc"
)

type NSMenuItem struct {
	objc.Object
}

func NewNSMenuItem(itemName string, action objc.Selector, keyEquivalent string) NSMenuItem {
	return NSMenuItem{objc.GetClass("NSMenuItem").SendMsg("alloc").SendMsg("initWithTitle:action:keyEquivalent:",
		NSStringFromString(itemName), action, NSStringFromString(keyEquivalent))}
}
