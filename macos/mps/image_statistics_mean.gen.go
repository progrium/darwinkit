// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageStatisticsMean] class.
var ImageStatisticsMeanClass = _ImageStatisticsMeanClass{objc.GetClass("MPSImageStatisticsMean")}

type _ImageStatisticsMeanClass struct {
	objc.Class
}

// An interface definition for the [ImageStatisticsMean] class.
type IImageStatisticsMean interface {
	IUnaryImageKernel
	ClipRectSource() metal.Region
	SetClipRectSource(value metal.Region)
}

// A kernel that computes the mean for a given region of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmean?language=objc
type ImageStatisticsMean struct {
	UnaryImageKernel
}

func ImageStatisticsMeanFrom(ptr unsafe.Pointer) ImageStatisticsMean {
	return ImageStatisticsMean{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageStatisticsMean) InitWithDevice(device metal.PDevice) ImageStatisticsMean {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageStatisticsMean](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmean/2867156-initwithdevice?language=objc
func NewImageStatisticsMeanWithDevice(device metal.PDevice) ImageStatisticsMean {
	instance := ImageStatisticsMeanClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageStatisticsMeanClass) Alloc() ImageStatisticsMean {
	rv := objc.Call[ImageStatisticsMean](ic, objc.Sel("alloc"))
	return rv
}

func ImageStatisticsMean_Alloc() ImageStatisticsMean {
	return ImageStatisticsMeanClass.Alloc()
}

func (ic _ImageStatisticsMeanClass) New() ImageStatisticsMean {
	rv := objc.Call[ImageStatisticsMean](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageStatisticsMean() ImageStatisticsMean {
	return ImageStatisticsMeanClass.New()
}

func (i_ ImageStatisticsMean) Init() ImageStatisticsMean {
	rv := objc.Call[ImageStatisticsMean](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageStatisticsMean) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageStatisticsMean {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageStatisticsMean](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageStatisticsMean_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageStatisticsMean {
	instance := ImageStatisticsMeanClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmean/2867093-cliprectsource?language=objc
func (i_ ImageStatisticsMean) ClipRectSource() metal.Region {
	rv := objc.Call[metal.Region](i_, objc.Sel("clipRectSource"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagestatisticsmean/2867093-cliprectsource?language=objc
func (i_ ImageStatisticsMean) SetClipRectSource(value metal.Region) {
	objc.Call[objc.Void](i_, objc.Sel("setClipRectSource:"), value)
}
