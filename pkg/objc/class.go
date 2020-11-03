// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

/*
#cgo LDFLAGS: -lobjc -framework Foundation
#define __OBJC2__ 1
#include <objc/runtime.h>
#include <objc/message.h>

static unsigned long key = 0xbadc0c0a;

void *GoObjc_GetClassByName(char *name) {
	return (void *) objc_getClass(name);
}

void *GoObjc_GetObjectClass(void *obj) {
	return (void *) object_getClass(obj);
}

void *GoObjc_AllocateClassPair(void *superCls, char *name) {
	return (void *) objc_allocateClassPair(superCls, name, 0);
}

void GoObjc_ClassAddMethod(void *subCls, void *sel, void *imp, char *typ) {
	class_addMethod(subCls, sel, imp, typ);
}

void GoObjc_SetInternal(void *obj, void *ptr) {
	objc_setAssociatedObject(obj, (void *)&key, ptr, OBJC_ASSOCIATION_ASSIGN);
}

void *GoObjc_GetInternal(void *obj) {
	return (void *) objc_getAssociatedObject(obj, (void *)&key);
}

void GoObjc_RegisterClass(void *cls) {
	objc_registerClassPair(cls);
}

char *GoObjc_GetClassName(void *cls) {
	return (char *) class_getName(cls);
}
*/
import "C"
import (
	"reflect"
	"strings"
	"unsafe"
)

type classInfo struct {
	typ       reflect.Type
	methodMap map[string]interface{}
	refs      map[uintptr]unsafe.Pointer
	setters   map[string]struct{}
}

func (ci classInfo) MethodForSelector(sel string) interface{} {
	return ci.methodMap[sel]
}

func (ci *classInfo) AddRef(ptr unsafe.Pointer) {
	if ci.refs == nil {
		ci.refs = make(map[uintptr]unsafe.Pointer)
	}
	ci.refs[uintptr(ptr)] = ptr
}

func (ci *classInfo) RemoveRef(ptr unsafe.Pointer) {
	if ci.refs != nil {
		delete(ci.refs, uintptr(ptr))
	}
}

var (
	classMap map[string]classInfo
)

func init() {
	classMap = make(map[string]classInfo)
}

// A Class represents a special Objective-C
// class Object.
type Class interface {
	Object

	// AddMethod registers a Go function to be called whenever
	// an instance of the class receives a message for the given
	// selector.
	//
	// The AddMethod call only works for classes created with
	// objc.NewClass. Method calls will only be received by the
	// Go function if instances of the class are instanciated
	// by calling objc.NewGoInstance.
	AddMethod(selector string, fn interface{})
}

// NewClass returns a new class. The value parameter must
// point to a value of the struct that is used to represent
// instances of the class in Go.
func NewClass(value interface{}) Class {
	typ := reflect.ValueOf(value).Type()

	// The tag of the first field contains a
	// description of the class: its name in
	// Objective-C and its super class.
	//
	// For example:
	//   `objc:"GOAppDelegate : NSObject"`
	field := typ.Field(0)
	descrStr := field.Tag.Get("objc")
	descr := strings.Split(descrStr, " : ")
	if len(descr) != 2 {
		panic("objc: bad description string for class " + typ.Name())
	}

	className := descr[0]
	superClassName := descr[1]
	superClass := GetClass(superClassName)

	ptr := C.GoObjc_AllocateClassPair(unsafe.Pointer(superClass.Pointer()), C.CString(className))
	if ptr == nil {
		panic("unable to AllocateClassPair")
	}

	// Check whether the class has any IBOutlets, and generate
	// appropriate setter methods so we can load NIBs on OS X.
	setters := map[string]struct{}{}
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Tag.Get("objc") == "IBOutlet" {
			if field.Type.Implements(objectInterfaceType) {
				setters["set"+field.Name+":"] = struct{}{}
				break
			} else {
				panic("IBOutlets must implement objc.Object")
			}
		}
	}

	// Register the IBOutlet setters.
	for setterSelector, _ := range setters {
		sel := selectorWithName(setterSelector)
		typeInfo := encVoid + encId + encSelector + encId
		C.GoObjc_ClassAddMethod(ptr, sel, methodCallTarget(), C.CString(typeInfo))
	}

	// Register the setValue:forKey: method for our custom IBOutlet handling
	// if the class has any IBOutlets. (Currently only relevant for the iOS runtime)
	if len(setters) > 0 {
		sel := selectorWithName("setValue:forKey:")
		typeInfo := encVoid + encId + encSelector + encId + encId
		C.GoObjc_ClassAddMethod(ptr, sel, methodCallTarget(), C.CString(typeInfo))
	}

	// Register the dealloc method to be able to properly remove the classInfo
	// reference to our internal pointer.
	sel := selectorWithName("dealloc")
	typeInfo := encVoid + encId + encSelector
	C.GoObjc_ClassAddMethod(ptr, sel, methodCallTarget(), C.CString(typeInfo))

	classMap[className] = classInfo{
		typ:       reflect.TypeOf(value),
		methodMap: make(map[string]interface{}),
		setters:   setters,
	}

	return object{ptr: uintptr(ptr)}
}

// Lookup a Class by name
func GetClass(name string) Class {
	return object{ptr: uintptr(C.GoObjc_GetClassByName(C.CString(name)))}
}

// Get the Class of an object. This equivalent to sending the
// class message to the object, but this is not always possible.
// (The SendMsg method, for example, needs to get the class of the
// object it is sending a message to. This would end in an infinite
// loop.)
func getObjectClass(obj Object) Class {
	classPtr := C.GoObjc_GetObjectClass(unsafe.Pointer(obj.Pointer()))
	return object{ptr: uintptr(classPtr)}
}

// RegisterClass registers a Class with the Objective-C runtime.
func RegisterClass(class Class) {
	C.GoObjc_RegisterClass(unsafe.Pointer(class.Pointer()))
}

// className returns the name of the Class represented by object.
func (cls object) className() string {
	return C.GoString(C.GoObjc_GetClassName(unsafe.Pointer(cls.Pointer())))
}

// AddMethod adds a new method to a Class.
func (cls object) AddMethod(selector string, fn interface{}) {
	clsName := cls.className()
	clsInfo := classMap[clsName]

	// Check if this method has already implicitly been
	// added by an IBOutlet tagged struct field.
	if _, isSetter := clsInfo.setters[selector]; isSetter {
		panic("objc: unable to add method '" + selector + "'; would shadow IBOutlet setter with same name.")
	}

	sel := selectorWithName(selector)
	typeInfo := funcTypeInfo(fn)
	C.GoObjc_ClassAddMethod(unsafe.Pointer(cls.Pointer()), sel, methodCallTarget(), C.CString(typeInfo))

	// Add the method to the class's method map
	clsInfo.methodMap[selector] = fn
}

// setInternalPointer sets an internal pointer on the object.
// This is used to implement correct method dispatch for
// Objective-C classes created from within Go.
func (obj object) setInternalPointer(value unsafe.Pointer) {
	C.GoObjc_SetInternal(unsafe.Pointer(obj.Pointer()), unsafe.Pointer(value))
}

// internalPointer returns the object's internal pointer.
// Must only be called on objects that are known to have
// an internal pointer set.
func (obj object) internalPointer() unsafe.Pointer {
	return C.GoObjc_GetInternal(unsafe.Pointer(obj.Pointer()))
}
