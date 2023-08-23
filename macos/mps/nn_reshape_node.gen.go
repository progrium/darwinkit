// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReshapeNode] class.
var NNReshapeNodeClass = _NNReshapeNodeClass{objc.GetClass("MPSNNReshapeNode")}

type _NNReshapeNodeClass struct {
	objc.Class
}

// An interface definition for the [NNReshapeNode] class.
type INNReshapeNode interface {
	INNFilterNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshapenode?language=objc
type NNReshapeNode struct {
	NNFilterNode
}

func NNReshapeNodeFrom(ptr unsafe.Pointer) NNReshapeNode {
	return NNReshapeNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (nc _NNReshapeNodeClass) NodeWithSourceResultWidthResultHeightResultFeatureChannels(source INNImageNode, resultWidth uint, resultHeight uint, resultFeatureChannels uint) NNReshapeNode {
	rv := objc.Call[NNReshapeNode](nc, objc.Sel("nodeWithSource:resultWidth:resultHeight:resultFeatureChannels:"), objc.Ptr(source), resultWidth, resultHeight, resultFeatureChannels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshapenode/3037421-nodewithsource?language=objc
func NNReshapeNode_NodeWithSourceResultWidthResultHeightResultFeatureChannels(source INNImageNode, resultWidth uint, resultHeight uint, resultFeatureChannels uint) NNReshapeNode {
	return NNReshapeNodeClass.NodeWithSourceResultWidthResultHeightResultFeatureChannels(source, resultWidth, resultHeight, resultFeatureChannels)
}

func (n_ NNReshapeNode) InitWithSourceResultWidthResultHeightResultFeatureChannels(source INNImageNode, resultWidth uint, resultHeight uint, resultFeatureChannels uint) NNReshapeNode {
	rv := objc.Call[NNReshapeNode](n_, objc.Sel("initWithSource:resultWidth:resultHeight:resultFeatureChannels:"), objc.Ptr(source), resultWidth, resultHeight, resultFeatureChannels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreshapenode/3037420-initwithsource?language=objc
func NewNNReshapeNodeWithSourceResultWidthResultHeightResultFeatureChannels(source INNImageNode, resultWidth uint, resultHeight uint, resultFeatureChannels uint) NNReshapeNode {
	instance := NNReshapeNodeClass.Alloc().InitWithSourceResultWidthResultHeightResultFeatureChannels(source, resultWidth, resultHeight, resultFeatureChannels)
	instance.Autorelease()
	return instance
}

func (nc _NNReshapeNodeClass) Alloc() NNReshapeNode {
	rv := objc.Call[NNReshapeNode](nc, objc.Sel("alloc"))
	return rv
}

func NNReshapeNode_Alloc() NNReshapeNode {
	return NNReshapeNodeClass.Alloc()
}

func (nc _NNReshapeNodeClass) New() NNReshapeNode {
	rv := objc.Call[NNReshapeNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReshapeNode() NNReshapeNode {
	return NNReshapeNodeClass.New()
}

func (n_ NNReshapeNode) Init() NNReshapeNode {
	rv := objc.Call[NNReshapeNode](n_, objc.Sel("init"))
	return rv
}
