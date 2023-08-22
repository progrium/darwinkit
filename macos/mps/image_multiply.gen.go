// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageMultiply] class.
var ImageMultiplyClass = _ImageMultiplyClass{objc.GetClass("MPSImageMultiply")}

type _ImageMultiplyClass struct {
	objc.Class
}

// An interface definition for the [ImageMultiply] class.
type IImageMultiply interface {
	IImageArithmetic
}

// A filter that returns the element-wise product of its two input images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagemultiply?language=objc
type ImageMultiply struct {
	ImageArithmetic
}

func ImageMultiplyFrom(ptr unsafe.Pointer) ImageMultiply {
	return ImageMultiply{
		ImageArithmetic: ImageArithmeticFrom(ptr),
	}
}

func (i_ ImageMultiply) InitWithDevice(device metal.PDevice) ImageMultiply {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageMultiply](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagemultiply/2866600-initwithdevice?language=objc
func NewImageMultiplyWithDevice(device metal.PDevice) ImageMultiply {
	instance := ImageMultiplyClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageMultiplyClass) Alloc() ImageMultiply {
	rv := objc.Call[ImageMultiply](ic, objc.Sel("alloc"))
	return rv
}

func ImageMultiply_Alloc() ImageMultiply {
	return ImageMultiplyClass.Alloc()
}

func (ic _ImageMultiplyClass) New() ImageMultiply {
	rv := objc.Call[ImageMultiply](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageMultiply() ImageMultiply {
	return ImageMultiplyClass.New()
}

func (i_ ImageMultiply) Init() ImageMultiply {
	rv := objc.Call[ImageMultiply](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageMultiply) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageMultiply {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageMultiply](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageMultiply_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageMultiply {
	instance := ImageMultiplyClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
