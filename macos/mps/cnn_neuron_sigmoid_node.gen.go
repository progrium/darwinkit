// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronSigmoidNode] class.
var CNNNeuronSigmoidNodeClass = _CNNNeuronSigmoidNodeClass{objc.GetClass("MPSCNNNeuronSigmoidNode")}

type _CNNNeuronSigmoidNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronSigmoidNode] class.
type ICNNNeuronSigmoidNode interface {
	ICNNNeuronNode
}

// A representation of a sigmoid neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsigmoidnode?language=objc
type CNNNeuronSigmoidNode struct {
	CNNNeuronNode
}

func CNNNeuronSigmoidNodeFrom(ptr unsafe.Pointer) CNNNeuronSigmoidNode {
	return CNNNeuronSigmoidNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}

func (cc _CNNNeuronSigmoidNodeClass) NodeWithSource(sourceNode INNImageNode) CNNNeuronSigmoidNode {
	rv := objc.Call[CNNNeuronSigmoidNode](cc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsigmoidnode/2866467-nodewithsource?language=objc
func CNNNeuronSigmoidNode_NodeWithSource(sourceNode INNImageNode) CNNNeuronSigmoidNode {
	return CNNNeuronSigmoidNodeClass.NodeWithSource(sourceNode)
}

func (c_ CNNNeuronSigmoidNode) InitWithSource(sourceNode INNImageNode) CNNNeuronSigmoidNode {
	rv := objc.Call[CNNNeuronSigmoidNode](c_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsigmoidnode/2921458-initwithsource?language=objc
func NewCNNNeuronSigmoidNodeWithSource(sourceNode INNImageNode) CNNNeuronSigmoidNode {
	instance := CNNNeuronSigmoidNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronSigmoidNodeClass) Alloc() CNNNeuronSigmoidNode {
	rv := objc.Call[CNNNeuronSigmoidNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronSigmoidNode_Alloc() CNNNeuronSigmoidNode {
	return CNNNeuronSigmoidNodeClass.Alloc()
}

func (cc _CNNNeuronSigmoidNodeClass) New() CNNNeuronSigmoidNode {
	rv := objc.Call[CNNNeuronSigmoidNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronSigmoidNode() CNNNeuronSigmoidNode {
	return CNNNeuronSigmoidNodeClass.New()
}

func (c_ CNNNeuronSigmoidNode) Init() CNNNeuronSigmoidNode {
	rv := objc.Call[CNNNeuronSigmoidNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNNeuronSigmoidNodeClass) NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronSigmoidNode {
	rv := objc.Call[CNNNeuronSigmoidNode](cc, objc.Sel("nodeWithSource:descriptor:"), objc.Ptr(sourceNode), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/3019333-nodewithsource?language=objc
func CNNNeuronSigmoidNode_NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronSigmoidNode {
	return CNNNeuronSigmoidNodeClass.NodeWithSourceDescriptor(sourceNode, descriptor)
}
