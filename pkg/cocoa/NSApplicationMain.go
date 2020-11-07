// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cocoa

/*
extern int NSApplicationMain(int argc, const char *argv[]);
void GoAppKit_NSApplicationMain() {
	NSApplicationMain(0, (void *)0);
}
*/
import "C"

func NSApplicationMain() {
	C.GoAppKit_NSApplicationMain()
}
