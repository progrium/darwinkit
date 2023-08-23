// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNGroupNormalizationGradientNode] class.
var CNNGroupNormalizationGradientNodeClass = _CNNGroupNormalizationGradientNodeClass{objc.GetClass("MPSCNNGroupNormalizationGradientNode")}

type _CNNGroupNormalizationGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNGroupNormalizationGradientNode] class.
type ICNNGroupNormalizationGradientNode interface {
	INNGradientFilterNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientnode?language=objc
type CNNGroupNormalizationGradientNode struct {
	NNGradientFilterNode
}

func CNNGroupNormalizationGradientNodeFrom(ptr unsafe.Pointer) CNNGroupNormalizationGradientNode {
	return CNNGroupNormalizationGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (c_ CNNGroupNormalizationGradientNode) InitWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNGroupNormalizationGradientNode {
	rv := objc.Call[CNNGroupNormalizationGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientnode/3152569-initwithsourcegradient?language=objc
func NewCNNGroupNormalizationGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNGroupNormalizationGradientNode {
	instance := CNNGroupNormalizationGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientState(sourceGradient, sourceImage, gradientState)
	instance.Autorelease()
	return instance
}

func (cc _CNNGroupNormalizationGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNGroupNormalizationGradientNode {
	rv := objc.Call[CNNGroupNormalizationGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientnode/3152570-nodewithsourcegradient?language=objc
func CNNGroupNormalizationGradientNode_NodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNGroupNormalizationGradientNode {
	return CNNGroupNormalizationGradientNodeClass.NodeWithSourceGradientSourceImageGradientState(sourceGradient, sourceImage, gradientState)
}

func (cc _CNNGroupNormalizationGradientNodeClass) Alloc() CNNGroupNormalizationGradientNode {
	rv := objc.Call[CNNGroupNormalizationGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNGroupNormalizationGradientNode_Alloc() CNNGroupNormalizationGradientNode {
	return CNNGroupNormalizationGradientNodeClass.Alloc()
}

func (cc _CNNGroupNormalizationGradientNodeClass) New() CNNGroupNormalizationGradientNode {
	rv := objc.Call[CNNGroupNormalizationGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNGroupNormalizationGradientNode() CNNGroupNormalizationGradientNode {
	return CNNGroupNormalizationGradientNodeClass.New()
}

func (c_ CNNGroupNormalizationGradientNode) Init() CNNGroupNormalizationGradientNode {
	rv := objc.Call[CNNGroupNormalizationGradientNode](c_, objc.Sel("init"))
	return rv
}
