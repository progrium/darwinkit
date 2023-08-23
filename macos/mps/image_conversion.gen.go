// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageConversion] class.
var ImageConversionClass = _ImageConversionClass{objc.GetClass("MPSImageConversion")}

type _ImageConversionClass struct {
	objc.Class
}

// An interface definition for the [ImageConversion] class.
type IImageConversion interface {
	IUnaryImageKernel
	SourceAlpha() AlphaType
	DestinationAlpha() AlphaType
}

// A filter that performs a conversion of color space, alpha, or pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageconversion?language=objc
type ImageConversion struct {
	UnaryImageKernel
}

func ImageConversionFrom(ptr unsafe.Pointer) ImageConversion {
	return ImageConversion{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageConversion) InitWithDeviceSrcAlphaDestAlphaBackgroundColorConversionInfo(device metal.PDevice, srcAlpha AlphaType, destAlpha AlphaType, backgroundColor *float64, conversionInfo coregraphics.ColorConversionInfoRef) ImageConversion {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageConversion](i_, objc.Sel("initWithDevice:srcAlpha:destAlpha:backgroundColor:conversionInfo:"), po0, srcAlpha, destAlpha, backgroundColor, conversionInfo)
	return rv
}

// Initializes a filter that can convert texture color space, alpha, and pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageconversion/2206722-initwithdevice?language=objc
func NewImageConversionWithDeviceSrcAlphaDestAlphaBackgroundColorConversionInfo(device metal.PDevice, srcAlpha AlphaType, destAlpha AlphaType, backgroundColor *float64, conversionInfo coregraphics.ColorConversionInfoRef) ImageConversion {
	instance := ImageConversionClass.Alloc().InitWithDeviceSrcAlphaDestAlphaBackgroundColorConversionInfo(device, srcAlpha, destAlpha, backgroundColor, conversionInfo)
	instance.Autorelease()
	return instance
}

func (ic _ImageConversionClass) Alloc() ImageConversion {
	rv := objc.Call[ImageConversion](ic, objc.Sel("alloc"))
	return rv
}

func ImageConversion_Alloc() ImageConversion {
	return ImageConversionClass.Alloc()
}

func (ic _ImageConversionClass) New() ImageConversion {
	rv := objc.Call[ImageConversion](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageConversion() ImageConversion {
	return ImageConversionClass.New()
}

func (i_ ImageConversion) Init() ImageConversion {
	rv := objc.Call[ImageConversion](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageConversion) InitWithDevice(device metal.PDevice) ImageConversion {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageConversion](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageConversionWithDevice(device metal.PDevice) ImageConversion {
	instance := ImageConversionClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageConversion) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageConversion {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageConversion](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageConversion_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageConversion {
	instance := ImageConversionClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// Premultiplication description for the source texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageconversion/1648518-sourcealpha?language=objc
func (i_ ImageConversion) SourceAlpha() AlphaType {
	rv := objc.Call[AlphaType](i_, objc.Sel("sourceAlpha"))
	return rv
}

// Premultiplication description for the destination texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageconversion/1648515-destinationalpha?language=objc
func (i_ ImageConversion) DestinationAlpha() AlphaType {
	rv := objc.Call[AlphaType](i_, objc.Sel("destinationAlpha"))
	return rv
}
