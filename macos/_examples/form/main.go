package main

import (
	"runtime"

	"github.com/progrium/macdriver/helper/layout"
	"github.com/progrium/macdriver/helper/widgets"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
)

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func initAndRun() {
	app := appkit.Application_SharedApplication()
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
	w := appkit.NewWindowWithSize(600, 400)
	w.SetTitle("Form")

	fv := widgets.NewFormView()
	fv.SetLabelWidth(100)
	fv.SetLabelAlignment(widgets.LabelAlignmentTrailing)
	fv.SetLabelControlSpacing(10)
	fv.AddRow("user", appkit.NewTextField())
	fv.SetLabelFont(appkit.FontClass.UserFontOfSize(13))
	fv.AddRow("password", appkit.NewSecureTextField())
	cb := appkit.NewCheckBox("")
	fv.AddRow("males", cb)
	fv.AddExpandRow()

	w.ContentView().AddSubview(fv)
	layout.PinEdgesToSuperView(fv, foundation.EdgeInsets{Top: 10, Left: 10, Bottom: 10, Right: 10})

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	ad := &appkit.ApplicationDelegate{}
	ad.SetApplicationDidFinishLaunching(func(foundation.Notification) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)
	})
	ad.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
	app.SetDelegate(ad)

	app.Run()
}

func main() {
	initAndRun()
}
