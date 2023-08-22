// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixDescriptor] class.
var MatrixDescriptorClass = _MatrixDescriptorClass{objc.GetClass("MPSMatrixDescriptor")}

type _MatrixDescriptorClass struct {
	objc.Class
}

// An interface definition for the [MatrixDescriptor] class.
type IMatrixDescriptor interface {
	objc.IObject
	RowBytes() uint
	SetRowBytes(value uint)
	Columns() uint
	SetColumns(value uint)
	DataType() DataType
	SetDataType(value DataType)
	Rows() uint
	SetRows(value uint)
	Matrices() uint
	MatrixBytes() uint
}

// A description of attributes used to create an MPS matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor?language=objc
type MatrixDescriptor struct {
	objc.Object
}

func MatrixDescriptorFrom(ptr unsafe.Pointer) MatrixDescriptor {
	return MatrixDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (mc _MatrixDescriptorClass) MatrixDescriptorWithDimensionsColumnsRowBytesDataType(rows uint, columns uint, rowBytes uint, dataType DataType) MatrixDescriptor {
	rv := objc.Call[MatrixDescriptor](mc, objc.Sel("matrixDescriptorWithDimensions:columns:rowBytes:dataType:"), rows, columns, rowBytes, dataType)
	return rv
}

// Creates a matrix descriptor with the specified dimensions and data type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143206-matrixdescriptorwithdimensions?language=objc
func MatrixDescriptor_MatrixDescriptorWithDimensionsColumnsRowBytesDataType(rows uint, columns uint, rowBytes uint, dataType DataType) MatrixDescriptor {
	return MatrixDescriptorClass.MatrixDescriptorWithDimensionsColumnsRowBytesDataType(rows, columns, rowBytes, dataType)
}

func (mc _MatrixDescriptorClass) MatrixDescriptorWithRowsColumnsMatricesRowBytesMatrixBytesDataType(rows uint, columns uint, matrices uint, rowBytes uint, matrixBytes uint, dataType DataType) MatrixDescriptor {
	rv := objc.Call[MatrixDescriptor](mc, objc.Sel("matrixDescriptorWithRows:columns:matrices:rowBytes:matrixBytes:dataType:"), rows, columns, matrices, rowBytes, matrixBytes, dataType)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2873350-matrixdescriptorwithrows?language=objc
func MatrixDescriptor_MatrixDescriptorWithRowsColumnsMatricesRowBytesMatrixBytesDataType(rows uint, columns uint, matrices uint, rowBytes uint, matrixBytes uint, dataType DataType) MatrixDescriptor {
	return MatrixDescriptorClass.MatrixDescriptorWithRowsColumnsMatricesRowBytesMatrixBytesDataType(rows, columns, matrices, rowBytes, matrixBytes, dataType)
}

func (mc _MatrixDescriptorClass) Alloc() MatrixDescriptor {
	rv := objc.Call[MatrixDescriptor](mc, objc.Sel("alloc"))
	return rv
}

func MatrixDescriptor_Alloc() MatrixDescriptor {
	return MatrixDescriptorClass.Alloc()
}

func (mc _MatrixDescriptorClass) New() MatrixDescriptor {
	rv := objc.Call[MatrixDescriptor](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixDescriptor() MatrixDescriptor {
	return MatrixDescriptorClass.New()
}

func (m_ MatrixDescriptor) Init() MatrixDescriptor {
	rv := objc.Call[MatrixDescriptor](m_, objc.Sel("init"))
	return rv
}

// Determines the recommended matrix row stride, in bytes, for a given number of columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143204-rowbytesfromcolumns?language=objc
func (mc _MatrixDescriptorClass) RowBytesFromColumnsDataType(columns uint, dataType DataType) uint {
	rv := objc.Call[uint](mc, objc.Sel("rowBytesFromColumns:dataType:"), columns, dataType)
	return rv
}

// Determines the recommended matrix row stride, in bytes, for a given number of columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143204-rowbytesfromcolumns?language=objc
func MatrixDescriptor_RowBytesFromColumnsDataType(columns uint, dataType DataType) uint {
	return MatrixDescriptorClass.RowBytesFromColumnsDataType(columns, dataType)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2873394-rowbytesforcolumns?language=objc
func (mc _MatrixDescriptorClass) RowBytesForColumnsDataType(columns uint, dataType DataType) uint {
	rv := objc.Call[uint](mc, objc.Sel("rowBytesForColumns:dataType:"), columns, dataType)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2873394-rowbytesforcolumns?language=objc
func MatrixDescriptor_RowBytesForColumnsDataType(columns uint, dataType DataType) uint {
	return MatrixDescriptorClass.RowBytesForColumnsDataType(columns, dataType)
}

// The stride, in bytes, between corresponding elements of consecutive rows in the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143199-rowbytes?language=objc
func (m_ MatrixDescriptor) RowBytes() uint {
	rv := objc.Call[uint](m_, objc.Sel("rowBytes"))
	return rv
}

// The stride, in bytes, between corresponding elements of consecutive rows in the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143199-rowbytes?language=objc
func (m_ MatrixDescriptor) SetRowBytes(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setRowBytes:"), value)
}

// The number of columns in the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143196-columns?language=objc
func (m_ MatrixDescriptor) Columns() uint {
	rv := objc.Call[uint](m_, objc.Sel("columns"))
	return rv
}

// The number of columns in the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143196-columns?language=objc
func (m_ MatrixDescriptor) SetColumns(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setColumns:"), value)
}

// The type of the values in the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143202-datatype?language=objc
func (m_ MatrixDescriptor) DataType() DataType {
	rv := objc.Call[DataType](m_, objc.Sel("dataType"))
	return rv
}

// The type of the values in the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143202-datatype?language=objc
func (m_ MatrixDescriptor) SetDataType(value DataType) {
	objc.Call[objc.Void](m_, objc.Sel("setDataType:"), value)
}

// The number of rows in the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143203-rows?language=objc
func (m_ MatrixDescriptor) Rows() uint {
	rv := objc.Call[uint](m_, objc.Sel("rows"))
	return rv
}

// The number of rows in the matrix. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2143203-rows?language=objc
func (m_ MatrixDescriptor) SetRows(value uint) {
	objc.Call[objc.Void](m_, objc.Sel("setRows:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2873351-matrices?language=objc
func (m_ MatrixDescriptor) Matrices() uint {
	rv := objc.Call[uint](m_, objc.Sel("matrices"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixdescriptor/2873387-matrixbytes?language=objc
func (m_ MatrixDescriptor) MatrixBytes() uint {
	rv := objc.Call[uint](m_, objc.Sel("matrixBytes"))
	return rv
}
