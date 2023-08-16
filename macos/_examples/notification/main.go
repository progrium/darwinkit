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

	nsbundle := foundation.Bundle_MainBundle().GetClass()
	objc.ReplaceMethod(nsbundle, objc.GetSelector("bundleIdentifier"), func(_ objc.IObject) string {
		return "com.example.fake2" // change this if you miss the allow notification
	})

	ad := &appkit.ApplicationDelegate{}
	ad.SetApplicationDidFinishLaunching(func(foundation.Notification) {

		notif := objc.CallMethod[objc.Object](objc.GetClass("NSUserNotification"), objc.GetSelector("new"))
		notif.Autorelease()
		objc.CallMethod[objc.Void](notif, objc.GetSelector("setTitle:"), "Hello, world!")
		objc.CallMethod[objc.Void](notif, objc.GetSelector("setInformativeText:"), "More text")

		center := objc.CallMethod[objc.Object](objc.GetClass("NSUserNotificationCenter"), objc.GetSelector("defaultUserNotificationCenter"))
		objc.CallMethod[objc.Void](center, objc.GetSelector("deliverNotification:"), notif)

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
