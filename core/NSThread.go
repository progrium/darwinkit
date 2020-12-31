// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import "github.com/progrium/macdriver/objc"

type NSThread struct {
	objc.Object
}

var NSThread_ = objc.Get("NSThread")

func NSThread_IsMainThread() bool {
	return NSThread_.Send("isMainThread") != nil
}
