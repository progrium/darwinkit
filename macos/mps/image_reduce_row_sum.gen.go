// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageReduceRowSum] class.
var ImageReduceRowSumClass = _ImageReduceRowSumClass{objc.GetClass("MPSImageReduceRowSum")}

type _ImageReduceRowSumClass struct {
	objc.Class
}

// An interface definition for the [ImageReduceRowSum] class.
type IImageReduceRowSum interface {
	IImageReduceUnary
}

// A filter that returns the sum of all values for a row in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducerowsum?language=objc
type ImageReduceRowSum struct {
	ImageReduceUnary
}

func ImageReduceRowSumFrom(ptr unsafe.Pointer) ImageReduceRowSum {
	return ImageReduceRowSum{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}

func (i_ ImageReduceRowSum) InitWithDevice(device metal.PDevice) ImageReduceRowSum {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageReduceRowSum](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducerowsum/2942334-initwithdevice?language=objc
func NewImageReduceRowSumWithDevice(device metal.PDevice) ImageReduceRowSum {
	instance := ImageReduceRowSumClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageReduceRowSumClass) Alloc() ImageReduceRowSum {
	rv := objc.Call[ImageReduceRowSum](ic, objc.Sel("alloc"))
	return rv
}

func ImageReduceRowSum_Alloc() ImageReduceRowSum {
	return ImageReduceRowSumClass.Alloc()
}

func (ic _ImageReduceRowSumClass) New() ImageReduceRowSum {
	rv := objc.Call[ImageReduceRowSum](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageReduceRowSum() ImageReduceRowSum {
	return ImageReduceRowSumClass.New()
}

func (i_ ImageReduceRowSum) Init() ImageReduceRowSum {
	rv := objc.Call[ImageReduceRowSum](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageReduceRowSum) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageReduceRowSum {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageReduceRowSum](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageReduceRowSum_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageReduceRowSum {
	instance := ImageReduceRowSumClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
