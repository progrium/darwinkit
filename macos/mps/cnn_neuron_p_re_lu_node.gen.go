// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronPReLUNode] class.
var CNNNeuronPReLUNodeClass = _CNNNeuronPReLUNodeClass{objc.GetClass("MPSCNNNeuronPReLUNode")}

type _CNNNeuronPReLUNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronPReLUNode] class.
type ICNNNeuronPReLUNode interface {
	ICNNNeuronNode
}

// A representation a PReLU neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronprelunode?language=objc
type CNNNeuronPReLUNode struct {
	CNNNeuronNode
}

func CNNNeuronPReLUNodeFrom(ptr unsafe.Pointer) CNNNeuronPReLUNode {
	return CNNNeuronPReLUNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}

func (cc _CNNNeuronPReLUNodeClass) NodeWithSourceAData(sourceNode INNImageNode, aData []byte) CNNNeuronPReLUNode {
	rv := objc.Call[CNNNeuronPReLUNode](cc, objc.Sel("nodeWithSource:aData:"), objc.Ptr(sourceNode), aData)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronprelunode/2921597-nodewithsource?language=objc
func CNNNeuronPReLUNode_NodeWithSourceAData(sourceNode INNImageNode, aData []byte) CNNNeuronPReLUNode {
	return CNNNeuronPReLUNodeClass.NodeWithSourceAData(sourceNode, aData)
}

func (c_ CNNNeuronPReLUNode) InitWithSourceAData(sourceNode INNImageNode, aData []byte) CNNNeuronPReLUNode {
	rv := objc.Call[CNNNeuronPReLUNode](c_, objc.Sel("initWithSource:aData:"), objc.Ptr(sourceNode), aData)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronprelunode/2921595-initwithsource?language=objc
func NewCNNNeuronPReLUNodeWithSourceAData(sourceNode INNImageNode, aData []byte) CNNNeuronPReLUNode {
	instance := CNNNeuronPReLUNodeClass.Alloc().InitWithSourceAData(sourceNode, aData)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronPReLUNodeClass) Alloc() CNNNeuronPReLUNode {
	rv := objc.Call[CNNNeuronPReLUNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronPReLUNode_Alloc() CNNNeuronPReLUNode {
	return CNNNeuronPReLUNodeClass.Alloc()
}

func (cc _CNNNeuronPReLUNodeClass) New() CNNNeuronPReLUNode {
	rv := objc.Call[CNNNeuronPReLUNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronPReLUNode() CNNNeuronPReLUNode {
	return CNNNeuronPReLUNodeClass.New()
}

func (c_ CNNNeuronPReLUNode) Init() CNNNeuronPReLUNode {
	rv := objc.Call[CNNNeuronPReLUNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNNeuronPReLUNodeClass) NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronPReLUNode {
	rv := objc.Call[CNNNeuronPReLUNode](cc, objc.Sel("nodeWithSource:descriptor:"), objc.Ptr(sourceNode), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/3019333-nodewithsource?language=objc
func CNNNeuronPReLUNode_NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronPReLUNode {
	return CNNNeuronPReLUNodeClass.NodeWithSourceDescriptor(sourceNode, descriptor)
}
