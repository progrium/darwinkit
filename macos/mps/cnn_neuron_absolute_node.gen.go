// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronAbsoluteNode] class.
var CNNNeuronAbsoluteNodeClass = _CNNNeuronAbsoluteNodeClass{objc.GetClass("MPSCNNNeuronAbsoluteNode")}

type _CNNNeuronAbsoluteNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronAbsoluteNode] class.
type ICNNNeuronAbsoluteNode interface {
	ICNNNeuronNode
}

// A representation of an absolute neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronabsolutenode?language=objc
type CNNNeuronAbsoluteNode struct {
	CNNNeuronNode
}

func CNNNeuronAbsoluteNodeFrom(ptr unsafe.Pointer) CNNNeuronAbsoluteNode {
	return CNNNeuronAbsoluteNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}

func (cc _CNNNeuronAbsoluteNodeClass) NodeWithSource(sourceNode INNImageNode) CNNNeuronAbsoluteNode {
	rv := objc.Call[CNNNeuronAbsoluteNode](cc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronabsolutenode/2866431-nodewithsource?language=objc
func CNNNeuronAbsoluteNode_NodeWithSource(sourceNode INNImageNode) CNNNeuronAbsoluteNode {
	return CNNNeuronAbsoluteNodeClass.NodeWithSource(sourceNode)
}

func (c_ CNNNeuronAbsoluteNode) InitWithSource(sourceNode INNImageNode) CNNNeuronAbsoluteNode {
	rv := objc.Call[CNNNeuronAbsoluteNode](c_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronabsolutenode/2921448-initwithsource?language=objc
func NewCNNNeuronAbsoluteNodeWithSource(sourceNode INNImageNode) CNNNeuronAbsoluteNode {
	instance := CNNNeuronAbsoluteNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronAbsoluteNodeClass) Alloc() CNNNeuronAbsoluteNode {
	rv := objc.Call[CNNNeuronAbsoluteNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronAbsoluteNode_Alloc() CNNNeuronAbsoluteNode {
	return CNNNeuronAbsoluteNodeClass.Alloc()
}

func (cc _CNNNeuronAbsoluteNodeClass) New() CNNNeuronAbsoluteNode {
	rv := objc.Call[CNNNeuronAbsoluteNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronAbsoluteNode() CNNNeuronAbsoluteNode {
	return CNNNeuronAbsoluteNodeClass.New()
}

func (c_ CNNNeuronAbsoluteNode) Init() CNNNeuronAbsoluteNode {
	rv := objc.Call[CNNNeuronAbsoluteNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNNeuronAbsoluteNodeClass) NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronAbsoluteNode {
	rv := objc.Call[CNNNeuronAbsoluteNode](cc, objc.Sel("nodeWithSource:descriptor:"), objc.Ptr(sourceNode), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/3019333-nodewithsource?language=objc
func CNNNeuronAbsoluteNode_NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronAbsoluteNode {
	return CNNNeuronAbsoluteNodeClass.NodeWithSourceDescriptor(sourceNode, descriptor)
}
