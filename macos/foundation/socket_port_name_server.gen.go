// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SocketPortNameServer] class.
var SocketPortNameServerClass = _SocketPortNameServerClass{objc.GetClass("NSSocketPortNameServer")}

type _SocketPortNameServerClass struct {
	objc.Class
}

// An interface definition for the [SocketPortNameServer] class.
type ISocketPortNameServer interface {
	IPortNameServer
}

// A port name server that takes and returns socket ports. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nssocketportnameserver?language=objc
type SocketPortNameServer struct {
	PortNameServer
}

func SocketPortNameServerFrom(ptr unsafe.Pointer) SocketPortNameServer {
	return SocketPortNameServer{
		PortNameServer: PortNameServerFrom(ptr),
	}
}

func (sc _SocketPortNameServerClass) Alloc() SocketPortNameServer {
	rv := objc.Call[SocketPortNameServer](sc, objc.Sel("alloc"))
	return rv
}

func SocketPortNameServer_Alloc() SocketPortNameServer {
	return SocketPortNameServerClass.Alloc()
}

func (sc _SocketPortNameServerClass) New() SocketPortNameServer {
	rv := objc.Call[SocketPortNameServer](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSocketPortNameServer() SocketPortNameServer {
	return SocketPortNameServerClass.New()
}

func (s_ SocketPortNameServer) Init() SocketPortNameServer {
	rv := objc.Call[SocketPortNameServer](s_, objc.Sel("init"))
	return rv
}
