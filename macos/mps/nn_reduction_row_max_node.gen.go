// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReductionRowMaxNode] class.
var NNReductionRowMaxNodeClass = _NNReductionRowMaxNodeClass{objc.GetClass("MPSNNReductionRowMaxNode")}

type _NNReductionRowMaxNodeClass struct {
	objc.Class
}

// An interface definition for the [NNReductionRowMaxNode] class.
type INNReductionRowMaxNode interface {
	INNUnaryReductionNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionrowmaxnode?language=objc
type NNReductionRowMaxNode struct {
	NNUnaryReductionNode
}

func NNReductionRowMaxNodeFrom(ptr unsafe.Pointer) NNReductionRowMaxNode {
	return NNReductionRowMaxNode{
		NNUnaryReductionNode: NNUnaryReductionNodeFrom(ptr),
	}
}

func (nc _NNReductionRowMaxNodeClass) Alloc() NNReductionRowMaxNode {
	rv := objc.Call[NNReductionRowMaxNode](nc, objc.Sel("alloc"))
	return rv
}

func NNReductionRowMaxNode_Alloc() NNReductionRowMaxNode {
	return NNReductionRowMaxNodeClass.Alloc()
}

func (nc _NNReductionRowMaxNodeClass) New() NNReductionRowMaxNode {
	rv := objc.Call[NNReductionRowMaxNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReductionRowMaxNode() NNReductionRowMaxNode {
	return NNReductionRowMaxNodeClass.New()
}

func (n_ NNReductionRowMaxNode) Init() NNReductionRowMaxNode {
	rv := objc.Call[NNReductionRowMaxNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNReductionRowMaxNodeClass) NodeWithSource(sourceNode INNImageNode) NNReductionRowMaxNode {
	rv := objc.Call[NNReductionRowMaxNode](nc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037425-nodewithsource?language=objc
func NNReductionRowMaxNode_NodeWithSource(sourceNode INNImageNode) NNReductionRowMaxNode {
	return NNReductionRowMaxNodeClass.NodeWithSource(sourceNode)
}

func (n_ NNReductionRowMaxNode) InitWithSource(sourceNode INNImageNode) NNReductionRowMaxNode {
	rv := objc.Call[NNReductionRowMaxNode](n_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037424-initwithsource?language=objc
func NewNNReductionRowMaxNodeWithSource(sourceNode INNImageNode) NNReductionRowMaxNode {
	instance := NNReductionRowMaxNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}
