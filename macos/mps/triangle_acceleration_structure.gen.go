// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TriangleAccelerationStructure] class.
var TriangleAccelerationStructureClass = _TriangleAccelerationStructureClass{objc.GetClass("MPSTriangleAccelerationStructure")}

type _TriangleAccelerationStructureClass struct {
	objc.Class
}

// An interface definition for the [TriangleAccelerationStructure] class.
type ITriangleAccelerationStructure interface {
	IPolygonAccelerationStructure
}

// An acceleration structure built over triangles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstriangleaccelerationstructure?language=objc
type TriangleAccelerationStructure struct {
	PolygonAccelerationStructure
}

func TriangleAccelerationStructureFrom(ptr unsafe.Pointer) TriangleAccelerationStructure {
	return TriangleAccelerationStructure{
		PolygonAccelerationStructure: PolygonAccelerationStructureFrom(ptr),
	}
}

func (tc _TriangleAccelerationStructureClass) Alloc() TriangleAccelerationStructure {
	rv := objc.Call[TriangleAccelerationStructure](tc, objc.Sel("alloc"))
	return rv
}

func TriangleAccelerationStructure_Alloc() TriangleAccelerationStructure {
	return TriangleAccelerationStructureClass.Alloc()
}

func (tc _TriangleAccelerationStructureClass) New() TriangleAccelerationStructure {
	rv := objc.Call[TriangleAccelerationStructure](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTriangleAccelerationStructure() TriangleAccelerationStructure {
	return TriangleAccelerationStructureClass.New()
}

func (t_ TriangleAccelerationStructure) Init() TriangleAccelerationStructure {
	rv := objc.Call[TriangleAccelerationStructure](t_, objc.Sel("init"))
	return rv
}

func (t_ TriangleAccelerationStructure) InitWithDevice(device metal.PDevice) TriangleAccelerationStructure {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[TriangleAccelerationStructure](t_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewTriangleAccelerationStructureWithDevice(device metal.PDevice) TriangleAccelerationStructure {
	instance := TriangleAccelerationStructureClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (t_ TriangleAccelerationStructure) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) TriangleAccelerationStructure {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[TriangleAccelerationStructure](t_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func TriangleAccelerationStructure_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) TriangleAccelerationStructure {
	instance := TriangleAccelerationStructureClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
