// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageSobel] class.
var ImageSobelClass = _ImageSobelClass{objc.GetClass("MPSImageSobel")}

type _ImageSobelClass struct {
	objc.Class
}

// An interface definition for the [ImageSobel] class.
type IImageSobel interface {
	IUnaryImageKernel
	ColorTransform() *float64
}

// A filter that convolves an image with the Sobel operator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesobel?language=objc
type ImageSobel struct {
	UnaryImageKernel
}

func ImageSobelFrom(ptr unsafe.Pointer) ImageSobel {
	return ImageSobel{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageSobel) InitWithDevice(device metal.PDevice) ImageSobel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageSobel](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a Sobel filter on a given device using the default color transform. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesobel/1618843-initwithdevice?language=objc
func NewImageSobelWithDevice(device metal.PDevice) ImageSobel {
	instance := ImageSobelClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageSobelClass) Alloc() ImageSobel {
	rv := objc.Call[ImageSobel](ic, objc.Sel("alloc"))
	return rv
}

func ImageSobel_Alloc() ImageSobel {
	return ImageSobelClass.Alloc()
}

func (ic _ImageSobelClass) New() ImageSobel {
	rv := objc.Call[ImageSobel](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageSobel() ImageSobel {
	return ImageSobelClass.New()
}

func (i_ ImageSobel) Init() ImageSobel {
	rv := objc.Call[ImageSobel](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageSobel) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageSobel {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageSobel](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageSobel_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageSobel {
	instance := ImageSobelClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// The color transform used to initialize the Sobel filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagesobel/1618777-colortransform?language=objc
func (i_ ImageSobel) ColorTransform() *float64 {
	rv := objc.Call[*float64](i_, objc.Sel("colorTransform"))
	return rv
}
