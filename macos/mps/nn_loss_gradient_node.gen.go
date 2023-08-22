// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNLossGradientNode] class.
var NNLossGradientNodeClass = _NNLossGradientNodeClass{objc.GetClass("MPSNNLossGradientNode")}

type _NNLossGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [NNLossGradientNode] class.
type INNLossGradientNode interface {
	INNGradientFilterNode
	NumberOfClasses() uint
	Weight() float64
	IsLabelsGradientFilter() bool
	PropertyCallBack() NNLossCallbackWrapper
	SetPropertyCallBack(value PNNLossCallback)
	SetPropertyCallBackObject(valueObject objc.IObject)
	Epsilon() float64
	Delta() float64
	ReductionType() CNNReductionType
	ReduceAcrossBatch() bool
	LossType() CNNLossType
	LabelSmoothing() float64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode?language=objc
type NNLossGradientNode struct {
	NNGradientFilterNode
}

func NNLossGradientNodeFrom(ptr unsafe.Pointer) NNLossGradientNode {
	return NNLossGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (nc _NNLossGradientNodeClass) NodeWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sourceNodes []INNImageNode, gradientState INNGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) NNLossGradientNode {
	rv := objc.Call[NNLossGradientNode](nc, objc.Sel("nodeWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceNodes, objc.Ptr(gradientState), objc.Ptr(descriptor), isLabelsGradientFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131863-nodewithsources?language=objc
func NNLossGradientNode_NodeWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sourceNodes []INNImageNode, gradientState INNGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) NNLossGradientNode {
	return NNLossGradientNodeClass.NodeWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sourceNodes, gradientState, descriptor, isLabelsGradientFilter)
}

func (n_ NNLossGradientNode) InitWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient INNImageNode, sourceImage INNImageNode, labels INNImageNode, gradientState INNGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) NNLossGradientNode {
	rv := objc.Call[NNLossGradientNode](n_, objc.Sel("initWithSourceGradient:sourceImage:labels:gradientState:lossDescriptor:isLabelsGradientFilter:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(labels), objc.Ptr(gradientState), objc.Ptr(descriptor), isLabelsGradientFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131855-initwithsourcegradient?language=objc
func NewNNLossGradientNodeWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient INNImageNode, sourceImage INNImageNode, labels INNImageNode, gradientState INNGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) NNLossGradientNode {
	instance := NNLossGradientNodeClass.Alloc().InitWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient, sourceImage, labels, gradientState, descriptor, isLabelsGradientFilter)
	instance.Autorelease()
	return instance
}

func (nc _NNLossGradientNodeClass) NodeWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient INNImageNode, sourceImage INNImageNode, labels INNImageNode, gradientState INNGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) NNLossGradientNode {
	rv := objc.Call[NNLossGradientNode](nc, objc.Sel("nodeWithSourceGradient:sourceImage:labels:gradientState:lossDescriptor:isLabelsGradientFilter:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(labels), objc.Ptr(gradientState), objc.Ptr(descriptor), isLabelsGradientFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131861-nodewithsourcegradient?language=objc
func NNLossGradientNode_NodeWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient INNImageNode, sourceImage INNImageNode, labels INNImageNode, gradientState INNGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) NNLossGradientNode {
	return NNLossGradientNodeClass.NodeWithSourceGradientSourceImageLabelsGradientStateLossDescriptorIsLabelsGradientFilter(sourceGradient, sourceImage, labels, gradientState, descriptor, isLabelsGradientFilter)
}

func (n_ NNLossGradientNode) InitWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sourceNodes []INNImageNode, gradientState INNGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) NNLossGradientNode {
	rv := objc.Call[NNLossGradientNode](n_, objc.Sel("initWithSources:gradientState:lossDescriptor:isLabelsGradientFilter:"), sourceNodes, objc.Ptr(gradientState), objc.Ptr(descriptor), isLabelsGradientFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131857-initwithsources?language=objc
func NewNNLossGradientNodeWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sourceNodes []INNImageNode, gradientState INNGradientStateNode, descriptor ICNNLossDescriptor, isLabelsGradientFilter bool) NNLossGradientNode {
	instance := NNLossGradientNodeClass.Alloc().InitWithSourcesGradientStateLossDescriptorIsLabelsGradientFilter(sourceNodes, gradientState, descriptor, isLabelsGradientFilter)
	instance.Autorelease()
	return instance
}

func (nc _NNLossGradientNodeClass) Alloc() NNLossGradientNode {
	rv := objc.Call[NNLossGradientNode](nc, objc.Sel("alloc"))
	return rv
}

func NNLossGradientNode_Alloc() NNLossGradientNode {
	return NNLossGradientNodeClass.Alloc()
}

func (nc _NNLossGradientNodeClass) New() NNLossGradientNode {
	rv := objc.Call[NNLossGradientNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNLossGradientNode() NNLossGradientNode {
	return NNLossGradientNodeClass.New()
}

func (n_ NNLossGradientNode) Init() NNLossGradientNode {
	rv := objc.Call[NNLossGradientNode](n_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131864-numberofclasses?language=objc
func (n_ NNLossGradientNode) NumberOfClasses() uint {
	rv := objc.Call[uint](n_, objc.Sel("numberOfClasses"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131867-weight?language=objc
func (n_ NNLossGradientNode) Weight() float64 {
	rv := objc.Call[float64](n_, objc.Sel("weight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131858-islabelsgradientfilter?language=objc
func (n_ NNLossGradientNode) IsLabelsGradientFilter() bool {
	rv := objc.Call[bool](n_, objc.Sel("isLabelsGradientFilter"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131865-propertycallback?language=objc
func (n_ NNLossGradientNode) PropertyCallBack() NNLossCallbackWrapper {
	rv := objc.Call[NNLossCallbackWrapper](n_, objc.Sel("propertyCallBack"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131865-propertycallback?language=objc
func (n_ NNLossGradientNode) SetPropertyCallBack(value PNNLossCallback) {
	po0 := objc.WrapAsProtocol("MPSNNLossCallback", value)
	objc.Call[objc.Void](n_, objc.Sel("setPropertyCallBack:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131865-propertycallback?language=objc
func (n_ NNLossGradientNode) SetPropertyCallBackObject(valueObject objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setPropertyCallBack:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131854-epsilon?language=objc
func (n_ NNLossGradientNode) Epsilon() float64 {
	rv := objc.Call[float64](n_, objc.Sel("epsilon"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131853-delta?language=objc
func (n_ NNLossGradientNode) Delta() float64 {
	rv := objc.Call[float64](n_, objc.Sel("delta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131866-reductiontype?language=objc
func (n_ NNLossGradientNode) ReductionType() CNNReductionType {
	rv := objc.Call[CNNReductionType](n_, objc.Sel("reductionType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3547988-reduceacrossbatch?language=objc
func (n_ NNLossGradientNode) ReduceAcrossBatch() bool {
	rv := objc.Call[bool](n_, objc.Sel("reduceAcrossBatch"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131860-losstype?language=objc
func (n_ NNLossGradientNode) LossType() CNNLossType {
	rv := objc.Call[CNNLossType](n_, objc.Sel("lossType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradientnode/3131859-labelsmoothing?language=objc
func (n_ NNLossGradientNode) LabelSmoothing() float64 {
	rv := objc.Call[float64](n_, objc.Sel("labelSmoothing"))
	return rv
}
