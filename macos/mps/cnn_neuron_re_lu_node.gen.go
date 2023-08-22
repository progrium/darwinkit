// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronReLUNode] class.
var CNNNeuronReLUNodeClass = _CNNNeuronReLUNodeClass{objc.GetClass("MPSCNNNeuronReLUNode")}

type _CNNNeuronReLUNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronReLUNode] class.
type ICNNNeuronReLUNode interface {
	ICNNNeuronNode
}

// A representation a ReLU neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunode?language=objc
type CNNNeuronReLUNode struct {
	CNNNeuronNode
}

func CNNNeuronReLUNodeFrom(ptr unsafe.Pointer) CNNNeuronReLUNode {
	return CNNNeuronReLUNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}

func (cc _CNNNeuronReLUNodeClass) NodeWithSource(sourceNode INNImageNode) CNNNeuronReLUNode {
	rv := objc.Call[CNNNeuronReLUNode](cc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunode/2921460-nodewithsource?language=objc
func CNNNeuronReLUNode_NodeWithSource(sourceNode INNImageNode) CNNNeuronReLUNode {
	return CNNNeuronReLUNodeClass.NodeWithSource(sourceNode)
}

func (c_ CNNNeuronReLUNode) InitWithSource(sourceNode INNImageNode) CNNNeuronReLUNode {
	rv := objc.Call[CNNNeuronReLUNode](c_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelunode/2921464-initwithsource?language=objc
func NewCNNNeuronReLUNodeWithSource(sourceNode INNImageNode) CNNNeuronReLUNode {
	instance := CNNNeuronReLUNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronReLUNodeClass) Alloc() CNNNeuronReLUNode {
	rv := objc.Call[CNNNeuronReLUNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronReLUNode_Alloc() CNNNeuronReLUNode {
	return CNNNeuronReLUNodeClass.Alloc()
}

func (cc _CNNNeuronReLUNodeClass) New() CNNNeuronReLUNode {
	rv := objc.Call[CNNNeuronReLUNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronReLUNode() CNNNeuronReLUNode {
	return CNNNeuronReLUNodeClass.New()
}

func (c_ CNNNeuronReLUNode) Init() CNNNeuronReLUNode {
	rv := objc.Call[CNNNeuronReLUNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNNeuronReLUNodeClass) NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronReLUNode {
	rv := objc.Call[CNNNeuronReLUNode](cc, objc.Sel("nodeWithSource:descriptor:"), objc.Ptr(sourceNode), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/3019333-nodewithsource?language=objc
func CNNNeuronReLUNode_NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronReLUNode {
	return CNNNeuronReLUNodeClass.NodeWithSourceDescriptor(sourceNode, descriptor)
}
