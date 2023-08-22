// AUTO-GENERATED CODE, DO NOT MODIFY

package mpsgraph

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/mps"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CreateSparseOpDescriptor] class.
var CreateSparseOpDescriptorClass = _CreateSparseOpDescriptorClass{objc.GetClass("MPSGraphCreateSparseOpDescriptor")}

type _CreateSparseOpDescriptorClass struct {
	objc.Class
}

// An interface definition for the [CreateSparseOpDescriptor] class.
type ICreateSparseOpDescriptor interface {
	objc.IObject
	SparseStorageType() SparseStorageType
	SetSparseStorageType(value SparseStorageType)
	DataType() mps.DataType
	SetDataType(value mps.DataType)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcreatesparseopdescriptor?language=objc
type CreateSparseOpDescriptor struct {
	objc.Object
}

func CreateSparseOpDescriptorFrom(ptr unsafe.Pointer) CreateSparseOpDescriptor {
	return CreateSparseOpDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (cc _CreateSparseOpDescriptorClass) DescriptorWithStorageTypeDataType(sparseStorageType SparseStorageType, dataType mps.DataType) CreateSparseOpDescriptor {
	rv := objc.Call[CreateSparseOpDescriptor](cc, objc.Sel("descriptorWithStorageType:dataType:"), sparseStorageType, dataType)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcreatesparseopdescriptor/3763061-descriptorwithstoragetype?language=objc
func CreateSparseOpDescriptor_DescriptorWithStorageTypeDataType(sparseStorageType SparseStorageType, dataType mps.DataType) CreateSparseOpDescriptor {
	return CreateSparseOpDescriptorClass.DescriptorWithStorageTypeDataType(sparseStorageType, dataType)
}

func (cc _CreateSparseOpDescriptorClass) Alloc() CreateSparseOpDescriptor {
	rv := objc.Call[CreateSparseOpDescriptor](cc, objc.Sel("alloc"))
	return rv
}

func CreateSparseOpDescriptor_Alloc() CreateSparseOpDescriptor {
	return CreateSparseOpDescriptorClass.Alloc()
}

func (cc _CreateSparseOpDescriptorClass) New() CreateSparseOpDescriptor {
	rv := objc.Call[CreateSparseOpDescriptor](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCreateSparseOpDescriptor() CreateSparseOpDescriptor {
	return CreateSparseOpDescriptorClass.New()
}

func (c_ CreateSparseOpDescriptor) Init() CreateSparseOpDescriptor {
	rv := objc.Call[CreateSparseOpDescriptor](c_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcreatesparseopdescriptor/3763062-sparsestoragetype?language=objc
func (c_ CreateSparseOpDescriptor) SparseStorageType() SparseStorageType {
	rv := objc.Call[SparseStorageType](c_, objc.Sel("sparseStorageType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcreatesparseopdescriptor/3763062-sparsestoragetype?language=objc
func (c_ CreateSparseOpDescriptor) SetSparseStorageType(value SparseStorageType) {
	objc.Call[objc.Void](c_, objc.Sel("setSparseStorageType:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcreatesparseopdescriptor/3763060-datatype?language=objc
func (c_ CreateSparseOpDescriptor) DataType() mps.DataType {
	rv := objc.Call[mps.DataType](c_, objc.Sel("dataType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshadersgraph/mpsgraphcreatesparseopdescriptor/3763060-datatype?language=objc
func (c_ CreateSparseOpDescriptor) SetDataType(value mps.DataType) {
	objc.Call[objc.Void](c_, objc.Sel("setDataType:"), value)
}
