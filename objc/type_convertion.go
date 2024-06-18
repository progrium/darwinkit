// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

// #import <stdint.h>
// #include "type_convertion.h"
// void* to_ns_string(const char* str);
// const char* to_c_string(void* ptr);
//
// void* to_ns_data(data);
// data to_c_bytes(void* ptr);
//
// void* to_ns_array(array array);
// array to_c_array(void* ptr);
//
// dict to_c_items(void* ptr);
// void* to_ns_dict(dict cDict);
import "C"

import (
	"fmt"
	"reflect"
	"unsafe"

	"github.com/progrium/darwinkit/objc/ffi"
)

var bytesType = reflect.TypeOf([]byte{})
var pointerHolderType = reflect.TypeOf((*Handle)(nil)).Elem()
var selectorType = reflect.TypeOf(Selector{})
var classType = reflect.TypeOf(Class{})

type reValue struct {
	// typ holds the type of the value represented by a Value.
	typ unsafe.Pointer

	// Pointer-valued data or, if flagIndir is set, pointer to data.
	// Valid when either flagIndir is set or typ.pointers() is true.
	ptr unsafe.Pointer

	flag uintptr
}

// flagIndir: val holds a pointer to the data
var flagIndir uintptr = 1 << 7

func ToNSString(s string) unsafe.Pointer {
	cs := C.CString(s)
	defer C.free(unsafe.Pointer(cs))
	return C.to_ns_string(cs)
}

func ToGoString(p unsafe.Pointer) string {
	if p == nil {
		return ""
	}
	return C.GoString(C.to_c_string(p))
}

func ToNSData(bytes []byte) unsafe.Pointer {
	if bytes == nil {
		return nil
	}
	size := len(bytes)
	var p unsafe.Pointer
	var d C.data
	if size == 0 {
		p = C.to_ns_data(d)
	} else {
		p = C.to_ns_data(C.data{
			len:  C.ulong(size),
			data: unsafe.Pointer(&bytes[0]),
		})
	}
	return p
}

func ToGoBytes(p unsafe.Pointer) []byte {
	if p == nil {
		return nil
	}
	d := C.to_c_bytes(p)
	size := int(d.len)
	if size <= 0 {
		return nil
	}
	bytes := unsafe.Slice((*byte)(d.data), size)
	newBytes := make([]byte, size)
	copy(newBytes, bytes)
	return newBytes
}

// slice/dict elements convert
func toNSElement(v reflect.Value) unsafe.Pointer {
	var t = v.Type()
	switch t.Kind() {
	case reflect.String:
		return ToNSString(v.String())
	case reflect.Slice:
		if t.Elem().Kind() == reflect.Uint8 {
			return ToNSData(v.Convert(bytesType).Interface().([]byte))
		} else {
			return ToNSArray(v)
		}
	case reflect.Interface:
		if t.AssignableTo(pointerHolderType) {
			return v.Interface().(Handle).Ptr()
		}
	case reflect.Struct:
		if t.Implements(pointerHolderType) {
			return v.Interface().(Handle).Ptr()
		}
	default:

	}
	panic("not supported types: " + t.Name())
}

// slice/dict elements convert
func toGoElement(ptr unsafe.Pointer, t reflect.Type) reflect.Value {
	switch t.Kind() {
	case reflect.String:
		return reflect.ValueOf(ToGoString(ptr)).Convert(t)
	case reflect.Slice:
		if t.Elem().Kind() == reflect.Uint8 {
			return reflect.ValueOf(ToGoBytes(ptr)).Convert(t)
		} else {
			return ToGoSlice(ptr, t)
		}
	case reflect.Struct:
		if t.Implements(pointerHolderType) {
			// objc object pointer holder struct should have same layout as a pointer
			return reflect.NewAt(t, unsafe.Pointer(&ptr)).Elem()
		}
	case reflect.Interface:
		if t.Implements(pointerHolderType) {
			// objc object pointer holder struct should have same layout as a pointer
			return reflect.NewAt(reflect.TypeOf(Object{}), unsafe.Pointer(&ptr)).Elem()
		}
	default:

	}
	panic("not supported types: " + t.Name())
}

func ToNSArray(slice reflect.Value) unsafe.Pointer {
	if slice.IsNil() {
		return nil
	}
	var cArray C.array
	if slice.Len() > 0 {
		cArrayData := make([]unsafe.Pointer, slice.Len())
		for i := 0; i < slice.Len(); i++ {
			cArrayData[i] = toNSElement(slice.Index(i))
		}
		cArray.data = unsafe.Pointer(&cArrayData[0])
		cArray.len = C.ulong(len(cArrayData))
	}
	return C.to_ns_array(cArray)
}

