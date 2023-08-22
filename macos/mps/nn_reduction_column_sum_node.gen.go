// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReductionColumnSumNode] class.
var NNReductionColumnSumNodeClass = _NNReductionColumnSumNodeClass{objc.GetClass("MPSNNReductionColumnSumNode")}

type _NNReductionColumnSumNodeClass struct {
	objc.Class
}

// An interface definition for the [NNReductionColumnSumNode] class.
type INNReductionColumnSumNode interface {
	INNUnaryReductionNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductioncolumnsumnode?language=objc
type NNReductionColumnSumNode struct {
	NNUnaryReductionNode
}

func NNReductionColumnSumNodeFrom(ptr unsafe.Pointer) NNReductionColumnSumNode {
	return NNReductionColumnSumNode{
		NNUnaryReductionNode: NNUnaryReductionNodeFrom(ptr),
	}
}

func (nc _NNReductionColumnSumNodeClass) Alloc() NNReductionColumnSumNode {
	rv := objc.Call[NNReductionColumnSumNode](nc, objc.Sel("alloc"))
	return rv
}

func NNReductionColumnSumNode_Alloc() NNReductionColumnSumNode {
	return NNReductionColumnSumNodeClass.Alloc()
}

func (nc _NNReductionColumnSumNodeClass) New() NNReductionColumnSumNode {
	rv := objc.Call[NNReductionColumnSumNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReductionColumnSumNode() NNReductionColumnSumNode {
	return NNReductionColumnSumNodeClass.New()
}

func (n_ NNReductionColumnSumNode) Init() NNReductionColumnSumNode {
	rv := objc.Call[NNReductionColumnSumNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNReductionColumnSumNodeClass) NodeWithSource(sourceNode INNImageNode) NNReductionColumnSumNode {
	rv := objc.Call[NNReductionColumnSumNode](nc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037425-nodewithsource?language=objc
func NNReductionColumnSumNode_NodeWithSource(sourceNode INNImageNode) NNReductionColumnSumNode {
	return NNReductionColumnSumNodeClass.NodeWithSource(sourceNode)
}

func (n_ NNReductionColumnSumNode) InitWithSource(sourceNode INNImageNode) NNReductionColumnSumNode {
	rv := objc.Call[NNReductionColumnSumNode](n_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037424-initwithsource?language=objc
func NewNNReductionColumnSumNodeWithSource(sourceNode INNImageNode) NNReductionColumnSumNode {
	instance := NNReductionColumnSumNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}
