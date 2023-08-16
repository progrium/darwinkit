// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MachBootstrapServer] class.
var MachBootstrapServerClass = _MachBootstrapServerClass{objc.GetClass("NSMachBootstrapServer")}

type _MachBootstrapServerClass struct {
	objc.Class
}

// An interface definition for the [MachBootstrapServer] class.
type IMachBootstrapServer interface {
	IPortNameServer
}

// A port name server that takes and returns Mach port objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmachbootstrapserver?language=objc
type MachBootstrapServer struct {
	PortNameServer
}

func MachBootstrapServerFrom(ptr unsafe.Pointer) MachBootstrapServer {
	return MachBootstrapServer{
		PortNameServer: PortNameServerFrom(ptr),
	}
}

func (mc _MachBootstrapServerClass) Alloc() MachBootstrapServer {
	rv := objc.Call[MachBootstrapServer](mc, objc.Sel("alloc"))
	return rv
}

func MachBootstrapServer_Alloc() MachBootstrapServer {
	return MachBootstrapServerClass.Alloc()
}

func (mc _MachBootstrapServerClass) New() MachBootstrapServer {
	rv := objc.Call[MachBootstrapServer](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMachBootstrapServer() MachBootstrapServer {
	return MachBootstrapServerClass.New()
}

func (m_ MachBootstrapServer) Init() MachBootstrapServer {
	rv := objc.Call[MachBootstrapServer](m_, objc.Sel("init"))
	return rv
}
