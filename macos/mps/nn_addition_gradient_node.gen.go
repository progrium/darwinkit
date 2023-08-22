// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNAdditionGradientNode] class.
var NNAdditionGradientNodeClass = _NNAdditionGradientNodeClass{objc.GetClass("MPSNNAdditionGradientNode")}

type _NNAdditionGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [NNAdditionGradientNode] class.
type INNAdditionGradientNode interface {
	INNArithmeticGradientNode
}

// A representation of a gradient addition operator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnadditiongradientnode?language=objc
type NNAdditionGradientNode struct {
	NNArithmeticGradientNode
}

func NNAdditionGradientNodeFrom(ptr unsafe.Pointer) NNAdditionGradientNode {
	return NNAdditionGradientNode{
		NNArithmeticGradientNode: NNArithmeticGradientNodeFrom(ptr),
	}
}

func (nc _NNAdditionGradientNodeClass) Alloc() NNAdditionGradientNode {
	rv := objc.Call[NNAdditionGradientNode](nc, objc.Sel("alloc"))
	return rv
}

func NNAdditionGradientNode_Alloc() NNAdditionGradientNode {
	return NNAdditionGradientNodeClass.Alloc()
}

func (nc _NNAdditionGradientNodeClass) New() NNAdditionGradientNode {
	rv := objc.Call[NNAdditionGradientNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNAdditionGradientNode() NNAdditionGradientNode {
	return NNAdditionGradientNodeClass.New()
}

func (n_ NNAdditionGradientNode) Init() NNAdditionGradientNode {
	rv := objc.Call[NNAdditionGradientNode](n_, objc.Sel("init"))
	return rv
}

func (n_ NNAdditionGradientNode) InitWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages []INNImageNode, filter INNFilterNode, isSecondarySourceFilter bool) NNAdditionGradientNode {
	rv := objc.Call[NNAdditionGradientNode](n_, objc.Sel("initWithGradientImages:forwardFilter:isSecondarySourceFilter:"), gradientImages, objc.Ptr(filter), isSecondarySourceFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952980-initwithgradientimages?language=objc
func NewNNAdditionGradientNodeWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages []INNImageNode, filter INNFilterNode, isSecondarySourceFilter bool) NNAdditionGradientNode {
	instance := NNAdditionGradientNodeClass.Alloc().InitWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages, filter, isSecondarySourceFilter)
	instance.Autorelease()
	return instance
}

func (n_ NNAdditionGradientNode) InitWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNAdditionGradientNode {
	rv := objc.Call[NNAdditionGradientNode](n_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), isSecondarySourceFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2956166-initwithsourcegradient?language=objc
func NewNNAdditionGradientNodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNAdditionGradientNode {
	instance := NNAdditionGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
	instance.Autorelease()
	return instance
}

func (nc _NNAdditionGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNAdditionGradientNode {
	rv := objc.Call[NNAdditionGradientNode](nc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), isSecondarySourceFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2956167-nodewithsourcegradient?language=objc
func NNAdditionGradientNode_NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNAdditionGradientNode {
	return NNAdditionGradientNodeClass.NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
}
