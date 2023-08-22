// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageReduceColumnMin] class.
var ImageReduceColumnMinClass = _ImageReduceColumnMinClass{objc.GetClass("MPSImageReduceColumnMin")}

type _ImageReduceColumnMinClass struct {
	objc.Class
}

// An interface definition for the [ImageReduceColumnMin] class.
type IImageReduceColumnMin interface {
	IImageReduceUnary
}

// A filter that returns the minimum value for each column in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducecolumnmin?language=objc
type ImageReduceColumnMin struct {
	ImageReduceUnary
}

func ImageReduceColumnMinFrom(ptr unsafe.Pointer) ImageReduceColumnMin {
	return ImageReduceColumnMin{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}

func (i_ ImageReduceColumnMin) InitWithDevice(device metal.PDevice) ImageReduceColumnMin {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageReduceColumnMin](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducecolumnmin/2942333-initwithdevice?language=objc
func NewImageReduceColumnMinWithDevice(device metal.PDevice) ImageReduceColumnMin {
	instance := ImageReduceColumnMinClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageReduceColumnMinClass) Alloc() ImageReduceColumnMin {
	rv := objc.Call[ImageReduceColumnMin](ic, objc.Sel("alloc"))
	return rv
}

func ImageReduceColumnMin_Alloc() ImageReduceColumnMin {
	return ImageReduceColumnMinClass.Alloc()
}

func (ic _ImageReduceColumnMinClass) New() ImageReduceColumnMin {
	rv := objc.Call[ImageReduceColumnMin](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageReduceColumnMin() ImageReduceColumnMin {
	return ImageReduceColumnMinClass.New()
}

func (i_ ImageReduceColumnMin) Init() ImageReduceColumnMin {
	rv := objc.Call[ImageReduceColumnMin](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageReduceColumnMin) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageReduceColumnMin {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageReduceColumnMin](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageReduceColumnMin_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageReduceColumnMin {
	instance := ImageReduceColumnMinClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
