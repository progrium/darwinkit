// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package appkit

import (
	. "github.com/progrium/macdriver/pkg/ns/Foundation"
	"github.com/progrium/macdriver/pkg/objc"
)

type NSNib struct {
	objc.Object
}

func NewNSNib(name string, bundle NSBundle) NSNib {
	return NSNib{objc.GetClass("NSNib").SendMsg("alloc").SendMsg("initWithNibNamed:bundle:",
		NSStringFromString(name), bundle)}
}

func (nib NSNib) InstantiateWithOwner(owner objc.Object) {
	nib.SendMsg("instantiateNibWithOwner:topLevelObjects:", owner, nil)
}
