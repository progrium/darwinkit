// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixFullyConnected] class.
var MatrixFullyConnectedClass = _MatrixFullyConnectedClass{objc.GetClass("MPSMatrixFullyConnected")}

type _MatrixFullyConnectedClass struct {
	objc.Class
}

// An interface definition for the [MatrixFullyConnected] class.
type IMatrixFullyConnected interface {
	IMatrixBinaryKernel
	NeuronParameterA() float64
	NeuronType() CNNNeuronType
	SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float64, parameterB float64, parameterC float64)
	NeuronParameterB() float64
	EncodeToCommandBufferInputMatrixWeightMatrixBiasVectorResultMatrix(commandBuffer metal.PCommandBuffer, inputMatrix IMatrix, weightMatrix IMatrix, biasVector IVector, resultMatrix IMatrix)
	EncodeToCommandBufferObjectInputMatrixWeightMatrixBiasVectorResultMatrix(commandBufferObject objc.IObject, inputMatrix IMatrix, weightMatrix IMatrix, biasVector IVector, resultMatrix IMatrix)
	NeuronParameterC() float64
	SourceOutputFeatureChannels() uint
	SetSourceOutputFeatureChannels(value uint)
	Alpha() float64
	SetAlpha(value float64)
	SourceNumberOfFeatureVectors() uint
	SetSourceNumberOfFeatureVectors(value uint)
	SourceInputFeatureChannels() uint
	SetSourceInputFeatureChannels(value uint)
}

// A kernel for applying a fully connected neural network layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected?language=objc
type MatrixFullyConnected struct {
	MatrixBinaryKernel
}

func MatrixFullyConnectedFrom(ptr unsafe.Pointer) MatrixFullyConnected {
	return MatrixFullyConnected{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}

func (m_ MatrixFullyConnected) InitWithDevice(device metal.PDevice) MatrixFullyConnected {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixFullyConnected](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935584-initwithdevice?language=objc
func NewMatrixFullyConnectedWithDevice(device metal.PDevice) MatrixFullyConnected {
	instance := MatrixFullyConnectedClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixFullyConnected) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixFullyConnected {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixFullyConnected](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935595-copywithzone?language=objc
func MatrixFullyConnected_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixFullyConnected {
	instance := MatrixFullyConnectedClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (mc _MatrixFullyConnectedClass) Alloc() MatrixFullyConnected {
	rv := objc.Call[MatrixFullyConnected](mc, objc.Sel("alloc"))
	return rv
}

func MatrixFullyConnected_Alloc() MatrixFullyConnected {
	return MatrixFullyConnectedClass.Alloc()
}

func (mc _MatrixFullyConnectedClass) New() MatrixFullyConnected {
	rv := objc.Call[MatrixFullyConnected](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixFullyConnected() MatrixFullyConnected {
	return MatrixFullyConnectedClass.New()
}

func (m_ MatrixFullyConnected) Init() MatrixFullyConnected {
	rv := objc.Call[MatrixFullyConnected](m_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935602-neuronparametera?language=objc
func (m_ MatrixFullyConnected) NeuronParameterA() float64 {
	rv := objc.Call[float64](m_, objc.Sel("neuronParameterA"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935588-neurontype?language=objc
func (m_ MatrixFullyConnected) NeuronType() CNNNeuronType {
	rv := objc.Call[CNNNeuronType](m_, objc.Sel("neuronType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935593-setneurontype?language=objc
func (m_ MatrixFullyConnected) SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float64, parameterB float64, parameterC float64) {
	objc.Call[objc.Void](m_, objc.Sel("setNeuronType:parameterA:parameterB:parameterC:"), neuronType, parameterA, parameterB, parameterC)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935591-neuronparameterb?language=objc
func (m_ MatrixFullyConnected) NeuronParameterB() float64 {
	rv := objc.Call[float64](m_, objc.Sel("neuronParameterB"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935596-encodetocommandbuffer?language=objc
func (m_ MatrixFullyConnected) EncodeToCommandBufferInputMatrixWeightMatrixBiasVectorResultMatrix(commandBuffer metal.PCommandBuffer, inputMatrix IMatrix, weightMatrix IMatrix, biasVector IVector, resultMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:inputMatrix:weightMatrix:biasVector:resultMatrix:"), po0, objc.Ptr(inputMatrix), objc.Ptr(weightMatrix), objc.Ptr(biasVector), objc.Ptr(resultMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935596-encodetocommandbuffer?language=objc
func (m_ MatrixFullyConnected) EncodeToCommandBufferObjectInputMatrixWeightMatrixBiasVectorResultMatrix(commandBufferObject objc.IObject, inputMatrix IMatrix, weightMatrix IMatrix, biasVector IVector, resultMatrix IMatrix) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:inputMatrix:weightMatrix:biasVector:resultMatrix:"), objc.Ptr(commandBufferObject), objc.Ptr(inputMatrix), objc.Ptr(weightMatrix), objc.Ptr(biasVector), objc.Ptr(resultMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935594-neuronparameterc?language=objc
func (m_ MatrixFullyConnected) NeuronParameterC() float64 {
	rv := objc.Call[float64](m_, objc.Sel("neuronParameterC"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935592-sourceoutputfeaturechannels?language=objc
func (m_ MatrixFullyConnected) SourceOutputFeatureChannels() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceOutputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935592-sourceoutputfeaturechannels?language=objc
func (m_ MatrixFullyConnected) SetSourceOutputFeatureChannels(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceOutputFeatureChannels:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935608-alpha?language=objc
func (m_ MatrixFullyConnected) Alpha() float64 {
	rv := objc.Call[float64](m_, objc.Sel("alpha"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935608-alpha?language=objc
func (m_ MatrixFullyConnected) SetAlpha(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setAlpha:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935609-sourcenumberoffeaturevectors?language=objc
func (m_ MatrixFullyConnected) SourceNumberOfFeatureVectors() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935609-sourcenumberoffeaturevectors?language=objc
func (m_ MatrixFullyConnected) SetSourceNumberOfFeatureVectors(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935597-sourceinputfeaturechannels?language=objc
func (m_ MatrixFullyConnected) SourceInputFeatureChannels() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceInputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnected/2935597-sourceinputfeaturechannels?language=objc
func (m_ MatrixFullyConnected) SetSourceInputFeatureChannels(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceInputFeatureChannels:"), value)
}
