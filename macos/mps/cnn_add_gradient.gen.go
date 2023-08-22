// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNAddGradient] class.
var CNNAddGradientClass = _CNNAddGradientClass{objc.GetClass("MPSCNNAddGradient")}

type _CNNAddGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNAddGradient] class.
type ICNNAddGradient interface {
	ICNNArithmeticGradient
}

// A gradient addition operator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnaddgradient?language=objc
type CNNAddGradient struct {
	CNNArithmeticGradient
}

func CNNAddGradientFrom(ptr unsafe.Pointer) CNNAddGradient {
	return CNNAddGradient{
		CNNArithmeticGradient: CNNArithmeticGradientFrom(ptr),
	}
}

func (c_ CNNAddGradient) InitWithDeviceIsSecondarySourceFilter(device metal.PDevice, isSecondarySourceFilter bool) CNNAddGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNAddGradient](c_, objc.Sel("initWithDevice:isSecondarySourceFilter:"), po0, isSecondarySourceFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnaddgradient/2956163-initwithdevice?language=objc
func NewCNNAddGradientWithDeviceIsSecondarySourceFilter(device metal.PDevice, isSecondarySourceFilter bool) CNNAddGradient {
	instance := CNNAddGradientClass.Alloc().InitWithDeviceIsSecondarySourceFilter(device, isSecondarySourceFilter)
	instance.Autorelease()
	return instance
}

func (cc _CNNAddGradientClass) Alloc() CNNAddGradient {
	rv := objc.Call[CNNAddGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNAddGradient_Alloc() CNNAddGradient {
	return CNNAddGradientClass.Alloc()
}

func (cc _CNNAddGradientClass) New() CNNAddGradient {
	rv := objc.Call[CNNAddGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNAddGradient() CNNAddGradient {
	return CNNAddGradientClass.New()
}

func (c_ CNNAddGradient) Init() CNNAddGradient {
	rv := objc.Call[CNNAddGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNAddGradient) InitWithDevice(device metal.PDevice) CNNAddGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNAddGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNAddGradientWithDevice(device metal.PDevice) CNNAddGradient {
	instance := CNNAddGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNAddGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNAddGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNAddGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNAddGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNAddGradient {
	instance := CNNAddGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
