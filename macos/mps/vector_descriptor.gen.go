// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VectorDescriptor] class.
var VectorDescriptorClass = _VectorDescriptorClass{objc.GetClass("MPSVectorDescriptor")}

type _VectorDescriptorClass struct {
	objc.Class
}

// An interface definition for the [VectorDescriptor] class.
type IVectorDescriptor interface {
	objc.IObject
	VectorBytes() uint
	Vectors() uint
	Length() uint
	SetLength(value uint)
	DataType() DataType
	SetDataType(value DataType)
}

// A description of the length and data type of a vector. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor?language=objc
type VectorDescriptor struct {
	objc.Object
}

func VectorDescriptorFrom(ptr unsafe.Pointer) VectorDescriptor {
	return VectorDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _VectorDescriptorClass) VectorDescriptorWithLengthDataType(length uint, dataType DataType) VectorDescriptor {
	rv := objc.Call[VectorDescriptor](vc, objc.Sel("vectorDescriptorWithLength:dataType:"), length, dataType)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873328-vectordescriptorwithlength?language=objc
func VectorDescriptor_VectorDescriptorWithLengthDataType(length uint, dataType DataType) VectorDescriptor {
	return VectorDescriptorClass.VectorDescriptorWithLengthDataType(length, dataType)
}

func (vc _VectorDescriptorClass) Alloc() VectorDescriptor {
	rv := objc.Call[VectorDescriptor](vc, objc.Sel("alloc"))
	return rv
}

func VectorDescriptor_Alloc() VectorDescriptor {
	return VectorDescriptorClass.Alloc()
}

func (vc _VectorDescriptorClass) New() VectorDescriptor {
	rv := objc.Call[VectorDescriptor](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVectorDescriptor() VectorDescriptor {
	return VectorDescriptorClass.New()
}

func (v_ VectorDescriptor) Init() VectorDescriptor {
	rv := objc.Call[VectorDescriptor](v_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873337-vectorbytesforlength?language=objc
func (vc _VectorDescriptorClass) VectorBytesForLengthDataType(length uint, dataType DataType) uint {
	rv := objc.Call[uint](vc, objc.Sel("vectorBytesForLength:dataType:"), length, dataType)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873337-vectorbytesforlength?language=objc
func VectorDescriptor_VectorBytesForLengthDataType(length uint, dataType DataType) uint {
	return VectorDescriptorClass.VectorBytesForLengthDataType(length, dataType)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873335-vectorbytes?language=objc
func (v_ VectorDescriptor) VectorBytes() uint {
	rv := objc.Call[uint](v_, objc.Sel("vectorBytes"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873333-vectors?language=objc
func (v_ VectorDescriptor) Vectors() uint {
	rv := objc.Call[uint](v_, objc.Sel("vectors"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873345-length?language=objc
func (v_ VectorDescriptor) Length() uint {
	rv := objc.Call[uint](v_, objc.Sel("length"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873345-length?language=objc
func (v_ VectorDescriptor) SetLength(value uint) {
	objc.Call[objc.Void](v_, objc.Sel("setLength:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873362-datatype?language=objc
func (v_ VectorDescriptor) DataType() DataType {
	rv := objc.Call[DataType](v_, objc.Sel("dataType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsvectordescriptor/2873362-datatype?language=objc
func (v_ VectorDescriptor) SetDataType(value DataType) {
	objc.Call[objc.Void](v_, objc.Sel("setDataType:"), value)
}
