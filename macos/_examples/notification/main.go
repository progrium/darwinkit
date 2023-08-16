package main

import (
	"runtime"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

func init() {
	runtime.LockOSThread()
}

func main() {
	app := appkit.Application_SharedApplication()

	nsbundle := foundation.Bundle_MainBundle().Class()
	objc.ReplaceMethod(nsbundle, objc.Sel("bundleIdentifier"), func(_ objc.IObject) string {
		return "com.example.fake2" // change this if you miss the allow notification
	})

	ad := &appkit.ApplicationDelegate{}
	ad.SetApplicationDidFinishLaunching(func(foundation.Notification) {

		notif := objc.Call[objc.Object](objc.GetClass("NSUserNotification"), objc.Sel("new"))
		notif.Autorelease()
		objc.Call[objc.Void](notif, objc.Sel("setTitle:"), "Hello, world!")
		objc.Call[objc.Void](notif, objc.Sel("setInformativeText:"), "More text")

		center := objc.Call[objc.Object](objc.GetClass("NSUserNotificationCenter"), objc.Sel("defaultUserNotificationCenter"))
		objc.Call[objc.Void](center, objc.Sel("deliverNotification:"), notif)

		//app.Terminate(nil)

	})
	ad.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
	app.SetDelegate(ad)

	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}
