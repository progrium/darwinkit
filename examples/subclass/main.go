package main

import (
	"log"

	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type TestView struct {
	objc.Object `objc:"TestView : NSView"`
}

func (v *TestView) acceptsFirstResponder() bool {
	return true
}

func (v *TestView) keyDown(event objc.Object) {
	log.Println("Keydown!!")
}

func main() {
	log.Println("Program started.")

	c := objc.NewClassFromStruct(TestView{})
	c.AddMethod("acceptsFirstResponder", (*TestView).acceptsFirstResponder)
	c.AddMethod("keyDown:", (*TestView).keyDown)
	objc.RegisterClass(c)

	app := cocoa.NSApp_WithDidLaunch(func(n objc.Object) {
		frame := core.Rect(400, 400, 300, 200)

		win := cocoa.NSWindow_Init(
			frame,
			cocoa.NSTitledWindowMask|
				cocoa.NSClosableWindowMask,
			cocoa.NSBackingStoreBuffered,
			false,
		)
		win.Retain()
		win.SetTitle("Hello world")
		win.SetLevel(cocoa.NSMainMenuWindowLevel + 2)

		view := objc.Get("TestView").Alloc().InitObject()
		win.Set("ContentView:", view)
		win.MakeKeyAndOrderFront(nil)

		log.Println("App started.")
	})
	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}
