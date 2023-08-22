// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNGramMatrixCalculationGradientNode] class.
var NNGramMatrixCalculationGradientNodeClass = _NNGramMatrixCalculationGradientNodeClass{objc.GetClass("MPSNNGramMatrixCalculationGradientNode")}

type _NNGramMatrixCalculationGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [NNGramMatrixCalculationGradientNode] class.
type INNGramMatrixCalculationGradientNode interface {
	INNGradientFilterNode
	Alpha() float64
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradientnode?language=objc
type NNGramMatrixCalculationGradientNode struct {
	NNGradientFilterNode
}

func NNGramMatrixCalculationGradientNodeFrom(ptr unsafe.Pointer) NNGramMatrixCalculationGradientNode {
	return NNGramMatrixCalculationGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (n_ NNGramMatrixCalculationGradientNode) InitWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNGramMatrixCalculationGradientNode {
	rv := objc.Call[NNGramMatrixCalculationGradientNode](n_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradientnode/3114089-initwithsourcegradient?language=objc
func NewNNGramMatrixCalculationGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNGramMatrixCalculationGradientNode {
	instance := NNGramMatrixCalculationGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientState(sourceGradient, sourceImage, gradientState)
	instance.Autorelease()
	return instance
}

func (nc _NNGramMatrixCalculationGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNGramMatrixCalculationGradientNode {
	rv := objc.Call[NNGramMatrixCalculationGradientNode](nc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradientnode/3114091-nodewithsourcegradient?language=objc
func NNGramMatrixCalculationGradientNode_NodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNGramMatrixCalculationGradientNode {
	return NNGramMatrixCalculationGradientNodeClass.NodeWithSourceGradientSourceImageGradientState(sourceGradient, sourceImage, gradientState)
}

func (nc _NNGramMatrixCalculationGradientNodeClass) Alloc() NNGramMatrixCalculationGradientNode {
	rv := objc.Call[NNGramMatrixCalculationGradientNode](nc, objc.Sel("alloc"))
	return rv
}

func NNGramMatrixCalculationGradientNode_Alloc() NNGramMatrixCalculationGradientNode {
	return NNGramMatrixCalculationGradientNodeClass.Alloc()
}

func (nc _NNGramMatrixCalculationGradientNodeClass) New() NNGramMatrixCalculationGradientNode {
	rv := objc.Call[NNGramMatrixCalculationGradientNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNGramMatrixCalculationGradientNode() NNGramMatrixCalculationGradientNode {
	return NNGramMatrixCalculationGradientNodeClass.New()
}

func (n_ NNGramMatrixCalculationGradientNode) Init() NNGramMatrixCalculationGradientNode {
	rv := objc.Call[NNGramMatrixCalculationGradientNode](n_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngrammatrixcalculationgradientnode/3114088-alpha?language=objc
func (n_ NNGramMatrixCalculationGradientNode) Alpha() float64 {
	rv := objc.Call[float64](n_, objc.Sel("alpha"))
	return rv
}
