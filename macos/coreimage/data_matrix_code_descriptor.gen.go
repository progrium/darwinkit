// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [DataMatrixCodeDescriptor] class.
var DataMatrixCodeDescriptorClass = _DataMatrixCodeDescriptorClass{objc.GetClass("CIDataMatrixCodeDescriptor")}

type _DataMatrixCodeDescriptorClass struct {
	objc.Class
}

// An interface definition for the [DataMatrixCodeDescriptor] class.
type IDataMatrixCodeDescriptor interface {
	IBarcodeDescriptor
	ColumnCount() int
	RowCount() int
	ErrorCorrectedPayload() []byte
	EccVersion() DataMatrixCodeECCVersion
}

// A concrete subclass of CIBarcodeDescriptor that represents a Data Matrix code symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidatamatrixcodedescriptor?language=objc
type DataMatrixCodeDescriptor struct {
	BarcodeDescriptor
}

func DataMatrixCodeDescriptorFrom(ptr unsafe.Pointer) DataMatrixCodeDescriptor {
	return DataMatrixCodeDescriptor{
		BarcodeDescriptor: BarcodeDescriptorFrom(ptr),
	}
}

func (d_ DataMatrixCodeDescriptor) InitWithPayloadRowCountColumnCountEccVersion(errorCorrectedPayload []byte, rowCount int, columnCount int, eccVersion DataMatrixCodeECCVersion) DataMatrixCodeDescriptor {
	rv := objc.Call[DataMatrixCodeDescriptor](d_, objc.Sel("initWithPayload:rowCount:columnCount:eccVersion:"), errorCorrectedPayload, rowCount, columnCount, eccVersion)
	return rv
}

// Initializes a descriptor that can be used as input to the CIBarcodeGenerator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidatamatrixcodedescriptor/2875201-initwithpayload?language=objc
func NewDataMatrixCodeDescriptorWithPayloadRowCountColumnCountEccVersion(errorCorrectedPayload []byte, rowCount int, columnCount int, eccVersion DataMatrixCodeECCVersion) DataMatrixCodeDescriptor {
	instance := DataMatrixCodeDescriptorClass.Alloc().InitWithPayloadRowCountColumnCountEccVersion(errorCorrectedPayload, rowCount, columnCount, eccVersion)
	instance.Autorelease()
	return instance
}

func (dc _DataMatrixCodeDescriptorClass) DescriptorWithPayloadRowCountColumnCountEccVersion(errorCorrectedPayload []byte, rowCount int, columnCount int, eccVersion DataMatrixCodeECCVersion) DataMatrixCodeDescriptor {
	rv := objc.Call[DataMatrixCodeDescriptor](dc, objc.Sel("descriptorWithPayload:rowCount:columnCount:eccVersion:"), errorCorrectedPayload, rowCount, columnCount, eccVersion)
	return rv
}

// Creates a Data Matrix code descriptor encoding the given payload and parameters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidatamatrixcodedescriptor/2875163-descriptorwithpayload?language=objc
func DataMatrixCodeDescriptor_DescriptorWithPayloadRowCountColumnCountEccVersion(errorCorrectedPayload []byte, rowCount int, columnCount int, eccVersion DataMatrixCodeECCVersion) DataMatrixCodeDescriptor {
	return DataMatrixCodeDescriptorClass.DescriptorWithPayloadRowCountColumnCountEccVersion(errorCorrectedPayload, rowCount, columnCount, eccVersion)
}

func (dc _DataMatrixCodeDescriptorClass) Alloc() DataMatrixCodeDescriptor {
	rv := objc.Call[DataMatrixCodeDescriptor](dc, objc.Sel("alloc"))
	return rv
}

func DataMatrixCodeDescriptor_Alloc() DataMatrixCodeDescriptor {
	return DataMatrixCodeDescriptorClass.Alloc()
}

func (dc _DataMatrixCodeDescriptorClass) New() DataMatrixCodeDescriptor {
	rv := objc.Call[DataMatrixCodeDescriptor](dc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewDataMatrixCodeDescriptor() DataMatrixCodeDescriptor {
	return DataMatrixCodeDescriptorClass.New()
}

func (d_ DataMatrixCodeDescriptor) Init() DataMatrixCodeDescriptor {
	rv := objc.Call[DataMatrixCodeDescriptor](d_, objc.Sel("init"))
	return rv
}

// The number of module columns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidatamatrixcodedescriptor/2875202-columncount?language=objc
func (d_ DataMatrixCodeDescriptor) ColumnCount() int {
	rv := objc.Call[int](d_, objc.Sel("columnCount"))
	return rv
}

// The number of module rows. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidatamatrixcodedescriptor/2875200-rowcount?language=objc
func (d_ DataMatrixCodeDescriptor) RowCount() int {
	rv := objc.Call[int](d_, objc.Sel("rowCount"))
	return rv
}

// The error-corrected payload that comprises the Data Matrix code symbol. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidatamatrixcodedescriptor/2875173-errorcorrectedpayload?language=objc
func (d_ DataMatrixCodeDescriptor) ErrorCorrectedPayload() []byte {
	rv := objc.Call[[]byte](d_, objc.Sel("errorCorrectedPayload"))
	return rv
}

// The Data Matrix code ECC version. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidatamatrixcodedescriptor/2875170-eccversion?language=objc
func (d_ DataMatrixCodeDescriptor) EccVersion() DataMatrixCodeECCVersion {
	rv := objc.Call[DataMatrixCodeECCVersion](d_, objc.Sel("eccVersion"))
	return rv
}
