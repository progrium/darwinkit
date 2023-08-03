// auto-generated code, do not modify
package webkit

import (
	"github.com/progrium/macdriver/objc"
)

type IUIDelegate interface {
	ImplementsWebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures() bool
	// optional
	WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures(webView WebView, configuration WebViewConfiguration, navigationAction NavigationAction, windowFeatures WindowFeatures) IWebView
	ImplementsWebViewDidClose() bool
	// optional
	WebViewDidClose(webView WebView)
	ImplementsWebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler() bool
	// optional
	WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler(webView WebView, message string, frame FrameInfo, completionHandler func())
	ImplementsWebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler() bool
	// optional
	WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler(webView WebView, message string, frame FrameInfo, completionHandler func(result bool))
	ImplementsWebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler() bool
	// optional
	WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler(webView WebView, prompt string, defaultText string, frame FrameInfo, completionHandler func(result string))
	ImplementsWebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler() bool
	// optional
	WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler(webView WebView, origin SecurityOrigin, frame FrameInfo, type_ MediaCaptureType, decisionHandler func(decision PermissionDecision))
}

type UIDelegate struct {
	_WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures                   func(webView WebView, configuration WebViewConfiguration, navigationAction NavigationAction, windowFeatures WindowFeatures) IWebView
	_WebViewDidClose                                                                          func(webView WebView)
	_WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler               func(webView WebView, message string, frame FrameInfo, completionHandler func())
	_WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler             func(webView WebView, message string, frame FrameInfo, completionHandler func(result bool))
	_WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler func(webView WebView, prompt string, defaultText string, frame FrameInfo, completionHandler func(result string))
	_WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler         func(webView WebView, origin SecurityOrigin, frame FrameInfo, type_ MediaCaptureType, decisionHandler func(decision PermissionDecision))
}

func (di *UIDelegate) ImplementsWebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures() bool {
	return di._WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures != nil
}

func (di *UIDelegate) SetWebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures(f func(webView WebView, configuration WebViewConfiguration, navigationAction NavigationAction, windowFeatures WindowFeatures) IWebView) {
	di._WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures = f
}

func (di *UIDelegate) WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures(webView WebView, configuration WebViewConfiguration, navigationAction NavigationAction, windowFeatures WindowFeatures) IWebView {
	return di._WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures(webView, configuration, navigationAction, windowFeatures)
}
func (di *UIDelegate) ImplementsWebViewDidClose() bool {
	return di._WebViewDidClose != nil
}

func (di *UIDelegate) SetWebViewDidClose(f func(webView WebView)) {
	di._WebViewDidClose = f
}

func (di *UIDelegate) WebViewDidClose(webView WebView) {
	di._WebViewDidClose(webView)
}
func (di *UIDelegate) ImplementsWebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler() bool {
	return di._WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler != nil
}

func (di *UIDelegate) SetWebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler(f func(webView WebView, message string, frame FrameInfo, completionHandler func())) {
	di._WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler = f
}

func (di *UIDelegate) WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler(webView WebView, message string, frame FrameInfo, completionHandler func()) {
	di._WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler(webView, message, frame, completionHandler)
}
func (di *UIDelegate) ImplementsWebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler() bool {
	return di._WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler != nil
}

func (di *UIDelegate) SetWebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler(f func(webView WebView, message string, frame FrameInfo, completionHandler func(result bool))) {
	di._WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler = f
}

func (di *UIDelegate) WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler(webView WebView, message string, frame FrameInfo, completionHandler func(result bool)) {
	di._WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler(webView, message, frame, completionHandler)
}
func (di *UIDelegate) ImplementsWebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler() bool {
	return di._WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler != nil
}

func (di *UIDelegate) SetWebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler(f func(webView WebView, prompt string, defaultText string, frame FrameInfo, completionHandler func(result string))) {
	di._WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler = f
}

func (di *UIDelegate) WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler(webView WebView, prompt string, defaultText string, frame FrameInfo, completionHandler func(result string)) {
	di._WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler(webView, prompt, defaultText, frame, completionHandler)
}
func (di *UIDelegate) ImplementsWebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler() bool {
	return di._WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler != nil
}

func (di *UIDelegate) SetWebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler(f func(webView WebView, origin SecurityOrigin, frame FrameInfo, type_ MediaCaptureType, decisionHandler func(decision PermissionDecision))) {
	di._WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler = f
}

