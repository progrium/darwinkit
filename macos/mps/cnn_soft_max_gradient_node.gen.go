// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNSoftMaxGradientNode] class.
var CNNSoftMaxGradientNodeClass = _CNNSoftMaxGradientNodeClass{objc.GetClass("MPSCNNSoftMaxGradientNode")}

type _CNNSoftMaxGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNSoftMaxGradientNode] class.
type ICNNSoftMaxGradientNode interface {
	INNGradientFilterNode
}

// A representation of a gradient softmax filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmaxgradientnode?language=objc
type CNNSoftMaxGradientNode struct {
	NNGradientFilterNode
}

func CNNSoftMaxGradientNodeFrom(ptr unsafe.Pointer) CNNSoftMaxGradientNode {
	return CNNSoftMaxGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (c_ CNNSoftMaxGradientNode) InitWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNSoftMaxGradientNode {
	rv := objc.Call[CNNSoftMaxGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmaxgradientnode/2948039-initwithsourcegradient?language=objc
func NewCNNSoftMaxGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNSoftMaxGradientNode {
	instance := CNNSoftMaxGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientState(sourceGradient, sourceImage, gradientState)
	instance.Autorelease()
	return instance
}

func (cc _CNNSoftMaxGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNSoftMaxGradientNode {
	rv := objc.Call[CNNSoftMaxGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnsoftmaxgradientnode/2947995-nodewithsourcegradient?language=objc
func CNNSoftMaxGradientNode_NodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNSoftMaxGradientNode {
	return CNNSoftMaxGradientNodeClass.NodeWithSourceGradientSourceImageGradientState(sourceGradient, sourceImage, gradientState)
}

func (cc _CNNSoftMaxGradientNodeClass) Alloc() CNNSoftMaxGradientNode {
	rv := objc.Call[CNNSoftMaxGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNSoftMaxGradientNode_Alloc() CNNSoftMaxGradientNode {
	return CNNSoftMaxGradientNodeClass.Alloc()
}

func (cc _CNNSoftMaxGradientNodeClass) New() CNNSoftMaxGradientNode {
	rv := objc.Call[CNNSoftMaxGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNSoftMaxGradientNode() CNNSoftMaxGradientNode {
	return CNNSoftMaxGradientNodeClass.New()
}

func (c_ CNNSoftMaxGradientNode) Init() CNNSoftMaxGradientNode {
	rv := objc.Call[CNNSoftMaxGradientNode](c_, objc.Sel("init"))
	return rv
}
