// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNBatchNormalizationGradientNode] class.
var CNNBatchNormalizationGradientNodeClass = _CNNBatchNormalizationGradientNodeClass{objc.GetClass("MPSCNNBatchNormalizationGradientNode")}

type _CNNBatchNormalizationGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNBatchNormalizationGradientNode] class.
type ICNNBatchNormalizationGradientNode interface {
	INNGradientFilterNode
}

// A representation of a gradient batch normalization kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradientnode?language=objc
type CNNBatchNormalizationGradientNode struct {
	NNGradientFilterNode
}

func CNNBatchNormalizationGradientNodeFrom(ptr unsafe.Pointer) CNNBatchNormalizationGradientNode {
	return CNNBatchNormalizationGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (c_ CNNBatchNormalizationGradientNode) InitWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNBatchNormalizationGradientNode {
	rv := objc.Call[CNNBatchNormalizationGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradientnode/2953939-initwithsourcegradient?language=objc
func NewCNNBatchNormalizationGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNBatchNormalizationGradientNode {
	instance := CNNBatchNormalizationGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientState(sourceGradient, sourceImage, gradientState)
	instance.Autorelease()
	return instance
}

func (cc _CNNBatchNormalizationGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNBatchNormalizationGradientNode {
	rv := objc.Call[CNNBatchNormalizationGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradientnode/2953943-nodewithsourcegradient?language=objc
func CNNBatchNormalizationGradientNode_NodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNBatchNormalizationGradientNode {
	return CNNBatchNormalizationGradientNodeClass.NodeWithSourceGradientSourceImageGradientState(sourceGradient, sourceImage, gradientState)
}

func (cc _CNNBatchNormalizationGradientNodeClass) Alloc() CNNBatchNormalizationGradientNode {
	rv := objc.Call[CNNBatchNormalizationGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNBatchNormalizationGradientNode_Alloc() CNNBatchNormalizationGradientNode {
	return CNNBatchNormalizationGradientNodeClass.Alloc()
}

func (cc _CNNBatchNormalizationGradientNodeClass) New() CNNBatchNormalizationGradientNode {
	rv := objc.Call[CNNBatchNormalizationGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNBatchNormalizationGradientNode() CNNBatchNormalizationGradientNode {
	return CNNBatchNormalizationGradientNodeClass.New()
}

func (c_ CNNBatchNormalizationGradientNode) Init() CNNBatchNormalizationGradientNode {
	rv := objc.Call[CNNBatchNormalizationGradientNode](c_, objc.Sel("init"))
	return rv
}
