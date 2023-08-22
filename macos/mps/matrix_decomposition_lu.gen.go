// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixDecompositionLU] class.
var MatrixDecompositionLUClass = _MatrixDecompositionLUClass{objc.GetClass("MPSMatrixDecompositionLU")}

type _MatrixDecompositionLUClass struct {
	objc.Class
}

// An interface definition for the [MatrixDecompositionLU] class.
type IMatrixDecompositionLU interface {
	IMatrixUnaryKernel
	EncodeToCommandBufferSourceMatrixResultMatrixPivotIndicesStatus(commandBuffer metal.PCommandBuffer, sourceMatrix IMatrix, resultMatrix IMatrix, pivotIndices IMatrix, status metal.PBuffer)
	EncodeToCommandBufferObjectSourceMatrixResultMatrixPivotIndicesStatusObject(commandBufferObject objc.IObject, sourceMatrix IMatrix, resultMatrix IMatrix, pivotIndices IMatrix, statusObject objc.IObject)
}

// A kernel for computing the LU factorization of a matrix using partial pivoting with row interchanges. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositionlu?language=objc
type MatrixDecompositionLU struct {
	MatrixUnaryKernel
}

func MatrixDecompositionLUFrom(ptr unsafe.Pointer) MatrixDecompositionLU {
	return MatrixDecompositionLU{
		MatrixUnaryKernel: MatrixUnaryKernelFrom(ptr),
	}
}

func (m_ MatrixDecompositionLU) InitWithDeviceRowsColumns(device metal.PDevice, rows uint, columns uint) MatrixDecompositionLU {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixDecompositionLU](m_, objc.Sel("initWithDevice:rows:columns:"), po0, rows, columns)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositionlu/2866960-initwithdevice?language=objc
func NewMatrixDecompositionLUWithDeviceRowsColumns(device metal.PDevice, rows uint, columns uint) MatrixDecompositionLU {
	instance := MatrixDecompositionLUClass.Alloc().InitWithDeviceRowsColumns(device, rows, columns)
	instance.Autorelease()
	return instance
}

func (mc _MatrixDecompositionLUClass) Alloc() MatrixDecompositionLU {
	rv := objc.Call[MatrixDecompositionLU](mc, objc.Sel("alloc"))
	return rv
}

func MatrixDecompositionLU_Alloc() MatrixDecompositionLU {
	return MatrixDecompositionLUClass.Alloc()
}

func (mc _MatrixDecompositionLUClass) New() MatrixDecompositionLU {
	rv := objc.Call[MatrixDecompositionLU](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixDecompositionLU() MatrixDecompositionLU {
	return MatrixDecompositionLUClass.New()
}

func (m_ MatrixDecompositionLU) Init() MatrixDecompositionLU {
	rv := objc.Call[MatrixDecompositionLU](m_, objc.Sel("init"))
	return rv
}

func (m_ MatrixDecompositionLU) InitWithDevice(device metal.PDevice) MatrixDecompositionLU {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixDecompositionLU](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewMatrixDecompositionLUWithDevice(device metal.PDevice) MatrixDecompositionLU {
	instance := MatrixDecompositionLUClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixDecompositionLU) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixDecompositionLU {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixDecompositionLU](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func MatrixDecompositionLU_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixDecompositionLU {
	instance := MatrixDecompositionLUClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositionlu/2867184-encodetocommandbuffer?language=objc
func (m_ MatrixDecompositionLU) EncodeToCommandBufferSourceMatrixResultMatrixPivotIndicesStatus(commandBuffer metal.PCommandBuffer, sourceMatrix IMatrix, resultMatrix IMatrix, pivotIndices IMatrix, status metal.PBuffer) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	po4 := objc.WrapAsProtocol("MTLBuffer", status)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:sourceMatrix:resultMatrix:pivotIndices:status:"), po0, objc.Ptr(sourceMatrix), objc.Ptr(resultMatrix), objc.Ptr(pivotIndices), po4)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdecompositionlu/2867184-encodetocommandbuffer?language=objc
func (m_ MatrixDecompositionLU) EncodeToCommandBufferObjectSourceMatrixResultMatrixPivotIndicesStatusObject(commandBufferObject objc.IObject, sourceMatrix IMatrix, resultMatrix IMatrix, pivotIndices IMatrix, statusObject objc.IObject) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:sourceMatrix:resultMatrix:pivotIndices:status:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceMatrix), objc.Ptr(resultMatrix), objc.Ptr(pivotIndices), objc.Ptr(statusObject))
}
