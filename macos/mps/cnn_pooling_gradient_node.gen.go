// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNPoolingGradientNode] class.
var CNNPoolingGradientNodeClass = _CNNPoolingGradientNodeClass{objc.GetClass("MPSCNNPoolingGradientNode")}

type _CNNPoolingGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNPoolingGradientNode] class.
type ICNNPoolingGradientNode interface {
	INNGradientFilterNode
	StrideInPixelsY() uint
	StrideInPixelsX() uint
	KernelHeight() uint
	KernelWidth() uint
}

// A representation of a gradient pooling kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode?language=objc
type CNNPoolingGradientNode struct {
	NNGradientFilterNode
}

func CNNPoolingGradientNodeFrom(ptr unsafe.Pointer) CNNPoolingGradientNode {
	return CNNPoolingGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (c_ CNNPoolingGradientNode) InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNPoolingGradientNode {
	po7 := objc.WrapAsProtocol("MPSNNPadding", paddingPolicy)
	rv := objc.Call[CNNPoolingGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, po7)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948011-initwithsourcegradient?language=objc
func NewCNNPoolingGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNPoolingGradientNode {
	instance := CNNPoolingGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
	instance.Autorelease()
	return instance
}

func (cc _CNNPoolingGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNPoolingGradientNode {
	po7 := objc.WrapAsProtocol("MPSNNPadding", paddingPolicy)
	rv := objc.Call[CNNPoolingGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, po7)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948045-nodewithsourcegradient?language=objc
func CNNPoolingGradientNode_NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNPoolingGradientNode {
	return CNNPoolingGradientNodeClass.NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
}

func (cc _CNNPoolingGradientNodeClass) Alloc() CNNPoolingGradientNode {
	rv := objc.Call[CNNPoolingGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNPoolingGradientNode_Alloc() CNNPoolingGradientNode {
	return CNNPoolingGradientNodeClass.Alloc()
}

func (cc _CNNPoolingGradientNodeClass) New() CNNPoolingGradientNode {
	rv := objc.Call[CNNPoolingGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNPoolingGradientNode() CNNPoolingGradientNode {
	return CNNPoolingGradientNodeClass.New()
}

func (c_ CNNPoolingGradientNode) Init() CNNPoolingGradientNode {
	rv := objc.Call[CNNPoolingGradientNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948048-strideinpixelsy?language=objc
func (c_ CNNPoolingGradientNode) StrideInPixelsY() uint {
	rv := objc.Call[uint](c_, objc.Sel("strideInPixelsY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948018-strideinpixelsx?language=objc
func (c_ CNNPoolingGradientNode) StrideInPixelsX() uint {
	rv := objc.Call[uint](c_, objc.Sel("strideInPixelsX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2947992-kernelheight?language=objc
func (c_ CNNPoolingGradientNode) KernelHeight() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelHeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948034-kernelwidth?language=objc
func (c_ CNNPoolingGradientNode) KernelWidth() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelWidth"))
	return rv
}
