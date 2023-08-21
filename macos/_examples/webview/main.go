package main

import (
	"embed"
	"fmt"
	"io/fs"
	"runtime"
	"time"

	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/webkit"
	"github.com/progrium/macdriver/objc"
)

//go:embed assets
var assetsFS embed.FS

func init() {
	runtime.LockOSThread()
}

func main() {
	app := appkit.ApplicationClass.SharedApplication()
	w := appkit.NewWindowWithSize(600, 400)
	w.SetTitle("Test FS WebView")

	_fs, _ := fs.Sub(assetsFS, "assets")

	configuration := webkit.NewWebViewConfiguration()
	gofsHandler := &webkit.FileSystemURLSchemeHandler{FS: _fs}
	configuration.SetURLSchemeHandlerForURLScheme(gofsHandler, "gofs")

	view := webkit.NewWebViewWithFrameConfiguration(foundation.Rect{}, configuration)
	webkit.AddScriptMessageHandlerWithReply(view, "greet", func(message objc.Object) (objc.Object, error) {
		param := message.Description()
		fmt.Println("greet handled")
		return foundation.NewStringWithString("hello: " + param).Object, nil
	})
	w.SetContentView(view)

	w.MakeKeyAndOrderFront(nil)
	w.Center()

	ad := &appkit.ApplicationDelegate{}
	ad.SetApplicationDidFinishLaunching(func(foundation.Notification) {
		app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
		app.ActivateIgnoringOtherApps(true)
		webkit.LoadURL(view, "gofs:/index.html")
	})
	ad.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
	app.SetDelegate(ad)

	go func() {
		time.Sleep(time.Second * 1)
		runtime.GC()
	}()
	app.Run()
}
