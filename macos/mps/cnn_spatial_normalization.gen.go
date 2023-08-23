// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNSpatialNormalization] class.
var CNNSpatialNormalizationClass = _CNNSpatialNormalizationClass{objc.GetClass("MPSCNNSpatialNormalization")}

type _CNNSpatialNormalizationClass struct {
	objc.Class
}

// An interface definition for the [CNNSpatialNormalization] class.
type ICNNSpatialNormalization interface {
	ICNNKernel
	Beta() float64
	SetBeta(value float64)
	Delta() float64
	SetDelta(value float64)
	Alpha() float64
	SetAlpha(value float64)
}

// A spatial normalization kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization?language=objc
type CNNSpatialNormalization struct {
	CNNKernel
}

func CNNSpatialNormalizationFrom(ptr unsafe.Pointer) CNNSpatialNormalization {
	return CNNSpatialNormalization{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (c_ CNNSpatialNormalization) InitWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNSpatialNormalization {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNSpatialNormalization](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), po0, kernelWidth, kernelHeight)
	return rv
}

// Initializes a spatial normalization kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/1648831-initwithdevice?language=objc
func NewCNNSpatialNormalizationWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNSpatialNormalization {
	instance := CNNSpatialNormalizationClass.Alloc().InitWithDeviceKernelWidthKernelHeight(device, kernelWidth, kernelHeight)
	instance.Autorelease()
	return instance
}

func (cc _CNNSpatialNormalizationClass) Alloc() CNNSpatialNormalization {
	rv := objc.Call[CNNSpatialNormalization](cc, objc.Sel("alloc"))
	return rv
}

func CNNSpatialNormalization_Alloc() CNNSpatialNormalization {
	return CNNSpatialNormalizationClass.Alloc()
}

func (cc _CNNSpatialNormalizationClass) New() CNNSpatialNormalization {
	rv := objc.Call[CNNSpatialNormalization](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNSpatialNormalization() CNNSpatialNormalization {
	return CNNSpatialNormalizationClass.New()
}

func (c_ CNNSpatialNormalization) Init() CNNSpatialNormalization {
	rv := objc.Call[CNNSpatialNormalization](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNSpatialNormalization) InitWithDevice(device metal.PDevice) CNNSpatialNormalization {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNSpatialNormalization](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNSpatialNormalizationWithDevice(device metal.PDevice) CNNSpatialNormalization {
	instance := CNNSpatialNormalizationClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNSpatialNormalization) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNSpatialNormalization {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNSpatialNormalization](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNSpatialNormalization_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNSpatialNormalization {
	instance := CNNSpatialNormalizationClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// The "beta" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/1648936-beta?language=objc
func (c_ CNNSpatialNormalization) Beta() float64 {
	rv := objc.Call[float64](c_, objc.Sel("beta"))
	return rv
}

// The "beta" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/1648936-beta?language=objc
func (c_ CNNSpatialNormalization) SetBeta(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setBeta:"), value)
}

// The "delta" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/1648933-delta?language=objc
func (c_ CNNSpatialNormalization) Delta() float64 {
	rv := objc.Call[float64](c_, objc.Sel("delta"))
	return rv
}

// The "delta" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/1648933-delta?language=objc
func (c_ CNNSpatialNormalization) SetDelta(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setDelta:"), value)
}

// The "alpha" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/1648825-alpha?language=objc
func (c_ CNNSpatialNormalization) Alpha() float64 {
	rv := objc.Call[float64](c_, objc.Sel("alpha"))
	return rv
}

// The "alpha" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnspatialnormalization/1648825-alpha?language=objc
func (c_ CNNSpatialNormalization) SetAlpha(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setAlpha:"), value)
}
