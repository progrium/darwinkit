package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"runtime"

	"github.com/progrium/macdriver/pkg/cocoa"
	"github.com/progrium/macdriver/pkg/core"
	"github.com/progrium/macdriver/pkg/objc"
)

type AppDelegate struct {
	objc.Object `objc:"AppDelegate : NSObject"`
}

func (delegate *AppDelegate) ApplicationDidFinishLaunching(notification objc.Object) {
	config := core.WKWebViewConfiguration_New()
	webview := core.WKWebView_Init(cocoa.NSScreen_Main().Frame(), config)
	wv := cocoa.NSWindow_Init(cocoa.NSScreen_Main().Frame(), cocoa.NSBorderlessWindowMask, cocoa.NSBackingStoreBuffered, false)

	wv.SetTitlebarAppearsTransparent(true)
	wv.SetTitleVisibility(cocoa.NSWindowTitleHidden)
	wv.Set("movableByWindowBackground:", true)
	wv.Set("opaque:", false)
	wv.Set("backgroundColor:", objc.Get("NSColor").Get("clearColor"))
	wv.Set("ignoresMouseEvents:", true)
	wv.Set("level:", cocoa.NSMainMenuWindowLevel+2)
	wv.SetContentView(webview)
	wv.MakeKeyAndOrderFront(wv)

	url := core.NSURL_Init("http://localhost:6545")
	req := core.NSURLRequest_Init(url)
	webview.LoadRequest(req)
	webview.Set("opaque:", false)
	webview.Set("backgroundColor:", objc.Get("NSColor").Get("clearColor"))
	webview.Send("setValue:forKey:", objc.Get("NSNumber").Send("numberWithBool:", false), core.String("drawsBackground"))

}

func main() {
	runtime.LockOSThread()

	c := objc.NewClass(AppDelegate{})
	c.AddMethod("applicationDidFinishLaunching:", (*AppDelegate).ApplicationDidFinishLaunching)
	objc.RegisterClass(c)

	go server()

	delegate := objc.Get("AppDelegate").Alloc().Init()
	app := cocoa.NSApp()
	app.SetDelegate(delegate)
	app.SetActivationPolicy(cocoa.NSApplicationActivationPolicyAccessory)
	app.ActivateIgnoringOtherApps(true)
	fmt.Println("running...")
	app.Run()

}

func server() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, `<html style="background-color: transparent;">
	  <body style="margin: 0px; background-color: transparent; border: 2px solid blue;">
		<div style="padding: 30px; background: red; position: absolute; left: 100px; top: 100px;">Hello</div>
	  </body>
	</html>`)
	})

	log.Fatal(http.ListenAndServe(":6545", nil))
}
