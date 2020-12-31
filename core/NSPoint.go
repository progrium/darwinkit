// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import "fmt"

type NSPoint struct {
	X float64
	Y float64
}

func (pt NSPoint) String() string {
	return fmt.Sprintf("(%v, %v)", pt.X, pt.Y)
}
