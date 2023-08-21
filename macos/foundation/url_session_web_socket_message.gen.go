// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [URLSessionWebSocketMessage] class.
var URLSessionWebSocketMessageClass = _URLSessionWebSocketMessageClass{objc.GetClass("NSURLSessionWebSocketMessage")}

type _URLSessionWebSocketMessageClass struct {
	objc.Class
}

// An interface definition for the [URLSessionWebSocketMessage] class.
type IURLSessionWebSocketMessage interface {
	objc.IObject
	Data() []byte
	String() string
	Type() URLSessionWebSocketMessageType
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsocketmessage?language=objc
type URLSessionWebSocketMessage struct {
	objc.Object
}

func URLSessionWebSocketMessageFrom(ptr unsafe.Pointer) URLSessionWebSocketMessage {
	return URLSessionWebSocketMessage{
		Object: objc.ObjectFrom(ptr),
	}
}

func (u_ URLSessionWebSocketMessage) InitWithData(data []byte) URLSessionWebSocketMessage {
	rv := objc.Call[URLSessionWebSocketMessage](u_, objc.Sel("initWithData:"), data)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsocketmessage/3181192-initwithdata?language=objc
func NewURLSessionWebSocketMessageWithData(data []byte) URLSessionWebSocketMessage {
	instance := URLSessionWebSocketMessageClass.Alloc().InitWithData(data)
	instance.Autorelease()
	return instance
}

func (u_ URLSessionWebSocketMessage) InitWithString(string_ string) URLSessionWebSocketMessage {
	rv := objc.Call[URLSessionWebSocketMessage](u_, objc.Sel("initWithString:"), string_)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsocketmessage/3181193-initwithstring?language=objc
func NewURLSessionWebSocketMessageWithString(string_ string) URLSessionWebSocketMessage {
	instance := URLSessionWebSocketMessageClass.Alloc().InitWithString(string_)
	instance.Autorelease()
	return instance
}

func (uc _URLSessionWebSocketMessageClass) Alloc() URLSessionWebSocketMessage {
	rv := objc.Call[URLSessionWebSocketMessage](uc, objc.Sel("alloc"))
	return rv
}

func URLSessionWebSocketMessage_Alloc() URLSessionWebSocketMessage {
	return URLSessionWebSocketMessageClass.Alloc()
}

func (uc _URLSessionWebSocketMessageClass) New() URLSessionWebSocketMessage {
	rv := objc.Call[URLSessionWebSocketMessage](uc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewURLSessionWebSocketMessage() URLSessionWebSocketMessage {
	return URLSessionWebSocketMessageClass.New()
}

func (u_ URLSessionWebSocketMessage) Init() URLSessionWebSocketMessage {
	rv := objc.Call[URLSessionWebSocketMessage](u_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsocketmessage/3181191-data?language=objc
func (u_ URLSessionWebSocketMessage) Data() []byte {
	rv := objc.Call[[]byte](u_, objc.Sel("data"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsocketmessage/3181194-string?language=objc
func (u_ URLSessionWebSocketMessage) String() string {
	rv := objc.Call[string](u_, objc.Sel("string"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsurlsessionwebsocketmessage/3181195-type?language=objc
func (u_ URLSessionWebSocketMessage) Type() URLSessionWebSocketMessageType {
	rv := objc.Call[URLSessionWebSocketMessageType](u_, objc.Sel("type"))
	return rv
}
