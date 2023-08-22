// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NDArray] class.
var NDArrayClass = _NDArrayClass{objc.GetClass("MPSNDArray")}

type _NDArrayClass struct {
	objc.Class
}

// An interface definition for the [NDArray] class.
type INDArray interface {
	objc.IObject
	ArrayViewWithCommandBufferDescriptorAliasing(cmdBuf metal.PCommandBuffer, descriptor INDArrayDescriptor, aliasing AliasingStrategy) NDArray
	ArrayViewWithCommandBufferObjectDescriptorAliasing(cmdBufObject objc.IObject, descriptor INDArrayDescriptor, aliasing AliasingStrategy) NDArray
	ResourceSize() uint
	WriteBytesStrideBytes(buffer unsafe.Pointer, strideBytesPerDimension *int)
	LengthOfDimension(dimensionIndex uint) uint
	ReadBytesStrideBytes(buffer unsafe.Pointer, strideBytesPerDimension *int)
	ExportDataWithCommandBufferToImagesOffset(cmdBuf metal.PCommandBuffer, images *foundation.Array, offset ImageCoordinate)
	ExportDataWithCommandBufferObjectToImagesOffset(cmdBufObject objc.IObject, images *foundation.Array, offset ImageCoordinate)
	SynchronizeOnCommandBuffer(commandBuffer metal.PCommandBuffer)
	SynchronizeOnCommandBufferObject(commandBufferObject objc.IObject)
	Descriptor() NDArrayDescriptor
	ImportDataWithCommandBufferFromImagesOffset(cmdBuf metal.PCommandBuffer, images *foundation.Array, offset ImageCoordinate)
	ImportDataWithCommandBufferObjectFromImagesOffset(cmdBufObject objc.IObject, images *foundation.Array, offset ImageCoordinate)
	Parent() NDArray
	Device() metal.DeviceWrapper
	DataTypeSize() uint
	DataType() DataType
	Label() string
	SetLabel(value string)
	NumberOfDimensions() uint
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray?language=objc
type NDArray struct {
	objc.Object
}

func NDArrayFrom(ptr unsafe.Pointer) NDArray {
	return NDArray{
		Object: objc.ObjectFrom(ptr),
	}
}

func (n_ NDArray) InitWithDeviceScalar(device metal.PDevice, value float64) NDArray {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArray](n_, objc.Sel("initWithDevice:scalar:"), po0, value)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114051-initwithdevice?language=objc
func NewNDArrayWithDeviceScalar(device metal.PDevice, value float64) NDArray {
	instance := NDArrayClass.Alloc().InitWithDeviceScalar(device, value)
	instance.Autorelease()
	return instance
}

func (nc _NDArrayClass) Alloc() NDArray {
	rv := objc.Call[NDArray](nc, objc.Sel("alloc"))
	return rv
}

func NDArray_Alloc() NDArray {
	return NDArrayClass.Alloc()
}

func (nc _NDArrayClass) New() NDArray {
	rv := objc.Call[NDArray](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNDArray() NDArray {
	return NDArrayClass.New()
}

func (n_ NDArray) Init() NDArray {
	rv := objc.Call[NDArray](n_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114040-arrayviewwithcommandbuffer?language=objc
func (n_ NDArray) ArrayViewWithCommandBufferDescriptorAliasing(cmdBuf metal.PCommandBuffer, descriptor INDArrayDescriptor, aliasing AliasingStrategy) NDArray {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[NDArray](n_, objc.Sel("arrayViewWithCommandBuffer:descriptor:aliasing:"), po0, objc.Ptr(descriptor), aliasing)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114040-arrayviewwithcommandbuffer?language=objc
func (n_ NDArray) ArrayViewWithCommandBufferObjectDescriptorAliasing(cmdBufObject objc.IObject, descriptor INDArrayDescriptor, aliasing AliasingStrategy) NDArray {
	rv := objc.Call[NDArray](n_, objc.Sel("arrayViewWithCommandBuffer:descriptor:aliasing:"), objc.Ptr(cmdBufObject), objc.Ptr(descriptor), aliasing)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114058-resourcesize?language=objc
func (n_ NDArray) ResourceSize() uint {
	rv := objc.Call[uint](n_, objc.Sel("resourceSize"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114060-writebytes?language=objc
func (n_ NDArray) WriteBytesStrideBytes(buffer unsafe.Pointer, strideBytesPerDimension *int) {
	objc.Call[objc.Void](n_, objc.Sel("writeBytes:strideBytes:"), buffer, strideBytesPerDimension)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114053-lengthofdimension?language=objc
func (n_ NDArray) LengthOfDimension(dimensionIndex uint) uint {
	rv := objc.Call[uint](n_, objc.Sel("lengthOfDimension:"), dimensionIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114057-readbytes?language=objc
func (n_ NDArray) ReadBytesStrideBytes(buffer unsafe.Pointer, strideBytesPerDimension *int) {
	objc.Call[objc.Void](n_, objc.Sel("readBytes:strideBytes:"), buffer, strideBytesPerDimension)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3152526-exportdatawithcommandbuffer?language=objc
func (n_ NDArray) ExportDataWithCommandBufferToImagesOffset(cmdBuf metal.PCommandBuffer, images *foundation.Array, offset ImageCoordinate) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	objc.Call[objc.Void](n_, objc.Sel("exportDataWithCommandBuffer:toImages:offset:"), po0, images, offset)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3152526-exportdatawithcommandbuffer?language=objc
func (n_ NDArray) ExportDataWithCommandBufferObjectToImagesOffset(cmdBufObject objc.IObject, images *foundation.Array, offset ImageCoordinate) {
	objc.Call[objc.Void](n_, objc.Sel("exportDataWithCommandBuffer:toImages:offset:"), objc.Ptr(cmdBufObject), images, offset)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114059-synchronizeoncommandbuffer?language=objc
func (n_ NDArray) SynchronizeOnCommandBuffer(commandBuffer metal.PCommandBuffer) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("synchronizeOnCommandBuffer:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114059-synchronizeoncommandbuffer?language=objc
func (n_ NDArray) SynchronizeOnCommandBufferObject(commandBufferObject objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("synchronizeOnCommandBuffer:"), objc.Ptr(commandBufferObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114044-descriptor?language=objc
func (n_ NDArray) Descriptor() NDArrayDescriptor {
	rv := objc.Call[NDArrayDescriptor](n_, objc.Sel("descriptor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3152527-importdatawithcommandbuffer?language=objc
func (n_ NDArray) ImportDataWithCommandBufferFromImagesOffset(cmdBuf metal.PCommandBuffer, images *foundation.Array, offset ImageCoordinate) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	objc.Call[objc.Void](n_, objc.Sel("importDataWithCommandBuffer:fromImages:offset:"), po0, images, offset)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3152527-importdatawithcommandbuffer?language=objc
func (n_ NDArray) ImportDataWithCommandBufferObjectFromImagesOffset(cmdBufObject objc.IObject, images *foundation.Array, offset ImageCoordinate) {
	objc.Call[objc.Void](n_, objc.Sel("importDataWithCommandBuffer:fromImages:offset:"), objc.Ptr(cmdBufObject), images, offset)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3131728-defaultallocator?language=objc
func (nc _NDArrayClass) DefaultAllocator() NDArrayAllocatorWrapper {
	rv := objc.Call[NDArrayAllocatorWrapper](nc, objc.Sel("defaultAllocator"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3131728-defaultallocator?language=objc
func NDArray_DefaultAllocator() NDArrayAllocatorWrapper {
	return NDArrayClass.DefaultAllocator()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114056-parent?language=objc
func (n_ NDArray) Parent() NDArray {
	rv := objc.Call[NDArray](n_, objc.Sel("parent"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114045-device?language=objc
func (n_ NDArray) Device() metal.DeviceWrapper {
	rv := objc.Call[metal.DeviceWrapper](n_, objc.Sel("device"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114042-datatypesize?language=objc
func (n_ NDArray) DataTypeSize() uint {
	rv := objc.Call[uint](n_, objc.Sel("dataTypeSize"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114041-datatype?language=objc
func (n_ NDArray) DataType() DataType {
	rv := objc.Call[DataType](n_, objc.Sel("dataType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114052-label?language=objc
func (n_ NDArray) Label() string {
	rv := objc.Call[string](n_, objc.Sel("label"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114052-label?language=objc
func (n_ NDArray) SetLabel(value string) {
	objc.Call[objc.Void](n_, objc.Sel("setLabel:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarray/3114055-numberofdimensions?language=objc
func (n_ NDArray) NumberOfDimensions() uint {
	rv := objc.Call[uint](n_, objc.Sel("numberOfDimensions"))
	return rv
}
