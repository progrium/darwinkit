// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

import (
	"reflect"
	"strconv"
	"strings"
)

func getTypeEncoding(t reflect.Type) string {
	switch t.Kind() {
	case reflect.Bool:
		return "B"
	case reflect.Int8:
		return "c"
	case reflect.Int16:
		return "s"
	case reflect.Int32:
		return "i"
	case reflect.Int64, reflect.Int:
		return "q"
	case reflect.Uint8:
		return "C"
	case reflect.Uint16:
		return "S"
	case reflect.Uint32:
		return "I"
	case reflect.Uint64, reflect.Uint:
		return "Q"
	case reflect.Float32:
		return "f"
	case reflect.Float64:
		return "d"
	case reflect.UnsafePointer:
		return "^v"
	case reflect.Pointer:
		switch t.Elem().Kind() {
		case reflect.Uint8, reflect.Int8:
			return "*"
		default:
			return "^" + getTypeEncoding(t.Elem())
		}

	case reflect.String, reflect.Slice, reflect.Map:
		return "@"
	case reflect.Array:
		return "[" + strconv.Itoa(t.Len()) + getTypeEncoding(t.Elem()) + "]"
	case reflect.Interface:
		if t.AssignableTo(pointerHolderType) {
			return "@"
		}
		panic("unsupported type:" + t.Name())
	case reflect.Struct:
		if t == classType {
			return "#"
		} else if t == selectorType {
			return ":"
		} else if t.Implements(pointerHolderType) {
			return "@"
		} else {
			return getStructTypeEncoding(t)
		}
	default:
		panic("unsupported type:" + t.Name())
	}
}

func getStructTypeEncoding(t reflect.Type) string {
	var sb strings.Builder
	sb.WriteString("{?=")
	for i := 0; i < t.NumField(); i++ {
		sb.WriteString(getTypeEncoding(t.Field(i).Type))
	}
	sb.WriteString("}")
	return sb.String()
}
