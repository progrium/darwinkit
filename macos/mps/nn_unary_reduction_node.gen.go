// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNUnaryReductionNode] class.
var NNUnaryReductionNodeClass = _NNUnaryReductionNodeClass{objc.GetClass("MPSNNUnaryReductionNode")}

type _NNUnaryReductionNodeClass struct {
	objc.Class
}

// An interface definition for the [NNUnaryReductionNode] class.
type INNUnaryReductionNode interface {
	INNFilterNode
	ClipRectSource() metal.Region
	SetClipRectSource(value metal.Region)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode?language=objc
type NNUnaryReductionNode struct {
	NNFilterNode
}

func NNUnaryReductionNodeFrom(ptr unsafe.Pointer) NNUnaryReductionNode {
	return NNUnaryReductionNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (nc _NNUnaryReductionNodeClass) NodeWithSource(sourceNode INNImageNode) NNUnaryReductionNode {
	rv := objc.Call[NNUnaryReductionNode](nc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037425-nodewithsource?language=objc
func NNUnaryReductionNode_NodeWithSource(sourceNode INNImageNode) NNUnaryReductionNode {
	return NNUnaryReductionNodeClass.NodeWithSource(sourceNode)
}

func (n_ NNUnaryReductionNode) InitWithSource(sourceNode INNImageNode) NNUnaryReductionNode {
	rv := objc.Call[NNUnaryReductionNode](n_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037424-initwithsource?language=objc
func NewNNUnaryReductionNodeWithSource(sourceNode INNImageNode) NNUnaryReductionNode {
	instance := NNUnaryReductionNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (nc _NNUnaryReductionNodeClass) Alloc() NNUnaryReductionNode {
	rv := objc.Call[NNUnaryReductionNode](nc, objc.Sel("alloc"))
	return rv
}

func NNUnaryReductionNode_Alloc() NNUnaryReductionNode {
	return NNUnaryReductionNodeClass.Alloc()
}

func (nc _NNUnaryReductionNodeClass) New() NNUnaryReductionNode {
	rv := objc.Call[NNUnaryReductionNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNUnaryReductionNode() NNUnaryReductionNode {
	return NNUnaryReductionNodeClass.New()
}

func (n_ NNUnaryReductionNode) Init() NNUnaryReductionNode {
	rv := objc.Call[NNUnaryReductionNode](n_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037423-cliprectsource?language=objc
func (n_ NNUnaryReductionNode) ClipRectSource() metal.Region {
	rv := objc.Call[metal.Region](n_, objc.Sel("clipRectSource"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037423-cliprectsource?language=objc
func (n_ NNUnaryReductionNode) SetClipRectSource(value metal.Region) {
	objc.Call[objc.Void](n_, objc.Sel("setClipRectSource:"), value)
}
