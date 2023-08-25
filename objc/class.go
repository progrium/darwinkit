// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

// #import <stdlib.h>
// #import <stdint.h>
// #import <stdbool.h>
//
// void* Objc_GetClass(const char* name);
// void* Objc_AllocateClassPair(void* superClass, const char* name, size_t extraBytes);
// void Objc_RegisterClassPair(void* class);
// void Objc_DisposeClassPair(void* class);
//
// void* Class_CreateInstance(void* cls, unsigned idxIvars);
// const char* Class_GetName(void *cls);
// void Class_SetVersion(void* cls, int);
// int Class_GetVersion(void* cls);
// void* Class_GetSuperClass(void* cls);
// bool Class_RespondsToSelectorAddMethod(void* cls, void* sel);
// bool Class_AddProtocol(void* cls, void* protocol);
// bool Class_AddMethod(void* cls, void* sel, void* imp, const char* types);
// void* Class_ReplaceMethod(void* cls, void* sel, void* imp, const char* types);
// void* Class_GetMethodImplementation(void* cls, void* sel);
// void* Class_GetMethodImplementationStret(void* cls, void* sel);
// void* Class_GetInstanceMethod(void* cls, void* sel);
// void* Class_GetClassMethod(void* cls, void* sel);
// void* Class_CopyMethodList(void* cls, unsigned int *outCount);
// void* Class_GetProperty(void* cls, const char *name);
// void* Class_CopyPropertyList(void* cls, unsigned int *outCount);
// bool Class_AddProperty(void* cls, const char *name, void* attributes, unsigned int attributeCount);
// void Class_ReplaceProperty(void* cls, const char *name, void* attributes, unsigned int attributeCount);
import "C"
import (
	"unsafe"
)

type IClass interface {
	Handle
	CreateInstance(idxIvars uint) Object
	Name() string
	SetVersion(version int)
	Version() int
	MetaClass() Class
	SuperClass() Class
	RespondsToSelector(sel Selector) bool
	AddMethod(sel Selector, imp IMP, types string) bool
	ReplaceMethod(sel Selector, imp IMP, types string) IMP
	MethodImplementation(sel Selector) IMP
	MethodImplementationStret(sel Selector) IMP
	InstanceMethod(sel Selector) Method
	ClassMethod(sel Selector) Method
	Property(name string) Property
	AddProperty(name string, attributes []PropertyAttribute) bool
	ReplaceProperty(name string, attributes []PropertyAttribute)
	AddProtocol(protocol Protocol) bool
	CopyMethodList() []Method
	CopyPropertyList() []Property
}

type Class struct {
	ptr unsafe.Pointer
}

func (c Class) Ptr() unsafe.Pointer {
	return c.ptr
}

// GetClass returns an Objective-C Class by name
func GetClass(name string) Class {
	nameStr := C.CString(name)
	defer C.free(unsafe.Pointer(nameStr))
	return Class{
		ptr: C.Objc_GetClass(nameStr),
	}
}

// Creates a new class and metaclass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418559-objc_allocateclasspair?language=objc
func AllocateClass(superClass Class, name string, extraBytes uint) Class {
	nameStr := C.CString(name)
	defer C.free(unsafe.Pointer(nameStr))
	return Class{
		ptr: C.Objc_AllocateClassPair(superClass.Ptr(), nameStr, C.size_t(extraBytes)),
	}
}

// Registers a class that was allocated using [AllocateClass] [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418603-objc_registerclasspair?language=objc
func RegisterClass(class IClass) {
	C.Objc_RegisterClassPair(class.Ptr())
}

// Destroys a class and its associated metaclass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418912-objc_disposeclasspair?language=objc
func DisposeClass(class IClass) {
	C.Objc_DisposeClassPair(class.Ptr())
}

func (c Class) CreateInstance(idxIvars uint) Object {
	ptr := C.Class_CreateInstance(c.ptr, C.uint(idxIvars))
	return ObjectFrom(ptr)
}

func (c Class) Name() string {
	cname := C.Class_GetName(c.ptr)
	name := C.GoString(cname)
	return name
}

