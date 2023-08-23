// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNDilatedPoolingMaxNode] class.
var CNNDilatedPoolingMaxNodeClass = _CNNDilatedPoolingMaxNodeClass{objc.GetClass("MPSCNNDilatedPoolingMaxNode")}

type _CNNDilatedPoolingMaxNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNDilatedPoolingMaxNode] class.
type ICNNDilatedPoolingMaxNode interface {
	INNFilterNode
	DilationRateY() uint
	DilationRateX() uint
}

// A representation of a dilated max pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxnode?language=objc
type CNNDilatedPoolingMaxNode struct {
	NNFilterNode
}

func CNNDilatedPoolingMaxNodeFrom(ptr unsafe.Pointer) CNNDilatedPoolingMaxNode {
	return CNNDilatedPoolingMaxNode{
		NNFilterNode: NNFilterNodeFrom(ptr),
	}
}

func (cc _CNNDilatedPoolingMaxNodeClass) NodeWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNDilatedPoolingMaxNode {
	rv := objc.Call[CNNDilatedPoolingMaxNode](cc, objc.Sel("nodeWithSource:filterSize:"), objc.Ptr(sourceNode), size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxnode/2873227-nodewithsource?language=objc
func CNNDilatedPoolingMaxNode_NodeWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNDilatedPoolingMaxNode {
	return CNNDilatedPoolingMaxNodeClass.NodeWithSourceFilterSize(sourceNode, size)
}

func (c_ CNNDilatedPoolingMaxNode) InitWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNDilatedPoolingMaxNode {
	rv := objc.Call[CNNDilatedPoolingMaxNode](c_, objc.Sel("initWithSource:filterSize:"), objc.Ptr(sourceNode), size)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxnode/2873240-initwithsource?language=objc
func NewCNNDilatedPoolingMaxNodeWithSourceFilterSize(sourceNode INNImageNode, size uint) CNNDilatedPoolingMaxNode {
	instance := CNNDilatedPoolingMaxNodeClass.Alloc().InitWithSourceFilterSize(sourceNode, size)
	instance.Autorelease()
	return instance
}

func (cc _CNNDilatedPoolingMaxNodeClass) Alloc() CNNDilatedPoolingMaxNode {
	rv := objc.Call[CNNDilatedPoolingMaxNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNDilatedPoolingMaxNode_Alloc() CNNDilatedPoolingMaxNode {
	return CNNDilatedPoolingMaxNodeClass.Alloc()
}

func (cc _CNNDilatedPoolingMaxNodeClass) New() CNNDilatedPoolingMaxNode {
	rv := objc.Call[CNNDilatedPoolingMaxNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNDilatedPoolingMaxNode() CNNDilatedPoolingMaxNode {
	return CNNDilatedPoolingMaxNodeClass.New()
}

func (c_ CNNDilatedPoolingMaxNode) Init() CNNDilatedPoolingMaxNode {
	rv := objc.Call[CNNDilatedPoolingMaxNode](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxnode/2887342-dilationratey?language=objc
func (c_ CNNDilatedPoolingMaxNode) DilationRateY() uint {
	rv := objc.Call[uint](c_, objc.Sel("dilationRateY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndilatedpoolingmaxnode/2887341-dilationratex?language=objc
func (c_ CNNDilatedPoolingMaxNode) DilationRateX() uint {
	rv := objc.Call[uint](c_, objc.Sel("dilationRateX"))
	return rv
}
