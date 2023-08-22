// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNOptimizerDescriptor] class.
var NNOptimizerDescriptorClass = _NNOptimizerDescriptorClass{objc.GetClass("MPSNNOptimizerDescriptor")}

type _NNOptimizerDescriptorClass struct {
	objc.Class
}

// An interface definition for the [NNOptimizerDescriptor] class.
type INNOptimizerDescriptor interface {
	objc.IObject
	GradientRescale() float64
	SetGradientRescale(value float64)
	ApplyGradientClipping() bool
	SetApplyGradientClipping(value bool)
	RegularizationScale() float64
	SetRegularizationScale(value float64)
	GradientClipMax() float64
	SetGradientClipMax(value float64)
	RegularizationType() NNRegularizationType
	SetRegularizationType(value NNRegularizationType)
	LearningRate() float64
	SetLearningRate(value float64)
	GradientClipMin() float64
	SetGradientClipMin(value float64)
}

// An object that specifies properties used by an optimizer kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor?language=objc
type NNOptimizerDescriptor struct {
	objc.Object
}

func NNOptimizerDescriptorFrom(ptr unsafe.Pointer) NNOptimizerDescriptor {
	return NNOptimizerDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (n_ NNOptimizerDescriptor) InitWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(learningRate float64, gradientRescale float64, regularizationType NNRegularizationType, regularizationScale float64) NNOptimizerDescriptor {
	rv := objc.Call[NNOptimizerDescriptor](n_, objc.Sel("initWithLearningRate:gradientRescale:regularizationType:regularizationScale:"), learningRate, gradientRescale, regularizationType, regularizationScale)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966727-initwithlearningrate?language=objc
func NewNNOptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(learningRate float64, gradientRescale float64, regularizationType NNRegularizationType, regularizationScale float64) NNOptimizerDescriptor {
	instance := NNOptimizerDescriptorClass.Alloc().InitWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(learningRate, gradientRescale, regularizationType, regularizationScale)
	instance.Autorelease()
	return instance
}

func (nc _NNOptimizerDescriptorClass) OptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(learningRate float64, gradientRescale float64, regularizationType NNRegularizationType, regularizationScale float64) NNOptimizerDescriptor {
	rv := objc.Call[NNOptimizerDescriptor](nc, objc.Sel("optimizerDescriptorWithLearningRate:gradientRescale:regularizationType:regularizationScale:"), learningRate, gradientRescale, regularizationType, regularizationScale)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966730-optimizerdescriptorwithlearningr?language=objc
func NNOptimizerDescriptor_OptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(learningRate float64, gradientRescale float64, regularizationType NNRegularizationType, regularizationScale float64) NNOptimizerDescriptor {
	return NNOptimizerDescriptorClass.OptimizerDescriptorWithLearningRateGradientRescaleRegularizationTypeRegularizationScale(learningRate, gradientRescale, regularizationType, regularizationScale)
}

func (nc _NNOptimizerDescriptorClass) Alloc() NNOptimizerDescriptor {
	rv := objc.Call[NNOptimizerDescriptor](nc, objc.Sel("alloc"))
	return rv
}

func NNOptimizerDescriptor_Alloc() NNOptimizerDescriptor {
	return NNOptimizerDescriptorClass.Alloc()
}

func (nc _NNOptimizerDescriptorClass) New() NNOptimizerDescriptor {
	rv := objc.Call[NNOptimizerDescriptor](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNOptimizerDescriptor() NNOptimizerDescriptor {
	return NNOptimizerDescriptorClass.New()
}

func (n_ NNOptimizerDescriptor) Init() NNOptimizerDescriptor {
	rv := objc.Call[NNOptimizerDescriptor](n_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966725-gradientrescale?language=objc
func (n_ NNOptimizerDescriptor) GradientRescale() float64 {
	rv := objc.Call[float64](n_, objc.Sel("gradientRescale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966725-gradientrescale?language=objc
func (n_ NNOptimizerDescriptor) SetGradientRescale(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setGradientRescale:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966722-applygradientclipping?language=objc
func (n_ NNOptimizerDescriptor) ApplyGradientClipping() bool {
	rv := objc.Call[bool](n_, objc.Sel("applyGradientClipping"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966722-applygradientclipping?language=objc
func (n_ NNOptimizerDescriptor) SetApplyGradientClipping(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setApplyGradientClipping:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966731-regularizationscale?language=objc
func (n_ NNOptimizerDescriptor) RegularizationScale() float64 {
	rv := objc.Call[float64](n_, objc.Sel("regularizationScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966731-regularizationscale?language=objc
func (n_ NNOptimizerDescriptor) SetRegularizationScale(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setRegularizationScale:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966723-gradientclipmax?language=objc
func (n_ NNOptimizerDescriptor) GradientClipMax() float64 {
	rv := objc.Call[float64](n_, objc.Sel("gradientClipMax"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966723-gradientclipmax?language=objc
func (n_ NNOptimizerDescriptor) SetGradientClipMax(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setGradientClipMax:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966732-regularizationtype?language=objc
func (n_ NNOptimizerDescriptor) RegularizationType() NNRegularizationType {
	rv := objc.Call[NNRegularizationType](n_, objc.Sel("regularizationType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966732-regularizationtype?language=objc
func (n_ NNOptimizerDescriptor) SetRegularizationType(value NNRegularizationType) {
	objc.Call[objc.Void](n_, objc.Sel("setRegularizationType:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966728-learningrate?language=objc
func (n_ NNOptimizerDescriptor) LearningRate() float64 {
	rv := objc.Call[float64](n_, objc.Sel("learningRate"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966728-learningrate?language=objc
func (n_ NNOptimizerDescriptor) SetLearningRate(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setLearningRate:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966724-gradientclipmin?language=objc
func (n_ NNOptimizerDescriptor) GradientClipMin() float64 {
	rv := objc.Call[float64](n_, objc.Sel("gradientClipMin"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnoptimizerdescriptor/2966724-gradientclipmin?language=objc
func (n_ NNOptimizerDescriptor) SetGradientClipMin(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setGradientClipMin:"), value)
}
