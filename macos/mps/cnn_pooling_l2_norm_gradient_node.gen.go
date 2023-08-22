// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNPoolingL2NormGradientNode] class.
var CNNPoolingL2NormGradientNodeClass = _CNNPoolingL2NormGradientNodeClass{objc.GetClass("MPSCNNPoolingL2NormGradientNode")}

type _CNNPoolingL2NormGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNPoolingL2NormGradientNode] class.
type ICNNPoolingL2NormGradientNode interface {
	ICNNPoolingGradientNode
}

// A representation of a gradient L2-norm pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingl2normgradientnode?language=objc
type CNNPoolingL2NormGradientNode struct {
	CNNPoolingGradientNode
}

func CNNPoolingL2NormGradientNodeFrom(ptr unsafe.Pointer) CNNPoolingL2NormGradientNode {
	return CNNPoolingL2NormGradientNode{
		CNNPoolingGradientNode: CNNPoolingGradientNodeFrom(ptr),
	}
}

func (cc _CNNPoolingL2NormGradientNodeClass) Alloc() CNNPoolingL2NormGradientNode {
	rv := objc.Call[CNNPoolingL2NormGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNPoolingL2NormGradientNode_Alloc() CNNPoolingL2NormGradientNode {
	return CNNPoolingL2NormGradientNodeClass.Alloc()
}

func (cc _CNNPoolingL2NormGradientNodeClass) New() CNNPoolingL2NormGradientNode {
	rv := objc.Call[CNNPoolingL2NormGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNPoolingL2NormGradientNode() CNNPoolingL2NormGradientNode {
	return CNNPoolingL2NormGradientNodeClass.New()
}

func (c_ CNNPoolingL2NormGradientNode) Init() CNNPoolingL2NormGradientNode {
	rv := objc.Call[CNNPoolingL2NormGradientNode](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNPoolingL2NormGradientNode) InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNPoolingL2NormGradientNode {
	po7 := objc.WrapAsProtocol("MPSNNPadding", paddingPolicy)
	rv := objc.Call[CNNPoolingL2NormGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, po7)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948011-initwithsourcegradient?language=objc
func NewCNNPoolingL2NormGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNPoolingL2NormGradientNode {
	instance := CNNPoolingL2NormGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
	instance.Autorelease()
	return instance
}

func (cc _CNNPoolingL2NormGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNPoolingL2NormGradientNode {
	po7 := objc.WrapAsProtocol("MPSNNPadding", paddingPolicy)
	rv := objc.Call[CNNPoolingL2NormGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, po7)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948045-nodewithsourcegradient?language=objc
func CNNPoolingL2NormGradientNode_NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNPoolingL2NormGradientNode {
	return CNNPoolingL2NormGradientNodeClass.NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
}
