// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageGuidedFilter] class.
var ImageGuidedFilterClass = _ImageGuidedFilterClass{objc.GetClass("MPSImageGuidedFilter")}

type _ImageGuidedFilterClass struct {
	objc.Class
}

// An interface definition for the [ImageGuidedFilter] class.
type IImageGuidedFilter interface {
	IKernel
	EncodeReconstructionToCommandBufferGuidanceTextureCoefficientsTextureDestinationTexture(commandBuffer metal.PCommandBuffer, guidanceTexture metal.PTexture, coefficientsTexture metal.PTexture, destinationTexture metal.PTexture)
	EncodeReconstructionToCommandBufferObjectGuidanceTextureObjectCoefficientsTextureObjectDestinationTextureObject(commandBufferObject objc.IObject, guidanceTextureObject objc.IObject, coefficientsTextureObject objc.IObject, destinationTextureObject objc.IObject)
	EncodeRegressionToCommandBufferSourceTextureGuidanceTextureWeightsTextureDestinationCoefficientsTexture(commandBuffer metal.PCommandBuffer, sourceTexture metal.PTexture, guidanceTexture metal.PTexture, weightsTexture metal.PTexture, destinationCoefficientsTexture metal.PTexture)
	EncodeRegressionToCommandBufferObjectSourceTextureObjectGuidanceTextureObjectWeightsTextureObjectDestinationCoefficientsTextureObject(commandBufferObject objc.IObject, sourceTextureObject objc.IObject, guidanceTextureObject objc.IObject, weightsTextureObject objc.IObject, destinationCoefficientsTextureObject objc.IObject)
	ReconstructOffset() float64
	SetReconstructOffset(value float64)
	Epsilon() float64
	SetEpsilon(value float64)
	ReconstructScale() float64
	SetReconstructScale(value float64)
	KernelDiameter() uint
}

// A filter that performs edge-aware filtering on an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter?language=objc
type ImageGuidedFilter struct {
	Kernel
}

func ImageGuidedFilterFrom(ptr unsafe.Pointer) ImageGuidedFilter {
	return ImageGuidedFilter{
		Kernel: KernelFrom(ptr),
	}
}

