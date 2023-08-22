// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixFullyConnectedGradient] class.
var MatrixFullyConnectedGradientClass = _MatrixFullyConnectedGradientClass{objc.GetClass("MPSMatrixFullyConnectedGradient")}

type _MatrixFullyConnectedGradientClass struct {
	objc.Class
}

// An interface definition for the [MatrixFullyConnectedGradient] class.
type IMatrixFullyConnectedGradient interface {
	IMatrixBinaryKernel
	EncodeGradientForWeightsAndBiasToCommandBufferGradientMatrixInputMatrixResultGradientForWeightMatrixResultGradientForBiasVector(commandBuffer metal.PCommandBuffer, gradientMatrix IMatrix, inputMatrix IMatrix, resultGradientForWeightMatrix IMatrix, resultGradientForBiasVector IVector)
	EncodeGradientForWeightsAndBiasToCommandBufferObjectGradientMatrixInputMatrixResultGradientForWeightMatrixResultGradientForBiasVector(commandBufferObject objc.IObject, gradientMatrix IMatrix, inputMatrix IMatrix, resultGradientForWeightMatrix IMatrix, resultGradientForBiasVector IVector)
	EncodeGradientForDataToCommandBufferGradientMatrixWeightMatrixResultGradientForDataMatrix(commandBuffer metal.PCommandBuffer, gradientMatrix IMatrix, weightMatrix IMatrix, resultGradientForDataMatrix IMatrix)
	EncodeGradientForDataToCommandBufferObjectGradientMatrixWeightMatrixResultGradientForDataMatrix(commandBufferObject objc.IObject, gradientMatrix IMatrix, weightMatrix IMatrix, resultGradientForDataMatrix IMatrix)
	SourceOutputFeatureChannels() uint
	SetSourceOutputFeatureChannels(value uint)
	Alpha() float64
	SetAlpha(value float64)
	SourceNumberOfFeatureVectors() uint
	SetSourceNumberOfFeatureVectors(value uint)
	SourceInputFeatureChannels() uint
	SetSourceInputFeatureChannels(value uint)
}

// A kernel for applying a fully gradient connected neural network layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient?language=objc
type MatrixFullyConnectedGradient struct {
	MatrixBinaryKernel
}

