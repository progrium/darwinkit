// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNSubtractionGradientNode] class.
var NNSubtractionGradientNodeClass = _NNSubtractionGradientNodeClass{objc.GetClass("MPSNNSubtractionGradientNode")}

type _NNSubtractionGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [NNSubtractionGradientNode] class.
type INNSubtractionGradientNode interface {
	INNArithmeticGradientNode
}

// A representation of a gradient subtraction operator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnsubtractiongradientnode?language=objc
type NNSubtractionGradientNode struct {
	NNArithmeticGradientNode
}

func NNSubtractionGradientNodeFrom(ptr unsafe.Pointer) NNSubtractionGradientNode {
	return NNSubtractionGradientNode{
		NNArithmeticGradientNode: NNArithmeticGradientNodeFrom(ptr),
	}
}

func (nc _NNSubtractionGradientNodeClass) Alloc() NNSubtractionGradientNode {
	rv := objc.Call[NNSubtractionGradientNode](nc, objc.Sel("alloc"))
	return rv
}

func NNSubtractionGradientNode_Alloc() NNSubtractionGradientNode {
	return NNSubtractionGradientNodeClass.Alloc()
}

func (nc _NNSubtractionGradientNodeClass) New() NNSubtractionGradientNode {
	rv := objc.Call[NNSubtractionGradientNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNSubtractionGradientNode() NNSubtractionGradientNode {
	return NNSubtractionGradientNodeClass.New()
}

func (n_ NNSubtractionGradientNode) Init() NNSubtractionGradientNode {
	rv := objc.Call[NNSubtractionGradientNode](n_, objc.Sel("init"))
	return rv
}

func (n_ NNSubtractionGradientNode) InitWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages []INNImageNode, filter INNFilterNode, isSecondarySourceFilter bool) NNSubtractionGradientNode {
	rv := objc.Call[NNSubtractionGradientNode](n_, objc.Sel("initWithGradientImages:forwardFilter:isSecondarySourceFilter:"), gradientImages, objc.Ptr(filter), isSecondarySourceFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2952980-initwithgradientimages?language=objc
func NewNNSubtractionGradientNodeWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages []INNImageNode, filter INNFilterNode, isSecondarySourceFilter bool) NNSubtractionGradientNode {
	instance := NNSubtractionGradientNodeClass.Alloc().InitWithGradientImagesForwardFilterIsSecondarySourceFilter(gradientImages, filter, isSecondarySourceFilter)
	instance.Autorelease()
	return instance
}

func (n_ NNSubtractionGradientNode) InitWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNSubtractionGradientNode {
	rv := objc.Call[NNSubtractionGradientNode](n_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), isSecondarySourceFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2956166-initwithsourcegradient?language=objc
func NewNNSubtractionGradientNodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNSubtractionGradientNode {
	instance := NNSubtractionGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
	instance.Autorelease()
	return instance
}

func (nc _NNSubtractionGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNSubtractionGradientNode {
	rv := objc.Call[NNSubtractionGradientNode](nc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:isSecondarySourceFilter:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), isSecondarySourceFilter)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientnode/2956167-nodewithsourcegradient?language=objc
func NNSubtractionGradientNode_NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNBinaryGradientStateNode, isSecondarySourceFilter bool) NNSubtractionGradientNode {
	return NNSubtractionGradientNodeClass.NodeWithSourceGradientSourceImageGradientStateIsSecondarySourceFilter(sourceGradient, sourceImage, gradientState, isSecondarySourceFilter)
}
