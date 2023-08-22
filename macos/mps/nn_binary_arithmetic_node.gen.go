// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNBinaryArithmeticNode] class.
var NNBinaryArithmeticNodeClass = _NNBinaryArithmeticNodeClass{objc.GetClass("MPSNNBinaryArithmeticNode")}

type _NNBinaryArithmeticNodeClass struct {
	objc.Class
}

// An interface definition for the [NNBinaryArithmeticNode] class.
type INNBinaryArithmeticNode interface {
	INNFilterNode
	GradientClass() objc.Class
	SecondaryScale() float64
	SetSecondaryScale(value float64)
	SecondaryStrideInPixelsY() uint
	SetSecondaryStrideInPixelsY(value uint)
	SecondaryStrideInPixelsX() uint
	SetSecondaryStrideInPixelsX(value uint)
	PrimaryStrideInPixelsY() uint
	SetPrimaryStrideInPixelsY(value uint)
	MaximumValue() float64
	SetMaximumValue(value float64)
	PrimaryStrideInPixelsX() uint
	SetPrimaryStrideInPixelsX(value uint)
	PrimaryScale() float64
	SetPrimaryScale(value float64)
	MinimumValue() float64
	SetMinimumValue(value float64)
	Bias() float64
	SetBias(value float64)
	SecondaryStrideInFeatureChannels() uint
	SetSecondaryStrideInFeatureChannels(value uint)
	PrimaryStrideInFeatureChannels() uint
	SetPrimaryStrideInFeatureChannels(value uint)
}

// Virtual base class for basic arithmetic nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode?language=objc
type NNBinaryArithmeticNode struct {
	NNFilterNode
}

