// Copyright 2021 Liu Dong. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

// #import <stdlib.h>
// #import <stdint.h>
// #import <stdbool.h>
// #import <objc/runtime.h>
//
// void* New_ProtocolImpl(void* class, uintptr_t goID);
// void* Objc_GetProtocol(const char* name);
// void* Objc_AllocateProtocol(const char* name);
// const char* Protocol_GetName(void *protocol);
// struct objc_method_description Protocol_GetMethodDescription(void* protocol, void* sel, bool required, bool instanceMethod);
// struct objc_method_description* Protocol_CopyMethodDescriptionList(void* protocol, bool required, bool instanceMethod, unsigned int *outCount);
// void* Protocol_CopyProtocolList(void* protocol, unsigned int *outCount);
// void* Protocol_GetProperty(void* protocol, const char *name, bool required, bool isInstanceProperty);
// void* Protocol_CopyPropertyList(void* protocol, unsigned int *outCount);
import "C"
import (
	"reflect"
	"runtime/cgo"
	"strings"
	"unsafe"

	"github.com/progrium/darwinkit/objc/ffi"
)

type IProtocol interface {
	Handle
	Name() string
	MethodDescription(sel Selector, required bool, instanceMethod bool) MethodDescription
	CopyMethodDescriptionList(required bool, instanceMethod bool) []MethodDescription
	CopyProtocolList() []Protocol
	Property(name string, required bool, isInstanceProperty bool) Property
	CopyPropertyList() []Property
}

type Protocol struct {
	ptr unsafe.Pointer
}

func (p Protocol) Ptr() unsafe.Pointer {
	return p.ptr
}

func GetProtocol(name string) Protocol {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return Protocol{
		ptr: C.Objc_GetProtocol(cname),
	}
}

func AllocateProtocol(name string) Protocol {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return Protocol{
		ptr: C.Objc_AllocateProtocol(cname),
	}
}

func (p Protocol) Name() string {
	cname := C.Protocol_GetName(p.ptr)
	name := C.GoString(cname)
	return name
}

func (p Protocol) MethodDescription(sel Selector, required bool, instanceMethod bool) MethodDescription {
	md := C.Protocol_GetMethodDescription(p.ptr, sel.ptr, C.bool(required), C.bool(instanceMethod))
	return MethodDescription{
		Name:  Selector{unsafe.Pointer(md.name)},
		Types: C.GoString(md.types),
	}
}

func (p Protocol) CopyMethodDescriptionList(required bool, instanceMethod bool) []MethodDescription {
	var count C.uint
	mdp := C.Protocol_CopyMethodDescriptionList(p.ptr, C.bool(required), C.bool(instanceMethod), &count)
	if mdp == nil {
		return nil
	}
	defer C.free(unsafe.Pointer(mdp))
	mds := unsafe.Slice(mdp, int(count))
	r := make([]MethodDescription, int(count))
	for i := 0; i < int(count); i++ {
		r[i] = MethodDescription{
			Name:  Selector{unsafe.Pointer(mds[i].name)},
			Types: C.GoString(mds[i].types),
		}
	}
	return r
}

func (p Protocol) CopyProtocolList() []Protocol {
	var count C.uint
	pp := C.Protocol_CopyProtocolList(p.Ptr(), &count)
	return convertToSliceAndFreePointer[Protocol](pp, int(count))
}

func (p Protocol) Property(name string, required bool, isInstanceProperty bool) Property {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return Property{
		ptr: C.Protocol_GetProperty(p.Ptr(), cname, C.bool(required), C.bool(isInstanceProperty)),
	}
}

func (p Protocol) CopyPropertyList() []Property {
	var outCount C.uint
	pp := C.Protocol_CopyPropertyList(p.Ptr(), &outCount)
	return convertToSliceAndFreePointer[Property](pp, int(outCount))
}

var classCache = SyncCache[string, *classInfo]{} //go protocol interface name to ClassInfo
var baseClass = GetClass("ProtocolImplBase")

type instanceInfo struct {
	instance     any        // the go instance
	protocolName string     // go protocol interface name
	classInfo    *classInfo // the class info for this instance
}

type classInfo struct {
	class       IClass
	methodInfos map[string]*methodInfo // selector name to MethodInfo
}

type methodInfo struct {
	required bool          // if is required protocol method
	hasFunc  reflect.Value // the hasXXX func for this method
}

func WrapAsProtocol[T any](protocolName string, d T) Object {
	t := reflect.TypeOf((*T)(nil)).Elem()
	if t.Kind() != reflect.Interface {
		panic("generic type T should be interface type")
	}

	dv := reflect.ValueOf(d)
	ci := createClass(dv.Type(), protocolName)
	ii := &instanceInfo{
		classInfo:    ci,
		instance:     d,
		protocolName: protocolName,
	}
	h := cgo.NewHandle(ii)
	return ObjectFrom(C.New_ProtocolImpl(ci.class.Ptr(), C.uintptr_t(h)))
}

