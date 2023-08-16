// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ExtensionContext] class.
var ExtensionContextClass = _ExtensionContextClass{objc.GetClass("NSExtensionContext")}

type _ExtensionContextClass struct {
	objc.Class
}

// An interface definition for the [ExtensionContext] class.
type IExtensionContext interface {
	objc.IObject
	DismissNotificationContentExtension()
	CompleteRequestReturningItemsCompletionHandler(items []objc.IObject, completionHandler func(expired bool))
	PerformNotificationDefaultAction()
	OpenURLCompletionHandler(URL IURL, completionHandler func(success bool))
	MediaPlayingStarted()
	CompleteRequestWithBroadcastURLSetupInfo(broadcastURL IURL, setupInfo map[string]objc.IObject)
	MediaPlayingPaused()
	CancelRequestWithError(error IError)
	LoadBroadcastingApplicationInfoWithCompletion(handler func(bundleID string, displayName string, appIcon objc.Object))
	InputItems() []objc.Object
}

// The host app context from which an app extension is invoked. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext?language=objc
type ExtensionContext struct {
	objc.Object
}

func ExtensionContextFrom(ptr unsafe.Pointer) ExtensionContext {
	return ExtensionContext{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ec _ExtensionContextClass) Alloc() ExtensionContext {
	rv := objc.Call[ExtensionContext](ec, objc.Sel("alloc"))
	return rv
}

func ExtensionContext_Alloc() ExtensionContext {
	return ExtensionContextClass.Alloc()
}

func (ec _ExtensionContextClass) New() ExtensionContext {
	rv := objc.Call[ExtensionContext](ec, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewExtensionContext() ExtensionContext {
	return ExtensionContextClass.New()
}

func (e_ ExtensionContext) Init() ExtensionContext {
	rv := objc.Call[ExtensionContext](e_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/2977624-dismissnotificationcontentextens?language=objc
func (e_ ExtensionContext) DismissNotificationContentExtension() {
	objc.Call[objc.Void](e_, objc.Sel("dismissNotificationContentExtension"))
}

// Tells the host app to complete the app extension request with an array of result items. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/1411301-completerequestreturningitems?language=objc
func (e_ ExtensionContext) CompleteRequestReturningItemsCompletionHandler(items []objc.IObject, completionHandler func(expired bool)) {
	objc.Call[objc.Void](e_, objc.Sel("completeRequestReturningItems:completionHandler:"), items, completionHandler)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/2968489-performnotificationdefaultaction?language=objc
func (e_ ExtensionContext) PerformNotificationDefaultAction() {
	objc.Call[objc.Void](e_, objc.Sel("performNotificationDefaultAction"))
}

// Asks the system to open a URL on behalf of the currently running app extension. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/1416791-openurl?language=objc
func (e_ ExtensionContext) OpenURLCompletionHandler(URL IURL, completionHandler func(success bool)) {
	objc.Call[objc.Void](e_, objc.Sel("openURL:completionHandler:"), objc.Ptr(URL), completionHandler)
}

// Tells the system that the Notification Content app extension began playing a media file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/1648523-mediaplayingstarted?language=objc
func (e_ ExtensionContext) MediaPlayingStarted() {
	objc.Call[objc.Void](e_, objc.Sel("mediaPlayingStarted"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/2891095-completerequestwithbroadcasturl?language=objc
func (e_ ExtensionContext) CompleteRequestWithBroadcastURLSetupInfo(broadcastURL IURL, setupInfo map[string]objc.IObject) {
	objc.Call[objc.Void](e_, objc.Sel("completeRequestWithBroadcastURL:setupInfo:"), objc.Ptr(broadcastURL), setupInfo)
}

// Tells the system that the Notification Content app extension stopped playing a media file. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/1648527-mediaplayingpaused?language=objc
func (e_ ExtensionContext) MediaPlayingPaused() {
	objc.Call[objc.Void](e_, objc.Sel("mediaPlayingPaused"))
}

// Tells the host app to cancel the app extension request, with a supplied error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/1412773-cancelrequestwitherror?language=objc
func (e_ ExtensionContext) CancelRequestWithError(error IError) {
	objc.Call[objc.Void](e_, objc.Sel("cancelRequestWithError:"), objc.Ptr(error))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/1845240-loadbroadcastingapplicationinfow?language=objc
func (e_ ExtensionContext) LoadBroadcastingApplicationInfoWithCompletion(handler func(bundleID string, displayName string, appIcon objc.Object)) {
	objc.Call[objc.Void](e_, objc.Sel("loadBroadcastingApplicationInfoWithCompletion:"), handler)
}

// The list of input NSExtensionItem objects associated with the context. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsextensioncontext/1414827-inputitems?language=objc
func (e_ ExtensionContext) InputItems() []objc.Object {
	rv := objc.Call[[]objc.Object](e_, objc.Sel("inputItems"))
	return rv
}
