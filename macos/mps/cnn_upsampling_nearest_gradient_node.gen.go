// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNUpsamplingNearestGradientNode] class.
var CNNUpsamplingNearestGradientNodeClass = _CNNUpsamplingNearestGradientNodeClass{objc.GetClass("MPSCNNUpsamplingNearestGradientNode")}

type _CNNUpsamplingNearestGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNUpsamplingNearestGradientNode] class.
type ICNNUpsamplingNearestGradientNode interface {
	INNGradientFilterNode
	ScaleFactorY() float64
	ScaleFactorX() float64
}

// A representation of a gradient nearest spatial upsampling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradientnode?language=objc
type CNNUpsamplingNearestGradientNode struct {
	NNGradientFilterNode
}

func CNNUpsamplingNearestGradientNodeFrom(ptr unsafe.Pointer) CNNUpsamplingNearestGradientNode {
	return CNNUpsamplingNearestGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (c_ CNNUpsamplingNearestGradientNode) InitWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, scaleFactorX float64, scaleFactorY float64) CNNUpsamplingNearestGradientNode {
	rv := objc.Call[CNNUpsamplingNearestGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), scaleFactorX, scaleFactorY)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradientnode/2947983-initwithsourcegradient?language=objc
func NewCNNUpsamplingNearestGradientNodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, scaleFactorX float64, scaleFactorY float64) CNNUpsamplingNearestGradientNode {
	instance := CNNUpsamplingNearestGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient, sourceImage, gradientState, scaleFactorX, scaleFactorY)
	instance.Autorelease()
	return instance
}

func (cc _CNNUpsamplingNearestGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, scaleFactorX float64, scaleFactorY float64) CNNUpsamplingNearestGradientNode {
	rv := objc.Call[CNNUpsamplingNearestGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:scaleFactorX:scaleFactorY:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), scaleFactorX, scaleFactorY)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradientnode/2948021-nodewithsourcegradient?language=objc
func CNNUpsamplingNearestGradientNode_NodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, scaleFactorX float64, scaleFactorY float64) CNNUpsamplingNearestGradientNode {
	return CNNUpsamplingNearestGradientNodeClass.NodeWithSourceGradientSourceImageGradientStateScaleFactorXScaleFactorY(sourceGradient, sourceImage, gradientState, scaleFactorX, scaleFactorY)
}

func (cc _CNNUpsamplingNearestGradientNodeClass) Alloc() CNNUpsamplingNearestGradientNode {
	rv := objc.Call[CNNUpsamplingNearestGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNUpsamplingNearestGradientNode_Alloc() CNNUpsamplingNearestGradientNode {
	return CNNUpsamplingNearestGradientNodeClass.Alloc()
}

func (cc _CNNUpsamplingNearestGradientNodeClass) New() CNNUpsamplingNearestGradientNode {
	rv := objc.Call[CNNUpsamplingNearestGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNUpsamplingNearestGradientNode() CNNUpsamplingNearestGradientNode {
	return CNNUpsamplingNearestGradientNodeClass.New()
}

func (c_ CNNUpsamplingNearestGradientNode) Init() CNNUpsamplingNearestGradientNode {
	rv := objc.Call[CNNUpsamplingNearestGradientNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradientnode/2948035-scalefactory?language=objc
func (c_ CNNUpsamplingNearestGradientNode) ScaleFactorY() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleFactorY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradientnode/2948024-scalefactorx?language=objc
func (c_ CNNUpsamplingNearestGradientNode) ScaleFactorX() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleFactorX"))
	return rv
}
