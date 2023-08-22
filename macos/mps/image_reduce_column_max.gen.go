// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageReduceColumnMax] class.
var ImageReduceColumnMaxClass = _ImageReduceColumnMaxClass{objc.GetClass("MPSImageReduceColumnMax")}

type _ImageReduceColumnMaxClass struct {
	objc.Class
}

// An interface definition for the [ImageReduceColumnMax] class.
type IImageReduceColumnMax interface {
	IImageReduceUnary
}

// A filter that returns the maximum value for each column in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducecolumnmax?language=objc
type ImageReduceColumnMax struct {
	ImageReduceUnary
}

func ImageReduceColumnMaxFrom(ptr unsafe.Pointer) ImageReduceColumnMax {
	return ImageReduceColumnMax{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}

func (i_ ImageReduceColumnMax) InitWithDevice(device metal.PDevice) ImageReduceColumnMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageReduceColumnMax](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducecolumnmax/2942318-initwithdevice?language=objc
func NewImageReduceColumnMaxWithDevice(device metal.PDevice) ImageReduceColumnMax {
	instance := ImageReduceColumnMaxClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageReduceColumnMaxClass) Alloc() ImageReduceColumnMax {
	rv := objc.Call[ImageReduceColumnMax](ic, objc.Sel("alloc"))
	return rv
}

func ImageReduceColumnMax_Alloc() ImageReduceColumnMax {
	return ImageReduceColumnMaxClass.Alloc()
}

func (ic _ImageReduceColumnMaxClass) New() ImageReduceColumnMax {
	rv := objc.Call[ImageReduceColumnMax](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageReduceColumnMax() ImageReduceColumnMax {
	return ImageReduceColumnMaxClass.New()
}

func (i_ ImageReduceColumnMax) Init() ImageReduceColumnMax {
	rv := objc.Call[ImageReduceColumnMax](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageReduceColumnMax) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageReduceColumnMax {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageReduceColumnMax](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageReduceColumnMax_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageReduceColumnMax {
	instance := ImageReduceColumnMaxClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
