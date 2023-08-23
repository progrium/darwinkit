// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNormalizationNode] class.
var CNNNormalizationNodeClass = _CNNNormalizationNodeClass{objc.GetClass("MPSCNNNormalizationNode")}

type _CNNNormalizationNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNormalizationNode] class.
type ICNNNormalizationNode interface {
	INNFilterNode
	Beta() float64
	SetBeta(value float64)
	Delta() float64
	SetDelta(value float64)
	Alpha() float64
	SetAlpha(value float64)
}

// Virtual base class for CNN normalization nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode?language=objc
type CNNNormalizationNode struct {
	NNFilterNode
}

func CNNNormalizationNodeFrom(ptr unsafe.Pointer) CNNNormalizationNode {
	return CNNNormalizationNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (cc _CNNNormalizationNodeClass) NodeWithSource(sourceNode INNImageNode) CNNNormalizationNode {
	rv := objc.Call[CNNNormalizationNode](cc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866460-nodewithsource?language=objc
func CNNNormalizationNode_NodeWithSource(sourceNode INNImageNode) CNNNormalizationNode {
	return CNNNormalizationNodeClass.NodeWithSource(sourceNode)
}

func (c_ CNNNormalizationNode) InitWithSource(sourceNode INNImageNode) CNNNormalizationNode {
	rv := objc.Call[CNNNormalizationNode](c_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866425-initwithsource?language=objc
func NewCNNNormalizationNodeWithSource(sourceNode INNImageNode) CNNNormalizationNode {
	instance := CNNNormalizationNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (cc _CNNNormalizationNodeClass) Alloc() CNNNormalizationNode {
	rv := objc.Call[CNNNormalizationNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNNormalizationNode_Alloc() CNNNormalizationNode {
	return CNNNormalizationNodeClass.Alloc()
}

func (cc _CNNNormalizationNodeClass) New() CNNNormalizationNode {
	rv := objc.Call[CNNNormalizationNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNormalizationNode() CNNNormalizationNode {
	return CNNNormalizationNodeClass.New()
}

func (c_ CNNNormalizationNode) Init() CNNNormalizationNode {
	rv := objc.Call[CNNNormalizationNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866497-beta?language=objc
func (c_ CNNNormalizationNode) Beta() float64 {
	rv := objc.Call[float64](c_, objc.Sel("beta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866497-beta?language=objc
func (c_ CNNNormalizationNode) SetBeta(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setBeta:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866482-delta?language=objc
func (c_ CNNNormalizationNode) Delta() float64 {
	rv := objc.Call[float64](c_, objc.Sel("delta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866482-delta?language=objc
func (c_ CNNNormalizationNode) SetDelta(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setDelta:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866474-alpha?language=objc
func (c_ CNNNormalizationNode) Alpha() float64 {
	rv := objc.Call[float64](c_, objc.Sel("alpha"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866474-alpha?language=objc
func (c_ CNNNormalizationNode) SetAlpha(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setAlpha:"), value)
}
