// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// A render command in an indirect command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectrendercommand?language=objc
type PIndirectRenderCommand interface {
	// optional
	DrawIndexedPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetInstanceCountBaseInstanceTessellationFactorBufferTessellationFactorBufferOffsetTessellationFactorBufferInstanceStride(numberOfPatchControlPoints uint, patchStart uint, patchCount uint, patchIndexBuffer BufferWrapper, patchIndexBufferOffset uint, controlPointIndexBuffer BufferWrapper, controlPointIndexBufferOffset uint, instanceCount uint, baseInstance uint, buffer BufferWrapper, offset uint, instanceStride uint)
	HasDrawIndexedPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetInstanceCountBaseInstanceTessellationFactorBufferTessellationFactorBufferOffsetTessellationFactorBufferInstanceStride() bool

	// optional
	SetVertexBufferOffsetAtIndex(buffer BufferWrapper, offset uint, index uint)
	HasSetVertexBufferOffsetAtIndex() bool

	// optional
	DrawPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetInstanceCountBaseInstanceTessellationFactorBufferTessellationFactorBufferOffsetTessellationFactorBufferInstanceStride(numberOfPatchControlPoints uint, patchStart uint, patchCount uint, patchIndexBuffer BufferWrapper, patchIndexBufferOffset uint, instanceCount uint, baseInstance uint, buffer BufferWrapper, offset uint, instanceStride uint)
	HasDrawPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetInstanceCountBaseInstanceTessellationFactorBufferTessellationFactorBufferOffsetTessellationFactorBufferInstanceStride() bool

	// optional
	DrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance(primitiveType PrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint, baseInstance uint)
	HasDrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance() bool

	// optional
	SetFragmentBufferOffsetAtIndex(buffer BufferWrapper, offset uint, index uint)
	HasSetFragmentBufferOffsetAtIndex() bool

	// optional
	SetRenderPipelineState(pipelineState RenderPipelineStateWrapper)
	HasSetRenderPipelineState() bool

	// optional
	Reset()
	HasReset() bool

	// optional
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCountBaseVertexBaseInstance(primitiveType PrimitiveType, indexCount uint, indexType IndexType, indexBuffer BufferWrapper, indexBufferOffset uint, instanceCount uint, baseVertex int, baseInstance uint)
	HasDrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCountBaseVertexBaseInstance() bool
}

// A concrete type wrapper for the [PIndirectRenderCommand] protocol.
type IndirectRenderCommandWrapper struct {
	objc.Object
}

func (i_ IndirectRenderCommandWrapper) HasDrawIndexedPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetInstanceCountBaseInstanceTessellationFactorBufferTessellationFactorBufferOffsetTessellationFactorBufferInstanceStride() bool {
	return i_.RespondsToSelector(objc.Sel("drawIndexedPatches:patchStart:patchCount:patchIndexBuffer:patchIndexBufferOffset:controlPointIndexBuffer:controlPointIndexBufferOffset:instanceCount:baseInstance:tessellationFactorBuffer:tessellationFactorBufferOffset:tessellationFactorBufferInstanceStride:"))
}

// Encodes a command to render a number of instances of tessellated patches, using a control point index buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectrendercommand/2967433-drawindexedpatches?language=objc
func (i_ IndirectRenderCommandWrapper) DrawIndexedPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetInstanceCountBaseInstanceTessellationFactorBufferTessellationFactorBufferOffsetTessellationFactorBufferInstanceStride(numberOfPatchControlPoints uint, patchStart uint, patchCount uint, patchIndexBuffer PBuffer, patchIndexBufferOffset uint, controlPointIndexBuffer PBuffer, controlPointIndexBufferOffset uint, instanceCount uint, baseInstance uint, buffer PBuffer, offset uint, instanceStride uint) {
	po3 := objc.WrapAsProtocol("MTLBuffer", patchIndexBuffer)
	po5 := objc.WrapAsProtocol("MTLBuffer", controlPointIndexBuffer)
	po9 := objc.WrapAsProtocol("MTLBuffer", buffer)
	objc.Call[objc.Void](i_, objc.Sel("drawIndexedPatches:patchStart:patchCount:patchIndexBuffer:patchIndexBufferOffset:controlPointIndexBuffer:controlPointIndexBufferOffset:instanceCount:baseInstance:tessellationFactorBuffer:tessellationFactorBufferOffset:tessellationFactorBufferInstanceStride:"), numberOfPatchControlPoints, patchStart, patchCount, po3, patchIndexBufferOffset, po5, controlPointIndexBufferOffset, instanceCount, baseInstance, po9, offset, instanceStride)
}

func (i_ IndirectRenderCommandWrapper) HasSetVertexBufferOffsetAtIndex() bool {
	return i_.RespondsToSelector(objc.Sel("setVertexBuffer:offset:atIndex:"))
}

// Sets a vertex buffer argument for the command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectrendercommand/2967438-setvertexbuffer?language=objc
func (i_ IndirectRenderCommandWrapper) SetVertexBufferOffsetAtIndex(buffer PBuffer, offset uint, index uint) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	objc.Call[objc.Void](i_, objc.Sel("setVertexBuffer:offset:atIndex:"), po0, offset, index)
}

