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
		notification := NSUserNotification{NSUserNotification_.Alloc().Init()}
		notification.Set("title:", core.NSString_FromString("Hello, world!"))
		notification.Set("informativeText:", core.NSString_FromString("More text"))

		center := NSUserNotificationCenter{NSUserNotificationCenter_.Send("defaultUserNotificationCenter")}
		center.Send("deliverNotification:", notification)
		notification.Release()
	})

	nsbundle := cocoa.NSBundle_Main().Send("class").(objc.Class)
	objc.TODO_RegisterClassInMap(nsbundle)

	nsbundle.AddMethod("__bundleIdentifier", func(_ objc.Object) objc.Object {
		return core.NSString_FromString("com.example.fake")
	})
	nsbundle.Swizzle("bundleIdentifier", "__bundleIdentifier")

	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
	app.Run()
}
