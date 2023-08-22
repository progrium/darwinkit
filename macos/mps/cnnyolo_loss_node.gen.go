// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNYOLOLossNode] class.
var CNNYOLOLossNodeClass = _CNNYOLOLossNodeClass{objc.GetClass("MPSCNNYOLOLossNode")}

type _CNNYOLOLossNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNYOLOLossNode] class.
type ICNNYOLOLossNode interface {
	INNFilterNode
	InputLabels() NNLabelsNode
}

// A representation of a YOLO loss kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossnode?language=objc
type CNNYOLOLossNode struct {
	NNFilterNode
}

func CNNYOLOLossNodeFrom(ptr unsafe.Pointer) CNNYOLOLossNode {
	return CNNYOLOLossNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (cc _CNNYOLOLossNodeClass) NodeWithSourceLossDescriptor(source INNImageNode, descriptor ICNNYOLOLossDescriptor) CNNYOLOLossNode {
	rv := objc.Call[CNNYOLOLossNode](cc, objc.Sel("nodeWithSource:lossDescriptor:"), objc.Ptr(source), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossnode/2976516-nodewithsource?language=objc
func CNNYOLOLossNode_NodeWithSourceLossDescriptor(source INNImageNode, descriptor ICNNYOLOLossDescriptor) CNNYOLOLossNode {
	return CNNYOLOLossNodeClass.NodeWithSourceLossDescriptor(source, descriptor)
}

func (c_ CNNYOLOLossNode) InitWithSourceLossDescriptor(source INNImageNode, descriptor ICNNYOLOLossDescriptor) CNNYOLOLossNode {
	rv := objc.Call[CNNYOLOLossNode](c_, objc.Sel("initWithSource:lossDescriptor:"), objc.Ptr(source), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossnode/2976514-initwithsource?language=objc
func NewCNNYOLOLossNodeWithSourceLossDescriptor(source INNImageNode, descriptor ICNNYOLOLossDescriptor) CNNYOLOLossNode {
	instance := CNNYOLOLossNodeClass.Alloc().InitWithSourceLossDescriptor(source, descriptor)
	instance.Autorelease()
	return instance
}

func (cc _CNNYOLOLossNodeClass) Alloc() CNNYOLOLossNode {
	rv := objc.Call[CNNYOLOLossNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNYOLOLossNode_Alloc() CNNYOLOLossNode {
	return CNNYOLOLossNodeClass.Alloc()
}

func (cc _CNNYOLOLossNodeClass) New() CNNYOLOLossNode {
	rv := objc.Call[CNNYOLOLossNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNYOLOLossNode() CNNYOLOLossNode {
	return CNNYOLOLossNodeClass.New()
}

func (c_ CNNYOLOLossNode) Init() CNNYOLOLossNode {
	rv := objc.Call[CNNYOLOLossNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnyololossnode/2976515-inputlabels?language=objc
func (c_ CNNYOLOLossNode) InputLabels() NNLabelsNode {
	rv := objc.Call[NNLabelsNode](c_, objc.Sel("inputLabels"))
	return rv
}
