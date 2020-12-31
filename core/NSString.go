// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package core

import (
	"reflect"
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

const (
	NSUTF8StringEncoding = 4
)

type NSString struct {
	objc.Object
}

var NSString_ = objc.Get("NSString")

func NSString_FromString(str string) NSString {
	hdrp := (*reflect.StringHeader)(unsafe.Pointer(&str))
	obj := NSString_.Alloc().Send("initWithBytes:length:encoding:", hdrp.Data, hdrp.Len, NSUTF8StringEncoding)
	return NSString_FromObject(obj)
}

func NSString_FromObject(obj objc.Object) NSString {
	return NSString{obj}
}

func (s NSString) SizeWithAttributes(attrs NSDictionary) NSSize {
	size := NSSize{}
	s.Send("sizeWithAttributes:", attrs, &size)
	return size
}
