package main

import (
	"embed"
	"fmt"
	"io/fs"

	"github.com/progrium/macdriver/macos"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/webkit"
	"github.com/progrium/macdriver/objc"
)

//go:embed assets
var assetsFS embed.FS

func main() {
	macos.RunApp(launched)
}

func launched(app appkit.Application, delegate *appkit.ApplicationDelegate) {
	w := appkit.NewWindowWithSize(600, 400)
	objc.Retain(&w)

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

	webkit.LoadURL(view, "gofs:/index.html")

	delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)

}
