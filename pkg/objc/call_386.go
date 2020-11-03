// Copyright (c) 2013 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

/*
extern unsigned long GoObjc_CallTargetFrameSetup;
extern void GoObjc_Fst(void *);
*/
import "C"
import (
	"math"
	"reflect"
	"unsafe"
)

// i386frameFetcher implements the logic needed
// to fetch arguments from a i386 stack frame.
type i386frameFetcher struct {
	args *[100]uintptr
	off  int
}

// frameFetcher returns a new i386frameFetcher that
// wraps an i386 stack frame.
func frameFetcher(stack uintptr) i386frameFetcher {
	frame := (*[100]uintptr)(unsafe.Pointer(stack))
	return i386frameFetcher{
		args: frame,
		off:  0,
	}
}

// Next returns the next argument from the wrapped
// i386 stack frame.
func (ff *i386frameFetcher) Next() uintptr {
	if ff.off < len(ff.args) {
		val := ff.args[ff.off]
		ff.off++
		return val
	}
	panic("invalid fetch")
}

// methodCallTarget returns a pointer to the entry point
// that the Objective-C runtime must call to reach an
// exported method from Go.
func methodCallTarget() unsafe.Pointer {
	return unsafe.Pointer(&C.GoObjc_CallTargetFrameSetup)
}

// setIBOutletValue attempts to assign the Objective-C object represented
// by 'value' to the field named 'name' in the Go struct represented by 'obj'.
//
// The function sends the 'retain' message to the Objective-C object represented
// by 'value' if the assignment was successful.
//
// If the assignment operation fails, this function will raise a runtime panic.
func setIBOutletValue(obj reflect.Value, name string, value Object) {
	// Find an IBOutlet with the name keyName.
	val := obj.Elem()
	typ := val.Type()
	fieldIdx := -1
	for i := 0; i < typ.NumField(); i++ {
		field := typ.Field(i)
		if field.Tag.Get("objc") == "IBOutlet" && field.Type.Implements(objectInterfaceType) {
			if field.Name == name {
				fieldIdx = i
				break
			}
		}
	}

	// If we couldn't find a matching field, simply panic.
	// We should only run into this is there is a bug in the package.
	if fieldIdx == -1 {
		panic("objc: bad setter for IBOutlet field '" + name + "'")
	}

	fieldVal := val.Field(fieldIdx)
	fieldVal.Set(reflect.ValueOf(value))
	value.Retain()
}

