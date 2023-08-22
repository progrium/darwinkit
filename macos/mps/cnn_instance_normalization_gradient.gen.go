// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNInstanceNormalizationGradient] class.
var CNNInstanceNormalizationGradientClass = _CNNInstanceNormalizationGradientClass{objc.GetClass("MPSCNNInstanceNormalizationGradient")}

type _CNNInstanceNormalizationGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNInstanceNormalizationGradient] class.
type ICNNInstanceNormalizationGradient interface {
	ICNNGradientKernel
}

// A gradient instance normalization kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradient?language=objc
type CNNInstanceNormalizationGradient struct {
	CNNGradientKernel
}

func CNNInstanceNormalizationGradientFrom(ptr unsafe.Pointer) CNNInstanceNormalizationGradient {
	return CNNInstanceNormalizationGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (cc _CNNInstanceNormalizationGradientClass) Alloc() CNNInstanceNormalizationGradient {
	rv := objc.Call[CNNInstanceNormalizationGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNInstanceNormalizationGradient_Alloc() CNNInstanceNormalizationGradient {
	return CNNInstanceNormalizationGradientClass.Alloc()
}

func (cc _CNNInstanceNormalizationGradientClass) New() CNNInstanceNormalizationGradient {
	rv := objc.Call[CNNInstanceNormalizationGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNInstanceNormalizationGradient() CNNInstanceNormalizationGradient {
	return CNNInstanceNormalizationGradientClass.New()
}

func (c_ CNNInstanceNormalizationGradient) Init() CNNInstanceNormalizationGradient {
	rv := objc.Call[CNNInstanceNormalizationGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNInstanceNormalizationGradient) InitWithDevice(device metal.PDevice) CNNInstanceNormalizationGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNInstanceNormalizationGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNInstanceNormalizationGradientWithDevice(device metal.PDevice) CNNInstanceNormalizationGradient {
	instance := CNNInstanceNormalizationGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNInstanceNormalizationGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNInstanceNormalizationGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNInstanceNormalizationGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNInstanceNormalizationGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNInstanceNormalizationGradient {
	instance := CNNInstanceNormalizationGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
