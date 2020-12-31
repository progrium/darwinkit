// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import "github.com/progrium/macdriver/objc"

type NSAutoreleasePool struct {
	objc.Object
}

var NSAutoreleasePool_ = objc.Get("NSAutoreleasePool")

func NSAutoreleasePool_New() NSAutoreleasePool {
	return NSAutoreleasePool{NSAutoreleasePool_.Alloc().Init()}
}
