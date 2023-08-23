// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNGramMatrixCalculationNode] class.
var NNGramMatrixCalculationNodeClass = _NNGramMatrixCalculationNodeClass{objc.GetClass("MPSNNGramMatrixCalculationNode")}

type _NNGramMatrixCalculationNodeClass struct {
	objc.Class
}

// An interface definition for the [NNGramMatrixCalculationNode] class.
type INNGramMatrixCalculationNode interface {
	INNFilterNode
	PropertyCallBack() NNGramMatrixCallbackWrapper
	SetPropertyCallBack(value PNNGramMatrixCallback)
	SetPropertyCallBackObject(valueObject objc.IObject)
	Alpha() float64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationnode?language=objc
type NNGramMatrixCalculationNode struct {
	NNFilterNode
}

func NNGramMatrixCalculationNodeFrom(ptr unsafe.Pointer) NNGramMatrixCalculationNode {
	return NNGramMatrixCalculationNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (nc _NNGramMatrixCalculationNodeClass) NodeWithSource(sourceNode INNImageNode) NNGramMatrixCalculationNode {
	rv := objc.Call[NNGramMatrixCalculationNode](nc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationnode/3114097-nodewithsource?language=objc
func NNGramMatrixCalculationNode_NodeWithSource(sourceNode INNImageNode) NNGramMatrixCalculationNode {
	return NNGramMatrixCalculationNodeClass.NodeWithSource(sourceNode)
}

func (n_ NNGramMatrixCalculationNode) InitWithSource(sourceNode INNImageNode) NNGramMatrixCalculationNode {
	rv := objc.Call[NNGramMatrixCalculationNode](n_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationnode/3114095-initwithsource?language=objc
func NewNNGramMatrixCalculationNodeWithSource(sourceNode INNImageNode) NNGramMatrixCalculationNode {
	instance := NNGramMatrixCalculationNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (nc _NNGramMatrixCalculationNodeClass) Alloc() NNGramMatrixCalculationNode {
	rv := objc.Call[NNGramMatrixCalculationNode](nc, objc.Sel("alloc"))
	return rv
}

func NNGramMatrixCalculationNode_Alloc() NNGramMatrixCalculationNode {
	return NNGramMatrixCalculationNodeClass.Alloc()
}

func (nc _NNGramMatrixCalculationNodeClass) New() NNGramMatrixCalculationNode {
	rv := objc.Call[NNGramMatrixCalculationNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNGramMatrixCalculationNode() NNGramMatrixCalculationNode {
	return NNGramMatrixCalculationNodeClass.New()
}

func (n_ NNGramMatrixCalculationNode) Init() NNGramMatrixCalculationNode {
	rv := objc.Call[NNGramMatrixCalculationNode](n_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationnode/3131844-propertycallback?language=objc
func (n_ NNGramMatrixCalculationNode) PropertyCallBack() NNGramMatrixCallbackWrapper {
	rv := objc.Call[NNGramMatrixCallbackWrapper](n_, objc.Sel("propertyCallBack"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationnode/3131844-propertycallback?language=objc
func (n_ NNGramMatrixCalculationNode) SetPropertyCallBack(value PNNGramMatrixCallback) {
	po0 := objc.WrapAsProtocol("MPSNNGramMatrixCallback", value)
	objc.Call[objc.Void](n_, objc.Sel("setPropertyCallBack:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationnode/3131844-propertycallback?language=objc
func (n_ NNGramMatrixCalculationNode) SetPropertyCallBackObject(valueObject objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setPropertyCallBack:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationnode/3114094-alpha?language=objc
func (n_ NNGramMatrixCalculationNode) Alpha() float64 {
	rv := objc.Call[float64](n_, objc.Sel("alpha"))
	return rv
}
