// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNMultiply] class.
var CNNMultiplyClass = _CNNMultiplyClass{objc.GetClass("MPSCNNMultiply")}

type _CNNMultiplyClass struct {
	objc.Class
}

// An interface definition for the [CNNMultiply] class.
type ICNNMultiply interface {
	ICNNArithmetic
}

// A multiply operator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiply?language=objc
type CNNMultiply struct {
	CNNArithmetic
}

func CNNMultiplyFrom(ptr unsafe.Pointer) CNNMultiply {
	return CNNMultiply{
		CNNArithmetic: CNNArithmeticFrom(ptr),
	}
}

func (c_ CNNMultiply) InitWithDevice(device metal.PDevice) CNNMultiply {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNMultiply](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnmultiply/2942507-initwithdevice?language=objc
func NewCNNMultiplyWithDevice(device metal.PDevice) CNNMultiply {
	instance := CNNMultiplyClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (cc _CNNMultiplyClass) Alloc() CNNMultiply {
	rv := objc.Call[CNNMultiply](cc, objc.Sel("alloc"))
	return rv
}

func CNNMultiply_Alloc() CNNMultiply {
	return CNNMultiplyClass.Alloc()
}

func (cc _CNNMultiplyClass) New() CNNMultiply {
	rv := objc.Call[CNNMultiply](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNMultiply() CNNMultiply {
	return CNNMultiplyClass.New()
}

func (c_ CNNMultiply) Init() CNNMultiply {
	rv := objc.Call[CNNMultiply](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNMultiply) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNMultiply {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNMultiply](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNMultiply_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNMultiply {
	instance := CNNMultiplyClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
