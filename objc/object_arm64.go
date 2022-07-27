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
	// fixme(mkrautz): 64-bit only only; also not sure if
	// this check is even valid for IEEE floats.
	if obj.ptr&0xffffffff00000000 == 0 {
		return float64(math.Float32frombits(uint32(obj.ptr)))
	}
	return math.Float64frombits(uint64(obj.ptr))
}
