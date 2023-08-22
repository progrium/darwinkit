// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageReduceRowMin] class.
var ImageReduceRowMinClass = _ImageReduceRowMinClass{objc.GetClass("MPSImageReduceRowMin")}

type _ImageReduceRowMinClass struct {
	objc.Class
}

// An interface definition for the [ImageReduceRowMin] class.
type IImageReduceRowMin interface {
	IImageReduceUnary
}

// A filter that returns the minimum value for each row in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducerowmin?language=objc
type ImageReduceRowMin struct {
	ImageReduceUnary
}

func ImageReduceRowMinFrom(ptr unsafe.Pointer) ImageReduceRowMin {
	return ImageReduceRowMin{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}

func (i_ ImageReduceRowMin) InitWithDevice(device metal.PDevice) ImageReduceRowMin {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageReduceRowMin](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducerowmin/2942325-initwithdevice?language=objc
func NewImageReduceRowMinWithDevice(device metal.PDevice) ImageReduceRowMin {
	instance := ImageReduceRowMinClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageReduceRowMinClass) Alloc() ImageReduceRowMin {
	rv := objc.Call[ImageReduceRowMin](ic, objc.Sel("alloc"))
	return rv
}

func ImageReduceRowMin_Alloc() ImageReduceRowMin {
	return ImageReduceRowMinClass.Alloc()
}

func (ic _ImageReduceRowMinClass) New() ImageReduceRowMin {
	rv := objc.Call[ImageReduceRowMin](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageReduceRowMin() ImageReduceRowMin {
	return ImageReduceRowMinClass.New()
}

func (i_ ImageReduceRowMin) Init() ImageReduceRowMin {
	rv := objc.Call[ImageReduceRowMin](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageReduceRowMin) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageReduceRowMin {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageReduceRowMin](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageReduceRowMin_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageReduceRowMin {
	instance := ImageReduceRowMinClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
