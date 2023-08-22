// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixSoftMax] class.
var MatrixSoftMaxClass = _MatrixSoftMaxClass{objc.GetClass("MPSMatrixSoftMax")}

type _MatrixSoftMaxClass struct {
	objc.Class
}

// An interface definition for the [MatrixSoftMax] class.
type IMatrixSoftMax interface {
	IMatrixUnaryKernel
	EncodeToCommandBufferInputMatrixResultMatrix(commandBuffer metal.PCommandBuffer, inputMatrix IMatrix, resultMatrix IMatrix)
	EncodeToCommandBufferObjectInputMatrixResultMatrix(commandBufferObject objc.IObject, inputMatrix IMatrix, resultMatrix IMatrix)
	SourceRows() uint
	SetSourceRows(value uint)
	SourceColumns() uint
	SetSourceColumns(value uint)
}

// A softmax kernel that operates on matrices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax?language=objc
type MatrixSoftMax struct {
	MatrixUnaryKernel
}

func MatrixSoftMaxFrom(ptr unsafe.Pointer) MatrixSoftMax {
	return MatrixSoftMax{
		MatrixUnaryKernel: MatrixUnaryKernelFrom(ptr),
	}
}

func (m_ MatrixSoftMax) InitWithDevice(device metal.PDevice) MatrixSoftMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixSoftMax](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935562-initwithdevice?language=objc
func NewMatrixSoftMaxWithDevice(device metal.PDevice) MatrixSoftMax {
	instance := MatrixSoftMaxClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixSoftMax) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixSoftMax {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixSoftMax](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935566-copywithzone?language=objc
func MatrixSoftMax_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixSoftMax {
	instance := MatrixSoftMaxClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (mc _MatrixSoftMaxClass) Alloc() MatrixSoftMax {
	rv := objc.Call[MatrixSoftMax](mc, objc.Sel("alloc"))
	return rv
}

func MatrixSoftMax_Alloc() MatrixSoftMax {
	return MatrixSoftMaxClass.Alloc()
}

func (mc _MatrixSoftMaxClass) New() MatrixSoftMax {
	rv := objc.Call[MatrixSoftMax](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixSoftMax() MatrixSoftMax {
	return MatrixSoftMaxClass.New()
}

func (m_ MatrixSoftMax) Init() MatrixSoftMax {
	rv := objc.Call[MatrixSoftMax](m_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935563-encodetocommandbuffer?language=objc
func (m_ MatrixSoftMax) EncodeToCommandBufferInputMatrixResultMatrix(commandBuffer metal.PCommandBuffer, inputMatrix IMatrix, resultMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:inputMatrix:resultMatrix:"), po0, objc.Ptr(inputMatrix), objc.Ptr(resultMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935563-encodetocommandbuffer?language=objc
func (m_ MatrixSoftMax) EncodeToCommandBufferObjectInputMatrixResultMatrix(commandBufferObject objc.IObject, inputMatrix IMatrix, resultMatrix IMatrix) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:inputMatrix:resultMatrix:"), objc.Ptr(commandBufferObject), objc.Ptr(inputMatrix), objc.Ptr(resultMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935561-sourcerows?language=objc
func (m_ MatrixSoftMax) SourceRows() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceRows"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935561-sourcerows?language=objc
func (m_ MatrixSoftMax) SetSourceRows(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceRows:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935560-sourcecolumns?language=objc
func (m_ MatrixSoftMax) SourceColumns() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceColumns"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935560-sourcecolumns?language=objc
func (m_ MatrixSoftMax) SetSourceColumns(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceColumns:"), value)
}
