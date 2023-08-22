// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [Vector] class.
var VectorClass = _VectorClass{objc.GetClass("MPSVector")}

type _VectorClass struct {
	objc.Class
}

// An interface definition for the [Vector] class.
type IVector interface {
	objc.IObject
	ResourceSize() uint
	SynchronizeOnCommandBuffer(commandBuffer metal.PCommandBuffer)
	SynchronizeOnCommandBufferObject(commandBufferObject objc.IObject)
	Device() metal.DeviceWrapper
	VectorBytes() uint
	Data() metal.BufferWrapper
	Vectors() uint
	Length() uint
	Offset() uint
	DataType() DataType
}

// A 1D array of data that stores the data's values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector?language=objc
type Vector struct {
	objc.Object
}

func VectorFrom(ptr unsafe.Pointer) Vector {
	return Vector{
		Object: objc.ObjectFrom(ptr),
	}
}

func (v_ Vector) InitWithDeviceDescriptor(device metal.PDevice, descriptor IVectorDescriptor) Vector {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[Vector](v_, objc.Sel("initWithDevice:descriptor:"), po0, objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2942566-initwithdevice?language=objc
func NewVectorWithDeviceDescriptor(device metal.PDevice, descriptor IVectorDescriptor) Vector {
	instance := VectorClass.Alloc().InitWithDeviceDescriptor(device, descriptor)
	instance.Autorelease()
	return instance
}

func (v_ Vector) InitWithBufferDescriptor(buffer metal.PBuffer, descriptor IVectorDescriptor) Vector {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	rv := objc.Call[Vector](v_, objc.Sel("initWithBuffer:descriptor:"), po0, objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873346-initwithbuffer?language=objc
func NewVectorWithBufferDescriptor(buffer metal.PBuffer, descriptor IVectorDescriptor) Vector {
	instance := VectorClass.Alloc().InitWithBufferDescriptor(buffer, descriptor)
	instance.Autorelease()
	return instance
}

func (vc _VectorClass) Alloc() Vector {
	rv := objc.Call[Vector](vc, objc.Sel("alloc"))
	return rv
}

func Vector_Alloc() Vector {
	return VectorClass.Alloc()
}

func (vc _VectorClass) New() Vector {
	rv := objc.Call[Vector](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVector() Vector {
	return VectorClass.New()
}

func (v_ Vector) Init() Vector {
	rv := objc.Call[Vector](v_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2942570-resourcesize?language=objc
func (v_ Vector) ResourceSize() uint {
	rv := objc.Call[uint](v_, objc.Sel("resourceSize"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2942568-synchronizeoncommandbuffer?language=objc
func (v_ Vector) SynchronizeOnCommandBuffer(commandBuffer metal.PCommandBuffer) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](v_, objc.Sel("synchronizeOnCommandBuffer:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2942568-synchronizeoncommandbuffer?language=objc
func (v_ Vector) SynchronizeOnCommandBufferObject(commandBufferObject objc.IObject) {
	objc.Call[objc.Void](v_, objc.Sel("synchronizeOnCommandBuffer:"), objc.Ptr(commandBufferObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873338-device?language=objc
func (v_ Vector) Device() metal.DeviceWrapper {
	rv := objc.Call[metal.DeviceWrapper](v_, objc.Sel("device"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873340-vectorbytes?language=objc
func (v_ Vector) VectorBytes() uint {
	rv := objc.Call[uint](v_, objc.Sel("vectorBytes"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873393-data?language=objc
func (v_ Vector) Data() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](v_, objc.Sel("data"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873388-vectors?language=objc
func (v_ Vector) Vectors() uint {
	rv := objc.Call[uint](v_, objc.Sel("vectors"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873392-length?language=objc
func (v_ Vector) Length() uint {
	rv := objc.Call[uint](v_, objc.Sel("length"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/3375741-offset?language=objc
func (v_ Vector) Offset() uint {
	rv := objc.Call[uint](v_, objc.Sel("offset"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvector/2873336-datatype?language=objc
func (v_ Vector) DataType() DataType {
	rv := objc.Call[DataType](v_, objc.Sel("dataType"))
	return rv
}
