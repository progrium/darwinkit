// Copyright (c) 2013 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package uikit

import (
	"unsafe"

	. "github.com/progrium/macdriver/pkg/ns/Foundation"
)

/*
#include <stdlib.h>

extern int UIApplicationMain(int argc, char *argv[], char *appCls, char *appDelegateCls);

int GOUIKit_UIApplicationMain(void *appCls, void *appDelegateCls) {
	return UIApplicationMain(0, NULL, appCls, appDelegateCls);
}
*/
import "C"

func UIApplicationMain(appClass string, appDelegateClass string) int {
	appClassNSString := NSStringFromString(appClass)
	appDelegateClassNSString := NSStringFromString(appDelegateClass)
	return int(C.GOUIKit_UIApplicationMain(unsafe.Pointer(appClassNSString.Pointer()), unsafe.Pointer(appDelegateClassNSString.Pointer())))
}
