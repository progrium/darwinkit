// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReductionFeatureChannelsMaxNode] class.
var NNReductionFeatureChannelsMaxNodeClass = _NNReductionFeatureChannelsMaxNodeClass{objc.GetClass("MPSNNReductionFeatureChannelsMaxNode")}

type _NNReductionFeatureChannelsMaxNodeClass struct {
	objc.Class
}

// An interface definition for the [NNReductionFeatureChannelsMaxNode] class.
type INNReductionFeatureChannelsMaxNode interface {
	INNUnaryReductionNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionfeaturechannelsmaxnode?language=objc
type NNReductionFeatureChannelsMaxNode struct {
	NNUnaryReductionNode
}

func NNReductionFeatureChannelsMaxNodeFrom(ptr unsafe.Pointer) NNReductionFeatureChannelsMaxNode {
	return NNReductionFeatureChannelsMaxNode{
		NNUnaryReductionNode: NNUnaryReductionNodeFrom(ptr),
	}
}

func (nc _NNReductionFeatureChannelsMaxNodeClass) Alloc() NNReductionFeatureChannelsMaxNode {
	rv := objc.Call[NNReductionFeatureChannelsMaxNode](nc, objc.Sel("alloc"))
	return rv
}

func NNReductionFeatureChannelsMaxNode_Alloc() NNReductionFeatureChannelsMaxNode {
	return NNReductionFeatureChannelsMaxNodeClass.Alloc()
}

func (nc _NNReductionFeatureChannelsMaxNodeClass) New() NNReductionFeatureChannelsMaxNode {
	rv := objc.Call[NNReductionFeatureChannelsMaxNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReductionFeatureChannelsMaxNode() NNReductionFeatureChannelsMaxNode {
	return NNReductionFeatureChannelsMaxNodeClass.New()
}

func (n_ NNReductionFeatureChannelsMaxNode) Init() NNReductionFeatureChannelsMaxNode {
	rv := objc.Call[NNReductionFeatureChannelsMaxNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNReductionFeatureChannelsMaxNodeClass) NodeWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsMaxNode {
	rv := objc.Call[NNReductionFeatureChannelsMaxNode](nc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037425-nodewithsource?language=objc
func NNReductionFeatureChannelsMaxNode_NodeWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsMaxNode {
	return NNReductionFeatureChannelsMaxNodeClass.NodeWithSource(sourceNode)
}

func (n_ NNReductionFeatureChannelsMaxNode) InitWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsMaxNode {
	rv := objc.Call[NNReductionFeatureChannelsMaxNode](n_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037424-initwithsource?language=objc
func NewNNReductionFeatureChannelsMaxNodeWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsMaxNode {
	instance := NNReductionFeatureChannelsMaxNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}
