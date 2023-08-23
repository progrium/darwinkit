// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageThresholdBinary] class.
var ImageThresholdBinaryClass = _ImageThresholdBinaryClass{objc.GetClass("MPSImageThresholdBinary")}

type _ImageThresholdBinaryClass struct {
	objc.Class
}

// An interface definition for the [ImageThresholdBinary] class.
type IImageThresholdBinary interface {
	IUnaryImageKernel
	MaximumValue() float64
	Transform() *float64
	ThresholdValue() float64
}

// A filter that returns a specified value for each pixel with a value greater than a specified threshold or 0 otherwise. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary?language=objc
type ImageThresholdBinary struct {
	UnaryImageKernel
}

func ImageThresholdBinaryFrom(ptr unsafe.Pointer) ImageThresholdBinary {
	return ImageThresholdBinary{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageThresholdBinary) InitWithDeviceThresholdValueMaximumValueLinearGrayColorTransform(device metal.PDevice, thresholdValue float64, maximumValue float64, transform *float64) ImageThresholdBinary {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageThresholdBinary](i_, objc.Sel("initWithDevice:thresholdValue:maximumValue:linearGrayColorTransform:"), po0, thresholdValue, maximumValue, transform)
	return rv
}

// Initializes the kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary/1618855-initwithdevice?language=objc
func NewImageThresholdBinaryWithDeviceThresholdValueMaximumValueLinearGrayColorTransform(device metal.PDevice, thresholdValue float64, maximumValue float64, transform *float64) ImageThresholdBinary {
	instance := ImageThresholdBinaryClass.Alloc().InitWithDeviceThresholdValueMaximumValueLinearGrayColorTransform(device, thresholdValue, maximumValue, transform)
	instance.Autorelease()
	return instance
}

func (ic _ImageThresholdBinaryClass) Alloc() ImageThresholdBinary {
	rv := objc.Call[ImageThresholdBinary](ic, objc.Sel("alloc"))
	return rv
}

func ImageThresholdBinary_Alloc() ImageThresholdBinary {
	return ImageThresholdBinaryClass.Alloc()
}

func (ic _ImageThresholdBinaryClass) New() ImageThresholdBinary {
	rv := objc.Call[ImageThresholdBinary](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageThresholdBinary() ImageThresholdBinary {
	return ImageThresholdBinaryClass.New()
}

func (i_ ImageThresholdBinary) Init() ImageThresholdBinary {
	rv := objc.Call[ImageThresholdBinary](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageThresholdBinary) InitWithDevice(device metal.PDevice) ImageThresholdBinary {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageThresholdBinary](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageThresholdBinaryWithDevice(device metal.PDevice) ImageThresholdBinary {
	instance := ImageThresholdBinaryClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageThresholdBinary) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageThresholdBinary {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageThresholdBinary](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageThresholdBinary_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageThresholdBinary {
	instance := ImageThresholdBinaryClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// The maximum value used to initialize the threshold filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary/1618852-maximumvalue?language=objc
func (i_ ImageThresholdBinary) MaximumValue() float64 {
	rv := objc.Call[float64](i_, objc.Sel("maximumValue"))
	return rv
}

// The color transform used to initialize the threshold filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary/1618744-transform?language=objc
func (i_ ImageThresholdBinary) Transform() *float64 {
	rv := objc.Call[*float64](i_, objc.Sel("transform"))
	return rv
}

// The threshold value used to initialize the threshold filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagethresholdbinary/1618851-thresholdvalue?language=objc
func (i_ ImageThresholdBinary) ThresholdValue() float64 {
	rv := objc.Call[float64](i_, objc.Sel("thresholdValue"))
	return rv
}
