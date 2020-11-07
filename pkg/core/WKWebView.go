package core

import (
	"github.com/progrium/macdriver/pkg/objc"
)

type WKWebView struct {
	objc.Object
}

var WKWebView_ = objc.Get("WKWebView")

func WKWebView_Init(frame NSRect, config WKWebViewConfiguration) WKWebView {
	return WKWebView{WKWebView_.Alloc().Send("initWithFrame:configuration:", frame, config)}
}

func (wv WKWebView) LoadRequest(req NSURLRequest) {
	wv.Send("loadRequest:", req)
}
