// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"github.com/progrium/macdriver/objc"
)

// An interface for responding to messages from JavaScript code running in a webpage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkscriptmessagehandlerwithreply?language=objc
type PScriptMessageHandlerWithReply interface {
	// optional
	UserContentControllerDidReceiveScriptMessageReplyHandler(userContentController UserContentController, message ScriptMessage, replyHandler func(reply objc.Object, errorMessage string))
	HasUserContentControllerDidReceiveScriptMessageReplyHandler() bool
}

// A concrete type wrapper for the [PScriptMessageHandlerWithReply] protocol.
type ScriptMessageHandlerWithReplyWrapper struct {
	objc.Object
}

func (s_ ScriptMessageHandlerWithReplyWrapper) HasUserContentControllerDidReceiveScriptMessageReplyHandler() bool {
	return s_.RespondsToSelector(objc.Sel("userContentController:didReceiveScriptMessage:replyHandler:"))
}

// Tells the handler that a webpage sent a script message that included a reply. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkscriptmessagehandlerwithreply/3585111-usercontentcontroller?language=objc
func (s_ ScriptMessageHandlerWithReplyWrapper) UserContentControllerDidReceiveScriptMessageReplyHandler(userContentController IUserContentController, message IScriptMessage, replyHandler func(reply objc.Object, errorMessage string)) {
	objc.Call[objc.Void](s_, objc.Sel("userContentController:didReceiveScriptMessage:replyHandler:"), objc.Ptr(userContentController), objc.Ptr(message), replyHandler)
}
