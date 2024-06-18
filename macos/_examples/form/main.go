package main

import (
	"github.com/progrium/darwinkit/helper/layout"
	"github.com/progrium/darwinkit/helper/widgets"
	"github.com/progrium/darwinkit/macos"
	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
	"github.com/progrium/darwinkit/objc"
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

	w.ContentView().AddSubview(fv.GridView)
	layout.PinEdgesToSuperView(fv, foundation.EdgeInsets{Top: 10, Left: 10, Bottom: 10, Right: 10})

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
}
