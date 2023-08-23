// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixRandomDistributionDescriptor] class.
var MatrixRandomDistributionDescriptorClass = _MatrixRandomDistributionDescriptorClass{objc.GetClass("MPSMatrixRandomDistributionDescriptor")}

type _MatrixRandomDistributionDescriptorClass struct {
	objc.Class
}

// An interface definition for the [MatrixRandomDistributionDescriptor] class.
type IMatrixRandomDistributionDescriptor interface {
	objc.IObject
	DistributionType() MatrixRandomDistribution
	SetDistributionType(value MatrixRandomDistribution)
	Minimum() float64
	SetMinimum(value float64)
	StandardDeviation() float64
	SetStandardDeviation(value float64)
	Maximum() float64
	SetMaximum(value float64)
	Mean() float64
	SetMean(value float64)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor?language=objc
type MatrixRandomDistributionDescriptor struct {
	objc.Object
}

func MatrixRandomDistributionDescriptorFrom(ptr unsafe.Pointer) MatrixRandomDistributionDescriptor {
	return MatrixRandomDistributionDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MatrixRandomDistributionDescriptorClass) Alloc() MatrixRandomDistributionDescriptor {
	rv := objc.Call[MatrixRandomDistributionDescriptor](mc, objc.Sel("alloc"))
	return rv
}

func MatrixRandomDistributionDescriptor_Alloc() MatrixRandomDistributionDescriptor {
	return MatrixRandomDistributionDescriptorClass.Alloc()
}

func (mc _MatrixRandomDistributionDescriptorClass) New() MatrixRandomDistributionDescriptor {
	rv := objc.Call[MatrixRandomDistributionDescriptor](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixRandomDistributionDescriptor() MatrixRandomDistributionDescriptor {
	return MatrixRandomDistributionDescriptorClass.New()
}

func (m_ MatrixRandomDistributionDescriptor) Init() MatrixRandomDistributionDescriptor {
	rv := objc.Call[MatrixRandomDistributionDescriptor](m_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242862-uniformdistributiondescriptorwit?language=objc
func (mc _MatrixRandomDistributionDescriptorClass) UniformDistributionDescriptorWithMinimumMaximum(minimum float64, maximum float64) MatrixRandomDistributionDescriptor {
	rv := objc.Call[MatrixRandomDistributionDescriptor](mc, objc.Sel("uniformDistributionDescriptorWithMinimum:maximum:"), minimum, maximum)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242862-uniformdistributiondescriptorwit?language=objc
func MatrixRandomDistributionDescriptor_UniformDistributionDescriptorWithMinimumMaximum(minimum float64, maximum float64) MatrixRandomDistributionDescriptor {
	return MatrixRandomDistributionDescriptorClass.UniformDistributionDescriptorWithMinimumMaximum(minimum, maximum)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3547979-normaldistributiondescriptorwith?language=objc
func (mc _MatrixRandomDistributionDescriptorClass) NormalDistributionDescriptorWithMeanStandardDeviation(mean float64, standardDeviation float64) MatrixRandomDistributionDescriptor {
	rv := objc.Call[MatrixRandomDistributionDescriptor](mc, objc.Sel("normalDistributionDescriptorWithMean:standardDeviation:"), mean, standardDeviation)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3547979-normaldistributiondescriptorwith?language=objc
func MatrixRandomDistributionDescriptor_NormalDistributionDescriptorWithMeanStandardDeviation(mean float64, standardDeviation float64) MatrixRandomDistributionDescriptor {
	return MatrixRandomDistributionDescriptorClass.NormalDistributionDescriptorWithMeanStandardDeviation(mean, standardDeviation)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242856-defaultdistributiondescriptor?language=objc
func (mc _MatrixRandomDistributionDescriptorClass) DefaultDistributionDescriptor() MatrixRandomDistributionDescriptor {
	rv := objc.Call[MatrixRandomDistributionDescriptor](mc, objc.Sel("defaultDistributionDescriptor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242856-defaultdistributiondescriptor?language=objc
func MatrixRandomDistributionDescriptor_DefaultDistributionDescriptor() MatrixRandomDistributionDescriptor {
	return MatrixRandomDistributionDescriptorClass.DefaultDistributionDescriptor()
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242857-distributiontype?language=objc
func (m_ MatrixRandomDistributionDescriptor) DistributionType() MatrixRandomDistribution {
	rv := objc.Call[MatrixRandomDistribution](m_, objc.Sel("distributionType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242857-distributiontype?language=objc
func (m_ MatrixRandomDistributionDescriptor) SetDistributionType(value MatrixRandomDistribution) {
	objc.Call[objc.Void](m_, objc.Sel("setDistributionType:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242860-minimum?language=objc
func (m_ MatrixRandomDistributionDescriptor) Minimum() float64 {
	rv := objc.Call[float64](m_, objc.Sel("minimum"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242860-minimum?language=objc
func (m_ MatrixRandomDistributionDescriptor) SetMinimum(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setMinimum:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242861-standarddeviation?language=objc
func (m_ MatrixRandomDistributionDescriptor) StandardDeviation() float64 {
	rv := objc.Call[float64](m_, objc.Sel("standardDeviation"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242861-standarddeviation?language=objc
func (m_ MatrixRandomDistributionDescriptor) SetStandardDeviation(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setStandardDeviation:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242858-maximum?language=objc
func (m_ MatrixRandomDistributionDescriptor) Maximum() float64 {
	rv := objc.Call[float64](m_, objc.Sel("maximum"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242858-maximum?language=objc
func (m_ MatrixRandomDistributionDescriptor) SetMaximum(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setMaximum:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242859-mean?language=objc
func (m_ MatrixRandomDistributionDescriptor) Mean() float64 {
	rv := objc.Call[float64](m_, objc.Sel("mean"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomdistributiondescriptor/3242859-mean?language=objc
func (m_ MatrixRandomDistributionDescriptor) SetMean(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setMean:"), value)
}
