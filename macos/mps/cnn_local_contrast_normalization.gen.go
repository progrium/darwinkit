// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNLocalContrastNormalization] class.
var CNNLocalContrastNormalizationClass = _CNNLocalContrastNormalizationClass{objc.GetClass("MPSCNNLocalContrastNormalization")}

type _CNNLocalContrastNormalizationClass struct {
	objc.Class
}

// An interface definition for the [CNNLocalContrastNormalization] class.
type ICNNLocalContrastNormalization interface {
	ICNNKernel
	Ps() float64
	SetPs(value float64)
	P0() float64
	SetP0(value float64)
	Beta() float64
	SetBeta(value float64)
	Delta() float64
	SetDelta(value float64)
	Pm() float64
	SetPm(value float64)
	Alpha() float64
	SetAlpha(value float64)
}

// A local-contrast normalization kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization?language=objc
type CNNLocalContrastNormalization struct {
	CNNKernel
}

func CNNLocalContrastNormalizationFrom(ptr unsafe.Pointer) CNNLocalContrastNormalization {
	return CNNLocalContrastNormalization{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (c_ CNNLocalContrastNormalization) InitWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNLocalContrastNormalization {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNLocalContrastNormalization](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), po0, kernelWidth, kernelHeight)
	return rv
}

// Initializes a local contrast normalization kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648924-initwithdevice?language=objc
func NewCNNLocalContrastNormalizationWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNLocalContrastNormalization {
	instance := CNNLocalContrastNormalizationClass.Alloc().InitWithDeviceKernelWidthKernelHeight(device, kernelWidth, kernelHeight)
	instance.Autorelease()
	return instance
}

func (cc _CNNLocalContrastNormalizationClass) Alloc() CNNLocalContrastNormalization {
	rv := objc.Call[CNNLocalContrastNormalization](cc, objc.Sel("alloc"))
	return rv
}

func CNNLocalContrastNormalization_Alloc() CNNLocalContrastNormalization {
	return CNNLocalContrastNormalizationClass.Alloc()
}

func (cc _CNNLocalContrastNormalizationClass) New() CNNLocalContrastNormalization {
	rv := objc.Call[CNNLocalContrastNormalization](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNLocalContrastNormalization() CNNLocalContrastNormalization {
	return CNNLocalContrastNormalizationClass.New()
}

func (c_ CNNLocalContrastNormalization) Init() CNNLocalContrastNormalization {
	rv := objc.Call[CNNLocalContrastNormalization](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNLocalContrastNormalization) InitWithDevice(device metal.PDevice) CNNLocalContrastNormalization {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNLocalContrastNormalization](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNLocalContrastNormalizationWithDevice(device metal.PDevice) CNNLocalContrastNormalization {
	instance := CNNLocalContrastNormalizationClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNLocalContrastNormalization) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNLocalContrastNormalization {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNLocalContrastNormalization](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNLocalContrastNormalization_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNLocalContrastNormalization {
	instance := CNNLocalContrastNormalizationClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// The "ps" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648942-ps?language=objc
func (c_ CNNLocalContrastNormalization) Ps() float64 {
	rv := objc.Call[float64](c_, objc.Sel("ps"))
	return rv
}

// The "ps" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648942-ps?language=objc
func (c_ CNNLocalContrastNormalization) SetPs(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setPs:"), value)
}

// The "p0" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648953-p0?language=objc
func (c_ CNNLocalContrastNormalization) P0() float64 {
	rv := objc.Call[float64](c_, objc.Sel("p0"))
	return rv
}

// The "p0" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648953-p0?language=objc
func (c_ CNNLocalContrastNormalization) SetP0(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setP0:"), value)
}

// The "beta" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648905-beta?language=objc
func (c_ CNNLocalContrastNormalization) Beta() float64 {
	rv := objc.Call[float64](c_, objc.Sel("beta"))
	return rv
}

// The "beta" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648905-beta?language=objc
func (c_ CNNLocalContrastNormalization) SetBeta(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setBeta:"), value)
}

// The "delta" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648812-delta?language=objc
func (c_ CNNLocalContrastNormalization) Delta() float64 {
	rv := objc.Call[float64](c_, objc.Sel("delta"))
	return rv
}

// The "delta" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648812-delta?language=objc
func (c_ CNNLocalContrastNormalization) SetDelta(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setDelta:"), value)
}

// The "pm" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648907-pm?language=objc
func (c_ CNNLocalContrastNormalization) Pm() float64 {
	rv := objc.Call[float64](c_, objc.Sel("pm"))
	return rv
}

// The "pm" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648907-pm?language=objc
func (c_ CNNLocalContrastNormalization) SetPm(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setPm:"), value)
}

// The "alpha" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648923-alpha?language=objc
func (c_ CNNLocalContrastNormalization) Alpha() float64 {
	rv := objc.Call[float64](c_, objc.Sel("alpha"))
	return rv
}

// The "alpha" variable of the kernel function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalization/1648923-alpha?language=objc
func (c_ CNNLocalContrastNormalization) SetAlpha(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setAlpha:"), value)
}
