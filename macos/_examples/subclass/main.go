package main

import (
	"log"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type TestView struct {
	objc.Object
}

func (v TestView) acceptsFirstResponder() bool {
	return true
}

func (v TestView) keyDown(event appkit.Event) {
	log.Println("Keydown:", v.Class().Name(), event.Class().Name())
}

func main() {
	log.Println("Program started.")

	c := objc.AllocateClass(objc.GetClass("NSView"), "TestView", 0)
	objc.AddMethod(c, objc.Sel("acceptsFirstResponder"), (TestView).acceptsFirstResponder)
	objc.AddMethod(c, objc.Sel("keyDown:"), (TestView).keyDown)
	objc.RegisterClass(c)

	app := appkit.Application_SharedApplication()

	ad := &appkit.ApplicationDelegate{}
	ad.SetApplicationDidFinishLaunching(func(foundation.Notification) {

		frame := rectOf(400, 400, 300, 200)

		win := appkit.NewWindowWithContentRectStyleMaskBackingDefer(
			frame,
			appkit.WindowStyleMaskTitled|appkit.WindowStyleMaskClosable,
			appkit.BackingStoreBuffered,
			false,
		)
		objc.Retain(&win)
		win.SetTitle("Hello world")
		win.SetLevel(appkit.MainMenuWindowLevel + 2)

		view := c.CreateInstance(0)
		objc.Call[objc.Void](win, objc.Sel("setContentView:"), view)
		win.MakeKeyAndOrderFront(nil)

		log.Println("App started.")

	})
	ad.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
	app.SetDelegate(ad)
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}

func rectOf(x, y, width, height float64) foundation.Rect {
	return foundation.Rect{Origin: foundation.Point{X: x, Y: y}, Size: foundation.Size{Width: width, Height: height}}
}