func (c Class) MetaClass() Class {
	return ObjectFrom(c.ptr).Class()
}

func (c Class) SetVersion(version int) {
	C.Class_SetVersion(c.ptr, C.int(version))
}

func (c Class) Version() int {
	return int(C.Class_GetVersion(c.ptr))
}

func (c Class) SuperClass() Class {
	ptr := C.Class_GetSuperClass(c.ptr)
	return Class{ptr}
}

func (c Class) RespondsToSelector(sel Selector) bool {
	r := C.Class_RespondsToSelectorAddMethod(c.ptr, sel.ptr)
	return bool(r)
}

func (c Class) AddMethod(sel Selector, imp IMP, types string) bool {
	ctypes := C.CString(types)
	defer C.free(unsafe.Pointer(ctypes))
	r := C.Class_AddMethod(c.ptr, sel.ptr, imp.ptr, ctypes)
	return bool(r)
}

func (c Class) ReplaceMethod(sel Selector, imp IMP, types string) IMP {
	ctypes := C.CString(types)
	defer C.free(unsafe.Pointer(ctypes))
	r := C.Class_ReplaceMethod(c.ptr, sel.ptr, imp.ptr, ctypes)
	return IMP{ptr: r}
}

func (c Class) MethodImplementation(sel Selector) IMP {
	r := C.Class_GetMethodImplementation(c.ptr, sel.ptr)
	return IMP{ptr: r}
}

func (c Class) MethodImplementationStret(sel Selector) IMP {
	r := C.Class_GetMethodImplementationStret(c.ptr, sel.ptr)
	return IMP{ptr: r}
}

func (c Class) InstanceMethod(sel Selector) Method {
	r := C.Class_GetInstanceMethod(c.ptr, sel.ptr)
	return Method{ptr: r}
}

func (c Class) ClassMethod(sel Selector) Method {
	r := C.Class_GetClassMethod(c.ptr, sel.ptr)
	return Method{ptr: r}
}

func (c Class) CopyMethodList() []Method {
	var count C.uint
	rp := C.Class_CopyMethodList(c.ptr, &count)
	return convertToSliceAndFreePointer[Method](rp, int(count))
}

func (c Class) AddProtocol(protocol Protocol) bool {
	r := C.Class_AddProtocol(c.ptr, protocol.ptr)
	return bool(r)
}

func (c Class) Property(name string) Property {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return Property{
		ptr: C.Class_GetProperty(c.Ptr(), cname),
	}
}

func (c Class) CopyPropertyList() []Property {
	var outCount C.uint
	pp := C.Class_CopyPropertyList(c.Ptr(), &outCount)
	return convertToSliceAndFreePointer[Property](pp, int(outCount))
}

func convertToSliceAndFreePointer[T Handle](p unsafe.Pointer, count int) []T {
	if p == nil {
		return nil
	}
	defer C.free(p)
	ps := unsafe.Slice((*unsafe.Pointer)(unsafe.Pointer(p)), count)
	slice := make([]T, count)
	for i := 0; i < int(count); i++ {
		slice[i] = ForceCast[unsafe.Pointer, T](ps[i])
	}
	return slice
}

func (c Class) AddProperty(name string, attributes []PropertyAttribute) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cattributes := convertToObjcPropertyAttributes(attributes)
	defer func() {
		for _, ca := range cattributes {
			C.free(unsafe.Pointer(ca.name))
			C.free(unsafe.Pointer(ca.value))
		}
	}()
	r := C.Class_AddProperty(c.Ptr(), cname, unsafe.Pointer(&cattributes[0]), C.uint(len(attributes)))
	return bool(r)
}

func (c Class) ReplaceProperty(name string, attributes []PropertyAttribute) {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cattributes := convertToObjcPropertyAttributes(attributes)
	defer func() {
		for _, ca := range cattributes {
			C.free(unsafe.Pointer(ca.name))
			C.free(unsafe.Pointer(ca.value))
		}
	}()
	C.Class_ReplaceProperty(c.Ptr(), cname, unsafe.Pointer(&cattributes[0]), C.uint(len(attributes)))
}
