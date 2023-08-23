// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNKernel] class.
var CNNKernelClass = _CNNKernelClass{objc.GetClass("MPSCNNKernel")}

type _CNNKernelClass struct {
	objc.Class
}

// An interface definition for the [CNNKernel] class.
type ICNNKernel interface {
	IKernel
	BatchEncodingStorageSizeForSourceImageSourceStatesDestinationImage(sourceImage *foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) uint
	TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer metal.PCommandBuffer, sourceImage IImage, sourceStates []IState, destinationImage IImage) State
	TemporaryResultStateForCommandBufferObjectSourceImageSourceStatesDestinationImage(commandBufferObject objc.IObject, sourceImage IImage, sourceStates []IState, destinationImage IImage) State
	AppendBatchBarrier() bool
	TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer metal.PCommandBuffer, sourceImage *foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) *foundation.Array
	TemporaryResultStateBatchForCommandBufferObjectSourceImageSourceStatesDestinationImage(commandBufferObject objc.IObject, sourceImage *foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) *foundation.Array
	DestinationImageDescriptorForSourceImagesSourceStates(sourceImages []IImage, sourceStates []IState) ImageDescriptor
	EncodingStorageSizeForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates []IState, destinationImage IImage) uint
	EncodeBatchToCommandBufferSourceImages(commandBuffer metal.PCommandBuffer, sourceImages *foundation.Array) *foundation.Array
	EncodeBatchToCommandBufferObjectSourceImages(commandBufferObject objc.IObject, sourceImages *foundation.Array) *foundation.Array
	ResultStateBatchForSourceImageSourceStatesDestinationImage(sourceImage *foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) *foundation.Array
	EncodeToCommandBufferSourceImage(commandBuffer metal.PCommandBuffer, sourceImage IImage) Image
	EncodeToCommandBufferObjectSourceImage(commandBufferObject objc.IObject, sourceImage IImage) Image
	IsResultStateReusedAcrossBatch() bool
	ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates []IState, destinationImage IImage) State
	SourceFeatureChannelMaxCount() uint
	SetSourceFeatureChannelMaxCount(value uint)
	IsStateModified() bool
	Padding() NNPaddingWrapper
	SetPadding(value PNNPadding)
	SetPaddingObject(valueObject objc.IObject)
	SourceFeatureChannelOffset() uint
	SetSourceFeatureChannelOffset(value uint)
	StrideInPixelsY() uint
	StrideInPixelsX() uint
	EdgeMode() ImageEdgeMode
	SetEdgeMode(value ImageEdgeMode)
	Offset() Offset
	SetOffset(value Offset)
	IsBackwards() bool
	DilationRateY() uint
	KernelHeight() uint
	ClipRect() metal.Region
	SetClipRect(value metal.Region)
	KernelWidth() uint
	DestinationImageAllocator() ImageAllocatorWrapper
	SetDestinationImageAllocator(value PImageAllocator)
	SetDestinationImageAllocatorObject(valueObject objc.IObject)
	DestinationFeatureChannelOffset() uint
	SetDestinationFeatureChannelOffset(value uint)
	DilationRateX() uint
}

// Base class for neural network layers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel?language=objc
type CNNKernel struct {
	Kernel
}

func CNNKernelFrom(ptr unsafe.Pointer) CNNKernel {
	return CNNKernel{
		Kernel: KernelFrom(ptr),
	}
}

