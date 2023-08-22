// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReductionFeatureChannelsMeanNode] class.
var NNReductionFeatureChannelsMeanNodeClass = _NNReductionFeatureChannelsMeanNodeClass{objc.GetClass("MPSNNReductionFeatureChannelsMeanNode")}

type _NNReductionFeatureChannelsMeanNodeClass struct {
	objc.Class
}

// An interface definition for the [NNReductionFeatureChannelsMeanNode] class.
type INNReductionFeatureChannelsMeanNode interface {
	INNUnaryReductionNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionfeaturechannelsmeannode?language=objc
type NNReductionFeatureChannelsMeanNode struct {
	NNUnaryReductionNode
}

func NNReductionFeatureChannelsMeanNodeFrom(ptr unsafe.Pointer) NNReductionFeatureChannelsMeanNode {
	return NNReductionFeatureChannelsMeanNode{
		NNUnaryReductionNode: NNUnaryReductionNodeFrom(ptr),
	}
}

func (nc _NNReductionFeatureChannelsMeanNodeClass) Alloc() NNReductionFeatureChannelsMeanNode {
	rv := objc.Call[NNReductionFeatureChannelsMeanNode](nc, objc.Sel("alloc"))
	return rv
}

func NNReductionFeatureChannelsMeanNode_Alloc() NNReductionFeatureChannelsMeanNode {
	return NNReductionFeatureChannelsMeanNodeClass.Alloc()
}

func (nc _NNReductionFeatureChannelsMeanNodeClass) New() NNReductionFeatureChannelsMeanNode {
	rv := objc.Call[NNReductionFeatureChannelsMeanNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReductionFeatureChannelsMeanNode() NNReductionFeatureChannelsMeanNode {
	return NNReductionFeatureChannelsMeanNodeClass.New()
}

func (n_ NNReductionFeatureChannelsMeanNode) Init() NNReductionFeatureChannelsMeanNode {
	rv := objc.Call[NNReductionFeatureChannelsMeanNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNReductionFeatureChannelsMeanNodeClass) NodeWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsMeanNode {
	rv := objc.Call[NNReductionFeatureChannelsMeanNode](nc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037425-nodewithsource?language=objc
func NNReductionFeatureChannelsMeanNode_NodeWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsMeanNode {
	return NNReductionFeatureChannelsMeanNodeClass.NodeWithSource(sourceNode)
}

func (n_ NNReductionFeatureChannelsMeanNode) InitWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsMeanNode {
	rv := objc.Call[NNReductionFeatureChannelsMeanNode](n_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037424-initwithsource?language=objc
func NewNNReductionFeatureChannelsMeanNodeWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsMeanNode {
	instance := NNReductionFeatureChannelsMeanNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}
