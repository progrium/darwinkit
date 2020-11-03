// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foundation

import "github.com/progrium/macdriver/pkg/objc"

type NSRunLoop struct {
	objc.Object
}

func NSRunLoopCurrent() NSRunLoop {
	return NSRunLoop{objc.GetClass("NSRunLoop").SendMsg("currentRunLoop")}
}

func NSRunLoopMain() NSRunLoop {
	return NSRunLoop{objc.GetClass("NSRunLoop").SendMsg("mainRunLoop")}
}

func (rl NSRunLoop) Run() {
	rl.SendMsg("run")
}
