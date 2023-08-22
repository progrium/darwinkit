// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNInitialGradient] class.
var NNInitialGradientClass = _NNInitialGradientClass{objc.GetClass("MPSNNInitialGradient")}

type _NNInitialGradientClass struct {
	objc.Class
}

// An interface definition for the [NNInitialGradient] class.
type INNInitialGradient interface {
	ICNNKernel
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnninitialgradient?language=objc
type NNInitialGradient struct {
	CNNKernel
}

func NNInitialGradientFrom(ptr unsafe.Pointer) NNInitialGradient {
	return NNInitialGradient{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (n_ NNInitialGradient) InitWithDevice(device metal.PDevice) NNInitialGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNInitialGradient](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnninitialgradient/3131809-initwithdevice?language=objc
func NewNNInitialGradientWithDevice(device metal.PDevice) NNInitialGradient {
	instance := NNInitialGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNInitialGradientClass) Alloc() NNInitialGradient {
	rv := objc.Call[NNInitialGradient](nc, objc.Sel("alloc"))
	return rv
}

func NNInitialGradient_Alloc() NNInitialGradient {
	return NNInitialGradientClass.Alloc()
}

func (nc _NNInitialGradientClass) New() NNInitialGradient {
	rv := objc.Call[NNInitialGradient](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNInitialGradient() NNInitialGradient {
	return NNInitialGradientClass.New()
}

func (n_ NNInitialGradient) Init() NNInitialGradient {
	rv := objc.Call[NNInitialGradient](n_, objc.Sel("init"))
	return rv
}

func (n_ NNInitialGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNInitialGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNInitialGradient](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNInitialGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNInitialGradient {
	instance := NNInitialGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
