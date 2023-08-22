// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNMultiplicationGradientNode] class.
var NNMultiplicationGradientNodeClass = _NNMultiplicationGradientNodeClass{objc.GetClass("MPSNNMultiplicationGradientNode")}

type _NNMultiplicationGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [NNMultiplicationGradientNode] class.
type INNMultiplicationGradientNode interface {
	INNArithmeticGradientNode
}

// A representation of a gradient multiplication operator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnmultiplicationgradientnode?language=objc
type NNMultiplicationGradientNode struct {
	NNArithmeticGradientNode
}

func NNMultiplicationGradientNodeFrom(ptr unsafe.Pointer) NNMultiplicationGradientNode {
	return NNMultiplicationGradientNode{
		NNArithmeticGradientNode: NNArithmeticGradientNodeFrom(ptr),
	}
}

func (nc _NNMultiplicationGradientNodeClass) Alloc() NNMultiplicationGradientNode {
	rv := objc.Call[NNMultiplicationGradientNode](nc, objc.Sel("alloc"))
	return rv
}

func NNMultiplicationGradientNode_Alloc() NNMultiplicationGradientNode {
	return NNMultiplicationGradientNodeClass.Alloc()
}

func (nc _NNMultiplicationGradientNodeClass) New() NNMultiplicationGradientNode {
	rv := objc.Call[NNMultiplicationGradientNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNMultiplicationGradientNode() NNMultiplicationGradientNode {
	return NNMultiplicationGradientNodeClass.New()
}

func (n_ NNMultiplicationGradientNode) Init() NNMultiplicationGradientNode {
	rv := objc.Call[NNMultiplicationGradientNode](n_, objc.Sel("init"))
	return rv
}

func (n_ NNMultiplicationGradientNode) InitWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages []INNImageNode, filter INNFilterNode, isSecondarySourceFilter bool) NNMultiplicationGradientNode {
	rv := objc.Call[NNMultiplicationGradientNode](n_, objc.Sel("initWithGradientImages:forwardFilter:isSecondarySourceFilter:"), gradientImages, objc.Ptr(filter), isSecondarySourceFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952980-initwithgradientimages?language=objc
func NewNNMultiplicationGradientNodeWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages []INNImageNode, filter INNFilterNode, isSecondarySourceFilter bool) NNMultiplicationGradientNode {
	instance := NNMultiplicationGradientNodeClass.Alloc().InitWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages, filter, isSecondarySourceFilter)
	instance.Autorelease()
	return instance
}

func (n_ NNMultiplicationGradientNode) InitWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNMultiplicationGradientNode {
	rv := objc.Call[NNMultiplicationGradientNode](n_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), isSecondarySourceFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2956166-initwithsourcegradient?language=objc
func NewNNMultiplicationGradientNodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNMultiplicationGradientNode {
	instance := NNMultiplicationGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
	instance.Autorelease()
	return instance
}

func (nc _NNMultiplicationGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNMultiplicationGradientNode {
	rv := objc.Call[NNMultiplicationGradientNode](nc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), isSecondarySourceFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2956167-nodewithsourcegradient?language=objc
func NNMultiplicationGradientNode_NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNMultiplicationGradientNode {
	return NNMultiplicationGradientNodeClass.NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
}
