// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AccelerationStructureGeometryDescriptor] class.
var AccelerationStructureGeometryDescriptorClass = _AccelerationStructureGeometryDescriptorClass{objc.GetClass("MTLAccelerationStructureGeometryDescriptor")}

type _AccelerationStructureGeometryDescriptorClass struct {
	objc.Class
}

// An interface definition for the [AccelerationStructureGeometryDescriptor] class.
type IAccelerationStructureGeometryDescriptor interface {
	objc.IObject
	AllowDuplicateIntersectionFunctionInvocation() bool
	SetAllowDuplicateIntersectionFunctionInvocation(value bool)
	Opaque() bool
	SetOpaque(value bool)
	Label() string
	SetLabel(value string)
	IntersectionFunctionTableOffset() uint
	SetIntersectionFunctionTableOffset(value uint)
}

// A base class for descriptors that contain geometry data to convert into a ray-tracing acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor?language=objc
type AccelerationStructureGeometryDescriptor struct {
	objc.Object
}

func AccelerationStructureGeometryDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureGeometryDescriptor {
	return AccelerationStructureGeometryDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (ac _AccelerationStructureGeometryDescriptorClass) Alloc() AccelerationStructureGeometryDescriptor {
	rv := objc.Call[AccelerationStructureGeometryDescriptor](ac, objc.Sel("alloc"))
	return rv
}

func AccelerationStructureGeometryDescriptor_Alloc() AccelerationStructureGeometryDescriptor {
	return AccelerationStructureGeometryDescriptorClass.Alloc()
}

func (ac _AccelerationStructureGeometryDescriptorClass) New() AccelerationStructureGeometryDescriptor {
	rv := objc.Call[AccelerationStructureGeometryDescriptor](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAccelerationStructureGeometryDescriptor() AccelerationStructureGeometryDescriptor {
	return AccelerationStructureGeometryDescriptorClass.New()
}

func (a_ AccelerationStructureGeometryDescriptor) Init() AccelerationStructureGeometryDescriptor {
	rv := objc.Call[AccelerationStructureGeometryDescriptor](a_, objc.Sel("init"))
	return rv
}

// A Boolean value that indicates whether Metal calls the ray-intersection test more than once per primitive on the structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/3666578-allowduplicateintersectionfuncti?language=objc
func (a_ AccelerationStructureGeometryDescriptor) AllowDuplicateIntersectionFunctionInvocation() bool {
	rv := objc.Call[bool](a_, objc.Sel("allowDuplicateIntersectionFunctionInvocation"))
	return rv
}

// A Boolean value that indicates whether Metal calls the ray-intersection test more than once per primitive on the structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/3666578-allowduplicateintersectionfuncti?language=objc
func (a_ AccelerationStructureGeometryDescriptor) SetAllowDuplicateIntersectionFunctionInvocation(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setAllowDuplicateIntersectionFunctionInvocation:"), value)
}

// A Boolean value that determines whether the geometry data in the acceleration structure needs to skip triangle-intersection tests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/3666579-opaque?language=objc
func (a_ AccelerationStructureGeometryDescriptor) Opaque() bool {
	rv := objc.Call[bool](a_, objc.Sel("opaque"))
	return rv
}

// A Boolean value that determines whether the geometry data in the acceleration structure needs to skip triangle-intersection tests. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/3666579-opaque?language=objc
func (a_ AccelerationStructureGeometryDescriptor) SetOpaque(value bool) {
	objc.Call[objc.Void](a_, objc.Sel("setOpaque:"), value)
}

// A label for the geometry structure, suitable for debugging. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/3750460-label?language=objc
func (a_ AccelerationStructureGeometryDescriptor) Label() string {
	rv := objc.Call[string](a_, objc.Sel("label"))
	return rv
}

// A label for the geometry structure, suitable for debugging. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/3750460-label?language=objc
func (a_ AccelerationStructureGeometryDescriptor) SetLabel(value string) {
	objc.Call[objc.Void](a_, objc.Sel("setLabel:"), value)
}

// An index into the intersection table for determining which intersection function Metal calls when it intersects a ray with the acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/3566524-intersectionfunctiontableoffset?language=objc
func (a_ AccelerationStructureGeometryDescriptor) IntersectionFunctionTableOffset() uint {
	rv := objc.Call[uint](a_, objc.Sel("intersectionFunctionTableOffset"))
	return rv
}

// An index into the intersection table for determining which intersection function Metal calls when it intersects a ray with the acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuregeometrydescriptor/3566524-intersectionfunctiontableoffset?language=objc
func (a_ AccelerationStructureGeometryDescriptor) SetIntersectionFunctionTableOffset(value uint) {
	objc.Call[objc.Void](a_, objc.Sel("setIntersectionFunctionTableOffset:"), value)
}
