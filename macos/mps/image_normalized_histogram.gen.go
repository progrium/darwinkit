// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageNormalizedHistogram] class.
var ImageNormalizedHistogramClass = _ImageNormalizedHistogramClass{objc.GetClass("MPSImageNormalizedHistogram")}

type _ImageNormalizedHistogramClass struct {
	objc.Class
}

// An interface definition for the [ImageNormalizedHistogram] class.
type IImageNormalizedHistogram interface {
	IKernel
	HistogramSizeForSourceFormat(sourceFormat metal.PixelFormat) uint
	EncodeToCommandBufferSourceTextureMinmaxTextureHistogramHistogramOffset(commandBuffer metal.PCommandBuffer, source metal.PTexture, minmaxTexture metal.PTexture, histogram metal.PBuffer, histogramOffset uint)
	EncodeToCommandBufferObjectSourceTextureObjectMinmaxTextureObjectHistogramObjectHistogramOffset(commandBufferObject objc.IObject, sourceObject objc.IObject, minmaxTextureObject objc.IObject, histogramObject objc.IObject, histogramOffset uint)
	HistogramInfo() ImageHistogramInfo
	ZeroHistogram() bool
	SetZeroHistogram(value bool)
	ClipRectSource() metal.Region
	SetClipRectSource(value metal.Region)
}

// A filter that computes the normalized histogram of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram?language=objc
type ImageNormalizedHistogram struct {
	Kernel
}

func ImageNormalizedHistogramFrom(ptr unsafe.Pointer) ImageNormalizedHistogram {
	return ImageNormalizedHistogram{
		Kernel: KernelFrom(ptr),
	}
}

func (i_ ImageNormalizedHistogram) InitWithDeviceHistogramInfo(device metal.PDevice, histogramInfo *ImageHistogramInfo) ImageNormalizedHistogram {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageNormalizedHistogram](i_, objc.Sel("initWithDevice:histogramInfo:"), po0, histogramInfo)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019326-initwithdevice?language=objc
func NewImageNormalizedHistogramWithDeviceHistogramInfo(device metal.PDevice, histogramInfo *ImageHistogramInfo) ImageNormalizedHistogram {
	instance := ImageNormalizedHistogramClass.Alloc().InitWithDeviceHistogramInfo(device, histogramInfo)
	instance.Autorelease()
	return instance
}

func (ic _ImageNormalizedHistogramClass) Alloc() ImageNormalizedHistogram {
	rv := objc.Call[ImageNormalizedHistogram](ic, objc.Sel("alloc"))
	return rv
}

func ImageNormalizedHistogram_Alloc() ImageNormalizedHistogram {
	return ImageNormalizedHistogramClass.Alloc()
}

func (ic _ImageNormalizedHistogramClass) New() ImageNormalizedHistogram {
	rv := objc.Call[ImageNormalizedHistogram](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageNormalizedHistogram() ImageNormalizedHistogram {
	return ImageNormalizedHistogramClass.New()
}

func (i_ ImageNormalizedHistogram) Init() ImageNormalizedHistogram {
	rv := objc.Call[ImageNormalizedHistogram](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageNormalizedHistogram) InitWithDevice(device metal.PDevice) ImageNormalizedHistogram {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageNormalizedHistogram](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewImageNormalizedHistogramWithDevice(device metal.PDevice) ImageNormalizedHistogram {
	instance := ImageNormalizedHistogramClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageNormalizedHistogram) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageNormalizedHistogram {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageNormalizedHistogram](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageNormalizedHistogram_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageNormalizedHistogram {
	instance := ImageNormalizedHistogramClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019324-histogramsizeforsourceformat?language=objc
func (i_ ImageNormalizedHistogram) HistogramSizeForSourceFormat(sourceFormat metal.PixelFormat) uint {
	rv := objc.Call[uint](i_, objc.Sel("histogramSizeForSourceFormat:"), sourceFormat)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019322-encodetocommandbuffer?language=objc
func (i_ ImageNormalizedHistogram) EncodeToCommandBufferSourceTextureMinmaxTextureHistogramHistogramOffset(commandBuffer metal.PCommandBuffer, source metal.PTexture, minmaxTexture metal.PTexture, histogram metal.PBuffer, histogramOffset uint) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	po1 := objc.WrapAsProtocol("MTLTexture", source)
	po2 := objc.WrapAsProtocol("MTLTexture", minmaxTexture)
	po3 := objc.WrapAsProtocol("MTLBuffer", histogram)
	objc.Call[objc.Void](i_, objc.Sel("encodeToCommandBuffer:sourceTexture:minmaxTexture:histogram:histogramOffset:"), po0, po1, po2, po3, histogramOffset)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019322-encodetocommandbuffer?language=objc
func (i_ ImageNormalizedHistogram) EncodeToCommandBufferObjectSourceTextureObjectMinmaxTextureObjectHistogramObjectHistogramOffset(commandBufferObject objc.IObject, sourceObject objc.IObject, minmaxTextureObject objc.IObject, histogramObject objc.IObject, histogramOffset uint) {
	objc.Call[objc.Void](i_, objc.Sel("encodeToCommandBuffer:sourceTexture:minmaxTexture:histogram:histogramOffset:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceObject), objc.Ptr(minmaxTextureObject), objc.Ptr(histogramObject), histogramOffset)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019323-histograminfo?language=objc
func (i_ ImageNormalizedHistogram) HistogramInfo() ImageHistogramInfo {
	rv := objc.Call[ImageHistogramInfo](i_, objc.Sel("histogramInfo"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019327-zerohistogram?language=objc
func (i_ ImageNormalizedHistogram) ZeroHistogram() bool {
	rv := objc.Call[bool](i_, objc.Sel("zeroHistogram"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019327-zerohistogram?language=objc
func (i_ ImageNormalizedHistogram) SetZeroHistogram(value bool) {
	objc.Call[objc.Void](i_, objc.Sel("setZeroHistogram:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019321-cliprectsource?language=objc
func (i_ ImageNormalizedHistogram) ClipRectSource() metal.Region {
	rv := objc.Call[metal.Region](i_, objc.Sel("clipRectSource"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagenormalizedhistogram/3019321-cliprectsource?language=objc
func (i_ ImageNormalizedHistogram) SetClipRectSource(value metal.Region) {
	objc.Call[objc.Void](i_, objc.Sel("setClipRectSource:"), value)
}
