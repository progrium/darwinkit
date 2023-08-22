// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReductionFeatureChannelsArgumentMaxNode] class.
var NNReductionFeatureChannelsArgumentMaxNodeClass = _NNReductionFeatureChannelsArgumentMaxNodeClass{objc.GetClass("MPSNNReductionFeatureChannelsArgumentMaxNode")}

type _NNReductionFeatureChannelsArgumentMaxNodeClass struct {
	objc.Class
}

// An interface definition for the [NNReductionFeatureChannelsArgumentMaxNode] class.
type INNReductionFeatureChannelsArgumentMaxNode interface {
	INNUnaryReductionNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionfeaturechannelsargumentmaxnode?language=objc
type NNReductionFeatureChannelsArgumentMaxNode struct {
	NNUnaryReductionNode
}

func NNReductionFeatureChannelsArgumentMaxNodeFrom(ptr unsafe.Pointer) NNReductionFeatureChannelsArgumentMaxNode {
	return NNReductionFeatureChannelsArgumentMaxNode{
		NNUnaryReductionNode: NNUnaryReductionNodeFrom(ptr),
	}
}

func (nc _NNReductionFeatureChannelsArgumentMaxNodeClass) Alloc() NNReductionFeatureChannelsArgumentMaxNode {
	rv := objc.Call[NNReductionFeatureChannelsArgumentMaxNode](nc, objc.Sel("alloc"))
	return rv
}

func NNReductionFeatureChannelsArgumentMaxNode_Alloc() NNReductionFeatureChannelsArgumentMaxNode {
	return NNReductionFeatureChannelsArgumentMaxNodeClass.Alloc()
}

func (nc _NNReductionFeatureChannelsArgumentMaxNodeClass) New() NNReductionFeatureChannelsArgumentMaxNode {
	rv := objc.Call[NNReductionFeatureChannelsArgumentMaxNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReductionFeatureChannelsArgumentMaxNode() NNReductionFeatureChannelsArgumentMaxNode {
	return NNReductionFeatureChannelsArgumentMaxNodeClass.New()
}

func (n_ NNReductionFeatureChannelsArgumentMaxNode) Init() NNReductionFeatureChannelsArgumentMaxNode {
	rv := objc.Call[NNReductionFeatureChannelsArgumentMaxNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNReductionFeatureChannelsArgumentMaxNodeClass) NodeWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsArgumentMaxNode {
	rv := objc.Call[NNReductionFeatureChannelsArgumentMaxNode](nc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037425-nodewithsource?language=objc
func NNReductionFeatureChannelsArgumentMaxNode_NodeWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsArgumentMaxNode {
	return NNReductionFeatureChannelsArgumentMaxNodeClass.NodeWithSource(sourceNode)
}

func (n_ NNReductionFeatureChannelsArgumentMaxNode) InitWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsArgumentMaxNode {
	rv := objc.Call[NNReductionFeatureChannelsArgumentMaxNode](n_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037424-initwithsource?language=objc
func NewNNReductionFeatureChannelsArgumentMaxNodeWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsArgumentMaxNode {
	instance := NNReductionFeatureChannelsArgumentMaxNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}
