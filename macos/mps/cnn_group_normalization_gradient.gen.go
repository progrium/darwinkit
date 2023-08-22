// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNGroupNormalizationGradient] class.
var CNNGroupNormalizationGradientClass = _CNNGroupNormalizationGradientClass{objc.GetClass("MPSCNNGroupNormalizationGradient")}

type _CNNGroupNormalizationGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNGroupNormalizationGradient] class.
type ICNNGroupNormalizationGradient interface {
	ICNNGradientKernel
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradient?language=objc
type CNNGroupNormalizationGradient struct {
	CNNGradientKernel
}

func CNNGroupNormalizationGradientFrom(ptr unsafe.Pointer) CNNGroupNormalizationGradient {
	return CNNGroupNormalizationGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (cc _CNNGroupNormalizationGradientClass) Alloc() CNNGroupNormalizationGradient {
	rv := objc.Call[CNNGroupNormalizationGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNGroupNormalizationGradient_Alloc() CNNGroupNormalizationGradient {
	return CNNGroupNormalizationGradientClass.Alloc()
}

func (cc _CNNGroupNormalizationGradientClass) New() CNNGroupNormalizationGradient {
	rv := objc.Call[CNNGroupNormalizationGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNGroupNormalizationGradient() CNNGroupNormalizationGradient {
	return CNNGroupNormalizationGradientClass.New()
}

func (c_ CNNGroupNormalizationGradient) Init() CNNGroupNormalizationGradient {
	rv := objc.Call[CNNGroupNormalizationGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNGroupNormalizationGradient) InitWithDevice(device metal.PDevice) CNNGroupNormalizationGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNGroupNormalizationGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNGroupNormalizationGradientWithDevice(device metal.PDevice) CNNGroupNormalizationGradient {
	instance := CNNGroupNormalizationGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNGroupNormalizationGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNGroupNormalizationGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNGroupNormalizationGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNGroupNormalizationGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNGroupNormalizationGradient {
	instance := CNNGroupNormalizationGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
