// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronNode] class.
var CNNNeuronNodeClass = _CNNNeuronNodeClass{objc.GetClass("MPSCNNNeuronNode")}

type _CNNNeuronNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronNode] class.
type ICNNNeuronNode interface {
	INNFilterNode
	A() float64
	C() float64
	B() float64
}

// The virtual base class for MPS CNN neuron nodes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode?language=objc
type CNNNeuronNode struct {
	NNFilterNode
}

func CNNNeuronNodeFrom(ptr unsafe.Pointer) CNNNeuronNode {
	return CNNNeuronNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (cc _CNNNeuronNodeClass) NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronNode {
	rv := objc.Call[CNNNeuronNode](cc, objc.Sel("nodeWithSource:descriptor:"), objc.Ptr(sourceNode), objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/3019333-nodewithsource?language=objc
func CNNNeuronNode_NodeWithSourceDescriptor(sourceNode INNImageNode, descriptor INNNeuronDescriptor) CNNNeuronNode {
	return CNNNeuronNodeClass.NodeWithSourceDescriptor(sourceNode, descriptor)
}

func (cc _CNNNeuronNodeClass) Alloc() CNNNeuronNode {
	rv := objc.Call[CNNNeuronNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronNode_Alloc() CNNNeuronNode {
	return CNNNeuronNodeClass.Alloc()
}

func (cc _CNNNeuronNodeClass) New() CNNNeuronNode {
	rv := objc.Call[CNNNeuronNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronNode() CNNNeuronNode {
	return CNNNeuronNodeClass.New()
}

func (c_ CNNNeuronNode) Init() CNNNeuronNode {
	rv := objc.Call[CNNNeuronNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/2921459-a?language=objc
func (c_ CNNNeuronNode) A() float64 {
	rv := objc.Call[float64](c_, objc.Sel("a"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/2935553-c?language=objc
func (c_ CNNNeuronNode) C() float64 {
	rv := objc.Call[float64](c_, objc.Sel("c"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronnode/2921461-b?language=objc
func (c_ CNNNeuronNode) B() float64 {
	rv := objc.Call[float64](c_, objc.Sel("b"))
	return rv
}
