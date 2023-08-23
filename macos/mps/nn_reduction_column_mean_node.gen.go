// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReductionColumnMeanNode] class.
var NNReductionColumnMeanNodeClass = _NNReductionColumnMeanNodeClass{objc.GetClass("MPSNNReductionColumnMeanNode")}

type _NNReductionColumnMeanNodeClass struct {
	objc.Class
}

// An interface definition for the [NNReductionColumnMeanNode] class.
type INNReductionColumnMeanNode interface {
	INNUnaryReductionNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductioncolumnmeannode?language=objc
type NNReductionColumnMeanNode struct {
	NNUnaryReductionNode
}

func NNReductionColumnMeanNodeFrom(ptr unsafe.Pointer) NNReductionColumnMeanNode {
	return NNReductionColumnMeanNode{
		NNUnaryReductionNode: NNUnaryReductionNodeFrom(ptr),
	}
}

func (nc _NNReductionColumnMeanNodeClass) Alloc() NNReductionColumnMeanNode {
	rv := objc.Call[NNReductionColumnMeanNode](nc, objc.Sel("alloc"))
	return rv
}

func NNReductionColumnMeanNode_Alloc() NNReductionColumnMeanNode {
	return NNReductionColumnMeanNodeClass.Alloc()
}

func (nc _NNReductionColumnMeanNodeClass) New() NNReductionColumnMeanNode {
	rv := objc.Call[NNReductionColumnMeanNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReductionColumnMeanNode() NNReductionColumnMeanNode {
	return NNReductionColumnMeanNodeClass.New()
}

func (n_ NNReductionColumnMeanNode) Init() NNReductionColumnMeanNode {
	rv := objc.Call[NNReductionColumnMeanNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNReductionColumnMeanNodeClass) NodeWithSource(sourceNode INNImageNode) NNReductionColumnMeanNode {
	rv := objc.Call[NNReductionColumnMeanNode](nc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037425-nodewithsource?language=objc
func NNReductionColumnMeanNode_NodeWithSource(sourceNode INNImageNode) NNReductionColumnMeanNode {
	return NNReductionColumnMeanNodeClass.NodeWithSource(sourceNode)
}

func (n_ NNReductionColumnMeanNode) InitWithSource(sourceNode INNImageNode) NNReductionColumnMeanNode {
	rv := objc.Call[NNReductionColumnMeanNode](n_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037424-initwithsource?language=objc
func NewNNReductionColumnMeanNodeWithSource(sourceNode INNImageNode) NNReductionColumnMeanNode {
	instance := NNReductionColumnMeanNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}
