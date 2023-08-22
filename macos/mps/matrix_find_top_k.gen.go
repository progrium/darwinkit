// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixFindTopK] class.
var MatrixFindTopKClass = _MatrixFindTopKClass{objc.GetClass("MPSMatrixFindTopK")}

type _MatrixFindTopKClass struct {
	objc.Class
}

// An interface definition for the [MatrixFindTopK] class.
type IMatrixFindTopK interface {
	IMatrixUnaryKernel
	EncodeToCommandBufferInputMatrixResultIndexMatrixResultValueMatrix(commandBuffer metal.PCommandBuffer, inputMatrix IMatrix, resultIndexMatrix IMatrix, resultValueMatrix IMatrix)
	EncodeToCommandBufferObjectInputMatrixResultIndexMatrixResultValueMatrix(commandBufferObject objc.IObject, inputMatrix IMatrix, resultIndexMatrix IMatrix, resultValueMatrix IMatrix)
	SourceRows() uint
	SetSourceRows(value uint)
	SourceColumns() uint
	SetSourceColumns(value uint)
	NumberOfTopKValues() uint
	SetNumberOfTopKValues(value uint)
	IndexOffset() uint
	SetIndexOffset(value uint)
}

// A kernel for computing the top-K values and their corresponding indices in a matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk?language=objc
type MatrixFindTopK struct {
	MatrixUnaryKernel
}

func MatrixFindTopKFrom(ptr unsafe.Pointer) MatrixFindTopK {
	return MatrixFindTopK{
		MatrixUnaryKernel: MatrixUnaryKernelFrom(ptr),
	}
}

func (m_ MatrixFindTopK) InitWithDeviceNumberOfTopKValues(device metal.PDevice, numberOfTopKValues uint) MatrixFindTopK {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixFindTopK](m_, objc.Sel("initWithDevice:numberOfTopKValues:"), po0, numberOfTopKValues)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935575-initwithdevice?language=objc
func NewMatrixFindTopKWithDeviceNumberOfTopKValues(device metal.PDevice, numberOfTopKValues uint) MatrixFindTopK {
	instance := MatrixFindTopKClass.Alloc().InitWithDeviceNumberOfTopKValues(device, numberOfTopKValues)
	instance.Autorelease()
	return instance
}

func (m_ MatrixFindTopK) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixFindTopK {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixFindTopK](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935581-copywithzone?language=objc
func MatrixFindTopK_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixFindTopK {
	instance := MatrixFindTopKClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (mc _MatrixFindTopKClass) Alloc() MatrixFindTopK {
	rv := objc.Call[MatrixFindTopK](mc, objc.Sel("alloc"))
	return rv
}

func MatrixFindTopK_Alloc() MatrixFindTopK {
	return MatrixFindTopKClass.Alloc()
}

func (mc _MatrixFindTopKClass) New() MatrixFindTopK {
	rv := objc.Call[MatrixFindTopK](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixFindTopK() MatrixFindTopK {
	return MatrixFindTopKClass.New()
}

func (m_ MatrixFindTopK) Init() MatrixFindTopK {
	rv := objc.Call[MatrixFindTopK](m_, objc.Sel("init"))
	return rv
}

func (m_ MatrixFindTopK) InitWithDevice(device metal.PDevice) MatrixFindTopK {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixFindTopK](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewMatrixFindTopKWithDevice(device metal.PDevice) MatrixFindTopK {
	instance := MatrixFindTopKClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935579-encodetocommandbuffer?language=objc
func (m_ MatrixFindTopK) EncodeToCommandBufferInputMatrixResultIndexMatrixResultValueMatrix(commandBuffer metal.PCommandBuffer, inputMatrix IMatrix, resultIndexMatrix IMatrix, resultValueMatrix IMatrix) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:inputMatrix:resultIndexMatrix:resultValueMatrix:"), po0, objc.Ptr(inputMatrix), objc.Ptr(resultIndexMatrix), objc.Ptr(resultValueMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935579-encodetocommandbuffer?language=objc
func (m_ MatrixFindTopK) EncodeToCommandBufferObjectInputMatrixResultIndexMatrixResultValueMatrix(commandBufferObject objc.IObject, inputMatrix IMatrix, resultIndexMatrix IMatrix, resultValueMatrix IMatrix) {
	objc.Call[objc.Void](m_, objc.Sel("encodeToCommandBuffer:inputMatrix:resultIndexMatrix:resultValueMatrix:"), objc.Ptr(commandBufferObject), objc.Ptr(inputMatrix), objc.Ptr(resultIndexMatrix), objc.Ptr(resultValueMatrix))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935580-sourcerows?language=objc
func (m_ MatrixFindTopK) SourceRows() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceRows"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935580-sourcerows?language=objc
func (m_ MatrixFindTopK) SetSourceRows(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceRows:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935573-sourcecolumns?language=objc
func (m_ MatrixFindTopK) SourceColumns() uint {
	rv := objc.Call[uint](m_, objc.Sel("sourceColumns"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935573-sourcecolumns?language=objc
func (m_ MatrixFindTopK) SetSourceColumns(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setSourceColumns:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935577-numberoftopkvalues?language=objc
func (m_ MatrixFindTopK) NumberOfTopKValues() uint {
	rv := objc.Call[uint](m_, objc.Sel("numberOfTopKValues"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935577-numberoftopkvalues?language=objc
func (m_ MatrixFindTopK) SetNumberOfTopKValues(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setNumberOfTopKValues:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935574-indexoffset?language=objc
func (m_ MatrixFindTopK) IndexOffset() uint {
	rv := objc.Call[uint](m_, objc.Sel("indexOffset"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixfindtopk/2935574-indexoffset?language=objc
func (m_ MatrixFindTopK) SetIndexOffset(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setIndexOffset:"), value)
}
