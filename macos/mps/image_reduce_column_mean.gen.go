// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageReduceColumnMean] class.
var ImageReduceColumnMeanClass = _ImageReduceColumnMeanClass{objc.GetClass("MPSImageReduceColumnMean")}

type _ImageReduceColumnMeanClass struct {
	objc.Class
}

// An interface definition for the [ImageReduceColumnMean] class.
type IImageReduceColumnMean interface {
	IImageReduceUnary
}

// A filter that returns the mean value for each column in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducecolumnmean?language=objc
type ImageReduceColumnMean struct {
	ImageReduceUnary
}

func ImageReduceColumnMeanFrom(ptr unsafe.Pointer) ImageReduceColumnMean {
	return ImageReduceColumnMean{
		ImageReduceUnary: ImageReduceUnaryFrom(ptr),
	}
}

func (i_ ImageReduceColumnMean) InitWithDevice(device metal.PDevice) ImageReduceColumnMean {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageReduceColumnMean](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagereducecolumnmean/2942331-initwithdevice?language=objc
func NewImageReduceColumnMeanWithDevice(device metal.PDevice) ImageReduceColumnMean {
	instance := ImageReduceColumnMeanClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageReduceColumnMeanClass) Alloc() ImageReduceColumnMean {
	rv := objc.Call[ImageReduceColumnMean](ic, objc.Sel("alloc"))
	return rv
}

func ImageReduceColumnMean_Alloc() ImageReduceColumnMean {
	return ImageReduceColumnMeanClass.Alloc()
}

func (ic _ImageReduceColumnMeanClass) New() ImageReduceColumnMean {
	rv := objc.Call[ImageReduceColumnMean](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageReduceColumnMean() ImageReduceColumnMean {
	return ImageReduceColumnMeanClass.New()
}

func (i_ ImageReduceColumnMean) Init() ImageReduceColumnMean {
	rv := objc.Call[ImageReduceColumnMean](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageReduceColumnMean) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageReduceColumnMean {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageReduceColumnMean](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageReduceColumnMean_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageReduceColumnMean {
	instance := ImageReduceColumnMeanClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
