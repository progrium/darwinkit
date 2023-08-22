// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [PolygonAccelerationStructure] class.
var PolygonAccelerationStructureClass = _PolygonAccelerationStructureClass{objc.GetClass("MPSPolygonAccelerationStructure")}

type _PolygonAccelerationStructureClass struct {
	objc.Class
}

// An interface definition for the [PolygonAccelerationStructure] class.
type IPolygonAccelerationStructure interface {
	IAccelerationStructure
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpspolygonaccelerationstructure?language=objc
type PolygonAccelerationStructure struct {
	AccelerationStructure
}

func PolygonAccelerationStructureFrom(ptr unsafe.Pointer) PolygonAccelerationStructure {
	return PolygonAccelerationStructure{
		AccelerationStructure: AccelerationStructureFrom(ptr),
	}
}

func (pc _PolygonAccelerationStructureClass) Alloc() PolygonAccelerationStructure {
	rv := objc.Call[PolygonAccelerationStructure](pc, objc.Sel("alloc"))
	return rv
}

func PolygonAccelerationStructure_Alloc() PolygonAccelerationStructure {
	return PolygonAccelerationStructureClass.Alloc()
}

func (pc _PolygonAccelerationStructureClass) New() PolygonAccelerationStructure {
	rv := objc.Call[PolygonAccelerationStructure](pc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewPolygonAccelerationStructure() PolygonAccelerationStructure {
	return PolygonAccelerationStructureClass.New()
}

func (p_ PolygonAccelerationStructure) Init() PolygonAccelerationStructure {
	rv := objc.Call[PolygonAccelerationStructure](p_, objc.Sel("init"))
	return rv
}

func (p_ PolygonAccelerationStructure) InitWithDevice(device metal.PDevice) PolygonAccelerationStructure {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[PolygonAccelerationStructure](p_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewPolygonAccelerationStructureWithDevice(device metal.PDevice) PolygonAccelerationStructure {
	instance := PolygonAccelerationStructureClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (p_ PolygonAccelerationStructure) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) PolygonAccelerationStructure {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[PolygonAccelerationStructure](p_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func PolygonAccelerationStructure_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) PolygonAccelerationStructure {
	instance := PolygonAccelerationStructureClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
