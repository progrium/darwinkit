// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [InstanceAccelerationStructureDescriptor] class.
var InstanceAccelerationStructureDescriptorClass = _InstanceAccelerationStructureDescriptorClass{objc.GetClass("MTLInstanceAccelerationStructureDescriptor")}

type _InstanceAccelerationStructureDescriptorClass struct {
	objc.Class
}

// An interface definition for the [InstanceAccelerationStructureDescriptor] class.
type IInstanceAccelerationStructureDescriptor interface {
	IAccelerationStructureDescriptor
	InstanceDescriptorBuffer() BufferWrapper
	SetInstanceDescriptorBuffer(value PBuffer)
	SetInstanceDescriptorBufferObject(valueObject objc.IObject)
	MotionTransformBuffer() BufferWrapper
	SetMotionTransformBuffer(value PBuffer)
	SetMotionTransformBufferObject(valueObject objc.IObject)
	InstanceDescriptorType() AccelerationStructureInstanceDescriptorType
	SetInstanceDescriptorType(value AccelerationStructureInstanceDescriptorType)
	InstanceDescriptorBufferOffset() uint
	SetInstanceDescriptorBufferOffset(value uint)
	InstancedAccelerationStructures() []AccelerationStructureWrapper
	SetInstancedAccelerationStructures(value []PAccelerationStructure)
	InstanceDescriptorStride() uint
	SetInstanceDescriptorStride(value uint)
	InstanceCount() uint
	SetInstanceCount(value uint)
	MotionTransformCount() uint
	SetMotionTransformCount(value uint)
	MotionTransformBufferOffset() uint
	SetMotionTransformBufferOffset(value uint)
}

// A description of an acceleration structure that derives from instances of primitive acceleration structures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor?language=objc
type InstanceAccelerationStructureDescriptor struct {
	AccelerationStructureDescriptor
}

func InstanceAccelerationStructureDescriptorFrom(ptr unsafe.Pointer) InstanceAccelerationStructureDescriptor {
	return InstanceAccelerationStructureDescriptor{
		AccelerationStructureDescriptor: AccelerationStructureDescriptorFrom(ptr),
	}
}

func (ic _InstanceAccelerationStructureDescriptorClass) Descriptor() InstanceAccelerationStructureDescriptor {
	rv := objc.Call[InstanceAccelerationStructureDescriptor](ic, objc.Sel("descriptor"))
	return rv
}

// Creates an instance descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3553884-descriptor?language=objc
func InstanceAccelerationStructureDescriptor_Descriptor() InstanceAccelerationStructureDescriptor {
	return InstanceAccelerationStructureDescriptorClass.Descriptor()
}

func (ic _InstanceAccelerationStructureDescriptorClass) Alloc() InstanceAccelerationStructureDescriptor {
	rv := objc.Call[InstanceAccelerationStructureDescriptor](ic, objc.Sel("alloc"))
	return rv
}

func InstanceAccelerationStructureDescriptor_Alloc() InstanceAccelerationStructureDescriptor {
	return InstanceAccelerationStructureDescriptorClass.Alloc()
}