func MatrixFullyConnectedGradientFrom(ptr unsafe.Pointer) MatrixFullyConnectedGradient {
	return MatrixFullyConnectedGradient{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}

func (m_ MatrixFullyConnectedGradient) InitWithDevice(device metal.PDevice) MatrixFullyConnectedGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixFullyConnectedGradient](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966668-initwithdevice?language=objc
func NewMatrixFullyConnectedGradientWithDevice(device metal.PDevice) MatrixFullyConnectedGradient {
	instance := MatrixFullyConnectedGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixFullyConnectedGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixFullyConnectedGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixFullyConnectedGradient](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966664-copywithzone?language=objc
func MatrixFullyConnectedGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixFullyConnectedGradient {
	instance := MatrixFullyConnectedGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (mc _MatrixFullyConnectedGradientClass) Alloc() MatrixFullyConnectedGradient {
	rv := objc.Call[MatrixFullyConnectedGradient](mc, objc.Sel("alloc"))
	return rv
}

func MatrixFullyConnectedGradient_Alloc() MatrixFullyConnectedGradient {
	return MatrixFullyConnectedGradientClass.Alloc()
}

func (mc _MatrixFullyConnectedGradientClass) New() MatrixFullyConnectedGradient {
	rv := objc.Call[MatrixFullyConnectedGradient](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixFullyConnectedGradient() MatrixFullyConnectedGradient {
	return MatrixFullyConnectedGradientClass.New()
}

func (m_ MatrixFullyConnectedGradient) Init() MatrixFullyConnectedGradient {
	rv := objc.Call[MatrixFullyConnectedGradient](m_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966666-encodegradientforweightsandbiast?language=objc
func (m_ MatrixFullyConnectedGradient) EncodeGradientForWeightsAndBiasToCommandBufferGradientMatrixInputMatrixResultGradientForWeightMatrixResultGradientForBiasVector(commandBuffer metal.PCommandBuffer, gradientMatrix IMatrix, inputMatrix IMatrix, resultGradientForWeightMatrix IMatrix, resultGradientForBiasVector IVector) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeGradientForWeightsAndBiasToCommandBuffer:gradientMatrix:inputMatrix:resultGradientForWeightMatrix:resultGradientForBiasVector:"), po0, objc.Ptr(gradientMatrix), objc.Ptr(inputMatrix), objc.Ptr(resultGradientForWeightMatrix), objc.Ptr(resultGradientForBiasVector))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966666-encodegradientforweightsandbiast?language=objc
func (m_ MatrixFullyConnectedGradient) EncodeGradientForWeightsAndBiasToCommandBufferObjectGradientMatrixInputMatrixResultGradientForWeightMatrixResultGradientForBiasVector(commandBufferObject objc.IObject, gradientMatrix IMatrix, inputMatrix IMatrix, resultGradientForWeightMatrix IMatrix, resultGradientForBiasVector IVector) {
	objc.Call[objc.Void](m_, objc.Sel("encodeGradientForWeightsAndBiasToCommandBuffer:gradientMatrix:inputMatrix:resultGradientForWeightMatrix:resultGradientForBiasVector:"), objc.Ptr(commandBufferObject), objc.Ptr(gradientMatrix), objc.Ptr(inputMatrix), objc.Ptr(resultGradientForWeightMatrix), objc.Ptr(resultGradientForBiasVector))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966665-encodegradientfordatatocommandbu?language=objc
func (m_ MatrixFullyConnectedGradient) EncodeGradientForDataToCommandBufferGradientMatrixWeightMatrixResultGradientForDataMatrix(commandBuffer metal.PCommandBuffer, gradientMatrix IMatrix, weightMatrix IMatrix, resultGradientForDataMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeGradientForDataToCommandBuffer:gradientMatrix:weightMatrix:resultGradientForDataMatrix:"), po0, objc.Ptr(gradientMatrix), objc.Ptr(weightMatrix), objc.Ptr(resultGradientForDataMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966665-encodegradientfordatatocommandbu?language=objc
func (m_ MatrixFullyConnectedGradient) EncodeGradientForDataToCommandBufferObjectGradientMatrixWeightMatrixResultGradientForDataMatrix(commandBufferObject objc.IObject, gradientMatrix IMatrix, weightMatrix IMatrix, resultGradientForDataMatrix IMatrix) {
	objc.Call[objc.Void](m_, objc.Sel("encodeGradientForDataToCommandBuffer:gradientMatrix:weightMatrix:resultGradientForDataMatrix:"), objc.Ptr(commandBufferObject), objc.Ptr(gradientMatrix), objc.Ptr(weightMatrix), objc.Ptr(resultGradientForDataMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966671-sourceoutputfeaturechannels?language=objc
func (m_ MatrixFullyConnectedGradient) SourceOutputFeatureChannels() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceOutputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966671-sourceoutputfeaturechannels?language=objc
func (m_ MatrixFullyConnectedGradient) SetSourceOutputFeatureChannels(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceOutputFeatureChannels:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966663-alpha?language=objc
func (m_ MatrixFullyConnectedGradient) Alpha() float64 {
	rv := objc.Call[float64](m_, objc.Sel("alpha"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966663-alpha?language=objc
func (m_ MatrixFullyConnectedGradient) SetAlpha(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setAlpha:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966670-sourcenumberoffeaturevectors?language=objc
func (m_ MatrixFullyConnectedGradient) SourceNumberOfFeatureVectors() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceNumberOfFeatureVectors"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966670-sourcenumberoffeaturevectors?language=objc
func (m_ MatrixFullyConnectedGradient) SetSourceNumberOfFeatureVectors(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceNumberOfFeatureVectors:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966669-sourceinputfeaturechannels?language=objc
func (m_ MatrixFullyConnectedGradient) SourceInputFeatureChannels() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceInputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfullyconnectedgradient/2966669-sourceinputfeaturechannels?language=objc
func (m_ MatrixFullyConnectedGradient) SetSourceInputFeatureChannels(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceInputFeatureChannels:"), value)
}
