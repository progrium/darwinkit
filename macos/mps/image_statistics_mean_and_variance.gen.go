// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageStatisticsMeanAndVariance] class.
var ImageStatisticsMeanAndVarianceClass = _ImageStatisticsMeanAndVarianceClass{objc.GetClass("MPSImageStatisticsMeanAndVariance")}

type _ImageStatisticsMeanAndVarianceClass struct {
	objc.Class
}

// An interface definition for the [ImageStatisticsMeanAndVariance] class.
type IImageStatisticsMeanAndVariance interface {
	IUnaryImageKernel
	ClipRectSource() metal.Region
	SetClipRectSource(value metal.Region)
}

// A kernel that computes the mean and variance for a given region of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmeanandvariance?language=objc
type ImageStatisticsMeanAndVariance struct {
	UnaryImageKernel
}

func ImageStatisticsMeanAndVarianceFrom(ptr unsafe.Pointer) ImageStatisticsMeanAndVariance {
	return ImageStatisticsMeanAndVariance{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageStatisticsMeanAndVariance) InitWithDevice(device metal.PDevice) ImageStatisticsMeanAndVariance {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageStatisticsMeanAndVariance](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmeanandvariance/2867165-initwithdevice?language=objc
func NewImageStatisticsMeanAndVarianceWithDevice(device metal.PDevice) ImageStatisticsMeanAndVariance {
	instance := ImageStatisticsMeanAndVarianceClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageStatisticsMeanAndVarianceClass) Alloc() ImageStatisticsMeanAndVariance {
	rv := objc.Call[ImageStatisticsMeanAndVariance](ic, objc.Sel("alloc"))
	return rv
}

func ImageStatisticsMeanAndVariance_Alloc() ImageStatisticsMeanAndVariance {
	return ImageStatisticsMeanAndVarianceClass.Alloc()
}

func (ic _ImageStatisticsMeanAndVarianceClass) New() ImageStatisticsMeanAndVariance {
	rv := objc.Call[ImageStatisticsMeanAndVariance](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageStatisticsMeanAndVariance() ImageStatisticsMeanAndVariance {
	return ImageStatisticsMeanAndVarianceClass.New()
}

func (i_ ImageStatisticsMeanAndVariance) Init() ImageStatisticsMeanAndVariance {
	rv := objc.Call[ImageStatisticsMeanAndVariance](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageStatisticsMeanAndVariance) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageStatisticsMeanAndVariance {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageStatisticsMeanAndVariance](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageStatisticsMeanAndVariance_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageStatisticsMeanAndVariance {
	instance := ImageStatisticsMeanAndVarianceClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmeanandvariance/2867131-cliprectsource?language=objc
func (i_ ImageStatisticsMeanAndVariance) ClipRectSource() metal.Region {
	rv := objc.Call[metal.Region](i_, objc.Sel("clipRectSource"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmeanandvariance/2867131-cliprectsource?language=objc
func (i_ ImageStatisticsMeanAndVariance) SetClipRectSource(value metal.Region) {
	objc.Call[objc.Void](i_, objc.Sel("setClipRectSource:"), value)
}
