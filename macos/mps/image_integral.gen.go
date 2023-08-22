// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageIntegral] class.
var ImageIntegralClass = _ImageIntegralClass{objc.GetClass("MPSImageIntegral")}

type _ImageIntegralClass struct {
	objc.Class
}

// An interface definition for the [ImageIntegral] class.
type IImageIntegral interface {
	IUnaryImageKernel
}

// A filter that calculates the sum of pixels over a specified region in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageintegral?language=objc
type ImageIntegral struct {
	UnaryImageKernel
}

func ImageIntegralFrom(ptr unsafe.Pointer) ImageIntegral {
	return ImageIntegral{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (ic _ImageIntegralClass) Alloc() ImageIntegral {
	rv := objc.Call[ImageIntegral](ic, objc.Sel("alloc"))
	return rv
}

func ImageIntegral_Alloc() ImageIntegral {
	return ImageIntegralClass.Alloc()
}

func (ic _ImageIntegralClass) New() ImageIntegral {
	rv := objc.Call[ImageIntegral](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageIntegral() ImageIntegral {
	return ImageIntegralClass.New()
}

func (i_ ImageIntegral) Init() ImageIntegral {
	rv := objc.Call[ImageIntegral](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageIntegral) InitWithDevice(device metal.PDevice) ImageIntegral {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageIntegral](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageIntegralWithDevice(device metal.PDevice) ImageIntegral {
	instance := ImageIntegralClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageIntegral) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageIntegral {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageIntegral](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageIntegral_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageIntegral {
	instance := ImageIntegralClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
