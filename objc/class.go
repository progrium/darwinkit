// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

/*
#cgo LDFLAGS: -lobjc -framework Foundation
#define __OBJC2__ 1
#include <stdlib.h>
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

void GoObjc_Swizzle(void *cls, void *sel1, void *sel2) {
	Method m1 = class_getInstanceMethod(cls, sel1);
	Method m2 = class_getInstanceMethod(cls, sel2);
	method_exchangeImplementations(m1, m2);
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

	// Swizzle swaps the implementation of two methods on a class
	// so that messages sent to selectorA are received by the
	// implementation of selectorB and vice-versa.
	Swizzle(selectorA, selectorB string)
}

func NewClass(classname string, superclass string) Class {
	superClass := GetClass(superclass)

	ptr := allocateClassPair(unsafe.Pointer(superClass.Pointer()), classname)
	if ptr == nil {
		panic("unable to AllocateClassPair")
	}

	// Register the dealloc method to be able to properly remove the classInfo
	// reference to our internal pointer.
	classAddMethod(ptr, "dealloc", encVoid+encId+encSelector)

	lazilyRegisterClassInMap(classname)
	return object{ptr: uintptr(ptr)}
}

func lazilyRegisterClassInMap(className string) {
	if _, found := classMap[className]; found {
		return
	}

	classMap[className] = classInfo{
		typ:       reflect.TypeOf(struct{ Object }{}),
		methodMap: map[string]interface{}{},
		setters:   map[string]struct{}{},
	}
}

// NewClass returns a new class. The value parameter must
// point to a value of the struct that is used to represent
// instances of the class in Go.
func NewClassFromStruct(value interface{}) Class {
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

	ptr := allocateClassPair(unsafe.Pointer(superClass.Pointer()), className)
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
	for setterSelector := range setters {
		classAddMethod(ptr, setterSelector, encVoid+encId+encSelector+encId)
	}

	// Register the setValue:forKey: method for our custom IBOutlet handling
	// if the class has any IBOutlets. (Currently only relevant for the iOS runtime)
	if len(setters) > 0 {
		classAddMethod(ptr, "setValue:forKey:", encVoid+encId+encSelector+encId+encId)
	}

	// Register the dealloc method to be able to properly remove the classInfo
	// reference to our internal pointer.
	classAddMethod(ptr, "dealloc", encVoid+encId+encSelector)

	classMap[className] = classInfo{
		typ:       reflect.TypeOf(value),
		methodMap: make(map[string]interface{}),
		setters:   setters,
	}

	return object{ptr: uintptr(ptr)}
}

func allocateClassPair(ptr unsafe.Pointer, className string) unsafe.Pointer {
	cstr := C.CString(className)
	defer C.free(unsafe.Pointer(cstr))
	return C.GoObjc_AllocateClassPair(ptr, cstr)
}

func classAddMethod(ptr unsafe.Pointer, sel string, typeInfo string) {
	typeCStr := C.CString(typeInfo)
	defer C.free(unsafe.Pointer(typeCStr))
	C.GoObjc_ClassAddMethod(ptr, selectorWithName(sel), methodCallTarget(), typeCStr)
}

// Get looks up a class by name.
func Get(name string) Class {
	cstr := C.CString(name)
	defer C.free(unsafe.Pointer(cstr))
	return object{ptr: uintptr(C.GoObjc_GetClassByName(cstr))}
}

// deprecated
func GetClass(name string) Class {
	return Get(name)
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

	classAddMethod(unsafe.Pointer(cls.Pointer()), selector, funcTypeInfo(fn))

	// Add the method to the class's method map
	clsInfo.methodMap[selector] = fn
}

// Swizzle swaps the implementation of two methods.
func (cls object) Swizzle(selectorA, selectorB string) {
	selA := selectorWithName(selectorA)
	selB := selectorWithName(selectorB)
	C.GoObjc_Swizzle(unsafe.Pointer(cls.Pointer()), selA, selB)

	clsName := cls.className()
	clsInfo := classMap[clsName]
	mm := clsInfo.methodMap

	mm[selectorA], mm[selectorB] = mm[selectorB], mm[selectorA]
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
