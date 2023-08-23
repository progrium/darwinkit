// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNUpsamplingNearestNode] class.
var CNNUpsamplingNearestNodeClass = _CNNUpsamplingNearestNodeClass{objc.GetClass("MPSCNNUpsamplingNearestNode")}

type _CNNUpsamplingNearestNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNUpsamplingNearestNode] class.
type ICNNUpsamplingNearestNode interface {
	INNFilterNode
	ScaleFactorY() float64
	ScaleFactorX() float64
}

// A representation of a nearest spatial upsampling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestnode?language=objc
type CNNUpsamplingNearestNode struct {
	NNFilterNode
}

func CNNUpsamplingNearestNodeFrom(ptr unsafe.Pointer) CNNUpsamplingNearestNode {
	return CNNUpsamplingNearestNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (cc _CNNUpsamplingNearestNodeClass) NodeWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode INNImageNode, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingNearestNode {
	rv := objc.Call[CNNUpsamplingNearestNode](cc, objc.Sel("nodeWithSource:integerScaleFactorX:integerScaleFactorY:"), objc.Ptr(sourceNode), integerScaleFactorX, integerScaleFactorY)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestnode/2875985-nodewithsource?language=objc
func CNNUpsamplingNearestNode_NodeWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode INNImageNode, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingNearestNode {
	return CNNUpsamplingNearestNodeClass.NodeWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode, integerScaleFactorX, integerScaleFactorY)
}

func (c_ CNNUpsamplingNearestNode) InitWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode INNImageNode, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingNearestNode {
	rv := objc.Call[CNNUpsamplingNearestNode](c_, objc.Sel("initWithSource:integerScaleFactorX:integerScaleFactorY:"), objc.Ptr(sourceNode), integerScaleFactorX, integerScaleFactorY)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestnode/2875222-initwithsource?language=objc
func NewCNNUpsamplingNearestNodeWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode INNImageNode, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingNearestNode {
	instance := CNNUpsamplingNearestNodeClass.Alloc().InitWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode, integerScaleFactorX, integerScaleFactorY)
	instance.Autorelease()
	return instance
}

func (cc _CNNUpsamplingNearestNodeClass) Alloc() CNNUpsamplingNearestNode {
	rv := objc.Call[CNNUpsamplingNearestNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNUpsamplingNearestNode_Alloc() CNNUpsamplingNearestNode {
	return CNNUpsamplingNearestNodeClass.Alloc()
}

func (cc _CNNUpsamplingNearestNodeClass) New() CNNUpsamplingNearestNode {
	rv := objc.Call[CNNUpsamplingNearestNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNUpsamplingNearestNode() CNNUpsamplingNearestNode {
	return CNNUpsamplingNearestNodeClass.New()
}

func (c_ CNNUpsamplingNearestNode) Init() CNNUpsamplingNearestNode {
	rv := objc.Call[CNNUpsamplingNearestNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestnode/2875155-scalefactory?language=objc
func (c_ CNNUpsamplingNearestNode) ScaleFactorY() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleFactorY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestnode/2875209-scalefactorx?language=objc
func (c_ CNNUpsamplingNearestNode) ScaleFactorX() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleFactorX"))
	return rv
}
