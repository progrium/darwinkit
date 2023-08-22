// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixNeuron] class.
var MatrixNeuronClass = _MatrixNeuronClass{objc.GetClass("MPSMatrixNeuron")}

type _MatrixNeuronClass struct {
	objc.Class
}

// An interface definition for the [MatrixNeuron] class.
type IMatrixNeuron interface {
	IMatrixUnaryKernel
	NeuronParameterA() float64
	NeuronType() CNNNeuronType
	SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float64, parameterB float64, parameterC float64)
	NeuronParameterB() float64
	EncodeToCommandBufferInputMatrixBiasVectorResultMatrix(commandBuffer metal.PCommandBuffer, inputMatrix IMatrix, biasVector IVector, resultMatrix IMatrix)
	EncodeToCommandBufferObjectInputMatrixBiasVectorResultMatrix(commandBufferObject objc.IObject, inputMatrix IMatrix, biasVector IVector, resultMatrix IMatrix)
	SetNeuronToPReLUWithParametersA(A []byte)
	NeuronParameterC() float64
	Alpha() float64
	SetAlpha(value float64)
	SourceNumberOfFeatureVectors() uint
	SetSourceNumberOfFeatureVectors(value uint)
	SourceInputFeatureChannels() uint
	SetSourceInputFeatureChannels(value uint)
}

// A neuron activation kernel that operates on matrices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron?language=objc
type MatrixNeuron struct {
	MatrixUnaryKernel
}

func MatrixNeuronFrom(ptr unsafe.Pointer) MatrixNeuron {
	return MatrixNeuron{
		MatrixUnaryKernel: MatrixUnaryKernelFrom(ptr),
	}
}

func (m_ MatrixNeuron) InitWithDevice(device metal.PDevice) MatrixNeuron {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixNeuron](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935603-initwithdevice?language=objc
func NewMatrixNeuronWithDevice(device metal.PDevice) MatrixNeuron {
	instance := MatrixNeuronClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixNeuron) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixNeuron {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixNeuron](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935604-copywithzone?language=objc
func MatrixNeuron_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixNeuron {
	instance := MatrixNeuronClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (mc _MatrixNeuronClass) Alloc() MatrixNeuron {
	rv := objc.Call[MatrixNeuron](mc, objc.Sel("alloc"))
	return rv
}

func MatrixNeuron_Alloc() MatrixNeuron {
	return MatrixNeuronClass.Alloc()
}

func (mc _MatrixNeuronClass) New() MatrixNeuron {
	rv := objc.Call[MatrixNeuron](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixNeuron() MatrixNeuron {
	return MatrixNeuronClass.New()
}

func (m_ MatrixNeuron) Init() MatrixNeuron {
	rv := objc.Call[MatrixNeuron](m_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935583-neuronparametera?language=objc
func (m_ MatrixNeuron) NeuronParameterA() float64 {
	rv := objc.Call[float64](m_, objc.Sel("neuronParameterA"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935587-neurontype?language=objc
func (m_ MatrixNeuron) NeuronType() CNNNeuronType {
	rv := objc.Call[CNNNeuronType](m_, objc.Sel("neuronType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935590-setneurontype?language=objc
func (m_ MatrixNeuron) SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float64, parameterB float64, parameterC float64) {
	objc.Call[objc.Void](m_, objc.Sel("setNeuronType:parameterA:parameterB:parameterC:"), neuronType, parameterA, parameterB, parameterC)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935585-neuronparameterb?language=objc
func (m_ MatrixNeuron) NeuronParameterB() float64 {
	rv := objc.Call[float64](m_, objc.Sel("neuronParameterB"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935606-encodetocommandbuffer?language=objc
func (m_ MatrixNeuron) EncodeToCommandBufferInputMatrixBiasVectorResultMatrix(commandBuffer metal.PCommandBuffer, inputMatrix IMatrix, biasVector IVector, resultMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:inputMatrix:biasVector:resultMatrix:"), po0, objc.Ptr(inputMatrix), objc.Ptr(biasVector), objc.Ptr(resultMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935606-encodetocommandbuffer?language=objc
func (m_ MatrixNeuron) EncodeToCommandBufferObjectInputMatrixBiasVectorResultMatrix(commandBufferObject objc.IObject, inputMatrix IMatrix, biasVector IVector, resultMatrix IMatrix) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:inputMatrix:biasVector:resultMatrix:"), objc.Ptr(commandBufferObject), objc.Ptr(inputMatrix), objc.Ptr(biasVector), objc.Ptr(resultMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935610-setneurontopreluwithparametersa?language=objc
func (m_ MatrixNeuron) SetNeuronToPReLUWithParametersA(A []byte) {
	objc.Call[objc.Void](m_, objc.Sel("setNeuronToPReLUWithParametersA:"), A)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935598-neuronparameterc?language=objc
func (m_ MatrixNeuron) NeuronParameterC() float64 {
	rv := objc.Call[float64](m_, objc.Sel("neuronParameterC"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935605-alpha?language=objc
func (m_ MatrixNeuron) Alpha() float64 {
	rv := objc.Call[float64](m_, objc.Sel("alpha"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935605-alpha?language=objc
func (m_ MatrixNeuron) SetAlpha(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setAlpha:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935607-sourcenumberoffeaturevectors?language=objc
func (m_ MatrixNeuron) SourceNumberOfFeatureVectors() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935607-sourcenumberoffeaturevectors?language=objc
func (m_ MatrixNeuron) SetSourceNumberOfFeatureVectors(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935599-sourceinputfeaturechannels?language=objc
func (m_ MatrixNeuron) SourceInputFeatureChannels() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceInputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixneuron/2935599-sourceinputfeaturechannels?language=objc
func (m_ MatrixNeuron) SetSourceInputFeatureChannels(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceInputFeatureChannels:"), value)
}
