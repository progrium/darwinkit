// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReductionFeatureChannelsMinNode] class.
var NNReductionFeatureChannelsMinNodeClass = _NNReductionFeatureChannelsMinNodeClass{objc.GetClass("MPSNNReductionFeatureChannelsMinNode")}

type _NNReductionFeatureChannelsMinNodeClass struct {
	objc.Class
}

// An interface definition for the [NNReductionFeatureChannelsMinNode] class.
type INNReductionFeatureChannelsMinNode interface {
	INNUnaryReductionNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionfeaturechannelsminnode?language=objc
type NNReductionFeatureChannelsMinNode struct {
	NNUnaryReductionNode
}

func NNReductionFeatureChannelsMinNodeFrom(ptr unsafe.Pointer) NNReductionFeatureChannelsMinNode {
	return NNReductionFeatureChannelsMinNode{
		NNUnaryReductionNode: NNUnaryReductionNodeFrom(ptr),
	}
}

func (nc _NNReductionFeatureChannelsMinNodeClass) Alloc() NNReductionFeatureChannelsMinNode {
	rv := objc.Call[NNReductionFeatureChannelsMinNode](nc, objc.Sel("alloc"))
	return rv
}

func NNReductionFeatureChannelsMinNode_Alloc() NNReductionFeatureChannelsMinNode {
	return NNReductionFeatureChannelsMinNodeClass.Alloc()
}

func (nc _NNReductionFeatureChannelsMinNodeClass) New() NNReductionFeatureChannelsMinNode {
	rv := objc.Call[NNReductionFeatureChannelsMinNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReductionFeatureChannelsMinNode() NNReductionFeatureChannelsMinNode {
	return NNReductionFeatureChannelsMinNodeClass.New()
}

func (n_ NNReductionFeatureChannelsMinNode) Init() NNReductionFeatureChannelsMinNode {
	rv := objc.Call[NNReductionFeatureChannelsMinNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNReductionFeatureChannelsMinNodeClass) NodeWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsMinNode {
	rv := objc.Call[NNReductionFeatureChannelsMinNode](nc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037425-nodewithsource?language=objc
func NNReductionFeatureChannelsMinNode_NodeWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsMinNode {
	return NNReductionFeatureChannelsMinNodeClass.NodeWithSource(sourceNode)
}

func (n_ NNReductionFeatureChannelsMinNode) InitWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsMinNode {
	rv := objc.Call[NNReductionFeatureChannelsMinNode](n_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037424-initwithsource?language=objc
func NewNNReductionFeatureChannelsMinNodeWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsMinNode {
	instance := NNReductionFeatureChannelsMinNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}
