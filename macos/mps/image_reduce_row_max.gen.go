// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageReduceRowMax] class.
var ImageReduceRowMaxClass = _ImageReduceRowMaxClass{objc.GetClass("MPSImageReduceRowMax")}

type _ImageReduceRowMaxClass struct {
	objc.Class
}

// An interface definition for the [ImageReduceRowMax] class.
type IImageReduceRowMax interface {
	IImageReduceUnary
}

// A filter that returns the maximum value for each row in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducerowmax?language=objc
type ImageReduceRowMax struct {
	ImageReduceUnary
}

func ImageReduceRowMaxFrom(ptr unsafe.Pointer) ImageReduceRowMax {
	return ImageReduceRowMax{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}

func (i_ ImageReduceRowMax) InitWithDevice(device metal.PDevice) ImageReduceRowMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageReduceRowMax](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducerowmax/2942328-initwithdevice?language=objc
func NewImageReduceRowMaxWithDevice(device metal.PDevice) ImageReduceRowMax {
	instance := ImageReduceRowMaxClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageReduceRowMaxClass) Alloc() ImageReduceRowMax {
	rv := objc.Call[ImageReduceRowMax](ic, objc.Sel("alloc"))
	return rv
}

func ImageReduceRowMax_Alloc() ImageReduceRowMax {
	return ImageReduceRowMaxClass.Alloc()
}

func (ic _ImageReduceRowMaxClass) New() ImageReduceRowMax {
	rv := objc.Call[ImageReduceRowMax](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageReduceRowMax() ImageReduceRowMax {
	return ImageReduceRowMaxClass.New()
}

func (i_ ImageReduceRowMax) Init() ImageReduceRowMax {
	rv := objc.Call[ImageReduceRowMax](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageReduceRowMax) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageReduceRowMax {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageReduceRowMax](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageReduceRowMax_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageReduceRowMax {
	instance := ImageReduceRowMaxClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
