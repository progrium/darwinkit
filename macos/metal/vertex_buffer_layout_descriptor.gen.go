// Code generated by DarwinKit. DO NOT EDIT.

package metal

import (
	"unsafe"

	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [VertexBufferLayoutDescriptor] class.
var VertexBufferLayoutDescriptorClass = _VertexBufferLayoutDescriptorClass{objc.GetClass("MTLVertexBufferLayoutDescriptor")}

type _VertexBufferLayoutDescriptorClass struct {
	objc.Class
}

// An interface definition for the [VertexBufferLayoutDescriptor] class.
type IVertexBufferLayoutDescriptor interface {
	objc.IObject
	StepFunction() VertexStepFunction
	SetStepFunction(value VertexStepFunction)
	StepRate() uint
	SetStepRate(value uint)
	Stride() uint
	SetStride(value uint)
}

// An object that configures how a render pipeline fetches data to send to the vertex function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexbufferlayoutdescriptor?language=objc
type VertexBufferLayoutDescriptor struct {
	objc.Object
}

func VertexBufferLayoutDescriptorFrom(ptr unsafe.Pointer) VertexBufferLayoutDescriptor {
	return VertexBufferLayoutDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (vc _VertexBufferLayoutDescriptorClass) Alloc() VertexBufferLayoutDescriptor {
	rv := objc.Call[VertexBufferLayoutDescriptor](vc, objc.Sel("alloc"))
	return rv
}

func (vc _VertexBufferLayoutDescriptorClass) New() VertexBufferLayoutDescriptor {
	rv := objc.Call[VertexBufferLayoutDescriptor](vc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewVertexBufferLayoutDescriptor() VertexBufferLayoutDescriptor {
	return VertexBufferLayoutDescriptorClass.New()
}

func (v_ VertexBufferLayoutDescriptor) Init() VertexBufferLayoutDescriptor {
	rv := objc.Call[VertexBufferLayoutDescriptor](v_, objc.Sel("init"))
	return rv
}

// The circumstances under which the vertex and its attributes are presented to the vertex function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexbufferlayoutdescriptor/1515341-stepfunction?language=objc
func (v_ VertexBufferLayoutDescriptor) StepFunction() VertexStepFunction {
	rv := objc.Call[VertexStepFunction](v_, objc.Sel("stepFunction"))
	return rv
}

// The circumstances under which the vertex and its attributes are presented to the vertex function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexbufferlayoutdescriptor/1515341-stepfunction?language=objc
func (v_ VertexBufferLayoutDescriptor) SetStepFunction(value VertexStepFunction) {
	objc.Call[objc.Void](v_, objc.Sel("setStepFunction:"), value)
}

// The interval at which the vertex and its attributes are presented to the vertex function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexbufferlayoutdescriptor/1516148-steprate?language=objc
func (v_ VertexBufferLayoutDescriptor) StepRate() uint {
	rv := objc.Call[uint](v_, objc.Sel("stepRate"))
	return rv
}

// The interval at which the vertex and its attributes are presented to the vertex function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexbufferlayoutdescriptor/1516148-steprate?language=objc
func (v_ VertexBufferLayoutDescriptor) SetStepRate(value uint) {
	objc.Call[objc.Void](v_, objc.Sel("setStepRate:"), value)
}

// The number of bytes between the first byte of two consecutive vertices in a buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexbufferlayoutdescriptor/1515441-stride?language=objc
func (v_ VertexBufferLayoutDescriptor) Stride() uint {
	rv := objc.Call[uint](v_, objc.Sel("stride"))
	return rv
}

// The number of bytes between the first byte of two consecutive vertices in a buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlvertexbufferlayoutdescriptor/1515441-stride?language=objc
func (v_ VertexBufferLayoutDescriptor) SetStride(value uint) {
	objc.Call[objc.Void](v_, objc.Sel("setStride:"), value)
}
