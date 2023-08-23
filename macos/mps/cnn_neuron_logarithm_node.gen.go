// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronLogarithmNode] class.
var CNNNeuronLogarithmNodeClass = _CNNNeuronLogarithmNodeClass{objc.GetClass("MPSCNNNeuronLogarithmNode")}

type _CNNNeuronLogarithmNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronLogarithmNode] class.
type ICNNNeuronLogarithmNode interface {
	ICNNNeuronNode
}

// A representation of a logarithm neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlogarithmnode?language=objc
type CNNNeuronLogarithmNode struct {
	CNNNeuronNode
}

func CNNNeuronLogarithmNodeFrom(ptr unsafe.Pointer) CNNNeuronLogarithmNode {
	return CNNNeuronLogarithmNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}

func (cc _CNNNeuronLogarithmNodeClass) NodeWithSource(sourceNode INNImageNode) CNNNeuronLogarithmNode {
	rv := objc.Call[CNNNeuronLogarithmNode](cc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlogarithmnode/2951931-nodewithsource?language=objc
func CNNNeuronLogarithmNode_NodeWithSource(sourceNode INNImageNode) CNNNeuronLogarithmNode {
	return CNNNeuronLogarithmNodeClass.NodeWithSource(sourceNode)
}

func (c_ CNNNeuronLogarithmNode) InitWithSource(sourceNode INNImageNode) CNNNeuronLogarithmNode {
	rv := objc.Call[CNNNeuronLogarithmNode](c_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlogarithmnode/2951949-initwithsource?language=objc
func NewCNNNeuronLogarithmNodeWithSource(sourceNode INNImageNode) CNNNeuronLogarithmNode {
	instance := CNNNeuronLogarithmNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronLogarithmNodeClass) Alloc() CNNNeuronLogarithmNode {
	rv := objc.Call[CNNNeuronLogarithmNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronLogarithmNode_Alloc() CNNNeuronLogarithmNode {
	return CNNNeuronLogarithmNodeClass.Alloc()
}

func (cc _CNNNeuronLogarithmNodeClass) New() CNNNeuronLogarithmNode {
	rv := objc.Call[CNNNeuronLogarithmNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronLogarithmNode() CNNNeuronLogarithmNode {
	return CNNNeuronLogarithmNodeClass.New()
}

func (c_ CNNNeuronLogarithmNode) Init() CNNNeuronLogarithmNode {
	rv := objc.Call[CNNNeuronLogarithmNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNNeuronLogarithmNodeClass) NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronLogarithmNode {
	rv := objc.Call[CNNNeuronLogarithmNode](cc, objc.Sel("nodeWithSource:descriptor:"), objc.Ptr(sourceNode), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/3019333-nodewithsource?language=objc
func CNNNeuronLogarithmNode_NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronLogarithmNode {
	return CNNNeuronLogarithmNodeClass.NodeWithSourceDescriptor(sourceNode, descriptor)
}
