// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReductionFeatureChannelsArgumentMinNode] class.
var NNReductionFeatureChannelsArgumentMinNodeClass = _NNReductionFeatureChannelsArgumentMinNodeClass{objc.GetClass("MPSNNReductionFeatureChannelsArgumentMinNode")}

type _NNReductionFeatureChannelsArgumentMinNodeClass struct {
	objc.Class
}

// An interface definition for the [NNReductionFeatureChannelsArgumentMinNode] class.
type INNReductionFeatureChannelsArgumentMinNode interface {
	INNUnaryReductionNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionfeaturechannelsargumentminnode?language=objc
type NNReductionFeatureChannelsArgumentMinNode struct {
	NNUnaryReductionNode
}

func NNReductionFeatureChannelsArgumentMinNodeFrom(ptr unsafe.Pointer) NNReductionFeatureChannelsArgumentMinNode {
	return NNReductionFeatureChannelsArgumentMinNode{
		NNUnaryReductionNode: NNUnaryReductionNodeFrom(ptr),
	}
}

func (nc _NNReductionFeatureChannelsArgumentMinNodeClass) Alloc() NNReductionFeatureChannelsArgumentMinNode {
	rv := objc.Call[NNReductionFeatureChannelsArgumentMinNode](nc, objc.Sel("alloc"))
	return rv
}

func NNReductionFeatureChannelsArgumentMinNode_Alloc() NNReductionFeatureChannelsArgumentMinNode {
	return NNReductionFeatureChannelsArgumentMinNodeClass.Alloc()
}

func (nc _NNReductionFeatureChannelsArgumentMinNodeClass) New() NNReductionFeatureChannelsArgumentMinNode {
	rv := objc.Call[NNReductionFeatureChannelsArgumentMinNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReductionFeatureChannelsArgumentMinNode() NNReductionFeatureChannelsArgumentMinNode {
	return NNReductionFeatureChannelsArgumentMinNodeClass.New()
}

func (n_ NNReductionFeatureChannelsArgumentMinNode) Init() NNReductionFeatureChannelsArgumentMinNode {
	rv := objc.Call[NNReductionFeatureChannelsArgumentMinNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNReductionFeatureChannelsArgumentMinNodeClass) NodeWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsArgumentMinNode {
	rv := objc.Call[NNReductionFeatureChannelsArgumentMinNode](nc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037425-nodewithsource?language=objc
func NNReductionFeatureChannelsArgumentMinNode_NodeWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsArgumentMinNode {
	return NNReductionFeatureChannelsArgumentMinNodeClass.NodeWithSource(sourceNode)
}

func (n_ NNReductionFeatureChannelsArgumentMinNode) InitWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsArgumentMinNode {
	rv := objc.Call[NNReductionFeatureChannelsArgumentMinNode](n_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037424-initwithsource?language=objc
func NewNNReductionFeatureChannelsArgumentMinNodeWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsArgumentMinNode {
	instance := NNReductionFeatureChannelsArgumentMinNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}
