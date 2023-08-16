// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [XPCListenerEndpoint] class.
var XPCListenerEndpointClass = _XPCListenerEndpointClass{objc.GetClass("NSXPCListenerEndpoint")}

type _XPCListenerEndpointClass struct {
	objc.Class
}

// An interface definition for the [XPCListenerEndpoint] class.
type IXPCListenerEndpoint interface {
	objc.IObject
}

// An object that names a specific XPC listener. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsxpclistenerendpoint?language=objc
type XPCListenerEndpoint struct {
	objc.Object
}

func XPCListenerEndpointFrom(ptr unsafe.Pointer) XPCListenerEndpoint {
	return XPCListenerEndpoint{
		Object: objc.ObjectFrom(ptr),
	}
}

func (xc _XPCListenerEndpointClass) Alloc() XPCListenerEndpoint {
	rv := objc.Call[XPCListenerEndpoint](xc, objc.Sel("alloc"))
	return rv
}

func XPCListenerEndpoint_Alloc() XPCListenerEndpoint {
	return XPCListenerEndpointClass.Alloc()
}

func (xc _XPCListenerEndpointClass) New() XPCListenerEndpoint {
	rv := objc.Call[XPCListenerEndpoint](xc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewXPCListenerEndpoint() XPCListenerEndpoint {
	return XPCListenerEndpointClass.New()
}

func (x_ XPCListenerEndpoint) Init() XPCListenerEndpoint {
	rv := objc.Call[XPCListenerEndpoint](x_, objc.Sel("init"))
	return rv
}
