// macOS Frameworks
//
// This package also contains platform specific helpers.
package macos

import (
	"runtime"

	"github.com/progrium/darwinkit/macos/appkit"
	"github.com/progrium/darwinkit/macos/foundation"
)

// RunApp builds a delegate, sets it on the shared application, and runs the event loop.
// It uses didLaunch to set ApplicationDidFinishLaunching on the delegate.
// It also wraps run with runtime.LockOSThread() and should be called from main.
func RunApp(didLaunch func(appkit.Application, *appkit.ApplicationDelegate)) {
	runtime.LockOSThread()
	defer runtime.UnlockOSThread()
	app := appkit.Application_SharedApplication()
	delegate := &appkit.ApplicationDelegate{}
	delegate.SetApplicationDidFinishLaunching(func(notification foundation.Notification) {
		didLaunch(app, delegate)
	})
	app.SetDelegate(delegate)
	app.Run()
}
