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
	w.SetTitle("Test widgets")

	sv := appkit.NewVerticalStackView()
	w.SetContentView(sv)

	webView := webkit.NewWebView()
	webView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	webView.LoadHTMLStringBaseURL(html, foundation.URLClass.URLWithString(url))
	sv.AddViewInGravity(webView, appkit.StackViewGravityTop)

	snapshotButton := appkit.NewButtonWithTitle("capture")

	snapshotWin := appkit.NewWindowWithSize(0, 0)
	objc.Retain(&snapshotWin)
	snapshotWin.SetTitle("Test widgets")

	snapshotWebView := webkit.NewWebView()
	snapshotWebView.SetTranslatesAutoresizingMaskIntoConstraints(false)
	snapshotWin.SetContentView(snapshotWebView)

	navigationDelegate := &webkit.NavigationDelegate{}
	navigationDelegate.SetWebViewDidFinishNavigation(func(webView webkit.WebView, navigation webkit.Navigation) {
		dispatch.MainQueue().DispatchAsync(func() {
			script := `var rect = {"width":document.body.scrollWidth, "height":document.body.scrollHeight}; rect`
			webView.EvaluateJavaScriptCompletionHandler(script, func(value objc.Object, err foundation.Error) {
				rect := foundation.DictToMap[string, foundation.Number](foundation.DictionaryFrom(value.Ptr()))
				width := rect["width"].DoubleValue()
				height := rect["height"].DoubleValue()
				snapshotWin.SetFrameDisplay(foundation.Rect{Size: foundation.Size{Width: width, Height: height}}, true)
				snapshotWebView.LoadHTMLStringBaseURL(html, foundation.URLClass.URLWithString(url))
			})
		})
	})
	webView.SetNavigationDelegate(navigationDelegate)

	ssnd := &webkit.NavigationDelegate{}
	ssnd.SetWebViewDidFinishNavigation(func(webView webkit.WebView, navigation webkit.Navigation) {
		snapshotButton.SetEnabled(true)
	})
	snapshotWebView.SetNavigationDelegate(ssnd)

	action.Set(snapshotButton, func(sender objc.Object) {
		snapshotWebView.TakeSnapshotWithConfigurationCompletionHandler(nil, func(image appkit.Image, err foundation.Error) {
			imageRef := image.CGImageForProposedRectContextHints(nil, nil, nil)
			imageRepo := appkit.NewBitmapImageRepWithCGImage(imageRef)
			imageRepo.SetSize(image.Size())
			pngData := imageRepo.RepresentationUsingTypeProperties(appkit.BitmapImageFileTypePNG, nil)

			if err := os.WriteFile("webview_screenshot.png", pngData, os.ModePerm); err != nil {
				fmt.Println("write image to file error:", err)
			} else {
				fmt.Println("image captured to webview_screenshot.png")
			}
		})
	})
	snapshotButton.SetEnabled(false)

	sv.AddViewInGravity(snapshotButton, appkit.StackViewGravityTop)

	wd := &appkit.WindowDelegate{}
	wd.SetWindowWillClose(func(notification foundation.Notification) {
		snapshotWin.Close()
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
