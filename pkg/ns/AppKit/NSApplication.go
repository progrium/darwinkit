// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package appkit

/*
extern int NSApplicationMain(int argc, const char *argv[]);
void GoAppKit_NSApplicationMain() {
	NSApplicationMain(0, (void *)0);
}
*/
import "C"

import "github.com/progrium/macdriver/pkg/objc"

type NSApplication struct {
	objc.Object
}

func NSSharedApplication() NSApplication {
	return NSApplication{objc.GetClass("NSApplication").SendMsg("sharedApplication")}
}

func (app NSApplication) Run() {
	app.SendMsg("run")
}

func (app NSApplication) SetDelegate(delegate objc.Object) {
	app.SendMsg("setDelegate:", delegate)
}

func (app NSApplication) Delegate() objc.Object {
	return app.SendMsg("delegate")
}

func (app NSApplication) SetMainMenu(menu NSMenu) {
	app.SendMsg("setMainMenu:", menu)
}

func (app NSApplication) MainMenu() NSMenu {
	return NSMenu{app.SendMsg("mainMenu")}
}

func NSApplicationMain() {
	C.GoAppKit_NSApplicationMain()
}
