// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"github.com/progrium/macdriver/objc"
)

// A protocol for loading resources with URL schemes that WebKit doesn't handle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkurlschemehandler?language=objc
type PURLSchemeHandler interface {
	// optional
	WebViewStartURLSchemeTask(webView WebView, urlSchemeTask URLSchemeTaskWrapper)
	HasWebViewStartURLSchemeTask() bool
}

// A concrete type wrapper for the [PURLSchemeHandler] protocol.
type URLSchemeHandlerWrapper struct {
	objc.Object
}

func (u_ URLSchemeHandlerWrapper) HasWebViewStartURLSchemeTask() bool {
	return u_.RespondsToSelector(objc.Sel("webView:startURLSchemeTask:"))
}

// Asks your handler to begin loading the data for the specified resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkurlschemehandler/2890834-webview?language=objc
func (u_ URLSchemeHandlerWrapper) WebViewStartURLSchemeTask(webView IWebView, urlSchemeTask PURLSchemeTask) {
	po1 := objc.WrapAsProtocol("WKURLSchemeTask", urlSchemeTask)
	objc.Call[objc.Void](u_, objc.Sel("webView:startURLSchemeTask:"), objc.Ptr(webView), po1)
}
