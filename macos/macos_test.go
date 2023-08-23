package macos_test

import (
	"fmt"
	"runtime"

	"github.com/progrium/macdriver/macos"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/webkit"
	"github.com/progrium/macdriver/objc"
)

// You normally want to start an application using RunApp.
// This is a simple webview window app. Note that we need to retain the window object.
func Example_mainWithRunApp() {
	// RunApp should only be called from main on the startup thread
	macos.RunApp(func(app appkit.Application, delegate *appkit.ApplicationDelegate) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)

		url := foundation.URL_URLWithString("http://progrium.com")
		req := foundation.NewURLRequestWithURL(url)
		frame := foundation.Rect{Size: foundation.Size{1440, 900}}

		config := webkit.NewWebViewConfiguration()
		wv := webkit.NewWebViewWithFrameConfiguration(frame, config)
		wv.LoadRequest(req)

		w := appkit.NewWindowWithContentRectStyleMaskBackingDefer(frame,
			appkit.ClosableWindowMask|appkit.TitledWindowMask,
			appkit.BackingStoreBuffered, false)
		objc.Retain(&w)
		w.SetContentView(wv)
		w.MakeKeyAndOrderFront(w)
		w.Center()

		delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
			return true
		})
	})
}

// You can use appkit APIs directly to start an application, but you need to
// lock the main thread and build your own application delegate if desired.
func Example_mainWithoutRunApp() {
	runtime.LockOSThread()
	app := appkit.Application_SharedApplication()
	delegate := &appkit.ApplicationDelegate{}
	delegate.SetApplicationDidFinishLaunching(func(notification foundation.Notification) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)

		// ... app started code ...

	})
	app.SetDelegate(delegate)
	app.Run()
}

// You can use some Apple frameworks without starting an application,
// but you need to lock the main thread and wrap your code with an autorelease pool.
func Example_mainWithoutApplication() {
	runtime.LockOSThread()
	objc.WithAutoreleasePool(func() {

		ws := appkit.Workspace_SharedWorkspace()

		for _, app := range ws.RunningApplications() {
			fmt.Println(app.LocalizedName())
		}

	})
}
