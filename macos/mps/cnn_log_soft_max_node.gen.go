// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNLogSoftMaxNode] class.
var CNNLogSoftMaxNodeClass = _CNNLogSoftMaxNodeClass{objc.GetClass("MPSCNNLogSoftMaxNode")}

type _CNNLogSoftMaxNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNLogSoftMaxNode] class.
type ICNNLogSoftMaxNode interface {
	INNFilterNode
}

// A representation of a logarithmic softmax filter kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxnode?language=objc
type CNNLogSoftMaxNode struct {
	NNFilterNode
}

func CNNLogSoftMaxNodeFrom(ptr unsafe.Pointer) CNNLogSoftMaxNode {
	return CNNLogSoftMaxNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (cc _CNNLogSoftMaxNodeClass) NodeWithSource(sourceNode INNImageNode) CNNLogSoftMaxNode {
	rv := objc.Call[CNNLogSoftMaxNode](cc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxnode/2866434-nodewithsource?language=objc
func CNNLogSoftMaxNode_NodeWithSource(sourceNode INNImageNode) CNNLogSoftMaxNode {
	return CNNLogSoftMaxNodeClass.NodeWithSource(sourceNode)
}

func (c_ CNNLogSoftMaxNode) InitWithSource(sourceNode INNImageNode) CNNLogSoftMaxNode {
	rv := objc.Call[CNNLogSoftMaxNode](c_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlogsoftmaxnode/2866457-initwithsource?language=objc
func NewCNNLogSoftMaxNodeWithSource(sourceNode INNImageNode) CNNLogSoftMaxNode {
	instance := CNNLogSoftMaxNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (cc _CNNLogSoftMaxNodeClass) Alloc() CNNLogSoftMaxNode {
	rv := objc.Call[CNNLogSoftMaxNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNLogSoftMaxNode_Alloc() CNNLogSoftMaxNode {
	return CNNLogSoftMaxNodeClass.Alloc()
}

func (cc _CNNLogSoftMaxNodeClass) New() CNNLogSoftMaxNode {
	rv := objc.Call[CNNLogSoftMaxNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNLogSoftMaxNode() CNNLogSoftMaxNode {
	return CNNLogSoftMaxNodeClass.New()
}

func (c_ CNNLogSoftMaxNode) Init() CNNLogSoftMaxNode {
	rv := objc.Call[CNNLogSoftMaxNode](c_, objc.Sel("init"))
	return rv
}
