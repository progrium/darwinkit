// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNMultiplyGradient] class.
var CNNMultiplyGradientClass = _CNNMultiplyGradientClass{objc.GetClass("MPSCNNMultiplyGradient")}

type _CNNMultiplyGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNMultiplyGradient] class.
type ICNNMultiplyGradient interface {
	ICNNArithmeticGradient
}

// A gradient multiply operator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiplygradient?language=objc
type CNNMultiplyGradient struct {
	CNNArithmeticGradient
}

func CNNMultiplyGradientFrom(ptr unsafe.Pointer) CNNMultiplyGradient {
	return CNNMultiplyGradient{
		CNNArithmeticGradient: CNNArithmeticGradientFrom(ptr),
	}
}

func (c_ CNNMultiplyGradient) InitWithDeviceIsSecondarySourceFilter(device metal.PDevice, isSecondarySourceFilter bool) CNNMultiplyGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNMultiplyGradient](c_, objc.Sel("initWithDevice:isSecondarySourceFilter:"), po0, isSecondarySourceFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiplygradient/2956164-initwithdevice?language=objc
func NewCNNMultiplyGradientWithDeviceIsSecondarySourceFilter(device metal.PDevice, isSecondarySourceFilter bool) CNNMultiplyGradient {
	instance := CNNMultiplyGradientClass.Alloc().InitWithDeviceIsSecondarySourceFilter(device, isSecondarySourceFilter)
	instance.Autorelease()
	return instance
}

func (cc _CNNMultiplyGradientClass) Alloc() CNNMultiplyGradient {
	rv := objc.Call[CNNMultiplyGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNMultiplyGradient_Alloc() CNNMultiplyGradient {
	return CNNMultiplyGradientClass.Alloc()
}

func (cc _CNNMultiplyGradientClass) New() CNNMultiplyGradient {
	rv := objc.Call[CNNMultiplyGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNMultiplyGradient() CNNMultiplyGradient {
	return CNNMultiplyGradientClass.New()
}

func (c_ CNNMultiplyGradient) Init() CNNMultiplyGradient {
	rv := objc.Call[CNNMultiplyGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNMultiplyGradient) InitWithDevice(device metal.PDevice) CNNMultiplyGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNMultiplyGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNMultiplyGradientWithDevice(device metal.PDevice) CNNMultiplyGradient {
	instance := CNNMultiplyGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNMultiplyGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNMultiplyGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNMultiplyGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNMultiplyGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNMultiplyGradient {
	instance := CNNMultiplyGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
