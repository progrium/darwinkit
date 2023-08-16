// AUTO-GENERATED CODE, DO NOT MODIFY

package webkit

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ScriptMessage] class.
var ScriptMessageClass = _ScriptMessageClass{objc.GetClass("WKScriptMessage")}

type _ScriptMessageClass struct {
	objc.Class
}

// An interface definition for the [ScriptMessage] class.
type IScriptMessage interface {
	objc.IObject
	Name() string
	Body() objc.Object
	FrameInfo() FrameInfo
	World() ContentWorld
}

// An object that encapsulates a message sent by JavaScript code from a webpage. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkscriptmessage?language=objc
type ScriptMessage struct {
	objc.Object
}

func ScriptMessageFrom(ptr unsafe.Pointer) ScriptMessage {
	return ScriptMessage{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _ScriptMessageClass) Alloc() ScriptMessage {
	rv := objc.Call[ScriptMessage](sc, objc.Sel("alloc"))
	return rv
}

func ScriptMessage_Alloc() ScriptMessage {
	return ScriptMessageClass.Alloc()
}

func (sc _ScriptMessageClass) New() ScriptMessage {
	rv := objc.Call[ScriptMessage](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewScriptMessage() ScriptMessage {
	return ScriptMessageClass.New()
}

func (s_ ScriptMessage) Init() ScriptMessage {
	rv := objc.Call[ScriptMessage](s_, objc.Sel("init"))
	return rv
}

// The name of the message handler to which the message is sent. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkscriptmessage/1417908-name?language=objc
func (s_ ScriptMessage) Name() string {
	rv := objc.Call[string](s_, objc.Sel("name"))
	return rv
}

// The body of the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkscriptmessage/1417901-body?language=objc
func (s_ ScriptMessage) Body() objc.Object {
	rv := objc.Call[objc.Object](s_, objc.Sel("body"))
	return rv
}

// The frame that sent the message. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkscriptmessage/1417906-frameinfo?language=objc
func (s_ ScriptMessage) FrameInfo() FrameInfo {
	rv := objc.Call[FrameInfo](s_, objc.Sel("frameInfo"))
	return rv
}

// The namespace in which the JavaScript code executes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/webkit/wkscriptmessage/3585109-world?language=objc
func (s_ ScriptMessage) World() ContentWorld {
	rv := objc.Call[ContentWorld](s_, objc.Sel("world"))
	return rv
}
