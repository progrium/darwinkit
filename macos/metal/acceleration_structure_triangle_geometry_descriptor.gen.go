// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [AccelerationStructureTriangleGeometryDescriptor] class.
var AccelerationStructureTriangleGeometryDescriptorClass = _AccelerationStructureTriangleGeometryDescriptorClass{objc.GetClass("MTLAccelerationStructureTriangleGeometryDescriptor")}

type _AccelerationStructureTriangleGeometryDescriptorClass struct {
	objc.Class
}

// An interface definition for the [AccelerationStructureTriangleGeometryDescriptor] class.
type IAccelerationStructureTriangleGeometryDescriptor interface {
	IAccelerationStructureGeometryDescriptor
	VertexBuffer() BufferWrapper
	SetVertexBuffer(value PBuffer)
	SetVertexBufferObject(valueObject objc.IObject)
	TriangleCount() uint
	SetTriangleCount(value uint)
	IndexBuffer() BufferWrapper
	SetIndexBuffer(value PBuffer)
	SetIndexBufferObject(valueObject objc.IObject)
	IndexBufferOffset() uint
	SetIndexBufferOffset(value uint)
	IndexType() IndexType
	SetIndexType(value IndexType)
	VertexStride() uint
	SetVertexStride(value uint)
	VertexBufferOffset() uint
	SetVertexBufferOffset(value uint)
}

// A description of a list of triangle primitives to turn into an acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor?language=objc
type AccelerationStructureTriangleGeometryDescriptor struct {
	AccelerationStructureGeometryDescriptor
}

func AccelerationStructureTriangleGeometryDescriptorFrom(ptr unsafe.Pointer) AccelerationStructureTriangleGeometryDescriptor {
	return AccelerationStructureTriangleGeometryDescriptor{
		AccelerationStructureGeometryDescriptor: AccelerationStructureGeometryDescriptorFrom(ptr),
	}
}

func (ac _AccelerationStructureTriangleGeometryDescriptorClass) Descriptor() AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Call[AccelerationStructureTriangleGeometryDescriptor](ac, objc.Sel("descriptor"))
	return rv
}

// Creates a new triangle descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/3553869-descriptor?language=objc
func AccelerationStructureTriangleGeometryDescriptor_Descriptor() AccelerationStructureTriangleGeometryDescriptor {
	return AccelerationStructureTriangleGeometryDescriptorClass.Descriptor()
}

func (ac _AccelerationStructureTriangleGeometryDescriptorClass) Alloc() AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Call[AccelerationStructureTriangleGeometryDescriptor](ac, objc.Sel("alloc"))
	return rv
}

func AccelerationStructureTriangleGeometryDescriptor_Alloc() AccelerationStructureTriangleGeometryDescriptor {
	return AccelerationStructureTriangleGeometryDescriptorClass.Alloc()
}

