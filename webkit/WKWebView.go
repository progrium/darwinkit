package webkit

import (
	"github.com/progrium/macdriver/cocoa"
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type WKWebView struct {
	cocoa.NSView
}

var wkWebView = objc.Get("WKWebView")

func WKWebView_Init(frame core.NSRect, config WKWebViewConfiguration) WKWebView {
	return WKWebView{cocoa.NSView{wkWebView.Alloc().Send("initWithFrame:configuration:", frame, config)}}
}

func (wv WKWebView) LoadRequest(req core.NSURLRequest) {
	wv.Send("loadRequest:", req)
}

func (wv WKWebView) Reload(sender objc.Object) {
	wv.Send("reload:", sender)
}
