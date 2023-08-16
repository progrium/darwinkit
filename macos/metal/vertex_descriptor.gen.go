// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VertexDescriptor] class.
var VertexDescriptorClass = _VertexDescriptorClass{objc.GetClass("MTLVertexDescriptor")}

type _VertexDescriptorClass struct {
	objc.Class
}

// An interface definition for the [VertexDescriptor] class.
type IVertexDescriptor interface {
	objc.IObject
	Reset()
	Layouts() VertexBufferLayoutDescriptorArray
	Attributes() VertexAttributeDescriptorArray
}

// An object that describes how to organize and map data to a vertex function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexdescriptor?language=objc
type VertexDescriptor struct {
	objc.Object
}

func VertexDescriptorFrom(ptr unsafe.Pointer) VertexDescriptor {
	return VertexDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _VertexDescriptorClass) Alloc() VertexDescriptor {
	rv := objc.Call[VertexDescriptor](vc, objc.Sel("alloc"))
	return rv
}

func VertexDescriptor_Alloc() VertexDescriptor {
	return VertexDescriptorClass.Alloc()
}

func (vc _VertexDescriptorClass) New() VertexDescriptor {
	rv := objc.Call[VertexDescriptor](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVertexDescriptor() VertexDescriptor {
	return VertexDescriptorClass.New()
}

func (v_ VertexDescriptor) Init() VertexDescriptor {
	rv := objc.Call[VertexDescriptor](v_, objc.Sel("init"))
	return rv
}

// Creates and returns a new vertex descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexdescriptor/1550756-vertexdescriptor?language=objc
func (vc _VertexDescriptorClass) VertexDescriptor() VertexDescriptor {
	rv := objc.Call[VertexDescriptor](vc, objc.Sel("vertexDescriptor"))
	return rv
}

// Creates and returns a new vertex descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexdescriptor/1550756-vertexdescriptor?language=objc
func VertexDescriptor_VertexDescriptor() VertexDescriptor {
	return VertexDescriptorClass.VertexDescriptor()
}

// Resets the default state for the vertex descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexdescriptor/1515988-reset?language=objc
func (v_ VertexDescriptor) Reset() {
	objc.Call[objc.Void](v_, objc.Sel("reset"))
}

// An array of state data that describes how data are fetched by a vertex shader function when rendering primitives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexdescriptor/1515480-layouts?language=objc
func (v_ VertexDescriptor) Layouts() VertexBufferLayoutDescriptorArray {
	rv := objc.Call[VertexBufferLayoutDescriptorArray](v_, objc.Sel("layouts"))
	return rv
}

// An array of state data that describes how vertex attribute data is stored in memory and is mapped to arguments for a vertex shader function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexdescriptor/1515921-attributes?language=objc
func (v_ VertexDescriptor) Attributes() VertexAttributeDescriptorArray {
	rv := objc.Call[VertexAttributeDescriptorArray](v_, objc.Sel("attributes"))
	return rv
}
