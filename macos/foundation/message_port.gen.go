// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MessagePort] class.
var MessagePortClass = _MessagePortClass{objc.GetClass("NSMessagePort")}

type _MessagePortClass struct {
	objc.Class
}

// An interface definition for the [MessagePort] class.
type IMessagePort interface {
	IPort
}

// A port that can be used as an endpoint for distributed object connections (or raw messaging). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmessageport?language=objc
type MessagePort struct {
	Port
}

func MessagePortFrom(ptr unsafe.Pointer) MessagePort {
	return MessagePort{
		Port: PortFrom(ptr),
	}
}

func (mc _MessagePortClass) Alloc() MessagePort {
	rv := objc.Call[MessagePort](mc, objc.Sel("alloc"))
	return rv
}

func MessagePort_Alloc() MessagePort {
	return MessagePortClass.Alloc()
}

func (mc _MessagePortClass) New() MessagePort {
	rv := objc.Call[MessagePort](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMessagePort() MessagePort {
	return MessagePortClass.New()
}

func (m_ MessagePort) Init() MessagePort {
	rv := objc.Call[MessagePort](m_, objc.Sel("init"))
	return rv
}
