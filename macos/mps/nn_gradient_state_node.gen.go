// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNGradientStateNode] class.
var NNGradientStateNodeClass = _NNGradientStateNodeClass{objc.GetClass("MPSNNGradientStateNode")}

type _NNGradientStateNodeClass struct {
	objc.Class
}

// An interface definition for the [NNGradientStateNode] class.
type INNGradientStateNode interface {
	INNStateNode
}

// A representation of the state created to record the properties of a gradient kernel at the time it was encoded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngradientstatenode?language=objc
type NNGradientStateNode struct {
	NNStateNode
}

func NNGradientStateNodeFrom(ptr unsafe.Pointer) NNGradientStateNode {
	return NNGradientStateNode{
		NNStateNode: NNStateNodeFrom(ptr),
	}
}

func (nc _NNGradientStateNodeClass) Alloc() NNGradientStateNode {
	rv := objc.Call[NNGradientStateNode](nc, objc.Sel("alloc"))
	return rv
}

func NNGradientStateNode_Alloc() NNGradientStateNode {
	return NNGradientStateNodeClass.Alloc()
}

func (nc _NNGradientStateNodeClass) New() NNGradientStateNode {
	rv := objc.Call[NNGradientStateNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNGradientStateNode() NNGradientStateNode {
	return NNGradientStateNodeClass.New()
}

func (n_ NNGradientStateNode) Init() NNGradientStateNode {
	rv := objc.Call[NNGradientStateNode](n_, objc.Sel("init"))
	return rv
}
