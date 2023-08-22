// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageReduceColumnSum] class.
var ImageReduceColumnSumClass = _ImageReduceColumnSumClass{objc.GetClass("MPSImageReduceColumnSum")}

type _ImageReduceColumnSumClass struct {
	objc.Class
}

// An interface definition for the [ImageReduceColumnSum] class.
type IImageReduceColumnSum interface {
	IImageReduceUnary
}

// A filter that returns the sum of all values for a column in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducecolumnsum?language=objc
type ImageReduceColumnSum struct {
	ImageReduceUnary
}

func ImageReduceColumnSumFrom(ptr unsafe.Pointer) ImageReduceColumnSum {
	return ImageReduceColumnSum{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}

func (i_ ImageReduceColumnSum) InitWithDevice(device metal.PDevice) ImageReduceColumnSum {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageReduceColumnSum](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducecolumnsum/2942321-initwithdevice?language=objc
func NewImageReduceColumnSumWithDevice(device metal.PDevice) ImageReduceColumnSum {
	instance := ImageReduceColumnSumClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageReduceColumnSumClass) Alloc() ImageReduceColumnSum {
	rv := objc.Call[ImageReduceColumnSum](ic, objc.Sel("alloc"))
	return rv
}

func ImageReduceColumnSum_Alloc() ImageReduceColumnSum {
	return ImageReduceColumnSumClass.Alloc()
}

func (ic _ImageReduceColumnSumClass) New() ImageReduceColumnSum {
	rv := objc.Call[ImageReduceColumnSum](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageReduceColumnSum() ImageReduceColumnSum {
	return ImageReduceColumnSumClass.New()
}

func (i_ ImageReduceColumnSum) Init() ImageReduceColumnSum {
	rv := objc.Call[ImageReduceColumnSum](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageReduceColumnSum) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageReduceColumnSum {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageReduceColumnSum](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageReduceColumnSum_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageReduceColumnSum {
	instance := ImageReduceColumnSumClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
