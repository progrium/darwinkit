// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNSoftMax] class.
var CNNSoftMaxClass = _CNNSoftMaxClass{objc.GetClass("MPSCNNSoftMax")}

type _CNNSoftMaxClass struct {
	objc.Class
}

// An interface definition for the [CNNSoftMax] class.
type ICNNSoftMax interface {
	ICNNKernel
}

// A neural transfer function that is useful for classification tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmax?language=objc
type CNNSoftMax struct {
	CNNKernel
}

func CNNSoftMaxFrom(ptr unsafe.Pointer) CNNSoftMax {
	return CNNSoftMax{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (cc _CNNSoftMaxClass) Alloc() CNNSoftMax {
	rv := objc.Call[CNNSoftMax](cc, objc.Sel("alloc"))
	return rv
}

func CNNSoftMax_Alloc() CNNSoftMax {
	return CNNSoftMaxClass.Alloc()
}

func (cc _CNNSoftMaxClass) New() CNNSoftMax {
	rv := objc.Call[CNNSoftMax](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNSoftMax() CNNSoftMax {
	return CNNSoftMaxClass.New()
}

func (c_ CNNSoftMax) Init() CNNSoftMax {
	rv := objc.Call[CNNSoftMax](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNSoftMax) InitWithDevice(device metal.PDevice) CNNSoftMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNSoftMax](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNSoftMaxWithDevice(device metal.PDevice) CNNSoftMax {
	instance := CNNSoftMaxClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNSoftMax) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNSoftMax {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNSoftMax](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNSoftMax_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNSoftMax {
	instance := CNNSoftMaxClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
