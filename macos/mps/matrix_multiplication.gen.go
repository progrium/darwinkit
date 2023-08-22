// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixMultiplication] class.
var MatrixMultiplicationClass = _MatrixMultiplicationClass{objc.GetClass("MPSMatrixMultiplication")}

type _MatrixMultiplicationClass struct {
	objc.Class
}

// An interface definition for the [MatrixMultiplication] class.
type IMatrixMultiplication interface {
	IKernel
	EncodeToCommandBufferLeftMatrixRightMatrixResultMatrix(commandBuffer metal.PCommandBuffer, leftMatrix IMatrix, rightMatrix IMatrix, resultMatrix IMatrix)
	EncodeToCommandBufferObjectLeftMatrixRightMatrixResultMatrix(commandBufferObject objc.IObject, leftMatrix IMatrix, rightMatrix IMatrix, resultMatrix IMatrix)
	ResultMatrixOrigin() metal.Origin
	SetResultMatrixOrigin(value metal.Origin)
	RightMatrixOrigin() metal.Origin
	SetRightMatrixOrigin(value metal.Origin)
	BatchStart() uint
	SetBatchStart(value uint)
	BatchSize() uint
	SetBatchSize(value uint)
	LeftMatrixOrigin() metal.Origin
	SetLeftMatrixOrigin(value metal.Origin)
}

// A matrix multiplication kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication?language=objc
type MatrixMultiplication struct {
	Kernel
}

func MatrixMultiplicationFrom(ptr unsafe.Pointer) MatrixMultiplication {
	return MatrixMultiplication{
		Kernel: KernelFrom(ptr),
	}
}

func (m_ MatrixMultiplication) InitWithDeviceResultRowsResultColumnsInteriorColumns(device metal.PDevice, resultRows uint, resultColumns uint, interiorColumns uint) MatrixMultiplication {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixMultiplication](m_, objc.Sel("initWithDevice:resultRows:resultColumns:interiorColumns:"), po0, resultRows, resultColumns, interiorColumns)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2909034-initwithdevice?language=objc
func NewMatrixMultiplicationWithDeviceResultRowsResultColumnsInteriorColumns(device metal.PDevice, resultRows uint, resultColumns uint, interiorColumns uint) MatrixMultiplication {
	instance := MatrixMultiplicationClass.Alloc().InitWithDeviceResultRowsResultColumnsInteriorColumns(device, resultRows, resultColumns, interiorColumns)
	instance.Autorelease()
	return instance
}

func (mc _MatrixMultiplicationClass) Alloc() MatrixMultiplication {
	rv := objc.Call[MatrixMultiplication](mc, objc.Sel("alloc"))
	return rv
}

func MatrixMultiplication_Alloc() MatrixMultiplication {
	return MatrixMultiplicationClass.Alloc()
}

func (mc _MatrixMultiplicationClass) New() MatrixMultiplication {
	rv := objc.Call[MatrixMultiplication](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixMultiplication() MatrixMultiplication {
	return MatrixMultiplicationClass.New()
}

func (m_ MatrixMultiplication) Init() MatrixMultiplication {
	rv := objc.Call[MatrixMultiplication](m_, objc.Sel("init"))
	return rv
}

func (m_ MatrixMultiplication) InitWithDevice(device metal.PDevice) MatrixMultiplication {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixMultiplication](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewMatrixMultiplicationWithDevice(device metal.PDevice) MatrixMultiplication {
	instance := MatrixMultiplicationClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixMultiplication) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixMultiplication {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixMultiplication](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func MatrixMultiplication_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixMultiplication {
	instance := MatrixMultiplicationClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// Encodes a matrix multiplication kernel to a command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2147848-encodetocommandbuffer?language=objc
func (m_ MatrixMultiplication) EncodeToCommandBufferLeftMatrixRightMatrixResultMatrix(commandBuffer metal.PCommandBuffer, leftMatrix IMatrix, rightMatrix IMatrix, resultMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:leftMatrix:rightMatrix:resultMatrix:"), po0, objc.Ptr(leftMatrix), objc.Ptr(rightMatrix), objc.Ptr(resultMatrix))
}

// Encodes a matrix multiplication kernel to a command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2147848-encodetocommandbuffer?language=objc
func (m_ MatrixMultiplication) EncodeToCommandBufferObjectLeftMatrixRightMatrixResultMatrix(commandBufferObject objc.IObject, leftMatrix IMatrix, rightMatrix IMatrix, resultMatrix IMatrix) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:leftMatrix:rightMatrix:resultMatrix:"), objc.Ptr(commandBufferObject), objc.Ptr(leftMatrix), objc.Ptr(rightMatrix), objc.Ptr(resultMatrix))
}

// The origin of the result matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2147847-resultmatrixorigin?language=objc
func (m_ MatrixMultiplication) ResultMatrixOrigin() metal.Origin {
	rv := objc.Call[metal.Origin](m_, objc.Sel("resultMatrixOrigin"))
	return rv
}

// The origin of the result matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2147847-resultmatrixorigin?language=objc
func (m_ MatrixMultiplication) SetResultMatrixOrigin(value metal.Origin) {
	objc.Call[objc.Void](m_, objc.Sel("setResultMatrixOrigin:"), value)
}

// The origin of the right input matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2147851-rightmatrixorigin?language=objc
func (m_ MatrixMultiplication) RightMatrixOrigin() metal.Origin {
	rv := objc.Call[metal.Origin](m_, objc.Sel("rightMatrixOrigin"))
	return rv
}

// The origin of the right input matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2147851-rightmatrixorigin?language=objc
func (m_ MatrixMultiplication) SetRightMatrixOrigin(value metal.Origin) {
	objc.Call[objc.Void](m_, objc.Sel("setRightMatrixOrigin:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2873081-batchstart?language=objc
func (m_ MatrixMultiplication) BatchStart() uint {
	rv := objc.Call[uint](m_, objc.Sel("batchStart"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2873081-batchstart?language=objc
func (m_ MatrixMultiplication) SetBatchStart(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setBatchStart:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2873082-batchsize?language=objc
func (m_ MatrixMultiplication) BatchSize() uint {
	rv := objc.Call[uint](m_, objc.Sel("batchSize"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2873082-batchsize?language=objc
func (m_ MatrixMultiplication) SetBatchSize(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setBatchSize:"), value)
}

// The origin of the left input matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2147846-leftmatrixorigin?language=objc
func (m_ MatrixMultiplication) LeftMatrixOrigin() metal.Origin {
	rv := objc.Call[metal.Origin](m_, objc.Sel("leftMatrixOrigin"))
	return rv
}

// The origin of the left input matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixmultiplication/2147846-leftmatrixorigin?language=objc
func (m_ MatrixMultiplication) SetLeftMatrixOrigin(value metal.Origin) {
	objc.Call[objc.Void](m_, objc.Sel("setLeftMatrixOrigin:"), value)
}
