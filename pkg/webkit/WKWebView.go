package webkit

import (
	"github.com/progrium/macdriver/pkg/cocoa"
	"github.com/progrium/macdriver/pkg/core"
	"github.com/progrium/macdriver/pkg/objc"
)

type WKWebView struct {
	cocoa.NSView
}

var WKWebView_ = objc.Get("WKWebView")

func WKWebView_Init(frame core.NSRect, config WKWebViewConfiguration) WKWebView {
	return WKWebView{cocoa.NSView{WKWebView_.Alloc().Send("initWithFrame:configuration:", frame, config)}}
}

func (wv WKWebView) LoadRequest(req core.NSURLRequest) {
	wv.Send("loadRequest:", req)
}

func (wv WKWebView) Reload(sender objc.Object) {
	wv.Send("reload:", sender)
}
