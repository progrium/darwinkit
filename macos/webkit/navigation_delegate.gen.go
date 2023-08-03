// auto-generated code, do not modify
package webkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type INavigationDelegate interface {
	ImplementsWebViewDecidePolicyForNavigationActionPreferencesDecisionHandler() bool
	// optional
	WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler(webView WebView, navigationAction NavigationAction, preferences WebpagePreferences, decisionHandler func(param1 NavigationActionPolicy, param2 WebpagePreferences))
	ImplementsWebViewDecidePolicyForNavigationActionDecisionHandler() bool
	// optional
	WebViewDecidePolicyForNavigationActionDecisionHandler(webView WebView, navigationAction NavigationAction, decisionHandler func(param1 NavigationActionPolicy))
	ImplementsWebViewDecidePolicyForNavigationResponseDecisionHandler() bool
	// optional
	WebViewDecidePolicyForNavigationResponseDecisionHandler(webView WebView, navigationResponse NavigationResponse, decisionHandler func(param1 NavigationResponsePolicy))
	ImplementsWebViewDidStartProvisionalNavigation() bool
	// optional
	WebViewDidStartProvisionalNavigation(webView WebView, navigation Navigation)
	ImplementsWebViewDidReceiveServerRedirectForProvisionalNavigation() bool
	// optional
	WebViewDidReceiveServerRedirectForProvisionalNavigation(webView WebView, navigation Navigation)
	ImplementsWebViewDidCommitNavigation() bool
	// optional
	WebViewDidCommitNavigation(webView WebView, navigation Navigation)
	ImplementsWebViewDidFinishNavigation() bool
	// optional
	WebViewDidFinishNavigation(webView WebView, navigation Navigation)
	ImplementsWebViewDidFailNavigationWithError() bool
	// optional
	WebViewDidFailNavigationWithError(webView WebView, navigation Navigation, error foundation.Error)
	ImplementsWebViewDidFailProvisionalNavigationWithError() bool
	// optional
	WebViewDidFailProvisionalNavigationWithError(webView WebView, navigation Navigation, error foundation.Error)
	ImplementsWebViewWebContentProcessDidTerminate() bool
	// optional
	WebViewWebContentProcessDidTerminate(webView WebView)
	ImplementsWebViewNavigationResponseDidBecomeDownload() bool
	// optional
	WebViewNavigationResponseDidBecomeDownload(webView WebView, navigationResponse NavigationResponse, download Download)
	ImplementsWebViewNavigationActionDidBecomeDownload() bool
	// optional
	WebViewNavigationActionDidBecomeDownload(webView WebView, navigationAction NavigationAction, download Download)
}

type NavigationDelegate struct {
	_WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler func(webView WebView, navigationAction NavigationAction, preferences WebpagePreferences, decisionHandler func(param1 NavigationActionPolicy, param2 WebpagePreferences))
	_WebViewDecidePolicyForNavigationActionDecisionHandler            func(webView WebView, navigationAction NavigationAction, decisionHandler func(param1 NavigationActionPolicy))
	_WebViewDecidePolicyForNavigationResponseDecisionHandler          func(webView WebView, navigationResponse NavigationResponse, decisionHandler func(param1 NavigationResponsePolicy))
	_WebViewDidStartProvisionalNavigation                             func(webView WebView, navigation Navigation)
	_WebViewDidReceiveServerRedirectForProvisionalNavigation          func(webView WebView, navigation Navigation)
	_WebViewDidCommitNavigation                                       func(webView WebView, navigation Navigation)
	_WebViewDidFinishNavigation                                       func(webView WebView, navigation Navigation)
	_WebViewDidFailNavigationWithError                                func(webView WebView, navigation Navigation, error foundation.Error)
	_WebViewDidFailProvisionalNavigationWithError                     func(webView WebView, navigation Navigation, error foundation.Error)
	_WebViewWebContentProcessDidTerminate                             func(webView WebView)
	_WebViewNavigationResponseDidBecomeDownload                       func(webView WebView, navigationResponse NavigationResponse, download Download)
	_WebViewNavigationActionDidBecomeDownload                         func(webView WebView, navigationAction NavigationAction, download Download)
}

func (di *NavigationDelegate) ImplementsWebViewDecidePolicyForNavigationActionPreferencesDecisionHandler() bool {
	return di._WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler != nil
}

