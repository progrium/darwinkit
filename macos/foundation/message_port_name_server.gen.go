// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MessagePortNameServer] class.
var MessagePortNameServerClass = _MessagePortNameServerClass{objc.GetClass("NSMessagePortNameServer")}

type _MessagePortNameServerClass struct {
	objc.Class
}

// An interface definition for the [MessagePortNameServer] class.
type IMessagePortNameServer interface {
	IPortNameServer
}

// A server takes and returns message ports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmessageportnameserver?language=objc
type MessagePortNameServer struct {
	PortNameServer
}

func MessagePortNameServerFrom(ptr unsafe.Pointer) MessagePortNameServer {
	return MessagePortNameServer{
		PortNameServer: PortNameServerFrom(ptr),
	}
}

func (mc _MessagePortNameServerClass) Alloc() MessagePortNameServer {
	rv := objc.Call[MessagePortNameServer](mc, objc.Sel("alloc"))
	return rv
}

func MessagePortNameServer_Alloc() MessagePortNameServer {
	return MessagePortNameServerClass.Alloc()
}

func (mc _MessagePortNameServerClass) New() MessagePortNameServer {
	rv := objc.Call[MessagePortNameServer](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMessagePortNameServer() MessagePortNameServer {
	return MessagePortNameServerClass.New()
}

func (m_ MessagePortNameServer) Init() MessagePortNameServer {
	rv := objc.Call[MessagePortNameServer](m_, objc.Sel("init"))
	return rv
}
