package webkit

import (
	"github.com/progrium/macdriver/core"
	"github.com/progrium/macdriver/objc"
)

type WKWebView struct {
	gen_WKWebView
}

func WKWebView_Init(frame core.NSRect, config WKWebViewConfiguration) WKWebView {
	return WKWebView_alloc().InitWithFrame_configuration__asWKWebView(frame, config)
}

func (wv WKWebView) LoadRequest(req core.NSURLRequest) {
	wv.LoadRequest_(req)
}

// FIXME this would conflict with the `reload` selector that doesn't take a
// parameter and returns WKNavigation. Should we standardize on that one
// instead?
func (wv WKWebView) Reload(sender objc.Object) {
	wv.Reload_(sender)
}

func (wv WKWebView) LoadHTMLString(html core.NSString, url core.NSURL) {
	wv.LoadHTMLString_baseURL_(html, url)
}
