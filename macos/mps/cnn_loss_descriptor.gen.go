// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNLossDescriptor] class.
var CNNLossDescriptorClass = _CNNLossDescriptorClass{objc.GetClass("MPSCNNLossDescriptor")}

type _CNNLossDescriptorClass struct {
	objc.Class
}

// An interface definition for the [CNNLossDescriptor] class.
type ICNNLossDescriptor interface {
	objc.IObject
	NumberOfClasses() uint
	SetNumberOfClasses(value uint)
	Weight() float64
	SetWeight(value float64)
	Epsilon() float64
	SetEpsilon(value float64)
	Delta() float64
	SetDelta(value float64)
	ReductionType() CNNReductionType
	SetReductionType(value CNNReductionType)
	ReduceAcrossBatch() bool
	SetReduceAcrossBatch(value bool)
	LossType() CNNLossType
	SetLossType(value CNNLossType)
	LabelSmoothing() float64
	SetLabelSmoothing(value float64)
}

// An object that specifies properties used by a loss kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor?language=objc
type CNNLossDescriptor struct {
	objc.Object
}

func CNNLossDescriptorFrom(ptr unsafe.Pointer) CNNLossDescriptor {
	return CNNLossDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CNNLossDescriptorClass) Alloc() CNNLossDescriptor {
	rv := objc.Call[CNNLossDescriptor](cc, objc.Sel("alloc"))
	return rv
}

func CNNLossDescriptor_Alloc() CNNLossDescriptor {
	return CNNLossDescriptorClass.Alloc()
}

func (cc _CNNLossDescriptorClass) New() CNNLossDescriptor {
	rv := objc.Call[CNNLossDescriptor](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNLossDescriptor() CNNLossDescriptor {
	return CNNLossDescriptorClass.New()
}

func (c_ CNNLossDescriptor) Init() CNNLossDescriptor {
	rv := objc.Call[CNNLossDescriptor](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942373-cnnlossdescriptorwithtype?language=objc
func (cc _CNNLossDescriptorClass) CnnLossDescriptorWithTypeReductionType(lossType CNNLossType, reductionType CNNReductionType) CNNLossDescriptor {
	rv := objc.Call[CNNLossDescriptor](cc, objc.Sel("cnnLossDescriptorWithType:reductionType:"), lossType, reductionType)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942373-cnnlossdescriptorwithtype?language=objc
func CNNLossDescriptor_CnnLossDescriptorWithTypeReductionType(lossType CNNLossType, reductionType CNNReductionType) CNNLossDescriptor {
	return CNNLossDescriptorClass.CnnLossDescriptorWithTypeReductionType(lossType, reductionType)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942382-numberofclasses?language=objc
func (c_ CNNLossDescriptor) NumberOfClasses() uint {
	rv := objc.Call[uint](c_, objc.Sel("numberOfClasses"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942382-numberofclasses?language=objc
func (c_ CNNLossDescriptor) SetNumberOfClasses(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setNumberOfClasses:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942367-weight?language=objc
func (c_ CNNLossDescriptor) Weight() float64 {
	rv := objc.Call[float64](c_, objc.Sel("weight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942367-weight?language=objc
func (c_ CNNLossDescriptor) SetWeight(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setWeight:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942362-epsilon?language=objc
func (c_ CNNLossDescriptor) Epsilon() float64 {
	rv := objc.Call[float64](c_, objc.Sel("epsilon"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942362-epsilon?language=objc
func (c_ CNNLossDescriptor) SetEpsilon(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setEpsilon:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942378-delta?language=objc
func (c_ CNNLossDescriptor) Delta() float64 {
	rv := objc.Call[float64](c_, objc.Sel("delta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942378-delta?language=objc
func (c_ CNNLossDescriptor) SetDelta(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setDelta:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942388-reductiontype?language=objc
func (c_ CNNLossDescriptor) ReductionType() CNNReductionType {
	rv := objc.Call[CNNReductionType](c_, objc.Sel("reductionType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942388-reductiontype?language=objc
func (c_ CNNLossDescriptor) SetReductionType(value CNNReductionType) {
	objc.Call[objc.Void](c_, objc.Sel("setReductionType:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/3547982-reduceacrossbatch?language=objc
func (c_ CNNLossDescriptor) ReduceAcrossBatch() bool {
	rv := objc.Call[bool](c_, objc.Sel("reduceAcrossBatch"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/3547982-reduceacrossbatch?language=objc
func (c_ CNNLossDescriptor) SetReduceAcrossBatch(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setReduceAcrossBatch:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942381-losstype?language=objc
func (c_ CNNLossDescriptor) LossType() CNNLossType {
	rv := objc.Call[CNNLossType](c_, objc.Sel("lossType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942381-losstype?language=objc
func (c_ CNNLossDescriptor) SetLossType(value CNNLossType) {
	objc.Call[objc.Void](c_, objc.Sel("setLossType:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942369-labelsmoothing?language=objc
func (c_ CNNLossDescriptor) LabelSmoothing() float64 {
	rv := objc.Call[float64](c_, objc.Sel("labelSmoothing"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossdescriptor/2942369-labelsmoothing?language=objc
func (c_ CNNLossDescriptor) SetLabelSmoothing(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setLabelSmoothing:"), value)
}
