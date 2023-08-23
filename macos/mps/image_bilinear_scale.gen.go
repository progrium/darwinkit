// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageBilinearScale] class.
var ImageBilinearScaleClass = _ImageBilinearScaleClass{objc.GetClass("MPSImageBilinearScale")}

type _ImageBilinearScaleClass struct {
	objc.Class
}

// An interface definition for the [ImageBilinearScale] class.
type IImageBilinearScale interface {
	IImageScale
}

// A filter that resizes and changes the aspect ratio of an image using Bilinear resampling. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagebilinearscale?language=objc
type ImageBilinearScale struct {
	ImageScale
}

func ImageBilinearScaleFrom(ptr unsafe.Pointer) ImageBilinearScale {
	return ImageBilinearScale{
		ImageScale: ImageScaleFrom(ptr),
	}
}

func (i_ ImageBilinearScale) InitWithDevice(device metal.PDevice) ImageBilinearScale {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageBilinearScale](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagebilinearscale/2881184-initwithdevice?language=objc
func NewImageBilinearScaleWithDevice(device metal.PDevice) ImageBilinearScale {
	instance := ImageBilinearScaleClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageBilinearScaleClass) Alloc() ImageBilinearScale {
	rv := objc.Call[ImageBilinearScale](ic, objc.Sel("alloc"))
	return rv
}

func ImageBilinearScale_Alloc() ImageBilinearScale {
	return ImageBilinearScaleClass.Alloc()
}

func (ic _ImageBilinearScaleClass) New() ImageBilinearScale {
	rv := objc.Call[ImageBilinearScale](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageBilinearScale() ImageBilinearScale {
	return ImageBilinearScaleClass.New()
}

func (i_ ImageBilinearScale) Init() ImageBilinearScale {
	rv := objc.Call[ImageBilinearScale](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageBilinearScale) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageBilinearScale {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageBilinearScale](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageBilinearScale_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageBilinearScale {
	instance := ImageBilinearScaleClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
