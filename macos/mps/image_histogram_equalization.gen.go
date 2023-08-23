// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageHistogramEqualization] class.
var ImageHistogramEqualizationClass = _ImageHistogramEqualizationClass{objc.GetClass("MPSImageHistogramEqualization")}

type _ImageHistogramEqualizationClass struct {
	objc.Class
}

// An interface definition for the [ImageHistogramEqualization] class.
type IImageHistogramEqualization interface {
	IUnaryImageKernel
	EncodeTransformToCommandBufferSourceTextureHistogramHistogramOffset(commandBuffer metal.PCommandBuffer, source metal.PTexture, histogram metal.PBuffer, histogramOffset uint)
	EncodeTransformToCommandBufferObjectSourceTextureObjectHistogramObjectHistogramOffset(commandBufferObject objc.IObject, sourceObject objc.IObject, histogramObject objc.IObject, histogramOffset uint)
	HistogramInfo() ImageHistogramInfo
}

// A filter that equalizes the histogram of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramequalization?language=objc
type ImageHistogramEqualization struct {
	UnaryImageKernel
}

func ImageHistogramEqualizationFrom(ptr unsafe.Pointer) ImageHistogramEqualization {
	return ImageHistogramEqualization{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageHistogramEqualization) InitWithDeviceHistogramInfo(device metal.PDevice, histogramInfo *ImageHistogramInfo) ImageHistogramEqualization {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageHistogramEqualization](i_, objc.Sel("initWithDevice:histogramInfo:"), po0, histogramInfo)
	return rv
}

// Initializes a histogram with specific information. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramequalization/1618856-initwithdevice?language=objc
func NewImageHistogramEqualizationWithDeviceHistogramInfo(device metal.PDevice, histogramInfo *ImageHistogramInfo) ImageHistogramEqualization {
	instance := ImageHistogramEqualizationClass.Alloc().InitWithDeviceHistogramInfo(device, histogramInfo)
	instance.Autorelease()
	return instance
}

func (ic _ImageHistogramEqualizationClass) Alloc() ImageHistogramEqualization {
	rv := objc.Call[ImageHistogramEqualization](ic, objc.Sel("alloc"))
	return rv
}

func ImageHistogramEqualization_Alloc() ImageHistogramEqualization {
	return ImageHistogramEqualizationClass.Alloc()
}

func (ic _ImageHistogramEqualizationClass) New() ImageHistogramEqualization {
	rv := objc.Call[ImageHistogramEqualization](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageHistogramEqualization() ImageHistogramEqualization {
	return ImageHistogramEqualizationClass.New()
}

func (i_ ImageHistogramEqualization) Init() ImageHistogramEqualization {
	rv := objc.Call[ImageHistogramEqualization](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageHistogramEqualization) InitWithDevice(device metal.PDevice) ImageHistogramEqualization {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageHistogramEqualization](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageHistogramEqualizationWithDevice(device metal.PDevice) ImageHistogramEqualization {
	instance := ImageHistogramEqualizationClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageHistogramEqualization) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageHistogramEqualization {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageHistogramEqualization](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageHistogramEqualization_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageHistogramEqualization {
	instance := ImageHistogramEqualizationClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// Encodes the transform function to a command buffer using a compute command encoder. The transform function computes the equalization lookup table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramequalization/1618746-encodetransformtocommandbuffer?language=objc
func (i_ ImageHistogramEqualization) EncodeTransformToCommandBufferSourceTextureHistogramHistogramOffset(commandBuffer metal.PCommandBuffer, source metal.PTexture, histogram metal.PBuffer, histogramOffset uint) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	po1 := objc.WrapAsProtocol("MTLTexture", source)
	po2 := objc.WrapAsProtocol("MTLBuffer", histogram)
	objc.Call[objc.Void](i_, objc.Sel("encodeTransformToCommandBuffer:sourceTexture:histogram:histogramOffset:"), po0, po1, po2, histogramOffset)
}

// Encodes the transform function to a command buffer using a compute command encoder. The transform function computes the equalization lookup table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramequalization/1618746-encodetransformtocommandbuffer?language=objc
func (i_ ImageHistogramEqualization) EncodeTransformToCommandBufferObjectSourceTextureObjectHistogramObjectHistogramOffset(commandBufferObject objc.IObject, sourceObject objc.IObject, histogramObject objc.IObject, histogramOffset uint) {
	objc.Call[objc.Void](i_, objc.Sel("encodeTransformToCommandBuffer:sourceTexture:histogram:histogramOffset:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceObject), objc.Ptr(histogramObject), histogramOffset)
}

// A structure describing the histogram content. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagehistogramequalization/1618775-histograminfo?language=objc
func (i_ ImageHistogramEqualization) HistogramInfo() ImageHistogramInfo {
	rv := objc.Call[ImageHistogramInfo](i_, objc.Sel("histogramInfo"))
	return rv
}
