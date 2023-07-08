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
	DefaultDelegate = objc.Get("DefaultDelegate").Alloc().InitObject()
}

type NSApplication struct {
	gen_NSApplication
}

func NSApplication_New() NSApplication {
	return NSApplication_Alloc().Init()
}

func NSApp() NSApplication {
	return NSApplication_SharedApplication()
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
	app.gen_NSApplication.Terminate(nil)
}
func (app NSApplication) SetActivationPolicy(policy int) {
	app.gen_NSApplication.SetActivationPolicy(core.NSInteger(policy))
}
