// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronELUNode] class.
var CNNNeuronELUNodeClass = _CNNNeuronELUNodeClass{objc.GetClass("MPSCNNNeuronELUNode")}

type _CNNNeuronELUNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronELUNode] class.
type ICNNNeuronELUNode interface {
	ICNNNeuronNode
}

// A representation of a parametric ELU neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronelunode?language=objc
type CNNNeuronELUNode struct {
	CNNNeuronNode
}

func CNNNeuronELUNodeFrom(ptr unsafe.Pointer) CNNNeuronELUNode {
	return CNNNeuronELUNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}

func (cc _CNNNeuronELUNodeClass) NodeWithSource(sourceNode INNImageNode) CNNNeuronELUNode {
	rv := objc.Call[CNNNeuronELUNode](cc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronelunode/2921452-nodewithsource?language=objc
func CNNNeuronELUNode_NodeWithSource(sourceNode INNImageNode) CNNNeuronELUNode {
	return CNNNeuronELUNodeClass.NodeWithSource(sourceNode)
}

func (c_ CNNNeuronELUNode) InitWithSource(sourceNode INNImageNode) CNNNeuronELUNode {
	rv := objc.Call[CNNNeuronELUNode](c_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronelunode/2921447-initwithsource?language=objc
func NewCNNNeuronELUNodeWithSource(sourceNode INNImageNode) CNNNeuronELUNode {
	instance := CNNNeuronELUNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronELUNodeClass) Alloc() CNNNeuronELUNode {
	rv := objc.Call[CNNNeuronELUNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronELUNode_Alloc() CNNNeuronELUNode {
	return CNNNeuronELUNodeClass.Alloc()
}

func (cc _CNNNeuronELUNodeClass) New() CNNNeuronELUNode {
	rv := objc.Call[CNNNeuronELUNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronELUNode() CNNNeuronELUNode {
	return CNNNeuronELUNodeClass.New()
}

func (c_ CNNNeuronELUNode) Init() CNNNeuronELUNode {
	rv := objc.Call[CNNNeuronELUNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNNeuronELUNodeClass) NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronELUNode {
	rv := objc.Call[CNNNeuronELUNode](cc, objc.Sel("nodeWithSource:descriptor:"), objc.Ptr(sourceNode), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/3019333-nodewithsource?language=objc
func CNNNeuronELUNode_NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronELUNode {
	return CNNNeuronELUNodeClass.NodeWithSourceDescriptor(sourceNode, descriptor)
}