func (di *NavigationDelegate) SetWebViewDecidePolicyForNavigationActionPreferencesDecisionHandler(f func(webView WebView, navigationAction NavigationAction, preferences WebpagePreferences, decisionHandler func(param1 NavigationActionPolicy, param2 WebpagePreferences))) {
	di._WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler = f
}

func (di *NavigationDelegate) WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler(webView WebView, navigationAction NavigationAction, preferences WebpagePreferences, decisionHandler func(param1 NavigationActionPolicy, param2 WebpagePreferences)) {
	di._WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler(webView, navigationAction, preferences, decisionHandler)
}
func (di *NavigationDelegate) ImplementsWebViewDecidePolicyForNavigationActionDecisionHandler() bool {
	return di._WebViewDecidePolicyForNavigationActionDecisionHandler != nil
}

func (di *NavigationDelegate) SetWebViewDecidePolicyForNavigationActionDecisionHandler(f func(webView WebView, navigationAction NavigationAction, decisionHandler func(param1 NavigationActionPolicy))) {
	di._WebViewDecidePolicyForNavigationActionDecisionHandler = f
}

func (di *NavigationDelegate) WebViewDecidePolicyForNavigationActionDecisionHandler(webView WebView, navigationAction NavigationAction, decisionHandler func(param1 NavigationActionPolicy)) {
	di._WebViewDecidePolicyForNavigationActionDecisionHandler(webView, navigationAction, decisionHandler)
}
func (di *NavigationDelegate) ImplementsWebViewDecidePolicyForNavigationResponseDecisionHandler() bool {
	return di._WebViewDecidePolicyForNavigationResponseDecisionHandler != nil
}

func (di *NavigationDelegate) SetWebViewDecidePolicyForNavigationResponseDecisionHandler(f func(webView WebView, navigationResponse NavigationResponse, decisionHandler func(param1 NavigationResponsePolicy))) {
	di._WebViewDecidePolicyForNavigationResponseDecisionHandler = f
}

func (di *NavigationDelegate) WebViewDecidePolicyForNavigationResponseDecisionHandler(webView WebView, navigationResponse NavigationResponse, decisionHandler func(param1 NavigationResponsePolicy)) {
	di._WebViewDecidePolicyForNavigationResponseDecisionHandler(webView, navigationResponse, decisionHandler)
}
func (di *NavigationDelegate) ImplementsWebViewDidStartProvisionalNavigation() bool {
	return di._WebViewDidStartProvisionalNavigation != nil
}

func (di *NavigationDelegate) SetWebViewDidStartProvisionalNavigation(f func(webView WebView, navigation Navigation)) {
	di._WebViewDidStartProvisionalNavigation = f
}

func (di *NavigationDelegate) WebViewDidStartProvisionalNavigation(webView WebView, navigation Navigation) {
	di._WebViewDidStartProvisionalNavigation(webView, navigation)
}
func (di *NavigationDelegate) ImplementsWebViewDidReceiveServerRedirectForProvisionalNavigation() bool {
	return di._WebViewDidReceiveServerRedirectForProvisionalNavigation != nil
}

func (di *NavigationDelegate) SetWebViewDidReceiveServerRedirectForProvisionalNavigation(f func(webView WebView, navigation Navigation)) {
	di._WebViewDidReceiveServerRedirectForProvisionalNavigation = f
}

func (di *NavigationDelegate) WebViewDidReceiveServerRedirectForProvisionalNavigation(webView WebView, navigation Navigation) {
	di._WebViewDidReceiveServerRedirectForProvisionalNavigation(webView, navigation)
}
func (di *NavigationDelegate) ImplementsWebViewDidCommitNavigation() bool {
	return di._WebViewDidCommitNavigation != nil
}

func (di *NavigationDelegate) SetWebViewDidCommitNavigation(f func(webView WebView, navigation Navigation)) {
	di._WebViewDidCommitNavigation = f
}

func (di *NavigationDelegate) WebViewDidCommitNavigation(webView WebView, navigation Navigation) {
	di._WebViewDidCommitNavigation(webView, navigation)
}
func (di *NavigationDelegate) ImplementsWebViewDidFinishNavigation() bool {
	return di._WebViewDidFinishNavigation != nil
}

func (di *NavigationDelegate) SetWebViewDidFinishNavigation(f func(webView WebView, navigation Navigation)) {
	di._WebViewDidFinishNavigation = f
}

