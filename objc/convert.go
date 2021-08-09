package objc

import "unsafe"

// #cgo LDFLAGS: -lobjc
// #include <objc/objc.h>
import "C"

func asBool(v interface{}) bool {
	return asInt(v) > 0
}

func asInt(v interface{}) int64 {
	switch v := v.(type) {
	case bool:
		if v {
			return 1
		}
		return 0
	case int:
		return int64(v)
	case uint:
		return int64(v)
	case int8:
		return int64(v)
	case uint8:
		return int64(v)
	case int16:
		return int64(v)
	case uint16:
		return int64(v)
	case int32:
		return int64(v)
	case uint32:
		return int64(v)
	case int64:
		return int64(v)
	case uint64:
		return int64(v)
	case float32:
		return int64(v)
	case float64:
		return int64(v)
	}
	panic("not an int")
}

func asUint(v interface{}) uint64 {
	switch v := v.(type) {
	case bool:
		if v {
			return 1
		}
		return 0
	case int:
		return uint64(v)
	case uint:
		return uint64(v)
	case int8:
		return uint64(v)
	case uint8:
		return uint64(v)
	case int16:
		return uint64(v)
	case uint16:
		return uint64(v)
	case int32:
		return uint64(v)
	case uint32:
		return uint64(v)
	case int64:
		return uint64(v)
	case uint64:
		return uint64(v)
	case float32:
		return uint64(v)
	case float64:
		return uint64(v)
	}
	panic("not an int")
}

func asFloat(v interface{}) float64 {
	switch v := v.(type) {
	case bool:
		if v {
			return 1
		}
		return 0
	case int:
		return float64(v)
	case uint:
		return float64(v)
	case int8:
		return float64(v)
	case uint8:
		return float64(v)
	case int16:
		return float64(v)
	case uint16:
		return float64(v)
	case int32:
		return float64(v)
	case uint32:
		return float64(v)
	case int64:
		return float64(v)
	case uint64:
		return float64(v)
	case float32:
		return float64(v)
	case float64:
		return float64(v)
	}
	panic("not a float")
}

func asObject(v interface{}) Object {
	switch v := v.(type) {
	case uintptr:
		return ObjectPtr(v)
	case unsafe.Pointer:
		return ObjectPtr(uintptr(v))
	case C.id:
		return ObjectPtr(uintptr(unsafe.Pointer(v)))
	case Object:
		return v
	}
	panic("not an object")
}

func asClass(v interface{}) Class {
	switch v := v.(type) {
	case uintptr:
		return ClassPtr(v)
	case unsafe.Pointer:
		return ClassPtr(uintptr(v))
	case C.id:
		return ClassPtr(uintptr(unsafe.Pointer(v)))
	case C.Class:
		return ClassPtr(uintptr(unsafe.Pointer(v)))
	case Class:
		return v
	}
	panic("not an object")
}

func asSelector(v interface{}) Selector {
	switch v := v.(type) {
	case uintptr:
		return SelectorPtr(unsafe.Pointer(v))
	case unsafe.Pointer:
		return SelectorPtr(v)
	case string:
		return Sel(v)
	case C.SEL:
		return SelectorPtr(unsafe.Pointer(v))
	case Selector:
		return v
	}
	panic("not a selector")
}
