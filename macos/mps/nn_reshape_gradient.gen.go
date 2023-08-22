// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReshapeGradient] class.
var NNReshapeGradientClass = _NNReshapeGradientClass{objc.GetClass("MPSNNReshapeGradient")}

type _NNReshapeGradientClass struct {
	objc.Class
}

// An interface definition for the [NNReshapeGradient] class.
type INNReshapeGradient interface {
	ICNNGradientKernel
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshapegradient?language=objc
type NNReshapeGradient struct {
	CNNGradientKernel
}

func NNReshapeGradientFrom(ptr unsafe.Pointer) NNReshapeGradient {
	return NNReshapeGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (n_ NNReshapeGradient) InitWithDevice(device metal.PDevice) NNReshapeGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReshapeGradient](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshapegradient/3037439-initwithdevice?language=objc
func NewNNReshapeGradientWithDevice(device metal.PDevice) NNReshapeGradient {
	instance := NNReshapeGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNReshapeGradientClass) Alloc() NNReshapeGradient {
	rv := objc.Call[NNReshapeGradient](nc, objc.Sel("alloc"))
	return rv
}

func NNReshapeGradient_Alloc() NNReshapeGradient {
	return NNReshapeGradientClass.Alloc()
}

func (nc _NNReshapeGradientClass) New() NNReshapeGradient {
	rv := objc.Call[NNReshapeGradient](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReshapeGradient() NNReshapeGradient {
	return NNReshapeGradientClass.New()
}

func (n_ NNReshapeGradient) Init() NNReshapeGradient {
	rv := objc.Call[NNReshapeGradient](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReshapeGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReshapeGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReshapeGradient](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReshapeGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReshapeGradient {
	instance := NNReshapeGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
