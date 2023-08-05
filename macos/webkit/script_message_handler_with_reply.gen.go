// AUTO-GENERATED CODE, DO NOT MODIFY
package webkit

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

type IScriptMessageHandlerWithReply interface {
	// required
	UserContentControllerDidReceiveScriptMessageReplyHandler(userContentController UserContentController, message ScriptMessage, replyHandler func(reply objc.Object, errorMessage foundation.String))
}

type ScriptMessageHandlerWithReplyWrapper struct {
	objc.Object
}

func (s_ ScriptMessageHandlerWithReplyWrapper) UserContentControllerDidReceiveScriptMessageReplyHandler(userContentController IUserContentController, message IScriptMessage, replyHandler func(reply objc.Object, errorMessage foundation.String)) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("userContentController:didReceiveScriptMessage:replyHandler:"), objc.ExtractPtr(userContentController), objc.ExtractPtr(message), replyHandler)
}
