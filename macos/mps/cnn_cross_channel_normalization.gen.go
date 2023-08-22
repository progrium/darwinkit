// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNCrossChannelNormalization] class.
var CNNCrossChannelNormalizationClass = _CNNCrossChannelNormalizationClass{objc.GetClass("MPSCNNCrossChannelNormalization")}

type _CNNCrossChannelNormalizationClass struct {
	objc.Class
}

// An interface definition for the [CNNCrossChannelNormalization] class.
type ICNNCrossChannelNormalization interface {
	ICNNKernel
	KernelSize() uint
	Beta() float64
	SetBeta(value float64)
	Delta() float64
	SetDelta(value float64)
	Alpha() float64
	SetAlpha(value float64)
}

// A normalization kernel applied across feature channels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalization?language=objc
type CNNCrossChannelNormalization struct {
	CNNKernel
}

func CNNCrossChannelNormalizationFrom(ptr unsafe.Pointer) CNNCrossChannelNormalization {
	return CNNCrossChannelNormalization{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (c_ CNNCrossChannelNormalization) InitWithDeviceKernelSize(device metal.PDevice, kernelSize uint) CNNCrossChannelNormalization {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNCrossChannelNormalization](c_, objc.Sel("initWithDevice:kernelSize:"), po0, kernelSize)
	return rv
}

// Initializes a normalization kernel in a channel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalization/1648834-initwithdevice?language=objc
func NewCNNCrossChannelNormalizationWithDeviceKernelSize(device metal.PDevice, kernelSize uint) CNNCrossChannelNormalization {
	instance := CNNCrossChannelNormalizationClass.Alloc().InitWithDeviceKernelSize(device, kernelSize)
	instance.Autorelease()
	return instance
}

func (cc _CNNCrossChannelNormalizationClass) Alloc() CNNCrossChannelNormalization {
	rv := objc.Call[CNNCrossChannelNormalization](cc, objc.Sel("alloc"))
	return rv
}

func CNNCrossChannelNormalization_Alloc() CNNCrossChannelNormalization {
	return CNNCrossChannelNormalizationClass.Alloc()
}

func (cc _CNNCrossChannelNormalizationClass) New() CNNCrossChannelNormalization {
	rv := objc.Call[CNNCrossChannelNormalization](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNCrossChannelNormalization() CNNCrossChannelNormalization {
	return CNNCrossChannelNormalizationClass.New()
}

func (c_ CNNCrossChannelNormalization) Init() CNNCrossChannelNormalization {
	rv := objc.Call[CNNCrossChannelNormalization](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNCrossChannelNormalization) InitWithDevice(device metal.PDevice) CNNCrossChannelNormalization {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNCrossChannelNormalization](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNCrossChannelNormalizationWithDevice(device metal.PDevice) CNNCrossChannelNormalization {
	instance := CNNCrossChannelNormalizationClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNCrossChannelNormalization) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNCrossChannelNormalization {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNCrossChannelNormalization](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNCrossChannelNormalization_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNCrossChannelNormalization {
	instance := CNNCrossChannelNormalizationClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// The size of the square kernel window. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalization/1648811-kernelsize?language=objc
func (c_ CNNCrossChannelNormalization) KernelSize() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelSize"))
	return rv
}

// The "beta" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalization/1648879-beta?language=objc
func (c_ CNNCrossChannelNormalization) Beta() float64 {
	rv := objc.Call[float64](c_, objc.Sel("beta"))
	return rv
}

// The "beta" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalization/1648879-beta?language=objc
func (c_ CNNCrossChannelNormalization) SetBeta(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setBeta:"), value)
}

// The "delta" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalization/1648881-delta?language=objc
func (c_ CNNCrossChannelNormalization) Delta() float64 {
	rv := objc.Call[float64](c_, objc.Sel("delta"))
	return rv
}

// The "delta" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalization/1648881-delta?language=objc
func (c_ CNNCrossChannelNormalization) SetDelta(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setDelta:"), value)
}

// The "alpha" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalization/1648896-alpha?language=objc
func (c_ CNNCrossChannelNormalization) Alpha() float64 {
	rv := objc.Call[float64](c_, objc.Sel("alpha"))
	return rv
}

// The "alpha" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalization/1648896-alpha?language=objc
func (c_ CNNCrossChannelNormalization) SetAlpha(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setAlpha:"), value)
}
