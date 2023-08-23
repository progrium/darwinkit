// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageThresholdBinaryInverse] class.
var ImageThresholdBinaryInverseClass = _ImageThresholdBinaryInverseClass{objc.GetClass("MPSImageThresholdBinaryInverse")}

type _ImageThresholdBinaryInverseClass struct {
	objc.Class
}

// An interface definition for the [ImageThresholdBinaryInverse] class.
type IImageThresholdBinaryInverse interface {
	IUnaryImageKernel
	MaximumValue() float64
	Transform() *float64
	ThresholdValue() float64
}

// A filter that returns 0 for each pixel with a value greater than a specified threshold or a specified value otherwise. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinaryinverse?language=objc
type ImageThresholdBinaryInverse struct {
	UnaryImageKernel
}

func ImageThresholdBinaryInverseFrom(ptr unsafe.Pointer) ImageThresholdBinaryInverse {
	return ImageThresholdBinaryInverse{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageThresholdBinaryInverse) InitWithDeviceThresholdValueMaximumValueLinearGrayColorTransform(device metal.PDevice, thresholdValue float64, maximumValue float64, transform *float64) ImageThresholdBinaryInverse {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageThresholdBinaryInverse](i_, objc.Sel("initWithDevice:thresholdValue:maximumValue:linearGrayColorTransform:"), po0, thresholdValue, maximumValue, transform)
	return rv
}

// Initializes the kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinaryinverse/1618903-initwithdevice?language=objc
func NewImageThresholdBinaryInverseWithDeviceThresholdValueMaximumValueLinearGrayColorTransform(device metal.PDevice, thresholdValue float64, maximumValue float64, transform *float64) ImageThresholdBinaryInverse {
	instance := ImageThresholdBinaryInverseClass.Alloc().InitWithDeviceThresholdValueMaximumValueLinearGrayColorTransform(device, thresholdValue, maximumValue, transform)
	instance.Autorelease()
	return instance
}

func (ic _ImageThresholdBinaryInverseClass) Alloc() ImageThresholdBinaryInverse {
	rv := objc.Call[ImageThresholdBinaryInverse](ic, objc.Sel("alloc"))
	return rv
}

func ImageThresholdBinaryInverse_Alloc() ImageThresholdBinaryInverse {
	return ImageThresholdBinaryInverseClass.Alloc()
}

func (ic _ImageThresholdBinaryInverseClass) New() ImageThresholdBinaryInverse {
	rv := objc.Call[ImageThresholdBinaryInverse](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageThresholdBinaryInverse() ImageThresholdBinaryInverse {
	return ImageThresholdBinaryInverseClass.New()
}

func (i_ ImageThresholdBinaryInverse) Init() ImageThresholdBinaryInverse {
	rv := objc.Call[ImageThresholdBinaryInverse](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageThresholdBinaryInverse) InitWithDevice(device metal.PDevice) ImageThresholdBinaryInverse {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageThresholdBinaryInverse](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageThresholdBinaryInverseWithDevice(device metal.PDevice) ImageThresholdBinaryInverse {
	instance := ImageThresholdBinaryInverseClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageThresholdBinaryInverse) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageThresholdBinaryInverse {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageThresholdBinaryInverse](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageThresholdBinaryInverse_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageThresholdBinaryInverse {
	instance := ImageThresholdBinaryInverseClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// The maximum value used to initialize the threshold filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinaryinverse/1618906-maximumvalue?language=objc
func (i_ ImageThresholdBinaryInverse) MaximumValue() float64 {
	rv := objc.Call[float64](i_, objc.Sel("maximumValue"))
	return rv
}

// The color transform used to initialize the threshold filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinaryinverse/1618904-transform?language=objc
func (i_ ImageThresholdBinaryInverse) Transform() *float64 {
	rv := objc.Call[*float64](i_, objc.Sel("transform"))
	return rv
}

// The threshold value used to initialize the threshold filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinaryinverse/1618845-thresholdvalue?language=objc
func (i_ ImageThresholdBinaryInverse) ThresholdValue() float64 {
	rv := objc.Call[float64](i_, objc.Sel("thresholdValue"))
	return rv
}
