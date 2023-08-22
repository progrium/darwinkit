// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNOptimizer] class.
var NNOptimizerClass = _NNOptimizerClass{objc.GetClass("MPSNNOptimizer")}

type _NNOptimizerClass struct {
	objc.Class
}

// An interface definition for the [NNOptimizer] class.
type INNOptimizer interface {
	IKernel
	SetLearningRate(newLearningRate float64)
	GradientRescale() float64
	ApplyGradientClipping() bool
	SetApplyGradientClipping(value bool)
	RegularizationScale() float64
	GradientClipMax() float64
	RegularizationType() NNRegularizationType
	LearningRate() float64
	GradientClipMin() float64
}

// The base class for optimization layers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer?language=objc
type NNOptimizer struct {
	Kernel
}

func NNOptimizerFrom(ptr unsafe.Pointer) NNOptimizer {
	return NNOptimizer{
		Kernel: KernelFrom(ptr),
	}
}

func (nc _NNOptimizerClass) Alloc() NNOptimizer {
	rv := objc.Call[NNOptimizer](nc, objc.Sel("alloc"))
	return rv
}

func NNOptimizer_Alloc() NNOptimizer {
	return NNOptimizerClass.Alloc()
}

func (nc _NNOptimizerClass) New() NNOptimizer {
	rv := objc.Call[NNOptimizer](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNOptimizer() NNOptimizer {
	return NNOptimizerClass.New()
}

func (n_ NNOptimizer) Init() NNOptimizer {
	rv := objc.Call[NNOptimizer](n_, objc.Sel("init"))
	return rv
}

func (n_ NNOptimizer) InitWithDevice(device metal.PDevice) NNOptimizer {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizer](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewNNOptimizerWithDevice(device metal.PDevice) NNOptimizer {
	instance := NNOptimizerClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (n_ NNOptimizer) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNOptimizer {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNOptimizer](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNOptimizer_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNOptimizer {
	instance := NNOptimizerClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966712-setlearningrate?language=objc
func (n_ NNOptimizer) SetLearningRate(newLearningRate float64) {
	objc.Call[objc.Void](n_, objc.Sel("setLearningRate:"), newLearningRate)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966708-gradientrescale?language=objc
func (n_ NNOptimizer) GradientRescale() float64 {
	rv := objc.Call[float64](n_, objc.Sel("gradientRescale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966705-applygradientclipping?language=objc
func (n_ NNOptimizer) ApplyGradientClipping() bool {
	rv := objc.Call[bool](n_, objc.Sel("applyGradientClipping"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966705-applygradientclipping?language=objc
func (n_ NNOptimizer) SetApplyGradientClipping(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setApplyGradientClipping:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966710-regularizationscale?language=objc
func (n_ NNOptimizer) RegularizationScale() float64 {
	rv := objc.Call[float64](n_, objc.Sel("regularizationScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966706-gradientclipmax?language=objc
func (n_ NNOptimizer) GradientClipMax() float64 {
	rv := objc.Call[float64](n_, objc.Sel("gradientClipMax"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966711-regularizationtype?language=objc
func (n_ NNOptimizer) RegularizationType() NNRegularizationType {
	rv := objc.Call[NNRegularizationType](n_, objc.Sel("regularizationType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966709-learningrate?language=objc
func (n_ NNOptimizer) LearningRate() float64 {
	rv := objc.Call[float64](n_, objc.Sel("learningRate"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizer/2966707-gradientclipmin?language=objc
func (n_ NNOptimizer) GradientClipMin() float64 {
	rv := objc.Call[float64](n_, objc.Sel("gradientClipMin"))
	return rv
}
