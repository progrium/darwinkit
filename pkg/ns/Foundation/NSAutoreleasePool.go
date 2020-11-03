// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foundation

import "github.com/progrium/macdriver/pkg/objc"

type NSAutoreleasePool struct {
	objc.Object
}

func NewNSAutoreleasePool() NSAutoreleasePool {
	return NSAutoreleasePool{objc.GetClass("NSAutoreleasePool").Alloc().Init()}
}
