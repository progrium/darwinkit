// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package appkit

import (
	"github.com/progrium/macdriver/pkg/cocoa"
	"github.com/progrium/macdriver/pkg/objc"
)

type NSImage struct {
	objc.Object
}

func NSImage_ImageNamed(name string) objc.Object {
	return NSImage{objc.GetClass("NSImage").SendMsg("imageNamed:", cocoa.String(name))}
}
