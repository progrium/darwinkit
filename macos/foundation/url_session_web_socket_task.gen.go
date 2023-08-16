// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLSessionWebSocketTask] class.
var URLSessionWebSocketTaskClass = _URLSessionWebSocketTaskClass{objc.GetClass("NSURLSessionWebSocketTask")}

type _URLSessionWebSocketTaskClass struct {
	objc.Class
}

// An interface definition for the [URLSessionWebSocketTask] class.
type IURLSessionWebSocketTask interface {
	IURLSessionTask
	ReceiveMessageWithCompletionHandler(completionHandler func(message URLSessionWebSocketMessage, error Error))
	SendMessageCompletionHandler(message IURLSessionWebSocketMessage, completionHandler func(error Error))
	CancelWithCloseCodeReason(closeCode URLSessionWebSocketCloseCode, reason []byte)
	SendPingWithPongReceiveHandler(pongReceiveHandler func(error Error))
	CloseReason() []byte
	MaximumMessageSize() int
	SetMaximumMessageSize(value int)
	CloseCode() URLSessionWebSocketCloseCode
}

// A URL session task that communicates over the WebSockets protocol standard. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsockettask?language=objc
type URLSessionWebSocketTask struct {
	URLSessionTask
}

func URLSessionWebSocketTaskFrom(ptr unsafe.Pointer) URLSessionWebSocketTask {
	return URLSessionWebSocketTask{
		URLSessionTask: URLSessionTaskFrom(ptr),
	}
}

func (uc _URLSessionWebSocketTaskClass) Alloc() URLSessionWebSocketTask {
	rv := objc.Call[URLSessionWebSocketTask](uc, objc.Sel("alloc"))
	return rv
}

func URLSessionWebSocketTask_Alloc() URLSessionWebSocketTask {
	return URLSessionWebSocketTaskClass.Alloc()
}

func (uc _URLSessionWebSocketTaskClass) New() URLSessionWebSocketTask {
	rv := objc.Call[URLSessionWebSocketTask](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLSessionWebSocketTask() URLSessionWebSocketTask {
	return URLSessionWebSocketTaskClass.New()
}

func (u_ URLSessionWebSocketTask) Init() URLSessionWebSocketTask {
	rv := objc.Call[URLSessionWebSocketTask](u_, objc.Sel("init"))
	return rv
}

// Reads a WebSocket message once all the frames of the message are available. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsockettask/3181204-receivemessagewithcompletionhand?language=objc
func (u_ URLSessionWebSocketTask) ReceiveMessageWithCompletionHandler(completionHandler func(message URLSessionWebSocketMessage, error Error)) {
	objc.Call[objc.Void](u_, objc.Sel("receiveMessageWithCompletionHandler:"), completionHandler)
}

// Sends a WebSocket message, receiving the result in a completion handler. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsockettask/3181205-sendmessage?language=objc
func (u_ URLSessionWebSocketTask) SendMessageCompletionHandler(message IURLSessionWebSocketMessage, completionHandler func(error Error)) {
	objc.Call[objc.Void](u_, objc.Sel("sendMessage:completionHandler:"), objc.Ptr(message), completionHandler)
}

// Sends a close frame with the given close code and optional close reason. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsockettask/3181200-cancelwithclosecode?language=objc
func (u_ URLSessionWebSocketTask) CancelWithCloseCodeReason(closeCode URLSessionWebSocketCloseCode, reason []byte) {
	objc.Call[objc.Void](u_, objc.Sel("cancelWithCloseCode:reason:"), closeCode, reason)
}

// Sends a ping frame from the client side, with a closure to receive the pong from the server endpoint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsockettask/3181206-sendpingwithpongreceivehandler?language=objc
func (u_ URLSessionWebSocketTask) SendPingWithPongReceiveHandler(pongReceiveHandler func(error Error)) {
	objc.Call[objc.Void](u_, objc.Sel("sendPingWithPongReceiveHandler:"), pongReceiveHandler)
}

// A block of data that provides further information about why a connection closed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsockettask/3181202-closereason?language=objc
func (u_ URLSessionWebSocketTask) CloseReason() []byte {
	rv := objc.Call[[]byte](u_, objc.Sel("closeReason"))
	return rv
}

// The maximum number of bytes to buffer before the receive call fails with an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsockettask/3181203-maximummessagesize?language=objc
func (u_ URLSessionWebSocketTask) MaximumMessageSize() int {
	rv := objc.Call[int](u_, objc.Sel("maximumMessageSize"))
	return rv
}

// The maximum number of bytes to buffer before the receive call fails with an error. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsockettask/3181203-maximummessagesize?language=objc
func (u_ URLSessionWebSocketTask) SetMaximumMessageSize(value int) {
	objc.Call[objc.Void](u_, objc.Sel("setMaximumMessageSize:"), value)
}

// A code that indicates the reason a connection closed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsockettask/3181201-closecode?language=objc
func (u_ URLSessionWebSocketTask) CloseCode() URLSessionWebSocketCloseCode {
	rv := objc.Call[URLSessionWebSocketCloseCode](u_, objc.Sel("closeCode"))
	return rv
}
