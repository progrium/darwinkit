// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNDivide] class.
var CNNDivideClass = _CNNDivideClass{objc.GetClass("MPSCNNDivide")}

type _CNNDivideClass struct {
	objc.Class
}

// An interface definition for the [CNNDivide] class.
type ICNNDivide interface {
	ICNNArithmetic
}

// A division operator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndivide?language=objc
type CNNDivide struct {
	CNNArithmetic
}

func CNNDivideFrom(ptr unsafe.Pointer) CNNDivide {
	return CNNDivide{
		CNNArithmetic: CNNArithmeticFrom(ptr),
	}
}

func (c_ CNNDivide) InitWithDevice(device metal.PDevice) CNNDivide {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDivide](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndivide/2942508-initwithdevice?language=objc
func NewCNNDivideWithDevice(device metal.PDevice) CNNDivide {
	instance := CNNDivideClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (cc _CNNDivideClass) Alloc() CNNDivide {
	rv := objc.Call[CNNDivide](cc, objc.Sel("alloc"))
	return rv
}

func CNNDivide_Alloc() CNNDivide {
	return CNNDivideClass.Alloc()
}

func (cc _CNNDivideClass) New() CNNDivide {
	rv := objc.Call[CNNDivide](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNDivide() CNNDivide {
	return CNNDivideClass.New()
}

func (c_ CNNDivide) Init() CNNDivide {
	rv := objc.Call[CNNDivide](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNDivide) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNDivide {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDivide](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNDivide_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNDivide {
	instance := CNNDivideClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
