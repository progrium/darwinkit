// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [GRUDescriptor] class.
var GRUDescriptorClass = _GRUDescriptorClass{objc.GetClass("MPSGRUDescriptor")}

type _GRUDescriptorClass struct {
	objc.Class
}

// An interface definition for the [GRUDescriptor] class.
type IGRUDescriptor interface {
	IRNNDescriptor
	InputGateInputWeights() CNNConvolutionDataSourceWrapper
	SetInputGateInputWeights(value PCNNConvolutionDataSource)
	SetInputGateInputWeightsObject(valueObject objc.IObject)
	RecurrentGateRecurrentWeights() CNNConvolutionDataSourceWrapper
	SetRecurrentGateRecurrentWeights(value PCNNConvolutionDataSource)
	SetRecurrentGateRecurrentWeightsObject(valueObject objc.IObject)
	OutputGateInputGateWeights() CNNConvolutionDataSourceWrapper
	SetOutputGateInputGateWeights(value PCNNConvolutionDataSource)
	SetOutputGateInputGateWeightsObject(valueObject objc.IObject)
	GatePnormValue() float64
	SetGatePnormValue(value float64)
	OutputGateRecurrentWeights() CNNConvolutionDataSourceWrapper
	SetOutputGateRecurrentWeights(value PCNNConvolutionDataSource)
	SetOutputGateRecurrentWeightsObject(valueObject objc.IObject)
	FlipOutputGates() bool
	SetFlipOutputGates(value bool)
	RecurrentGateInputWeights() CNNConvolutionDataSourceWrapper
	SetRecurrentGateInputWeights(value PCNNConvolutionDataSource)
	SetRecurrentGateInputWeightsObject(valueObject objc.IObject)
	InputGateRecurrentWeights() CNNConvolutionDataSourceWrapper
	SetInputGateRecurrentWeights(value PCNNConvolutionDataSource)
	SetInputGateRecurrentWeightsObject(valueObject objc.IObject)
	OutputGateInputWeights() CNNConvolutionDataSourceWrapper
	SetOutputGateInputWeights(value PCNNConvolutionDataSource)
	SetOutputGateInputWeightsObject(valueObject objc.IObject)
}

// A description of a gated recurrent unit block or layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor?language=objc
type GRUDescriptor struct {
	RNNDescriptor
}

func GRUDescriptorFrom(ptr unsafe.Pointer) GRUDescriptor {
	return GRUDescriptor{
		RNNDescriptor: RNNDescriptorFrom(ptr),
	}
}

