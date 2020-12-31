// Copyright 2013 Mikkel Krautz. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package variadic

import "unsafe"

/*
#include <stdarg.h>

void *arg0(void *arg0, ...) {
	return arg0;
}

void *argN(int n, ...) {
	va_list v;
	va_start(v, n);
	int skip = n-1;
	while (skip > 0) {
		va_arg(v, void *);
		skip--;
	}
	void *val = va_arg(v, void *);
	return val;
}

void *arg0fn() {
	return arg0;
}

void *argNfn() {
	return argN;
}

float floatret(void *arg0) {
	return 3.141592f;
}

void *floatretfn() {
	return floatret;
}

double doubleret(void *arg0) {
	return 3.141592;
}

void *doubleretfn() {
	return doubleret;
}
*/
import "C"

func arg0fn() unsafe.Pointer {
	return unsafe.Pointer(C.arg0fn())
}

func argNfn() unsafe.Pointer {
	return unsafe.Pointer(C.argNfn())
}

func floatret() unsafe.Pointer {
	return unsafe.Pointer(C.floatretfn())
}

func doubleret() unsafe.Pointer {
	return unsafe.Pointer(C.doubleretfn())
}