// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNDropoutGradient] class.
var CNNDropoutGradientClass = _CNNDropoutGradientClass{objc.GetClass("MPSCNNDropoutGradient")}

type _CNNDropoutGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNDropoutGradient] class.
type ICNNDropoutGradient interface {
	ICNNGradientKernel
	KeepProbability() float64
	MaskStrideInPixels() metal.Size
	Seed() uint
}

// A gradient dropout filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradient?language=objc
type CNNDropoutGradient struct {
	CNNGradientKernel
}

func CNNDropoutGradientFrom(ptr unsafe.Pointer) CNNDropoutGradient {
	return CNNDropoutGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (c_ CNNDropoutGradient) InitWithDeviceKeepProbabilitySeedMaskStrideInPixels(device metal.PDevice, keepProbability float64, seed uint, maskStrideInPixels metal.Size) CNNDropoutGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDropoutGradient](c_, objc.Sel("initWithDevice:keepProbability:seed:maskStrideInPixels:"), po0, keepProbability, seed, maskStrideInPixels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradient/2942518-initwithdevice?language=objc
func NewCNNDropoutGradientWithDeviceKeepProbabilitySeedMaskStrideInPixels(device metal.PDevice, keepProbability float64, seed uint, maskStrideInPixels metal.Size) CNNDropoutGradient {
	instance := CNNDropoutGradientClass.Alloc().InitWithDeviceKeepProbabilitySeedMaskStrideInPixels(device, keepProbability, seed, maskStrideInPixels)
	instance.Autorelease()
	return instance
}

func (cc _CNNDropoutGradientClass) Alloc() CNNDropoutGradient {
	rv := objc.Call[CNNDropoutGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNDropoutGradient_Alloc() CNNDropoutGradient {
	return CNNDropoutGradientClass.Alloc()
}

func (cc _CNNDropoutGradientClass) New() CNNDropoutGradient {
	rv := objc.Call[CNNDropoutGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNDropoutGradient() CNNDropoutGradient {
	return CNNDropoutGradientClass.New()
}

func (c_ CNNDropoutGradient) Init() CNNDropoutGradient {
	rv := objc.Call[CNNDropoutGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNDropoutGradient) InitWithDevice(device metal.PDevice) CNNDropoutGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDropoutGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNDropoutGradientWithDevice(device metal.PDevice) CNNDropoutGradient {
	instance := CNNDropoutGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNDropoutGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNDropoutGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDropoutGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNDropoutGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNDropoutGradient {
	instance := CNNDropoutGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradient/2942520-keepprobability?language=objc
func (c_ CNNDropoutGradient) KeepProbability() float64 {
	rv := objc.Call[float64](c_, objc.Sel("keepProbability"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradient/2942515-maskstrideinpixels?language=objc
func (c_ CNNDropoutGradient) MaskStrideInPixels() metal.Size {
	rv := objc.Call[metal.Size](c_, objc.Sel("maskStrideInPixels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradient/2942528-seed?language=objc
func (c_ CNNDropoutGradient) Seed() uint {
	rv := objc.Call[uint](c_, objc.Sel("seed"))
	return rv
}
