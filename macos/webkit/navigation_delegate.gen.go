// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"github.com/progrium/macdriver/objc"
)

// Methods for accepting or rejecting navigation changes, and for tracking the progress of navigation requests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationdelegate?language=objc
type PNavigationDelegate interface {
	// optional
	WebViewWebContentProcessDidTerminate(webView WebView)
	HasWebViewWebContentProcessDidTerminate() bool

	// optional
	WebViewDidFinishNavigation(webView WebView, navigation Navigation)
	HasWebViewDidFinishNavigation() bool
}

// A delegate implementation builder for the [PNavigationDelegate] protocol.
type NavigationDelegate struct {
	_WebViewWebContentProcessDidTerminate func(webView WebView)
	_WebViewDidFinishNavigation           func(webView WebView, navigation Navigation)
}

func (di *NavigationDelegate) HasWebViewWebContentProcessDidTerminate() bool {
	return di._WebViewWebContentProcessDidTerminate != nil
}

// Tells the delegate that the web view’s content process was terminated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationdelegate/1455639-webviewwebcontentprocessdidtermi?language=objc
func (di *NavigationDelegate) SetWebViewWebContentProcessDidTerminate(f func(webView WebView)) {
	di._WebViewWebContentProcessDidTerminate = f
}

// Tells the delegate that the web view’s content process was terminated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationdelegate/1455639-webviewwebcontentprocessdidtermi?language=objc
func (di *NavigationDelegate) WebViewWebContentProcessDidTerminate(webView WebView) {
	di._WebViewWebContentProcessDidTerminate(webView)
}
func (di *NavigationDelegate) HasWebViewDidFinishNavigation() bool {
	return di._WebViewDidFinishNavigation != nil
}

// Tells the delegate that navigation is complete. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationdelegate/1455629-webview?language=objc
func (di *NavigationDelegate) SetWebViewDidFinishNavigation(f func(webView WebView, navigation Navigation)) {
	di._WebViewDidFinishNavigation = f
}

// Tells the delegate that navigation is complete. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationdelegate/1455629-webview?language=objc
func (di *NavigationDelegate) WebViewDidFinishNavigation(webView WebView, navigation Navigation) {
	di._WebViewDidFinishNavigation(webView, navigation)
}

// A concrete type wrapper for the [PNavigationDelegate] protocol.
type NavigationDelegateWrapper struct {
	objc.Object
}

func (n_ NavigationDelegateWrapper) HasWebViewWebContentProcessDidTerminate() bool {
	return n_.RespondsToSelector(objc.Sel("webViewWebContentProcessDidTerminate:"))
}

// Tells the delegate that the web view’s content process was terminated. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationdelegate/1455639-webviewwebcontentprocessdidtermi?language=objc
func (n_ NavigationDelegateWrapper) WebViewWebContentProcessDidTerminate(webView IWebView) {
	objc.Call[objc.Void](n_, objc.Sel("webViewWebContentProcessDidTerminate:"), objc.Ptr(webView))
}

func (n_ NavigationDelegateWrapper) HasWebViewDidFinishNavigation() bool {
	return n_.RespondsToSelector(objc.Sel("webView:didFinishNavigation:"))
}

// Tells the delegate that navigation is complete. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wknavigationdelegate/1455629-webview?language=objc
func (n_ NavigationDelegateWrapper) WebViewDidFinishNavigation(webView IWebView, navigation INavigation) {
	objc.Call[objc.Void](n_, objc.Sel("webView:didFinishNavigation:"), objc.Ptr(webView), objc.Ptr(navigation))
}