func (i_ IndirectRenderCommandWrapper) HasDrawPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetInstanceCountBaseInstanceTessellationFactorBufferTessellationFactorBufferOffsetTessellationFactorBufferInstanceStride() bool {
	return i_.RespondsToSelector(objc.Sel("drawPatches:patchStart:patchCount:patchIndexBuffer:patchIndexBufferOffset:instanceCount:baseInstance:tessellationFactorBuffer:tessellationFactorBufferOffset:tessellationFactorBufferInstanceStride:"))
}

// Encodes a command to render a number of instances of tessellated patches. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectrendercommand/2967435-drawpatches?language=objc
func (i_ IndirectRenderCommandWrapper) DrawPatchesPatchStartPatchCountPatchIndexBufferPatchIndexBufferOffsetInstanceCountBaseInstanceTessellationFactorBufferTessellationFactorBufferOffsetTessellationFactorBufferInstanceStride(numberOfPatchControlPoints uint, patchStart uint, patchCount uint, patchIndexBuffer PBuffer, patchIndexBufferOffset uint, instanceCount uint, baseInstance uint, buffer PBuffer, offset uint, instanceStride uint) {
	po3 := objc.WrapAsProtocol("MTLBuffer", patchIndexBuffer)
	po7 := objc.WrapAsProtocol("MTLBuffer", buffer)
	objc.Call[objc.Void](i_, objc.Sel("drawPatches:patchStart:patchCount:patchIndexBuffer:patchIndexBufferOffset:instanceCount:baseInstance:tessellationFactorBuffer:tessellationFactorBufferOffset:tessellationFactorBufferInstanceStride:"), numberOfPatchControlPoints, patchStart, patchCount, po3, patchIndexBufferOffset, instanceCount, baseInstance, po7, offset, instanceStride)
}

func (i_ IndirectRenderCommandWrapper) HasDrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance() bool {
	return i_.RespondsToSelector(objc.Sel("drawPrimitives:vertexStart:vertexCount:instanceCount:baseInstance:"))
}

// Encodes a command to render a number of instances of primitives using vertex data in contiguous array elements, starting from the base instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectrendercommand/2966620-drawprimitives?language=objc
func (i_ IndirectRenderCommandWrapper) DrawPrimitivesVertexStartVertexCountInstanceCountBaseInstance(primitiveType PrimitiveType, vertexStart uint, vertexCount uint, instanceCount uint, baseInstance uint) {
	objc.Call[objc.Void](i_, objc.Sel("drawPrimitives:vertexStart:vertexCount:instanceCount:baseInstance:"), primitiveType, vertexStart, vertexCount, instanceCount, baseInstance)
}

func (i_ IndirectRenderCommandWrapper) HasSetFragmentBufferOffsetAtIndex() bool {
	return i_.RespondsToSelector(objc.Sel("setFragmentBuffer:offset:atIndex:"))
}

// Sets a fragment buffer argument for the command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectrendercommand/2967437-setfragmentbuffer?language=objc
func (i_ IndirectRenderCommandWrapper) SetFragmentBufferOffsetAtIndex(buffer PBuffer, offset uint, index uint) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	objc.Call[objc.Void](i_, objc.Sel("setFragmentBuffer:offset:atIndex:"), po0, offset, index)
}

func (i_ IndirectRenderCommandWrapper) HasSetRenderPipelineState() bool {
	return i_.RespondsToSelector(objc.Sel("setRenderPipelineState:"))
}

// Sets the render pipeline state object used by the command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectrendercommand/2966625-setrenderpipelinestate?language=objc
func (i_ IndirectRenderCommandWrapper) SetRenderPipelineState(pipelineState PRenderPipelineState) {
	po0 := objc.WrapAsProtocol("MTLRenderPipelineState", pipelineState)
	objc.Call[objc.Void](i_, objc.Sel("setRenderPipelineState:"), po0)
}

func (i_ IndirectRenderCommandWrapper) HasReset() bool {
	return i_.RespondsToSelector(objc.Sel("reset"))
}

// Resets the command to its default state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectrendercommand/2981030-reset?language=objc
func (i_ IndirectRenderCommandWrapper) Reset() {
	objc.Call[objc.Void](i_, objc.Sel("reset"))
}

func (i_ IndirectRenderCommandWrapper) HasDrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCountBaseVertexBaseInstance() bool {
	return i_.RespondsToSelector(objc.Sel("drawIndexedPrimitives:indexCount:indexType:indexBuffer:indexBufferOffset:instanceCount:baseVertex:baseInstance:"))
}

// Encodes a command to render a number of instances of primitives using an index list specified in a buffer, starting from the base vertex of the base instance. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectrendercommand/2966618-drawindexedprimitives?language=objc
func (i_ IndirectRenderCommandWrapper) DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCountBaseVertexBaseInstance(primitiveType PrimitiveType, indexCount uint, indexType IndexType, indexBuffer PBuffer, indexBufferOffset uint, instanceCount uint, baseVertex int, baseInstance uint) {
	po3 := objc.WrapAsProtocol("MTLBuffer", indexBuffer)
	objc.Call[objc.Void](i_, objc.Sel("drawIndexedPrimitives:indexCount:indexType:indexBuffer:indexBufferOffset:instanceCount:baseVertex:baseInstance:"), primitiveType, indexCount, indexType, po3, indexBufferOffset, instanceCount, baseVertex, baseInstance)
}
