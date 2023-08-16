// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"github.com/progrium/macdriver/objc"
)

// An interface for receiving messages from JavaScript code running in a webpage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkscriptmessagehandler?language=objc
type PScriptMessageHandler interface {
	// optional
	UserContentControllerDidReceiveScriptMessage(userContentController UserContentController, message ScriptMessage)
	HasUserContentControllerDidReceiveScriptMessage() bool
}

// A concrete type wrapper for the [PScriptMessageHandler] protocol.
type ScriptMessageHandlerWrapper struct {
	objc.Object
}

func (s_ ScriptMessageHandlerWrapper) HasUserContentControllerDidReceiveScriptMessage() bool {
	return s_.RespondsToSelector(objc.Sel("userContentController:didReceiveScriptMessage:"))
}

// Tells the handler that a webpage sent a script message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkscriptmessagehandler/1396222-usercontentcontroller?language=objc
func (s_ ScriptMessageHandlerWrapper) UserContentControllerDidReceiveScriptMessage(userContentController IUserContentController, message IScriptMessage) {
	objc.Call[objc.Void](s_, objc.Sel("userContentController:didReceiveScriptMessage:"), objc.Ptr(userContentController), objc.Ptr(message))
}
