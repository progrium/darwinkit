// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageStatisticsMinAndMax] class.
var ImageStatisticsMinAndMaxClass = _ImageStatisticsMinAndMaxClass{objc.GetClass("MPSImageStatisticsMinAndMax")}

type _ImageStatisticsMinAndMaxClass struct {
	objc.Class
}

// An interface definition for the [ImageStatisticsMinAndMax] class.
type IImageStatisticsMinAndMax interface {
	IUnaryImageKernel
	ClipRectSource() metal.Region
	SetClipRectSource(value metal.Region)
}

// A kernel that computes the minimum and maximum pixel values for a given region of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsminandmax?language=objc
type ImageStatisticsMinAndMax struct {
	UnaryImageKernel
}

func ImageStatisticsMinAndMaxFrom(ptr unsafe.Pointer) ImageStatisticsMinAndMax {
	return ImageStatisticsMinAndMax{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageStatisticsMinAndMax) InitWithDevice(device metal.PDevice) ImageStatisticsMinAndMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageStatisticsMinAndMax](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsminandmax/2867125-initwithdevice?language=objc
func NewImageStatisticsMinAndMaxWithDevice(device metal.PDevice) ImageStatisticsMinAndMax {
	instance := ImageStatisticsMinAndMaxClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageStatisticsMinAndMaxClass) Alloc() ImageStatisticsMinAndMax {
	rv := objc.Call[ImageStatisticsMinAndMax](ic, objc.Sel("alloc"))
	return rv
}

func ImageStatisticsMinAndMax_Alloc() ImageStatisticsMinAndMax {
	return ImageStatisticsMinAndMaxClass.Alloc()
}

func (ic _ImageStatisticsMinAndMaxClass) New() ImageStatisticsMinAndMax {
	rv := objc.Call[ImageStatisticsMinAndMax](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageStatisticsMinAndMax() ImageStatisticsMinAndMax {
	return ImageStatisticsMinAndMaxClass.New()
}

func (i_ ImageStatisticsMinAndMax) Init() ImageStatisticsMinAndMax {
	rv := objc.Call[ImageStatisticsMinAndMax](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageStatisticsMinAndMax) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageStatisticsMinAndMax {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageStatisticsMinAndMax](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageStatisticsMinAndMax_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageStatisticsMinAndMax {
	instance := ImageStatisticsMinAndMaxClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsminandmax/2867045-cliprectsource?language=objc
func (i_ ImageStatisticsMinAndMax) ClipRectSource() metal.Region {
	rv := objc.Call[metal.Region](i_, objc.Sel("clipRectSource"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsminandmax/2867045-cliprectsource?language=objc
func (i_ ImageStatisticsMinAndMax) SetClipRectSource(value metal.Region) {
	objc.Call[objc.Void](i_, objc.Sel("setClipRectSource:"), value)
}
