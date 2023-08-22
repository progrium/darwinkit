// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNInitialGradientNode] class.
var NNInitialGradientNodeClass = _NNInitialGradientNodeClass{objc.GetClass("MPSNNInitialGradientNode")}

type _NNInitialGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [NNInitialGradientNode] class.
type INNInitialGradientNode interface {
	INNFilterNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnninitialgradientnode?language=objc
type NNInitialGradientNode struct {
	NNFilterNode
}

func NNInitialGradientNodeFrom(ptr unsafe.Pointer) NNInitialGradientNode {
	return NNInitialGradientNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (nc _NNInitialGradientNodeClass) NodeWithSource(source INNImageNode) NNInitialGradientNode {
	rv := objc.Call[NNInitialGradientNode](nc, objc.Sel("nodeWithSource:"), objc.Ptr(source))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnninitialgradientnode/3131849-nodewithsource?language=objc
func NNInitialGradientNode_NodeWithSource(source INNImageNode) NNInitialGradientNode {
	return NNInitialGradientNodeClass.NodeWithSource(source)
}

func (n_ NNInitialGradientNode) InitWithSource(source INNImageNode) NNInitialGradientNode {
	rv := objc.Call[NNInitialGradientNode](n_, objc.Sel("initWithSource:"), objc.Ptr(source))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnninitialgradientnode/3131848-initwithsource?language=objc
func NewNNInitialGradientNodeWithSource(source INNImageNode) NNInitialGradientNode {
	instance := NNInitialGradientNodeClass.Alloc().InitWithSource(source)
	instance.Autorelease()
	return instance
}

func (nc _NNInitialGradientNodeClass) Alloc() NNInitialGradientNode {
	rv := objc.Call[NNInitialGradientNode](nc, objc.Sel("alloc"))
	return rv
}

func NNInitialGradientNode_Alloc() NNInitialGradientNode {
	return NNInitialGradientNodeClass.Alloc()
}

func (nc _NNInitialGradientNodeClass) New() NNInitialGradientNode {
	rv := objc.Call[NNInitialGradientNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNInitialGradientNode() NNInitialGradientNode {
	return NNInitialGradientNodeClass.New()
}

func (n_ NNInitialGradientNode) Init() NNInitialGradientNode {
	rv := objc.Call[NNInitialGradientNode](n_, objc.Sel("init"))
	return rv
}
