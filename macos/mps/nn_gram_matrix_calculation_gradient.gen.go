// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNGramMatrixCalculationGradient] class.
var NNGramMatrixCalculationGradientClass = _NNGramMatrixCalculationGradientClass{objc.GetClass("MPSNNGramMatrixCalculationGradient")}

type _NNGramMatrixCalculationGradientClass struct {
	objc.Class
}

// An interface definition for the [NNGramMatrixCalculationGradient] class.
type INNGramMatrixCalculationGradient interface {
	ICNNGradientKernel
	Alpha() float64
	SetAlpha(value float64)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradient?language=objc
type NNGramMatrixCalculationGradient struct {
	CNNGradientKernel
}

func NNGramMatrixCalculationGradientFrom(ptr unsafe.Pointer) NNGramMatrixCalculationGradient {
	return NNGramMatrixCalculationGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (n_ NNGramMatrixCalculationGradient) InitWithDevice(device metal.PDevice) NNGramMatrixCalculationGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNGramMatrixCalculationGradient](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradient/3114084-initwithdevice?language=objc
func NewNNGramMatrixCalculationGradientWithDevice(device metal.PDevice) NNGramMatrixCalculationGradient {
	instance := NNGramMatrixCalculationGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNGramMatrixCalculationGradientClass) Alloc() NNGramMatrixCalculationGradient {
	rv := objc.Call[NNGramMatrixCalculationGradient](nc, objc.Sel("alloc"))
	return rv
}

func NNGramMatrixCalculationGradient_Alloc() NNGramMatrixCalculationGradient {
	return NNGramMatrixCalculationGradientClass.Alloc()
}

func (nc _NNGramMatrixCalculationGradientClass) New() NNGramMatrixCalculationGradient {
	rv := objc.Call[NNGramMatrixCalculationGradient](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNGramMatrixCalculationGradient() NNGramMatrixCalculationGradient {
	return NNGramMatrixCalculationGradientClass.New()
}

func (n_ NNGramMatrixCalculationGradient) Init() NNGramMatrixCalculationGradient {
	rv := objc.Call[NNGramMatrixCalculationGradient](n_, objc.Sel("init"))
	return rv
}

func (n_ NNGramMatrixCalculationGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNGramMatrixCalculationGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNGramMatrixCalculationGradient](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNGramMatrixCalculationGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNGramMatrixCalculationGradient {
	instance := NNGramMatrixCalculationGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradient/3114082-alpha?language=objc
func (n_ NNGramMatrixCalculationGradient) Alpha() float64 {
	rv := objc.Call[float64](n_, objc.Sel("alpha"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradient/3114082-alpha?language=objc
func (n_ NNGramMatrixCalculationGradient) SetAlpha(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setAlpha:"), value)
}