func (di *NavigationDelegate) WebViewDidFinishNavigation(webView WebView, navigation Navigation) {
	di._WebViewDidFinishNavigation(webView, navigation)
}
func (di *NavigationDelegate) ImplementsWebViewDidFailNavigationWithError() bool {
	return di._WebViewDidFailNavigationWithError != nil
}

func (di *NavigationDelegate) SetWebViewDidFailNavigationWithError(f func(webView WebView, navigation Navigation, error foundation.Error)) {
	di._WebViewDidFailNavigationWithError = f
}

func (di *NavigationDelegate) WebViewDidFailNavigationWithError(webView WebView, navigation Navigation, error foundation.Error) {
	di._WebViewDidFailNavigationWithError(webView, navigation, error)
}
func (di *NavigationDelegate) ImplementsWebViewDidFailProvisionalNavigationWithError() bool {
	return di._WebViewDidFailProvisionalNavigationWithError != nil
}

func (di *NavigationDelegate) SetWebViewDidFailProvisionalNavigationWithError(f func(webView WebView, navigation Navigation, error foundation.Error)) {
	di._WebViewDidFailProvisionalNavigationWithError = f
}

func (di *NavigationDelegate) WebViewDidFailProvisionalNavigationWithError(webView WebView, navigation Navigation, error foundation.Error) {
	di._WebViewDidFailProvisionalNavigationWithError(webView, navigation, error)
}
func (di *NavigationDelegate) ImplementsWebViewWebContentProcessDidTerminate() bool {
	return di._WebViewWebContentProcessDidTerminate != nil
}

func (di *NavigationDelegate) SetWebViewWebContentProcessDidTerminate(f func(webView WebView)) {
	di._WebViewWebContentProcessDidTerminate = f
}

func (di *NavigationDelegate) WebViewWebContentProcessDidTerminate(webView WebView) {
	di._WebViewWebContentProcessDidTerminate(webView)
}
func (di *NavigationDelegate) ImplementsWebViewNavigationResponseDidBecomeDownload() bool {
	return di._WebViewNavigationResponseDidBecomeDownload != nil
}

func (di *NavigationDelegate) SetWebViewNavigationResponseDidBecomeDownload(f func(webView WebView, navigationResponse NavigationResponse, download Download)) {
	di._WebViewNavigationResponseDidBecomeDownload = f
}

func (di *NavigationDelegate) WebViewNavigationResponseDidBecomeDownload(webView WebView, navigationResponse NavigationResponse, download Download) {
	di._WebViewNavigationResponseDidBecomeDownload(webView, navigationResponse, download)
}
func (di *NavigationDelegate) ImplementsWebViewNavigationActionDidBecomeDownload() bool {
	return di._WebViewNavigationActionDidBecomeDownload != nil
}

func (di *NavigationDelegate) SetWebViewNavigationActionDidBecomeDownload(f func(webView WebView, navigationAction NavigationAction, download Download)) {
	di._WebViewNavigationActionDidBecomeDownload = f
}

func (di *NavigationDelegate) WebViewNavigationActionDidBecomeDownload(webView WebView, navigationAction NavigationAction, download Download) {
	di._WebViewNavigationActionDidBecomeDownload(webView, navigationAction, download)
}

type NavigationDelegateWrapper struct {
	objc.Object
}

func (n_ NavigationDelegateWrapper) ImplementsWebViewDecidePolicyForNavigationActionPreferencesDecisionHandler() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:decidePolicyForNavigationAction:preferences:decisionHandler:"))
}

func (n_ NavigationDelegateWrapper) WebViewDecidePolicyForNavigationActionPreferencesDecisionHandler(webView IWebView, navigationAction INavigationAction, preferences IWebpagePreferences, decisionHandler func(param1 NavigationActionPolicy, param2 WebpagePreferences)) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:decidePolicyForNavigationAction:preferences:decisionHandler:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigationAction), objc.ExtractPtr(preferences), decisionHandler)
}

func (n_ NavigationDelegateWrapper) ImplementsWebViewDecidePolicyForNavigationActionDecisionHandler() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:decidePolicyForNavigationAction:decisionHandler:"))
}

