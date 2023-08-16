// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PortNameServer] class.
var PortNameServerClass = _PortNameServerClass{objc.GetClass("NSPortNameServer")}

type _PortNameServerClass struct {
	objc.Class
}

// An interface definition for the [PortNameServer] class.
type IPortNameServer interface {
	objc.IObject
}

// An object-oriented interface to the port registration service used by the distributed objects system. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsportnameserver?language=objc
type PortNameServer struct {
	objc.Object
}

func PortNameServerFrom(ptr unsafe.Pointer) PortNameServer {
	return PortNameServer{
		Object: objc.ObjectFrom(ptr),
	}
}

func (pc _PortNameServerClass) Alloc() PortNameServer {
	rv := objc.Call[PortNameServer](pc, objc.Sel("alloc"))
	return rv
}

func PortNameServer_Alloc() PortNameServer {
	return PortNameServerClass.Alloc()
}

func (pc _PortNameServerClass) New() PortNameServer {
	rv := objc.Call[PortNameServer](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPortNameServer() PortNameServer {
	return PortNameServerClass.New()
}

func (p_ PortNameServer) Init() PortNameServer {
	rv := objc.Call[PortNameServer](p_, objc.Sel("init"))
	return rv
}
