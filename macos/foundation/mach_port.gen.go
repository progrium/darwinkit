// AUTO-GENERATED CODE, DO NOT MODIFY

package foundation

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MachPort] class.
var MachPortClass = _MachPortClass{objc.GetClass("NSMachPort")}

type _MachPortClass struct {
	objc.Class
}

// An interface definition for the [MachPort] class.
type IMachPort interface {
	IPort
	MachPort() uint32
}

// A port that can be used as an endpoint for distributed object connections (or raw messaging). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmachport?language=objc
type MachPort struct {
	Port
}

func MachPortFrom(ptr unsafe.Pointer) MachPort {
	return MachPort{
		Port: PortFrom(ptr),
	}
}

func (m_ MachPort) InitWithMachPort(machPort uint32) MachPort {
	rv := objc.Call[MachPort](m_, objc.Sel("initWithMachPort:"), machPort)
	return rv
}

// Initializes a newly allocated NSMachPort object with a given Mach port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmachport/1399499-initwithmachport?language=objc
func NewMachPortWithMachPort(machPort uint32) MachPort {
	instance := MachPortClass.Alloc().InitWithMachPort(machPort)
	instance.Autorelease()
	return instance
}

func (mc _MachPortClass) Alloc() MachPort {
	rv := objc.Call[MachPort](mc, objc.Sel("alloc"))
	return rv
}

func MachPort_Alloc() MachPort {
	return MachPortClass.Alloc()
}

func (mc _MachPortClass) New() MachPort {
	rv := objc.Call[MachPort](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMachPort() MachPort {
	return MachPortClass.New()
}

func (m_ MachPort) Init() MachPort {
	rv := objc.Call[MachPort](m_, objc.Sel("init"))
	return rv
}

// Creates and returns a port object configured with the given Mach port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmachport/1399478-portwithmachport?language=objc
func (mc _MachPortClass) PortWithMachPort(machPort uint32) Port {
	rv := objc.Call[Port](mc, objc.Sel("portWithMachPort:"), machPort)
	return rv
}

// Creates and returns a port object configured with the given Mach port. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmachport/1399478-portwithmachport?language=objc
func MachPort_PortWithMachPort(machPort uint32) Port {
	return MachPortClass.PortWithMachPort(machPort)
}

// The Mach port used by the receiver, represented as an integer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/foundation/nsmachport/1399539-machport?language=objc
func (m_ MachPort) MachPort() uint32 {
	rv := objc.Call[uint32](m_, objc.Sel("machPort"))
	return rv
}
