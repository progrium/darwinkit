// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VertexAttributeDescriptor] class.
var VertexAttributeDescriptorClass = _VertexAttributeDescriptorClass{objc.GetClass("MTLVertexAttributeDescriptor")}

type _VertexAttributeDescriptorClass struct {
	objc.Class
}

// An interface definition for the [VertexAttributeDescriptor] class.
type IVertexAttributeDescriptor interface {
	objc.IObject
	BufferIndex() uint
	SetBufferIndex(value uint)
	Offset() uint
	SetOffset(value uint)
	Format() VertexFormat
	SetFormat(value VertexFormat)
}

// An object that determines how to store attribute data in memory and map it to the arguments of a vertex function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattributedescriptor?language=objc
type VertexAttributeDescriptor struct {
	objc.Object
}

func VertexAttributeDescriptorFrom(ptr unsafe.Pointer) VertexAttributeDescriptor {
	return VertexAttributeDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _VertexAttributeDescriptorClass) Alloc() VertexAttributeDescriptor {
	rv := objc.Call[VertexAttributeDescriptor](vc, objc.Sel("alloc"))
	return rv
}

func VertexAttributeDescriptor_Alloc() VertexAttributeDescriptor {
	return VertexAttributeDescriptorClass.Alloc()
}

func (vc _VertexAttributeDescriptorClass) New() VertexAttributeDescriptor {
	rv := objc.Call[VertexAttributeDescriptor](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVertexAttributeDescriptor() VertexAttributeDescriptor {
	return VertexAttributeDescriptorClass.New()
}

func (v_ VertexAttributeDescriptor) Init() VertexAttributeDescriptor {
	rv := objc.Call[VertexAttributeDescriptor](v_, objc.Sel("init"))
	return rv
}

// The index in the argument table for the associated vertex buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattributedescriptor/1515502-bufferindex?language=objc
func (v_ VertexAttributeDescriptor) BufferIndex() uint {
	rv := objc.Call[uint](v_, objc.Sel("bufferIndex"))
	return rv
}

// The index in the argument table for the associated vertex buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattributedescriptor/1515502-bufferindex?language=objc
func (v_ VertexAttributeDescriptor) SetBufferIndex(value uint) {
	objc.Call[objc.Void](v_, objc.Sel("setBufferIndex:"), value)
}

// The location of an attribute in vertex data, determined by the byte offset from the start of the vertex data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattributedescriptor/1515785-offset?language=objc
func (v_ VertexAttributeDescriptor) Offset() uint {
	rv := objc.Call[uint](v_, objc.Sel("offset"))
	return rv
}

// The location of an attribute in vertex data, determined by the byte offset from the start of the vertex data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattributedescriptor/1515785-offset?language=objc
func (v_ VertexAttributeDescriptor) SetOffset(value uint) {
	objc.Call[objc.Void](v_, objc.Sel("setOffset:"), value)
}

// The format of the vertex attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattributedescriptor/1516081-format?language=objc
func (v_ VertexAttributeDescriptor) Format() VertexFormat {
	rv := objc.Call[VertexFormat](v_, objc.Sel("format"))
	return rv
}

// The format of the vertex attribute. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexattributedescriptor/1516081-format?language=objc
func (v_ VertexAttributeDescriptor) SetFormat(value VertexFormat) {
	objc.Call[objc.Void](v_, objc.Sel("setFormat:"), value)
}
