// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"github.com/progrium/macdriver/objc"
)

type NSDictionary struct {
	objc.Object
}

var NSDictionary_ = objc.Get("NSDictionary")

func NSDictionary_New() NSDictionary {
	return NSDictionary{NSDictionary_.Alloc().Init()}
}

func NSDictionary_Init(valueKeys ...interface{}) NSDictionary {
	return NSDictionary{NSDictionary_.Alloc().Send("initWithObjectsAndKeys:", valueKeys...)}
}

func (d NSDictionary) ObjectForKey(key objc.Object) objc.Object {
	return d.Send("objectForKey:", key)
}
