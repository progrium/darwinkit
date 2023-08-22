// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNMultiaryGradientStateNode] class.
var NNMultiaryGradientStateNodeClass = _NNMultiaryGradientStateNodeClass{objc.GetClass("MPSNNMultiaryGradientStateNode")}

type _NNMultiaryGradientStateNodeClass struct {
	objc.Class
}

// An interface definition for the [NNMultiaryGradientStateNode] class.
type INNMultiaryGradientStateNode interface {
	INNStateNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnmultiarygradientstatenode?language=objc
type NNMultiaryGradientStateNode struct {
	NNStateNode
}

func NNMultiaryGradientStateNodeFrom(ptr unsafe.Pointer) NNMultiaryGradientStateNode {
	return NNMultiaryGradientStateNode{
		NNStateNode: NNStateNodeFrom(ptr),
	}
}

func (nc _NNMultiaryGradientStateNodeClass) Alloc() NNMultiaryGradientStateNode {
	rv := objc.Call[NNMultiaryGradientStateNode](nc, objc.Sel("alloc"))
	return rv
}

func NNMultiaryGradientStateNode_Alloc() NNMultiaryGradientStateNode {
	return NNMultiaryGradientStateNodeClass.Alloc()
}

func (nc _NNMultiaryGradientStateNodeClass) New() NNMultiaryGradientStateNode {
	rv := objc.Call[NNMultiaryGradientStateNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNMultiaryGradientStateNode() NNMultiaryGradientStateNode {
	return NNMultiaryGradientStateNodeClass.New()
}

func (n_ NNMultiaryGradientStateNode) Init() NNMultiaryGradientStateNode {
	rv := objc.Call[NNMultiaryGradientStateNode](n_, objc.Sel("init"))
	return rv
}
