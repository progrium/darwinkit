// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VertexAttribute] class.
var VertexAttributeClass = _VertexAttributeClass{objc.GetClass("MTLVertexAttribute")}

type _VertexAttributeClass struct {
	objc.Class
}

// An interface definition for the [VertexAttribute] class.
type IVertexAttribute interface {
	objc.IObject
	Name() string
	IsActive() bool
	IsPatchControlPointData() bool
	IsPatchData() bool
	AttributeType() DataType
	AttributeIndex() uint
}

// An object that represents an attribute of a vertex function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute?language=objc
type VertexAttribute struct {
	objc.Object
}

func VertexAttributeFrom(ptr unsafe.Pointer) VertexAttribute {
	return VertexAttribute{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _VertexAttributeClass) Alloc() VertexAttribute {
	rv := objc.Call[VertexAttribute](vc, objc.Sel("alloc"))
	return rv
}

func VertexAttribute_Alloc() VertexAttribute {
	return VertexAttributeClass.Alloc()
}

func (vc _VertexAttributeClass) New() VertexAttribute {
	rv := objc.Call[VertexAttribute](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVertexAttribute() VertexAttribute {
	return VertexAttributeClass.New()
}

func (v_ VertexAttribute) Init() VertexAttribute {
	rv := objc.Call[VertexAttribute](v_, objc.Sel("init"))
	return rv
}

// The name of the attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/1515447-name?language=objc
func (v_ VertexAttribute) Name() string {
	rv := objc.Call[string](v_, objc.Sel("name"))
	return rv
}

// A Boolean value that indicates whether this vertex attribute is active. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/1515994-active?language=objc
func (v_ VertexAttribute) IsActive() bool {
	rv := objc.Call[bool](v_, objc.Sel("isActive"))
	return rv
}

// A Boolean value that indicates whether this vertex attribute represents control point data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/1640013-patchcontrolpointdata?language=objc
func (v_ VertexAttribute) IsPatchControlPointData() bool {
	rv := objc.Call[bool](v_, objc.Sel("isPatchControlPointData"))
	return rv
}

// A Boolean value that indicates whether this vertex attribute represents patch data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/1640002-patchdata?language=objc
func (v_ VertexAttribute) IsPatchData() bool {
	rv := objc.Call[bool](v_, objc.Sel("isPatchData"))
	return rv
}

// The data type for the attribute, as declared in Metal shader source code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/1516002-attributetype?language=objc
func (v_ VertexAttribute) AttributeType() DataType {
	rv := objc.Call[DataType](v_, objc.Sel("attributeType"))
	return rv
}

// The index of the attribute, as declared in Metal shader source code. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattribute/1516285-attributeindex?language=objc
func (v_ VertexAttribute) AttributeIndex() uint {
	rv := objc.Call[uint](v_, objc.Sel("attributeIndex"))
	return rv
}
