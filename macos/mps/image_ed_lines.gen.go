// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageEDLines] class.
var ImageEDLinesClass = _ImageEDLinesClass{objc.GetClass("MPSImageEDLines")}

type _ImageEDLinesClass struct {
	objc.Class
}

// An interface definition for the [ImageEDLines] class.
type IImageEDLines interface {
	IKernel
	EncodeToCommandBufferSourceTextureDestinationTextureEndpointBufferEndpointOffset(commandBuffer metal.PCommandBuffer, source metal.PTexture, dest metal.PTexture, endpointBuffer metal.PBuffer, endpointOffset uint)
	EncodeToCommandBufferObjectSourceTextureObjectDestinationTextureObjectEndpointBufferObjectEndpointOffset(commandBufferObject objc.IObject, sourceObject objc.IObject, destObject objc.IObject, endpointBufferObject objc.IObject, endpointOffset uint)
	MaxLines() uint
	SetMaxLines(value uint)
	MinLineLength() int
	SetMinLineLength(value int)
	LineErrorThreshold() float64
	SetLineErrorThreshold(value float64)
	GaussianSigma() float64
	DetailRatio() int
	SetDetailRatio(value int)
	GradientThreshold() float64
	SetGradientThreshold(value float64)
	ClipRectSource() metal.Region
	SetClipRectSource(value metal.Region)
	MergeLocalityThreshold() float64
	SetMergeLocalityThreshold(value float64)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines?language=objc
type ImageEDLines struct {
	Kernel
}

func ImageEDLinesFrom(ptr unsafe.Pointer) ImageEDLines {
	return ImageEDLines{
		Kernel: KernelFrom(ptr),
	}
}

func (i_ ImageEDLines) InitWithDeviceGaussianSigmaMinLineLengthMaxLinesDetailRatioGradientThresholdLineErrorThresholdMergeLocalityThreshold(device metal.PDevice, gaussianSigma float64, minLineLength int, maxLines uint, detailRatio int, gradientThreshold float64, lineErrorThreshold float64, mergeLocalityThreshold float64) ImageEDLines {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageEDLines](i_, objc.Sel("initWithDevice:gaussianSigma:minLineLength:maxLines:detailRatio:gradientThreshold:lineErrorThreshold:mergeLocalityThreshold:"), po0, gaussianSigma, minLineLength, maxLines, detailRatio, gradientThreshold, lineErrorThreshold, mergeLocalityThreshold)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618921-initwithdevice?language=objc
func NewImageEDLinesWithDeviceGaussianSigmaMinLineLengthMaxLinesDetailRatioGradientThresholdLineErrorThresholdMergeLocalityThreshold(device metal.PDevice, gaussianSigma float64, minLineLength int, maxLines uint, detailRatio int, gradientThreshold float64, lineErrorThreshold float64, mergeLocalityThreshold float64) ImageEDLines {
	instance := ImageEDLinesClass.Alloc().InitWithDeviceGaussianSigmaMinLineLengthMaxLinesDetailRatioGradientThresholdLineErrorThresholdMergeLocalityThreshold(device, gaussianSigma, minLineLength, maxLines, detailRatio, gradientThreshold, lineErrorThreshold, mergeLocalityThreshold)
	instance.Autorelease()
	return instance
}

func (ic _ImageEDLinesClass) Alloc() ImageEDLines {
	rv := objc.Call[ImageEDLines](ic, objc.Sel("alloc"))
	return rv
}

func ImageEDLines_Alloc() ImageEDLines {
	return ImageEDLinesClass.Alloc()
}

func (ic _ImageEDLinesClass) New() ImageEDLines {
	rv := objc.Call[ImageEDLines](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageEDLines() ImageEDLines {
	return ImageEDLinesClass.New()
}

func (i_ ImageEDLines) Init() ImageEDLines {
	rv := objc.Call[ImageEDLines](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageEDLines) InitWithDevice(device metal.PDevice) ImageEDLines {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageEDLines](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewImageEDLinesWithDevice(device metal.PDevice) ImageEDLines {
	instance := ImageEDLinesClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageEDLines) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageEDLines {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageEDLines](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageEDLines_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageEDLines {
	instance := ImageEDLinesClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618917-encodetocommandbuffer?language=objc
func (i_ ImageEDLines) EncodeToCommandBufferSourceTextureDestinationTextureEndpointBufferEndpointOffset(commandBuffer metal.PCommandBuffer, source metal.PTexture, dest metal.PTexture, endpointBuffer metal.PBuffer, endpointOffset uint) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	po1 := objc.WrapAsProtocol("MTLTexture", source)
	po2 := objc.WrapAsProtocol("MTLTexture", dest)
	po3 := objc.WrapAsProtocol("MTLBuffer", endpointBuffer)
	objc.Call[objc.Void](i_, objc.Sel("encodeToCommandBuffer:sourceTexture:destinationTexture:endpointBuffer:endpointOffset:"), po0, po1, po2, po3, endpointOffset)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618917-encodetocommandbuffer?language=objc
func (i_ ImageEDLines) EncodeToCommandBufferObjectSourceTextureObjectDestinationTextureObjectEndpointBufferObjectEndpointOffset(commandBufferObject objc.IObject, sourceObject objc.IObject, destObject objc.IObject, endpointBufferObject objc.IObject, endpointOffset uint) {
	objc.Call[objc.Void](i_, objc.Sel("encodeToCommandBuffer:sourceTexture:destinationTexture:endpointBuffer:endpointOffset:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceObject), objc.Ptr(destObject), objc.Ptr(endpointBufferObject), endpointOffset)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618923-maxlines?language=objc
func (i_ ImageEDLines) MaxLines() uint {
	rv := objc.Call[uint](i_, objc.Sel("maxLines"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618923-maxlines?language=objc
func (i_ ImageEDLines) SetMaxLines(value uint) {
	objc.Call[objc.Void](i_, objc.Sel("setMaxLines:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618925-minlinelength?language=objc
func (i_ ImageEDLines) MinLineLength() int {
	rv := objc.Call[int](i_, objc.Sel("minLineLength"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618925-minlinelength?language=objc
func (i_ ImageEDLines) SetMinLineLength(value int) {
	objc.Call[objc.Void](i_, objc.Sel("setMinLineLength:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618922-lineerrorthreshold?language=objc
func (i_ ImageEDLines) LineErrorThreshold() float64 {
	rv := objc.Call[float64](i_, objc.Sel("lineErrorThreshold"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618922-lineerrorthreshold?language=objc
func (i_ ImageEDLines) SetLineErrorThreshold(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setLineErrorThreshold:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618918-gaussiansigma?language=objc
func (i_ ImageEDLines) GaussianSigma() float64 {
	rv := objc.Call[float64](i_, objc.Sel("gaussianSigma"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618916-detailratio?language=objc
func (i_ ImageEDLines) DetailRatio() int {
	rv := objc.Call[int](i_, objc.Sel("detailRatio"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618916-detailratio?language=objc
func (i_ ImageEDLines) SetDetailRatio(value int) {
	objc.Call[objc.Void](i_, objc.Sel("setDetailRatio:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618919-gradientthreshold?language=objc
func (i_ ImageEDLines) GradientThreshold() float64 {
	rv := objc.Call[float64](i_, objc.Sel("gradientThreshold"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618919-gradientthreshold?language=objc
func (i_ ImageEDLines) SetGradientThreshold(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setGradientThreshold:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618915-cliprectsource?language=objc
func (i_ ImageEDLines) ClipRectSource() metal.Region {
	rv := objc.Call[metal.Region](i_, objc.Sel("clipRectSource"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618915-cliprectsource?language=objc
func (i_ ImageEDLines) SetClipRectSource(value metal.Region) {
	objc.Call[objc.Void](i_, objc.Sel("setClipRectSource:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618924-mergelocalitythreshold?language=objc
func (i_ ImageEDLines) MergeLocalityThreshold() float64 {
	rv := objc.Call[float64](i_, objc.Sel("mergeLocalityThreshold"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageedlines/3618924-mergelocalitythreshold?language=objc
func (i_ ImageEDLines) SetMergeLocalityThreshold(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setMergeLocalityThreshold:"), value)
}
