// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNScaleNode] class.
var NNScaleNodeClass = _NNScaleNodeClass{objc.GetClass("MPSNNScaleNode")}

type _NNScaleNodeClass struct {
	objc.Class
}

// An interface definition for the [NNScaleNode] class.
type INNScaleNode interface {
	INNFilterNode
}

// Abstract node representing an image resampling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode?language=objc
type NNScaleNode struct {
	NNFilterNode
}

func NNScaleNodeFrom(ptr unsafe.Pointer) NNScaleNode {
	return NNScaleNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (nc _NNScaleNodeClass) NodeWithSourceOutputSize(sourceNode INNImageNode, size metal.Size) NNScaleNode {
	rv := objc.Call[NNScaleNode](nc, objc.Sel("nodeWithSource:outputSize:"), objc.Ptr(sourceNode), size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode/2915280-nodewithsource?language=objc
func NNScaleNode_NodeWithSourceOutputSize(sourceNode INNImageNode, size metal.Size) NNScaleNode {
	return NNScaleNodeClass.NodeWithSourceOutputSize(sourceNode, size)
}

func (n_ NNScaleNode) InitWithSourceOutputSize(sourceNode INNImageNode, size metal.Size) NNScaleNode {
	rv := objc.Call[NNScaleNode](n_, objc.Sel("initWithSource:outputSize:"), objc.Ptr(sourceNode), size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnscalenode/2915285-initwithsource?language=objc
func NewNNScaleNodeWithSourceOutputSize(sourceNode INNImageNode, size metal.Size) NNScaleNode {
	instance := NNScaleNodeClass.Alloc().InitWithSourceOutputSize(sourceNode, size)
	instance.Autorelease()
	return instance
}

func (nc _NNScaleNodeClass) Alloc() NNScaleNode {
	rv := objc.Call[NNScaleNode](nc, objc.Sel("alloc"))
	return rv
}

func NNScaleNode_Alloc() NNScaleNode {
	return NNScaleNodeClass.Alloc()
}

func (nc _NNScaleNodeClass) New() NNScaleNode {
	rv := objc.Call[NNScaleNode](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNScaleNode() NNScaleNode {
	return NNScaleNodeClass.New()
}

func (n_ NNScaleNode) Init() NNScaleNode {
	rv := objc.Call[NNScaleNode](n_, objc.Sel("init"))
	return rv
}
