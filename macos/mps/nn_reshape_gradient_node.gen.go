// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReshapeGradientNode] class.
var NNReshapeGradientNodeClass = _NNReshapeGradientNodeClass{objc.GetClass("MPSNNReshapeGradientNode")}

type _NNReshapeGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [NNReshapeGradientNode] class.
type INNReshapeGradientNode interface {
	INNGradientFilterNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshapegradientnode?language=objc
type NNReshapeGradientNode struct {
	NNGradientFilterNode
}

func NNReshapeGradientNodeFrom(ptr unsafe.Pointer) NNReshapeGradientNode {
	return NNReshapeGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (n_ NNReshapeGradientNode) InitWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNReshapeGradientNode {
	rv := objc.Call[NNReshapeGradientNode](n_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshapegradientnode/3037417-initwithsourcegradient?language=objc
func NewNNReshapeGradientNodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNReshapeGradientNode {
	instance := NNReshapeGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientState(sourceGradient, sourceImage, gradientState)
	instance.Autorelease()
	return instance
}

func (nc _NNReshapeGradientNodeClass) NodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNReshapeGradientNode {
	rv := objc.Call[NNReshapeGradientNode](nc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshapegradientnode/3037418-nodewithsourcegradient?language=objc
func NNReshapeGradientNode_NodeWithSourceGradientSourceImageGradientState(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode) NNReshapeGradientNode {
	return NNReshapeGradientNodeClass.NodeWithSourceGradientSourceImageGradientState(sourceGradient, sourceImage, gradientState)
}

func (nc _NNReshapeGradientNodeClass) Alloc() NNReshapeGradientNode {
	rv := objc.Call[NNReshapeGradientNode](nc, objc.Sel("alloc"))
	return rv
}

func NNReshapeGradientNode_Alloc() NNReshapeGradientNode {
	return NNReshapeGradientNodeClass.Alloc()
}

func (nc _NNReshapeGradientNodeClass) New() NNReshapeGradientNode {
	rv := objc.Call[NNReshapeGradientNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReshapeGradientNode() NNReshapeGradientNode {
	return NNReshapeGradientNodeClass.New()
}

func (n_ NNReshapeGradientNode) Init() NNReshapeGradientNode {
	rv := objc.Call[NNReshapeGradientNode](n_, objc.Sel("init"))
	return rv
}