func ToGoSlice(ptr unsafe.Pointer, sliceType reflect.Type) reflect.Value {
	if ptr == nil {
		return reflect.New(sliceType).Elem()
	}
	ca := C.to_c_array(ptr)
	if ca.len > 0 {
		defer C.free(ca.data)
	}
	ptrSlice := unsafe.Slice((*unsafe.Pointer)(ca.data), int(ca.len))
	var slice = reflect.MakeSlice(sliceType, len(ptrSlice), len(ptrSlice))
	for idx, ptr := range ptrSlice {
		slice.Index(idx).Set(toGoElement(ptr, sliceType.Elem()))
	}
	return slice
}

func ToNSDict(m reflect.Value) unsafe.Pointer {
	if m.IsNil() {
		return nil
	}
	var cDict C.dict
	if m.Len() > 0 {
		keyData := make([]unsafe.Pointer, m.Len())
		valueData := make([]unsafe.Pointer, m.Len())
		for idx, k := range m.MapKeys() {
			v := m.MapIndex(k)
			keyData[idx] = toNSElement(k)
			valueData[idx] = toNSElement(v)
		}
		cDict.key_data = unsafe.Pointer(&keyData[0])
		cDict.value_data = unsafe.Pointer(&valueData[0])
		cDict.len = C.ulong(m.Len())
	}
	return C.to_ns_dict(cDict)
}

func ToGoMap(ptr unsafe.Pointer, mapType reflect.Type) reflect.Value {
	if ptr == nil {
		return reflect.New(mapType).Elem()
	}
	cDict := C.to_c_items(ptr)
	if cDict.len > 0 {
		defer C.free(cDict.key_data)
		defer C.free(cDict.value_data)
	}
	keys := unsafe.Slice((*unsafe.Pointer)(cDict.key_data), int(cDict.len))
	values := unsafe.Slice((*unsafe.Pointer)(cDict.value_data), int(cDict.len))
	var m = reflect.MakeMap(mapType)
	for idx, k := range keys {
		v := values[idx]
		m.SetMapIndex(toGoElement(k, mapType.Key()), toGoElement(v, mapType.Elem()))
	}
	return m
}

// convertToGoValue convert objc value to go value.
// param p: the pointer to objc value
// param t: the go value type
func convertToGoValue(p unsafe.Pointer, t reflect.Type) reflect.Value {
	switch t.Kind() {
	case reflect.Bool:
		return reflect.ValueOf(*((*uint8)(p)) == 1).Convert(t)
	case reflect.Int8:
		return reflect.ValueOf(*((*int8)(p))).Convert(t)
	case reflect.Int16:
		return reflect.ValueOf(*((*int16)(p))).Convert(t)
	case reflect.Int32:
		return reflect.ValueOf(*((*int32)(p))).Convert(t)
	case reflect.Int, reflect.Int64:
		return reflect.ValueOf(*((*int64)(p))).Convert(t)
	case reflect.Uint8:
		return reflect.ValueOf(*((*uint8)(p))).Convert(t)
	case reflect.Uint16:
		return reflect.ValueOf(*((*uint16)(p))).Convert(t)
	case reflect.Uint32:
		return reflect.ValueOf(*((*uint32)(p))).Convert(t)
	case reflect.Uint, reflect.Uint64, reflect.Uintptr:
		return reflect.ValueOf(*((*uint64)(p))).Convert(t)
	case reflect.Float32:
		return reflect.ValueOf(*((*float32)(p))).Convert(t)
	case reflect.Float64:
		return reflect.ValueOf(*((*float64)(p))).Convert(t)
	case reflect.Pointer, reflect.UnsafePointer:
		return reflect.ValueOf(*(*unsafe.Pointer)(p)).Convert(t)
	case reflect.String:
		return reflect.ValueOf(ToGoString(*(*unsafe.Pointer)(p))).Convert(t)
	case reflect.Slice:
		if t.Elem().Kind() == reflect.Uint8 {
			return reflect.ValueOf(ToGoBytes(*(*unsafe.Pointer)(p)))
		} else {
			return reflect.ValueOf(ToGoSlice(*(*unsafe.Pointer)(p), t).Interface())
		}
	case reflect.Map:
		return ToGoMap(*(*unsafe.Pointer)(p), t)
	case reflect.Struct:
		return reflect.NewAt(t, p).Elem()
	case reflect.Func:
		rv := wrapBlockInGoFunc(*(*unsafe.Pointer)(p), t)
		return rv
	default:
		if t.Implements(pointerHolderType) {
			// objc object pointer holder struct should have same layout as a pointer
			return reflect.NewAt(t, p).Elem()
		}
		panic(fmt.Sprintf("not support type: %v %v", t, t.Kind()))
	}
}