func (di *UIDelegate) WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler(webView WebView, origin SecurityOrigin, frame FrameInfo, type_ MediaCaptureType, decisionHandler func(decision PermissionDecision)) {
	di._WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler(webView, origin, frame, type_, decisionHandler)
}

type UIDelegateWrapper struct {
	objc.Object
}

func (u_ UIDelegateWrapper) ImplementsWebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures() bool {
	return u_.RespondsToSelector(objc.GetSelector("webView:createWebViewWithConfiguration:forNavigationAction:windowFeatures:"))
}

func (u_ UIDelegateWrapper) WebViewCreateWebViewWithConfigurationForNavigationActionWindowFeatures(webView IWebView, configuration IWebViewConfiguration, navigationAction INavigationAction, windowFeatures IWindowFeatures) WebView {
	rv := objc.CallMethod[WebView](u_, objc.GetSelector("webView:createWebViewWithConfiguration:forNavigationAction:windowFeatures:"), objc.ExtractPtr(webView), objc.ExtractPtr(configuration), objc.ExtractPtr(navigationAction), objc.ExtractPtr(windowFeatures))
	return rv
}

func (u_ UIDelegateWrapper) ImplementsWebViewDidClose() bool {
	return u_.RespondsToSelector(objc.GetSelector("webViewDidClose:"))
}

func (u_ UIDelegateWrapper) WebViewDidClose(webView IWebView) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("webViewDidClose:"), objc.ExtractPtr(webView))
}

func (u_ UIDelegateWrapper) ImplementsWebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler() bool {
	return u_.RespondsToSelector(objc.GetSelector("webView:runJavaScriptAlertPanelWithMessage:initiatedByFrame:completionHandler:"))
}

func (u_ UIDelegateWrapper) WebViewRunJavaScriptAlertPanelWithMessageInitiatedByFrameCompletionHandler(webView IWebView, message string, frame IFrameInfo, completionHandler func()) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("webView:runJavaScriptAlertPanelWithMessage:initiatedByFrame:completionHandler:"), objc.ExtractPtr(webView), message, objc.ExtractPtr(frame), completionHandler)
}

func (u_ UIDelegateWrapper) ImplementsWebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler() bool {
	return u_.RespondsToSelector(objc.GetSelector("webView:runJavaScriptConfirmPanelWithMessage:initiatedByFrame:completionHandler:"))
}

func (u_ UIDelegateWrapper) WebViewRunJavaScriptConfirmPanelWithMessageInitiatedByFrameCompletionHandler(webView IWebView, message string, frame IFrameInfo, completionHandler func(result bool)) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("webView:runJavaScriptConfirmPanelWithMessage:initiatedByFrame:completionHandler:"), objc.ExtractPtr(webView), message, objc.ExtractPtr(frame), completionHandler)
}

func (u_ UIDelegateWrapper) ImplementsWebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler() bool {
	return u_.RespondsToSelector(objc.GetSelector("webView:runJavaScriptTextInputPanelWithPrompt:defaultText:initiatedByFrame:completionHandler:"))
}

func (u_ UIDelegateWrapper) WebViewRunJavaScriptTextInputPanelWithPromptDefaultTextInitiatedByFrameCompletionHandler(webView IWebView, prompt string, defaultText string, frame IFrameInfo, completionHandler func(result string)) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("webView:runJavaScriptTextInputPanelWithPrompt:defaultText:initiatedByFrame:completionHandler:"), objc.ExtractPtr(webView), prompt, defaultText, objc.ExtractPtr(frame), completionHandler)
}

func (u_ UIDelegateWrapper) ImplementsWebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler() bool {
	return u_.RespondsToSelector(objc.GetSelector("webView:requestMediaCapturePermissionForOrigin:initiatedByFrame:type:decisionHandler:"))
}

func (u_ UIDelegateWrapper) WebViewRequestMediaCapturePermissionForOriginInitiatedByFrameTypeDecisionHandler(webView IWebView, origin ISecurityOrigin, frame IFrameInfo, type_ MediaCaptureType, decisionHandler func(decision PermissionDecision)) {
	objc.CallMethod[objc.Void](u_, objc.GetSelector("webView:requestMediaCapturePermissionForOrigin:initiatedByFrame:type:decisionHandler:"), objc.ExtractPtr(webView), objc.ExtractPtr(origin), objc.ExtractPtr(frame), type_, decisionHandler)
}