//export goMethodCallEntryPoint
func goMethodCallEntryPoint(p uintptr) uintptr {
	fetcher := frameFetcher(p)

	// Hijack the stack pointer to perform a little
	// float-return magic. We have to use the 387 stack
	// to return float values.
	//
	// We re-use the stack slot for this function's
	// p parameter as a pointer to a 64-bit 'double precision'
	// floating point value that we will move to the st0
	// register (where floating point return values should go.)
	//
	// For non-float cases, this slot should be zeroed out
	// so we know not to touch the 387 floating point stack.
	stk := (*[10]uintptr)(unsafe.Pointer(p - 4))
	stk[0] = 0

	// Skip, skip.
	fetcher.Next()
	fetcher.Next()

	obj := object{ptr: fetcher.Next()}
	sel := stringFromSelector(unsafe.Pointer(fetcher.Next()))

	clsName := object{ptr: getObjectClass(obj).Pointer()}.className()
	clsInfo := classMap[clsName]
	method := clsInfo.MethodForSelector(sel)

	// Check if we have an internal pointer set for this object.
	// If not, make it happen.
	internalPtr := obj.internalPointer()
	if internalPtr == nil {
		// Allocate the Go struct.
		val := reflect.New(clsInfo.typ)
		ptr := unsafe.Pointer(val.Pointer())
		// Add a reference in the classInfo. This ensures we have
		// a reference to our Go struct somewhere in Go land, which
		// makes the garbage collector not collect it under our feet.
		clsInfo.AddRef(ptr)
		// Set the internalPointer so we can easily access
		// the instance's Go struct pointer.
		obj.setInternalPointer(ptr)
		internalPtr = ptr
		// Finally, update the Go struct's embedded objc.Object to
		// point to the actual Objective-C instance.
		structVal := val.Elem()
		if structVal.Kind() == reflect.Struct {
			objectVal := structVal.FieldByName("Object")
			if objectVal.IsValid() {
				objectVal.Set(reflect.ValueOf(obj))
			}
		}
	}
	objVal := reflect.NewAt(clsInfo.typ, internalPtr)

	// Check if the invoked selector is a setter for IBOutlets.
	if _, isSetter := clsInfo.setters[sel]; isSetter {
		valuePtr := fetcher.Next()
		keyName := sel[3:]                    // strip 'set'
		keyName = keyName[0 : len(keyName)-1] // strip ':'
		setIBOutletValue(objVal, keyName, object{ptr: valuePtr})
		return 0
	}

	// Our own internal override for setValue:forKey: in order
	// to support key-value coding. (For IBOutlets on iOS)
	if sel == "setValue:forKey:" && method == nil {
		// We only support Object values, so fetching
		// Ints here is OK.
		valuePtr := fetcher.Next()
		keyPtr := fetcher.Next()

		// We don't export any NSString-based functionality
		// in package objc, except for the String() method
		// on object.  It calls the object's decription method,
		// which for NSStrings returns the string itself (or at
		// least something that has been good enough for now!).
		keyName := object{ptr: keyPtr}.String()

		setIBOutletValue(objVal, keyName, object{ptr: valuePtr})
		return 0
	}

	// The default dealloc implementation. This is called
	// if a class doesn't register its own custom dealloc
	// method.
	if sel == "dealloc" && method == nil {
		clsInfo.RemoveRef(internalPtr)
		obj.SendSuperMsg("dealloc")
		return 0
	}

	methodVal := reflect.ValueOf(method)

	// First argument should point to the Go method's proper receiver.
	// That's stored in the internalPointer, so fetch that.
	args := []reflect.Value{objVal}

	// Take care of the rest of the arguments
	mt := reflect.TypeOf(method)
	for i := 1; i < mt.NumIn(); i++ {
		typ := mt.In(i)

		if typ.Implements(objectInterfaceType) {
			args = append(args, reflect.ValueOf(object{ptr: fetcher.Next()}))
			continue
		} else if typ.Implements(selectorInterfaceType) {
			sel := selector(stringFromSelector(unsafe.Pointer(fetcher.Next())))
			args = append(args, reflect.ValueOf(sel))
			continue
		}

		switch typ.Kind() {
		case reflect.Int:
			args = append(args, reflect.ValueOf(int(fetcher.Next())))
		case reflect.Int8:
			args = append(args, reflect.ValueOf(int8(fetcher.Next())))
		case reflect.Int16:
			args = append(args, reflect.ValueOf(int16(fetcher.Next())))
		case reflect.Int32:
			args = append(args, reflect.ValueOf(int32(fetcher.Next())))
		case reflect.Int64:
			v1 := uint64(fetcher.Next())
			v2 := uint64(fetcher.Next())
			u64 := (v1 & 0xffffffff) | ((v2 << 32) & 0xffffffff00000000)
			args = append(args, reflect.ValueOf(int64(u64)))

		case reflect.Uint8:
			args = append(args, reflect.ValueOf(uint8(fetcher.Next())))
		case reflect.Uint16:
			args = append(args, reflect.ValueOf(uint16(fetcher.Next())))
		case reflect.Uint32:
			args = append(args, reflect.ValueOf(uint32(fetcher.Next())))
		case reflect.Uint64:
			v1 := uint64(fetcher.Next())
			v2 := uint64(fetcher.Next())
			u64 := (v1 & 0xffffffff) | ((v2 << 32) & 0xffffffff00000000)
			args = append(args, reflect.ValueOf(u64))
		case reflect.Uintptr:
			args = append(args, reflect.ValueOf(fetcher.Next()))

		case reflect.Float32:
			f32 := math.Float32frombits(uint32(fetcher.Next()))
			args = append(args, reflect.ValueOf(f32))
		case reflect.Float64:
			v1 := uint64(fetcher.Next())
			v2 := uint64(fetcher.Next())
			u64 := (v1 & 0xffffffff) | ((v2 << 32) & 0xffffffff00000000)
			args = append(args, reflect.ValueOf(math.Float64frombits(u64)))

		case reflect.Bool:
			val := fetcher.Next() != 0
			args = append(args, reflect.ValueOf(val))

		case reflect.Ptr:
			ptrAddr := unsafe.Pointer(uintptr(fetcher.Next()))
			args = append(args, reflect.NewAt(typ.Elem(), ptrAddr))

		default:
			panic("call: unhandled arg")
		}
	}

	retVals := methodVal.Call(args)

	// If a custom dealloc method has been registered, we
	// still need to remove the reference to our Go struct
	// pointer in order for the GC to kick in. Do that now.
	if sel == "dealloc" {
		clsInfo.RemoveRef(internalPtr)
	}

	if len(retVals) > 0 {
		val := retVals[0]
		switch val.Kind() {
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			return uintptr(val.Int())
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
			return uintptr(val.Uint())
		case reflect.Bool:
			if val.Bool() {
				return 1
			} else {
				return 0
			}
		case reflect.Float32, reflect.Float64:
			f64 := val.Float()
			stk[0] = p + 8
			u64 := math.Float64bits(f64)
			stk[3] = uintptr(u64 & 0xffffffff)
			stk[4] = uintptr((u64 >> 32) & 0xffffffff)
			return 0
		case reflect.Interface:
			if obj, ok := val.Interface().(Object); ok {
				return obj.Pointer()
			}
			panic("call: bad interface return value")
		default:
			panic("call: unknown return value")
		}
	}

	return 0
}
