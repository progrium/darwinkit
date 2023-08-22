// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNBinaryGradientStateNode] class.
var NNBinaryGradientStateNodeClass = _NNBinaryGradientStateNodeClass{objc.GetClass("MPSNNBinaryGradientStateNode")}

type _NNBinaryGradientStateNodeClass struct {
	objc.Class
}

// An interface definition for the [NNBinaryGradientStateNode] class.
type INNBinaryGradientStateNode interface {
	INNStateNode
}

// A representation of the state created to record the properties of a binary gradient kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinarygradientstatenode?language=objc
type NNBinaryGradientStateNode struct {
	NNStateNode
}

func NNBinaryGradientStateNodeFrom(ptr unsafe.Pointer) NNBinaryGradientStateNode {
	return NNBinaryGradientStateNode{
		NNStateNode: NNStateNodeFrom(ptr),
	}
}

func (nc _NNBinaryGradientStateNodeClass) Alloc() NNBinaryGradientStateNode {
	rv := objc.Call[NNBinaryGradientStateNode](nc, objc.Sel("alloc"))
	return rv
}

func NNBinaryGradientStateNode_Alloc() NNBinaryGradientStateNode {
	return NNBinaryGradientStateNodeClass.Alloc()
}

func (nc _NNBinaryGradientStateNodeClass) New() NNBinaryGradientStateNode {
	rv := objc.Call[NNBinaryGradientStateNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNBinaryGradientStateNode() NNBinaryGradientStateNode {
	return NNBinaryGradientStateNodeClass.New()
}

func (n_ NNBinaryGradientStateNode) Init() NNBinaryGradientStateNode {
	rv := objc.Call[NNBinaryGradientStateNode](n_, objc.Sel("init"))
	return rv
}
