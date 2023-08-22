// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNConvolutionTransposeNode] class.
var CNNConvolutionTransposeNodeClass = _CNNConvolutionTransposeNodeClass{objc.GetClass("MPSCNNConvolutionTransposeNode")}

type _CNNConvolutionTransposeNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNConvolutionTransposeNode] class.
type ICNNConvolutionTransposeNode interface {
	ICNNConvolutionNode
}

// A representation of a transposed convolution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposenode?language=objc
type CNNConvolutionTransposeNode struct {
	CNNConvolutionNode
}

func CNNConvolutionTransposeNodeFrom(ptr unsafe.Pointer) CNNConvolutionTransposeNode {
	return CNNConvolutionTransposeNode{
		CNNConvolutionNode: CNNConvolutionNodeFrom(ptr),
	}
}

func (cc _CNNConvolutionTransposeNodeClass) NodeWithSourceConvolutionGradientStateWeights(sourceNode INNImageNode, convolutionGradientState ICNNConvolutionGradientStateNode, weights PCNNConvolutionDataSource) CNNConvolutionTransposeNode {
	po2 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNConvolutionTransposeNode](cc, objc.Sel("nodeWithSource:convolutionGradientState:weights:"), objc.Ptr(sourceNode), objc.Ptr(convolutionGradientState), po2)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposenode/2942636-nodewithsource?language=objc
func CNNConvolutionTransposeNode_NodeWithSourceConvolutionGradientStateWeights(sourceNode INNImageNode, convolutionGradientState ICNNConvolutionGradientStateNode, weights PCNNConvolutionDataSource) CNNConvolutionTransposeNode {
	return CNNConvolutionTransposeNodeClass.NodeWithSourceConvolutionGradientStateWeights(sourceNode, convolutionGradientState, weights)
}

func (c_ CNNConvolutionTransposeNode) InitWithSourceConvolutionGradientStateWeights(sourceNode INNImageNode, convolutionGradientState ICNNConvolutionGradientStateNode, weights PCNNConvolutionDataSource) CNNConvolutionTransposeNode {
	po2 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNConvolutionTransposeNode](c_, objc.Sel("initWithSource:convolutionGradientState:weights:"), objc.Ptr(sourceNode), objc.Ptr(convolutionGradientState), po2)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposenode/2942641-initwithsource?language=objc
func NewCNNConvolutionTransposeNodeWithSourceConvolutionGradientStateWeights(sourceNode INNImageNode, convolutionGradientState ICNNConvolutionGradientStateNode, weights PCNNConvolutionDataSource) CNNConvolutionTransposeNode {
	instance := CNNConvolutionTransposeNodeClass.Alloc().InitWithSourceConvolutionGradientStateWeights(sourceNode, convolutionGradientState, weights)
	instance.Autorelease()
	return instance
}

func (cc _CNNConvolutionTransposeNodeClass) Alloc() CNNConvolutionTransposeNode {
	rv := objc.Call[CNNConvolutionTransposeNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNConvolutionTransposeNode_Alloc() CNNConvolutionTransposeNode {
	return CNNConvolutionTransposeNodeClass.Alloc()
}

func (cc _CNNConvolutionTransposeNodeClass) New() CNNConvolutionTransposeNode {
	rv := objc.Call[CNNConvolutionTransposeNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNConvolutionTransposeNode() CNNConvolutionTransposeNode {
	return CNNConvolutionTransposeNodeClass.New()
}

func (c_ CNNConvolutionTransposeNode) Init() CNNConvolutionTransposeNode {
	rv := objc.Call[CNNConvolutionTransposeNode](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNConvolutionTransposeNodeClass) NodeWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNConvolutionTransposeNode {
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNConvolutionTransposeNode](cc, objc.Sel("nodeWithSource:weights:"), objc.Ptr(sourceNode), po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2866436-nodewithsource?language=objc
func CNNConvolutionTransposeNode_NodeWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNConvolutionTransposeNode {
	return CNNConvolutionTransposeNodeClass.NodeWithSourceWeights(sourceNode, weights)
}

func (c_ CNNConvolutionTransposeNode) InitWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNConvolutionTransposeNode {
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNConvolutionTransposeNode](c_, objc.Sel("initWithSource:weights:"), objc.Ptr(sourceNode), po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2866470-initwithsource?language=objc
func NewCNNConvolutionTransposeNodeWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNConvolutionTransposeNode {
	instance := CNNConvolutionTransposeNodeClass.Alloc().InitWithSourceWeights(sourceNode, weights)
	instance.Autorelease()
	return instance
}
