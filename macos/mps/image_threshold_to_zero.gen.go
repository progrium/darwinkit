// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageThresholdToZero] class.
var ImageThresholdToZeroClass = _ImageThresholdToZeroClass{objc.GetClass("MPSImageThresholdToZero")}

type _ImageThresholdToZeroClass struct {
	objc.Class
}

// An interface definition for the [ImageThresholdToZero] class.
type IImageThresholdToZero interface {
	IUnaryImageKernel
	Transform() *float64
	ThresholdValue() float64
}

// A filter that returns the original value for each pixel with a value greater than a specified threshold or 0 otherwise. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozero?language=objc
type ImageThresholdToZero struct {
	UnaryImageKernel
}

func ImageThresholdToZeroFrom(ptr unsafe.Pointer) ImageThresholdToZero {
	return ImageThresholdToZero{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageThresholdToZero) InitWithDeviceThresholdValueLinearGrayColorTransform(device metal.PDevice, thresholdValue float64, transform *float64) ImageThresholdToZero {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageThresholdToZero](i_, objc.Sel("initWithDevice:thresholdValue:linearGrayColorTransform:"), po0, thresholdValue, transform)
	return rv
}

// Initializes the kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozero/1618865-initwithdevice?language=objc
func NewImageThresholdToZeroWithDeviceThresholdValueLinearGrayColorTransform(device metal.PDevice, thresholdValue float64, transform *float64) ImageThresholdToZero {
	instance := ImageThresholdToZeroClass.Alloc().InitWithDeviceThresholdValueLinearGrayColorTransform(device, thresholdValue, transform)
	instance.Autorelease()
	return instance
}

func (ic _ImageThresholdToZeroClass) Alloc() ImageThresholdToZero {
	rv := objc.Call[ImageThresholdToZero](ic, objc.Sel("alloc"))
	return rv
}

func ImageThresholdToZero_Alloc() ImageThresholdToZero {
	return ImageThresholdToZeroClass.Alloc()
}

func (ic _ImageThresholdToZeroClass) New() ImageThresholdToZero {
	rv := objc.Call[ImageThresholdToZero](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageThresholdToZero() ImageThresholdToZero {
	return ImageThresholdToZeroClass.New()
}

func (i_ ImageThresholdToZero) Init() ImageThresholdToZero {
	rv := objc.Call[ImageThresholdToZero](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageThresholdToZero) InitWithDevice(device metal.PDevice) ImageThresholdToZero {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageThresholdToZero](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageThresholdToZeroWithDevice(device metal.PDevice) ImageThresholdToZero {
	instance := ImageThresholdToZeroClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageThresholdToZero) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageThresholdToZero {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageThresholdToZero](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageThresholdToZero_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageThresholdToZero {
	instance := ImageThresholdToZeroClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// The color transform used to initialize the threshold filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozero/1618823-transform?language=objc
func (i_ ImageThresholdToZero) Transform() *float64 {
	rv := objc.Call[*float64](i_, objc.Sel("transform"))
	return rv
}

// The threshold value used to initialize the threshold filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdtozero/1618767-thresholdvalue?language=objc
func (i_ ImageThresholdToZero) ThresholdValue() float64 {
	rv := objc.Call[float64](i_, objc.Sel("thresholdValue"))
	return rv
}
