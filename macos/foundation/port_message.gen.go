// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PortMessage] class.
var PortMessageClass = _PortMessageClass{objc.GetClass("NSPortMessage")}

type _PortMessageClass struct {
	objc.Class
}

// An interface definition for the [PortMessage] class.
type IPortMessage interface {
	objc.IObject
	SendBeforeDate(date IDate) bool
	Components() []objc.Object
	SendPort() Port
	Msgid() uint32
	SetMsgid(value uint32)
	ReceivePort() Port
}

// A low-level, operating system-independent type for inter-application (and inter-thread) messages. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsportmessage?language=objc
type PortMessage struct {
	objc.Object
}

func PortMessageFrom(ptr unsafe.Pointer) PortMessage {
	return PortMessage{
		Object: objc.ObjectFrom(ptr),
	}
}

func (p_ PortMessage) InitWithSendPortReceivePortComponents(sendPort IPort, replyPort IPort, components []objc.IObject) PortMessage {
	rv := objc.Call[PortMessage](p_, objc.Sel("initWithSendPort:receivePort:components:"), objc.Ptr(sendPort), objc.Ptr(replyPort), components)
	return rv
}

// Initializes a newly allocated NSPortMessage object to send given data on a given port and to receiver replies on another given port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsportmessage/1417387-initwithsendport?language=objc
func PortMessage_InitWithSendPortReceivePortComponents(sendPort IPort, replyPort IPort, components []objc.IObject) PortMessage {
	return PortMessageClass.Alloc().InitWithSendPortReceivePortComponents(sendPort, replyPort, components)
}

func (pc _PortMessageClass) Alloc() PortMessage {
	rv := objc.Call[PortMessage](pc, objc.Sel("alloc"))
	return rv
}

func PortMessage_Alloc() PortMessage {
	return PortMessageClass.Alloc()
}

func (pc _PortMessageClass) New() PortMessage {
	rv := objc.Call[PortMessage](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPortMessage() PortMessage {
	return PortMessageClass.New()
}

func (p_ PortMessage) Init() PortMessage {
	rv := objc.Call[PortMessage](p_, objc.Sel("init"))
	return rv
}

// Attempts to send the message before aDate, returning YES if successful or NO if the operation times out. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsportmessage/1407464-sendbeforedate?language=objc
func (p_ PortMessage) SendBeforeDate(date IDate) bool {
	rv := objc.Call[bool](p_, objc.Sel("sendBeforeDate:"), objc.Ptr(date))
	return rv
}

// Returns the data components of the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsportmessage/1407377-components?language=objc
func (p_ PortMessage) Components() []objc.Object {
	rv := objc.Call[[]objc.Object](p_, objc.Sel("components"))
	return rv
}

// For an outgoing message, returns the port the receiver will send itself through. For an incoming message, returns the port replies to the receiver should be sent through. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsportmessage/1417234-sendport?language=objc
func (p_ PortMessage) SendPort() Port {
	rv := objc.Call[Port](p_, objc.Sel("sendPort"))
	return rv
}

// Returns the identifier for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsportmessage/1407880-msgid?language=objc
func (p_ PortMessage) Msgid() uint32 {
	rv := objc.Call[uint32](p_, objc.Sel("msgid"))
	return rv
}

// Returns the identifier for the receiver. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsportmessage/1407880-msgid?language=objc
func (p_ PortMessage) SetMsgid(value uint32) {
	objc.Call[objc.Void](p_, objc.Sel("setMsgid:"), value)
}

// For an outgoing message, returns the port on which replies to the receiver will arrive. For an incoming message, returns the port the receiver did arrive on. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsportmessage/1413908-receiveport?language=objc
func (p_ PortMessage) ReceivePort() Port {
	rv := objc.Call[Port](p_, objc.Sel("receivePort"))
	return rv
}
