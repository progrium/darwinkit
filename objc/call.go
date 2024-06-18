package objc

import (
	"reflect"
	"runtime/cgo"
	"strings"
	"unsafe"

	"github.com/progrium/darwinkit/objc/ffi"
)

// Call sends a message to the Handle and returns a value
func Call[T any](o Handle, selector Selector, params ...any) T {
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

	class := ObjectFrom(ptr).Class()
	var imp IMP
	if ffi.IsStruct(retType) {
		imp = class.MethodImplementationStret(selector)
	} else {
		imp = class.MethodImplementation(selector)
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

// AddMethod adds an instance method using a Go function.
// The first argument of the Go function should be the object instance,
// the second argument should be the method selector.
func AddMethod(class IClass, sel Selector, f any) bool {
	rf := reflect.ValueOf(f)

	typeEncoding := _getMethodTypeEncoding(rf.Type(), false)

	imp, handle := wrapGoFuncAsMethodIMP(rf)
	flag := class.AddMethod(sel, imp, typeEncoding)
	if !flag {
		handle.Delete()
	}
	return flag
}

// ReplaceMethod replaces an instance method using a Go function.
// The first argument of the Go function should be the object instance,
// the second argument should be the method selector.
func ReplaceMethod(class IClass, sel Selector, f any) {
	rf := reflect.ValueOf(f)
	typeEncoding := _getMethodTypeEncoding(rf.Type(), false)

	imp, handle := wrapGoFuncAsMethodIMP(rf)
	_ = handle
	oldIMP := class.ReplaceMethod(sel, imp, typeEncoding)
	if oldIMP.ptr != nil {
		//
	}
}

// AddClassMethod adds a class method using a Go function.
// The first argument of the Go function should be the class,
// the second argument should be the method selector.
func AddClassMethod(class IClass, sel Selector, f any) bool {
	rf := reflect.ValueOf(f)
	typeEncoding := _getMethodTypeEncoding(rf.Type(), true)

	imp, handle := wrapGoFuncAsMethodIMP(rf)
	metaClass := class.MetaClass()
	if metaClass.Ptr() == nil {
		panic("no meta class")
	}
	flag := metaClass.AddMethod(sel, imp, typeEncoding)
	if !flag {
		handle.Delete()
	}
	return flag
}

// ReplaceClassMethod replaces a class method using a Go function.
// The first argument of the Go function should be the class,
// the second argument should be the method selector.
func ReplaceClassMethod(class IClass, sel Selector, f any) {
	rf := reflect.ValueOf(f)
	typeEncoding := _getMethodTypeEncoding(rf.Type(), true)

	imp, handle := wrapGoFuncAsMethodIMP(rf)
	_ = handle

	oldIMP := ObjectFrom(class.Ptr()).Class().ReplaceMethod(sel, imp, typeEncoding)
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

	return IMPFrom(fn), handle
}
