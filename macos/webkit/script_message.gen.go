// auto-generated code, do not modify
package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

var ScriptMessageClass = _ScriptMessageClass{objc.GetClass("WKScriptMessage")}

type _ScriptMessageClass struct {
	objc.Class
}

type IScriptMessage interface {
	objc.IObject
	Body() objc.Object
	FrameInfo() FrameInfo
	WebView() WebView
	World() ContentWorld
	Name() string
}

type ScriptMessage struct {
	objc.Object
}

func MakeScriptMessage(ptr unsafe.Pointer) ScriptMessage {
	return ScriptMessage{
		Object: objc.MakeObject(ptr),
	}
}

func (sc _ScriptMessageClass) Alloc() ScriptMessage {
	rv := objc.CallMethod[ScriptMessage](sc, objc.GetSelector("alloc"))
	return rv
}

func ScriptMessage_Alloc() ScriptMessage {
	return ScriptMessageClass.Alloc()
}

func (sc _ScriptMessageClass) New() ScriptMessage {
	rv := objc.CallMethod[ScriptMessage](sc, objc.GetSelector("new"))
	rv.Autorelease()
	return rv
}

func NewScriptMessage() ScriptMessage {
	return ScriptMessageClass.New()
}

func ScriptMessage_New() ScriptMessage {
	return ScriptMessageClass.New()
}

func (s_ ScriptMessage) Init() ScriptMessage {
	rv := objc.CallMethod[ScriptMessage](s_, objc.GetSelector("init"))
	return rv
}

func ScriptMessage_Init() ScriptMessage {
	return ScriptMessageClass.Alloc().Init()
}

func (s_ ScriptMessage) Body() objc.Object {
	rv := objc.CallMethod[objc.Object](s_, objc.GetSelector("body"))
	return rv
}

func (s_ ScriptMessage) FrameInfo() FrameInfo {
	rv := objc.CallMethod[FrameInfo](s_, objc.GetSelector("frameInfo"))
	return rv
}

func (s_ ScriptMessage) WebView() WebView {
	rv := objc.CallMethod[WebView](s_, objc.GetSelector("webView"))
	return rv
}

func (s_ ScriptMessage) World() ContentWorld {
	rv := objc.CallMethod[ContentWorld](s_, objc.GetSelector("world"))
	return rv
}

func (s_ ScriptMessage) Name() string {
	rv := objc.CallMethod[string](s_, objc.GetSelector("name"))
	return rv
}
