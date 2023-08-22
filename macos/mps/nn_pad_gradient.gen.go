// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNPadGradient] class.
var NNPadGradientClass = _NNPadGradientClass{objc.GetClass("MPSNNPadGradient")}

type _NNPadGradientClass struct {
	objc.Class
}

// An interface definition for the [NNPadGradient] class.
type INNPadGradient interface {
	ICNNGradientKernel
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadgradient?language=objc
type NNPadGradient struct {
	CNNGradientKernel
}

func NNPadGradientFrom(ptr unsafe.Pointer) NNPadGradient {
	return NNPadGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (n_ NNPadGradient) InitWithDevice(device metal.PDevice) NNPadGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNPadGradient](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadgradient/3037436-initwithdevice?language=objc
func NewNNPadGradientWithDevice(device metal.PDevice) NNPadGradient {
	instance := NNPadGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNPadGradientClass) Alloc() NNPadGradient {
	rv := objc.Call[NNPadGradient](nc, objc.Sel("alloc"))
	return rv
}

func NNPadGradient_Alloc() NNPadGradient {
	return NNPadGradientClass.Alloc()
}

func (nc _NNPadGradientClass) New() NNPadGradient {
	rv := objc.Call[NNPadGradient](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNPadGradient() NNPadGradient {
	return NNPadGradientClass.New()
}

func (n_ NNPadGradient) Init() NNPadGradient {
	rv := objc.Call[NNPadGradient](n_, objc.Sel("init"))
	return rv
}

func (n_ NNPadGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNPadGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNPadGradient](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNPadGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNPadGradient {
	instance := NNPadGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
