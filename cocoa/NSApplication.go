// Copyright (c) 2012 The 'objc' Package Authors. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package cocoa

import "github.com/progrium/macdriver/objc"

const (
	NSApplicationActivationPolicyRegular    = 0
	NSApplicationActivationPolicyAccessory  = 1
	NSApplicationActivationPolicyProhibited = 2
)

var (
	DefaultDelegate      objc.Object
	DefaultDelegateClass objc.Class

	TerminateAfterWindowsClose = true
)

func init() {
	DefaultDelegateClass = objc.NewClass("DefaultDelegate", "NSObject")
	DefaultDelegateClass.AddMethod("applicationShouldTerminateAfterLastWindowClosed:", func(notification objc.Object) bool {
		return TerminateAfterWindowsClose
	})
	objc.RegisterClass(DefaultDelegateClass)
	DefaultDelegate = objc.Get("DefaultDelegate").Alloc().Init()
}

type NSApplication struct {
	objc.Object
}

var NSApplication_ = objc.Get("NSApplication")

func NSApplication_New() NSApplication {
	return NSApplication{NSApplication_.Alloc().Init()}
}

func NSApp() NSApplication {
	return NSApplication{NSApplication_.Send("sharedApplication")}
}

func NSApp_WithDidLaunch(cb func(notification objc.Object)) NSApplication {
	DefaultDelegateClass.AddMethod("applicationDidFinishLaunching:", cb)
	app := NSApp()
	app.SetDelegate(DefaultDelegate)
	return app
}

func (app NSApplication) Run() {
	app.Send("run")
}

func (app NSApplication) Terminate() {
	app.Send("terminate:", nil)
}

func (app NSApplication) SetDelegate(delegate objc.Object) {
	app.Send("setDelegate:", delegate)
}

func (app NSApplication) Delegate() objc.Object {
	return app.Send("delegate")
}

func (app NSApplication) SetMainMenu(menu NSMenu) {
	app.Send("setMainMenu:", menu)
}

func (app NSApplication) SetActivationPolicy(policy int) {
	app.Send("setActivationPolicy:", policy)
}

func (app NSApplication) ActivateIgnoringOtherApps(flag bool) {
	app.Send("activateIgnoringOtherApps:", flag)
}

func (app NSApplication) MainMenu() NSMenu {
	return NSMenu{app.Send("mainMenu")}
}
