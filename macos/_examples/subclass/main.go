package main

import (
	"log"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type CustomView struct {
	appkit.View `objc:"NSView"`
}

func (v CustomView) AcceptsFirstResponder() bool {
	return true
}

func (v CustomView) KeyDown(event appkit.Event) {
	log.Println("Keydown:", v.Class().Name(), event.Class().Name())
}

type CustomSplitView struct {
	appkit.SplitView `objc:"NSSplitView"`
}

func (v CustomSplitView) DividerThickness() float64 {
	return 10.0
}

func (v CustomSplitView) DividerColor() appkit.Color {
	return appkit.Color_BlackColor()
}

func main() {
	log.Println("Program started.")

	CustomViewClass := objc.NewClass[CustomView](
		objc.Sel("acceptsFirstResponder"),
		objc.Sel("keyDown:"),
	)
	objc.RegisterClass(CustomViewClass)

	CustomSplitViewClass := objc.NewClass[CustomSplitView](
		objc.Sel("dividerThickness"),
		objc.Sel("dividerColor"),
	)
	objc.RegisterClass(CustomSplitViewClass)

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

		view := CustomSplitViewClass.New().InitWithFrame(frame)
		view.SetVertical(true)

		neatView := CustomViewClass.New().InitWithFrame(rectOf(0, 0, 150, 99))
		coolView := CustomViewClass.New().InitWithFrame(rectOf(10, 0, 150, 99))
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
