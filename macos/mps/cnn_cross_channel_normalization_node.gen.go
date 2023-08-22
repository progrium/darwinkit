// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNCrossChannelNormalizationNode] class.
var CNNCrossChannelNormalizationNodeClass = _CNNCrossChannelNormalizationNodeClass{objc.GetClass("MPSCNNCrossChannelNormalizationNode")}

type _CNNCrossChannelNormalizationNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNCrossChannelNormalizationNode] class.
type ICNNCrossChannelNormalizationNode interface {
	ICNNNormalizationNode
	KernelSizeInFeatureChannels() uint
	SetKernelSizeInFeatureChannels(value uint)
}

// A representation of a normalization kernel across feature channels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationnode?language=objc
type CNNCrossChannelNormalizationNode struct {
	CNNNormalizationNode
}

func CNNCrossChannelNormalizationNodeFrom(ptr unsafe.Pointer) CNNCrossChannelNormalizationNode {
	return CNNCrossChannelNormalizationNode{
		CNNNormalizationNode: CNNNormalizationNodeFrom(ptr),
	}
}

func (cc _CNNCrossChannelNormalizationNodeClass) NodeWithSourceKernelSize(sourceNode INNImageNode, kernelSize uint) CNNCrossChannelNormalizationNode {
	rv := objc.Call[CNNCrossChannelNormalizationNode](cc, objc.Sel("nodeWithSource:kernelSize:"), objc.Ptr(sourceNode), kernelSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationnode/2866476-nodewithsource?language=objc
func CNNCrossChannelNormalizationNode_NodeWithSourceKernelSize(sourceNode INNImageNode, kernelSize uint) CNNCrossChannelNormalizationNode {
	return CNNCrossChannelNormalizationNodeClass.NodeWithSourceKernelSize(sourceNode, kernelSize)
}

func (c_ CNNCrossChannelNormalizationNode) InitWithSource(sourceNode INNImageNode) CNNCrossChannelNormalizationNode {
	rv := objc.Call[CNNCrossChannelNormalizationNode](c_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationnode/2866459-initwithsource?language=objc
func NewCNNCrossChannelNormalizationNodeWithSource(sourceNode INNImageNode) CNNCrossChannelNormalizationNode {
	instance := CNNCrossChannelNormalizationNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (cc _CNNCrossChannelNormalizationNodeClass) Alloc() CNNCrossChannelNormalizationNode {
	rv := objc.Call[CNNCrossChannelNormalizationNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNCrossChannelNormalizationNode_Alloc() CNNCrossChannelNormalizationNode {
	return CNNCrossChannelNormalizationNodeClass.Alloc()
}

func (cc _CNNCrossChannelNormalizationNodeClass) New() CNNCrossChannelNormalizationNode {
	rv := objc.Call[CNNCrossChannelNormalizationNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNCrossChannelNormalizationNode() CNNCrossChannelNormalizationNode {
	return CNNCrossChannelNormalizationNodeClass.New()
}

func (c_ CNNCrossChannelNormalizationNode) Init() CNNCrossChannelNormalizationNode {
	rv := objc.Call[CNNCrossChannelNormalizationNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNCrossChannelNormalizationNodeClass) NodeWithSource(sourceNode INNImageNode) CNNCrossChannelNormalizationNode {
	rv := objc.Call[CNNCrossChannelNormalizationNode](cc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866460-nodewithsource?language=objc
func CNNCrossChannelNormalizationNode_NodeWithSource(sourceNode INNImageNode) CNNCrossChannelNormalizationNode {
	return CNNCrossChannelNormalizationNodeClass.NodeWithSource(sourceNode)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationnode/2866419-kernelsizeinfeaturechannels?language=objc
func (c_ CNNCrossChannelNormalizationNode) KernelSizeInFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelSizeInFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationnode/2866419-kernelsizeinfeaturechannels?language=objc
func (c_ CNNCrossChannelNormalizationNode) SetKernelSizeInFeatureChannels(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setKernelSizeInFeatureChannels:"), value)
}
