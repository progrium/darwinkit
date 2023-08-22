// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageIntegralOfSquares] class.
var ImageIntegralOfSquaresClass = _ImageIntegralOfSquaresClass{objc.GetClass("MPSImageIntegralOfSquares")}

type _ImageIntegralOfSquaresClass struct {
	objc.Class
}

// An interface definition for the [ImageIntegralOfSquares] class.
type IImageIntegralOfSquares interface {
	IUnaryImageKernel
}

// A filter that calculates the sum of squared pixels over a specified region in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageintegralofsquares?language=objc
type ImageIntegralOfSquares struct {
	UnaryImageKernel
}

func ImageIntegralOfSquaresFrom(ptr unsafe.Pointer) ImageIntegralOfSquares {
	return ImageIntegralOfSquares{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (ic _ImageIntegralOfSquaresClass) Alloc() ImageIntegralOfSquares {
	rv := objc.Call[ImageIntegralOfSquares](ic, objc.Sel("alloc"))
	return rv
}

func ImageIntegralOfSquares_Alloc() ImageIntegralOfSquares {
	return ImageIntegralOfSquaresClass.Alloc()
}

func (ic _ImageIntegralOfSquaresClass) New() ImageIntegralOfSquares {
	rv := objc.Call[ImageIntegralOfSquares](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageIntegralOfSquares() ImageIntegralOfSquares {
	return ImageIntegralOfSquaresClass.New()
}

func (i_ ImageIntegralOfSquares) Init() ImageIntegralOfSquares {
	rv := objc.Call[ImageIntegralOfSquares](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageIntegralOfSquares) InitWithDevice(device metal.PDevice) ImageIntegralOfSquares {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageIntegralOfSquares](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageIntegralOfSquaresWithDevice(device metal.PDevice) ImageIntegralOfSquares {
	instance := ImageIntegralOfSquaresClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageIntegralOfSquares) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageIntegralOfSquares {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageIntegralOfSquares](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageIntegralOfSquares_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageIntegralOfSquares {
	instance := ImageIntegralOfSquaresClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
