// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNInstanceNormalizationGradientNode] class.
var CNNInstanceNormalizationGradientNodeClass = _CNNInstanceNormalizationGradientNodeClass{objc.GetClass("MPSCNNInstanceNormalizationGradientNode")}

type _CNNInstanceNormalizationGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNInstanceNormalizationGradientNode] class.
type ICNNInstanceNormalizationGradientNode interface {
	INNGradientFilterNode
}

// A representation of a gradient instance normalization kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientnode?language=objc
type CNNInstanceNormalizationGradientNode struct {
	NNGradientFilterNode
}

func CNNInstanceNormalizationGradientNodeFrom(ptr unsafe.Pointer) CNNInstanceNormalizationGradientNode {
	return CNNInstanceNormalizationGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (c_ CNNInstanceNormalizationGradientNode) InitWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNInstanceNormalizationGradientNode {
	rv := objc.Call[CNNInstanceNormalizationGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientnode/2951954-initwithsourcegradient?language=objc
func NewCNNInstanceNormalizationGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNInstanceNormalizationGradientNode {
	instance := CNNInstanceNormalizationGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientState(sourceGradient, sourceImage, gradientState)
	instance.Autorelease()
	return instance
}

func (cc _CNNInstanceNormalizationGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNInstanceNormalizationGradientNode {
	rv := objc.Call[CNNInstanceNormalizationGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientnode/2951932-nodewithsourcegradient?language=objc
func CNNInstanceNormalizationGradientNode_NodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNInstanceNormalizationGradientNode {
	return CNNInstanceNormalizationGradientNodeClass.NodeWithSourceGradientSourceImageGradientState(sourceGradient, sourceImage, gradientState)
}

func (cc _CNNInstanceNormalizationGradientNodeClass) Alloc() CNNInstanceNormalizationGradientNode {
	rv := objc.Call[CNNInstanceNormalizationGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNInstanceNormalizationGradientNode_Alloc() CNNInstanceNormalizationGradientNode {
	return CNNInstanceNormalizationGradientNodeClass.Alloc()
}

func (cc _CNNInstanceNormalizationGradientNodeClass) New() CNNInstanceNormalizationGradientNode {
	rv := objc.Call[CNNInstanceNormalizationGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNInstanceNormalizationGradientNode() CNNInstanceNormalizationGradientNode {
	return CNNInstanceNormalizationGradientNodeClass.New()
}

func (c_ CNNInstanceNormalizationGradientNode) Init() CNNInstanceNormalizationGradientNode {
	rv := objc.Call[CNNInstanceNormalizationGradientNode](c_, objc.Sel("init"))
	return rv
}
