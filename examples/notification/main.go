package main

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type NSUserNotification struct {
	objc.Object
}

var NSUserNotification_ = objc.Get("NSUserNotification")

type NSUserNotificationCenter struct {
	objc.Object
}

var NSUserNotificationCenter_ = objc.Get("NSUserNotificationCenter")

func main() {
	app := cocoa.NSApp_WithDidLaunch(func(_ objc.Object) {
		notification := NSUserNotification{NSUserNotification_.Alloc().InitObject()}
		notification.Set("title:", core.String("Hello, world!"))
		notification.Set("informativeText:", core.String("More text"))

		center := NSUserNotificationCenter{NSUserNotificationCenter_.Send("defaultUserNotificationCenter")}
		center.Send("deliverNotification:", notification)
		notification.Release()
	})

	nsbundle := cocoa.NSBundle_Main().Class()

	nsbundle.AddMethod("__bundleIdentifier", func(_ objc.Object) objc.Object {
		return core.String("com.example.fake")
	})
	nsbundle.Swizzle("bundleIdentifier", "__bundleIdentifier")

	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}
