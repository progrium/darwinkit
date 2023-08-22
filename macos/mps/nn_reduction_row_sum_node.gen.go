// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReductionRowSumNode] class.
var NNReductionRowSumNodeClass = _NNReductionRowSumNodeClass{objc.GetClass("MPSNNReductionRowSumNode")}

type _NNReductionRowSumNodeClass struct {
	objc.Class
}

// An interface definition for the [NNReductionRowSumNode] class.
type INNReductionRowSumNode interface {
	INNUnaryReductionNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionrowsumnode?language=objc
type NNReductionRowSumNode struct {
	NNUnaryReductionNode
}

func NNReductionRowSumNodeFrom(ptr unsafe.Pointer) NNReductionRowSumNode {
	return NNReductionRowSumNode{
		NNUnaryReductionNode: NNUnaryReductionNodeFrom(ptr),
	}
}

func (nc _NNReductionRowSumNodeClass) Alloc() NNReductionRowSumNode {
	rv := objc.Call[NNReductionRowSumNode](nc, objc.Sel("alloc"))
	return rv
}

func NNReductionRowSumNode_Alloc() NNReductionRowSumNode {
	return NNReductionRowSumNodeClass.Alloc()
}

func (nc _NNReductionRowSumNodeClass) New() NNReductionRowSumNode {
	rv := objc.Call[NNReductionRowSumNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReductionRowSumNode() NNReductionRowSumNode {
	return NNReductionRowSumNodeClass.New()
}

func (n_ NNReductionRowSumNode) Init() NNReductionRowSumNode {
	rv := objc.Call[NNReductionRowSumNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNReductionRowSumNodeClass) NodeWithSource(sourceNode INNImageNode) NNReductionRowSumNode {
	rv := objc.Call[NNReductionRowSumNode](nc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037425-nodewithsource?language=objc
func NNReductionRowSumNode_NodeWithSource(sourceNode INNImageNode) NNReductionRowSumNode {
	return NNReductionRowSumNodeClass.NodeWithSource(sourceNode)
}

func (n_ NNReductionRowSumNode) InitWithSource(sourceNode INNImageNode) NNReductionRowSumNode {
	rv := objc.Call[NNReductionRowSumNode](n_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037424-initwithsource?language=objc
func NewNNReductionRowSumNodeWithSource(sourceNode INNImageNode) NNReductionRowSumNode {
	instance := NNReductionRowSumNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}
