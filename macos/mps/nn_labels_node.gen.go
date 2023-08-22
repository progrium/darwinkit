// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNLabelsNode] class.
var NNLabelsNodeClass = _NNLabelsNodeClass{objc.GetClass("MPSNNLabelsNode")}

type _NNLabelsNodeClass struct {
	objc.Class
}

// An interface definition for the [NNLabelsNode] class.
type INNLabelsNode interface {
	INNStateNode
}

// A placeholder node denoting the per-element weight buffer used by loss and gradient loss kernels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlabelsnode?language=objc
type NNLabelsNode struct {
	NNStateNode
}

func NNLabelsNodeFrom(ptr unsafe.Pointer) NNLabelsNode {
	return NNLabelsNode{
		NNStateNode: NNStateNodeFrom(ptr),
	}
}

func (nc _NNLabelsNodeClass) Alloc() NNLabelsNode {
	rv := objc.Call[NNLabelsNode](nc, objc.Sel("alloc"))
	return rv
}

func NNLabelsNode_Alloc() NNLabelsNode {
	return NNLabelsNodeClass.Alloc()
}

func (nc _NNLabelsNodeClass) New() NNLabelsNode {
	rv := objc.Call[NNLabelsNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNLabelsNode() NNLabelsNode {
	return NNLabelsNodeClass.New()
}

func (n_ NNLabelsNode) Init() NNLabelsNode {
	rv := objc.Call[NNLabelsNode](n_, objc.Sel("init"))
	return rv
}
