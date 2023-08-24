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

func (v TestView) dividerThickness() float64 {
	return 10.0
}

func (v TestView) dividerColor() appkit.Color {
	return appkit.Color_BlackColor()
}

func main() {
	log.Println("Program started.")

	// Create a SplitView subclass using AllocateClass
	SplitViewClass := objc.AllocateClass(objc.GetClass("NSSplitView"), "TestSplitView", 0)
	objc.AddMethod(SplitViewClass, objc.Sel("acceptsFirstResponder"), (TestView).acceptsFirstResponder)
	objc.AddMethod(SplitViewClass, objc.Sel("keyDown:"), (TestView).keyDown)

	// Implement these methods for the dividerThickness and dividerColor properties on the subclass
	objc.AddMethod(SplitViewClass, objc.Sel("dividerThickness"), (TestView).dividerThickness)
	objc.AddMethod(SplitViewClass, objc.Sel("dividerColor"), (TestView).dividerColor)

	objc.RegisterClass(SplitViewClass)

	ViewClass := objc.AllocateClass(objc.GetClass("NSView"), "TestView", 0)
	objc.AddMethod(ViewClass, objc.Sel("acceptsFirstResponder"), (TestView).acceptsFirstResponder)
	objc.AddMethod(ViewClass, objc.Sel("keyDown:"), (TestView).keyDown)
	objc.RegisterClass(ViewClass)

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

		view := appkit.SplitViewFrom(SplitViewClass.CreateInstance(0).Ptr()).InitWithFrame(frame)
		view.SetVertical(true)

		neatView := appkit.ViewFrom(ViewClass.CreateInstance(0).Ptr()).InitWithFrame(rectOf(0, 0, 150, 99))
		coolView := appkit.ViewFrom(ViewClass.CreateInstance(0).Ptr()).InitWithFrame(rectOf(10, 0, 150, 99))
		neatView.AddSubview(appkit.NewLabel("NEAT"))
		coolView.AddSubview(appkit.NewLabel("COOL"))

		// Add subviews
		view.AddArrangedSubview(neatView)
		view.AddArrangedSubview(coolView)

		win.SetContentView(view)
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
