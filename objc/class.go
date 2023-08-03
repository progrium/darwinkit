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
	"reflect"
	"runtime/cgo"
	"strings"
	"unsafe"

	"github.com/progrium/macdriver/objc/ffi"
)

type IClass interface {
	Pointer
	CreateInstance(idxIvars uint) Object
	GetName() string
	SetVersion(version int)
	GetVersion() int
	GetClass() Class
	GetSuperClass() Class
	RespondsToSelector(sel Selector) bool
	AddMethod(sel Selector, imp IMP, types string) bool
	ReplaceMethod(sel Selector, imp IMP, types string) IMP
	GetMethodImplementation(sel Selector) IMP
	GetMethodImplementationStret(sel Selector) IMP
	GetInstanceMethod(sel Selector) Method
	GetClassMethod(sel Selector) Method
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

// GetClass get and return an objc Class by name
func GetClass(name string) Class {
	nameStr := C.CString(name)
	defer C.free(unsafe.Pointer(nameStr))
	return Class{
		ptr: C.Objc_GetClass(nameStr),
	}
}

func AllocateClassPair(superClass Class, name string, extraBytes uint) Class {
	nameStr := C.CString(name)
	defer C.free(unsafe.Pointer(nameStr))
	return Class{
		ptr: C.Objc_AllocateClassPair(superClass.Ptr(), nameStr, C.size_t(extraBytes)),
	}
}

func RegisterClassPair(class Class) {
	C.Objc_RegisterClassPair(class.Ptr())
}

func DisposeClassPair(class Class) {
	C.Objc_DisposeClassPair(class.Ptr())
}

func (c Class) CreateInstance(idxIvars uint) Object {
	ptr := C.Class_CreateInstance(c.ptr, C.uint(idxIvars))
	return MakeObject(ptr)
}

func (c Class) GetName() string {
	cname := C.Class_GetName(c.ptr)
	name := C.GoString(cname)
	return name
}

// GetClass returns class's meta class.
func (c Class) GetClass() Class {
	return MakeObject(c.ptr).GetClass()
}

func (c Class) SetVersion(version int) {
	C.Class_SetVersion(c.ptr, C.int(version))
}

func (c Class) GetVersion() int {
	return int(C.Class_GetVersion(c.ptr))
}

func (c Class) GetSuperClass() Class {
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

func (c Class) GetMethodImplementation(sel Selector) IMP {
	r := C.Class_GetMethodImplementation(c.ptr, sel.ptr)
	return IMP{ptr: r}
}

func (c Class) GetMethodImplementationStret(sel Selector) IMP {
	r := C.Class_GetMethodImplementationStret(c.ptr, sel.ptr)
	return IMP{ptr: r}
}

func (c Class) GetInstanceMethod(sel Selector) Method {
	r := C.Class_GetInstanceMethod(c.ptr, sel.ptr)
	return Method{ptr: r}
}

func (c Class) GetClassMethod(sel Selector) Method {
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

func (c Class) GetProperty(name string) Property {
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

func convertToSliceAndFreePointer[T Pointer](p unsafe.Pointer, count int) []T {
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

// type T: the ret value type
func CallMethod[T any](o Pointer, selector Selector, params ...any) T {
	ptr := o.Ptr()
	if ptr == nil {
		panic("object is nil")
	}
	argc := len(params)

	var args = make([]unsafe.Pointer, argc+2)
	var argTypes = make([]*ffi.Type, argc+2)
	args[0] = unsafe.Pointer(&ptr)
	argTypes[0] = ffi.TypePointer
	args[1] = unsafe.Pointer(&selector.ptr)
	argTypes[1] = ffi.TypePointer
	for i := 0; i < argc; i++ {
		v := reflect.ValueOf(params[i])
		args[i+2] = convertToObjcValue(v)
		if !v.IsValid() {
			argTypes[i+2] = ffi.TypePointer
		} else {
			argTypes[i+2] = toFFIType(v.Type())
		}
	}

	var ret T
	var retPtr unsafe.Pointer
	rt := reflect.TypeOf(ret)
	if directPointer(rt) {
		retPtr = unsafe.Pointer(&ret)
	} else {
		var i uintptr
		retPtr = unsafe.Pointer(&i)
	}
	retType := toFFIType(rt)

	class := MakeObject(ptr).GetClass()
	var imp IMP
	if ffi.IsStruct(retType) {
		imp = class.GetMethodImplementationStret(selector)
	} else {
		imp = class.GetMethodImplementation(selector)
	}
	if imp.ptr == nil {
		return ret
	}
	cif, status := ffi.PrepCIF(retType, argTypes)
	if status != ffi.OK {
		panic("ffi prep cif status not ok")
	}
	ffi.Call(cif, imp.ptr, retPtr, args)
	if directPointer(rt) {
		return *(*T)(retPtr)
	}

	return convertToGoValue(retPtr, reflect.TypeOf(ret)).Interface().(T)
}

// type that do not need convert
func directPointer(t reflect.Type) bool {
	switch t.Kind() {
	case reflect.String, reflect.Slice, reflect.Map:
		return false
	default:
		return true
	}
}

// AddMethod add a new objc instance method with a go function.
// The first param of go function should be the object instance, the second param should be the method selector.
func AddMethod(class Class, sel Selector, f any) bool {
	rf := reflect.ValueOf(f)

	typeEncoding := _getMethodTypeEncoding(rf.Type(), false)

	imp, handle := wrapGoFuncAsMethodIMP(rf)
	flag := class.AddMethod(sel, imp, typeEncoding)
	if !flag {
		handle.Delete()
	}
	return flag
}

// ReplaceMethod replace objc instance method with a go function.
// The first param of go function should be the object instance, the second param should be the method selector.
func ReplaceMethod(class Class, sel Selector, f any) {
	rf := reflect.ValueOf(f)
	typeEncoding := _getMethodTypeEncoding(rf.Type(), false)

	imp, handle := wrapGoFuncAsMethodIMP(rf)
	_ = handle
	oldIMP := class.ReplaceMethod(sel, imp, typeEncoding)
	if oldIMP.ptr != nil {
		//
	}
}

// AddClassMethod add a new objc class method with a go function.
// The first param of go function should be the class, the second param should be the method selector.
func AddClassMethod(class Class, sel Selector, f any) bool {
	rf := reflect.ValueOf(f)
	typeEncoding := _getMethodTypeEncoding(rf.Type(), true)

	imp, handle := wrapGoFuncAsMethodIMP(rf)
	metaClass := class.GetClass()
	if metaClass.Ptr() == nil {
		panic("no meta class")
	}
	flag := metaClass.AddMethod(sel, imp, typeEncoding)
	if !flag {
		handle.Delete()
	}
	return flag
}

// ReplaceClassMethod replace objc class method with a go function.
func ReplaceClassMethod(class Class, sel Selector, f any) {
	rf := reflect.ValueOf(f)
	typeEncoding := _getMethodTypeEncoding(rf.Type(), true)

	imp, handle := wrapGoFuncAsMethodIMP(rf)
	_ = handle

	oldIMP := MakeObject(class.Ptr()).GetClass().ReplaceMethod(sel, imp, typeEncoding)
	if oldIMP.Ptr() != nil {

	}
}

func _getMethodTypeEncoding(ft reflect.Type, classMethod bool) string {
	if ft.Kind() != reflect.Func {
		panic("not func type")
	}
	if ft.NumOut() > 1 {
		panic("to many return values")
	}
	if ft.NumIn() < 1 || !ft.In(0).AssignableTo(pointerHolderType) {
		panic("first param must be the objc object or class")
	}

	var sb strings.Builder
	if ft.NumOut() == 0 {
		sb.WriteByte('v')
	} else {
		sb.WriteString(getTypeEncoding(ft.Out(0)))
	}
	if classMethod {
		sb.WriteString("#") // class self as first parameter
	} else {
		sb.WriteString("@") // instance self as first parameter
	}
	sb.WriteString(":") // selector
	// skip first go param
	for i := 1; i < ft.NumIn(); i++ {
		sb.WriteString(getTypeEncoding(ft.In(i)))
	}
	return sb.String()
}

// the first param of go func must be o object instance or a class instance;
// other params are mapped from objc method params.
func wrapGoFuncAsMethodIMP(rf reflect.Value) (imp IMP, handle cgo.Handle) {
	if rf.Kind() != reflect.Func {
		panic("f should be a func")
	}
	rt := rf.Type()
	if rt.NumOut() > 1 {
		panic("too many return value")
	}

	goArgc := rt.NumIn()
	var objcArgTypes = make([]*ffi.Type, goArgc+1)
	objcArgTypes[0] = ffi.TypePointer // objc instance or class pointer
	objcArgTypes[1] = ffi.TypePointer // selector pointer
	// skip selector param(the second param of objc method)
	for i := 1; i < goArgc; i++ {
		objcArgTypes[i+1] = toFFIType(rt.In(i))
	}

	var retType *ffi.Type
	if rt.NumOut() == 0 {
		retType = ffi.TypeVoid
	} else {
		retType = toFFIType(rt.Out(0))
	}

	cif, status := ffi.PrepCIF(retType, objcArgTypes)
	if status != ffi.OK {
		panic("ffi prep cif status not ok")
	}

	fn, handle, status := ffi.CreateClosure(cif, func(cif *ffi.CIF, ret unsafe.Pointer, objcArgs []unsafe.Pointer) {
		var goArgs = make([]reflect.Value, len(objcArgs)-1)
		goArgs[0] = convertToGoValue(objcArgs[0], rt.In(0))
		// skip selector param(the second param of objc method)
		for i := 1; i < len(goArgs); i++ {
			goArgs[i] = convertToGoValue(objcArgs[i+1], rt.In(i))
		}
		results := rf.Call(goArgs)
		if len(results) == 1 {
			setGoValueToObjcPointer(results[0], ret)
		}
	})
	if status != ffi.OK {
		panic("ffi prep closure status not ok")
	}

	return MakeIMP(fn), handle
}
