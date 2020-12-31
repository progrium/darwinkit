// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cocoa

import (
	"github.com/progrium/macdriver/objc"
)

type NSImageView struct {
	objc.Object
}

func NSImageView_New() NSImageView {
	return NSImageView{objc.Get("NSImageView").Alloc().Init()}
}

func (imgView NSImageView) SetImage(img objc.Object) {
	imgView.Send("setImage:", img)
}

func (imgView NSImageView) Image() NSImage {
	return NSImage{imgView.Send("image")}
}
