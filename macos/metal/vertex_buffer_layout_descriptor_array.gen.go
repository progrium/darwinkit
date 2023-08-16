// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [VertexBufferLayoutDescriptorArray] class.
var VertexBufferLayoutDescriptorArrayClass = _VertexBufferLayoutDescriptorArrayClass{objc.GetClass("MTLVertexBufferLayoutDescriptorArray")}

type _VertexBufferLayoutDescriptorArrayClass struct {
	objc.Class
}

// An interface definition for the [VertexBufferLayoutDescriptorArray] class.
type IVertexBufferLayoutDescriptorArray interface {
	objc.IObject
	ObjectAtIndexedSubscript(index uint) VertexBufferLayoutDescriptor
	SetObjectAtIndexedSubscript(bufferDesc IVertexBufferLayoutDescriptor, index uint)
}

// An array of vertex buffer layout descriptor objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexbufferlayoutdescriptorarray?language=objc
type VertexBufferLayoutDescriptorArray struct {
	objc.Object
}

func VertexBufferLayoutDescriptorArrayFrom(ptr unsafe.Pointer) VertexBufferLayoutDescriptorArray {
	return VertexBufferLayoutDescriptorArray{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _VertexBufferLayoutDescriptorArrayClass) Alloc() VertexBufferLayoutDescriptorArray {
	rv := objc.Call[VertexBufferLayoutDescriptorArray](vc, objc.Sel("alloc"))
	return rv
}

func VertexBufferLayoutDescriptorArray_Alloc() VertexBufferLayoutDescriptorArray {
	return VertexBufferLayoutDescriptorArrayClass.Alloc()
}

func (vc _VertexBufferLayoutDescriptorArrayClass) New() VertexBufferLayoutDescriptorArray {
	rv := objc.Call[VertexBufferLayoutDescriptorArray](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVertexBufferLayoutDescriptorArray() VertexBufferLayoutDescriptorArray {
	return VertexBufferLayoutDescriptorArrayClass.New()
}

func (v_ VertexBufferLayoutDescriptorArray) Init() VertexBufferLayoutDescriptorArray {
	rv := objc.Call[VertexBufferLayoutDescriptorArray](v_, objc.Sel("init"))
	return rv
}

// Returns the state of the specified vertex buffer layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexbufferlayoutdescriptorarray/1516230-objectatindexedsubscript?language=objc
func (v_ VertexBufferLayoutDescriptorArray) ObjectAtIndexedSubscript(index uint) VertexBufferLayoutDescriptor {
	rv := objc.Call[VertexBufferLayoutDescriptor](v_, objc.Sel("objectAtIndexedSubscript:"), index)
	return rv
}

// Sets the state of the specified vertex buffer layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexbufferlayoutdescriptorarray/1550758-setobject?language=objc
func (v_ VertexBufferLayoutDescriptorArray) SetObjectAtIndexedSubscript(bufferDesc IVertexBufferLayoutDescriptor, index uint) {
	objc.Call[objc.Void](v_, objc.Sel("setObject:atIndexedSubscript:"), objc.Ptr(bufferDesc), index)
}
