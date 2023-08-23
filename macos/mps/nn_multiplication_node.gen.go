// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNMultiplicationNode] class.
var NNMultiplicationNodeClass = _NNMultiplicationNodeClass{objc.GetClass("MPSNNMultiplicationNode")}

type _NNMultiplicationNodeClass struct {
	objc.Class
}

// An interface definition for the [NNMultiplicationNode] class.
type INNMultiplicationNode interface {
	INNBinaryArithmeticNode
}

// A representation of a multiplication operator. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnmultiplicationnode?language=objc
type NNMultiplicationNode struct {
	NNBinaryArithmeticNode
}

func NNMultiplicationNodeFrom(ptr unsafe.Pointer) NNMultiplicationNode {
	return NNMultiplicationNode{
		NNBinaryArithmeticNode: NNBinaryArithmeticNodeFrom(ptr),
	}
}

func (nc _NNMultiplicationNodeClass) Alloc() NNMultiplicationNode {
	rv := objc.Call[NNMultiplicationNode](nc, objc.Sel("alloc"))
	return rv
}

func NNMultiplicationNode_Alloc() NNMultiplicationNode {
	return NNMultiplicationNodeClass.Alloc()
}

func (nc _NNMultiplicationNodeClass) New() NNMultiplicationNode {
	rv := objc.Call[NNMultiplicationNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNMultiplicationNode() NNMultiplicationNode {
	return NNMultiplicationNodeClass.New()
}

func (n_ NNMultiplicationNode) Init() NNMultiplicationNode {
	rv := objc.Call[NNMultiplicationNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNMultiplicationNodeClass) NodeWithSources(sourceNodes []INNImageNode) NNMultiplicationNode {
	rv := objc.Call[NNMultiplicationNode](nc, objc.Sel("nodeWithSources:"), sourceNodes)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890829-nodewithsources?language=objc
func NNMultiplicationNode_NodeWithSources(sourceNodes []INNImageNode) NNMultiplicationNode {
	return NNMultiplicationNodeClass.NodeWithSources(sourceNodes)
}

func (nc _NNMultiplicationNodeClass) NodeWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNMultiplicationNode {
	rv := objc.Call[NNMultiplicationNode](nc, objc.Sel("nodeWithLeftSource:rightSource:"), objc.Ptr(left), objc.Ptr(right))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890830-nodewithleftsource?language=objc
func NNMultiplicationNode_NodeWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNMultiplicationNode {
	return NNMultiplicationNodeClass.NodeWithLeftSourceRightSource(left, right)
}

func (n_ NNMultiplicationNode) InitWithSources(sourceNodes []INNImageNode) NNMultiplicationNode {
	rv := objc.Call[NNMultiplicationNode](n_, objc.Sel("initWithSources:"), sourceNodes)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890820-initwithsources?language=objc
func NewNNMultiplicationNodeWithSources(sourceNodes []INNImageNode) NNMultiplicationNode {
	instance := NNMultiplicationNodeClass.Alloc().InitWithSources(sourceNodes)
	instance.Autorelease()
	return instance
}

func (n_ NNMultiplicationNode) InitWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNMultiplicationNode {
	rv := objc.Call[NNMultiplicationNode](n_, objc.Sel("initWithLeftSource:rightSource:"), objc.Ptr(left), objc.Ptr(right))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890825-initwithleftsource?language=objc
func NewNNMultiplicationNodeWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNMultiplicationNode {
	instance := NNMultiplicationNodeClass.Alloc().InitWithLeftSourceRightSource(left, right)
	instance.Autorelease()
	return instance
}
