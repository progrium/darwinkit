// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixSum] class.
var MatrixSumClass = _MatrixSumClass{objc.GetClass("MPSMatrixSum")}

type _MatrixSumClass struct {
	objc.Class
}

// An interface definition for the [MatrixSum] class.
type IMatrixSum interface {
	IKernel
	NeuronType() CNNNeuronType
	SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float64, parameterB float64, parameterC float64)
	EncodeToCommandBufferSourceMatricesResultMatrixScaleVectorOffsetVectorBiasVectorStartIndex(buffer metal.PCommandBuffer, sourceMatrices []IMatrix, resultMatrix IMatrix, scaleVector IVector, offsetVector IVector, biasVector IVector, startIndex uint)
	EncodeToCommandBufferObjectSourceMatricesResultMatrixScaleVectorOffsetVectorBiasVectorStartIndex(bufferObject objc.IObject, sourceMatrices []IMatrix, resultMatrix IMatrix, scaleVector IVector, offsetVector IVector, biasVector IVector, startIndex uint)
	NeuronParameterA() float64
	ResultMatrixOrigin() metal.Origin
	SetResultMatrixOrigin(value metal.Origin)
	Transpose() bool
	Columns() uint
	Count() uint
	Rows() uint
	NeuronParameterB() float64
	NeuronParameterC() float64
}

// A kernel for performing a pointwise summation of a matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum?language=objc
type MatrixSum struct {
	Kernel
}

func MatrixSumFrom(ptr unsafe.Pointer) MatrixSum {
	return MatrixSum{
		Kernel: KernelFrom(ptr),
	}
}

func (m_ MatrixSum) InitWithDeviceCountRowsColumnsTranspose(device metal.PDevice, count uint, rows uint, columns uint, transpose bool) MatrixSum {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixSum](m_, objc.Sel("initWithDevice:count:rows:columns:transpose:"), po0, count, rows, columns, transpose)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935623-initwithdevice?language=objc
func NewMatrixSumWithDeviceCountRowsColumnsTranspose(device metal.PDevice, count uint, rows uint, columns uint, transpose bool) MatrixSum {
	instance := MatrixSumClass.Alloc().InitWithDeviceCountRowsColumnsTranspose(device, count, rows, columns, transpose)
	instance.Autorelease()
	return instance
}

func (mc _MatrixSumClass) Alloc() MatrixSum {
	rv := objc.Call[MatrixSum](mc, objc.Sel("alloc"))
	return rv
}

func MatrixSum_Alloc() MatrixSum {
	return MatrixSumClass.Alloc()
}

func (mc _MatrixSumClass) New() MatrixSum {
	rv := objc.Call[MatrixSum](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixSum() MatrixSum {
	return MatrixSumClass.New()
}

func (m_ MatrixSum) Init() MatrixSum {
	rv := objc.Call[MatrixSum](m_, objc.Sel("init"))
	return rv
}

func (m_ MatrixSum) InitWithDevice(device metal.PDevice) MatrixSum {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixSum](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewMatrixSumWithDevice(device metal.PDevice) MatrixSum {
	instance := MatrixSumClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixSum) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixSum {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixSum](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func MatrixSum_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixSum {
	instance := MatrixSumClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935625-neurontype?language=objc
func (m_ MatrixSum) NeuronType() CNNNeuronType {
	rv := objc.Call[CNNNeuronType](m_, objc.Sel("neuronType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935617-setneurontype?language=objc
func (m_ MatrixSum) SetNeuronTypeParameterAParameterBParameterC(neuronType CNNNeuronType, parameterA float64, parameterB float64, parameterC float64) {
	objc.Call[objc.Void](m_, objc.Sel("setNeuronType:parameterA:parameterB:parameterC:"), neuronType, parameterA, parameterB, parameterC)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935613-encodetocommandbuffer?language=objc
func (m_ MatrixSum) EncodeToCommandBufferSourceMatricesResultMatrixScaleVectorOffsetVectorBiasVectorStartIndex(buffer metal.PCommandBuffer, sourceMatrices []IMatrix, resultMatrix IMatrix, scaleVector IVector, offsetVector IVector, biasVector IVector, startIndex uint) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", buffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:sourceMatrices:resultMatrix:scaleVector:offsetVector:biasVector:startIndex:"), po0, sourceMatrices, objc.Ptr(resultMatrix), objc.Ptr(scaleVector), objc.Ptr(offsetVector), objc.Ptr(biasVector), startIndex)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935613-encodetocommandbuffer?language=objc
func (m_ MatrixSum) EncodeToCommandBufferObjectSourceMatricesResultMatrixScaleVectorOffsetVectorBiasVectorStartIndex(bufferObject objc.IObject, sourceMatrices []IMatrix, resultMatrix IMatrix, scaleVector IVector, offsetVector IVector, biasVector IVector, startIndex uint) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:sourceMatrices:resultMatrix:scaleVector:offsetVector:biasVector:startIndex:"), objc.Ptr(bufferObject), sourceMatrices, objc.Ptr(resultMatrix), objc.Ptr(scaleVector), objc.Ptr(offsetVector), objc.Ptr(biasVector), startIndex)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935624-neuronparametera?language=objc
func (m_ MatrixSum) NeuronParameterA() float64 {
	rv := objc.Call[float64](m_, objc.Sel("neuronParameterA"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/3152564-resultmatrixorigin?language=objc
func (m_ MatrixSum) ResultMatrixOrigin() metal.Origin {
	rv := objc.Call[metal.Origin](m_, objc.Sel("resultMatrixOrigin"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/3152564-resultmatrixorigin?language=objc
func (m_ MatrixSum) SetResultMatrixOrigin(value metal.Origin) {
	objc.Call[objc.Void](m_, objc.Sel("setResultMatrixOrigin:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935621-transpose?language=objc
func (m_ MatrixSum) Transpose() bool {
	rv := objc.Call[bool](m_, objc.Sel("transpose"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935615-columns?language=objc
func (m_ MatrixSum) Columns() uint {
	rv := objc.Call[uint](m_, objc.Sel("columns"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935620-count?language=objc
func (m_ MatrixSum) Count() uint {
	rv := objc.Call[uint](m_, objc.Sel("count"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935622-rows?language=objc
func (m_ MatrixSum) Rows() uint {
	rv := objc.Call[uint](m_, objc.Sel("rows"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935616-neuronparameterb?language=objc
func (m_ MatrixSum) NeuronParameterB() float64 {
	rv := objc.Call[float64](m_, objc.Sel("neuronParameterB"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsum/2935618-neuronparameterc?language=objc
func (m_ MatrixSum) NeuronParameterC() float64 {
	rv := objc.Call[float64](m_, objc.Sel("neuronParameterC"))
	return rv
}
