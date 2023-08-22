// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNLogSoftMaxGradient] class.
var CNNLogSoftMaxGradientClass = _CNNLogSoftMaxGradientClass{objc.GetClass("MPSCNNLogSoftMaxGradient")}

type _CNNLogSoftMaxGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNLogSoftMaxGradient] class.
type ICNNLogSoftMaxGradient interface {
	ICNNGradientKernel
}

// A gradient logarithmic softmax filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxgradient?language=objc
type CNNLogSoftMaxGradient struct {
	CNNGradientKernel
}

func CNNLogSoftMaxGradientFrom(ptr unsafe.Pointer) CNNLogSoftMaxGradient {
	return CNNLogSoftMaxGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (c_ CNNLogSoftMaxGradient) InitWithDevice(device metal.PDevice) CNNLogSoftMaxGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNLogSoftMaxGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxgradient/2942622-initwithdevice?language=objc
func NewCNNLogSoftMaxGradientWithDevice(device metal.PDevice) CNNLogSoftMaxGradient {
	instance := CNNLogSoftMaxGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (cc _CNNLogSoftMaxGradientClass) Alloc() CNNLogSoftMaxGradient {
	rv := objc.Call[CNNLogSoftMaxGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNLogSoftMaxGradient_Alloc() CNNLogSoftMaxGradient {
	return CNNLogSoftMaxGradientClass.Alloc()
}

func (cc _CNNLogSoftMaxGradientClass) New() CNNLogSoftMaxGradient {
	rv := objc.Call[CNNLogSoftMaxGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNLogSoftMaxGradient() CNNLogSoftMaxGradient {
	return CNNLogSoftMaxGradientClass.New()
}

func (c_ CNNLogSoftMaxGradient) Init() CNNLogSoftMaxGradient {
	rv := objc.Call[CNNLogSoftMaxGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNLogSoftMaxGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNLogSoftMaxGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNLogSoftMaxGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNLogSoftMaxGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNLogSoftMaxGradient {
	instance := CNNLogSoftMaxGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
