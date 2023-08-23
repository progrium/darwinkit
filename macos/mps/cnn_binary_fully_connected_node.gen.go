// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNBinaryFullyConnectedNode] class.
var CNNBinaryFullyConnectedNodeClass = _CNNBinaryFullyConnectedNodeClass{objc.GetClass("MPSCNNBinaryFullyConnectedNode")}

type _CNNBinaryFullyConnectedNodeClass struct {
	objc.Class
}

// An interface definition for the [CNNBinaryFullyConnectedNode] class.
type ICNNBinaryFullyConnectedNode interface {
	ICNNBinaryConvolutionNode
}

// A representation of a fully connected convolution layer with binary weights and optionally binarized input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryfullyconnectednode?language=objc
type CNNBinaryFullyConnectedNode struct {
	CNNBinaryConvolutionNode
}

func CNNBinaryFullyConnectedNodeFrom(ptr unsafe.Pointer) CNNBinaryFullyConnectedNode {
	return CNNBinaryFullyConnectedNode{
		CNNBinaryConvolutionNode: CNNBinaryConvolutionNodeFrom(ptr),
	}
}

func (cc _CNNBinaryFullyConnectedNodeClass) NodeWithSourceWeightsScaleValueTypeFlags(sourceNode INNImageNode, weights PCNNConvolutionDataSource, scaleValue float64, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryFullyConnectedNode {
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNBinaryFullyConnectedNode](cc, objc.Sel("nodeWithSource:weights:scaleValue:type:flags:"), objc.Ptr(sourceNode), po1, scaleValue, type_, flags)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryfullyconnectednode/2866469-nodewithsource?language=objc
func CNNBinaryFullyConnectedNode_NodeWithSourceWeightsScaleValueTypeFlags(sourceNode INNImageNode, weights PCNNConvolutionDataSource, scaleValue float64, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryFullyConnectedNode {
	return CNNBinaryFullyConnectedNodeClass.NodeWithSourceWeightsScaleValueTypeFlags(sourceNode, weights, scaleValue, type_, flags)
}

func (c_ CNNBinaryFullyConnectedNode) InitWithSourceWeightsOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags(sourceNode INNImageNode, weights PCNNConvolutionDataSource, outputBiasTerms *float64, outputScaleTerms *float64, inputBiasTerms *float64, inputScaleTerms *float64, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryFullyConnectedNode {
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNBinaryFullyConnectedNode](c_, objc.Sel("initWithSource:weights:outputBiasTerms:outputScaleTerms:inputBiasTerms:inputScaleTerms:type:flags:"), objc.Ptr(sourceNode), po1, outputBiasTerms, outputScaleTerms, inputBiasTerms, inputScaleTerms, type_, flags)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryfullyconnectednode/2942637-initwithsource?language=objc
func NewCNNBinaryFullyConnectedNodeWithSourceWeightsOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags(sourceNode INNImageNode, weights PCNNConvolutionDataSource, outputBiasTerms *float64, outputScaleTerms *float64, inputBiasTerms *float64, inputScaleTerms *float64, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryFullyConnectedNode {
	instance := CNNBinaryFullyConnectedNodeClass.Alloc().InitWithSourceWeightsOutputBiasTermsOutputScaleTermsInputBiasTermsInputScaleTermsTypeFlags(sourceNode, weights, outputBiasTerms, outputScaleTerms, inputBiasTerms, inputScaleTerms, type_, flags)
	instance.Autorelease()
	return instance
}

func (cc _CNNBinaryFullyConnectedNodeClass) Alloc() CNNBinaryFullyConnectedNode {
	rv := objc.Call[CNNBinaryFullyConnectedNode](cc, objc.Sel("alloc"))
	return rv
}

func CNNBinaryFullyConnectedNode_Alloc() CNNBinaryFullyConnectedNode {
	return CNNBinaryFullyConnectedNodeClass.Alloc()
}

func (cc _CNNBinaryFullyConnectedNodeClass) New() CNNBinaryFullyConnectedNode {
	rv := objc.Call[CNNBinaryFullyConnectedNode](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNBinaryFullyConnectedNode() CNNBinaryFullyConnectedNode {
	return CNNBinaryFullyConnectedNodeClass.New()
}

func (c_ CNNBinaryFullyConnectedNode) Init() CNNBinaryFullyConnectedNode {
	rv := objc.Call[CNNBinaryFullyConnectedNode](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNBinaryFullyConnectedNode) InitWithSourceWeightsScaleValueTypeFlags(sourceNode INNImageNode, weights PCNNConvolutionDataSource, scaleValue float64, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryFullyConnectedNode {
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNBinaryFullyConnectedNode](c_, objc.Sel("initWithSource:weights:scaleValue:type:flags:"), objc.Ptr(sourceNode), po1, scaleValue, type_, flags)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolutionnode/2866509-initwithsource?language=objc
func NewCNNBinaryFullyConnectedNodeWithSourceWeightsScaleValueTypeFlags(sourceNode INNImageNode, weights PCNNConvolutionDataSource, scaleValue float64, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryFullyConnectedNode {
	instance := CNNBinaryFullyConnectedNodeClass.Alloc().InitWithSourceWeightsScaleValueTypeFlags(sourceNode, weights, scaleValue, type_, flags)
	instance.Autorelease()
	return instance
}

func (cc _CNNBinaryFullyConnectedNodeClass) NodeWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNBinaryFullyConnectedNode {
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNBinaryFullyConnectedNode](cc, objc.Sel("nodeWithSource:weights:"), objc.Ptr(sourceNode), po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2866436-nodewithsource?language=objc
func CNNBinaryFullyConnectedNode_NodeWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNBinaryFullyConnectedNode {
	return CNNBinaryFullyConnectedNodeClass.NodeWithSourceWeights(sourceNode, weights)
}

func (c_ CNNBinaryFullyConnectedNode) InitWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNBinaryFullyConnectedNode {
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNBinaryFullyConnectedNode](c_, objc.Sel("initWithSource:weights:"), objc.Ptr(sourceNode), po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionnode/2866470-initwithsource?language=objc
func NewCNNBinaryFullyConnectedNodeWithSourceWeights(sourceNode INNImageNode, weights PCNNConvolutionDataSource) CNNBinaryFullyConnectedNode {
	instance := CNNBinaryFullyConnectedNodeClass.Alloc().InitWithSourceWeights(sourceNode, weights)
	instance.Autorelease()
	return instance
}
