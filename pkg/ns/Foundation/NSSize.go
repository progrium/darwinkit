// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foundation

import "fmt"

type NSSize struct {
	Width  float32
	Height float32
}

func (sz NSSize) String() string {
	return fmt.Sprintf("(%v, %v)", sz.Width, sz.Height)
}
