// Copyright (c) 2012-2022 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"fmt"
	"reflect"
)

func typeInfoForType(typ reflect.Type) string {
	if typ.Implements(classInterfaceType) {
		return encClass
	} else if typ.Implements(objectInterfaceType) {
		return encId
	} else if typ.Implements(selectorInterfaceType) {
		return encSelector
	}

	kind := typ.Kind()
	switch kind {
	case reflect.Bool:
		return encBool
	case reflect.Int:
		return encInt
	case reflect.Int8:
		return encChar
	case reflect.Int16:
		return encShort
	case reflect.Int32:
		return encInt
	case reflect.Int64:
		return encULong
	case reflect.Uint:
		return encUInt
	case reflect.Uint8:
		return encUChar
	case reflect.Uint16:
		return encUShort
	case reflect.Uint32:
		return encUInt
	case reflect.Uint64:
		return encULong
	case reflect.Uintptr:
		return encPtr
	case reflect.Float32:
		return encFloat
	case reflect.Float64:
		return encDouble
	case reflect.Ptr:
		return encPtr
	}

	panic("typeinfo: unhandled/invalid kind " + fmt.Sprintf("%v", kind) + " " + fmt.Sprintf("%v", typ))
}
