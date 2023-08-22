// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNLogSoftMaxGradientNode] class.
var CNNLogSoftMaxGradientNodeClass = _CNNLogSoftMaxGradientNodeClass{objc.GetClass("MPSCNNLogSoftMaxGradientNode")}

type _CNNLogSoftMaxGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNLogSoftMaxGradientNode] class.
type ICNNLogSoftMaxGradientNode interface {
	INNGradientFilterNode
}

// A representation of a gradient logarithmic softmax filter kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxgradientnode?language=objc
type CNNLogSoftMaxGradientNode struct {
	NNGradientFilterNode
}

func CNNLogSoftMaxGradientNodeFrom(ptr unsafe.Pointer) CNNLogSoftMaxGradientNode {
	return CNNLogSoftMaxGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (c_ CNNLogSoftMaxGradientNode) InitWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNLogSoftMaxGradientNode {
	rv := objc.Call[CNNLogSoftMaxGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxgradientnode/2947971-initwithsourcegradient?language=objc
func NewCNNLogSoftMaxGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNLogSoftMaxGradientNode {
	instance := CNNLogSoftMaxGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientState(sourceGradient, sourceImage, gradientState)
	instance.Autorelease()
	return instance
}

func (cc _CNNLogSoftMaxGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNLogSoftMaxGradientNode {
	rv := objc.Call[CNNLogSoftMaxGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxgradientnode/2947974-nodewithsourcegradient?language=objc
func CNNLogSoftMaxGradientNode_NodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) CNNLogSoftMaxGradientNode {
	return CNNLogSoftMaxGradientNodeClass.NodeWithSourceGradientSourceImageGradientState(sourceGradient, sourceImage, gradientState)
}

func (cc _CNNLogSoftMaxGradientNodeClass) Alloc() CNNLogSoftMaxGradientNode {
	rv := objc.Call[CNNLogSoftMaxGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNLogSoftMaxGradientNode_Alloc() CNNLogSoftMaxGradientNode {
	return CNNLogSoftMaxGradientNodeClass.Alloc()
}

func (cc _CNNLogSoftMaxGradientNodeClass) New() CNNLogSoftMaxGradientNode {
	rv := objc.Call[CNNLogSoftMaxGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNLogSoftMaxGradientNode() CNNLogSoftMaxGradientNode {
	return CNNLogSoftMaxGradientNodeClass.New()
}

func (c_ CNNLogSoftMaxGradientNode) Init() CNNLogSoftMaxGradientNode {
	rv := objc.Call[CNNLogSoftMaxGradientNode](c_, objc.Sel("init"))
	return rv
}
