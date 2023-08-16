// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [BufferLayoutDescriptorArray] class.
var BufferLayoutDescriptorArrayClass = _BufferLayoutDescriptorArrayClass{objc.GetClass("MTLBufferLayoutDescriptorArray")}

type _BufferLayoutDescriptorArrayClass struct {
	objc.Class
}

// An interface definition for the [BufferLayoutDescriptorArray] class.
type IBufferLayoutDescriptorArray interface {
	objc.IObject
	ObjectAtIndexedSubscript(index uint) BufferLayoutDescriptor
	SetObjectAtIndexedSubscript(bufferDesc IBufferLayoutDescriptor, index uint)
}

// An array of buffer layout descriptor objects. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutdescriptorarray?language=objc
type BufferLayoutDescriptorArray struct {
	objc.Object
}

func BufferLayoutDescriptorArrayFrom(ptr unsafe.Pointer) BufferLayoutDescriptorArray {
	return BufferLayoutDescriptorArray{
		Object: objc.ObjectFrom(ptr),
	}
}

func (bc _BufferLayoutDescriptorArrayClass) Alloc() BufferLayoutDescriptorArray {
	rv := objc.Call[BufferLayoutDescriptorArray](bc, objc.Sel("alloc"))
	return rv
}

func BufferLayoutDescriptorArray_Alloc() BufferLayoutDescriptorArray {
	return BufferLayoutDescriptorArrayClass.Alloc()
}

func (bc _BufferLayoutDescriptorArrayClass) New() BufferLayoutDescriptorArray {
	rv := objc.Call[BufferLayoutDescriptorArray](bc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewBufferLayoutDescriptorArray() BufferLayoutDescriptorArray {
	return BufferLayoutDescriptorArrayClass.New()
}

func (b_ BufferLayoutDescriptorArray) Init() BufferLayoutDescriptorArray {
	rv := objc.Call[BufferLayoutDescriptorArray](b_, objc.Sel("init"))
	return rv
}

// Returns the state of the specified buffer layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutdescriptorarray/2097228-objectatindexedsubscript?language=objc
func (b_ BufferLayoutDescriptorArray) ObjectAtIndexedSubscript(index uint) BufferLayoutDescriptor {
	rv := objc.Call[BufferLayoutDescriptor](b_, objc.Sel("objectAtIndexedSubscript:"), index)
	return rv
}

// Sets the state of the specified buffer layout. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlbufferlayoutdescriptorarray/2097295-setobject?language=objc
func (b_ BufferLayoutDescriptorArray) SetObjectAtIndexedSubscript(bufferDesc IBufferLayoutDescriptor, index uint) {
	objc.Call[objc.Void](b_, objc.Sel("setObject:atIndexedSubscript:"), objc.Ptr(bufferDesc), index)
}
