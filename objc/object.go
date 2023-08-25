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
	"reflect"
	"unsafe"
)

// Handle is an interface for holding an Objective-C pointer
type Handle interface {
	// Ptr returns the underlying unsafe.Pointer
	Ptr() unsafe.Pointer
}

// An interface definition for the [Object] class.
type IObject interface {
	Handle
	IsNil() bool

	Class() Class
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

// Ptr returns unsafe.Pointer or nil
func Ptr(o Handle) unsafe.Pointer {
	if o == nil {
		return nil
	}
	return o.Ptr()
}

func setPtr(obj any, ptr unsafe.Pointer) {
	if o, ok := obj.(*Object); ok {
		o.ptr = ptr
	} else {
		if objPtr := findObjectPointer(reflect.ValueOf(obj)); objPtr != nil {
			objPtr.ptr = ptr
		} else {
			panic("unable to find embedded object")
		}
	}
}

func findObjectPointer(v reflect.Value) *Object {
	if v.Kind() == reflect.Ptr {
		v = v.Elem()
	}
	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		if t.Field(i).Type == reflect.TypeOf(Object{}) {
			ptr := v.Field(i).Addr().Interface().(*Object)
			return ptr
		}
		if t.Field(i).Anonymous {
			f := findObjectPointer(v.Field(i))
			if f != nil {
				return f
			}
		}
	}
	return nil
}

// The root class of most Objective-C class hierarchies, from which subclasses inherit a basic interface to the runtime system and the ability to behave as Objective-C objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject?language=objc
type Object struct {
	ptr unsafe.Pointer
}

// Returns a Boolean value that indicates whether the receiver is an instance of given class or an instance of any class that inherits from that class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418956-nsobject/1418511-iskindofclass
func (o Object) IsKindOfClass(class Class) bool {
	return bool(C.Object_IsKindOfClass(o.Ptr(), class.Ptr()))
}

// Returns a Boolean value that indicates whether the receiver is an instance of a given class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418956-nsobject/1418766-ismemberofclass
func (o Object) IsMemberOfClass(class Class) bool {
	return bool(C.Object_IsMemberOfClass(o.Ptr(), class.Ptr()))
}

// Returns a Boolean value that indicates whether the receiver implements or inherits a method that can respond to a specified message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418956-nsobject/1418583-respondstoselector
func (o Object) RespondsToSelector(sel Selector) bool {
	return bool(C.Object_RespondsToSelector(o.Ptr(), sel.ptr))
}

// Returns a Boolean value that indicates whether the receiver conforms to a given protocol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418956-nsobject/1418515-conformstoprotocol
func (o Object) ConformsToProtocol(protocol Protocol) bool {
	return bool(C.Object_ConformsToProtocol(o.Ptr(), protocol.ptr))
}

// Make an [Object] from an unsafe.Pointer
func ObjectFrom(ptr unsafe.Pointer) Object {
	return Object{ptr}
}

// IsNil returns true if the underlying pointer is nil
func (o Object) IsNil() bool {
	return o.ptr == nil
}

// Ptr returns the underlying unsafe.Pointer
func (o Object) Ptr() unsafe.Pointer {
	return o.ptr
}

// Instantiate a new [Object] with Autorelease called
func NewObject() Object {
	o := ObjectFrom(C.C_NSObject_NewObject())
	o.Autorelease()
	return o
}

// Returns the class object for the receiver’s class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418956-nsobject/1571949-class
func (o Object) Class() Class {
	return Class{C.Object_GetClass(o.Ptr())}
}

// Returns a Boolean value that indicates whether the receiver does not descend from NSObject. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418956-nsobject/1418528-isproxy/
func (o Object) IsProxy() bool {
	return bool(C.Object_IsProxy(o.Ptr()))
}

// func (o Object) Copy() Object {
// 	v := ObjectFrom(C.C_NSObject_Copy(o.Ptr()))
// 	v.Autorelease()
// 	return v
// }

// func (o Object) MutableCopy() Object {
// 	v := ObjectFrom(C.C_NSObject_MutableCopy(o.Ptr()))
// 	v.Autorelease()
// 	return v
// }

// Sends a specified message to the receiver and returns the result of the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418956-nsobject/1418867-performselector?language=objc
func (o Object) PerformSelector(sel Selector) Object {
	rp := C.Object_PerformSelector(o.Ptr(), sel.Ptr())
	return ObjectFrom(rp)
}

// Sends a message to the receiver with an object as the argument. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418956-nsobject/1418764-performselector?language=objc
func (o Object) PerformSelectorWithObject(sel Selector, object Object) Object {
	var param = Ptr(object)
	rp := C.Object_PerformSelector_WithObject(o.Ptr(), sel.Ptr(), param)
	return ObjectFrom(rp)
}

// Sends a message to the receiver with two objects as arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418956-nsobject/1418667-performselector?language=objc
func (o Object) PerformSelectorWithObjectWithObject(sel Selector, obj1, obj2 Object) Object {
	param1 := Ptr(obj1)
	param2 := Ptr(obj2)
	rp := C.Object_PerformSelector_WithObject_WithObject(o.Ptr(), sel.Ptr(), param1, param2)
	return ObjectFrom(rp)
}

// "Do not use this method." [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418956-nsobject/1571952-retaincount?language=objc
func (o Object) RetainCount() uint {
	return uint(C.Object_RetainCount(o.Ptr()))
}

// Increments the receiver’s reference count. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418956-nsobject/1571946-retain?language=objc
func (o Object) Retain() Object {
	return ObjectFrom(C.Object_Retain(o.Ptr()))
}

// Decrements the receiver’s reference count. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418956-nsobject/1571957-release?language=objc
func (o Object) Release() {
	C.Object_Release(o.Ptr())
}

// Decrements the receiver’s retain count at the end of the current autorelease pool block. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/1418956-nsobject/1571951-autorelease?language=objc
func (o Object) Autorelease() Object {
	return ObjectFrom(C.Object_Autorelease(o.Ptr()))
}

// Deallocates the memory occupied by the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject/1571947-dealloc?language=objc
func (o Object) Dealloc() {
	C.Object_Dealloc(o.Ptr())
}

// Returns a string that represents the contents of the receiving class. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/objectivec/nsobject/1418799-description?language=objc
func (o Object) Description() string {
	return C.GoString(C.Object_Description(o.Ptr()))
}
