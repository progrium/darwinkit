package macdriver

import (
	"fmt"
	"log"
	"runtime"

	"github.com/progrium/macdriver/pkg/cocoa"
	"github.com/progrium/macdriver/pkg/core"
	"github.com/progrium/macdriver/pkg/objc"
)

var effect objc.Object
var webview core.WKWebView

func init() {
	defer runtime.LockOSThread()

	c := objc.NewClass(AppDelegate{})
	c.AddMethod("applicationDidFinishLaunching:", (*AppDelegate).ApplicationDidFinishLaunching)
	c.AddMethod("applicationShouldTerminateAfterLastWindowClosed:", (*AppDelegate).ApplicationShouldTerminateAfterLastWindowClosed)
	c.AddMethod("applicationWillFinishLaunching:", (*AppDelegate).ApplicationWillFinishLaunching)
	c.AddMethod("foobar:", (*AppDelegate).Foobar)
	objc.RegisterClass(c)

}

type AppDelegate struct {
	objc.Object `objc:"AppDelegate : NSObject"`
}

func (delegate *AppDelegate) Foobar() {
	log.Println("FOOBAR")
	url := core.NSURL_Init("http://localhost:8080/bgtest.html")
	req := core.NSURLRequest_Init(url)
	webview.LoadRequest(req)
	webview.Set("opaque:", false)
	webview.Set("backgroundColor:", objc.Get("NSColor").Get("clearColor"))
	webview.Send("setValue:forKey:", objc.Get("NSNumber").Send("numberWithBool:", false), core.String("drawsBackground"))
}

func (delegate *AppDelegate) ApplicationShouldTerminateAfterLastWindowClosed(sender objc.Object) bool {
	return true
}

func (delegate *AppDelegate) ApplicationWillFinishLaunching(notification objc.Object) {
	cocoa.NSApp().SetActivationPolicy(0)
}

func (delegate *AppDelegate) ApplicationDidFinishLaunching(notification objc.Object) {

	t := cocoa.NSTextView_Init(core.Rect(0, 0, 1400, 300))
	t.SetString("Hello again")
	t.SetSelectable(false)
	t.SetRichText(false)
	t.SetEditable(false)
	t.SetFieldEditor(false)
	t.SetImportsGraphics(false)
	t.SetDrawsBackground(false)
	t.SetFont(cocoa.NSFont_Init("Helvetica", 258.0))
	t.SetAlignment(cocoa.NSTextAlignmentCenter)

	effect = objc.Get("NSVisualEffectView").Alloc().Init()
	effect.Set("translatesAutoresizingMaskIntoConstraints:", false)
	effect.Set("state:", 1)
	effect.Set("blendingMode:", cocoa.NSVisualEffectBlendingModeBehindWindow)
	effect.Set("material:", 2)
	effect.Set("wantsLayer:", true)
	effect.Get("layer").Set("masksToBounds:", true)
	effect.Get("layer").Set("cornerRadius:", 16.0)

	win := cocoa.NSWindow_Init(core.Rect(0, 0, 1400, 300), cocoa.NSTitledWindowMask, cocoa.NSBackingStoreBuffered, false)
	win.Set("movableByWindowBackground:", true)
	win.Set("titlebarAppearsTransparent:", true)
	win.Set("titleVisibility:", cocoa.NSWindowTitleHidden)
	win.Set("opaque:", false)
	win.Set("backgroundColor:", objc.Get("NSColor").Get("clearColor"))
	win.Center()
	win.SetContentView(effect)
	win.ContentView().Send("addSubview:positioned:relativeTo:", t, cocoa.NSWindowAbove, nil)

	win.SetTitle("Hello world!!")
	win.MakeKeyAndOrderFront(win)

	webview = MakeWebView()

	wv := cocoa.NSWindow_Init(core.Rect(0, 0, 1400, 300), cocoa.NSTitledWindowMask|cocoa.NSClosableWindowMask|cocoa.NSMiniaturizableWindowMask|cocoa.NSResizableWindowMask, cocoa.NSBackingStoreBuffered, false)
	wv.Set("movableByWindowBackground:", true)
	wv.Set("opaque:", false)
	wv.Set("backgroundColor:", objc.Get("NSColor").Get("clearColor"))
	wv.Set("ignoresMouseEvents:", true)
	wv.SetContentView(webview)
	wv.MakeKeyAndOrderFront(wv)

	cocoa.NSApp().SetMainMenu(MakeMenu())

	//log.Println(w2.ContentRectForFrameRect(cocoa.Rect(200.0, 300.0, 200.0, 300.0)), w2.IsVisible())
	//debug.FontTest()

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
	core.NSAutoreleasePool_New()

	app := cocoa.NSApp()
	delegate := objc.Get("AppDelegate").Alloc().Init()

	statusBarItem := objc.Get("NSStatusBar").Send("systemStatusBar").Send("statusItemWithLength:", -1.0)
	statusBarItem.Send("button").Send("setTitle:", core.String("Hello world"))
	statusBarItem.Send("setTarget:", delegate)
	statusBarItem.Send("setAction:", objc.Sel("foobar:"))

	app.SetDelegate(delegate)
	fmt.Println("running...")
	app.Run()
}

func MakeWebView() core.WKWebView {
	config := core.WKWebViewConfiguration_New()
	wv := core.WKWebView_Init(core.Rect(0, 0, 1400, 300), config)

	return wv
}

func MakeMenu() cocoa.NSMenu {
	mainMenu := cocoa.NSMenu_New()
	mainMenu.Autorelease()

	mainAppItem := cocoa.NSMenuItem_New()
	mainAppItem.Autorelease()

	mainFileItem := cocoa.NSMenuItem_New()
	mainFileItem.Autorelease()

	mainMenu.AddItem(mainAppItem)
	mainMenu.AddItem(mainFileItem)

	fileMenu := cocoa.NSMenu_Init("File")
	fileMenu.Autorelease()
	mainFileItem.SetSubmenu(fileMenu)

	appMenu := cocoa.NSMenu_Init("App")
	appMenu.Autorelease()
	mainAppItem.SetSubmenu(appMenu)

	quitItem := cocoa.NSMenuItem_New()
	quitItem.Send("setKeyEquivalent:", core.String("q"))
	quitItem.Send("setTitle:", core.String("Quit"))
	quitItem.Send("setAction:", objc.Sel("terminate:"))
	quitItem.Autorelease()

	quitItem2 := cocoa.NSMenuItem_New()
	quitItem2.Send("setTitle:", core.String("Foobar"))
	quitItem2.Send("setAction:", objc.Sel("foobar:"))
	quitItem2.Autorelease()

	fileMenu.AddItem(quitItem2)
	appMenu.AddItem(quitItem)

	return mainMenu
}
