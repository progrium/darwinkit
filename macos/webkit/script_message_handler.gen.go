// auto-generated code, do not modify
package webkit

import (
	"github.com/progrium/macdriver/objc"
)

type IScriptMessageHandler interface {
	// required
	UserContentControllerDidReceiveScriptMessage(userContentController UserContentController, message ScriptMessage)
}

type ScriptMessageHandlerWrapper struct {
	objc.Object
}

func (s_ ScriptMessageHandlerWrapper) UserContentControllerDidReceiveScriptMessage(userContentController IUserContentController, message IScriptMessage) {
	objc.CallMethod[objc.Void](s_, objc.GetSelector("userContentController:didReceiveScriptMessage:"), objc.ExtractPtr(userContentController), objc.ExtractPtr(message))
}
