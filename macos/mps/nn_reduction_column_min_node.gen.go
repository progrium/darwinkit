// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReductionColumnMinNode] class.
var NNReductionColumnMinNodeClass = _NNReductionColumnMinNodeClass{objc.GetClass("MPSNNReductionColumnMinNode")}

type _NNReductionColumnMinNodeClass struct {
	objc.Class
}

// An interface definition for the [NNReductionColumnMinNode] class.
type INNReductionColumnMinNode interface {
	INNUnaryReductionNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductioncolumnminnode?language=objc
type NNReductionColumnMinNode struct {
	NNUnaryReductionNode
}

func NNReductionColumnMinNodeFrom(ptr unsafe.Pointer) NNReductionColumnMinNode {
	return NNReductionColumnMinNode{
		NNUnaryReductionNode: NNUnaryReductionNodeFrom(ptr),
	}
}

func (nc _NNReductionColumnMinNodeClass) Alloc() NNReductionColumnMinNode {
	rv := objc.Call[NNReductionColumnMinNode](nc, objc.Sel("alloc"))
	return rv
}

func NNReductionColumnMinNode_Alloc() NNReductionColumnMinNode {
	return NNReductionColumnMinNodeClass.Alloc()
}

func (nc _NNReductionColumnMinNodeClass) New() NNReductionColumnMinNode {
	rv := objc.Call[NNReductionColumnMinNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReductionColumnMinNode() NNReductionColumnMinNode {
	return NNReductionColumnMinNodeClass.New()
}

func (n_ NNReductionColumnMinNode) Init() NNReductionColumnMinNode {
	rv := objc.Call[NNReductionColumnMinNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNReductionColumnMinNodeClass) NodeWithSource(sourceNode INNImageNode) NNReductionColumnMinNode {
	rv := objc.Call[NNReductionColumnMinNode](nc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037425-nodewithsource?language=objc
func NNReductionColumnMinNode_NodeWithSource(sourceNode INNImageNode) NNReductionColumnMinNode {
	return NNReductionColumnMinNodeClass.NodeWithSource(sourceNode)
}

func (n_ NNReductionColumnMinNode) InitWithSource(sourceNode INNImageNode) NNReductionColumnMinNode {
	rv := objc.Call[NNReductionColumnMinNode](n_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037424-initwithsource?language=objc
func NewNNReductionColumnMinNodeWithSource(sourceNode INNImageNode) NNReductionColumnMinNode {
	instance := NNReductionColumnMinNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}
