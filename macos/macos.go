package macos

import (
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
)

// Launch builds a delegate, sets it on the shared application and runs.
// It uses didLaunch to set ApplicationDidFinishLaunching on the delegate.
func Launch(didLaunch func(appkit.Application, *appkit.ApplicationDelegate)) {
	app := appkit.Application_SharedApplication()
	delegate := &appkit.ApplicationDelegate{}
	delegate.SetApplicationDidFinishLaunching(func(notification foundation.Notification) {
		didLaunch(app, delegate)
	})
	app.SetDelegate(delegate)
	app.Run()
}
