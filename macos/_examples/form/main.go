package main

import (
	"github.com/progrium/macdriver/helper/layout"
	"github.com/progrium/macdriver/helper/widgets"
	"github.com/progrium/macdriver/macos"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

func main() {
	macos.RunApp(launched)
}

func launched(app appkit.Application, delegate *appkit.ApplicationDelegate) {
	w := appkit.NewWindowWithSize(600, 400)
	objc.Retain(&w)
	w.SetTitle("Form")

	fv := widgets.NewFormView()
	fv.SetLabelWidth(100)
	fv.SetLabelAlignment(widgets.LabelAlignmentTrailing)
	fv.SetLabelControlSpacing(10)
	fv.AddRow("user", appkit.NewTextField())
	fv.SetLabelFont(appkit.Font_UserFontOfSize(13))
	fv.AddRow("password", appkit.NewSecureTextField())
	cb := appkit.NewCheckBox("")
	fv.AddRow("males", cb)
	fv.AddExpandRow()

	w.ContentView().AddSubview(fv)
	layout.PinEdgesToSuperView(fv, foundation.EdgeInsets{Top: 10, Left: 10, Bottom: 10, Right: 10})

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
}
