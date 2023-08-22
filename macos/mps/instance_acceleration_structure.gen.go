// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [InstanceAccelerationStructure] class.
var InstanceAccelerationStructureClass = _InstanceAccelerationStructureClass{objc.GetClass("MPSInstanceAccelerationStructure")}

type _InstanceAccelerationStructureClass struct {
	objc.Class
}

// An interface definition for the [InstanceAccelerationStructure] class.
type IInstanceAccelerationStructure interface {
	IAccelerationStructure
}

// An acceleration structure built over instances of other acceleration structures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsinstanceaccelerationstructure?language=objc
type InstanceAccelerationStructure struct {
	AccelerationStructure
}

func InstanceAccelerationStructureFrom(ptr unsafe.Pointer) InstanceAccelerationStructure {
	return InstanceAccelerationStructure{
		AccelerationStructure: AccelerationStructureFrom(ptr),
	}
}

func (ic _InstanceAccelerationStructureClass) Alloc() InstanceAccelerationStructure {
	rv := objc.Call[InstanceAccelerationStructure](ic, objc.Sel("alloc"))
	return rv
}

func InstanceAccelerationStructure_Alloc() InstanceAccelerationStructure {
	return InstanceAccelerationStructureClass.Alloc()
}

func (ic _InstanceAccelerationStructureClass) New() InstanceAccelerationStructure {
	rv := objc.Call[InstanceAccelerationStructure](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewInstanceAccelerationStructure() InstanceAccelerationStructure {
	return InstanceAccelerationStructureClass.New()
}

func (i_ InstanceAccelerationStructure) Init() InstanceAccelerationStructure {
	rv := objc.Call[InstanceAccelerationStructure](i_, objc.Sel("init"))
	return rv
}

func (i_ InstanceAccelerationStructure) InitWithDevice(device metal.PDevice) InstanceAccelerationStructure {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[InstanceAccelerationStructure](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewInstanceAccelerationStructureWithDevice(device metal.PDevice) InstanceAccelerationStructure {
	instance := InstanceAccelerationStructureClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ InstanceAccelerationStructure) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) InstanceAccelerationStructure {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[InstanceAccelerationStructure](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func InstanceAccelerationStructure_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) InstanceAccelerationStructure {
	instance := InstanceAccelerationStructureClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
