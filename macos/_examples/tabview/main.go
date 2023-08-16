package main

import (
	"runtime"
	"strconv"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
)

// Arrange that main.main runs on main thread.
func init() {
	runtime.LockOSThread()
}

func initAndRun() {
	app := appkit.Application_SharedApplication()
	w := appkit.NewWindowWithSize(720, 440)
	w.SetTitle("Decoder")

	tabView := appkit.NewTabView()
	tabView.SetTranslatesAutoresizingMaskIntoConstraints(false)

	// add tabs
	tabView.AddTabViewItem(createNewView(1))
	tabView.AddTabViewItem(createNewView(2))

	w.SetContentView(tabView)

	ad := &appkit.ApplicationDelegate{}
	ad.SetApplicationDidFinishLaunching(func(foundation.Notification) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)
	})
	ad.SetApplicationWillFinishLaunching(func(foundation.Notification) {
		w.SetFrameAutosaveName("tab-test")
	})
	ad.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
	app.SetDelegate(ad)

	w.Center()
	w.MakeKeyAndOrderFront(nil)
	app.Run()
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

func main() {
	initAndRun()
}
