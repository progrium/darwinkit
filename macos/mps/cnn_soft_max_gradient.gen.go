// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNSoftMaxGradient] class.
var CNNSoftMaxGradientClass = _CNNSoftMaxGradientClass{objc.GetClass("MPSCNNSoftMaxGradient")}

type _CNNSoftMaxGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNSoftMaxGradient] class.
type ICNNSoftMaxGradient interface {
	ICNNGradientKernel
}

// A gradient softmax filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmaxgradient?language=objc
type CNNSoftMaxGradient struct {
	CNNGradientKernel
}

func CNNSoftMaxGradientFrom(ptr unsafe.Pointer) CNNSoftMaxGradient {
	return CNNSoftMaxGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (c_ CNNSoftMaxGradient) InitWithDevice(device metal.PDevice) CNNSoftMaxGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNSoftMaxGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmaxgradient/2942620-initwithdevice?language=objc
func NewCNNSoftMaxGradientWithDevice(device metal.PDevice) CNNSoftMaxGradient {
	instance := CNNSoftMaxGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (cc _CNNSoftMaxGradientClass) Alloc() CNNSoftMaxGradient {
	rv := objc.Call[CNNSoftMaxGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNSoftMaxGradient_Alloc() CNNSoftMaxGradient {
	return CNNSoftMaxGradientClass.Alloc()
}

func (cc _CNNSoftMaxGradientClass) New() CNNSoftMaxGradient {
	rv := objc.Call[CNNSoftMaxGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNSoftMaxGradient() CNNSoftMaxGradient {
	return CNNSoftMaxGradientClass.New()
}

func (c_ CNNSoftMaxGradient) Init() CNNSoftMaxGradient {
	rv := objc.Call[CNNSoftMaxGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNSoftMaxGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNSoftMaxGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNSoftMaxGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNSoftMaxGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNSoftMaxGradient {
	instance := CNNSoftMaxGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
