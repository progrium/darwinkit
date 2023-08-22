// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReductionRowMinNode] class.
var NNReductionRowMinNodeClass = _NNReductionRowMinNodeClass{objc.GetClass("MPSNNReductionRowMinNode")}

type _NNReductionRowMinNodeClass struct {
	objc.Class
}

// An interface definition for the [NNReductionRowMinNode] class.
type INNReductionRowMinNode interface {
	INNUnaryReductionNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionrowminnode?language=objc
type NNReductionRowMinNode struct {
	NNUnaryReductionNode
}

func NNReductionRowMinNodeFrom(ptr unsafe.Pointer) NNReductionRowMinNode {
	return NNReductionRowMinNode{
		NNUnaryReductionNode: NNUnaryReductionNodeFrom(ptr),
	}
}

func (nc _NNReductionRowMinNodeClass) Alloc() NNReductionRowMinNode {
	rv := objc.Call[NNReductionRowMinNode](nc, objc.Sel("alloc"))
	return rv
}

func NNReductionRowMinNode_Alloc() NNReductionRowMinNode {
	return NNReductionRowMinNodeClass.Alloc()
}

func (nc _NNReductionRowMinNodeClass) New() NNReductionRowMinNode {
	rv := objc.Call[NNReductionRowMinNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReductionRowMinNode() NNReductionRowMinNode {
	return NNReductionRowMinNodeClass.New()
}

func (n_ NNReductionRowMinNode) Init() NNReductionRowMinNode {
	rv := objc.Call[NNReductionRowMinNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNReductionRowMinNodeClass) NodeWithSource(sourceNode INNImageNode) NNReductionRowMinNode {
	rv := objc.Call[NNReductionRowMinNode](nc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037425-nodewithsource?language=objc
func NNReductionRowMinNode_NodeWithSource(sourceNode INNImageNode) NNReductionRowMinNode {
	return NNReductionRowMinNodeClass.NodeWithSource(sourceNode)
}

func (n_ NNReductionRowMinNode) InitWithSource(sourceNode INNImageNode) NNReductionRowMinNode {
	rv := objc.Call[NNReductionRowMinNode](n_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037424-initwithsource?language=objc
func NewNNReductionRowMinNodeWithSource(sourceNode INNImageNode) NNReductionRowMinNode {
	instance := NNReductionRowMinNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}
