// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"github.com/progrium/macdriver/objc"
)

// The methods for presenting native user interface elements on behalf of a webpage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate?language=objc
type PUIDelegate interface {
	// optional
	WebViewDidClose(webView WebView)
	HasWebViewDidClose() bool
}

// A delegate implementation builder for the [PUIDelegate] protocol.
type UIDelegate struct {
	_WebViewDidClose func(webView WebView)
}

func (di *UIDelegate) HasWebViewDidClose() bool {
	return di._WebViewDidClose != nil
}

// Notifies your app that the DOM window closed successfully. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1537390-webviewdidclose?language=objc
func (di *UIDelegate) SetWebViewDidClose(f func(webView WebView)) {
	di._WebViewDidClose = f
}

// Notifies your app that the DOM window closed successfully. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1537390-webviewdidclose?language=objc
func (di *UIDelegate) WebViewDidClose(webView WebView) {
	di._WebViewDidClose(webView)
}

// A concrete type wrapper for the [PUIDelegate] protocol.
type UIDelegateWrapper struct {
	objc.Object
}

func (u_ UIDelegateWrapper) HasWebViewDidClose() bool {
	return u_.RespondsToSelector(objc.Sel("webViewDidClose:"))
}

// Notifies your app that the DOM window closed successfully. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkuidelegate/1537390-webviewdidclose?language=objc
func (u_ UIDelegateWrapper) WebViewDidClose(webView IWebView) {
	objc.Call[objc.Void](u_, objc.Sel("webViewDidClose:"), objc.Ptr(webView))
}
