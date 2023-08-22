// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNFullyConnectedGradientNode] class.
var CNNFullyConnectedGradientNodeClass = _CNNFullyConnectedGradientNodeClass{objc.GetClass("MPSCNNFullyConnectedGradientNode")}

type _CNNFullyConnectedGradientNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNFullyConnectedGradientNode] class.
type ICNNFullyConnectedGradientNode interface {
	ICNNConvolutionGradientNode
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnectedgradientnode?language=objc
type CNNFullyConnectedGradientNode struct {
	CNNConvolutionGradientNode
}

func CNNFullyConnectedGradientNodeFrom(ptr unsafe.Pointer) CNNFullyConnectedGradientNode {
	return CNNFullyConnectedGradientNode{
		CNNConvolutionGradientNode: CNNConvolutionGradientNodeFrom(ptr),
	}
}

func (c_ CNNFullyConnectedGradientNode) InitWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState ICNNConvolutionGradientStateNode, weights PCNNConvolutionDataSource) CNNFullyConnectedGradientNode {
	po3 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNFullyConnectedGradientNode](c_, objc.Sel("initWithSourceGradient:sourceImage:convolutionGradientState:weights:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), po3)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnectedgradientnode/3152566-initwithsourcegradient?language=objc
func NewCNNFullyConnectedGradientNodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState ICNNConvolutionGradientStateNode, weights PCNNConvolutionDataSource) CNNFullyConnectedGradientNode {
	instance := CNNFullyConnectedGradientNodeClass.Alloc().InitWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient, sourceImage, gradientState, weights)
	instance.Autorelease()
	return instance
}

func (cc _CNNFullyConnectedGradientNodeClass) NodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState ICNNConvolutionGradientStateNode, weights PCNNConvolutionDataSource) CNNFullyConnectedGradientNode {
	po3 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNFullyConnectedGradientNode](cc, objc.Sel("nodeWithSourceGradient:sourceImage:convolutionGradientState:weights:"), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(gradientState), po3)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnectedgradientnode/3152567-nodewithsourcegradient?language=objc
func CNNFullyConnectedGradientNode_NodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient INNImageNode, sourceImage INNImageNode, gradientState ICNNConvolutionGradientStateNode, weights PCNNConvolutionDataSource) CNNFullyConnectedGradientNode {
	return CNNFullyConnectedGradientNodeClass.NodeWithSourceGradientSourceImageConvolutionGradientStateWeights(sourceGradient, sourceImage, gradientState, weights)
}

func (cc _CNNFullyConnectedGradientNodeClass) Alloc() CNNFullyConnectedGradientNode {
	rv := objc.Call[CNNFullyConnectedGradientNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNFullyConnectedGradientNode_Alloc() CNNFullyConnectedGradientNode {
	return CNNFullyConnectedGradientNodeClass.Alloc()
}

func (cc _CNNFullyConnectedGradientNodeClass) New() CNNFullyConnectedGradientNode {
	rv := objc.Call[CNNFullyConnectedGradientNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNFullyConnectedGradientNode() CNNFullyConnectedGradientNode {
	return CNNFullyConnectedGradientNodeClass.New()
}

func (c_ CNNFullyConnectedGradientNode) Init() CNNFullyConnectedGradientNode {
	rv := objc.Call[CNNFullyConnectedGradientNode](c_, objc.Sel("init"))
	return rv
}
