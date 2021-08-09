package main

import (
	"reflect"
	"sort"
	"strings"
	"unsafe"
)

/*
#cgo CFLAGS: -x objective-c -Wno-incompatible-pointer-types
#cgo LDFLAGS: -framework Cocoa
#define OBJC2_UNAVAILABLE
#include <objc/objc.h>
#include <objc/runtime.h>
#include <Foundation/Foundation.h>

bool isaNSObject(Class cls) {
	while (cls && cls != [NSObject class]) {
		cls = class_getSuperclass(cls);
	}

	return !!cls;
}
*/
import "C"

func mostCommonSignatures() []string {
	type Count struct {
		Encoding string
		Count    int
	}
	var counts []Count
	lup := map[string]int{}

	for _, class := range getClasses() {
		if skipClass(class) {
			continue
		}

		for _, method := range getMethods(class) {
			if skipMethod(method) {
				continue
			}

			enc := simplifyTypeEncoding(C.GoString(C.method_getTypeEncoding(method)))
			i, ok := lup[enc]
			if ok {
				counts[i].Count++
				continue
			}

			// Skip unsupported types
			if strings.Contains(enc, "{") || strings.Contains(enc, "?") {
				continue
			}

			lup[enc] = len(counts)
			counts = append(counts, Count{enc, 1})
		}
	}

	top := make([]string, *itemCount)
	sort.Slice(counts, func(i, j int) bool { return counts[i].Count > counts[j].Count })
	for i := range top {
		top[i] = counts[i].Encoding
	}

	return top
}

func skipClass(class C.Class) bool {
	name := C.GoString(C.class_getName(class))
	if strings.HasPrefix(name, "_") {
		return true
	}

	return *onlyNSObject && !bool(C.isaNSObject(class))
}

func skipMethod(method C.Method) bool {
	name := C.GoString(C.sel_getName(C.method_getName(method)))
	return strings.HasPrefix(name, "_")
}

func getClasses() []C.Class {
	n := C.objc_getClassList(nil, 0)
	classes := make([]C.Class, n)
	C.objc_getClassList(&classes[0], n)
	return classes
}

func getMethods(class C.Class) []C.Method {
	var count C.uint
	ptr := C.class_copyMethodList(class, &count)
	defer C.free(unsafe.Pointer(ptr))

	var v []C.Method
	vh := (*reflect.SliceHeader)(unsafe.Pointer(&v))
	vh.Data = uintptr(unsafe.Pointer(ptr))
	vh.Len = int(count)
	vh.Cap = int(count)

	u := make([]C.Method, count)
	copy(u, v)
	return u
}

// simplifyTypeInfo returns a simplified typeInfo representation
// with C specifiers and stack information stripped out.
func simplifyTypeEncoding(typeInfo string) string {
	ti := typeInfo
	sti := []rune{}
	for _, r := range ti {
		if r >= '0' && r <= '9' {
			continue
		}
		if r == 'r' {
			continue
		}
		// fixme(mkrautz): What is V? The NSObject release method uses V.
		if r == 'V' {
			continue
		}
		sti = append(sti, r)
	}
	return string(sti)
}
