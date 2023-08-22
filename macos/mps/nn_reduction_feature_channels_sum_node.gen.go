// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReductionFeatureChannelsSumNode] class.
var NNReductionFeatureChannelsSumNodeClass = _NNReductionFeatureChannelsSumNodeClass{objc.GetClass("MPSNNReductionFeatureChannelsSumNode")}

type _NNReductionFeatureChannelsSumNodeClass struct {
	objc.Class
}

// An interface definition for the [NNReductionFeatureChannelsSumNode] class.
type INNReductionFeatureChannelsSumNode interface {
	INNUnaryReductionNode
	Weight() float64
	SetWeight(value float64)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionfeaturechannelssumnode?language=objc
type NNReductionFeatureChannelsSumNode struct {
	NNUnaryReductionNode
}

func NNReductionFeatureChannelsSumNodeFrom(ptr unsafe.Pointer) NNReductionFeatureChannelsSumNode {
	return NNReductionFeatureChannelsSumNode{
		NNUnaryReductionNode: NNUnaryReductionNodeFrom(ptr),
	}
}

func (nc _NNReductionFeatureChannelsSumNodeClass) Alloc() NNReductionFeatureChannelsSumNode {
	rv := objc.Call[NNReductionFeatureChannelsSumNode](nc, objc.Sel("alloc"))
	return rv
}

func NNReductionFeatureChannelsSumNode_Alloc() NNReductionFeatureChannelsSumNode {
	return NNReductionFeatureChannelsSumNodeClass.Alloc()
}

func (nc _NNReductionFeatureChannelsSumNodeClass) New() NNReductionFeatureChannelsSumNode {
	rv := objc.Call[NNReductionFeatureChannelsSumNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReductionFeatureChannelsSumNode() NNReductionFeatureChannelsSumNode {
	return NNReductionFeatureChannelsSumNodeClass.New()
}

func (n_ NNReductionFeatureChannelsSumNode) Init() NNReductionFeatureChannelsSumNode {
	rv := objc.Call[NNReductionFeatureChannelsSumNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNReductionFeatureChannelsSumNodeClass) NodeWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsSumNode {
	rv := objc.Call[NNReductionFeatureChannelsSumNode](nc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037425-nodewithsource?language=objc
func NNReductionFeatureChannelsSumNode_NodeWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsSumNode {
	return NNReductionFeatureChannelsSumNodeClass.NodeWithSource(sourceNode)
}

func (n_ NNReductionFeatureChannelsSumNode) InitWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsSumNode {
	rv := objc.Call[NNReductionFeatureChannelsSumNode](n_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnunaryreductionnode/3037424-initwithsource?language=objc
func NewNNReductionFeatureChannelsSumNodeWithSource(sourceNode INNImageNode) NNReductionFeatureChannelsSumNode {
	instance := NNReductionFeatureChannelsSumNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionfeaturechannelssumnode/3037407-weight?language=objc
func (n_ NNReductionFeatureChannelsSumNode) Weight() float64 {
	rv := objc.Call[float64](n_, objc.Sel("weight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreductionfeaturechannelssumnode/3037407-weight?language=objc
func (n_ NNReductionFeatureChannelsSumNode) SetWeight(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setWeight:"), value)
}
