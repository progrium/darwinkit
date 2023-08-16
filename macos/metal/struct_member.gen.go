// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [StructMember] class.
var StructMemberClass = _StructMemberClass{objc.GetClass("MTLStructMember")}

type _StructMemberClass struct {
	objc.Class
}

// An interface definition for the [StructMember] class.
type IStructMember interface {
	objc.IObject
	ArrayType() ArrayType
	PointerType() PointerType
	StructType() StructType
	TextureReferenceType() TextureReferenceType
	Name() string
	ArgumentIndex() uint
	Offset() uint
	DataType() DataType
}

// An object that provides information about a field in a structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstructmember?language=objc
type StructMember struct {
	objc.Object
}

func StructMemberFrom(ptr unsafe.Pointer) StructMember {
	return StructMember{
		Object: objc.ObjectFrom(ptr),
	}
}

func (sc _StructMemberClass) Alloc() StructMember {
	rv := objc.Call[StructMember](sc, objc.Sel("alloc"))
	return rv
}

func StructMember_Alloc() StructMember {
	return StructMemberClass.Alloc()
}

func (sc _StructMemberClass) New() StructMember {
	rv := objc.Call[StructMember](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStructMember() StructMember {
	return StructMemberClass.New()
}

func (s_ StructMember) Init() StructMember {
	rv := objc.Call[StructMember](s_, objc.Sel("init"))
	return rv
}

// Provides a description of the underlying array when the struct member holds an array. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstructmember/1461887-arraytype?language=objc
func (s_ StructMember) ArrayType() ArrayType {
	rv := objc.Call[ArrayType](s_, objc.Sel("arrayType"))
	return rv
}

// Provides a description of the underlying pointer when the struct member holds a pointer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstructmember/2915751-pointertype?language=objc
func (s_ StructMember) PointerType() PointerType {
	rv := objc.Call[PointerType](s_, objc.Sel("pointerType"))
	return rv
}

// Provides a description of the underlying struct when the struct member holds a struct. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstructmember/1462011-structtype?language=objc
func (s_ StructMember) StructType() StructType {
	rv := objc.Call[StructType](s_, objc.Sel("structType"))
	return rv
}

// Provides a description of the underlying texture when the struct member holds a texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstructmember/2915753-texturereferencetype?language=objc
func (s_ StructMember) TextureReferenceType() TextureReferenceType {
	rv := objc.Call[TextureReferenceType](s_, objc.Sel("textureReferenceType"))
	return rv
}

// The name of the struct member. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstructmember/1461944-name?language=objc
func (s_ StructMember) Name() string {
	rv := objc.Call[string](s_, objc.Sel("name"))
	return rv
}

// The index in the argument table that corresponds to the struct member. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstructmember/2915754-argumentindex?language=objc
func (s_ StructMember) ArgumentIndex() uint {
	rv := objc.Call[uint](s_, objc.Sel("argumentIndex"))
	return rv
}

// The location of this member relative to the start of its struct, in bytes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstructmember/1461970-offset?language=objc
func (s_ StructMember) Offset() uint {
	rv := objc.Call[uint](s_, objc.Sel("offset"))
	return rv
}

// The data type of the struct member. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstructmember/1461972-datatype?language=objc
func (s_ StructMember) DataType() DataType {
	rv := objc.Call[DataType](s_, objc.Sel("dataType"))
	return rv
}
