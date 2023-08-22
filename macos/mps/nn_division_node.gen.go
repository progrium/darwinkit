// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNDivisionNode] class.
var NNDivisionNodeClass = _NNDivisionNodeClass{objc.GetClass("MPSNNDivisionNode")}

type _NNDivisionNodeClass struct {
	objc.Class
}

// An interface definition for the [NNDivisionNode] class.
type INNDivisionNode interface {
	INNBinaryArithmeticNode
}

// A representation of a division operator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnndivisionnode?language=objc
type NNDivisionNode struct {
	NNBinaryArithmeticNode
}

func NNDivisionNodeFrom(ptr unsafe.Pointer) NNDivisionNode {
	return NNDivisionNode{
		NNBinaryArithmeticNode: NNBinaryArithmeticNodeFrom(ptr),
	}
}

func (nc _NNDivisionNodeClass) Alloc() NNDivisionNode {
	rv := objc.Call[NNDivisionNode](nc, objc.Sel("alloc"))
	return rv
}

func NNDivisionNode_Alloc() NNDivisionNode {
	return NNDivisionNodeClass.Alloc()
}

func (nc _NNDivisionNodeClass) New() NNDivisionNode {
	rv := objc.Call[NNDivisionNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNDivisionNode() NNDivisionNode {
	return NNDivisionNodeClass.New()
}

func (n_ NNDivisionNode) Init() NNDivisionNode {
	rv := objc.Call[NNDivisionNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNDivisionNodeClass) NodeWithSources(sourceNodes []INNImageNode) NNDivisionNode {
	rv := objc.Call[NNDivisionNode](nc, objc.Sel("nodeWithSources:"), sourceNodes)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890829-nodewithsources?language=objc
func NNDivisionNode_NodeWithSources(sourceNodes []INNImageNode) NNDivisionNode {
	return NNDivisionNodeClass.NodeWithSources(sourceNodes)
}

func (nc _NNDivisionNodeClass) NodeWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNDivisionNode {
	rv := objc.Call[NNDivisionNode](nc, objc.Sel("nodeWithLeftSource:rightSource:"), objc.Ptr(left), objc.Ptr(right))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890830-nodewithleftsource?language=objc
func NNDivisionNode_NodeWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNDivisionNode {
	return NNDivisionNodeClass.NodeWithLeftSourceRightSource(left, right)
}

func (n_ NNDivisionNode) InitWithSources(sourceNodes []INNImageNode) NNDivisionNode {
	rv := objc.Call[NNDivisionNode](n_, objc.Sel("initWithSources:"), sourceNodes)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890820-initwithsources?language=objc
func NewNNDivisionNodeWithSources(sourceNodes []INNImageNode) NNDivisionNode {
	instance := NNDivisionNodeClass.Alloc().InitWithSources(sourceNodes)
	instance.Autorelease()
	return instance
}

func (n_ NNDivisionNode) InitWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNDivisionNode {
	rv := objc.Call[NNDivisionNode](n_, objc.Sel("initWithLeftSource:rightSource:"), objc.Ptr(left), objc.Ptr(right))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890825-initwithleftsource?language=objc
func NewNNDivisionNodeWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNDivisionNode {
	instance := NNDivisionNodeClass.Alloc().InitWithLeftSourceRightSource(left, right)
	instance.Autorelease()
	return instance
}
