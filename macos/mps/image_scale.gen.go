// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageScale] class.
var ImageScaleClass = _ImageScaleClass{objc.GetClass("MPSImageScale")}

type _ImageScaleClass struct {
	objc.Class
}

// An interface definition for the [ImageScale] class.
type IImageScale interface {
	IUnaryImageKernel
	ScaleTransform() *ScaleTransform
	SetScaleTransform(value *ScaleTransform)
}

// A filter that resizes and changes the aspect ratio of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagescale?language=objc
type ImageScale struct {
	UnaryImageKernel
}

func ImageScaleFrom(ptr unsafe.Pointer) ImageScale {
	return ImageScale{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageScale) InitWithDevice(device metal.PDevice) ImageScale {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageScale](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagescale/2881186-initwithdevice?language=objc
func NewImageScaleWithDevice(device metal.PDevice) ImageScale {
	instance := ImageScaleClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageScaleClass) Alloc() ImageScale {
	rv := objc.Call[ImageScale](ic, objc.Sel("alloc"))
	return rv
}

func ImageScale_Alloc() ImageScale {
	return ImageScaleClass.Alloc()
}

func (ic _ImageScaleClass) New() ImageScale {
	rv := objc.Call[ImageScale](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageScale() ImageScale {
	return ImageScaleClass.New()
}

func (i_ ImageScale) Init() ImageScale {
	rv := objc.Call[ImageScale](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageScale) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageScale {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageScale](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageScale_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageScale {
	instance := ImageScaleClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagescale/2881183-scaletransform?language=objc
func (i_ ImageScale) ScaleTransform() *ScaleTransform {
	rv := objc.Call[*ScaleTransform](i_, objc.Sel("scaleTransform"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagescale/2881183-scaletransform?language=objc
func (i_ ImageScale) SetScaleTransform(value *ScaleTransform) {
	objc.Call[objc.Void](i_, objc.Sel("setScaleTransform:"), value)
}
