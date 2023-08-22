// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNCrossChannelNormalizationGradientNode] class.
var CNNCrossChannelNormalizationGradientNodeClass = _CNNCrossChannelNormalizationGradientNodeClass{objc.GetClass("MPSCNNCrossChannelNormalizationGradientNode")}

type _CNNCrossChannelNormalizationGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNCrossChannelNormalizationGradientNode] class.
type ICNNCrossChannelNormalizationGradientNode interface {
	INNGradientFilterNode
	KernelSize() uint
}

// A representation of a gradient normalization kernel applied across feature channels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradientnode?language=objc
type CNNCrossChannelNormalizationGradientNode struct {
	NNGradientFilterNode
}

func CNNCrossChannelNormalizationGradientNodeFrom(ptr unsafe.Pointer) CNNCrossChannelNormalizationGradientNode {
	return CNNCrossChannelNormalizationGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (c_ CNNCrossChannelNormalizationGradientNode) InitWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelSize uint) CNNCrossChannelNormalizationGradientNode {
	rv := objc.Call[CNNCrossChannelNormalizationGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelSize:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), kernelSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradientnode/2948043-initwithsourcegradient?language=objc
func NewCNNCrossChannelNormalizationGradientNodeWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelSize uint) CNNCrossChannelNormalizationGradientNode {
	instance := CNNCrossChannelNormalizationGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient, sourceImage, gradientState, kernelSize)
	instance.Autorelease()
	return instance
}

func (cc _CNNCrossChannelNormalizationGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelSize uint) CNNCrossChannelNormalizationGradientNode {
	rv := objc.Call[CNNCrossChannelNormalizationGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelSize:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), kernelSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradientnode/2948032-nodewithsourcegradient?language=objc
func CNNCrossChannelNormalizationGradientNode_NodeWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelSize uint) CNNCrossChannelNormalizationGradientNode {
	return CNNCrossChannelNormalizationGradientNodeClass.NodeWithSourceGradientSourceImageGradientStateKernelSize(sourceGradient, sourceImage, gradientState, kernelSize)
}

func (cc _CNNCrossChannelNormalizationGradientNodeClass) Alloc() CNNCrossChannelNormalizationGradientNode {
	rv := objc.Call[CNNCrossChannelNormalizationGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNCrossChannelNormalizationGradientNode_Alloc() CNNCrossChannelNormalizationGradientNode {
	return CNNCrossChannelNormalizationGradientNodeClass.Alloc()
}

func (cc _CNNCrossChannelNormalizationGradientNodeClass) New() CNNCrossChannelNormalizationGradientNode {
	rv := objc.Call[CNNCrossChannelNormalizationGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNCrossChannelNormalizationGradientNode() CNNCrossChannelNormalizationGradientNode {
	return CNNCrossChannelNormalizationGradientNodeClass.New()
}

func (c_ CNNCrossChannelNormalizationGradientNode) Init() CNNCrossChannelNormalizationGradientNode {
	rv := objc.Call[CNNCrossChannelNormalizationGradientNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnncrosschannelnormalizationgradientnode/2948049-kernelsize?language=objc
func (c_ CNNCrossChannelNormalizationGradientNode) KernelSize() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelSize"))
	return rv
}
