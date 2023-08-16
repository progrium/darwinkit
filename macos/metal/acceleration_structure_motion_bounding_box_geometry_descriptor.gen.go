// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AccelerationStructureMotionBoundingBoxGeometryDescriptor] class.
var AccelerationStructureMotionBoundingBoxGeometryDescriptorClass = _AccelerationStructureMotionBoundingBoxGeometryDescriptorClass{objc.GetClass("MTLAccelerationStructureMotionBoundingBoxGeometryDescriptor")}

type _AccelerationStructureMotionBoundingBoxGeometryDescriptorClass struct {
	objc.Class
}

// An interface definition for the [AccelerationStructureMotionBoundingBoxGeometryDescriptor] class.
type IAccelerationStructureMotionBoundingBoxGeometryDescriptor interface {
	IAccelerationStructureGeometryDescriptor
	BoundingBoxStride() uint
	SetBoundingBoxStride(value uint)
	BoundingBoxCount() uint
	SetBoundingBoxCount(value uint)
	BoundingBoxBuffers() []MotionKeyframeData
	SetBoundingBoxBuffers(value []IMotionKeyframeData)
}

// A description of a list of bounding boxes, as motion keyframe data, to turn into an acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotionboundingboxgeometrydescriptor?language=objc
type AccelerationStructureMotionBoundingBoxGeometryDescriptor struct {
	AccelerationStructureGeometryDescriptor
}

func AccelerationStructureMotionBoundingBoxGeometryDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	return AccelerationStructureMotionBoundingBoxGeometryDescriptor{
		AccelerationStructureGeometryDescriptor: AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

func (ac _AccelerationStructureMotionBoundingBoxGeometryDescriptorClass) Descriptor() AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Call[AccelerationStructureMotionBoundingBoxGeometryDescriptor](ac, objc.Sel("descriptor"))
	return rv
}

// Creates a new bounding box descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotionboundingboxgeometrydescriptor/3750469-descriptor?language=objc
func AccelerationStructureMotionBoundingBoxGeometryDescriptor_Descriptor() AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	return AccelerationStructureMotionBoundingBoxGeometryDescriptorClass.Descriptor()
}

func (ac _AccelerationStructureMotionBoundingBoxGeometryDescriptorClass) Alloc() AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Call[AccelerationStructureMotionBoundingBoxGeometryDescriptor](ac, objc.Sel("alloc"))
	return rv
}

func AccelerationStructureMotionBoundingBoxGeometryDescriptor_Alloc() AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	return AccelerationStructureMotionBoundingBoxGeometryDescriptorClass.Alloc()
}

func (ac _AccelerationStructureMotionBoundingBoxGeometryDescriptorClass) New() AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Call[AccelerationStructureMotionBoundingBoxGeometryDescriptor](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAccelerationStructureMotionBoundingBoxGeometryDescriptor() AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	return AccelerationStructureMotionBoundingBoxGeometryDescriptorClass.New()
}

func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) Init() AccelerationStructureMotionBoundingBoxGeometryDescriptor {
	rv := objc.Call[AccelerationStructureMotionBoundingBoxGeometryDescriptor](a_, objc.Sel("init"))
	return rv
}

// The stride, in bytes, between bounding boxes in each buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotionboundingboxgeometrydescriptor/3750468-boundingboxstride?language=objc
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxStride() uint {
	rv := objc.Call[uint](a_, objc.Sel("boundingBoxStride"))
	return rv
}

// The stride, in bytes, between bounding boxes in each buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotionboundingboxgeometrydescriptor/3750468-boundingboxstride?language=objc
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxStride(value uint) {
	objc.Call[objc.Void](a_, objc.Sel("setBoundingBoxStride:"), value)
}

// The number of bounding boxes in each bounding box buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotionboundingboxgeometrydescriptor/3750467-boundingboxcount?language=objc
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxCount() uint {
	rv := objc.Call[uint](a_, objc.Sel("boundingBoxCount"))
	return rv
}

// The number of bounding boxes in each bounding box buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotionboundingboxgeometrydescriptor/3750467-boundingboxcount?language=objc
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxCount(value uint) {
	objc.Call[objc.Void](a_, objc.Sel("setBoundingBoxCount:"), value)
}

// A array of motion keyframes, each containing bounding box data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotionboundingboxgeometrydescriptor/3750466-boundingboxbuffers?language=objc
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) BoundingBoxBuffers() []MotionKeyframeData {
	rv := objc.Call[[]MotionKeyframeData](a_, objc.Sel("boundingBoxBuffers"))
	return rv
}

// A array of motion keyframes, each containing bounding box data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuremotionboundingboxgeometrydescriptor/3750466-boundingboxbuffers?language=objc
func (a_ AccelerationStructureMotionBoundingBoxGeometryDescriptor) SetBoundingBoxBuffers(value []IMotionKeyframeData) {
	objc.Call[objc.Void](a_, objc.Sel("setBoundingBoxBuffers:"), value)
}
