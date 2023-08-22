// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReductionColumnMaxNode] class.
var NNReductionColumnMaxNodeClass = _NNReductionColumnMaxNodeClass{objc.GetClass("MPSNNReductionColumnMaxNode")}

type _NNReductionColumnMaxNodeClass struct {
	objc.Class
}

// An interface definition for the [NNReductionColumnMaxNode] class.
type INNReductionColumnMaxNode interface {
	INNUnaryReductionNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductioncolumnmaxnode?language=objc
type NNReductionColumnMaxNode struct {
	NNUnaryReductionNode
}

func NNReductionColumnMaxNodeFrom(ptr unsafe.Pointer) NNReductionColumnMaxNode {
	return NNReductionColumnMaxNode{
		NNUnaryReductionNode: NNUnaryReductionNodeFrom(ptr),
	}
}

func (nc _NNReductionColumnMaxNodeClass) Alloc() NNReductionColumnMaxNode {
	rv := objc.Call[NNReductionColumnMaxNode](nc, objc.Sel("alloc"))
	return rv
}

func NNReductionColumnMaxNode_Alloc() NNReductionColumnMaxNode {
	return NNReductionColumnMaxNodeClass.Alloc()
}

func (nc _NNReductionColumnMaxNodeClass) New() NNReductionColumnMaxNode {
	rv := objc.Call[NNReductionColumnMaxNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReductionColumnMaxNode() NNReductionColumnMaxNode {
	return NNReductionColumnMaxNodeClass.New()
}

func (n_ NNReductionColumnMaxNode) Init() NNReductionColumnMaxNode {
	rv := objc.Call[NNReductionColumnMaxNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNReductionColumnMaxNodeClass) NodeWithSource(sourceNode INNImageNode) NNReductionColumnMaxNode {
	rv := objc.Call[NNReductionColumnMaxNode](nc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037425-nodewithsource?language=objc
func NNReductionColumnMaxNode_NodeWithSource(sourceNode INNImageNode) NNReductionColumnMaxNode {
	return NNReductionColumnMaxNodeClass.NodeWithSource(sourceNode)
}

func (n_ NNReductionColumnMaxNode) InitWithSource(sourceNode INNImageNode) NNReductionColumnMaxNode {
	rv := objc.Call[NNReductionColumnMaxNode](n_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037424-initwithsource?language=objc
func NewNNReductionColumnMaxNodeWithSource(sourceNode INNImageNode) NNReductionColumnMaxNode {
	instance := NNReductionColumnMaxNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}
