// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixBatchNormalization] class.
var MatrixBatchNormalizationClass = _MatrixBatchNormalizationClass{objc.GetClass("MPSMatrixBatchNormalization")}

type _MatrixBatchNormalizationClass struct {
	objc.Class
}

// An interface definition for the [MatrixBatchNormalization] class.
type IMatrixBatchNormalization interface {
	IMatrixUnaryKernel
	NeuronParameterA() float64
	NeuronType() CNNNeuronType
	SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float64, parameterB float64, parameterC float64)
	NeuronParameterB() float64
	EncodeToCommandBufferInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultMatrix(commandBuffer metal.PCommandBuffer, inputMatrix IMatrix, meanVector IVector, varianceVector IVector, gammaVector IVector, betaVector IVector, resultMatrix IMatrix)
	EncodeToCommandBufferObjectInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultMatrix(commandBufferObject objc.IObject, inputMatrix IMatrix, meanVector IVector, varianceVector IVector, gammaVector IVector, betaVector IVector, resultMatrix IMatrix)
	NeuronParameterC() float64
	ComputeStatistics() bool
	SetComputeStatistics(value bool)
	Epsilon() float64
	SetEpsilon(value float64)
	SourceNumberOfFeatureVectors() uint
	SetSourceNumberOfFeatureVectors(value uint)
	SourceInputFeatureChannels() uint
	SetSourceInputFeatureChannels(value uint)
}

// A batch normalization kernel that operates on matrices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization?language=objc
type MatrixBatchNormalization struct {
	MatrixUnaryKernel
}

func MatrixBatchNormalizationFrom(ptr unsafe.Pointer) MatrixBatchNormalization {
	return MatrixBatchNormalization{
		MatrixUnaryKernel: MatrixUnaryKernelFrom(ptr),
	}
}

func (m_ MatrixBatchNormalization) InitWithDevice(device metal.PDevice) MatrixBatchNormalization {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixBatchNormalization](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980735-initwithdevice?language=objc
func NewMatrixBatchNormalizationWithDevice(device metal.PDevice) MatrixBatchNormalization {
	instance := MatrixBatchNormalizationClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixBatchNormalization) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixBatchNormalization {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixBatchNormalization](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980731-copywithzone?language=objc
func MatrixBatchNormalization_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixBatchNormalization {
	instance := MatrixBatchNormalizationClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (mc _MatrixBatchNormalizationClass) Alloc() MatrixBatchNormalization {
	rv := objc.Call[MatrixBatchNormalization](mc, objc.Sel("alloc"))
	return rv
}

func MatrixBatchNormalization_Alloc() MatrixBatchNormalization {
	return MatrixBatchNormalizationClass.Alloc()
}

func (mc _MatrixBatchNormalizationClass) New() MatrixBatchNormalization {
	rv := objc.Call[MatrixBatchNormalization](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixBatchNormalization() MatrixBatchNormalization {
	return MatrixBatchNormalizationClass.New()
}

func (m_ MatrixBatchNormalization) Init() MatrixBatchNormalization {
	rv := objc.Call[MatrixBatchNormalization](m_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980736-neuronparametera?language=objc
func (m_ MatrixBatchNormalization) NeuronParameterA() float64 {
	rv := objc.Call[float64](m_, objc.Sel("neuronParameterA"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980739-neurontype?language=objc
func (m_ MatrixBatchNormalization) NeuronType() CNNNeuronType {
	rv := objc.Call[CNNNeuronType](m_, objc.Sel("neuronType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980740-setneurontype?language=objc
func (m_ MatrixBatchNormalization) SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float64, parameterB float64, parameterC float64) {
	objc.Call[objc.Void](m_, objc.Sel("setNeuronType:parameterA:parameterB:parameterC:"), neuronType, parameterA, parameterB, parameterC)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980737-neuronparameterb?language=objc
func (m_ MatrixBatchNormalization) NeuronParameterB() float64 {
	rv := objc.Call[float64](m_, objc.Sel("neuronParameterB"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980732-encodetocommandbuffer?language=objc
func (m_ MatrixBatchNormalization) EncodeToCommandBufferInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultMatrix(commandBuffer metal.PCommandBuffer, inputMatrix IMatrix, meanVector IVector, varianceVector IVector, gammaVector IVector, betaVector IVector, resultMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:inputMatrix:meanVector:varianceVector:gammaVector:betaVector:resultMatrix:"), po0, objc.Ptr(inputMatrix), objc.Ptr(meanVector), objc.Ptr(varianceVector), objc.Ptr(gammaVector), objc.Ptr(betaVector), objc.Ptr(resultMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980732-encodetocommandbuffer?language=objc
func (m_ MatrixBatchNormalization) EncodeToCommandBufferObjectInputMatrixMeanVectorVarianceVectorGammaVectorBetaVectorResultMatrix(commandBufferObject objc.IObject, inputMatrix IMatrix, meanVector IVector, varianceVector IVector, gammaVector IVector, betaVector IVector, resultMatrix IMatrix) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:inputMatrix:meanVector:varianceVector:gammaVector:betaVector:resultMatrix:"), objc.Ptr(commandBufferObject), objc.Ptr(inputMatrix), objc.Ptr(meanVector), objc.Ptr(varianceVector), objc.Ptr(gammaVector), objc.Ptr(betaVector), objc.Ptr(resultMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980738-neuronparameterc?language=objc
func (m_ MatrixBatchNormalization) NeuronParameterC() float64 {
	rv := objc.Call[float64](m_, objc.Sel("neuronParameterC"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980730-computestatistics?language=objc
func (m_ MatrixBatchNormalization) ComputeStatistics() bool {
	rv := objc.Call[bool](m_, objc.Sel("computeStatistics"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980730-computestatistics?language=objc
func (m_ MatrixBatchNormalization) SetComputeStatistics(value bool) {
	objc.Call[objc.Void](m_, objc.Sel("setComputeStatistics:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980733-epsilon?language=objc
func (m_ MatrixBatchNormalization) Epsilon() float64 {
	rv := objc.Call[float64](m_, objc.Sel("epsilon"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980733-epsilon?language=objc
func (m_ MatrixBatchNormalization) SetEpsilon(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setEpsilon:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980742-sourcenumberoffeaturevectors?language=objc
func (m_ MatrixBatchNormalization) SourceNumberOfFeatureVectors() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980742-sourcenumberoffeaturevectors?language=objc
func (m_ MatrixBatchNormalization) SetSourceNumberOfFeatureVectors(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980741-sourceinputfeaturechannels?language=objc
func (m_ MatrixBatchNormalization) SourceInputFeatureChannels() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceInputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixbatchnormalization/2980741-sourceinputfeaturechannels?language=objc
func (m_ MatrixBatchNormalization) SetSourceInputFeatureChannels(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceInputFeatureChannels:"), value)
}
