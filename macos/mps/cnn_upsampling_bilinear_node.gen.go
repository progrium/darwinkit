// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNUpsamplingBilinearNode] class.
var CNNUpsamplingBilinearNodeClass = _CNNUpsamplingBilinearNodeClass{objc.GetClass("MPSCNNUpsamplingBilinearNode")}

type _CNNUpsamplingBilinearNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNUpsamplingBilinearNode] class.
type ICNNUpsamplingBilinearNode interface {
	INNFilterNode
	AlignCorners() bool
	ScaleFactorY() float64
	ScaleFactorX() float64
}

// A representation of a bilinear spatial upsampling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode?language=objc
type CNNUpsamplingBilinearNode struct {
	NNFilterNode
}

func CNNUpsamplingBilinearNodeFrom(ptr unsafe.Pointer) CNNUpsamplingBilinearNode {
	return CNNUpsamplingBilinearNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (cc _CNNUpsamplingBilinearNodeClass) NodeWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode INNImageNode, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingBilinearNode {
	rv := objc.Call[CNNUpsamplingBilinearNode](cc, objc.Sel("nodeWithSource:integerScaleFactorX:integerScaleFactorY:"), objc.Ptr(sourceNode), integerScaleFactorX, integerScaleFactorY)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2875987-nodewithsource?language=objc
func CNNUpsamplingBilinearNode_NodeWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode INNImageNode, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingBilinearNode {
	return CNNUpsamplingBilinearNodeClass.NodeWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode, integerScaleFactorX, integerScaleFactorY)
}

func (c_ CNNUpsamplingBilinearNode) InitWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode INNImageNode, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingBilinearNode {
	rv := objc.Call[CNNUpsamplingBilinearNode](c_, objc.Sel("initWithSource:integerScaleFactorX:integerScaleFactorY:"), objc.Ptr(sourceNode), integerScaleFactorX, integerScaleFactorY)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2875152-initwithsource?language=objc
func NewCNNUpsamplingBilinearNodeWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode INNImageNode, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingBilinearNode {
	instance := CNNUpsamplingBilinearNodeClass.Alloc().InitWithSourceIntegerScaleFactorXIntegerScaleFactorY(sourceNode, integerScaleFactorX, integerScaleFactorY)
	instance.Autorelease()
	return instance
}

func (cc _CNNUpsamplingBilinearNodeClass) Alloc() CNNUpsamplingBilinearNode {
	rv := objc.Call[CNNUpsamplingBilinearNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNUpsamplingBilinearNode_Alloc() CNNUpsamplingBilinearNode {
	return CNNUpsamplingBilinearNodeClass.Alloc()
}

func (cc _CNNUpsamplingBilinearNodeClass) New() CNNUpsamplingBilinearNode {
	rv := objc.Call[CNNUpsamplingBilinearNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNUpsamplingBilinearNode() CNNUpsamplingBilinearNode {
	return CNNUpsamplingBilinearNodeClass.New()
}

func (c_ CNNUpsamplingBilinearNode) Init() CNNUpsamplingBilinearNode {
	rv := objc.Call[CNNUpsamplingBilinearNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2966687-aligncorners?language=objc
func (c_ CNNUpsamplingBilinearNode) AlignCorners() bool {
	rv := objc.Call[bool](c_, objc.Sel("alignCorners"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2875150-scalefactory?language=objc
func (c_ CNNUpsamplingBilinearNode) ScaleFactorY() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleFactorY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilinearnode/2875153-scalefactorx?language=objc
func (c_ CNNUpsamplingBilinearNode) ScaleFactorX() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleFactorX"))
	return rv
}
