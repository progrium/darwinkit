// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNDilatedPoolingMaxGradientNode] class.
var CNNDilatedPoolingMaxGradientNodeClass = _CNNDilatedPoolingMaxGradientNodeClass{objc.GetClass("MPSCNNDilatedPoolingMaxGradientNode")}

type _CNNDilatedPoolingMaxGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNDilatedPoolingMaxGradientNode] class.
type ICNNDilatedPoolingMaxGradientNode interface {
	ICNNPoolingGradientNode
	DilationRateY() uint
	DilationRateX() uint
}

// A representation of a gradient dilated max pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxgradientnode?language=objc
type CNNDilatedPoolingMaxGradientNode struct {
	CNNPoolingGradientNode
}

func CNNDilatedPoolingMaxGradientNodeFrom(ptr unsafe.Pointer) CNNDilatedPoolingMaxGradientNode {
	return CNNDilatedPoolingMaxGradientNode{
		CNNPoolingGradientNode: CNNPoolingGradientNodeFrom(ptr),
	}
}

func (c_ CNNDilatedPoolingMaxGradientNode) InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, dilationRateX uint, dilationRateY uint) CNNDilatedPoolingMaxGradientNode {
	rv := objc.Call[CNNDilatedPoolingMaxGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:dilationRateX:dilationRateY:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, dilationRateX, dilationRateY)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxgradientnode/2948026-initwithsourcegradient?language=objc
func NewCNNDilatedPoolingMaxGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, dilationRateX uint, dilationRateY uint) CNNDilatedPoolingMaxGradientNode {
	instance := CNNDilatedPoolingMaxGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY(sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, dilationRateX, dilationRateY)
	instance.Autorelease()
	return instance
}

func (cc _CNNDilatedPoolingMaxGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, dilationRateX uint, dilationRateY uint) CNNDilatedPoolingMaxGradientNode {
	rv := objc.Call[CNNDilatedPoolingMaxGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:dilationRateX:dilationRateY:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, dilationRateX, dilationRateY)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxgradientnode/2948012-nodewithsourcegradient?language=objc
func CNNDilatedPoolingMaxGradientNode_NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, dilationRateX uint, dilationRateY uint) CNNDilatedPoolingMaxGradientNode {
	return CNNDilatedPoolingMaxGradientNodeClass.NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYDilationRateXDilationRateY(sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, dilationRateX, dilationRateY)
}

func (cc _CNNDilatedPoolingMaxGradientNodeClass) Alloc() CNNDilatedPoolingMaxGradientNode {
	rv := objc.Call[CNNDilatedPoolingMaxGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNDilatedPoolingMaxGradientNode_Alloc() CNNDilatedPoolingMaxGradientNode {
	return CNNDilatedPoolingMaxGradientNodeClass.Alloc()
}

func (cc _CNNDilatedPoolingMaxGradientNodeClass) New() CNNDilatedPoolingMaxGradientNode {
	rv := objc.Call[CNNDilatedPoolingMaxGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNDilatedPoolingMaxGradientNode() CNNDilatedPoolingMaxGradientNode {
	return CNNDilatedPoolingMaxGradientNodeClass.New()
}

func (c_ CNNDilatedPoolingMaxGradientNode) Init() CNNDilatedPoolingMaxGradientNode {
	rv := objc.Call[CNNDilatedPoolingMaxGradientNode](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNDilatedPoolingMaxGradientNode) InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNDilatedPoolingMaxGradientNode {
	po7 := objc.WrapAsProtocol("MPSNNPadding", paddingPolicy)
	rv := objc.Call[CNNDilatedPoolingMaxGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, po7)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948011-initwithsourcegradient?language=objc
func NewCNNDilatedPoolingMaxGradientNodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNDilatedPoolingMaxGradientNode {
	instance := CNNDilatedPoolingMaxGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
	instance.Autorelease()
	return instance
}

func (cc _CNNDilatedPoolingMaxGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNDilatedPoolingMaxGradientNode {
	po7 := objc.WrapAsProtocol("MPSNNPadding", paddingPolicy)
	rv := objc.Call[CNNDilatedPoolingMaxGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:paddingPolicy:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, po7)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradientnode/2948045-nodewithsourcegradient?language=objc
func CNNDilatedPoolingMaxGradientNode_NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint, paddingPolicy PNNPadding) CNNDilatedPoolingMaxGradientNode {
	return CNNDilatedPoolingMaxGradientNodeClass.NodeWithSourceGradientSourceImageGradientStateKernelWidthKernelHeightStrideInPixelsXStrideInPixelsYPaddingPolicy(sourceGradient, sourceImage, gradientState, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY, paddingPolicy)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxgradientnode/2948037-dilationratey?language=objc
func (c_ CNNDilatedPoolingMaxGradientNode) DilationRateY() uint {
	rv := objc.Call[uint](c_, objc.Sel("dilationRateY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxgradientnode/2947996-dilationratex?language=objc
func (c_ CNNDilatedPoolingMaxGradientNode) DilationRateX() uint {
	rv := objc.Call[uint](c_, objc.Sel("dilationRateX"))
	return rv
}
