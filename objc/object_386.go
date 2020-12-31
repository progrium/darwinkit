// Copyright (c) 2013 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import "math"

func (obj object) Uint() uint64 {
	return uint64(obj.ptr)
}

func (obj object) Int() int64 {
	return int64(obj.ptr)
}

func (obj object) Bool() bool {
	return obj.ptr == 1
}

func (obj object) Float() float64 {
	u32 := uint32(obj.ptr)
	f32 := math.Float32frombits(u32)
	return float64(f32)
}