func NNBinaryArithmeticNodeFrom(ptr unsafe.Pointer) NNBinaryArithmeticNode {
	return NNBinaryArithmeticNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (nc _NNBinaryArithmeticNodeClass) NodeWithSources(sourceNodes []INNImageNode) NNBinaryArithmeticNode {
	rv := objc.Call[NNBinaryArithmeticNode](nc, objc.Sel("nodeWithSources:"), sourceNodes)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890829-nodewithsources?language=objc
func NNBinaryArithmeticNode_NodeWithSources(sourceNodes []INNImageNode) NNBinaryArithmeticNode {
	return NNBinaryArithmeticNodeClass.NodeWithSources(sourceNodes)
}

func (nc _NNBinaryArithmeticNodeClass) NodeWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNBinaryArithmeticNode {
	rv := objc.Call[NNBinaryArithmeticNode](nc, objc.Sel("nodeWithLeftSource:rightSource:"), objc.Ptr(left), objc.Ptr(right))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890830-nodewithleftsource?language=objc
func NNBinaryArithmeticNode_NodeWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNBinaryArithmeticNode {
	return NNBinaryArithmeticNodeClass.NodeWithLeftSourceRightSource(left, right)
}

func (n_ NNBinaryArithmeticNode) InitWithSources(sourceNodes []INNImageNode) NNBinaryArithmeticNode {
	rv := objc.Call[NNBinaryArithmeticNode](n_, objc.Sel("initWithSources:"), sourceNodes)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890820-initwithsources?language=objc
func NewNNBinaryArithmeticNodeWithSources(sourceNodes []INNImageNode) NNBinaryArithmeticNode {
	instance := NNBinaryArithmeticNodeClass.Alloc().InitWithSources(sourceNodes)
	instance.Autorelease()
	return instance
}

func (n_ NNBinaryArithmeticNode) InitWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNBinaryArithmeticNode {
	rv := objc.Call[NNBinaryArithmeticNode](n_, objc.Sel("initWithLeftSource:rightSource:"), objc.Ptr(left), objc.Ptr(right))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890825-initwithleftsource?language=objc
func NewNNBinaryArithmeticNodeWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNBinaryArithmeticNode {
	instance := NNBinaryArithmeticNodeClass.Alloc().InitWithLeftSourceRightSource(left, right)
	instance.Autorelease()
	return instance
}

func (nc _NNBinaryArithmeticNodeClass) Alloc() NNBinaryArithmeticNode {
	rv := objc.Call[NNBinaryArithmeticNode](nc, objc.Sel("alloc"))
	return rv
}

func NNBinaryArithmeticNode_Alloc() NNBinaryArithmeticNode {
	return NNBinaryArithmeticNodeClass.Alloc()
}

func (nc _NNBinaryArithmeticNodeClass) New() NNBinaryArithmeticNode {
	rv := objc.Call[NNBinaryArithmeticNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNBinaryArithmeticNode() NNBinaryArithmeticNode {
	return NNBinaryArithmeticNodeClass.New()
}

func (n_ NNBinaryArithmeticNode) Init() NNBinaryArithmeticNode {
	rv := objc.Call[NNBinaryArithmeticNode](n_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952978-gradientclass?language=objc
func (n_ NNBinaryArithmeticNode) GradientClass() objc.Class {
	rv := objc.Call[objc.Class](n_, objc.Sel("gradientClass"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952976-secondaryscale?language=objc
func (n_ NNBinaryArithmeticNode) SecondaryScale() float64 {
	rv := objc.Call[float64](n_, objc.Sel("secondaryScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952976-secondaryscale?language=objc
func (n_ NNBinaryArithmeticNode) SetSecondaryScale(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setSecondaryScale:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952985-secondarystrideinpixelsy?language=objc
func (n_ NNBinaryArithmeticNode) SecondaryStrideInPixelsY() uint {
	rv := objc.Call[uint](n_, objc.Sel("secondaryStrideInPixelsY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952985-secondarystrideinpixelsy?language=objc
func (n_ NNBinaryArithmeticNode) SetSecondaryStrideInPixelsY(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setSecondaryStrideInPixelsY:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952972-secondarystrideinpixelsx?language=objc
func (n_ NNBinaryArithmeticNode) SecondaryStrideInPixelsX() uint {
	rv := objc.Call[uint](n_, objc.Sel("secondaryStrideInPixelsX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952972-secondarystrideinpixelsx?language=objc
func (n_ NNBinaryArithmeticNode) SetSecondaryStrideInPixelsX(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setSecondaryStrideInPixelsX:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952996-primarystrideinpixelsy?language=objc
func (n_ NNBinaryArithmeticNode) PrimaryStrideInPixelsY() uint {
	rv := objc.Call[uint](n_, objc.Sel("primaryStrideInPixelsY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952996-primarystrideinpixelsy?language=objc
func (n_ NNBinaryArithmeticNode) SetPrimaryStrideInPixelsY(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setPrimaryStrideInPixelsY:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952979-maximumvalue?language=objc
func (n_ NNBinaryArithmeticNode) MaximumValue() float64 {
	rv := objc.Call[float64](n_, objc.Sel("maximumValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952979-maximumvalue?language=objc
func (n_ NNBinaryArithmeticNode) SetMaximumValue(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setMaximumValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952973-primarystrideinpixelsx?language=objc
func (n_ NNBinaryArithmeticNode) PrimaryStrideInPixelsX() uint {
	rv := objc.Call[uint](n_, objc.Sel("primaryStrideInPixelsX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952973-primarystrideinpixelsx?language=objc
func (n_ NNBinaryArithmeticNode) SetPrimaryStrideInPixelsX(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setPrimaryStrideInPixelsX:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952966-primaryscale?language=objc
func (n_ NNBinaryArithmeticNode) PrimaryScale() float64 {
	rv := objc.Call[float64](n_, objc.Sel("primaryScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952966-primaryscale?language=objc
func (n_ NNBinaryArithmeticNode) SetPrimaryScale(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setPrimaryScale:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952970-minimumvalue?language=objc
func (n_ NNBinaryArithmeticNode) MinimumValue() float64 {
	rv := objc.Call[float64](n_, objc.Sel("minimumValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952970-minimumvalue?language=objc
func (n_ NNBinaryArithmeticNode) SetMinimumValue(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setMinimumValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952964-bias?language=objc
func (n_ NNBinaryArithmeticNode) Bias() float64 {
	rv := objc.Call[float64](n_, objc.Sel("bias"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952964-bias?language=objc
func (n_ NNBinaryArithmeticNode) SetBias(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setBias:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952974-secondarystrideinfeaturechannels?language=objc
func (n_ NNBinaryArithmeticNode) SecondaryStrideInFeatureChannels() uint {
	rv := objc.Call[uint](n_, objc.Sel("secondaryStrideInFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952974-secondarystrideinfeaturechannels?language=objc
func (n_ NNBinaryArithmeticNode) SetSecondaryStrideInFeatureChannels(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setSecondaryStrideInFeatureChannels:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952983-primarystrideinfeaturechannels?language=objc
func (n_ NNBinaryArithmeticNode) PrimaryStrideInFeatureChannels() uint {
	rv := objc.Call[uint](n_, objc.Sel("primaryStrideInFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2952983-primarystrideinfeaturechannels?language=objc
func (n_ NNBinaryArithmeticNode) SetPrimaryStrideInFeatureChannels(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setPrimaryStrideInFeatureChannels:"), value)
}
