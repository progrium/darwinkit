package main

import (
	"fmt"
	"os"

	"github.com/progrium/macdriver/dispatch"
	"github.com/progrium/macdriver/helper/action"
	"github.com/progrium/macdriver/macos"
	"github.com/progrium/macdriver/macos/appkit"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/webkit"
	"github.com/progrium/macdriver/objc"
)

func main() {
	macos.RunApp(launched)
}

func launched(app appkit.Application, delegate *appkit.ApplicationDelegate) {
	var html = "<html><h1>Hello World!</h1></html>"
	var url = "https://www.test.com"

	w := appkit.NewWindowWithSize(600, 400)
	objc.Retain(&w)
	w.SetTitle("Webshot Demo")
	sv := appkit.NewVerticalStackView()
	w.SetContentView(sv)

	wv := webkit.NewWebView()
	wv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	wv.LoadHTMLStringBaseURL(html, foundation.URL_URLWithString(url))
	sv.AddViewInGravity(wv, appkit.StackViewGravityTop)

	sw := appkit.NewWindowWithSize(0, 0)
	objc.Retain(&sw)
	swv := webkit.NewWebView()
	swv.SetTranslatesAutoresizingMaskIntoConstraints(false)
	sw.SetContentView(swv)

	navDelegate := &webkit.NavigationDelegate{}
	navDelegate.SetWebViewDidFinishNavigation(func(webView webkit.WebView, navigation webkit.Navigation) {
		dispatch.MainQueue().DispatchAsync(func() {
			script := `var rect = {"width":document.body.scrollWidth, "height":document.body.scrollHeight}; rect`
			webView.EvaluateJavaScriptCompletionHandler(script, func(value objc.Object, err foundation.Error) {
				rect := foundation.DictToMap[string, foundation.Number](foundation.DictionaryFrom(value.Ptr()))
				width := rect["width"].DoubleValue()
				height := rect["height"].DoubleValue()
				sw.SetFrameDisplay(foundation.Rect{Size: foundation.Size{Width: width, Height: height}}, true)
				swv.LoadHTMLStringBaseURL(html, foundation.URL_URLWithString(url))
			})
		})
	})
	wv.SetNavigationDelegate(navDelegate)

	button := appkit.NewButtonWithTitle("capture")

	snd := &webkit.NavigationDelegate{}
	snd.SetWebViewDidFinishNavigation(func(webView webkit.WebView, navigation webkit.Navigation) {
		button.SetEnabled(true)
	})
	swv.SetNavigationDelegate(snd)

	action.Set(button, func(sender objc.Object) {
		swv.TakeSnapshotWithConfigurationCompletionHandler(nil, func(image appkit.Image, err foundation.Error) {
			imgref := image.CGImageForProposedRectContextHints(nil, nil, nil)
			img := appkit.NewBitmapImageRepWithCGImage(imgref)
			img.SetSize(image.Size())
			png := img.RepresentationUsingTypeProperties(appkit.BitmapImageFileTypePNG, nil)
			if err := os.WriteFile("webview_screenshot.png", png, os.ModePerm); err != nil {
				fmt.Println("image write to file error:", err)
			} else {
				fmt.Println("image captured to webview_screenshot.png")
			}
		})
	})
	button.SetEnabled(false)
	sv.AddViewInGravity(button, appkit.StackViewGravityTop)

	wd := &appkit.WindowDelegate{}
	wd.SetWindowWillClose(func(notification foundation.Notification) {
		sw.Close()
	})
	w.SetDelegate(wd)
	w.MakeKeyAndOrderFront(nil)
	w.Center()

	delegate.SetApplicationShouldTerminateAfterLastWindowClosed(func(appkit.Application) bool {
		return true
	})
	app.SetActivationPolicy(appkit.ApplicationActivationPolicyRegular)
	app.ActivateIgnoringOtherApps(true)
}