func (gc _GRUDescriptorClass) CreateGRUDescriptorWithInputFeatureChannelsOutputFeatureChannels(inputFeatureChannels uint, outputFeatureChannels uint) GRUDescriptor {
	rv := objc.Call[GRUDescriptor](gc, objc.Sel("createGRUDescriptorWithInputFeatureChannels:outputFeatureChannels:"), inputFeatureChannels, outputFeatureChannels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865715-creategrudescriptorwithinputfeat?language=objc
func GRUDescriptor_CreateGRUDescriptorWithInputFeatureChannelsOutputFeatureChannels(inputFeatureChannels uint, outputFeatureChannels uint) GRUDescriptor {
	return GRUDescriptorClass.CreateGRUDescriptorWithInputFeatureChannelsOutputFeatureChannels(inputFeatureChannels, outputFeatureChannels)
}

func (gc _GRUDescriptorClass) Alloc() GRUDescriptor {
	rv := objc.Call[GRUDescriptor](gc, objc.Sel("alloc"))
	return rv
}

func GRUDescriptor_Alloc() GRUDescriptor {
	return GRUDescriptorClass.Alloc()
}

func (gc _GRUDescriptorClass) New() GRUDescriptor {
	rv := objc.Call[GRUDescriptor](gc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewGRUDescriptor() GRUDescriptor {
	return GRUDescriptorClass.New()
}

func (g_ GRUDescriptor) Init() GRUDescriptor {
	rv := objc.Call[GRUDescriptor](g_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865690-inputgateinputweights?language=objc
func (g_ GRUDescriptor) InputGateInputWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](g_, objc.Sel("inputGateInputWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865690-inputgateinputweights?language=objc
func (g_ GRUDescriptor) SetInputGateInputWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](g_, objc.Sel("setInputGateInputWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865690-inputgateinputweights?language=objc
func (g_ GRUDescriptor) SetInputGateInputWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](g_, objc.Sel("setInputGateInputWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865695-recurrentgaterecurrentweights?language=objc
func (g_ GRUDescriptor) RecurrentGateRecurrentWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](g_, objc.Sel("recurrentGateRecurrentWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865695-recurrentgaterecurrentweights?language=objc
func (g_ GRUDescriptor) SetRecurrentGateRecurrentWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](g_, objc.Sel("setRecurrentGateRecurrentWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865695-recurrentgaterecurrentweights?language=objc
func (g_ GRUDescriptor) SetRecurrentGateRecurrentWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](g_, objc.Sel("setRecurrentGateRecurrentWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2878270-outputgateinputgateweights?language=objc
func (g_ GRUDescriptor) OutputGateInputGateWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](g_, objc.Sel("outputGateInputGateWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2878270-outputgateinputgateweights?language=objc
func (g_ GRUDescriptor) SetOutputGateInputGateWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](g_, objc.Sel("setOutputGateInputGateWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2878270-outputgateinputgateweights?language=objc
func (g_ GRUDescriptor) SetOutputGateInputGateWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](g_, objc.Sel("setOutputGateInputGateWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2873332-gatepnormvalue?language=objc
func (g_ GRUDescriptor) GatePnormValue() float64 {
	rv := objc.Call[float64](g_, objc.Sel("gatePnormValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2873332-gatepnormvalue?language=objc
func (g_ GRUDescriptor) SetGatePnormValue(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setGatePnormValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865699-outputgaterecurrentweights?language=objc
func (g_ GRUDescriptor) OutputGateRecurrentWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](g_, objc.Sel("outputGateRecurrentWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865699-outputgaterecurrentweights?language=objc
func (g_ GRUDescriptor) SetOutputGateRecurrentWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](g_, objc.Sel("setOutputGateRecurrentWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865699-outputgaterecurrentweights?language=objc
func (g_ GRUDescriptor) SetOutputGateRecurrentWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](g_, objc.Sel("setOutputGateRecurrentWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2878271-flipoutputgates?language=objc
func (g_ GRUDescriptor) FlipOutputGates() bool {
	rv := objc.Call[bool](g_, objc.Sel("flipOutputGates"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2878271-flipoutputgates?language=objc
func (g_ GRUDescriptor) SetFlipOutputGates(value bool) {
	objc.Call[objc.Void](g_, objc.Sel("setFlipOutputGates:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865719-recurrentgateinputweights?language=objc
func (g_ GRUDescriptor) RecurrentGateInputWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](g_, objc.Sel("recurrentGateInputWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865719-recurrentgateinputweights?language=objc
func (g_ GRUDescriptor) SetRecurrentGateInputWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](g_, objc.Sel("setRecurrentGateInputWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865719-recurrentgateinputweights?language=objc
func (g_ GRUDescriptor) SetRecurrentGateInputWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](g_, objc.Sel("setRecurrentGateInputWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865724-inputgaterecurrentweights?language=objc
func (g_ GRUDescriptor) InputGateRecurrentWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](g_, objc.Sel("inputGateRecurrentWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865724-inputgaterecurrentweights?language=objc
func (g_ GRUDescriptor) SetInputGateRecurrentWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](g_, objc.Sel("setInputGateRecurrentWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865724-inputgaterecurrentweights?language=objc
func (g_ GRUDescriptor) SetInputGateRecurrentWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](g_, objc.Sel("setInputGateRecurrentWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865722-outputgateinputweights?language=objc
func (g_ GRUDescriptor) OutputGateInputWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](g_, objc.Sel("outputGateInputWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865722-outputgateinputweights?language=objc
func (g_ GRUDescriptor) SetOutputGateInputWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](g_, objc.Sel("setOutputGateInputWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsgrudescriptor/2865722-outputgateinputweights?language=objc
func (g_ GRUDescriptor) SetOutputGateInputWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](g_, objc.Sel("setOutputGateInputWeights:"), objc.Ptr(valueObject))
}