func createClass(t reflect.Type, protocolName string) *classInfo {
	if t.Kind() == reflect.Interface {
		panic("should not be interface type")
	}

	return classCache.Load(protocolName, func(key string) *classInfo {
		class := AllocateClass(baseClass, protocolName+"Adaptor", 0)
		protocol := GetProtocol(protocolName)
		class.AddProtocol(protocol)

		// check for conflicts using selectorToGoName
		var selectorToSuffix = map[string]string{}
		var selectorSeen = map[string]string{}
		for _, md := range getProtocolMethods(protocol) {
			sel := md.Name.Name()
			goSel := selectorToGoName(sel)
			if selectorSeen[goSel] != "" {
				// if conflict, mark the longer selector to add suffix
				if len(selectorSeen[goSel]) > len(sel) {
					selectorToSuffix[goSel] = selectorSeen[goSel]
				} else {
					selectorToSuffix[goSel] = sel
				}
			} else {
				selectorSeen[goSel] = sel
			}
		}

		var methodInfos = map[string]*methodInfo{} // selector name to method signature
		for _, md := range getProtocolMethods(protocol) {
			selector := md.Name
			selName := selector.Name()
			goFuncName := selectorToGoName(selName)
			if selectorToSuffix[goFuncName] == selName {
				goFuncName += "_"
			}
			goMethod, ok := t.MethodByName(goFuncName)
			if !ok {
				if md.required {
					panic("required method not implemented:" + selName)
				} else {
					continue
				}
			}
			addProtocolMethod(class, md, goMethod)
			hasFunc, _ := t.MethodByName("Has" + goFuncName)

			methodInfos[selName] = &methodInfo{
				required: md.required,
				hasFunc:  hasFunc.Func,
			}
		}

		RegisterClass(class)
		return &classInfo{
			class:       class,
			methodInfos: methodInfos,
		}

	})
}

type methodDescription struct {
	MethodDescription
	required       bool
	instanceMethod bool
}

func getProtocolMethods(protocol Protocol) []methodDescription {
	if protocol.Name() == "NSObject" {
		return nil
	}
	var mds []methodDescription
	for _, md := range protocol.CopyMethodDescriptionList(true, true) {
		mds = append(mds, methodDescription{
			MethodDescription: md,
			required:          true,
			instanceMethod:    true,
		})
	}

	for _, md := range protocol.CopyMethodDescriptionList(false, true) {
		mds = append(mds, methodDescription{
			MethodDescription: md,
			required:          false,
			instanceMethod:    true,
		})
	}

	for _, parent := range protocol.CopyProtocolList() {
		mds = append(mds, getProtocolMethods(parent)...)
	}
	return mds
}

func addProtocolMethod(class IClass, md methodDescription, method reflect.Method) {
	//TODO: class method
	if !md.instanceMethod {
		return
	}

	rt := method.Type
	if rt.NumOut() > 1 {
		panic("too many return value")
	}
	goArgc := rt.NumIn()
	var objcArgTypes = make([]*ffi.Type, goArgc+1)
	objcArgTypes[0] = ffi.TypePointer // objc instance or class pointer
	objcArgTypes[1] = ffi.TypePointer // selector pointer
	// first go fun param is receiver
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
		o := ObjectFrom(*(*unsafe.Pointer)(objcArgs[0]))
		handle := Call[uintptr](o, Sel("goID"))
		instance := cgo.Handle(handle).Value().(*instanceInfo)

		var goArgs = make([]reflect.Value, len(objcArgs)-1)
		goArgs[0] = reflect.ValueOf(instance.instance)
		// skip selector param(the second param of objc method)
		for i := 1; i < len(goArgs); i++ {
			goArgs[i] = convertToGoValue(objcArgs[i+1], rt.In(i))
		}
		results := method.Func.Call(goArgs)
		if len(results) == 1 {
			setGoValueToObjcPointer(results[0], ret)
		}
	})
	_ = handle // never free
	if status != ffi.OK {
		panic("ffi prep closure status not ok")
	}
	flag := class.AddMethod(md.Name, IMPFrom(fn), md.Types)
	if !flag {
		panic("add method to protocol failed")
	}
}

//export respondsTo
func respondsTo(goID uintptr, sel unsafe.Pointer) bool {
	selName := SelectorFrom(sel).Name()
	ii := cgo.Handle(goID).Value().(*instanceInfo)
	mi := ii.classInfo.methodInfos[selName]
	if mi == nil {
		return false
	}

	v := reflect.ValueOf(ii.instance)
	if mi.required {
		return true
	}

	return mi.hasFunc.Call([]reflect.Value{v})[0].Bool()
}

// menuWillOpen: -> MenuWillOpen
// menu:updateItem:atIndex:shouldCancel: -> MenuUpdateItemAtIndexShouldCancel
func selectorToGoName(sel string) string {
	var sb strings.Builder
	sb.Grow(len(sel))
	for _, item := range strings.Split(sel, ":") {
		if len(item) == 0 {
			continue
		}
		sb.WriteString(strings.ToUpper(item[:1]))
		sb.WriteString(item[1:])
	}
	return sb.String()
}
