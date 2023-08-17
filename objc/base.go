// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

// #import <stdlib.h>
// #import <stdint.h>
// #import <stdbool.h>
// #import <objc/runtime.h>
//
// void* IMP_ImplementationWithBlock(void* bp);
// void* IMP_GetBlock(void* ptr);
// bool IMP_RemoveBlock(void* ptr);
//
// void* Method_GetName(void* method);
// const char* Method_GetTypeEncoding(void* method);
// void* Method_GetImplementation(void* method);
// void* Method_SetImplementation(void* method, void* imp);
//
// const char* Property_GetName(void* property);
// const char* Property_GetAttributes(void* property);
// char* Property_CopyAttributeValue(void* property, const char* name);
// objc_property_attribute_t * Property_CopyAttributeList(void* property, unsigned int *outCount);
//
// void* Selector_SEL_RegisterName(const char* name);
// const char* Selector_SEL_GetName(void* ptr);
//
// void Objc_SetAssociatedObject(void* ptr, const void* key, void* valuePtr, uintptr_t policy);
// void* Objc_GetAssociatedObject(void* ptr, const void* key);
// void Objc_RemoveAssociatedObjects(void* ptr);
import "C"
import (
	"unsafe"
)

type IMP struct {
	ptr unsafe.Pointer
}

func (i IMP) Ptr() unsafe.Pointer {
	return i.ptr
}

func IMPFrom(ptr unsafe.Pointer) IMP {
	return IMP{ptr}
}

func IMPWithBlock(b Block) IMP {
	return IMP{
		ptr: C.IMP_ImplementationWithBlock(b.ptr),
	}
}

func (i IMP) Block() Block {
	return Block{
		ptr: C.IMP_GetBlock(i.ptr),
	}
}

func (i IMP) RemoveBlock() bool {
	return bool(C.IMP_RemoveBlock(i.ptr))
}

func convertToObjcPropertyAttributes(attributes []PropertyAttribute) []*C.objc_property_attribute_t {
	aps := make([]*C.objc_property_attribute_t, len(attributes))
	for i := 0; i < len(attributes); i++ {
		a := attributes[i]
		aps[i] = &C.objc_property_attribute_t{
			name:  C.CString(a.Name),
			value: C.CString(a.Value),
		}
	}
	return aps
}

type MethodDescription struct {
	Name  Selector
	Types string
}

type Property struct {
	ptr unsafe.Pointer
}

func (p Property) Ptr() unsafe.Pointer {
	return p.ptr
}

type PropertyAttribute struct {
	Name  string
	Value string
}

const (
	PropertyAttributeNameNonatomic = "N"
	PropertyAttributeNameStrong    = "&"
	PropertyAttributeNameRetain    = PropertyAttributeNameStrong
	PropertyAttributeNameWeak      = "W"
	PropertyAttributeNameReadonly  = "R"
	PropertyAttributeNameGetter    = "G"
	PropertyAttributeNameSetter    = "S"
	PropertyAttributeNameIvar      = "V"
	PropertyAttributeNameType      = "T"
)

func (p Property) Name() string {
	cs := C.Property_GetName(p.Ptr())
	return C.GoString(cs)
}

func (p Property) Attributes() string {
	cs := C.Property_GetAttributes(p.Ptr())
	return C.GoString(cs)
}

func (p Property) CopyAttributeValue(name string) string {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cs := C.Property_CopyAttributeValue(p.Ptr(), cname)
	if cs != nil {
		defer C.free(unsafe.Pointer(cs))
	}
	return C.GoString(cs)
}

func (p Property) CopyAttributeList() []PropertyAttribute {
	var outCount C.uint
	pap := C.Property_CopyAttributeList(p.Ptr(), &outCount)
	if outCount == 0 {
		return nil
	}
	defer C.free(unsafe.Pointer(pap))
	pas := unsafe.Slice(pap, int(outCount))
	slice := make([]PropertyAttribute, len(pas))
	for i := 0; i < len(pas); i++ {
		pa := pas[i]
		slice[i] = PropertyAttribute{
			Name:  C.GoString(pa.name),
			Value: C.GoString(pa.value),
		}
	}
	return slice
}

type Method struct {
	ptr unsafe.Pointer
}

func (m Method) Ptr() unsafe.Pointer {
	return m.ptr
}

func (m Method) Name() Selector {
	return Selector{
		ptr: C.Method_GetName(m.ptr),
	}
}

func (m Method) TypeEncoding() string {
	r := C.Method_GetTypeEncoding(m.ptr)
	return C.GoString(r)
}

func (m Method) Implementation() IMP {
	return IMP{
		ptr: C.Method_GetImplementation(m.ptr),
	}
}

func (m Method) SetImplementation(imp IMP) IMP {
	return IMP{
		ptr: C.Method_SetImplementation(m.ptr, imp.ptr),
	}
}

type Ivar struct {
	ptr unsafe.Pointer
}

func (i Ivar) Ptr() unsafe.Pointer {
	return i.Ptr()
}

type Category struct {
	ptr unsafe.Pointer
}

func (c Category) Ptr() unsafe.Pointer {
	return c.Ptr()
}

type Selector struct {
	ptr unsafe.Pointer
}

func SelectorFrom(ptr unsafe.Pointer) Selector {
	return Selector{ptr}
}

var selectorCache = SyncCache[string, Selector]{}

// Sel returns a cached method selector by name, registering it if needed.
func Sel(selName string) Selector {
	return selectorCache.Load(selName, func(selName string) Selector {
		return RegisterSelectorName(selName)
	})
}

// RegisterSelectorName registers a method with the Objective-C runtime system, maps the method name to a selector, and returns the selector value.
func RegisterSelectorName(name string) Selector {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return Selector{C.Selector_SEL_RegisterName(cname)}
}

func (s Selector) Ptr() unsafe.Pointer {
	return s.ptr
}

func (s Selector) Name() string {
	cstr := C.Selector_SEL_GetName(s.ptr)
	return C.GoString(cstr)
}

// Type to specify the behavior of an association. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/objc_associationpolicy?language=objc
type AssociationPolicy uintptr

const (
	ASSOCIATION_ASSIGN           = 0     // Specifies a weak reference to the associated object.
	ASSOCIATION_RETAIN_NONATOMIC = 1     // Specifies a strong reference to the associated object. The association is not made atomically.
	ASSOCIATION_COPY_NONATOMIC   = 3     // Specifies that the associated object is copied. The association is not made atomically.
	ASSOCIATION_RETAIN           = 01401 // Specifies a strong reference to the associated object. The association is made atomically.
	ASSOCIATION_COPY             = 01403 //Specifies that the associated object is copied. The association is made atomically.
)

// Sets an associated value for a given object using a given key and association policy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418509-objc_setassociatedobject?language=objc
func SetAssociatedObject(o IObject, key unsafe.Pointer, value IObject, policy AssociationPolicy) {
	C.Objc_SetAssociatedObject(o.Ptr(), key, value.Ptr(), C.uintptr_t(policy))
}

// Returns the value associated with a given object for a given key. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418865-objc_getassociatedobject?language=objc
func GetAssociatedObject(o IObject, key unsafe.Pointer) Object {
	return ObjectFrom(C.Objc_GetAssociatedObject(o.Ptr(), key))
}

// Removes all associations for a given object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418683-objc_removeassociatedobjects?language=objc
func RemoveAssociatedObjects(o IObject) {
	C.Objc_RemoveAssociatedObjects(o.Ptr())
}
