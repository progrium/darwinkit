package main

import (
	"strconv"

	"github.com/progrium/macdriver/macos"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

func main() {
	macos.RunApp(launched)
}

func launched(app appkit.Application, delegate *appkit.ApplicationDelegate) {
	w := appkit.NewWindowWithSize(720, 440)
	objc.Retain(&w)
	w.SetTitle("Decoder")

	tabView := appkit.NewTabView()
	tabView.SetTranslatesAutoresizingMaskIntoConstraints(false)

	// add tabs
	tabView.AddTabViewItem(createNewView(1))
	tabView.AddTabViewItem(createNewView(2))

	w.SetContentView(tabView)
	w.Center()
	w.MakeKeyAndOrderFront(nil)

	delegate.SetApplicationWillFinishLaunching(func(foundation.Notification) {
		w.SetFrameAutosaveName("tab-test")
	})
	delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)

}

func createNewView(idx int) appkit.ITabViewItem {
	sv := appkit.NewStackView()
	sv.SetTranslatesAutoresizingMaskIntoConstraints(true)
	sv.AddViewInGravity(appkit.NewButtonWithTitle("button"), appkit.StackViewGravityTop)
	sv.AddViewInGravity(appkit.NewTextField(), appkit.StackViewGravityTop)
	ti := appkit.NewTabViewItem()
	ti.SetLabel("Tab" + strconv.Itoa(idx))
	ti.SetView(sv)
	return ti
}
