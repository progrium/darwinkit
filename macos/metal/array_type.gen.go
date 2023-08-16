// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ArrayType] class.
var ArrayTypeClass = _ArrayTypeClass{objc.GetClass("MTLArrayType")}

type _ArrayTypeClass struct {
	objc.Class
}

// An interface definition for the [ArrayType] class.
type IArrayType interface {
	IType
	ElementPointerType() PointerType
	ElementArrayType() ArrayType
	ElementStructType() StructType
	ElementTextureReferenceType() TextureReferenceType
	ArgumentIndexStride() uint
	ArrayLength() uint
	Stride() uint
	ElementType() DataType
}

// A description of an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlarraytype?language=objc
type ArrayType struct {
	Type
}

func ArrayTypeFrom(ptr unsafe.Pointer) ArrayType {
	return ArrayType{
		Type: TypeFrom(ptr),
	}
}

func (ac _ArrayTypeClass) Alloc() ArrayType {
	rv := objc.Call[ArrayType](ac, objc.Sel("alloc"))
	return rv
}

func ArrayType_Alloc() ArrayType {
	return ArrayTypeClass.Alloc()
}

func (ac _ArrayTypeClass) New() ArrayType {
	rv := objc.Call[ArrayType](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewArrayType() ArrayType {
	return ArrayTypeClass.New()
}

func (a_ ArrayType) Init() ArrayType {
	rv := objc.Call[ArrayType](a_, objc.Sel("init"))
	return rv
}

// Provides a description of the underlying pointer type when an array holds pointers as its elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlarraytype/2915749-elementpointertype?language=objc
func (a_ ArrayType) ElementPointerType() PointerType {
	rv := objc.Call[PointerType](a_, objc.Sel("elementPointerType"))
	return rv
}

// Provides a description of the underlying type when an array holds other arrays as its elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlarraytype/1461963-elementarraytype?language=objc
func (a_ ArrayType) ElementArrayType() ArrayType {
	rv := objc.Call[ArrayType](a_, objc.Sel("elementArrayType"))
	return rv
}

// Provides a description of the underlying struct type when an array holds structs as its elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlarraytype/1461901-elementstructtype?language=objc
func (a_ ArrayType) ElementStructType() StructType {
	rv := objc.Call[StructType](a_, objc.Sel("elementStructType"))
	return rv
}

// Provides a description of the underlying texture type when an array holds textures as its elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlarraytype/2915750-elementtexturereferencetype?language=objc
func (a_ ArrayType) ElementTextureReferenceType() TextureReferenceType {
	rv := objc.Call[TextureReferenceType](a_, objc.Sel("elementTextureReferenceType"))
	return rv
}

// The stride, in bytes, between argument indices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlarraytype/2915752-argumentindexstride?language=objc
func (a_ ArrayType) ArgumentIndexStride() uint {
	rv := objc.Call[uint](a_, objc.Sel("argumentIndexStride"))
	return rv
}

// The number of elements in the array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlarraytype/1462043-arraylength?language=objc
func (a_ ArrayType) ArrayLength() uint {
	rv := objc.Call[uint](a_, objc.Sel("arrayLength"))
	return rv
}

// The stride between array elements, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlarraytype/1461926-stride?language=objc
func (a_ ArrayType) Stride() uint {
	rv := objc.Call[uint](a_, objc.Sel("stride"))
	return rv
}

// The data type of the array’s elements. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlarraytype/1462012-elementtype?language=objc
func (a_ ArrayType) ElementType() DataType {
	rv := objc.Call[DataType](a_, objc.Sel("elementType"))
	return rv
}