// convertToObjcValue convert go value to objc value, return the pointer to objc value, and the ffi-type
func convertToObjcValue(v reflect.Value) unsafe.Pointer {
	if !v.IsValid() {
		var p unsafe.Pointer = nil
		return unsafe.Pointer(&p)
	}

	rt := v.Type()
	switch rt.Kind() {
	case reflect.Bool:
		var cv uint8 = 0
		if v.Bool() {
			cv = 1
		}
		return unsafe.Pointer(&cv)
	case reflect.Int8:
		cv := int8(v.Int())
		return unsafe.Pointer(&cv)
	case reflect.Int16:
		cv := int16(v.Int())
		return unsafe.Pointer(&cv)
	case reflect.Int32:
		cv := int32(v.Int())
		return unsafe.Pointer(&cv)
	case reflect.Int, reflect.Int64:
		cv := int64(v.Int())
		return unsafe.Pointer(&cv)
	case reflect.Uint8:
		cv := uint8(v.Uint())
		return unsafe.Pointer(&cv)
	case reflect.Uint16:
		cv := uint16(v.Uint())
		return unsafe.Pointer(&cv)
	case reflect.Uint32:
		cv := uint32(v.Uint())
		return unsafe.Pointer(&cv)
	case reflect.Uint, reflect.Uint64, reflect.Uintptr:
		cv := uint64(v.Uint())
		return unsafe.Pointer(&cv)
	case reflect.Float32:
		cv := float32(v.Float())
		return unsafe.Pointer(&cv)
	case reflect.Float64:
		cv := float64(v.Float())
		return unsafe.Pointer(&cv)
	case reflect.Pointer:
		// allow use of *string
		if rt.Elem().Kind() == reflect.String {
			if v.IsNil() {
				var p unsafe.Pointer = nil
				return unsafe.Pointer(&p)
			}
			sp := ToNSString(v.Elem().String())
			return unsafe.Pointer(&sp)
		}
		// otherwise treat as unsafe pointer
		cv := v.UnsafePointer()
		return unsafe.Pointer(&cv)
	case reflect.UnsafePointer:
		cv := v.UnsafePointer()
		return unsafe.Pointer(&cv)
	case reflect.Interface:
		if v.Type().AssignableTo(pointerHolderType) {
			cv := v.Interface().(Handle).Ptr()
			return unsafe.Pointer(&cv)
		}
		panic(fmt.Sprintf("not support type: %T", v))
	case reflect.Struct:
		if v.Type().Implements(pointerHolderType) {
			cv := v.Interface().(Handle).Ptr()
			return unsafe.Pointer(&cv)
		}
		return getStructValuePointer(v)
	case reflect.String:
		sp := ToNSString(v.String())
		return unsafe.Pointer(&sp)
	case reflect.Slice:
		if rt.Elem().Kind() == reflect.Uint8 {
			dp := ToNSData(v.Bytes())
			return unsafe.Pointer(&dp)
		} else {
			sp := ToNSArray(v)
			return unsafe.Pointer(&sp)
		}
	case reflect.Map:
		dp := ToNSDict(v)
		return unsafe.Pointer(&dp)
	case reflect.Func:
		fp := CreateMallocBlock(v.Interface())
		return unsafe.Pointer(&fp.ptr)
	default:
		panic(fmt.Sprintf("not support type: %T, kind: %v", v, v.Kind()))
	}
}

