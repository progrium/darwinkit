// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNCrossChannelNormalizationGradient] class.
var CNNCrossChannelNormalizationGradientClass = _CNNCrossChannelNormalizationGradientClass{objc.GetClass("MPSCNNCrossChannelNormalizationGradient")}

type _CNNCrossChannelNormalizationGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNCrossChannelNormalizationGradient] class.
type ICNNCrossChannelNormalizationGradient interface {
	ICNNGradientKernel
	KernelSize() uint
	Beta() float64
	SetBeta(value float64)
	Delta() float64
	SetDelta(value float64)
	Alpha() float64
	SetAlpha(value float64)
}

// A gradient normalization kernel applied across feature channels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient?language=objc
type CNNCrossChannelNormalizationGradient struct {
	CNNGradientKernel
}

func CNNCrossChannelNormalizationGradientFrom(ptr unsafe.Pointer) CNNCrossChannelNormalizationGradient {
	return CNNCrossChannelNormalizationGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (c_ CNNCrossChannelNormalizationGradient) InitWithDeviceKernelSize(device metal.PDevice, kernelSize uint) CNNCrossChannelNormalizationGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNCrossChannelNormalizationGradient](c_, objc.Sel("initWithDevice:kernelSize:"), po0, kernelSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942463-initwithdevice?language=objc
func NewCNNCrossChannelNormalizationGradientWithDeviceKernelSize(device metal.PDevice, kernelSize uint) CNNCrossChannelNormalizationGradient {
	instance := CNNCrossChannelNormalizationGradientClass.Alloc().InitWithDeviceKernelSize(device, kernelSize)
	instance.Autorelease()
	return instance
}

func (cc _CNNCrossChannelNormalizationGradientClass) Alloc() CNNCrossChannelNormalizationGradient {
	rv := objc.Call[CNNCrossChannelNormalizationGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNCrossChannelNormalizationGradient_Alloc() CNNCrossChannelNormalizationGradient {
	return CNNCrossChannelNormalizationGradientClass.Alloc()
}

func (cc _CNNCrossChannelNormalizationGradientClass) New() CNNCrossChannelNormalizationGradient {
	rv := objc.Call[CNNCrossChannelNormalizationGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNCrossChannelNormalizationGradient() CNNCrossChannelNormalizationGradient {
	return CNNCrossChannelNormalizationGradientClass.New()
}

func (c_ CNNCrossChannelNormalizationGradient) Init() CNNCrossChannelNormalizationGradient {
	rv := objc.Call[CNNCrossChannelNormalizationGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNCrossChannelNormalizationGradient) InitWithDevice(device metal.PDevice) CNNCrossChannelNormalizationGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNCrossChannelNormalizationGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNCrossChannelNormalizationGradientWithDevice(device metal.PDevice) CNNCrossChannelNormalizationGradient {
	instance := CNNCrossChannelNormalizationGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNCrossChannelNormalizationGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNCrossChannelNormalizationGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNCrossChannelNormalizationGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNCrossChannelNormalizationGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNCrossChannelNormalizationGradient {
	instance := CNNCrossChannelNormalizationGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942468-kernelsize?language=objc
func (c_ CNNCrossChannelNormalizationGradient) KernelSize() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelSize"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942477-beta?language=objc
func (c_ CNNCrossChannelNormalizationGradient) Beta() float64 {
	rv := objc.Call[float64](c_, objc.Sel("beta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942477-beta?language=objc
func (c_ CNNCrossChannelNormalizationGradient) SetBeta(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setBeta:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942465-delta?language=objc
func (c_ CNNCrossChannelNormalizationGradient) Delta() float64 {
	rv := objc.Call[float64](c_, objc.Sel("delta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942465-delta?language=objc
func (c_ CNNCrossChannelNormalizationGradient) SetDelta(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setDelta:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942464-alpha?language=objc
func (c_ CNNCrossChannelNormalizationGradient) Alpha() float64 {
	rv := objc.Call[float64](c_, objc.Sel("alpha"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradient/2942464-alpha?language=objc
func (c_ CNNCrossChannelNormalizationGradient) SetAlpha(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setAlpha:"), value)
}
