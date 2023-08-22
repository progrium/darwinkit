// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronGeLUNode] class.
var CNNNeuronGeLUNodeClass = _CNNNeuronGeLUNodeClass{objc.GetClass("MPSCNNNeuronGeLUNode")}

type _CNNNeuronGeLUNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronGeLUNode] class.
type ICNNNeuronGeLUNode interface {
	ICNNNeuronNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongelunode?language=objc
type CNNNeuronGeLUNode struct {
	CNNNeuronNode
}

func CNNNeuronGeLUNodeFrom(ptr unsafe.Pointer) CNNNeuronGeLUNode {
	return CNNNeuronGeLUNode{
		CNNNeuronNode: CNNNeuronNodeFrom(ptr),
	}
}

func (cc _CNNNeuronGeLUNodeClass) NodeWithSource(sourceNode INNImageNode) CNNNeuronGeLUNode {
	rv := objc.Call[CNNNeuronGeLUNode](cc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongelunode/3237267-nodewithsource?language=objc
func CNNNeuronGeLUNode_NodeWithSource(sourceNode INNImageNode) CNNNeuronGeLUNode {
	return CNNNeuronGeLUNodeClass.NodeWithSource(sourceNode)
}

func (c_ CNNNeuronGeLUNode) InitWithSource(sourceNode INNImageNode) CNNNeuronGeLUNode {
	rv := objc.Call[CNNNeuronGeLUNode](c_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurongelunode/3237266-initwithsource?language=objc
func NewCNNNeuronGeLUNodeWithSource(sourceNode INNImageNode) CNNNeuronGeLUNode {
	instance := CNNNeuronGeLUNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronGeLUNodeClass) Alloc() CNNNeuronGeLUNode {
	rv := objc.Call[CNNNeuronGeLUNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronGeLUNode_Alloc() CNNNeuronGeLUNode {
	return CNNNeuronGeLUNodeClass.Alloc()
}

func (cc _CNNNeuronGeLUNodeClass) New() CNNNeuronGeLUNode {
	rv := objc.Call[CNNNeuronGeLUNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronGeLUNode() CNNNeuronGeLUNode {
	return CNNNeuronGeLUNodeClass.New()
}

func (c_ CNNNeuronGeLUNode) Init() CNNNeuronGeLUNode {
	rv := objc.Call[CNNNeuronGeLUNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNNeuronGeLUNodeClass) NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronGeLUNode {
	rv := objc.Call[CNNNeuronGeLUNode](cc, objc.Sel("nodeWithSource:descriptor:"), objc.Ptr(sourceNode), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/3019333-nodewithsource?language=objc
func CNNNeuronGeLUNode_NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronGeLUNode {
	return CNNNeuronGeLUNodeClass.NodeWithSourceDescriptor(sourceNode, descriptor)
}
