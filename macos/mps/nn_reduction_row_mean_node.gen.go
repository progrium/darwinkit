// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReductionRowMeanNode] class.
var NNReductionRowMeanNodeClass = _NNReductionRowMeanNodeClass{objc.GetClass("MPSNNReductionRowMeanNode")}

type _NNReductionRowMeanNodeClass struct {
	objc.Class
}

// An interface definition for the [NNReductionRowMeanNode] class.
type INNReductionRowMeanNode interface {
	INNUnaryReductionNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionrowmeannode?language=objc
type NNReductionRowMeanNode struct {
	NNUnaryReductionNode
}

func NNReductionRowMeanNodeFrom(ptr unsafe.Pointer) NNReductionRowMeanNode {
	return NNReductionRowMeanNode{
		NNUnaryReductionNode: NNUnaryReductionNodeFrom(ptr),
	}
}

func (nc _NNReductionRowMeanNodeClass) Alloc() NNReductionRowMeanNode {
	rv := objc.Call[NNReductionRowMeanNode](nc, objc.Sel("alloc"))
	return rv
}

func NNReductionRowMeanNode_Alloc() NNReductionRowMeanNode {
	return NNReductionRowMeanNodeClass.Alloc()
}

func (nc _NNReductionRowMeanNodeClass) New() NNReductionRowMeanNode {
	rv := objc.Call[NNReductionRowMeanNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReductionRowMeanNode() NNReductionRowMeanNode {
	return NNReductionRowMeanNodeClass.New()
}

func (n_ NNReductionRowMeanNode) Init() NNReductionRowMeanNode {
	rv := objc.Call[NNReductionRowMeanNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNReductionRowMeanNodeClass) NodeWithSource(sourceNode INNImageNode) NNReductionRowMeanNode {
	rv := objc.Call[NNReductionRowMeanNode](nc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037425-nodewithsource?language=objc
func NNReductionRowMeanNode_NodeWithSource(sourceNode INNImageNode) NNReductionRowMeanNode {
	return NNReductionRowMeanNodeClass.NodeWithSource(sourceNode)
}

func (n_ NNReductionRowMeanNode) InitWithSource(sourceNode INNImageNode) NNReductionRowMeanNode {
	rv := objc.Call[NNReductionRowMeanNode](n_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037424-initwithsource?language=objc
func NewNNReductionRowMeanNodeWithSource(sourceNode INNImageNode) NNReductionRowMeanNode {
	instance := NNReductionRowMeanNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}