func (c_ CNNKernel) InitWithDevice(device metal.PDevice) CNNKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNKernel](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNKernelWithDevice(device metal.PDevice) CNNKernel {
	instance := CNNKernelClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (cc _CNNKernelClass) Alloc() CNNKernel {
	rv := objc.Call[CNNKernel](cc, objc.Sel("alloc"))
	return rv
}

func CNNKernel_Alloc() CNNKernel {
	return CNNKernelClass.Alloc()
}

func (cc _CNNKernelClass) New() CNNKernel {
	rv := objc.Call[CNNKernel](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNKernel() CNNKernel {
	return CNNKernelClass.New()
}

func (c_ CNNKernel) Init() CNNKernel {
	rv := objc.Call[CNNKernel](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNKernel) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNKernel {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNKernel](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNKernel_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNKernel {
	instance := CNNKernelClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/3237263-batchencodingstoragesizeforsourc?language=objc
func (c_ CNNKernel) BatchEncodingStorageSizeForSourceImageSourceStatesDestinationImage(sourceImage *foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) uint {
	rv := objc.Call[uint](c_, objc.Sel("batchEncodingStorageSizeForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2947937-temporaryresultstateforcommandbu?language=objc
func (c_ CNNKernel) TemporaryResultStateForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer metal.PCommandBuffer, sourceImage IImage, sourceStates []IState, destinationImage IImage) State {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[State](c_, objc.Sel("temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:"), po0, objc.Ptr(sourceImage), sourceStates, objc.Ptr(destinationImage))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2947937-temporaryresultstateforcommandbu?language=objc
func (c_ CNNKernel) TemporaryResultStateForCommandBufferObjectSourceImageSourceStatesDestinationImage(commandBufferObject objc.IObject, sourceImage IImage, sourceStates []IState, destinationImage IImage) State {
	rv := objc.Call[State](c_, objc.Sel("temporaryResultStateForCommandBuffer:sourceImage:sourceStates:destinationImage:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceImage), sourceStates, objc.Ptr(destinationImage))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942671-appendbatchbarrier?language=objc
func (c_ CNNKernel) AppendBatchBarrier() bool {
	rv := objc.Call[bool](c_, objc.Sel("appendBatchBarrier"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2947939-temporaryresultstatebatchforcomm?language=objc
func (c_ CNNKernel) TemporaryResultStateBatchForCommandBufferSourceImageSourceStatesDestinationImage(commandBuffer metal.PCommandBuffer, sourceImage *foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) *foundation.Array {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[*foundation.Array](c_, objc.Sel("temporaryResultStateBatchForCommandBuffer:sourceImage:sourceStates:destinationImage:"), po0, sourceImage, sourceStates, destinationImage)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2947939-temporaryresultstatebatchforcomm?language=objc
func (c_ CNNKernel) TemporaryResultStateBatchForCommandBufferObjectSourceImageSourceStatesDestinationImage(commandBufferObject objc.IObject, sourceImage *foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) *foundation.Array {
	rv := objc.Call[*foundation.Array](c_, objc.Sel("temporaryResultStateBatchForCommandBuffer:sourceImage:sourceStates:destinationImage:"), objc.Ptr(commandBufferObject), sourceImage, sourceStates, destinationImage)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942661-destinationimagedescriptorforsou?language=objc
func (c_ CNNKernel) DestinationImageDescriptorForSourceImagesSourceStates(sourceImages []IImage, sourceStates []IState) ImageDescriptor {
	rv := objc.Call[ImageDescriptor](c_, objc.Sel("destinationImageDescriptorForSourceImages:sourceStates:"), sourceImages, sourceStates)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/3237264-encodingstoragesizeforsourceimag?language=objc
func (c_ CNNKernel) EncodingStorageSizeForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates []IState, destinationImage IImage) uint {
	rv := objc.Call[uint](c_, objc.Sel("encodingStorageSizeForSourceImage:sourceStates:destinationImage:"), objc.Ptr(sourceImage), sourceStates, objc.Ptr(destinationImage))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942651-encodebatchtocommandbuffer?language=objc
func (c_ CNNKernel) EncodeBatchToCommandBufferSourceImages(commandBuffer metal.PCommandBuffer, sourceImages *foundation.Array) *foundation.Array {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[*foundation.Array](c_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:"), po0, sourceImages)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942651-encodebatchtocommandbuffer?language=objc
func (c_ CNNKernel) EncodeBatchToCommandBufferObjectSourceImages(commandBufferObject objc.IObject, sourceImages *foundation.Array) *foundation.Array {
	rv := objc.Call[*foundation.Array](c_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:"), objc.Ptr(commandBufferObject), sourceImages)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2947931-resultstatebatchforsourceimage?language=objc
func (c_ CNNKernel) ResultStateBatchForSourceImageSourceStatesDestinationImage(sourceImage *foundation.Array, sourceStates []*foundation.Array, destinationImage *foundation.Array) *foundation.Array {
	rv := objc.Call[*foundation.Array](c_, objc.Sel("resultStateBatchForSourceImage:sourceStates:destinationImage:"), sourceImage, sourceStates, destinationImage)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865642-encodetocommandbuffer?language=objc
func (c_ CNNKernel) EncodeToCommandBufferSourceImage(commandBuffer metal.PCommandBuffer, sourceImage IImage) Image {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[Image](c_, objc.Sel("encodeToCommandBuffer:sourceImage:"), po0, objc.Ptr(sourceImage))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865642-encodetocommandbuffer?language=objc
func (c_ CNNKernel) EncodeToCommandBufferObjectSourceImage(commandBufferObject objc.IObject, sourceImage IImage) Image {
	rv := objc.Call[Image](c_, objc.Sel("encodeToCommandBuffer:sourceImage:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceImage))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942665-isresultstatereusedacrossbatch?language=objc
func (c_ CNNKernel) IsResultStateReusedAcrossBatch() bool {
	rv := objc.Call[bool](c_, objc.Sel("isResultStateReusedAcrossBatch"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2947932-resultstateforsourceimage?language=objc
func (c_ CNNKernel) ResultStateForSourceImageSourceStatesDestinationImage(sourceImage IImage, sourceStates []IState, destinationImage IImage) State {
	rv := objc.Call[State](c_, objc.Sel("resultStateForSourceImage:sourceStates:destinationImage:"), objc.Ptr(sourceImage), sourceStates, objc.Ptr(destinationImage))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2951917-sourcefeaturechannelmaxcount?language=objc
func (c_ CNNKernel) SourceFeatureChannelMaxCount() uint {
	rv := objc.Call[uint](c_, objc.Sel("sourceFeatureChannelMaxCount"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2951917-sourcefeaturechannelmaxcount?language=objc
func (c_ CNNKernel) SetSourceFeatureChannelMaxCount(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setSourceFeatureChannelMaxCount:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942673-isstatemodified?language=objc
func (c_ CNNKernel) IsStateModified() bool {
	rv := objc.Call[bool](c_, objc.Sel("isStateModified"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865657-padding?language=objc
func (c_ CNNKernel) Padding() NNPaddingWrapper {
	rv := objc.Call[NNPaddingWrapper](c_, objc.Sel("padding"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865657-padding?language=objc
func (c_ CNNKernel) SetPadding(value PNNPadding) {
	po0 := objc.WrapAsProtocol("MPSNNPadding", value)
	objc.Call[objc.Void](c_, objc.Sel("setPadding:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865657-padding?language=objc
func (c_ CNNKernel) SetPaddingObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setPadding:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942682-sourcefeaturechanneloffset?language=objc
func (c_ CNNKernel) SourceFeatureChannelOffset() uint {
	rv := objc.Call[uint](c_, objc.Sel("sourceFeatureChannelOffset"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942682-sourcefeaturechanneloffset?language=objc
func (c_ CNNKernel) SetSourceFeatureChannelOffset(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setSourceFeatureChannelOffset:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865644-strideinpixelsy?language=objc
func (c_ CNNKernel) StrideInPixelsY() uint {
	rv := objc.Call[uint](c_, objc.Sel("strideInPixelsY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865654-strideinpixelsx?language=objc
func (c_ CNNKernel) StrideInPixelsX() uint {
	rv := objc.Call[uint](c_, objc.Sel("strideInPixelsX"))
	return rv
}

// The edge mode to use when texture reads stray off the edge of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/1648826-edgemode?language=objc
func (c_ CNNKernel) EdgeMode() ImageEdgeMode {
	rv := objc.Call[ImageEdgeMode](c_, objc.Sel("edgeMode"))
	return rv
}

// The edge mode to use when texture reads stray off the edge of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/1648826-edgemode?language=objc
func (c_ CNNKernel) SetEdgeMode(value ImageEdgeMode) {
	objc.Call[objc.Void](c_, objc.Sel("setEdgeMode:"), value)
}

// The position of the destination image's clip rectangle origin, relative to the source image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/1648835-offset?language=objc
func (c_ CNNKernel) Offset() Offset {
	rv := objc.Call[Offset](c_, objc.Sel("offset"))
	return rv
}

// The position of the destination image's clip rectangle origin, relative to the source image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/1648835-offset?language=objc
func (c_ CNNKernel) SetOffset(value Offset) {
	objc.Call[objc.Void](c_, objc.Sel("setOffset:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865634-isbackwards?language=objc
func (c_ CNNKernel) IsBackwards() bool {
	rv := objc.Call[bool](c_, objc.Sel("isBackwards"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942679-dilationratey?language=objc
func (c_ CNNKernel) DilationRateY() uint {
	rv := objc.Call[uint](c_, objc.Sel("dilationRateY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865648-kernelheight?language=objc
func (c_ CNNKernel) KernelHeight() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelHeight"))
	return rv
}

// An optional clip rectangle to use when writing data. Only the pixels in the clip rectangle will be overwritten. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/1648911-cliprect?language=objc
func (c_ CNNKernel) ClipRect() metal.Region {
	rv := objc.Call[metal.Region](c_, objc.Sel("clipRect"))
	return rv
}

// An optional clip rectangle to use when writing data. Only the pixels in the clip rectangle will be overwritten. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/1648911-cliprect?language=objc
func (c_ CNNKernel) SetClipRect(value metal.Region) {
	objc.Call[objc.Void](c_, objc.Sel("setClipRect:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865637-kernelwidth?language=objc
func (c_ CNNKernel) KernelWidth() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelWidth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865650-destinationimageallocator?language=objc
func (c_ CNNKernel) DestinationImageAllocator() ImageAllocatorWrapper {
	rv := objc.Call[ImageAllocatorWrapper](c_, objc.Sel("destinationImageAllocator"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865650-destinationimageallocator?language=objc
func (c_ CNNKernel) SetDestinationImageAllocator(value PImageAllocator) {
	po0 := objc.WrapAsProtocol("MPSImageAllocator", value)
	objc.Call[objc.Void](c_, objc.Sel("setDestinationImageAllocator:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865650-destinationimageallocator?language=objc
func (c_ CNNKernel) SetDestinationImageAllocatorObject(valueObject objc.IObject) {
	objc.Call[objc.Void](c_, objc.Sel("setDestinationImageAllocator:"), objc.Ptr(valueObject))
}

// The number of channels in the destination image to skip before writing output data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2097550-destinationfeaturechanneloffset?language=objc
func (c_ CNNKernel) DestinationFeatureChannelOffset() uint {
	rv := objc.Call[uint](c_, objc.Sel("destinationFeatureChannelOffset"))
	return rv
}

// The number of channels in the destination image to skip before writing output data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2097550-destinationfeaturechanneloffset?language=objc
func (c_ CNNKernel) SetDestinationFeatureChannelOffset(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setDestinationFeatureChannelOffset:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2942669-dilationratex?language=objc
func (c_ CNNKernel) DilationRateX() uint {
	rv := objc.Call[uint](c_, objc.Sel("dilationRateX"))
	return rv
}
