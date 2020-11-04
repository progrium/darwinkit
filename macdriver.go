package macdriver

import (
	"fmt"
	"log"

	"github.com/progrium/macdriver/pkg/cocoa"
	appkit "github.com/progrium/macdriver/pkg/cocoa/AppKit"
	foundation "github.com/progrium/macdriver/pkg/cocoa/Foundation"
	"github.com/progrium/macdriver/pkg/objc"
)

func init() {
	// defer runtime.LockOSThread()
	c := objc.NewClass(AppDelegate{})
	c.AddMethod("applicationDidFinishLaunching:", (*AppDelegate).ApplicationDidFinishLaunching)
	c.AddMethod("applicationShouldTerminateAfterLastWindowClosed:", (*AppDelegate).ApplicationShouldTerminateAfterLastWindowClosed)
	c.AddMethod("applicationWillFinishLaunching:", (*AppDelegate).ApplicationWillFinishLaunching)
	c.AddMethod("foobar:", (*AppDelegate).Foobar)
	objc.RegisterClass(c)
}

type AppDelegate struct {
	objc.Object `objc:"GOAppDelegate : NSObject"`
	Window      appkit.NSWindow
}

func (delegate *AppDelegate) Foobar() {
	log.Println("FOOBAR")
}

func (delegate *AppDelegate) ApplicationShouldTerminateAfterLastWindowClosed(sender objc.Object) bool {
	return true
}

func (delegate *AppDelegate) ApplicationWillFinishLaunching(notification objc.Object) {
	//log.Printf("ApplicationWillFinishLaunching! %v", notification)

	appkit.NSApp().SetActivationPolicy(0)
}

func (delegate *AppDelegate) ApplicationDidFinishLaunching(notification objc.Object) {
	//log.Printf("ApplicationDidFinishLaunching! %v", notification)

	delegate.Window = appkit.NSWindow_Init(cocoa.Rect(200.0, 200.0, 600.0, 400.0),
		appkit.NSTitledWindowMask|appkit.NSClosableWindowMask|appkit.NSMiniaturizableWindowMask,
		appkit.NSBackingStoreBuffered,
		false,
	)
	delegate.Window.SetTitle("Hello world!!")
	delegate.Window.MakeKeyAndOrderFront(delegate.Window)

	appkit.NSApp().SetMainMenu(MakeMenu())

	// content := objc.GetClass("UNMutableNotificationContent").Alloc().Init()
	// content.SendMsg("setTitle:", cocoa.String("Title"))
	// content.SendMsg("setBody:", cocoa.String("This is the body"))
	// trigger := objc.GetClass("UNTimeIntervalNotificationTrigger").SendMsg("triggerWithTimeInterval:repeats:", 0, false)
	// req := objc.GetClass("UNNotificationRequest").SendMsg("requestWithIdentifier:content:trigger:", cocoa.String("test"), content, nil)
	// objc.GetClass("UNUserNotificationCenter").SendMsg("currentNotificationCenter").SendMsg("addNotificationRequest:", req)

	// notif := objc.GetClass("NSUSerNotification").Alloc().Init()
	// notif.SendMsg("setTitle:", cocoa.String("Hello there"))
	// objc.GetClass("NSUserNotificationCenter").SendMsg("defaultUserNotificationCenter").SendMsg("scheduleNotification:", notif)
}

// regular w/ titlebar (w/ auto dark mode)
// bg: solid, transparent, translucent
// corners: radius
// titlebar: regular, minimal, none

// fixed size
// min-max size
// resizable, resizable to grid?
// always on top
// hide, minimize, maximize

func Run() {
	foundation.NSAutoreleasePool_New()

	app := appkit.NSApp()
	delegate := objc.GetClass("GOAppDelegate").Alloc().Init()

	statusBarItem := objc.GetClass("NSStatusBar").SendMsg("systemStatusBar").SendMsg("statusItemWithLength:", -1.0)
	statusBarItem.SendMsg("button").SendMsg("setTitle:", cocoa.String("Hello world"))
	statusBarItem.SendMsg("setTarget:", delegate)
	statusBarItem.SendMsg("setAction:", objc.GetSelector("foobar:"))

	app.SetDelegate(delegate)
	fmt.Println("running...")
	app.Run()
}

func MakeMenu() appkit.NSMenu {
	mainMenu := appkit.NSMenu_New()
	mainMenu.AutoRelease()

	mainAppItem := appkit.NSMenuItem_New()
	mainAppItem.AutoRelease()

	mainFileItem := appkit.NSMenuItem_New()
	mainFileItem.AutoRelease()

	mainMenu.AddItem(mainAppItem)
	mainMenu.AddItem(mainFileItem)

	fileMenu := appkit.NSMenu_Init("File")
	fileMenu.AutoRelease()
	mainFileItem.SetSubmenu(fileMenu)

	appMenu := appkit.NSMenu_Init("App")
	appMenu.AutoRelease()
	mainAppItem.SetSubmenu(appMenu)

	quitItem := appkit.NSMenuItem_New()
	quitItem.SendMsg("setKeyEquivalent:", cocoa.String("q"))
	quitItem.SendMsg("setTitle:", cocoa.String("Quit"))
	quitItem.SendMsg("setAction:", objc.GetSelector("terminate:"))
	quitItem.AutoRelease()

	quitItem2 := appkit.NSMenuItem_New()
	quitItem2.SendMsg("setTitle:", cocoa.String("Foobar"))
	quitItem2.SendMsg("setAction:", objc.GetSelector("foobar:"))
	quitItem2.AutoRelease()

	fileMenu.AddItem(quitItem2)
	appMenu.AddItem(quitItem)

	return mainMenu
}
