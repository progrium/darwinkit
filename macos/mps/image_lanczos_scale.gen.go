// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageLanczosScale] class.
var ImageLanczosScaleClass = _ImageLanczosScaleClass{objc.GetClass("MPSImageLanczosScale")}

type _ImageLanczosScaleClass struct {
	objc.Class
}

// An interface definition for the [ImageLanczosScale] class.
type IImageLanczosScale interface {
	IImageScale
}

// A filter that resizes and changes the aspect ratio of an image using Lanczos resampling. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelanczosscale?language=objc
type ImageLanczosScale struct {
	ImageScale
}

func ImageLanczosScaleFrom(ptr unsafe.Pointer) ImageLanczosScale {
	return ImageLanczosScale{
		ImageScale: ImageScaleFrom(ptr),
	}
}

func (i_ ImageLanczosScale) InitWithDevice(device metal.PDevice) ImageLanczosScale {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageLanczosScale](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelanczosscale/2867001-initwithdevice?language=objc
func NewImageLanczosScaleWithDevice(device metal.PDevice) ImageLanczosScale {
	instance := ImageLanczosScaleClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageLanczosScaleClass) Alloc() ImageLanczosScale {
	rv := objc.Call[ImageLanczosScale](ic, objc.Sel("alloc"))
	return rv
}

func ImageLanczosScale_Alloc() ImageLanczosScale {
	return ImageLanczosScaleClass.Alloc()
}

func (ic _ImageLanczosScaleClass) New() ImageLanczosScale {
	rv := objc.Call[ImageLanczosScale](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageLanczosScale() ImageLanczosScale {
	return ImageLanczosScaleClass.New()
}

func (i_ ImageLanczosScale) Init() ImageLanczosScale {
	rv := objc.Call[ImageLanczosScale](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageLanczosScale) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageLanczosScale {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageLanczosScale](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageLanczosScale_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageLanczosScale {
	instance := ImageLanczosScaleClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