func (i_ ImageGuidedFilter) InitWithDeviceKernelDiameter(device metal.PDevice, kernelDiameter uint) ImageGuidedFilter {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageGuidedFilter](i_, objc.Sel("initWithDevice:kernelDiameter:"), po0, kernelDiameter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2951910-initwithdevice?language=objc
func NewImageGuidedFilterWithDeviceKernelDiameter(device metal.PDevice, kernelDiameter uint) ImageGuidedFilter {
	instance := ImageGuidedFilterClass.Alloc().InitWithDeviceKernelDiameter(device, kernelDiameter)
	instance.Autorelease()
	return instance
}

func (ic _ImageGuidedFilterClass) Alloc() ImageGuidedFilter {
	rv := objc.Call[ImageGuidedFilter](ic, objc.Sel("alloc"))
	return rv
}

func ImageGuidedFilter_Alloc() ImageGuidedFilter {
	return ImageGuidedFilterClass.Alloc()
}

func (ic _ImageGuidedFilterClass) New() ImageGuidedFilter {
	rv := objc.Call[ImageGuidedFilter](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageGuidedFilter() ImageGuidedFilter {
	return ImageGuidedFilterClass.New()
}

func (i_ ImageGuidedFilter) Init() ImageGuidedFilter {
	rv := objc.Call[ImageGuidedFilter](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageGuidedFilter) InitWithDevice(device metal.PDevice) ImageGuidedFilter {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageGuidedFilter](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewImageGuidedFilterWithDevice(device metal.PDevice) ImageGuidedFilter {
	instance := ImageGuidedFilterClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageGuidedFilter) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageGuidedFilter {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageGuidedFilter](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageGuidedFilter_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageGuidedFilter {
	instance := ImageGuidedFilterClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2951906-encodereconstructiontocommandbuf?language=objc
func (i_ ImageGuidedFilter) EncodeReconstructionToCommandBufferGuidanceTextureCoefficientsTextureDestinationTexture(commandBuffer metal.PCommandBuffer, guidanceTexture metal.PTexture, coefficientsTexture metal.PTexture, destinationTexture metal.PTexture) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	po1 := objc.WrapAsProtocol("MTLTexture", guidanceTexture)
	po2 := objc.WrapAsProtocol("MTLTexture", coefficientsTexture)
	po3 := objc.WrapAsProtocol("MTLTexture", destinationTexture)
	objc.Call[objc.Void](i_, objc.Sel("encodeReconstructionToCommandBuffer:guidanceTexture:coefficientsTexture:destinationTexture:"), po0, po1, po2, po3)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2951906-encodereconstructiontocommandbuf?language=objc
func (i_ ImageGuidedFilter) EncodeReconstructionToCommandBufferObjectGuidanceTextureObjectCoefficientsTextureObjectDestinationTextureObject(commandBufferObject objc.IObject, guidanceTextureObject objc.IObject, coefficientsTextureObject objc.IObject, destinationTextureObject objc.IObject) {
	objc.Call[objc.Void](i_, objc.Sel("encodeReconstructionToCommandBuffer:guidanceTexture:coefficientsTexture:destinationTexture:"), objc.Ptr(commandBufferObject), objc.Ptr(guidanceTextureObject), objc.Ptr(coefficientsTextureObject), objc.Ptr(destinationTextureObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2951907-encoderegressiontocommandbuffer?language=objc
func (i_ ImageGuidedFilter) EncodeRegressionToCommandBufferSourceTextureGuidanceTextureWeightsTextureDestinationCoefficientsTexture(commandBuffer metal.PCommandBuffer, sourceTexture metal.PTexture, guidanceTexture metal.PTexture, weightsTexture metal.PTexture, destinationCoefficientsTexture metal.PTexture) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	po1 := objc.WrapAsProtocol("MTLTexture", sourceTexture)
	po2 := objc.WrapAsProtocol("MTLTexture", guidanceTexture)
	po3 := objc.WrapAsProtocol("MTLTexture", weightsTexture)
	po4 := objc.WrapAsProtocol("MTLTexture", destinationCoefficientsTexture)
	objc.Call[objc.Void](i_, objc.Sel("encodeRegressionToCommandBuffer:sourceTexture:guidanceTexture:weightsTexture:destinationCoefficientsTexture:"), po0, po1, po2, po3, po4)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2951907-encoderegressiontocommandbuffer?language=objc
func (i_ ImageGuidedFilter) EncodeRegressionToCommandBufferObjectSourceTextureObjectGuidanceTextureObjectWeightsTextureObjectDestinationCoefficientsTextureObject(commandBufferObject objc.IObject, sourceTextureObject objc.IObject, guidanceTextureObject objc.IObject, weightsTextureObject objc.IObject, destinationCoefficientsTextureObject objc.IObject) {
	objc.Call[objc.Void](i_, objc.Sel("encodeRegressionToCommandBuffer:sourceTexture:guidanceTexture:weightsTexture:destinationCoefficientsTexture:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceTextureObject), objc.Ptr(guidanceTextureObject), objc.Ptr(weightsTextureObject), objc.Ptr(destinationCoefficientsTextureObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2953079-reconstructoffset?language=objc
func (i_ ImageGuidedFilter) ReconstructOffset() float64 {
	rv := objc.Call[float64](i_, objc.Sel("reconstructOffset"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2953079-reconstructoffset?language=objc
func (i_ ImageGuidedFilter) SetReconstructOffset(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setReconstructOffset:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2951908-epsilon?language=objc
func (i_ ImageGuidedFilter) Epsilon() float64 {
	rv := objc.Call[float64](i_, objc.Sel("epsilon"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2951908-epsilon?language=objc
func (i_ ImageGuidedFilter) SetEpsilon(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setEpsilon:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2953078-reconstructscale?language=objc
func (i_ ImageGuidedFilter) ReconstructScale() float64 {
	rv := objc.Call[float64](i_, objc.Sel("reconstructScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2953078-reconstructscale?language=objc
func (i_ ImageGuidedFilter) SetReconstructScale(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setReconstructScale:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageguidedfilter/2951909-kerneldiameter?language=objc
func (i_ ImageGuidedFilter) KernelDiameter() uint {
	rv := objc.Call[uint](i_, objc.Sel("kernelDiameter"))
	return rv
}
