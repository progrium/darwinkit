// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foundation

import "fmt"

type NSRect struct {
	Origin NSPoint
	Size   NSSize
}

func (r NSRect) String() string {
	return fmt.Sprintf("(%v, %v, %v, %v)", r.Origin.X, r.Origin.Y, r.Size.Width, r.Size.Height)
}

func NSRectMake(x, y, w, h float32) NSRect {
	return NSRect{
		NSPoint{
			x, y,
		},
		NSSize{
			w, h,
		},
	}
}