func setGoValueToObjcPointer(rv reflect.Value, p unsafe.Pointer) {

	rt := rv.Type()
	switch rt.Kind() {
	case reflect.Bool:
		if rv.Bool() {
			*(*uint8)(p) = 1
		} else {
			*(*uint8)(p) = 0
		}
	case reflect.Int8:
		*(*int8)(p) = int8(rv.Int())
	case reflect.Int16:
		*(*int16)(p) = int16(rv.Int())
	case reflect.Int32:
		*(*int32)(p) = int32(rv.Int())
	case reflect.Int, reflect.Int64:
		*(*int64)(p) = int64(rv.Int())
	case reflect.Uint8:
		*(*uint8)(p) = uint8(rv.Uint())
	case reflect.Uint16:
		*(*uint16)(p) = uint16(rv.Uint())
	case reflect.Uint32:
		*(*uint32)(p) = uint32(rv.Uint())
	case reflect.Uint, reflect.Uint64, reflect.Uintptr:
		*(*uint64)(p) = uint64(rv.Uint())
	case reflect.Float32:
		*(*float32)(p) = float32(rv.Float())
	case reflect.Float64:
		*(*float64)(p) = float64(rv.Float())
	case reflect.UnsafePointer, reflect.Pointer:
		*(*unsafe.Pointer)(p) = rv.UnsafePointer()
	case reflect.Interface:
		if rv.Type().AssignableTo(pointerHolderType) {
			*(*unsafe.Pointer)(p) = rv.Interface().(Handle).Ptr()
		} else {
			panic(fmt.Sprintf("not support type: %v", rv.Type()))
		}
	case reflect.Struct:
		if rv.Type().Implements(pointerHolderType) {
			*(*unsafe.Pointer)(p) = rv.Interface().(Handle).Ptr()
		} else {
			sp := getStructValuePointer(rv)
			size := rv.Type().Size()
			copy(unsafe.Slice((*byte)(p), size), unsafe.Slice((*byte)(sp), size))
		}
	case reflect.String:
		sp := ToNSString(rv.String())
		*(*unsafe.Pointer)(p) = sp
	case reflect.Slice:
		if rt.Elem().Kind() == reflect.Uint8 {
			dp := ToNSData(rv.Bytes())
			*(*unsafe.Pointer)(p) = dp
		} else {
			sp := ToNSArray(rv)
			*(*unsafe.Pointer)(p) = sp
		}
	case reflect.Map:
		dp := ToNSDict(rv)
		*(*unsafe.Pointer)(p) = dp
	case reflect.Func:
		fp := CreateMallocBlock(rv.Interface())
		*(*unsafe.Pointer)(p) = fp.ptr
	default:
		panic(fmt.Sprintf("not support type: %v, kind: %v", rv.Type(), rv.Kind()))
	}
}

func getStructValuePointer(value reflect.Value) unsafe.Pointer {
	vsp := (*reValue)(unsafe.Pointer(&value))

	if vsp.flag&flagIndir == 0 {
		return unsafe.Pointer(&vsp.ptr)
	} else {
		return vsp.ptr
	}
}

// toFFIType return the ffi type
func toFFIType(rt reflect.Type) *ffi.Type {
	switch rt.Kind() {
	case reflect.Bool:
		return ffi.TypeUint8
	case reflect.Int8:
		return ffi.TypeSint8
	case reflect.Int16:
		return ffi.TypeSint16
	case reflect.Int32:
		return ffi.TypeSint32
	case reflect.Int, reflect.Int64:
		return ffi.TypeSint64
	case reflect.Uint8:
		return ffi.TypeUint8
	case reflect.Uint16:
		return ffi.TypeUint16
	case reflect.Uint32:
		return ffi.TypeUint32
	case reflect.Uint, reflect.Uint64, reflect.Uintptr:
		return ffi.TypeUint64
	case reflect.Float32:
		return ffi.TypeFloat
	case reflect.Float64:
		return ffi.TypeDouble
	case reflect.UnsafePointer, reflect.Pointer:
		return ffi.TypePointer
	case reflect.Interface:
		if rt.AssignableTo(pointerHolderType) {
			return ffi.TypePointer
		}
		panic(fmt.Sprintf("not support type: %v", rt))
	case reflect.Struct:
		if rt.Size() == 0 {
			return ffi.TypeVoid
		}
		if rt.Implements(pointerHolderType) {
			return ffi.TypePointer
		}
		return getStructFFIType(rt)
	case reflect.String:
		return ffi.TypePointer
	case reflect.Slice:
		return ffi.TypePointer
	case reflect.Map:
		return ffi.TypePointer
	case reflect.Func:
		return ffi.TypePointer
	default:
		panic(fmt.Sprintf("not support type: %v, kind: %v", rt, rt.Kind()))
	}
}

func getStructFFIType(t reflect.Type) *ffi.Type {
	fn := t.NumField()
	var fts []*ffi.Type
	for i := 0; i < fn; i++ {
		ft := t.Field(i).Type
		switch ft.Kind() {
		case reflect.Array:
			for i := 0; i < ft.Len(); i++ {
				fts = append(fts, toFFIType(ft.Elem()))
			}
		default:
			fts = append(fts, toFFIType(ft))
		}
	}
	return ffi.MakeStructType(fts)
}
