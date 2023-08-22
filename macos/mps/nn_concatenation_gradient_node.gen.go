// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNConcatenationGradientNode] class.
var NNConcatenationGradientNodeClass = _NNConcatenationGradientNodeClass{objc.GetClass("MPSNNConcatenationGradientNode")}

type _NNConcatenationGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [NNConcatenationGradientNode] class.
type INNConcatenationGradientNode interface {
	INNGradientFilterNode
}

// A representation of the results from one or more gradient kernels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnconcatenationgradientnode?language=objc
type NNConcatenationGradientNode struct {
	NNGradientFilterNode
}

func NNConcatenationGradientNodeFrom(ptr unsafe.Pointer) NNConcatenationGradientNode {
	return NNConcatenationGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (n_ NNConcatenationGradientNode) InitWithSourceGradientSourceImageGradientState(gradientSourceNode INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNConcatenationGradientNode {
	rv := objc.Call[NNConcatenationGradientNode](n_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), objc.Ptr(gradientSourceNode), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnconcatenationgradientnode/2951934-initwithsourcegradient?language=objc
func NewNNConcatenationGradientNodeWithSourceGradientSourceImageGradientState(gradientSourceNode INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNConcatenationGradientNode {
	instance := NNConcatenationGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientState(gradientSourceNode, sourceImage, gradientState)
	instance.Autorelease()
	return instance
}

func (nc _NNConcatenationGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(gradientSourceNode INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNConcatenationGradientNode {
	rv := objc.Call[NNConcatenationGradientNode](nc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), objc.Ptr(gradientSourceNode), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnconcatenationgradientnode/2951948-nodewithsourcegradient?language=objc
func NNConcatenationGradientNode_NodeWithSourceGradientSourceImageGradientState(gradientSourceNode INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNConcatenationGradientNode {
	return NNConcatenationGradientNodeClass.NodeWithSourceGradientSourceImageGradientState(gradientSourceNode, sourceImage, gradientState)
}

func (nc _NNConcatenationGradientNodeClass) Alloc() NNConcatenationGradientNode {
	rv := objc.Call[NNConcatenationGradientNode](nc, objc.Sel("alloc"))
	return rv
}

func NNConcatenationGradientNode_Alloc() NNConcatenationGradientNode {
	return NNConcatenationGradientNodeClass.Alloc()
}

func (nc _NNConcatenationGradientNodeClass) New() NNConcatenationGradientNode {
	rv := objc.Call[NNConcatenationGradientNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNConcatenationGradientNode() NNConcatenationGradientNode {
	return NNConcatenationGradientNodeClass.New()
}

func (n_ NNConcatenationGradientNode) Init() NNConcatenationGradientNode {
	rv := objc.Call[NNConcatenationGradientNode](n_, objc.Sel("init"))
	return rv
}
