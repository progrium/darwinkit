// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNConvolutionTransposeGradientNode] class.
var CNNConvolutionTransposeGradientNodeClass = _CNNConvolutionTransposeGradientNodeClass{objc.GetClass("MPSCNNConvolutionTransposeGradientNode")}

type _CNNConvolutionTransposeGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNConvolutionTransposeGradientNode] class.
type ICNNConvolutionTransposeGradientNode interface {
	ICNNConvolutionGradientNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradientnode?language=objc
type CNNConvolutionTransposeGradientNode struct {
	CNNConvolutionGradientNode
}

func CNNConvolutionTransposeGradientNodeFrom(ptr unsafe.Pointer) CNNConvolutionTransposeGradientNode {
	return CNNConvolutionTransposeGradientNode{
		CNNConvolutionGradientNode: CNNConvolutionGradientNodeFrom(ptr),
	}
}

func (c_ CNNConvolutionTransposeGradientNode) InitWithSourceGradientSourceImageConvolutionTransposeGradientStateWeights(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState ICNNConvolutionTransposeGradientStateNode, weights PCNNConvolutionDataSource) CNNConvolutionTransposeGradientNode {
	po3 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNConvolutionTransposeGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:convolutionTransposeGradientState:weights:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), po3)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradientnode/3143550-initwithsourcegradient?language=objc
func NewCNNConvolutionTransposeGradientNodeWithSourceGradientSourceImageConvolutionTransposeGradientStateWeights(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState ICNNConvolutionTransposeGradientStateNode, weights PCNNConvolutionDataSource) CNNConvolutionTransposeGradientNode {
	instance := CNNConvolutionTransposeGradientNodeClass.Alloc().InitWithSourceGradientSourceImageConvolutionTransposeGradientStateWeights(sourceGradient, sourceImage, gradientState, weights)
	instance.Autorelease()
	return instance
}

func (cc _CNNConvolutionTransposeGradientNodeClass) NodeWithSourceGradientSourceImageConvolutionTransposeGradientStateWeights(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState ICNNConvolutionTransposeGradientStateNode, weights PCNNConvolutionDataSource) CNNConvolutionTransposeGradientNode {
	po3 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNConvolutionTransposeGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:convolutionTransposeGradientState:weights:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), po3)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradientnode/3143551-nodewithsourcegradient?language=objc
func CNNConvolutionTransposeGradientNode_NodeWithSourceGradientSourceImageConvolutionTransposeGradientStateWeights(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState ICNNConvolutionTransposeGradientStateNode, weights PCNNConvolutionDataSource) CNNConvolutionTransposeGradientNode {
	return CNNConvolutionTransposeGradientNodeClass.NodeWithSourceGradientSourceImageConvolutionTransposeGradientStateWeights(sourceGradient, sourceImage, gradientState, weights)
}

func (cc _CNNConvolutionTransposeGradientNodeClass) Alloc() CNNConvolutionTransposeGradientNode {
	rv := objc.Call[CNNConvolutionTransposeGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNConvolutionTransposeGradientNode_Alloc() CNNConvolutionTransposeGradientNode {
	return CNNConvolutionTransposeGradientNodeClass.Alloc()
}

func (cc _CNNConvolutionTransposeGradientNodeClass) New() CNNConvolutionTransposeGradientNode {
	rv := objc.Call[CNNConvolutionTransposeGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNConvolutionTransposeGradientNode() CNNConvolutionTransposeGradientNode {
	return CNNConvolutionTransposeGradientNodeClass.New()
}

func (c_ CNNConvolutionTransposeGradientNode) Init() CNNConvolutionTransposeGradientNode {
	rv := objc.Call[CNNConvolutionTransposeGradientNode](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNConvolutionTransposeGradientNode) InitWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState ICNNConvolutionGradientStateNode, weights PCNNConvolutionDataSource) CNNConvolutionTransposeGradientNode {
	po3 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNConvolutionTransposeGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:convolutionGradientState:weights:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), po3)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientnode/2947999-initwithsourcegradient?language=objc
func NewCNNConvolutionTransposeGradientNodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState ICNNConvolutionGradientStateNode, weights PCNNConvolutionDataSource) CNNConvolutionTransposeGradientNode {
	instance := CNNConvolutionTransposeGradientNodeClass.Alloc().InitWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient, sourceImage, gradientState, weights)
	instance.Autorelease()
	return instance
}

func (cc _CNNConvolutionTransposeGradientNodeClass) NodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState ICNNConvolutionGradientStateNode, weights PCNNConvolutionDataSource) CNNConvolutionTransposeGradientNode {
	po3 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNConvolutionTransposeGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:convolutionGradientState:weights:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), po3)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientnode/2947984-nodewithsourcegradient?language=objc
func CNNConvolutionTransposeGradientNode_NodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState ICNNConvolutionGradientStateNode, weights PCNNConvolutionDataSource) CNNConvolutionTransposeGradientNode {
	return CNNConvolutionTransposeGradientNodeClass.NodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient, sourceImage, gradientState, weights)
}
