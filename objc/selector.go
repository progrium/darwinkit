// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package objc

/*
#cgo LDFLAGS: -lobjc
#include <objc/runtime.h>

void *GoObjc_RegisterSelector(char *name) {
	return (void *) sel_registerName(name);
}

char *GoObjc_TypeInfoForMethod(void *cls, void *sel) {
	Method m = class_getInstanceMethod(cls, sel);
	return (char *) method_getTypeEncoding(m);
}

char *GoObjc_SelectorToString(void *sel) {
	return (char *) sel;
}
*/
import "C"
import (
	"sync"
	"unsafe"
)

// A Selector represents an Objective-C method selector.
type Selector interface {
	// Selector returns a string representation of
	// a selector.
	Selector() string

	// String returns the same string as Selector does.
	// It is only implemented to implement the Stringer
	// interface.
	String() string
}

// Type selector is the underlying implementation
// of the Selector interface. It is represented as
// a Go string.
type selector string

// Selector implements the Selector method of the
// Selector interface.
func (sel selector) Selector() string {
	return string(sel)
}

// String implements the String method of the
// Selector interface.
func (sel selector) String() string {
	return sel.Selector()
}

// GetSelector looks up a Selector by name.
func Sel(name string) Selector {
	return selector(name)
}

// deprecated
func GetSelector(name string) Selector {
	return Sel(name)
}

func RegisterSelector(name string) unsafe.Pointer {
	return selectorWithName(name)
}

// selectorWithName looks up a selector by name.
func selectorWithName(name string) unsafe.Pointer {
	return C.GoObjc_RegisterSelector(C.CString(name))
}

// stringFromSelector converts a selector to a Go string
func stringFromSelector(sel unsafe.Pointer) string {
	return C.GoString(C.GoObjc_SelectorToString(sel))
}

// simplifyTypeInfo returns a simplified typeInfo representation
// with C specifiers and stack information stripped out.
func simplifyTypeInfo(typeInfo string) string {
	ti := typeInfo
	sti := []byte{}
	for i := 0; i < len(ti); i++ {
		if ti[i] >= '0' && ti[i] <= '9' {
			continue
		}
		if string(ti[i]) == encConst {
			continue
		}
		// fixme(mkrautz): What is V? The NSObject release method uses V.
		if ti[i] == 'V' {
			continue
		}
		sti = append(sti, ti[i])
	}
	return string(sti)
}

var methodTypeInfo = struct {
	sync.RWMutex
	e map[[2]unsafe.Pointer]string
}{
	e: map[[2]unsafe.Pointer]string{},
}

// simpleTypeInfoForMethod fetches the type info for the method
// identified by obj's class and the given selector and returns
// it in a simplified form produced by the simplifyTypeInfo function.
func simpleTypeInfoForMethod(obj Object, selector string) string {
	cls := unsafe.Pointer(getObjectClass(obj).Pointer())
	sel := selectorWithName(selector)
	key := [2]unsafe.Pointer{cls, sel}

	methodTypeInfo.RLock()
	ti, ok := methodTypeInfo.e[key]
	methodTypeInfo.RUnlock()
	if ok {
		return ti
	}

	methodTypeInfo.Lock()
	defer methodTypeInfo.Unlock()

	ti, ok = methodTypeInfo.e[key]
	if ok {
		return ti
	}

	ti = C.GoString(C.GoObjc_TypeInfoForMethod(cls, sel))
	ti = simplifyTypeInfo(ti)
	methodTypeInfo.e[key] = ti
	return ti
}
