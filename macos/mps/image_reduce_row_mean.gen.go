// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageReduceRowMean] class.
var ImageReduceRowMeanClass = _ImageReduceRowMeanClass{objc.GetClass("MPSImageReduceRowMean")}

type _ImageReduceRowMeanClass struct {
	objc.Class
}

// An interface definition for the [ImageReduceRowMean] class.
type IImageReduceRowMean interface {
	IImageReduceUnary
}

// A filter that returns the mean value for each row in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducerowmean?language=objc
type ImageReduceRowMean struct {
	ImageReduceUnary
}

func ImageReduceRowMeanFrom(ptr unsafe.Pointer) ImageReduceRowMean {
	return ImageReduceRowMean{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}

func (i_ ImageReduceRowMean) InitWithDevice(device metal.PDevice) ImageReduceRowMean {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageReduceRowMean](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducerowmean/2942322-initwithdevice?language=objc
func NewImageReduceRowMeanWithDevice(device metal.PDevice) ImageReduceRowMean {
	instance := ImageReduceRowMeanClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageReduceRowMeanClass) Alloc() ImageReduceRowMean {
	rv := objc.Call[ImageReduceRowMean](ic, objc.Sel("alloc"))
	return rv
}

func ImageReduceRowMean_Alloc() ImageReduceRowMean {
	return ImageReduceRowMeanClass.Alloc()
}

func (ic _ImageReduceRowMeanClass) New() ImageReduceRowMean {
	rv := objc.Call[ImageReduceRowMean](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageReduceRowMean() ImageReduceRowMean {
	return ImageReduceRowMeanClass.New()
}

func (i_ ImageReduceRowMean) Init() ImageReduceRowMean {
	rv := objc.Call[ImageReduceRowMean](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageReduceRowMean) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageReduceRowMean {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageReduceRowMean](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageReduceRowMean_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageReduceRowMean {
	instance := ImageReduceRowMeanClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
