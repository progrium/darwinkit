// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNSubtract] class.
var CNNSubtractClass = _CNNSubtractClass{objc.GetClass("MPSCNNSubtract")}

type _CNNSubtractClass struct {
	objc.Class
}

// An interface definition for the [CNNSubtract] class.
type ICNNSubtract interface {
	ICNNArithmetic
}

// A subtraction operator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsubtract?language=objc
type CNNSubtract struct {
	CNNArithmetic
}

func CNNSubtractFrom(ptr unsafe.Pointer) CNNSubtract {
	return CNNSubtract{
		CNNArithmetic: CNNArithmeticFrom(ptr),
	}
}

func (c_ CNNSubtract) InitWithDevice(device metal.PDevice) CNNSubtract {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNSubtract](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsubtract/2942503-initwithdevice?language=objc
func NewCNNSubtractWithDevice(device metal.PDevice) CNNSubtract {
	instance := CNNSubtractClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (cc _CNNSubtractClass) Alloc() CNNSubtract {
	rv := objc.Call[CNNSubtract](cc, objc.Sel("alloc"))
	return rv
}

func CNNSubtract_Alloc() CNNSubtract {
	return CNNSubtractClass.Alloc()
}

func (cc _CNNSubtractClass) New() CNNSubtract {
	rv := objc.Call[CNNSubtract](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNSubtract() CNNSubtract {
	return CNNSubtractClass.New()
}

func (c_ CNNSubtract) Init() CNNSubtract {
	rv := objc.Call[CNNSubtract](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNSubtract) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNSubtract {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNSubtract](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNSubtract_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNSubtract {
	instance := CNNSubtractClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
