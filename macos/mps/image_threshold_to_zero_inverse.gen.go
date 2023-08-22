// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageThresholdToZeroInverse] class.
var ImageThresholdToZeroInverseClass = _ImageThresholdToZeroInverseClass{objc.GetClass("MPSImageThresholdToZeroInverse")}

type _ImageThresholdToZeroInverseClass struct {
	objc.Class
}

// An interface definition for the [ImageThresholdToZeroInverse] class.
type IImageThresholdToZeroInverse interface {
	IUnaryImageKernel
	Transform() *float64
	ThresholdValue() float64
}

// A filter that returns 0 for each pixel with a value greater than a specified threshold or the original value otherwise. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozeroinverse?language=objc
type ImageThresholdToZeroInverse struct {
	UnaryImageKernel
}

func ImageThresholdToZeroInverseFrom(ptr unsafe.Pointer) ImageThresholdToZeroInverse {
	return ImageThresholdToZeroInverse{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageThresholdToZeroInverse) InitWithDeviceThresholdValueLinearGrayColorTransform(device metal.PDevice, thresholdValue float64, transform *float64) ImageThresholdToZeroInverse {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageThresholdToZeroInverse](i_, objc.Sel("initWithDevice:thresholdValue:linearGrayColorTransform:"), po0, thresholdValue, transform)
	return rv
}

// Initializes the kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozeroinverse/1618911-initwithdevice?language=objc
func NewImageThresholdToZeroInverseWithDeviceThresholdValueLinearGrayColorTransform(device metal.PDevice, thresholdValue float64, transform *float64) ImageThresholdToZeroInverse {
	instance := ImageThresholdToZeroInverseClass.Alloc().InitWithDeviceThresholdValueLinearGrayColorTransform(device, thresholdValue, transform)
	instance.Autorelease()
	return instance
}

func (ic _ImageThresholdToZeroInverseClass) Alloc() ImageThresholdToZeroInverse {
	rv := objc.Call[ImageThresholdToZeroInverse](ic, objc.Sel("alloc"))
	return rv
}

func ImageThresholdToZeroInverse_Alloc() ImageThresholdToZeroInverse {
	return ImageThresholdToZeroInverseClass.Alloc()
}

func (ic _ImageThresholdToZeroInverseClass) New() ImageThresholdToZeroInverse {
	rv := objc.Call[ImageThresholdToZeroInverse](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageThresholdToZeroInverse() ImageThresholdToZeroInverse {
	return ImageThresholdToZeroInverseClass.New()
}

func (i_ ImageThresholdToZeroInverse) Init() ImageThresholdToZeroInverse {
	rv := objc.Call[ImageThresholdToZeroInverse](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageThresholdToZeroInverse) InitWithDevice(device metal.PDevice) ImageThresholdToZeroInverse {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageThresholdToZeroInverse](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageThresholdToZeroInverseWithDevice(device metal.PDevice) ImageThresholdToZeroInverse {
	instance := ImageThresholdToZeroInverseClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageThresholdToZeroInverse) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageThresholdToZeroInverse {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageThresholdToZeroInverse](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageThresholdToZeroInverse_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageThresholdToZeroInverse {
	instance := ImageThresholdToZeroInverseClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// The color transform used to initialize the threshold filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozeroinverse/1618828-transform?language=objc
func (i_ ImageThresholdToZeroInverse) Transform() *float64 {
	rv := objc.Call[*float64](i_, objc.Sel("transform"))
	return rv
}

// The threshold value used to initialize the threshold filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozeroinverse/1618914-thresholdvalue?language=objc
func (i_ ImageThresholdToZeroInverse) ThresholdValue() float64 {
	rv := objc.Call[float64](i_, objc.Sel("thresholdValue"))
	return rv
}
