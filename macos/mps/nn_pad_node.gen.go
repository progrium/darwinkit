// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNPadNode] class.
var NNPadNodeClass = _NNPadNodeClass{objc.GetClass("MPSNNPadNode")}

type _NNPadNodeClass struct {
	objc.Class
}

// An interface definition for the [NNPadNode] class.
type INNPadNode interface {
	INNFilterNode
	FillValue() float64
	SetFillValue(value float64)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadnode?language=objc
type NNPadNode struct {
	NNFilterNode
}

func NNPadNodeFrom(ptr unsafe.Pointer) NNPadNode {
	return NNPadNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (nc _NNPadNodeClass) NodeWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode(source INNImageNode, paddingSizeBefore ImageCoordinate, paddingSizeAfter ImageCoordinate, edgeMode ImageEdgeMode) NNPadNode {
	rv := objc.Call[NNPadNode](nc, objc.Sel("nodeWithSource:paddingSizeBefore:paddingSizeAfter:edgeMode:"), objc.Ptr(source), paddingSizeBefore, paddingSizeAfter, edgeMode)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadnode/3037396-nodewithsource?language=objc
func NNPadNode_NodeWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode(source INNImageNode, paddingSizeBefore ImageCoordinate, paddingSizeAfter ImageCoordinate, edgeMode ImageEdgeMode) NNPadNode {
	return NNPadNodeClass.NodeWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode(source, paddingSizeBefore, paddingSizeAfter, edgeMode)
}

func (n_ NNPadNode) InitWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode(source INNImageNode, paddingSizeBefore ImageCoordinate, paddingSizeAfter ImageCoordinate, edgeMode ImageEdgeMode) NNPadNode {
	rv := objc.Call[NNPadNode](n_, objc.Sel("initWithSource:paddingSizeBefore:paddingSizeAfter:edgeMode:"), objc.Ptr(source), paddingSizeBefore, paddingSizeAfter, edgeMode)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadnode/3037395-initwithsource?language=objc
func NewNNPadNodeWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode(source INNImageNode, paddingSizeBefore ImageCoordinate, paddingSizeAfter ImageCoordinate, edgeMode ImageEdgeMode) NNPadNode {
	instance := NNPadNodeClass.Alloc().InitWithSourcePaddingSizeBeforePaddingSizeAfterEdgeMode(source, paddingSizeBefore, paddingSizeAfter, edgeMode)
	instance.Autorelease()
	return instance
}

func (nc _NNPadNodeClass) Alloc() NNPadNode {
	rv := objc.Call[NNPadNode](nc, objc.Sel("alloc"))
	return rv
}

func NNPadNode_Alloc() NNPadNode {
	return NNPadNodeClass.Alloc()
}

func (nc _NNPadNodeClass) New() NNPadNode {
	rv := objc.Call[NNPadNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNPadNode() NNPadNode {
	return NNPadNodeClass.New()
}

func (n_ NNPadNode) Init() NNPadNode {
	rv := objc.Call[NNPadNode](n_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadnode/3037394-fillvalue?language=objc
func (n_ NNPadNode) FillValue() float64 {
	rv := objc.Call[float64](n_, objc.Sel("fillValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnpadnode/3037394-fillvalue?language=objc
func (n_ NNPadNode) SetFillValue(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setFillValue:"), value)
}