func (ac _AccelerationStructureTriangleGeometryDescriptorClass) New() AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Call[AccelerationStructureTriangleGeometryDescriptor](ac, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewAccelerationStructureTriangleGeometryDescriptor() AccelerationStructureTriangleGeometryDescriptor {
	return AccelerationStructureTriangleGeometryDescriptorClass.New()
}

func (a_ AccelerationStructureTriangleGeometryDescriptor) Init() AccelerationStructureTriangleGeometryDescriptor {
	rv := objc.Call[AccelerationStructureTriangleGeometryDescriptor](a_, objc.Sel("init"))
	return rv
}

// A buffer  that contains vertex data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/3553876-vertexbuffer?language=objc
func (a_ AccelerationStructureTriangleGeometryDescriptor) VertexBuffer() BufferWrapper {
	rv := objc.Call[BufferWrapper](a_, objc.Sel("vertexBuffer"))
	return rv
}

// A buffer  that contains vertex data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/3553876-vertexbuffer?language=objc
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetVertexBuffer(value PBuffer) {
	po0 := objc.WrapAsProtocol("MTLBuffer", value)
	objc.Call[objc.Void](a_, objc.Sel("setVertexBuffer:"), po0)
}

// A buffer  that contains vertex data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/3553876-vertexbuffer?language=objc
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetVertexBufferObject(valueObject objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setVertexBuffer:"), objc.Ptr(valueObject))
}

// The number of triangles in the buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/3553875-trianglecount?language=objc
func (a_ AccelerationStructureTriangleGeometryDescriptor) TriangleCount() uint {
	rv := objc.Call[uint](a_, objc.Sel("triangleCount"))
	return rv
}

// The number of triangles in the buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/3553875-trianglecount?language=objc
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetTriangleCount(value uint) {
	objc.Call[objc.Void](a_, objc.Sel("setTriangleCount:"), value)
}

// A buffer that contains indices for the vertices that compose the triangle list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/3553870-indexbuffer?language=objc
func (a_ AccelerationStructureTriangleGeometryDescriptor) IndexBuffer() BufferWrapper {
	rv := objc.Call[BufferWrapper](a_, objc.Sel("indexBuffer"))
	return rv
}

// A buffer that contains indices for the vertices that compose the triangle list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/3553870-indexbuffer?language=objc
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetIndexBuffer(value PBuffer) {
	po0 := objc.WrapAsProtocol("MTLBuffer", value)
	objc.Call[objc.Void](a_, objc.Sel("setIndexBuffer:"), po0)
}

// A buffer that contains indices for the vertices that compose the triangle list. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/3553870-indexbuffer?language=objc
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetIndexBufferObject(valueObject objc.IObject) {
	objc.Call[objc.Void](a_, objc.Sel("setIndexBuffer:"), objc.Ptr(valueObject))
}

// The offset, in bytes, to the first index in the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/3553871-indexbufferoffset?language=objc
func (a_ AccelerationStructureTriangleGeometryDescriptor) IndexBufferOffset() uint {
	rv := objc.Call[uint](a_, objc.Sel("indexBufferOffset"))
	return rv
}

// The offset, in bytes, to the first index in the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/3553871-indexbufferoffset?language=objc
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetIndexBufferOffset(value uint) {
	objc.Call[objc.Void](a_, objc.Sel("setIndexBufferOffset:"), value)
}

// The data type of indices in the index buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/3666587-indextype?language=objc
func (a_ AccelerationStructureTriangleGeometryDescriptor) IndexType() IndexType {
	rv := objc.Call[IndexType](a_, objc.Sel("indexType"))
	return rv
}

// The data type of indices in the index buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/3666587-indextype?language=objc
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetIndexType(value IndexType) {
	objc.Call[objc.Void](a_, objc.Sel("setIndexType:"), value)
}

// The stride, in bytes, between vertices in the vertex buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/3553878-vertexstride?language=objc
func (a_ AccelerationStructureTriangleGeometryDescriptor) VertexStride() uint {
	rv := objc.Call[uint](a_, objc.Sel("vertexStride"))
	return rv
}

// The stride, in bytes, between vertices in the vertex buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/3553878-vertexstride?language=objc
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetVertexStride(value uint) {
	objc.Call[objc.Void](a_, objc.Sel("setVertexStride:"), value)
}

// The offset, in bytes, for the first vertex in the vertex buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/3553877-vertexbufferoffset?language=objc
func (a_ AccelerationStructureTriangleGeometryDescriptor) VertexBufferOffset() uint {
	rv := objc.Call[uint](a_, objc.Sel("vertexBufferOffset"))
	return rv
}

// The offset, in bytes, for the first vertex in the vertex buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructuretrianglegeometrydescriptor/3553877-vertexbufferoffset?language=objc
func (a_ AccelerationStructureTriangleGeometryDescriptor) SetVertexBufferOffset(value uint) {
	objc.Call[objc.Void](a_, objc.Sel("setVertexBufferOffset:"), value)
}