func (ic _InstanceAccelerationStructureDescriptorClass) New() InstanceAccelerationStructureDescriptor {
	rv := objc.Call[InstanceAccelerationStructureDescriptor](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewInstanceAccelerationStructureDescriptor() InstanceAccelerationStructureDescriptor {
	return InstanceAccelerationStructureDescriptorClass.New()
}

func (i_ InstanceAccelerationStructureDescriptor) Init() InstanceAccelerationStructureDescriptor {
	rv := objc.Call[InstanceAccelerationStructureDescriptor](i_, objc.Sel("init"))
	return rv
}

// A buffer that contains descriptions of each instance in the acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3553886-instancedescriptorbuffer?language=objc
func (i_ InstanceAccelerationStructureDescriptor) InstanceDescriptorBuffer() BufferWrapper {
	rv := objc.Call[BufferWrapper](i_, objc.Sel("instanceDescriptorBuffer"))
	return rv
}

// A buffer that contains descriptions of each instance in the acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3553886-instancedescriptorbuffer?language=objc
func (i_ InstanceAccelerationStructureDescriptor) SetInstanceDescriptorBuffer(value PBuffer) {
	po0 := objc.WrapAsProtocol("MTLBuffer", value)
	objc.Call[objc.Void](i_, objc.Sel("setInstanceDescriptorBuffer:"), po0)
}

// A buffer that contains descriptions of each instance in the acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3553886-instancedescriptorbuffer?language=objc
func (i_ InstanceAccelerationStructureDescriptor) SetInstanceDescriptorBufferObject(valueObject objc.IObject) {
	objc.Call[objc.Void](i_, objc.Sel("setInstanceDescriptorBuffer:"), objc.Ptr(valueObject))
}

// A buffer that contains descriptions of each motion transform in the acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3750499-motiontransformbuffer?language=objc
func (i_ InstanceAccelerationStructureDescriptor) MotionTransformBuffer() BufferWrapper {
	rv := objc.Call[BufferWrapper](i_, objc.Sel("motionTransformBuffer"))
	return rv
}

// A buffer that contains descriptions of each motion transform in the acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3750499-motiontransformbuffer?language=objc
func (i_ InstanceAccelerationStructureDescriptor) SetMotionTransformBuffer(value PBuffer) {
	po0 := objc.WrapAsProtocol("MTLBuffer", value)
	objc.Call[objc.Void](i_, objc.Sel("setMotionTransformBuffer:"), po0)
}

// A buffer that contains descriptions of each motion transform in the acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3750499-motiontransformbuffer?language=objc
func (i_ InstanceAccelerationStructureDescriptor) SetMotionTransformBufferObject(valueObject objc.IObject) {
	objc.Call[objc.Void](i_, objc.Sel("setMotionTransformBuffer:"), objc.Ptr(valueObject))
}

// The format of the instance data in the descriptor buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3750498-instancedescriptortype?language=objc
func (i_ InstanceAccelerationStructureDescriptor) InstanceDescriptorType() AccelerationStructureInstanceDescriptorType {
	rv := objc.Call[AccelerationStructureInstanceDescriptorType](i_, objc.Sel("instanceDescriptorType"))
	return rv
}

// The format of the instance data in the descriptor buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3750498-instancedescriptortype?language=objc
func (i_ InstanceAccelerationStructureDescriptor) SetInstanceDescriptorType(value AccelerationStructureInstanceDescriptorType) {
	objc.Call[objc.Void](i_, objc.Sel("setInstanceDescriptorType:"), value)
}

// The offset, in bytes, to the descripton of the first instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3553887-instancedescriptorbufferoffset?language=objc
func (i_ InstanceAccelerationStructureDescriptor) InstanceDescriptorBufferOffset() uint {
	rv := objc.Call[uint](i_, objc.Sel("instanceDescriptorBufferOffset"))
	return rv
}

// The offset, in bytes, to the descripton of the first instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3553887-instancedescriptorbufferoffset?language=objc
func (i_ InstanceAccelerationStructureDescriptor) SetInstanceDescriptorBufferOffset(value uint) {
	objc.Call[objc.Void](i_, objc.Sel("setInstanceDescriptorBufferOffset:"), value)
}

// The bottom-level acceleration structures that instances use in the instance acceleration structure . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3553889-instancedaccelerationstructures?language=objc
func (i_ InstanceAccelerationStructureDescriptor) InstancedAccelerationStructures() []AccelerationStructureWrapper {
	rv := objc.Call[[]AccelerationStructureWrapper](i_, objc.Sel("instancedAccelerationStructures"))
	return rv
}

// The bottom-level acceleration structures that instances use in the instance acceleration structure . [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3553889-instancedaccelerationstructures?language=objc
func (i_ InstanceAccelerationStructureDescriptor) SetInstancedAccelerationStructures(value []PAccelerationStructure) {
	objc.Call[objc.Void](i_, objc.Sel("setInstancedAccelerationStructures:"), value)
}

// The stride, in bytes, between instance descriptions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3553888-instancedescriptorstride?language=objc
func (i_ InstanceAccelerationStructureDescriptor) InstanceDescriptorStride() uint {
	rv := objc.Call[uint](i_, objc.Sel("instanceDescriptorStride"))
	return rv
}

// The stride, in bytes, between instance descriptions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3553888-instancedescriptorstride?language=objc
func (i_ InstanceAccelerationStructureDescriptor) SetInstanceDescriptorStride(value uint) {
	objc.Call[objc.Void](i_, objc.Sel("setInstanceDescriptorStride:"), value)
}

// The number of instances in the instance descriptor buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3553885-instancecount?language=objc
func (i_ InstanceAccelerationStructureDescriptor) InstanceCount() uint {
	rv := objc.Call[uint](i_, objc.Sel("instanceCount"))
	return rv
}

// The number of instances in the instance descriptor buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3553885-instancecount?language=objc
func (i_ InstanceAccelerationStructureDescriptor) SetInstanceCount(value uint) {
	objc.Call[objc.Void](i_, objc.Sel("setInstanceCount:"), value)
}

// The number of motion transforms in the motion transform buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3750501-motiontransformcount?language=objc
func (i_ InstanceAccelerationStructureDescriptor) MotionTransformCount() uint {
	rv := objc.Call[uint](i_, objc.Sel("motionTransformCount"))
	return rv
}

// The number of motion transforms in the motion transform buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3750501-motiontransformcount?language=objc
func (i_ InstanceAccelerationStructureDescriptor) SetMotionTransformCount(value uint) {
	objc.Call[objc.Void](i_, objc.Sel("setMotionTransformCount:"), value)
}

// The offset, in bytes, to the descripton of the first motion transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3750500-motiontransformbufferoffset?language=objc
func (i_ InstanceAccelerationStructureDescriptor) MotionTransformBufferOffset() uint {
	rv := objc.Call[uint](i_, objc.Sel("motionTransformBufferOffset"))
	return rv
}

// The offset, in bytes, to the descripton of the first motion transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlinstanceaccelerationstructuredescriptor/3750500-motiontransformbufferoffset?language=objc
func (i_ InstanceAccelerationStructureDescriptor) SetMotionTransformBufferOffset(value uint) {
	objc.Call[objc.Void](i_, objc.Sel("setMotionTransformBufferOffset:"), value)
}
