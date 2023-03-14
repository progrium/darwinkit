// Copyright (c) 2012-2022 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

/*
extern unsigned long GoObjc_CallTargetFrameSetup;
*/
import "C"
import (
	"math"
	"reflect"
	"strings"
	"unsafe"
)

type uint128 struct {
	Upper uintptr
	Lower uintptr
}

// arm64frame represents the layout of the
// register set once it's pushed onto the stack
// when the Objective-C runtime calls our call target.
type arm64frame struct {
	x0 uintptr
	x1 uintptr
	x2 uintptr
	x3 uintptr
	x4 uintptr
	x5 uintptr
	x6 uintptr
	x7 uintptr
	q0 uint128
	q1 uint128
	q2 uint128
	q3 uint128
	q4 uint128
	q5 uint128
	q6 uint128
	q7 uint128
}

// arm64frameFetcher implements the logic needed
// to fetch arguments from an arm64frame in the
// correct order.
type arm64frameFetcher struct {
	frame  *arm64frame
	ints   *[8]uintptr
	floats *[8]uint128
	ioff   int
	foff   int
	soff   int
}

// frameFetcher returns a new arm64frameFetcher that
// wraps an existing arm64 frame.
func frameFetcher(frame *arm64frame) arm64frameFetcher {
	ints := (*[8]uintptr)(unsafe.Pointer(frame))
	floats := (*[8]uint128)(unsafe.Pointer(&frame.q0))
	return arm64frameFetcher{
		ints:   ints,
		floats: floats,
	}
}

// Int returns the next integer argument from the arm64frame
// wrapped by the frame fetcher.
func (ff *arm64frameFetcher) Int() uintptr {
	if ff.ioff < len(ff.ints) {
		val := ff.ints[ff.ioff]
		ff.ioff++
		return val
	}
	return ff.Stack()
}

// Float returns the next floating point argument from the arm64
// frame wrapped by the frame fetcher.
func (ff *arm64frameFetcher) Float() uintptr {
	if ff.foff < len(ff.floats) {
		val := ff.floats[ff.foff].Upper
		ff.foff++
		return val
	}
	return ff.Stack()
}

// Stack returns the next stack argument from arm64
// frame wrapped by the frame fetcher.
func (ff *arm64frameFetcher) Stack() uintptr {
	panic("arm64frameFetcher: stack fetching not supported")
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
	frame := (*arm64frame)(unsafe.Pointer(p))
	fetcher := frameFetcher(frame)

	obj := object{ptr: fetcher.Int()}
	sel := stringFromSelector(unsafe.Pointer(fetcher.Int()))

	clsName := object{ptr: getObjectClass(obj).Pointer()}.className()
	// Sometimes newer versions of macOS dynamically create a "NSKVONotifying_"
	// prefixed subclass, which we won't have registered in classMap. However,
	// since it's a subclass, we can probably get away with looking up the class
	// that was inherited from based on the name part after the prefix.
	if strings.HasPrefix(clsName, "NSKVONotifying_") {
		clsName = strings.Replace(clsName, "NSKVONotifying_", "", 1)
	}
	clsInfo := classMap[clsName]
	if clsInfo.typ == nil {
		panic("unable to find registered class: " + clsName)
	}
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
		valuePtr := fetcher.Int()
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
		valuePtr := fetcher.Int()
		keyPtr := fetcher.Int()

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
		obj.SendSuper("dealloc")
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
			args = append(args, reflect.ValueOf(object{ptr: fetcher.Int()}))
			continue
		} else if typ.Implements(selectorInterfaceType) {
			sel := selector(stringFromSelector(unsafe.Pointer(fetcher.Int())))
			args = append(args, reflect.ValueOf(sel))
			continue
		}

		switch typ.Kind() {
		case reflect.Int:
			args = append(args, reflect.ValueOf(int(fetcher.Int())))
		case reflect.Int8:
			args = append(args, reflect.ValueOf(int8(fetcher.Int())))
		case reflect.Int16:
			args = append(args, reflect.ValueOf(int16(fetcher.Int())))
		case reflect.Int32:
			args = append(args, reflect.ValueOf(int32(fetcher.Int())))
		case reflect.Int64:
			args = append(args, reflect.ValueOf(int64(fetcher.Int())))

		case reflect.Uint8:
			args = append(args, reflect.ValueOf(uint8(fetcher.Int())))
		case reflect.Uint16:
			args = append(args, reflect.ValueOf(uint16(fetcher.Int())))
		case reflect.Uint32:
			args = append(args, reflect.ValueOf(uint32(fetcher.Int())))
		case reflect.Uint64:
			args = append(args, reflect.ValueOf(uint64(fetcher.Int())))
		case reflect.Uintptr:
			args = append(args, reflect.ValueOf(fetcher.Int()))

		case reflect.Float32:
			args = append(args, reflect.ValueOf(math.Float32frombits(uint32(fetcher.Float()))))
		case reflect.Float64:
			args = append(args, reflect.ValueOf(math.Float64frombits(uint64(fetcher.Float()))))

		case reflect.Bool:
			val := fetcher.Int() != 0
			args = append(args, reflect.ValueOf(val))

		case reflect.Ptr:
			ptrAddr := unsafe.Pointer(uintptr(fetcher.Int()))
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
		case reflect.Float32:
			frame.q0.Upper = uintptr(math.Float32bits(float32(val.Float())))
			return 1
		case reflect.Float64:
			frame.q0.Upper = uintptr(math.Float64bits(val.Float()))
			return 1
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
