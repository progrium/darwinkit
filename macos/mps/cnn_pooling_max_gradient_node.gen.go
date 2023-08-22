// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNPoolingMaxGradientNode] class.
var CNNPoolingMaxGradientNodeClass = _CNNPoolingMaxGradientNodeClass{objc.GetClass("MPSCNNPoolingMaxGradientNode")}

type _CNNPoolingMaxGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNPoolingMaxGradientNode] class.
type ICNNPoolingMaxGradientNode interface {
	ICNNPoolingGradientNode
}

// A representation of a gradient max pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingmaxgradientnode?language=objc
type CNNPoolingMaxGradientNode struct {
	CNNPoolingGradientNode
}

func CNNPoolingMaxGradientNodeFrom(ptr unsafe.Pointer) CNNPoolingMaxGradientNode {
	return CNNPoolingMaxGradientNode{
		CNNPoolingGradientNode: CNNPoolingGradientNodeFrom(ptr),
	}
}

func (cc _CNNPoolingMaxGradientNodeClass) Alloc() CNNPoolingMaxGradientNode {
	rv := objc.Call[CNNPoolingMaxGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNPoolingMaxGradientNode_Alloc() CNNPoolingMaxGradientNode {
	return CNNPoolingMaxGradientNodeClass.Alloc()
}

func (cc _CNNPoolingMaxGradientNodeClass) New() CNNPoolingMaxGradientNode {
	rv := objc.Call[CNNPoolingMaxGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNPoolingMaxGradientNode() CNNPoolingMaxGradientNode {
	return CNNPoolingMaxGradientNodeClass.New()
}

func (c_ CNNPoolingMaxGradientNode) Init() CNNPoolingMaxGradientNode {
	rv := objc.Call[CNNPoolingMaxGradientNode](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNPoolingMaxGradientNode) InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNPoolingMaxGradientNode {
	po7 := objc.WrapAsProtocol("MPSNNPadding", paddingPolicy)
	rv := objc.Call[CNNPoolingMaxGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, po7)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948011-initwithsourcegradient?language=objc
func NewCNNPoolingMaxGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNPoolingMaxGradientNode {
	instance := CNNPoolingMaxGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
	instance.Autorelease()
	return instance
}

func (cc _CNNPoolingMaxGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNPoolingMaxGradientNode {
	po7 := objc.WrapAsProtocol("MPSNNPadding", paddingPolicy)
	rv := objc.Call[CNNPoolingMaxGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, po7)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948045-nodewithsourcegradient?language=objc
func CNNPoolingMaxGradientNode_NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNPoolingMaxGradientNode {
	return CNNPoolingMaxGradientNodeClass.NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
}
