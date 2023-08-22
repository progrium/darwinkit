// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageHistogramSpecification] class.
var ImageHistogramSpecificationClass = _ImageHistogramSpecificationClass{objc.GetClass("MPSImageHistogramSpecification")}

type _ImageHistogramSpecificationClass struct {
	objc.Class
}

// An interface definition for the [ImageHistogramSpecification] class.
type IImageHistogramSpecification interface {
	IUnaryImageKernel
	EncodeTransformToCommandBufferSourceTextureSourceHistogramSourceHistogramOffsetDesiredHistogramDesiredHistogramOffset(commandBuffer metal.PCommandBuffer, source metal.PTexture, sourceHistogram metal.PBuffer, sourceHistogramOffset uint, desiredHistogram metal.PBuffer, desiredHistogramOffset uint)
	EncodeTransformToCommandBufferObjectSourceTextureObjectSourceHistogramObjectSourceHistogramOffsetDesiredHistogramObjectDesiredHistogramOffset(commandBufferObject objc.IObject, sourceObject objc.IObject, sourceHistogramObject objc.IObject, sourceHistogramOffset uint, desiredHistogramObject objc.IObject, desiredHistogramOffset uint)
	HistogramInfo() ImageHistogramInfo
}

// A filter that performs a histogram specification operation on an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramspecification?language=objc
type ImageHistogramSpecification struct {
	UnaryImageKernel
}

func ImageHistogramSpecificationFrom(ptr unsafe.Pointer) ImageHistogramSpecification {
	return ImageHistogramSpecification{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageHistogramSpecification) InitWithDeviceHistogramInfo(device metal.PDevice, histogramInfo *ImageHistogramInfo) ImageHistogramSpecification {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageHistogramSpecification](i_, objc.Sel("initWithDevice:histogramInfo:"), po0, histogramInfo)
	return rv
}

// Initializes a histogram with specific information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramspecification/1618907-initwithdevice?language=objc
func NewImageHistogramSpecificationWithDeviceHistogramInfo(device metal.PDevice, histogramInfo *ImageHistogramInfo) ImageHistogramSpecification {
	instance := ImageHistogramSpecificationClass.Alloc().InitWithDeviceHistogramInfo(device, histogramInfo)
	instance.Autorelease()
	return instance
}

func (ic _ImageHistogramSpecificationClass) Alloc() ImageHistogramSpecification {
	rv := objc.Call[ImageHistogramSpecification](ic, objc.Sel("alloc"))
	return rv
}

func ImageHistogramSpecification_Alloc() ImageHistogramSpecification {
	return ImageHistogramSpecificationClass.Alloc()
}

func (ic _ImageHistogramSpecificationClass) New() ImageHistogramSpecification {
	rv := objc.Call[ImageHistogramSpecification](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageHistogramSpecification() ImageHistogramSpecification {
	return ImageHistogramSpecificationClass.New()
}

func (i_ ImageHistogramSpecification) Init() ImageHistogramSpecification {
	rv := objc.Call[ImageHistogramSpecification](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageHistogramSpecification) InitWithDevice(device metal.PDevice) ImageHistogramSpecification {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageHistogramSpecification](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageHistogramSpecificationWithDevice(device metal.PDevice) ImageHistogramSpecification {
	instance := ImageHistogramSpecificationClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageHistogramSpecification) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageHistogramSpecification {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageHistogramSpecification](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageHistogramSpecification_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageHistogramSpecification {
	instance := ImageHistogramSpecificationClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// Encodes the transform function to a command buffer using a compute command encoder. The transform function computes the equalization lookup table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramspecification/1618854-encodetransformtocommandbuffer?language=objc
func (i_ ImageHistogramSpecification) EncodeTransformToCommandBufferSourceTextureSourceHistogramSourceHistogramOffsetDesiredHistogramDesiredHistogramOffset(commandBuffer metal.PCommandBuffer, source metal.PTexture, sourceHistogram metal.PBuffer, sourceHistogramOffset uint, desiredHistogram metal.PBuffer, desiredHistogramOffset uint) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	po1 := objc.WrapAsProtocol("MTLTexture", source)
	po2 := objc.WrapAsProtocol("MTLBuffer", sourceHistogram)
	po4 := objc.WrapAsProtocol("MTLBuffer", desiredHistogram)
	objc.Call[objc.Void](i_, objc.Sel("encodeTransformToCommandBuffer:sourceTexture:sourceHistogram:sourceHistogramOffset:desiredHistogram:desiredHistogramOffset:"), po0, po1, po2, sourceHistogramOffset, po4, desiredHistogramOffset)
}

// Encodes the transform function to a command buffer using a compute command encoder. The transform function computes the equalization lookup table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramspecification/1618854-encodetransformtocommandbuffer?language=objc
func (i_ ImageHistogramSpecification) EncodeTransformToCommandBufferObjectSourceTextureObjectSourceHistogramObjectSourceHistogramOffsetDesiredHistogramObjectDesiredHistogramOffset(commandBufferObject objc.IObject, sourceObject objc.IObject, sourceHistogramObject objc.IObject, sourceHistogramOffset uint, desiredHistogramObject objc.IObject, desiredHistogramOffset uint) {
	objc.Call[objc.Void](i_, objc.Sel("encodeTransformToCommandBuffer:sourceTexture:sourceHistogram:sourceHistogramOffset:desiredHistogram:desiredHistogramOffset:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceObject), objc.Ptr(sourceHistogramObject), sourceHistogramOffset, objc.Ptr(desiredHistogramObject), desiredHistogramOffset)
}

// A structure describing the histogram content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramspecification/1618810-histograminfo?language=objc
func (i_ ImageHistogramSpecification) HistogramInfo() ImageHistogramInfo {
	rv := objc.Call[ImageHistogramInfo](i_, objc.Sel("histogramInfo"))
	return rv
}
