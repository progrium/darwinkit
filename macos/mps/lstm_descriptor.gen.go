// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [LSTMDescriptor] class.
var LSTMDescriptorClass = _LSTMDescriptorClass{objc.GetClass("MPSLSTMDescriptor")}

type _LSTMDescriptorClass struct {
	objc.Class
}

// An interface definition for the [LSTMDescriptor] class.
type ILSTMDescriptor interface {
	IRNNDescriptor
	InputGateInputWeights() CNNConvolutionDataSourceWrapper
	SetInputGateInputWeights(value PCNNConvolutionDataSource)
	SetInputGateInputWeightsObject(valueObject objc.IObject)
	CellToOutputNeuronParamA() float64
	SetCellToOutputNeuronParamA(value float64)
	ForgetGateInputWeights() CNNConvolutionDataSourceWrapper
	SetForgetGateInputWeights(value PCNNConvolutionDataSource)
	SetForgetGateInputWeightsObject(valueObject objc.IObject)
	MemoryWeightsAreDiagonal() bool
	SetMemoryWeightsAreDiagonal(value bool)
	CellGateMemoryWeights() CNNConvolutionDataSourceWrapper
	SetCellGateMemoryWeights(value PCNNConvolutionDataSource)
	SetCellGateMemoryWeightsObject(valueObject objc.IObject)
	ForgetGateRecurrentWeights() CNNConvolutionDataSourceWrapper
	SetForgetGateRecurrentWeights(value PCNNConvolutionDataSource)
	SetForgetGateRecurrentWeightsObject(valueObject objc.IObject)
	ForgetGateMemoryWeights() CNNConvolutionDataSourceWrapper
	SetForgetGateMemoryWeights(value PCNNConvolutionDataSource)
	SetForgetGateMemoryWeightsObject(valueObject objc.IObject)
	CellToOutputNeuronType() CNNNeuronType
	SetCellToOutputNeuronType(value CNNNeuronType)
	CellGateInputWeights() CNNConvolutionDataSourceWrapper
	SetCellGateInputWeights(value PCNNConvolutionDataSource)
	SetCellGateInputWeightsObject(valueObject objc.IObject)
	OutputGateRecurrentWeights() CNNConvolutionDataSourceWrapper
	SetOutputGateRecurrentWeights(value PCNNConvolutionDataSource)
	SetOutputGateRecurrentWeightsObject(valueObject objc.IObject)
	InputGateMemoryWeights() CNNConvolutionDataSourceWrapper
	SetInputGateMemoryWeights(value PCNNConvolutionDataSource)
	SetInputGateMemoryWeightsObject(valueObject objc.IObject)
	InputGateRecurrentWeights() CNNConvolutionDataSourceWrapper
	SetInputGateRecurrentWeights(value PCNNConvolutionDataSource)
	SetInputGateRecurrentWeightsObject(valueObject objc.IObject)
	CellToOutputNeuronParamC() float64
	SetCellToOutputNeuronParamC(value float64)
	CellGateRecurrentWeights() CNNConvolutionDataSourceWrapper
	SetCellGateRecurrentWeights(value PCNNConvolutionDataSource)
	SetCellGateRecurrentWeightsObject(valueObject objc.IObject)
	OutputGateMemoryWeights() CNNConvolutionDataSourceWrapper
	SetOutputGateMemoryWeights(value PCNNConvolutionDataSource)
	SetOutputGateMemoryWeightsObject(valueObject objc.IObject)
	OutputGateInputWeights() CNNConvolutionDataSourceWrapper
	SetOutputGateInputWeights(value PCNNConvolutionDataSource)
	SetOutputGateInputWeightsObject(valueObject objc.IObject)
	CellToOutputNeuronParamB() float64
	SetCellToOutputNeuronParamB(value float64)
}

// A description of a long short-term memory block or layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor?language=objc
type LSTMDescriptor struct {
	RNNDescriptor
}

func LSTMDescriptorFrom(ptr unsafe.Pointer) LSTMDescriptor {
	return LSTMDescriptor{
		RNNDescriptor: RNNDescriptorFrom(ptr),
	}
}

