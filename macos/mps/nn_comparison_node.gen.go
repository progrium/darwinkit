// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNComparisonNode] class.
var NNComparisonNodeClass = _NNComparisonNodeClass{objc.GetClass("MPSNNComparisonNode")}

type _NNComparisonNodeClass struct {
	objc.Class
}

// An interface definition for the [NNComparisonNode] class.
type INNComparisonNode interface {
	INNBinaryArithmeticNode
	ComparisonType() NNComparisonType
	SetComparisonType(value NNComparisonType)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncomparisonnode?language=objc
type NNComparisonNode struct {
	NNBinaryArithmeticNode
}

func NNComparisonNodeFrom(ptr unsafe.Pointer) NNComparisonNode {
	return NNComparisonNode{
		NNBinaryArithmeticNode: NNBinaryArithmeticNodeFrom(ptr),
	}
}

func (nc _NNComparisonNodeClass) Alloc() NNComparisonNode {
	rv := objc.Call[NNComparisonNode](nc, objc.Sel("alloc"))
	return rv
}

func NNComparisonNode_Alloc() NNComparisonNode {
	return NNComparisonNodeClass.Alloc()
}

func (nc _NNComparisonNodeClass) New() NNComparisonNode {
	rv := objc.Call[NNComparisonNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNComparisonNode() NNComparisonNode {
	return NNComparisonNodeClass.New()
}

func (n_ NNComparisonNode) Init() NNComparisonNode {
	rv := objc.Call[NNComparisonNode](n_, objc.Sel("init"))
	return rv
}

func (nc _NNComparisonNodeClass) NodeWithSources(sourceNodes []INNImageNode) NNComparisonNode {
	rv := objc.Call[NNComparisonNode](nc, objc.Sel("nodeWithSources:"), sourceNodes)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890829-nodewithsources?language=objc
func NNComparisonNode_NodeWithSources(sourceNodes []INNImageNode) NNComparisonNode {
	return NNComparisonNodeClass.NodeWithSources(sourceNodes)
}

func (nc _NNComparisonNodeClass) NodeWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNComparisonNode {
	rv := objc.Call[NNComparisonNode](nc, objc.Sel("nodeWithLeftSource:rightSource:"), objc.Ptr(left), objc.Ptr(right))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890830-nodewithleftsource?language=objc
func NNComparisonNode_NodeWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNComparisonNode {
	return NNComparisonNodeClass.NodeWithLeftSourceRightSource(left, right)
}

func (n_ NNComparisonNode) InitWithSources(sourceNodes []INNImageNode) NNComparisonNode {
	rv := objc.Call[NNComparisonNode](n_, objc.Sel("initWithSources:"), sourceNodes)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890820-initwithsources?language=objc
func NewNNComparisonNodeWithSources(sourceNodes []INNImageNode) NNComparisonNode {
	instance := NNComparisonNodeClass.Alloc().InitWithSources(sourceNodes)
	instance.Autorelease()
	return instance
}

func (n_ NNComparisonNode) InitWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNComparisonNode {
	rv := objc.Call[NNComparisonNode](n_, objc.Sel("initWithLeftSource:rightSource:"), objc.Ptr(left), objc.Ptr(right))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinaryarithmeticnode/2890825-initwithleftsource?language=objc
func NewNNComparisonNodeWithLeftSourceRightSource(left INNImageNode, right INNImageNode) NNComparisonNode {
	instance := NNComparisonNodeClass.Alloc().InitWithLeftSourceRightSource(left, right)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncomparisonnode/3037389-comparisontype?language=objc
func (n_ NNComparisonNode) ComparisonType() NNComparisonType {
	rv := objc.Call[NNComparisonType](n_, objc.Sel("comparisonType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncomparisonnode/3037389-comparisontype?language=objc
func (n_ NNComparisonNode) SetComparisonType(value NNComparisonType) {
	objc.Call[objc.Void](n_, objc.Sel("setComparisonType:"), value)
}
