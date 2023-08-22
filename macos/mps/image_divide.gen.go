// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageDivide] class.
var ImageDivideClass = _ImageDivideClass{objc.GetClass("MPSImageDivide")}

type _ImageDivideClass struct {
	objc.Class
}

// An interface definition for the [ImageDivide] class.
type IImageDivide interface {
	IImageArithmetic
}

// A filter that returns the element-wise quotient of its two input images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedivide?language=objc
type ImageDivide struct {
	ImageArithmetic
}

func ImageDivideFrom(ptr unsafe.Pointer) ImageDivide {
	return ImageDivide{
		ImageArithmetic: ImageArithmeticFrom(ptr),
	}
}

func (i_ ImageDivide) InitWithDevice(device metal.PDevice) ImageDivide {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageDivide](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedivide/2866606-initwithdevice?language=objc
func NewImageDivideWithDevice(device metal.PDevice) ImageDivide {
	instance := ImageDivideClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageDivideClass) Alloc() ImageDivide {
	rv := objc.Call[ImageDivide](ic, objc.Sel("alloc"))
	return rv
}

func ImageDivide_Alloc() ImageDivide {
	return ImageDivideClass.Alloc()
}

func (ic _ImageDivideClass) New() ImageDivide {
	rv := objc.Call[ImageDivide](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageDivide() ImageDivide {
	return ImageDivideClass.New()
}

func (i_ ImageDivide) Init() ImageDivide {
	rv := objc.Call[ImageDivide](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageDivide) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageDivide {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageDivide](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageDivide_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageDivide {
	instance := ImageDivideClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
