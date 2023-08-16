// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [StructType] class.
var StructTypeClass = _StructTypeClass{objc.GetClass("MTLStructType")}

type _StructTypeClass struct {
	objc.Class
}

// An interface definition for the [StructType] class.
type IStructType interface {
	IType
	MemberByName(name string) StructMember
	Members() []StructMember
}

// A description of a structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstructtype?language=objc
type StructType struct {
	Type
}

func StructTypeFrom(ptr unsafe.Pointer) StructType {
	return StructType{
		Type: TypeFrom(ptr),
	}
}

func (sc _StructTypeClass) Alloc() StructType {
	rv := objc.Call[StructType](sc, objc.Sel("alloc"))
	return rv
}

func StructType_Alloc() StructType {
	return StructTypeClass.Alloc()
}

func (sc _StructTypeClass) New() StructType {
	rv := objc.Call[StructType](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewStructType() StructType {
	return StructTypeClass.New()
}

func (s_ StructType) Init() StructType {
	rv := objc.Call[StructType](s_, objc.Sel("init"))
	return rv
}

// Provides a representation of a struct member. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstructtype/1462001-memberbyname?language=objc
func (s_ StructType) MemberByName(name string) StructMember {
	rv := objc.Call[StructMember](s_, objc.Sel("memberByName:"), name)
	return rv
}

// An array of objects that describe the fields in the struct. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlstructtype/1461975-members?language=objc
func (s_ StructType) Members() []StructMember {
	rv := objc.Call[[]StructMember](s_, objc.Sel("members"))
	return rv
}
