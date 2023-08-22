// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixSolveCholesky] class.
var MatrixSolveCholeskyClass = _MatrixSolveCholeskyClass{objc.GetClass("MPSMatrixSolveCholesky")}

type _MatrixSolveCholeskyClass struct {
	objc.Class
}

// An interface definition for the [MatrixSolveCholesky] class.
type IMatrixSolveCholesky interface {
	IMatrixBinaryKernel
	EncodeToCommandBufferSourceMatrixRightHandSideMatrixSolutionMatrix(commandBuffer metal.PCommandBuffer, sourceMatrix IMatrix, rightHandSideMatrix IMatrix, solutionMatrix IMatrix)
	EncodeToCommandBufferObjectSourceMatrixRightHandSideMatrixSolutionMatrix(commandBufferObject objc.IObject, sourceMatrix IMatrix, rightHandSideMatrix IMatrix, solutionMatrix IMatrix)
}

// A kernel for computing the solution of a linear system of equations using a Cholesky factorization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvecholesky?language=objc
type MatrixSolveCholesky struct {
	MatrixBinaryKernel
}

func MatrixSolveCholeskyFrom(ptr unsafe.Pointer) MatrixSolveCholesky {
	return MatrixSolveCholesky{
		MatrixBinaryKernel: MatrixBinaryKernelFrom(ptr),
	}
}

func (m_ MatrixSolveCholesky) InitWithDeviceUpperOrderNumberOfRightHandSides(device metal.PDevice, upper bool, order uint, numberOfRightHandSides uint) MatrixSolveCholesky {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixSolveCholesky](m_, objc.Sel("initWithDevice:upper:order:numberOfRightHandSides:"), po0, upper, order, numberOfRightHandSides)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvecholesky/2873006-initwithdevice?language=objc
func NewMatrixSolveCholeskyWithDeviceUpperOrderNumberOfRightHandSides(device metal.PDevice, upper bool, order uint, numberOfRightHandSides uint) MatrixSolveCholesky {
	instance := MatrixSolveCholeskyClass.Alloc().InitWithDeviceUpperOrderNumberOfRightHandSides(device, upper, order, numberOfRightHandSides)
	instance.Autorelease()
	return instance
}

func (mc _MatrixSolveCholeskyClass) Alloc() MatrixSolveCholesky {
	rv := objc.Call[MatrixSolveCholesky](mc, objc.Sel("alloc"))
	return rv
}

func MatrixSolveCholesky_Alloc() MatrixSolveCholesky {
	return MatrixSolveCholeskyClass.Alloc()
}

func (mc _MatrixSolveCholeskyClass) New() MatrixSolveCholesky {
	rv := objc.Call[MatrixSolveCholesky](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixSolveCholesky() MatrixSolveCholesky {
	return MatrixSolveCholeskyClass.New()
}

func (m_ MatrixSolveCholesky) Init() MatrixSolveCholesky {
	rv := objc.Call[MatrixSolveCholesky](m_, objc.Sel("init"))
	return rv
}

func (m_ MatrixSolveCholesky) InitWithDevice(device metal.PDevice) MatrixSolveCholesky {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixSolveCholesky](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewMatrixSolveCholeskyWithDevice(device metal.PDevice) MatrixSolveCholesky {
	instance := MatrixSolveCholeskyClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixSolveCholesky) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixSolveCholesky {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixSolveCholesky](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func MatrixSolveCholesky_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixSolveCholesky {
	instance := MatrixSolveCholeskyClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvecholesky/2866957-encodetocommandbuffer?language=objc
func (m_ MatrixSolveCholesky) EncodeToCommandBufferSourceMatrixRightHandSideMatrixSolutionMatrix(commandBuffer metal.PCommandBuffer, sourceMatrix IMatrix, rightHandSideMatrix IMatrix, solutionMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:sourceMatrix:rightHandSideMatrix:solutionMatrix:"), po0, objc.Ptr(sourceMatrix), objc.Ptr(rightHandSideMatrix), objc.Ptr(solutionMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsolvecholesky/2866957-encodetocommandbuffer?language=objc
func (m_ MatrixSolveCholesky) EncodeToCommandBufferObjectSourceMatrixRightHandSideMatrixSolutionMatrix(commandBufferObject objc.IObject, sourceMatrix IMatrix, rightHandSideMatrix IMatrix, solutionMatrix IMatrix) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:sourceMatrix:rightHandSideMatrix:solutionMatrix:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceMatrix), objc.Ptr(rightHandSideMatrix), objc.Ptr(solutionMatrix))
}
