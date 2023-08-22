// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNGradientFilterNode] class.
var NNGradientFilterNodeClass = _NNGradientFilterNodeClass{objc.GetClass("MPSNNGradientFilterNode")}

type _NNGradientFilterNodeClass struct {
	objc.Class
}

// An interface definition for the [NNGradientFilterNode] class.
type INNGradientFilterNode interface {
	INNFilterNode
}

// A representation of a gradient filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngradientfilternode?language=objc
type NNGradientFilterNode struct {
	NNFilterNode
}

func NNGradientFilterNodeFrom(ptr unsafe.Pointer) NNGradientFilterNode {
	return NNGradientFilterNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (nc _NNGradientFilterNodeClass) Alloc() NNGradientFilterNode {
	rv := objc.Call[NNGradientFilterNode](nc, objc.Sel("alloc"))
	return rv
}

func NNGradientFilterNode_Alloc() NNGradientFilterNode {
	return NNGradientFilterNodeClass.Alloc()
}

func (nc _NNGradientFilterNodeClass) New() NNGradientFilterNode {
	rv := objc.Call[NNGradientFilterNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNGradientFilterNode() NNGradientFilterNode {
	return NNGradientFilterNodeClass.New()
}

func (n_ NNGradientFilterNode) Init() NNGradientFilterNode {
	rv := objc.Call[NNGradientFilterNode](n_, objc.Sel("init"))
	return rv
}
