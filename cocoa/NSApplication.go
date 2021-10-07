package cocoa

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

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
	gen_NSApplication
}

func NSApplication_New() NSApplication {
	return NSApplication_alloc().Init_asNSApplication()
}

func NSApp() NSApplication {
	return NSApplication_sharedApplication()
}

func NSApp_WithDidLaunch(cb func(notification objc.Object)) NSApplication {
	DefaultDelegateClass.AddMethod("applicationDidFinishLaunching:", func(_, notification objc.Object) {
		cb(notification)
	})
	app := NSApp()
	app.SetDelegate(DefaultDelegate)
	return app
}

func (app NSApplication) Terminate() {
	app.Terminate_(nil)
}

func (app NSApplication) SetDelegate(delegate objc.Object) {
	app.SetDelegate_(delegate)
}

func (app NSApplication) Delegate() objc.Object {
	return app.gen_NSApplication.Delegate()
}

func (app NSApplication) SetMainMenu(menu NSMenu) {
	app.SetMainMenu_(menu)
}

func (app NSApplication) SetActivationPolicy(policy int) {
	app.SetActivationPolicy_(core.NSInteger(policy))
}

func (app NSApplication) ActivateIgnoringOtherApps(flag bool) {
	app.ActivateIgnoringOtherApps_(flag)
}
