// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageGaussianBlur] class.
var ImageGaussianBlurClass = _ImageGaussianBlurClass{objc.GetClass("MPSImageGaussianBlur")}

type _ImageGaussianBlurClass struct {
	objc.Class
}

// An interface definition for the [ImageGaussianBlur] class.
type IImageGaussianBlur interface {
	IUnaryImageKernel
	Sigma() float64
}

// A filter that convolves an image with a Gaussian blur of a given sigma in both the x and y directions. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagegaussianblur?language=objc
type ImageGaussianBlur struct {
	UnaryImageKernel
}

func ImageGaussianBlurFrom(ptr unsafe.Pointer) ImageGaussianBlur {
	return ImageGaussianBlur{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageGaussianBlur) InitWithDeviceSigma(device metal.PDevice, sigma float64) ImageGaussianBlur {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageGaussianBlur](i_, objc.Sel("initWithDevice:sigma:"), po0, sigma)
	return rv
}

// Initializes a Gaussian blur filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagegaussianblur/1618813-initwithdevice?language=objc
func NewImageGaussianBlurWithDeviceSigma(device metal.PDevice, sigma float64) ImageGaussianBlur {
	instance := ImageGaussianBlurClass.Alloc().InitWithDeviceSigma(device, sigma)
	instance.Autorelease()
	return instance
}

func (ic _ImageGaussianBlurClass) Alloc() ImageGaussianBlur {
	rv := objc.Call[ImageGaussianBlur](ic, objc.Sel("alloc"))
	return rv
}

func ImageGaussianBlur_Alloc() ImageGaussianBlur {
	return ImageGaussianBlurClass.Alloc()
}

func (ic _ImageGaussianBlurClass) New() ImageGaussianBlur {
	rv := objc.Call[ImageGaussianBlur](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageGaussianBlur() ImageGaussianBlur {
	return ImageGaussianBlurClass.New()
}

func (i_ ImageGaussianBlur) Init() ImageGaussianBlur {
	rv := objc.Call[ImageGaussianBlur](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageGaussianBlur) InitWithDevice(device metal.PDevice) ImageGaussianBlur {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageGaussianBlur](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageGaussianBlurWithDevice(device metal.PDevice) ImageGaussianBlur {
	instance := ImageGaussianBlurClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageGaussianBlur) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageGaussianBlur {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageGaussianBlur](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageGaussianBlur_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageGaussianBlur {
	instance := ImageGaussianBlurClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// The sigma value with which the filter was created. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagegaussianblur/1618850-sigma?language=objc
func (i_ ImageGaussianBlur) Sigma() float64 {
	rv := objc.Call[float64](i_, objc.Sel("sigma"))
	return rv
}
