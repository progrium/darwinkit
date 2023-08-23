// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNLocalContrastNormalizationNode] class.
var CNNLocalContrastNormalizationNodeClass = _CNNLocalContrastNormalizationNodeClass{objc.GetClass("MPSCNNLocalContrastNormalizationNode")}

type _CNNLocalContrastNormalizationNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNLocalContrastNormalizationNode] class.
type ICNNLocalContrastNormalizationNode interface {
	ICNNNormalizationNode
	Ps() float64
	SetPs(value float64)
	P0() float64
	SetP0(value float64)
	Pm() float64
	SetPm(value float64)
	KernelHeight() uint
	SetKernelHeight(value uint)
	KernelWidth() uint
	SetKernelWidth(value uint)
}

// A representation of a local-contrast normalization kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode?language=objc
type CNNLocalContrastNormalizationNode struct {
	CNNNormalizationNode
}

func CNNLocalContrastNormalizationNodeFrom(ptr unsafe.Pointer) CNNLocalContrastNormalizationNode {
	return CNNLocalContrastNormalizationNode{
		CNNNormalizationNode: CNNNormalizationNodeFrom(ptr),
	}
}

func (cc _CNNLocalContrastNormalizationNodeClass) NodeWithSourceKernelSize(sourceNode INNImageNode, kernelSize uint) CNNLocalContrastNormalizationNode {
	rv := objc.Call[CNNLocalContrastNormalizationNode](cc, objc.Sel("nodeWithSource:kernelSize:"), objc.Ptr(sourceNode), kernelSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866439-nodewithsource?language=objc
func CNNLocalContrastNormalizationNode_NodeWithSourceKernelSize(sourceNode INNImageNode, kernelSize uint) CNNLocalContrastNormalizationNode {
	return CNNLocalContrastNormalizationNodeClass.NodeWithSourceKernelSize(sourceNode, kernelSize)
}

func (c_ CNNLocalContrastNormalizationNode) InitWithSource(sourceNode INNImageNode) CNNLocalContrastNormalizationNode {
	rv := objc.Call[CNNLocalContrastNormalizationNode](c_, objc.Sel("initWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866454-initwithsource?language=objc
func NewCNNLocalContrastNormalizationNodeWithSource(sourceNode INNImageNode) CNNLocalContrastNormalizationNode {
	instance := CNNLocalContrastNormalizationNodeClass.Alloc().InitWithSource(sourceNode)
	instance.Autorelease()
	return instance
}

func (cc _CNNLocalContrastNormalizationNodeClass) Alloc() CNNLocalContrastNormalizationNode {
	rv := objc.Call[CNNLocalContrastNormalizationNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNLocalContrastNormalizationNode_Alloc() CNNLocalContrastNormalizationNode {
	return CNNLocalContrastNormalizationNodeClass.Alloc()
}

func (cc _CNNLocalContrastNormalizationNodeClass) New() CNNLocalContrastNormalizationNode {
	rv := objc.Call[CNNLocalContrastNormalizationNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNLocalContrastNormalizationNode() CNNLocalContrastNormalizationNode {
	return CNNLocalContrastNormalizationNodeClass.New()
}

func (c_ CNNLocalContrastNormalizationNode) Init() CNNLocalContrastNormalizationNode {
	rv := objc.Call[CNNLocalContrastNormalizationNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNLocalContrastNormalizationNodeClass) NodeWithSource(sourceNode INNImageNode) CNNLocalContrastNormalizationNode {
	rv := objc.Call[CNNLocalContrastNormalizationNode](cc, objc.Sel("nodeWithSource:"), objc.Ptr(sourceNode))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationnode/2866460-nodewithsource?language=objc
func CNNLocalContrastNormalizationNode_NodeWithSource(sourceNode INNImageNode) CNNLocalContrastNormalizationNode {
	return CNNLocalContrastNormalizationNodeClass.NodeWithSource(sourceNode)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866500-ps?language=objc
func (c_ CNNLocalContrastNormalizationNode) Ps() float64 {
	rv := objc.Call[float64](c_, objc.Sel("ps"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866500-ps?language=objc
func (c_ CNNLocalContrastNormalizationNode) SetPs(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setPs:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866510-p0?language=objc
func (c_ CNNLocalContrastNormalizationNode) P0() float64 {
	rv := objc.Call[float64](c_, objc.Sel("p0"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866510-p0?language=objc
func (c_ CNNLocalContrastNormalizationNode) SetP0(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setP0:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866404-pm?language=objc
func (c_ CNNLocalContrastNormalizationNode) Pm() float64 {
	rv := objc.Call[float64](c_, objc.Sel("pm"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866404-pm?language=objc
func (c_ CNNLocalContrastNormalizationNode) SetPm(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setPm:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866485-kernelheight?language=objc
func (c_ CNNLocalContrastNormalizationNode) KernelHeight() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelHeight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866485-kernelheight?language=objc
func (c_ CNNLocalContrastNormalizationNode) SetKernelHeight(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setKernelHeight:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866441-kernelwidth?language=objc
func (c_ CNNLocalContrastNormalizationNode) KernelWidth() uint {
	rv := objc.Call[uint](c_, objc.Sel("kernelWidth"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnlocalcontrastnormalizationnode/2866441-kernelwidth?language=objc
func (c_ CNNLocalContrastNormalizationNode) SetKernelWidth(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setKernelWidth:"), value)
}
