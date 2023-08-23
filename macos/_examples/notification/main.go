package main

import (
	"runtime"
	"time"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

func main() {
	runtime.LockOSThread()

	// notifications need a unique bundleIdentifier which we can define by
	// replacing the bundleIdentifier method on the default main bundle
	nsbundle := foundation.Bundle_MainBundle().Class()
	objc.ReplaceMethod(nsbundle, objc.Sel("bundleIdentifier"), func(_ objc.IObject) string {
		return "com.example.fake2" // change this if you miss the allow notification
	})

	objc.WithAutoreleasePool(func() {
		// this API is deprecated and we currently don't generate bindings for deprecated APIs,
		// so this is what using an API without bindings looks like.
		notif := objc.Call[objc.Object](objc.GetClass("NSUserNotification"), objc.Sel("new"))
		notif.Autorelease()
		objc.Call[objc.Void](notif, objc.Sel("setTitle:"), "Hello, world!")
		objc.Call[objc.Void](notif, objc.Sel("setInformativeText:"), "More text")

		center := objc.Call[objc.Object](objc.GetClass("NSUserNotificationCenter"), objc.Sel("defaultUserNotificationCenter"))
		objc.Call[objc.Void](center, objc.Sel("deliverNotification:"), notif)
	})

	// give notification center a moment
	<-time.After(1 * time.Second)
}
