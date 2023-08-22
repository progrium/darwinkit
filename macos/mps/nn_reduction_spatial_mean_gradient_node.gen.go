// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReductionSpatialMeanGradientNode] class.
var NNReductionSpatialMeanGradientNodeClass = _NNReductionSpatialMeanGradientNodeClass{objc.GetClass("MPSNNReductionSpatialMeanGradientNode")}

type _NNReductionSpatialMeanGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [NNReductionSpatialMeanGradientNode] class.
type INNReductionSpatialMeanGradientNode interface {
	INNGradientFilterNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionspatialmeangradientnode?language=objc
type NNReductionSpatialMeanGradientNode struct {
	NNGradientFilterNode
}

func NNReductionSpatialMeanGradientNodeFrom(ptr unsafe.Pointer) NNReductionSpatialMeanGradientNode {
	return NNReductionSpatialMeanGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (n_ NNReductionSpatialMeanGradientNode) InitWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNReductionSpatialMeanGradientNode {
	rv := objc.Call[NNReductionSpatialMeanGradientNode](n_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionspatialmeangradientnode/3037413-initwithsourcegradient?language=objc
func NewNNReductionSpatialMeanGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNReductionSpatialMeanGradientNode {
	instance := NNReductionSpatialMeanGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientState(sourceGradient, sourceImage, gradientState)
	instance.Autorelease()
	return instance
}

func (nc _NNReductionSpatialMeanGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNReductionSpatialMeanGradientNode {
	rv := objc.Call[NNReductionSpatialMeanGradientNode](nc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionspatialmeangradientnode/3037414-nodewithsourcegradient?language=objc
func NNReductionSpatialMeanGradientNode_NodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNReductionSpatialMeanGradientNode {
	return NNReductionSpatialMeanGradientNodeClass.NodeWithSourceGradientSourceImageGradientState(sourceGradient, sourceImage, gradientState)
}

func (nc _NNReductionSpatialMeanGradientNodeClass) Alloc() NNReductionSpatialMeanGradientNode {
	rv := objc.Call[NNReductionSpatialMeanGradientNode](nc, objc.Sel("alloc"))
	return rv
}

func NNReductionSpatialMeanGradientNode_Alloc() NNReductionSpatialMeanGradientNode {
	return NNReductionSpatialMeanGradientNodeClass.Alloc()
}

func (nc _NNReductionSpatialMeanGradientNodeClass) New() NNReductionSpatialMeanGradientNode {
	rv := objc.Call[NNReductionSpatialMeanGradientNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReductionSpatialMeanGradientNode() NNReductionSpatialMeanGradientNode {
	return NNReductionSpatialMeanGradientNodeClass.New()
}

func (n_ NNReductionSpatialMeanGradientNode) Init() NNReductionSpatialMeanGradientNode {
	rv := objc.Call[NNReductionSpatialMeanGradientNode](n_, objc.Sel("init"))
	return rv
}
