// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNPoolingAverageGradientNode] class.
var CNNPoolingAverageGradientNodeClass = _CNNPoolingAverageGradientNodeClass{objc.GetClass("MPSCNNPoolingAverageGradientNode")}

type _CNNPoolingAverageGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNPoolingAverageGradientNode] class.
type ICNNPoolingAverageGradientNode interface {
	ICNNPoolingGradientNode
}

// A representation of a gradient average pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradientnode?language=objc
type CNNPoolingAverageGradientNode struct {
	CNNPoolingGradientNode
}

func CNNPoolingAverageGradientNodeFrom(ptr unsafe.Pointer) CNNPoolingAverageGradientNode {
	return CNNPoolingAverageGradientNode{
		CNNPoolingGradientNode: CNNPoolingGradientNodeFrom(ptr),
	}
}

func (cc _CNNPoolingAverageGradientNodeClass) Alloc() CNNPoolingAverageGradientNode {
	rv := objc.Call[CNNPoolingAverageGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNPoolingAverageGradientNode_Alloc() CNNPoolingAverageGradientNode {
	return CNNPoolingAverageGradientNodeClass.Alloc()
}

func (cc _CNNPoolingAverageGradientNodeClass) New() CNNPoolingAverageGradientNode {
	rv := objc.Call[CNNPoolingAverageGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNPoolingAverageGradientNode() CNNPoolingAverageGradientNode {
	return CNNPoolingAverageGradientNodeClass.New()
}

func (c_ CNNPoolingAverageGradientNode) Init() CNNPoolingAverageGradientNode {
	rv := objc.Call[CNNPoolingAverageGradientNode](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNPoolingAverageGradientNode) InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNPoolingAverageGradientNode {
	po7 := objc.WrapAsProtocol("MPSNNPadding", paddingPolicy)
	rv := objc.Call[CNNPoolingAverageGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, po7)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948011-initwithsourcegradient?language=objc
func NewCNNPoolingAverageGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNPoolingAverageGradientNode {
	instance := CNNPoolingAverageGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
	instance.Autorelease()
	return instance
}

func (cc _CNNPoolingAverageGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNPoolingAverageGradientNode {
	po7 := objc.WrapAsProtocol("MPSNNPadding", paddingPolicy)
	rv := objc.Call[CNNPoolingAverageGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, po7)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948045-nodewithsourcegradient?language=objc
func CNNPoolingAverageGradientNode_NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNPoolingAverageGradientNode {
	return CNNPoolingAverageGradientNodeClass.NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
}
