// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNArithmeticGradientStateNode] class.
var NNArithmeticGradientStateNodeClass = _NNArithmeticGradientStateNodeClass{objc.GetClass("MPSNNArithmeticGradientStateNode")}

type _NNArithmeticGradientStateNodeClass struct {
	objc.Class
}

// An interface definition for the [NNArithmeticGradientStateNode] class.
type INNArithmeticGradientStateNode interface {
	INNBinaryGradientStateNode
}

// A representation of the clamp mask used by gradient arithmetic operators. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnarithmeticgradientstatenode?language=objc
type NNArithmeticGradientStateNode struct {
	NNBinaryGradientStateNode
}

func NNArithmeticGradientStateNodeFrom(ptr unsafe.Pointer) NNArithmeticGradientStateNode {
	return NNArithmeticGradientStateNode{
		NNBinaryGradientStateNode: NNBinaryGradientStateNodeFrom(ptr),
	}
}

func (nc _NNArithmeticGradientStateNodeClass) Alloc() NNArithmeticGradientStateNode {
	rv := objc.Call[NNArithmeticGradientStateNode](nc, objc.Sel("alloc"))
	return rv
}

func NNArithmeticGradientStateNode_Alloc() NNArithmeticGradientStateNode {
	return NNArithmeticGradientStateNodeClass.Alloc()
}

func (nc _NNArithmeticGradientStateNodeClass) New() NNArithmeticGradientStateNode {
	rv := objc.Call[NNArithmeticGradientStateNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNArithmeticGradientStateNode() NNArithmeticGradientStateNode {
	return NNArithmeticGradientStateNodeClass.New()
}

func (n_ NNArithmeticGradientStateNode) Init() NNArithmeticGradientStateNode {
	rv := objc.Call[NNArithmeticGradientStateNode](n_, objc.Sel("init"))
	return rv
}
