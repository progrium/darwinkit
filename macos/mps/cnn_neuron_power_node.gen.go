// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronPowerNode] class.
var CNNNeuronPowerNodeClass = _CNNNeuronPowerNodeClass{objc.GetClass("MPSCNNNeuronPowerNode")}

type _CNNNeuronPowerNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronPowerNode] class.
type ICNNNeuronPowerNode interface {
	ICNNNeuronNode
}

// A representation of a power neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronpowernode?language=objc
type CNNNeuronPowerNode struct {
	CNNNeuronNode
}

func CNNNeuronPowerNodeFrom(ptr unsafe.Pointer) CNNNeuronPowerNode {
	return CNNNeuronPowerNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}

func (cc _CNNNeuronPowerNodeClass) NodeWithSource(sourceNode INNImageNode) CNNNeuronPowerNode {
	rv := objc.Call[CNNNeuronPowerNode](cc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronpowernode/2951958-nodewithsource?language=objc
func CNNNeuronPowerNode_NodeWithSource(sourceNode INNImageNode) CNNNeuronPowerNode {
	return CNNNeuronPowerNodeClass.NodeWithSource(sourceNode)
}

func (c_ CNNNeuronPowerNode) InitWithSource(sourceNode INNImageNode) CNNNeuronPowerNode {
	rv := objc.Call[CNNNeuronPowerNode](c_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronpowernode/2951937-initwithsource?language=objc
func NewCNNNeuronPowerNodeWithSource(sourceNode INNImageNode) CNNNeuronPowerNode {
	instance := CNNNeuronPowerNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronPowerNodeClass) Alloc() CNNNeuronPowerNode {
	rv := objc.Call[CNNNeuronPowerNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronPowerNode_Alloc() CNNNeuronPowerNode {
	return CNNNeuronPowerNodeClass.Alloc()
}

func (cc _CNNNeuronPowerNodeClass) New() CNNNeuronPowerNode {
	rv := objc.Call[CNNNeuronPowerNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronPowerNode() CNNNeuronPowerNode {
	return CNNNeuronPowerNodeClass.New()
}

func (c_ CNNNeuronPowerNode) Init() CNNNeuronPowerNode {
	rv := objc.Call[CNNNeuronPowerNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNNeuronPowerNodeClass) NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronPowerNode {
	rv := objc.Call[CNNNeuronPowerNode](cc, objc.Sel("nodeWithSource:descriptor:"), objc.Ptr(sourceNode), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/3019333-nodewithsource?language=objc
func CNNNeuronPowerNode_NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronPowerNode {
	return CNNNeuronPowerNodeClass.NodeWithSourceDescriptor(sourceNode, descriptor)
}
