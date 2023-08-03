// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

// #import <stdlib.h>
// #import <stdint.h>
// #import <stdbool.h>
//
// void* C_NSObject_NewObject();
// bool Object_IsKindOfClass(void* ptr, void* classPtr);
// bool Object_IsMemberOfClass(void* ptr, void* classPtr);
// bool Object_RespondsToSelector(void* ptr, void* sel);
// bool Object_ConformsToProtocol(void* ptr, void* protocolPtr);
// void* Object_GetClass(void* ptr);
// bool Object_IsProxy(void* ptr);
// int Object_RetainCount(void* ptr);
// void* Object_Retain(void* ptr);
// void Object_Release(void* ptr);
// void* Object_Autorelease(void* ptr);
// void* C_NSObject_Copy(void* ptr);
// void* C_NSObject_MutableCopy(void* ptr);
// void Object_Dealloc(void* ptr);
// void* Object_PerformSelector(void* ptr, void* sel_p);
// void* Object_PerformSelector_WithObject(void* ptr, void* sel_p, void* param);
// void* Object_PerformSelector_WithObject_WithObject(void* ptr, void* sel_p, void* param1, void* param2);
// const char* Object_Description(void* ptr);
import "C"
import (
	"unsafe"
)

// Pointer is an interface for holding an objc pointer
type Pointer interface {
	// Ptr return the delegate objc pointer
	Ptr() unsafe.Pointer
}

// Object is interface for all objc Object type
type IObject interface {
	Pointer
	IsNil() bool

	GetClass() Class
	IsKindOfClass(class Class) bool
	IsMemberOfClass(class Class) bool
	RespondsToSelector(sel Selector) bool
	ConformsToProtocol(protocol Protocol) bool

	IsProxy() bool

	Retain() Object
	Release()
	Autorelease() Object
	RetainCount() uint

	// Copy() IObject
	// MutableCopy() IObject

	Dealloc()
	Description() string
}

// ExtractPtr return the objc ptr hold by Object. If is nil, or contains a nil pointer, return nil
func ExtractPtr(o Pointer) unsafe.Pointer {
	if o == nil {
		return nil
	}
	return o.Ptr()
}

type Object struct {
	ptr unsafe.Pointer
}

func (o Object) IsKindOfClass(class Class) bool {
	return bool(C.Object_IsKindOfClass(o.Ptr(), class.Ptr()))
}

func (o Object) IsMemberOfClass(class Class) bool {
	return bool(C.Object_IsMemberOfClass(o.Ptr(), class.Ptr()))
}

func (o Object) RespondsToSelector(sel Selector) bool {
	return bool(C.Object_RespondsToSelector(o.Ptr(), sel.ptr))
}

func (o Object) ConformsToProtocol(protocol Protocol) bool {
	return bool(C.Object_ConformsToProtocol(o.Ptr(), protocol.ptr))
}

func MakeObject(ptr unsafe.Pointer) Object {
	return Object{ptr}
}

func (o Object) IsNil() bool {
	return o.ptr == nil
}

func (o Object) Ptr() unsafe.Pointer {
	return o.ptr
}

func NewObject() Object {
	o := MakeObject(C.C_NSObject_NewObject())
	o.Autorelease()
	return o
}

func (o Object) GetClass() Class {
	return Class{C.Object_GetClass(o.Ptr())}
}

func (o Object) IsProxy() bool {
	return bool(C.Object_IsProxy(o.Ptr()))
}

// func (o Object) Copy() Object {
// 	v := MakeObject(C.C_NSObject_Copy(o.Ptr()))
// 	v.Autorelease()
// 	return v
// }

// func (o Object) MutableCopy() Object {
// 	v := MakeObject(C.C_NSObject_MutableCopy(o.Ptr()))
// 	v.Autorelease()
// 	return v
// }

func (o Object) PerformSelector(sel Selector) Object {
	rp := C.Object_PerformSelector(o.Ptr(), sel.Ptr())
	return MakeObject(rp)
}

func (o Object) PerformSelector_WithObject(sel Selector, object Object) Object {
	var param = ExtractPtr(object)
	rp := C.Object_PerformSelector_WithObject(o.Ptr(), sel.Ptr(), param)
	return MakeObject(rp)
}

func (o Object) PerformSelector_WithObject_WithObject(sel Selector, obj1, obj2 Object) Object {
	param1 := ExtractPtr(obj1)
	param2 := ExtractPtr(obj2)
	rp := C.Object_PerformSelector_WithObject_WithObject(o.Ptr(), sel.Ptr(), param1, param2)
	return MakeObject(rp)
}

func (o Object) RetainCount() uint {
	return uint(C.Object_RetainCount(o.Ptr()))
}

func (o Object) Retain() Object {
	return MakeObject(C.Object_Retain(o.Ptr()))
}

func (o Object) Release() {
	C.Object_Release(o.Ptr())
}

func (o Object) Autorelease() Object {
	return MakeObject(C.Object_Autorelease(o.Ptr()))
}

func (o Object) Dealloc() {
	C.Object_Dealloc(o.Ptr())
}

func (o Object) Description() string {
	return C.GoString(C.Object_Description(o.Ptr()))
}
