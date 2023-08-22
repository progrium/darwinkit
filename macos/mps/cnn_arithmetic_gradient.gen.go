// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNArithmeticGradient] class.
var CNNArithmeticGradientClass = _CNNArithmeticGradientClass{objc.GetClass("MPSCNNArithmeticGradient")}

type _CNNArithmeticGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNArithmeticGradient] class.
type ICNNArithmeticGradient interface {
	ICNNGradientKernel
	SecondaryScale() float64
	SetSecondaryScale(value float64)
	IsSecondarySourceFilter() bool
	MaximumValue() float64
	SetMaximumValue(value float64)
	PrimaryScale() float64
	SetPrimaryScale(value float64)
	MinimumValue() float64
	SetMinimumValue(value float64)
	Bias() float64
	SetBias(value float64)
	SecondaryStrideInFeatureChannels() uint
	SetSecondaryStrideInFeatureChannels(value uint)
}

// The base class for gradient arithmetic operators. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient?language=objc
type CNNArithmeticGradient struct {
	CNNGradientKernel
}

func CNNArithmeticGradientFrom(ptr unsafe.Pointer) CNNArithmeticGradient {
	return CNNArithmeticGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (cc _CNNArithmeticGradientClass) Alloc() CNNArithmeticGradient {
	rv := objc.Call[CNNArithmeticGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNArithmeticGradient_Alloc() CNNArithmeticGradient {
	return CNNArithmeticGradientClass.Alloc()
}

func (cc _CNNArithmeticGradientClass) New() CNNArithmeticGradient {
	rv := objc.Call[CNNArithmeticGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNArithmeticGradient() CNNArithmeticGradient {
	return CNNArithmeticGradientClass.New()
}

func (c_ CNNArithmeticGradient) Init() CNNArithmeticGradient {
	rv := objc.Call[CNNArithmeticGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNArithmeticGradient) InitWithDevice(device metal.PDevice) CNNArithmeticGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNArithmeticGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNArithmeticGradientWithDevice(device metal.PDevice) CNNArithmeticGradient {
	instance := CNNArithmeticGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNArithmeticGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNArithmeticGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNArithmeticGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNArithmeticGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNArithmeticGradient {
	instance := CNNArithmeticGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951856-secondaryscale?language=objc
func (c_ CNNArithmeticGradient) SecondaryScale() float64 {
	rv := objc.Call[float64](c_, objc.Sel("secondaryScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951856-secondaryscale?language=objc
func (c_ CNNArithmeticGradient) SetSecondaryScale(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setSecondaryScale:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951852-issecondarysourcefilter?language=objc
func (c_ CNNArithmeticGradient) IsSecondarySourceFilter() bool {
	rv := objc.Call[bool](c_, objc.Sel("isSecondarySourceFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951854-maximumvalue?language=objc
func (c_ CNNArithmeticGradient) MaximumValue() float64 {
	rv := objc.Call[float64](c_, objc.Sel("maximumValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951854-maximumvalue?language=objc
func (c_ CNNArithmeticGradient) SetMaximumValue(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setMaximumValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951861-primaryscale?language=objc
func (c_ CNNArithmeticGradient) PrimaryScale() float64 {
	rv := objc.Call[float64](c_, objc.Sel("primaryScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951861-primaryscale?language=objc
func (c_ CNNArithmeticGradient) SetPrimaryScale(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setPrimaryScale:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951858-minimumvalue?language=objc
func (c_ CNNArithmeticGradient) MinimumValue() float64 {
	rv := objc.Call[float64](c_, objc.Sel("minimumValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951858-minimumvalue?language=objc
func (c_ CNNArithmeticGradient) SetMinimumValue(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setMinimumValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951855-bias?language=objc
func (c_ CNNArithmeticGradient) Bias() float64 {
	rv := objc.Call[float64](c_, objc.Sel("bias"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951855-bias?language=objc
func (c_ CNNArithmeticGradient) SetBias(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setBias:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951853-secondarystrideinfeaturechannels?language=objc
func (c_ CNNArithmeticGradient) SecondaryStrideInFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("secondaryStrideInFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradient/2951853-secondarystrideinfeaturechannels?language=objc
func (c_ CNNArithmeticGradient) SetSecondaryStrideInFeatureChannels(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setSecondaryStrideInFeatureChannels:"), value)
}
