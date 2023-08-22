// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageReduceUnary] class.
var ImageReduceUnaryClass = _ImageReduceUnaryClass{objc.GetClass("MPSImageReduceUnary")}

type _ImageReduceUnaryClass struct {
	objc.Class
}

// An interface definition for the [ImageReduceUnary] class.
type IImageReduceUnary interface {
	IUnaryImageKernel
	ClipRectSource() metal.Region
	SetClipRectSource(value metal.Region)
}

// The base class for reduction filters that take a single source as input. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereduceunary?language=objc
type ImageReduceUnary struct {
	UnaryImageKernel
}

func ImageReduceUnaryFrom(ptr unsafe.Pointer) ImageReduceUnary {
	return ImageReduceUnary{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (ic _ImageReduceUnaryClass) Alloc() ImageReduceUnary {
	rv := objc.Call[ImageReduceUnary](ic, objc.Sel("alloc"))
	return rv
}

func ImageReduceUnary_Alloc() ImageReduceUnary {
	return ImageReduceUnaryClass.Alloc()
}

func (ic _ImageReduceUnaryClass) New() ImageReduceUnary {
	rv := objc.Call[ImageReduceUnary](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageReduceUnary() ImageReduceUnary {
	return ImageReduceUnaryClass.New()
}

func (i_ ImageReduceUnary) Init() ImageReduceUnary {
	rv := objc.Call[ImageReduceUnary](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageReduceUnary) InitWithDevice(device metal.PDevice) ImageReduceUnary {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageReduceUnary](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageReduceUnaryWithDevice(device metal.PDevice) ImageReduceUnary {
	instance := ImageReduceUnaryClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageReduceUnary) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageReduceUnary {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageReduceUnary](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageReduceUnary_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageReduceUnary {
	instance := ImageReduceUnaryClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereduceunary/2942332-cliprectsource?language=objc
func (i_ ImageReduceUnary) ClipRectSource() metal.Region {
	rv := objc.Call[metal.Region](i_, objc.Sel("clipRectSource"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereduceunary/2942332-cliprectsource?language=objc
func (i_ ImageReduceUnary) SetClipRectSource(value metal.Region) {
	objc.Call[objc.Void](i_, objc.Sel("setClipRectSource:"), value)
}
