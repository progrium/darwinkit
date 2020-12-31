// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSBundle struct {
	objc.Object
}

func NSBundle_Main() NSBundle {
	return NSBundle{objc.Get("NSBundle").Send("mainBundle")}
}

func (bundle NSBundle) InfoDictionary() core.NSDictionary {
	return core.NSDictionary{Object: bundle.Send("infoDictionary")}
}