func (lc _LSTMDescriptorClass) CreateLSTMDescriptorWithInputFeatureChannelsOutputFeatureChannels(inputFeatureChannels uint, outputFeatureChannels uint) LSTMDescriptor {
	rv := objc.Call[LSTMDescriptor](lc, objc.Sel("createLSTMDescriptorWithInputFeatureChannels:outputFeatureChannels:"), inputFeatureChannels, outputFeatureChannels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865681-createlstmdescriptorwithinputfea?language=objc
func LSTMDescriptor_CreateLSTMDescriptorWithInputFeatureChannelsOutputFeatureChannels(inputFeatureChannels uint, outputFeatureChannels uint) LSTMDescriptor {
	return LSTMDescriptorClass.CreateLSTMDescriptorWithInputFeatureChannelsOutputFeatureChannels(inputFeatureChannels, outputFeatureChannels)
}

func (lc _LSTMDescriptorClass) Alloc() LSTMDescriptor {
	rv := objc.Call[LSTMDescriptor](lc, objc.Sel("alloc"))
	return rv
}

func LSTMDescriptor_Alloc() LSTMDescriptor {
	return LSTMDescriptorClass.Alloc()
}

func (lc _LSTMDescriptorClass) New() LSTMDescriptor {
	rv := objc.Call[LSTMDescriptor](lc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewLSTMDescriptor() LSTMDescriptor {
	return LSTMDescriptorClass.New()
}

func (l_ LSTMDescriptor) Init() LSTMDescriptor {
	rv := objc.Call[LSTMDescriptor](l_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865684-inputgateinputweights?language=objc
func (l_ LSTMDescriptor) InputGateInputWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](l_, objc.Sel("inputGateInputWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865684-inputgateinputweights?language=objc
func (l_ LSTMDescriptor) SetInputGateInputWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](l_, objc.Sel("setInputGateInputWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865684-inputgateinputweights?language=objc
func (l_ LSTMDescriptor) SetInputGateInputWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setInputGateInputWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865744-celltooutputneuronparama?language=objc
func (l_ LSTMDescriptor) CellToOutputNeuronParamA() float64 {
	rv := objc.Call[float64](l_, objc.Sel("cellToOutputNeuronParamA"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865744-celltooutputneuronparama?language=objc
func (l_ LSTMDescriptor) SetCellToOutputNeuronParamA(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setCellToOutputNeuronParamA:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865734-forgetgateinputweights?language=objc
func (l_ LSTMDescriptor) ForgetGateInputWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](l_, objc.Sel("forgetGateInputWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865734-forgetgateinputweights?language=objc
func (l_ LSTMDescriptor) SetForgetGateInputWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](l_, objc.Sel("setForgetGateInputWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865734-forgetgateinputweights?language=objc
func (l_ LSTMDescriptor) SetForgetGateInputWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setForgetGateInputWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865712-memoryweightsarediagonal?language=objc
func (l_ LSTMDescriptor) MemoryWeightsAreDiagonal() bool {
	rv := objc.Call[bool](l_, objc.Sel("memoryWeightsAreDiagonal"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865712-memoryweightsarediagonal?language=objc
func (l_ LSTMDescriptor) SetMemoryWeightsAreDiagonal(value bool) {
	objc.Call[objc.Void](l_, objc.Sel("setMemoryWeightsAreDiagonal:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865683-cellgatememoryweights?language=objc
func (l_ LSTMDescriptor) CellGateMemoryWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](l_, objc.Sel("cellGateMemoryWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865683-cellgatememoryweights?language=objc
func (l_ LSTMDescriptor) SetCellGateMemoryWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](l_, objc.Sel("setCellGateMemoryWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865683-cellgatememoryweights?language=objc
func (l_ LSTMDescriptor) SetCellGateMemoryWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setCellGateMemoryWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865735-forgetgaterecurrentweights?language=objc
func (l_ LSTMDescriptor) ForgetGateRecurrentWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](l_, objc.Sel("forgetGateRecurrentWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865735-forgetgaterecurrentweights?language=objc
func (l_ LSTMDescriptor) SetForgetGateRecurrentWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](l_, objc.Sel("setForgetGateRecurrentWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865735-forgetgaterecurrentweights?language=objc
func (l_ LSTMDescriptor) SetForgetGateRecurrentWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setForgetGateRecurrentWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865689-forgetgatememoryweights?language=objc
func (l_ LSTMDescriptor) ForgetGateMemoryWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](l_, objc.Sel("forgetGateMemoryWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865689-forgetgatememoryweights?language=objc
func (l_ LSTMDescriptor) SetForgetGateMemoryWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](l_, objc.Sel("setForgetGateMemoryWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865689-forgetgatememoryweights?language=objc
func (l_ LSTMDescriptor) SetForgetGateMemoryWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setForgetGateMemoryWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865736-celltooutputneurontype?language=objc
func (l_ LSTMDescriptor) CellToOutputNeuronType() CNNNeuronType {
	rv := objc.Call[CNNNeuronType](l_, objc.Sel("cellToOutputNeuronType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865736-celltooutputneurontype?language=objc
func (l_ LSTMDescriptor) SetCellToOutputNeuronType(value CNNNeuronType) {
	objc.Call[objc.Void](l_, objc.Sel("setCellToOutputNeuronType:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865741-cellgateinputweights?language=objc
func (l_ LSTMDescriptor) CellGateInputWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](l_, objc.Sel("cellGateInputWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865741-cellgateinputweights?language=objc
func (l_ LSTMDescriptor) SetCellGateInputWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](l_, objc.Sel("setCellGateInputWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865741-cellgateinputweights?language=objc
func (l_ LSTMDescriptor) SetCellGateInputWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setCellGateInputWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865750-outputgaterecurrentweights?language=objc
func (l_ LSTMDescriptor) OutputGateRecurrentWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](l_, objc.Sel("outputGateRecurrentWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865750-outputgaterecurrentweights?language=objc
func (l_ LSTMDescriptor) SetOutputGateRecurrentWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](l_, objc.Sel("setOutputGateRecurrentWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865750-outputgaterecurrentweights?language=objc
func (l_ LSTMDescriptor) SetOutputGateRecurrentWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setOutputGateRecurrentWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865731-inputgatememoryweights?language=objc
func (l_ LSTMDescriptor) InputGateMemoryWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](l_, objc.Sel("inputGateMemoryWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865731-inputgatememoryweights?language=objc
func (l_ LSTMDescriptor) SetInputGateMemoryWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](l_, objc.Sel("setInputGateMemoryWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865731-inputgatememoryweights?language=objc
func (l_ LSTMDescriptor) SetInputGateMemoryWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setInputGateMemoryWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865747-inputgaterecurrentweights?language=objc
func (l_ LSTMDescriptor) InputGateRecurrentWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](l_, objc.Sel("inputGateRecurrentWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865747-inputgaterecurrentweights?language=objc
func (l_ LSTMDescriptor) SetInputGateRecurrentWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](l_, objc.Sel("setInputGateRecurrentWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865747-inputgaterecurrentweights?language=objc
func (l_ LSTMDescriptor) SetInputGateRecurrentWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setInputGateRecurrentWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2935551-celltooutputneuronparamc?language=objc
func (l_ LSTMDescriptor) CellToOutputNeuronParamC() float64 {
	rv := objc.Call[float64](l_, objc.Sel("cellToOutputNeuronParamC"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2935551-celltooutputneuronparamc?language=objc
func (l_ LSTMDescriptor) SetCellToOutputNeuronParamC(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setCellToOutputNeuronParamC:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865679-cellgaterecurrentweights?language=objc
func (l_ LSTMDescriptor) CellGateRecurrentWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](l_, objc.Sel("cellGateRecurrentWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865679-cellgaterecurrentweights?language=objc
func (l_ LSTMDescriptor) SetCellGateRecurrentWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](l_, objc.Sel("setCellGateRecurrentWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865679-cellgaterecurrentweights?language=objc
func (l_ LSTMDescriptor) SetCellGateRecurrentWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setCellGateRecurrentWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865688-outputgatememoryweights?language=objc
func (l_ LSTMDescriptor) OutputGateMemoryWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](l_, objc.Sel("outputGateMemoryWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865688-outputgatememoryweights?language=objc
func (l_ LSTMDescriptor) SetOutputGateMemoryWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](l_, objc.Sel("setOutputGateMemoryWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865688-outputgatememoryweights?language=objc
func (l_ LSTMDescriptor) SetOutputGateMemoryWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setOutputGateMemoryWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865701-outputgateinputweights?language=objc
func (l_ LSTMDescriptor) OutputGateInputWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](l_, objc.Sel("outputGateInputWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865701-outputgateinputweights?language=objc
func (l_ LSTMDescriptor) SetOutputGateInputWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](l_, objc.Sel("setOutputGateInputWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865701-outputgateinputweights?language=objc
func (l_ LSTMDescriptor) SetOutputGateInputWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](l_, objc.Sel("setOutputGateInputWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865694-celltooutputneuronparamb?language=objc
func (l_ LSTMDescriptor) CellToOutputNeuronParamB() float64 {
	rv := objc.Call[float64](l_, objc.Sel("cellToOutputNeuronParamB"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpslstmdescriptor/2865694-celltooutputneuronparamb?language=objc
func (l_ LSTMDescriptor) SetCellToOutputNeuronParamB(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setCellToOutputNeuronParamB:"), value)
}
