// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNAdd] class.
var CNNAddClass = _CNNAddClass{objc.GetClass("MPSCNNAdd")}

type _CNNAddClass struct {
	objc.Class
}

// An interface definition for the [CNNAdd] class.
type ICNNAdd interface {
	ICNNArithmetic
}

// An addition operator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnadd?language=objc
type CNNAdd struct {
	CNNArithmetic
}

func CNNAddFrom(ptr unsafe.Pointer) CNNAdd {
	return CNNAdd{
		CNNArithmetic: CNNArithmeticFrom(ptr),
	}
}

func (c_ CNNAdd) InitWithDevice(device metal.PDevice) CNNAdd {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNAdd](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnadd/2942501-initwithdevice?language=objc
func NewCNNAddWithDevice(device metal.PDevice) CNNAdd {
	instance := CNNAddClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (cc _CNNAddClass) Alloc() CNNAdd {
	rv := objc.Call[CNNAdd](cc, objc.Sel("alloc"))
	return rv
}

func CNNAdd_Alloc() CNNAdd {
	return CNNAddClass.Alloc()
}

func (cc _CNNAddClass) New() CNNAdd {
	rv := objc.Call[CNNAdd](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNAdd() CNNAdd {
	return CNNAddClass.New()
}

func (c_ CNNAdd) Init() CNNAdd {
	rv := objc.Call[CNNAdd](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNAdd) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNAdd {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNAdd](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNAdd_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNAdd {
	instance := CNNAddClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
