// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNSubtractGradient] class.
var CNNSubtractGradientClass = _CNNSubtractGradientClass{objc.GetClass("MPSCNNSubtractGradient")}

type _CNNSubtractGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNSubtractGradient] class.
type ICNNSubtractGradient interface {
	ICNNArithmeticGradient
}

// A gradient subtraction operator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsubtractgradient?language=objc
type CNNSubtractGradient struct {
	CNNArithmeticGradient
}

func CNNSubtractGradientFrom(ptr unsafe.Pointer) CNNSubtractGradient {
	return CNNSubtractGradient{
		CNNArithmeticGradient: CNNArithmeticGradientFrom(ptr),
	}
}

func (c_ CNNSubtractGradient) InitWithDeviceIsSecondarySourceFilter(device metal.PDevice, isSecondarySourceFilter bool) CNNSubtractGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNSubtractGradient](c_, objc.Sel("initWithDevice:isSecondarySourceFilter:"), po0, isSecondarySourceFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsubtractgradient/2956165-initwithdevice?language=objc
func NewCNNSubtractGradientWithDeviceIsSecondarySourceFilter(device metal.PDevice, isSecondarySourceFilter bool) CNNSubtractGradient {
	instance := CNNSubtractGradientClass.Alloc().InitWithDeviceIsSecondarySourceFilter(device, isSecondarySourceFilter)
	instance.Autorelease()
	return instance
}

func (cc _CNNSubtractGradientClass) Alloc() CNNSubtractGradient {
	rv := objc.Call[CNNSubtractGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNSubtractGradient_Alloc() CNNSubtractGradient {
	return CNNSubtractGradientClass.Alloc()
}

func (cc _CNNSubtractGradientClass) New() CNNSubtractGradient {
	rv := objc.Call[CNNSubtractGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNSubtractGradient() CNNSubtractGradient {
	return CNNSubtractGradientClass.New()
}

func (c_ CNNSubtractGradient) Init() CNNSubtractGradient {
	rv := objc.Call[CNNSubtractGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNSubtractGradient) InitWithDevice(device metal.PDevice) CNNSubtractGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNSubtractGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNSubtractGradientWithDevice(device metal.PDevice) CNNSubtractGradient {
	instance := CNNSubtractGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNSubtractGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNSubtractGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNSubtractGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNSubtractGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNSubtractGradient {
	instance := CNNSubtractGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
