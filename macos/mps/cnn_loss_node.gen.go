// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNLossNode] class.
var CNNLossNodeClass = _CNNLossNodeClass{objc.GetClass("MPSCNNLossNode")}

type _CNNLossNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNLossNode] class.
type ICNNLossNode interface {
	INNFilterNode
	InputLabels() NNLabelsNode
}

// A representation of a loss kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossnode?language=objc
type CNNLossNode struct {
	NNFilterNode
}

func CNNLossNodeFrom(ptr unsafe.Pointer) CNNLossNode {
	return CNNLossNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (cc _CNNLossNodeClass) NodeWithSourceLossDescriptor(source INNImageNode, descriptor ICNNLossDescriptor) CNNLossNode {
	rv := objc.Call[CNNLossNode](cc, objc.Sel("nodeWithSource:lossDescriptor:"), objc.Ptr(source), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossnode/2951956-nodewithsource?language=objc
func CNNLossNode_NodeWithSourceLossDescriptor(source INNImageNode, descriptor ICNNLossDescriptor) CNNLossNode {
	return CNNLossNodeClass.NodeWithSourceLossDescriptor(source, descriptor)
}

func (c_ CNNLossNode) InitWithSourceLossDescriptor(source INNImageNode, descriptor ICNNLossDescriptor) CNNLossNode {
	rv := objc.Call[CNNLossNode](c_, objc.Sel("initWithSource:lossDescriptor:"), objc.Ptr(source), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossnode/2951947-initwithsource?language=objc
func NewCNNLossNodeWithSourceLossDescriptor(source INNImageNode, descriptor ICNNLossDescriptor) CNNLossNode {
	instance := CNNLossNodeClass.Alloc().InitWithSourceLossDescriptor(source, descriptor)
	instance.Autorelease()
	return instance
}

func (cc _CNNLossNodeClass) Alloc() CNNLossNode {
	rv := objc.Call[CNNLossNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNLossNode_Alloc() CNNLossNode {
	return CNNLossNodeClass.Alloc()
}

func (cc _CNNLossNodeClass) New() CNNLossNode {
	rv := objc.Call[CNNLossNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNLossNode() CNNLossNode {
	return CNNLossNodeClass.New()
}

func (c_ CNNLossNode) Init() CNNLossNode {
	rv := objc.Call[CNNLossNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlossnode/2951942-inputlabels?language=objc
func (c_ CNNLossNode) InputLabels() NNLabelsNode {
	rv := objc.Call[NNLabelsNode](c_, objc.Sel("inputLabels"))
	return rv
}
