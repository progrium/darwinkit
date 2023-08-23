// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReductionSpatialMeanNode] class.
var NNReductionSpatialMeanNodeClass = _NNReductionSpatialMeanNodeClass{objc.GetClass("MPSNNReductionSpatialMeanNode")}

type _NNReductionSpatialMeanNodeClass struct {
	objc.Class
}

// An interface definition for the [NNReductionSpatialMeanNode] class.
type INNReductionSpatialMeanNode interface {
	INNUnaryReductionNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionspatialmeannode?language=objc
type NNReductionSpatialMeanNode struct {
	NNUnaryReductionNode
}

func NNReductionSpatialMeanNodeFrom(ptr unsafe.Pointer) NNReductionSpatialMeanNode {
	return NNReductionSpatialMeanNode{
		NNUnaryReductionNode: NNUnaryReductionNodeFrom(ptr),
	}
}

func (nc _NNReductionSpatialMeanNodeClass) Alloc() NNReductionSpatialMeanNode {
	rv := objc.Call[NNReductionSpatialMeanNode](nc, objc.Sel("alloc"))
	return rv
}

func NNReductionSpatialMeanNode_Alloc() NNReductionSpatialMeanNode {
	return NNReductionSpatialMeanNodeClass.Alloc()
}

func (nc _NNReductionSpatialMeanNodeClass) New() NNReductionSpatialMeanNode {
	rv := objc.Call[NNReductionSpatialMeanNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReductionSpatialMeanNode() NNReductionSpatialMeanNode {
	return NNReductionSpatialMeanNodeClass.New()
}

func (n_ NNReductionSpatialMeanNode) Init() NNReductionSpatialMeanNode {
	rv := objc.Call[NNReductionSpatialMeanNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNReductionSpatialMeanNodeClass) NodeWithSource(sourceNode INNImageNode) NNReductionSpatialMeanNode {
	rv := objc.Call[NNReductionSpatialMeanNode](nc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037425-nodewithsource?language=objc
func NNReductionSpatialMeanNode_NodeWithSource(sourceNode INNImageNode) NNReductionSpatialMeanNode {
	return NNReductionSpatialMeanNodeClass.NodeWithSource(sourceNode)
}

func (n_ NNReductionSpatialMeanNode) InitWithSource(sourceNode INNImageNode) NNReductionSpatialMeanNode {
	rv := objc.Call[NNReductionSpatialMeanNode](n_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037424-initwithsource?language=objc
func NewNNReductionSpatialMeanNodeWithSource(sourceNode INNImageNode) NNReductionSpatialMeanNode {
	instance := NNReductionSpatialMeanNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}
