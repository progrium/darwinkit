// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageSubtract] class.
var ImageSubtractClass = _ImageSubtractClass{objc.GetClass("MPSImageSubtract")}

type _ImageSubtractClass struct {
	objc.Class
}

// An interface definition for the [ImageSubtract] class.
type IImageSubtract interface {
	IImageArithmetic
}

// A filter that returns the element-wise difference of its two input images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesubtract?language=objc
type ImageSubtract struct {
	ImageArithmetic
}

func ImageSubtractFrom(ptr unsafe.Pointer) ImageSubtract {
	return ImageSubtract{
		ImageArithmetic: ImageArithmeticFrom(ptr),
	}
}

func (i_ ImageSubtract) InitWithDevice(device metal.PDevice) ImageSubtract {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageSubtract](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesubtract/2866613-initwithdevice?language=objc
func NewImageSubtractWithDevice(device metal.PDevice) ImageSubtract {
	instance := ImageSubtractClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageSubtractClass) Alloc() ImageSubtract {
	rv := objc.Call[ImageSubtract](ic, objc.Sel("alloc"))
	return rv
}

func ImageSubtract_Alloc() ImageSubtract {
	return ImageSubtractClass.Alloc()
}

func (ic _ImageSubtractClass) New() ImageSubtract {
	rv := objc.Call[ImageSubtract](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageSubtract() ImageSubtract {
	return ImageSubtractClass.New()
}

func (i_ ImageSubtract) Init() ImageSubtract {
	rv := objc.Call[ImageSubtract](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageSubtract) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageSubtract {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageSubtract](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageSubtract_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageSubtract {
	instance := ImageSubtractClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
