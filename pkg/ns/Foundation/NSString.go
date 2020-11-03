// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package foundation

import (
	"reflect"
	"unsafe"

	"github.com/progrium/macdriver/pkg/objc"
)

const (
	NSUTF8StringEncoding = 4
)

type NSString struct {
	objc.Object
}

func NSStringFromString(str string) NSString {
	hdrp := (*reflect.StringHeader)(unsafe.Pointer(&str))
	obj := objc.GetClass("NSString").SendMsg("alloc").SendMsg("initWithBytes:length:encoding:", hdrp.Data, hdrp.Len, NSUTF8StringEncoding)
	return NSStringFromObject(obj)
}

func NSStringFromObject(obj objc.Object) NSString {
	return NSString{obj}
}
