// auto-generated code, do not modify
package webkit

import (
	"github.com/progrium/macdriver/objc"
)

type IURLSchemeHandler interface {
	// required
	WebViewStartURLSchemeTask(webView WebView, urlSchemeTask URLSchemeTaskWrapper)
	// required
	WebViewStopURLSchemeTask(webView WebView, urlSchemeTask URLSchemeTaskWrapper)
}

type URLSchemeHandlerWrapper struct {
	objc.Object
}

func (u_ URLSchemeHandlerWrapper) WebViewStartURLSchemeTask(webView IWebView, urlSchemeTask IURLSchemeTask) {
	po := objc.WrapAsProtocol("WKURLSchemeTask", urlSchemeTask)
	objc.CallMethod[objc.Void](u_, objc.GetSelector("webView:startURLSchemeTask:"), objc.ExtractPtr(webView), po)
}

func (u_ URLSchemeHandlerWrapper) WebViewStopURLSchemeTask(webView IWebView, urlSchemeTask IURLSchemeTask) {
	po := objc.WrapAsProtocol("WKURLSchemeTask", urlSchemeTask)
	objc.CallMethod[objc.Void](u_, objc.GetSelector("webView:stopURLSchemeTask:"), objc.ExtractPtr(webView), po)
}
