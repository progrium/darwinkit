// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VertexAttributeDescriptorArray] class.
var VertexAttributeDescriptorArrayClass = _VertexAttributeDescriptorArrayClass{objc.GetClass("MTLVertexAttributeDescriptorArray")}

type _VertexAttributeDescriptorArrayClass struct {
	objc.Class
}

// An interface definition for the [VertexAttributeDescriptorArray] class.
type IVertexAttributeDescriptorArray interface {
	objc.IObject
	ObjectAtIndexedSubscript(index uint) VertexAttributeDescriptor
	SetObjectAtIndexedSubscript(attributeDesc IVertexAttributeDescriptor, index uint)
}

// An array of vertex attribute descriptor objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattributedescriptorarray?language=objc
type VertexAttributeDescriptorArray struct {
	objc.Object
}

func VertexAttributeDescriptorArrayFrom(ptr unsafe.Pointer) VertexAttributeDescriptorArray {
	return VertexAttributeDescriptorArray{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _VertexAttributeDescriptorArrayClass) Alloc() VertexAttributeDescriptorArray {
	rv := objc.Call[VertexAttributeDescriptorArray](vc, objc.Sel("alloc"))
	return rv
}

func VertexAttributeDescriptorArray_Alloc() VertexAttributeDescriptorArray {
	return VertexAttributeDescriptorArrayClass.Alloc()
}

func (vc _VertexAttributeDescriptorArrayClass) New() VertexAttributeDescriptorArray {
	rv := objc.Call[VertexAttributeDescriptorArray](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVertexAttributeDescriptorArray() VertexAttributeDescriptorArray {
	return VertexAttributeDescriptorArrayClass.New()
}

func (v_ VertexAttributeDescriptorArray) Init() VertexAttributeDescriptorArray {
	rv := objc.Call[VertexAttributeDescriptorArray](v_, objc.Sel("init"))
	return rv
}

// Returns the state of the specified vertex attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattributedescriptorarray/1516138-objectatindexedsubscript?language=objc
func (v_ VertexAttributeDescriptorArray) ObjectAtIndexedSubscript(index uint) VertexAttributeDescriptor {
	rv := objc.Call[VertexAttributeDescriptor](v_, objc.Sel("objectAtIndexedSubscript:"), index)
	return rv
}

// Sets state for the specified vertex attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattributedescriptorarray/1550757-setobject?language=objc
func (v_ VertexAttributeDescriptorArray) SetObjectAtIndexedSubscript(attributeDesc IVertexAttributeDescriptor, index uint) {
	objc.Call[objc.Void](v_, objc.Sel("setObject:atIndexedSubscript:"), objc.Ptr(attributeDesc), index)
}