func (n_ NavigationDelegateWrapper) WebViewDecidePolicyForNavigationActionDecisionHandler(webView IWebView, navigationAction INavigationAction, decisionHandler func(param1 NavigationActionPolicy)) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:decidePolicyForNavigationAction:decisionHandler:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigationAction), decisionHandler)
}

func (n_ NavigationDelegateWrapper) ImplementsWebViewDecidePolicyForNavigationResponseDecisionHandler() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:decidePolicyForNavigationResponse:decisionHandler:"))
}

func (n_ NavigationDelegateWrapper) WebViewDecidePolicyForNavigationResponseDecisionHandler(webView IWebView, navigationResponse INavigationResponse, decisionHandler func(param1 NavigationResponsePolicy)) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:decidePolicyForNavigationResponse:decisionHandler:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigationResponse), decisionHandler)
}

func (n_ NavigationDelegateWrapper) ImplementsWebViewDidStartProvisionalNavigation() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:didStartProvisionalNavigation:"))
}

func (n_ NavigationDelegateWrapper) WebViewDidStartProvisionalNavigation(webView IWebView, navigation INavigation) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:didStartProvisionalNavigation:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigation))
}

func (n_ NavigationDelegateWrapper) ImplementsWebViewDidReceiveServerRedirectForProvisionalNavigation() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:didReceiveServerRedirectForProvisionalNavigation:"))
}

func (n_ NavigationDelegateWrapper) WebViewDidReceiveServerRedirectForProvisionalNavigation(webView IWebView, navigation INavigation) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:didReceiveServerRedirectForProvisionalNavigation:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigation))
}

func (n_ NavigationDelegateWrapper) ImplementsWebViewDidCommitNavigation() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:didCommitNavigation:"))
}

func (n_ NavigationDelegateWrapper) WebViewDidCommitNavigation(webView IWebView, navigation INavigation) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:didCommitNavigation:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigation))
}

func (n_ NavigationDelegateWrapper) ImplementsWebViewDidFinishNavigation() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:didFinishNavigation:"))
}

func (n_ NavigationDelegateWrapper) WebViewDidFinishNavigation(webView IWebView, navigation INavigation) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:didFinishNavigation:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigation))
}

func (n_ NavigationDelegateWrapper) ImplementsWebViewDidFailNavigationWithError() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:didFailNavigation:withError:"))
}

func (n_ NavigationDelegateWrapper) WebViewDidFailNavigationWithError(webView IWebView, navigation INavigation, error foundation.IError) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:didFailNavigation:withError:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigation), objc.ExtractPtr(error))
}

func (n_ NavigationDelegateWrapper) ImplementsWebViewDidFailProvisionalNavigationWithError() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:didFailProvisionalNavigation:withError:"))
}

func (n_ NavigationDelegateWrapper) WebViewDidFailProvisionalNavigationWithError(webView IWebView, navigation INavigation, error foundation.IError) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:didFailProvisionalNavigation:withError:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigation), objc.ExtractPtr(error))
}

func (n_ NavigationDelegateWrapper) ImplementsWebViewWebContentProcessDidTerminate() bool {
	return n_.RespondsToSelector(objc.GetSelector("webViewWebContentProcessDidTerminate:"))
}

func (n_ NavigationDelegateWrapper) WebViewWebContentProcessDidTerminate(webView IWebView) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webViewWebContentProcessDidTerminate:"), objc.ExtractPtr(webView))
}

func (n_ NavigationDelegateWrapper) ImplementsWebViewNavigationResponseDidBecomeDownload() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:navigationResponse:didBecomeDownload:"))
}

func (n_ NavigationDelegateWrapper) WebViewNavigationResponseDidBecomeDownload(webView IWebView, navigationResponse INavigationResponse, download IDownload) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:navigationResponse:didBecomeDownload:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigationResponse), objc.ExtractPtr(download))
}

func (n_ NavigationDelegateWrapper) ImplementsWebViewNavigationActionDidBecomeDownload() bool {
	return n_.RespondsToSelector(objc.GetSelector("webView:navigationAction:didBecomeDownload:"))
}

func (n_ NavigationDelegateWrapper) WebViewNavigationActionDidBecomeDownload(webView IWebView, navigationAction INavigationAction, download IDownload) {
	objc.CallMethod[objc.Void](n_, objc.GetSelector("webView:navigationAction:didBecomeDownload:"), objc.ExtractPtr(webView), objc.ExtractPtr(navigationAction), objc.ExtractPtr(download))
}
