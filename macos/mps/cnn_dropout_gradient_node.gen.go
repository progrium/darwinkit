// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNDropoutGradientNode] class.
var CNNDropoutGradientNodeClass = _CNNDropoutGradientNodeClass{objc.GetClass("MPSCNNDropoutGradientNode")}

type _CNNDropoutGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNDropoutGradientNode] class.
type ICNNDropoutGradientNode interface {
	INNGradientFilterNode
	KeepProbability() float64
	MaskStrideInPixels() metal.Size
	Seed() uint
}

// A representation of a gradient dropout filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientnode?language=objc
type CNNDropoutGradientNode struct {
	NNGradientFilterNode
}

func CNNDropoutGradientNodeFrom(ptr unsafe.Pointer) CNNDropoutGradientNode {
	return CNNDropoutGradientNode{
		NNGradientFilterNode: NNGradientFilterNodeFrom(ptr),
	}
}

func (c_ CNNDropoutGradientNode) InitWithSourceGradientSourceImageGradientStateKeepProbabilitySeedMaskStrideInPixels(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, keepProbability float64, seed uint, maskStrideInPixels metal.Size) CNNDropoutGradientNode {
	rv := objc.Call[CNNDropoutGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:gradientState:keepProbability:seed:maskStrideInPixels:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), keepProbability, seed, maskStrideInPixels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientnode/2948001-initwithsourcegradient?language=objc
func NewCNNDropoutGradientNodeWithSourceGradientSourceImageGradientStateKeepProbabilitySeedMaskStrideInPixels(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, keepProbability float64, seed uint, maskStrideInPixels metal.Size) CNNDropoutGradientNode {
	instance := CNNDropoutGradientNodeClass.Alloc().InitWithSourceGradientSourceImageGradientStateKeepProbabilitySeedMaskStrideInPixels(sourceGradient, sourceImage, gradientState, keepProbability, seed, maskStrideInPixels)
	instance.Autorelease()
	return instance
}

func (cc _CNNDropoutGradientNodeClass) NodeWithSourceGradientSourceImageGradientStateKeepProbabilitySeedMaskStrideInPixels(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, keepProbability float64, seed uint, maskStrideInPixels metal.Size) CNNDropoutGradientNode {
	rv := objc.Call[CNNDropoutGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:gradientState:keepProbability:seed:maskStrideInPixels:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), keepProbability, seed, maskStrideInPixels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientnode/2947997-nodewithsourcegradient?language=objc
func CNNDropoutGradientNode_NodeWithSourceGradientSourceImageGradientStateKeepProbabilitySeedMaskStrideInPixels(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState INNGradientStateNode, keepProbability float64, seed uint, maskStrideInPixels metal.Size) CNNDropoutGradientNode {
	return CNNDropoutGradientNodeClass.NodeWithSourceGradientSourceImageGradientStateKeepProbabilitySeedMaskStrideInPixels(sourceGradient, sourceImage, gradientState, keepProbability, seed, maskStrideInPixels)
}

func (cc _CNNDropoutGradientNodeClass) Alloc() CNNDropoutGradientNode {
	rv := objc.Call[CNNDropoutGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNDropoutGradientNode_Alloc() CNNDropoutGradientNode {
	return CNNDropoutGradientNodeClass.Alloc()
}

func (cc _CNNDropoutGradientNodeClass) New() CNNDropoutGradientNode {
	rv := objc.Call[CNNDropoutGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNDropoutGradientNode() CNNDropoutGradientNode {
	return CNNDropoutGradientNodeClass.New()
}

func (c_ CNNDropoutGradientNode) Init() CNNDropoutGradientNode {
	rv := objc.Call[CNNDropoutGradientNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientnode/2947988-keepprobability?language=objc
func (c_ CNNDropoutGradientNode) KeepProbability() float64 {
	rv := objc.Call[float64](c_, objc.Sel("keepProbability"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientnode/2947972-maskstrideinpixels?language=objc
func (c_ CNNDropoutGradientNode) MaskStrideInPixels() metal.Size {
	rv := objc.Call[metal.Size](c_, objc.Sel("maskStrideInPixels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientnode/2948003-seed?language=objc
func (c_ CNNDropoutGradientNode) Seed() uint {
	rv := objc.Call[uint](c_, objc.Sel("seed"))
	return rv
}
