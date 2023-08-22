// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixSolveLU] class.
var MatrixSolveLUClass = _MatrixSolveLUClass{objc.GetClass("MPSMatrixSolveLU")}

type _MatrixSolveLUClass struct {
	objc.Class
}

// An interface definition for the [MatrixSolveLU] class.
type IMatrixSolveLU interface {
	IMatrixBinaryKernel
	EncodeToCommandBufferSourceMatrixRightHandSideMatrixPivotIndicesSolutionMatrix(commandBuffer metal.PCommandBuffer, sourceMatrix IMatrix, rightHandSideMatrix IMatrix, pivotIndices IMatrix, solutionMatrix IMatrix)
	EncodeToCommandBufferObjectSourceMatrixRightHandSideMatrixPivotIndicesSolutionMatrix(commandBufferObject objc.IObject, sourceMatrix IMatrix, rightHandSideMatrix IMatrix, pivotIndices IMatrix, solutionMatrix IMatrix)
}

// A kernel for computing the solution of a linear system of equations using an LU factorization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvelu?language=objc
type MatrixSolveLU struct {
	MatrixBinaryKernel
}

func MatrixSolveLUFrom(ptr unsafe.Pointer) MatrixSolveLU {
	return MatrixSolveLU{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}

func (m_ MatrixSolveLU) InitWithDeviceTransposeOrderNumberOfRightHandSides(device metal.PDevice, transpose bool, order uint, numberOfRightHandSides uint) MatrixSolveLU {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixSolveLU](m_, objc.Sel("initWithDevice:transpose:order:numberOfRightHandSides:"), po0, transpose, order, numberOfRightHandSides)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvelu/2873005-initwithdevice?language=objc
func NewMatrixSolveLUWithDeviceTransposeOrderNumberOfRightHandSides(device metal.PDevice, transpose bool, order uint, numberOfRightHandSides uint) MatrixSolveLU {
	instance := MatrixSolveLUClass.Alloc().InitWithDeviceTransposeOrderNumberOfRightHandSides(device, transpose, order, numberOfRightHandSides)
	instance.Autorelease()
	return instance
}

func (mc _MatrixSolveLUClass) Alloc() MatrixSolveLU {
	rv := objc.Call[MatrixSolveLU](mc, objc.Sel("alloc"))
	return rv
}

func MatrixSolveLU_Alloc() MatrixSolveLU {
	return MatrixSolveLUClass.Alloc()
}

func (mc _MatrixSolveLUClass) New() MatrixSolveLU {
	rv := objc.Call[MatrixSolveLU](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixSolveLU() MatrixSolveLU {
	return MatrixSolveLUClass.New()
}

func (m_ MatrixSolveLU) Init() MatrixSolveLU {
	rv := objc.Call[MatrixSolveLU](m_, objc.Sel("init"))
	return rv
}

func (m_ MatrixSolveLU) InitWithDevice(device metal.PDevice) MatrixSolveLU {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixSolveLU](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewMatrixSolveLUWithDevice(device metal.PDevice) MatrixSolveLU {
	instance := MatrixSolveLUClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixSolveLU) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixSolveLU {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixSolveLU](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func MatrixSolveLU_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixSolveLU {
	instance := MatrixSolveLUClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvelu/2867074-encodetocommandbuffer?language=objc
func (m_ MatrixSolveLU) EncodeToCommandBufferSourceMatrixRightHandSideMatrixPivotIndicesSolutionMatrix(commandBuffer metal.PCommandBuffer, sourceMatrix IMatrix, rightHandSideMatrix IMatrix, pivotIndices IMatrix, solutionMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:sourceMatrix:rightHandSideMatrix:pivotIndices:solutionMatrix:"), po0, objc.Ptr(sourceMatrix), objc.Ptr(rightHandSideMatrix), objc.Ptr(pivotIndices), objc.Ptr(solutionMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvelu/2867074-encodetocommandbuffer?language=objc
func (m_ MatrixSolveLU) EncodeToCommandBufferObjectSourceMatrixRightHandSideMatrixPivotIndicesSolutionMatrix(commandBufferObject objc.IObject, sourceMatrix IMatrix, rightHandSideMatrix IMatrix, pivotIndices IMatrix, solutionMatrix IMatrix) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:sourceMatrix:rightHandSideMatrix:pivotIndices:solutionMatrix:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceMatrix), objc.Ptr(rightHandSideMatrix), objc.Ptr(pivotIndices), objc.Ptr(solutionMatrix))
}
