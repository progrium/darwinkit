// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronSoftSignNode] class.
var CNNNeuronSoftSignNodeClass = _CNNNeuronSoftSignNodeClass{objc.GetClass("MPSCNNNeuronSoftSignNode")}

type _CNNNeuronSoftSignNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronSoftSignNode] class.
type ICNNNeuronSoftSignNode interface {
	ICNNNeuronNode
}

// A representation of a softsign neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftsignnode?language=objc
type CNNNeuronSoftSignNode struct {
	CNNNeuronNode
}

func CNNNeuronSoftSignNodeFrom(ptr unsafe.Pointer) CNNNeuronSoftSignNode {
	return CNNNeuronSoftSignNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}

func (cc _CNNNeuronSoftSignNodeClass) NodeWithSource(sourceNode INNImageNode) CNNNeuronSoftSignNode {
	rv := objc.Call[CNNNeuronSoftSignNode](cc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftsignnode/2866428-nodewithsource?language=objc
func CNNNeuronSoftSignNode_NodeWithSource(sourceNode INNImageNode) CNNNeuronSoftSignNode {
	return CNNNeuronSoftSignNodeClass.NodeWithSource(sourceNode)
}

func (c_ CNNNeuronSoftSignNode) InitWithSource(sourceNode INNImageNode) CNNNeuronSoftSignNode {
	rv := objc.Call[CNNNeuronSoftSignNode](c_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftsignnode/2921463-initwithsource?language=objc
func NewCNNNeuronSoftSignNodeWithSource(sourceNode INNImageNode) CNNNeuronSoftSignNode {
	instance := CNNNeuronSoftSignNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronSoftSignNodeClass) Alloc() CNNNeuronSoftSignNode {
	rv := objc.Call[CNNNeuronSoftSignNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronSoftSignNode_Alloc() CNNNeuronSoftSignNode {
	return CNNNeuronSoftSignNodeClass.Alloc()
}

func (cc _CNNNeuronSoftSignNodeClass) New() CNNNeuronSoftSignNode {
	rv := objc.Call[CNNNeuronSoftSignNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronSoftSignNode() CNNNeuronSoftSignNode {
	return CNNNeuronSoftSignNodeClass.New()
}

func (c_ CNNNeuronSoftSignNode) Init() CNNNeuronSoftSignNode {
	rv := objc.Call[CNNNeuronSoftSignNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNNeuronSoftSignNodeClass) NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronSoftSignNode {
	rv := objc.Call[CNNNeuronSoftSignNode](cc, objc.Sel("nodeWithSource:descriptor:"), objc.Ptr(sourceNode), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/3019333-nodewithsource?language=objc
func CNNNeuronSoftSignNode_NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronSoftSignNode {
	return CNNNeuronSoftSignNodeClass.NodeWithSourceDescriptor(sourceNode, descriptor)
}
