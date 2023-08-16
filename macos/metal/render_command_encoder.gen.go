// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// An interface that encodes a render pass into a command buffer, including all its draw calls and configuration. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder?language=objc
type PRenderCommandEncoder interface {
	// optional
	SetViewportsCount(viewports *Viewport, count uint)
	HasSetViewportsCount() bool

	// optional
	SetThreadgroupMemoryLengthOffsetAtIndex(length uint, offset uint, index uint)
	HasSetThreadgroupMemoryLengthOffsetAtIndex() bool

	// optional
	SetFragmentVisibleFunctionTableAtBufferIndex(functionTable VisibleFunctionTableWrapper, bufferIndex uint)
	HasSetFragmentVisibleFunctionTableAtBufferIndex() bool

	// optional
	SetFragmentTexturesWithRange(textures TextureWrapper, range_ foundation.Range)
	HasSetFragmentTexturesWithRange() bool

	// optional
	SetViewport(viewport Viewport)
	HasSetViewport() bool

	// optional
	DrawIndexedPatchesPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetIndirectBufferIndirectBufferOffset(numberOfPatchControlPoints uint, patchIndexBuffer BufferWrapper, patchIndexBufferOffset uint, controlPointIndexBuffer BufferWrapper, controlPointIndexBufferOffset uint, indirectBuffer BufferWrapper, indirectBufferOffset uint)
	HasDrawIndexedPatchesPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetIndirectBufferIndirectBufferOffset() bool

	// optional
	SetTileSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers SamplerStateWrapper, lodMinClamps *float64, lodMaxClamps *float64, range_ foundation.Range)
	HasSetTileSamplerStatesLodMinClampsLodMaxClampsWithRange() bool

	// optional
	SetFragmentTextureAtIndex(texture TextureWrapper, index uint)
	HasSetFragmentTextureAtIndex() bool

	// optional
	SetDepthStoreAction(storeAction StoreAction)
	HasSetDepthStoreAction() bool

	// optional
	SetFragmentSamplerStatesWithRange(samplers SamplerStateWrapper, range_ foundation.Range)
	HasSetFragmentSamplerStatesWithRange() bool

	// optional
	SetVisibilityResultModeOffset(mode VisibilityResultMode, offset uint)
	HasSetVisibilityResultModeOffset() bool

	// optional
	SetTileBufferOffsetAtIndex(offset uint, index uint)
	HasSetTileBufferOffsetAtIndex() bool

	// optional
	SetFragmentBufferOffsetAtIndex(offset uint, index uint)
	HasSetFragmentBufferOffsetAtIndex() bool

	// optional
	SetVertexBufferOffsetAtIndex_(buffer BufferWrapper, offset uint, index uint)
	HasSetVertexBufferOffsetAtIndex_() bool

	// optional
	SetVertexIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable IntersectionFunctionTableWrapper, bufferIndex uint)
	HasSetVertexIntersectionFunctionTableAtBufferIndex() bool

	// optional
	SetTileBufferOffsetAtIndex_(buffer BufferWrapper, offset uint, index uint)
	HasSetTileBufferOffsetAtIndex_() bool

	// optional
	SetTessellationFactorBufferOffsetInstanceStride(buffer BufferWrapper, offset uint, instanceStride uint)
	HasSetTessellationFactorBufferOffsetInstanceStride() bool

	// optional
	SetTileVisibleFunctionTableAtBufferIndex(functionTable VisibleFunctionTableWrapper, bufferIndex uint)
	HasSetTileVisibleFunctionTableAtBufferIndex() bool

	// optional
	DrawPatchesPatchIndexBufferPatchIndexBufferOffsetIndirectBufferIndirectBufferOffset(numberOfPatchControlPoints uint, patchIndexBuffer BufferWrapper, patchIndexBufferOffset uint, indirectBuffer BufferWrapper, indirectBufferOffset uint)
	HasDrawPatchesPatchIndexBufferPatchIndexBufferOffsetIndirectBufferIndirectBufferOffset() bool

	// optional
	UseResourceUsageStages(resource ResourceWrapper, usage ResourceUsage, stages RenderStages)
	HasUseResourceUsageStages() bool

	// optional
	DrawPrimitivesVertexStartVertexCount(primitiveType PrimitiveType, vertexStart uint, vertexCount uint)
	HasDrawPrimitivesVertexStartVertexCount() bool

	// optional
	SetVertexBuffersOffsetsWithRange(buffers BufferWrapper, offsets *uint, range_ foundation.Range)
	HasSetVertexBuffersOffsetsWithRange() bool

	// optional
	SetFragmentBufferOffsetAtIndex_(buffer BufferWrapper, offset uint, index uint)
	HasSetFragmentBufferOffsetAtIndex_() bool

	// optional
	SetDepthStencilState(depthStencilState DepthStencilStateWrapper)
	HasSetDepthStencilState() bool

	// optional
	SetVertexAccelerationStructureAtBufferIndex(accelerationStructure AccelerationStructureWrapper, bufferIndex uint)
	HasSetVertexAccelerationStructureAtBufferIndex() bool

	// optional
	SetVertexBytesLengthAtIndex(bytes unsafe.Pointer, length uint, index uint)
	HasSetVertexBytesLengthAtIndex() bool

	// optional
	SetVertexAmplificationCountViewMappings(count uint, viewMappings *VertexAmplificationViewMapping)
	HasSetVertexAmplificationCountViewMappings() bool

	// optional
	SetScissorRect(rect ScissorRect)
	HasSetScissorRect() bool

	// optional
	SetTileTextureAtIndex(texture TextureWrapper, index uint)
	HasSetTileTextureAtIndex() bool

	// optional
	SetVertexSamplerStateLodMinClampLodMaxClampAtIndex(sampler SamplerStateWrapper, lodMinClamp float64, lodMaxClamp float64, index uint)
	HasSetVertexSamplerStateLodMinClampLodMaxClampAtIndex() bool

	// optional
	SetStencilFrontReferenceValueBackReferenceValue(frontReferenceValue uint32, backReferenceValue uint32)
	HasSetStencilFrontReferenceValueBackReferenceValue() bool

	// optional
	SetVertexVisibleFunctionTablesWithBufferRange(functionTables VisibleFunctionTableWrapper, range_ foundation.Range)
	HasSetVertexVisibleFunctionTablesWithBufferRange() bool

	// optional
	SetStencilStoreAction(storeAction StoreAction)
	HasSetStencilStoreAction() bool

	// optional
	SetCullMode(cullMode CullMode)
	HasSetCullMode() bool

	// optional
	SetTriangleFillMode(fillMode TriangleFillMode)
	HasSetTriangleFillMode() bool

	// optional
	SetTileTexturesWithRange(textures TextureWrapper, range_ foundation.Range)
	HasSetTileTexturesWithRange() bool

	// optional
	MemoryBarrierWithResourcesCountAfterStagesBeforeStages(resources ResourceWrapper, count uint, after RenderStages, before RenderStages)
	HasMemoryBarrierWithResourcesCountAfterStagesBeforeStages() bool

	// optional
	SetFragmentAccelerationStructureAtBufferIndex(accelerationStructure AccelerationStructureWrapper, bufferIndex uint)
	HasSetFragmentAccelerationStructureAtBufferIndex() bool

	// optional
	SetDepthBiasSlopeScaleClamp(depthBias float64, slopeScale float64, clamp float64)
	HasSetDepthBiasSlopeScaleClamp() bool

	// optional
	SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer CounterSampleBufferWrapper, sampleIndex uint, barrier bool)
	HasSampleCountersInBufferAtSampleIndexWithBarrier() bool

	// optional
	SetTileIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable IntersectionFunctionTableWrapper, bufferIndex uint)
	HasSetTileIntersectionFunctionTableAtBufferIndex() bool

	// optional
	SetRenderPipelineState(pipelineState RenderPipelineStateWrapper)
	HasSetRenderPipelineState() bool

	// optional
	SetTileBytesLengthAtIndex(bytes unsafe.Pointer, length uint, index uint)
	HasSetTileBytesLengthAtIndex() bool

	// optional
	ExecuteCommandsInBufferWithRange(indirectCommandBuffer IndirectCommandBufferWrapper, executionRange foundation.Range)
	HasExecuteCommandsInBufferWithRange() bool

	// optional
	SetFragmentBytesLengthAtIndex(bytes unsafe.Pointer, length uint, index uint)
	HasSetFragmentBytesLengthAtIndex() bool

	// optional
	SetTileBuffersOffsetsWithRange(buffers BufferWrapper, offsets *uint, range_ foundation.Range)
	HasSetTileBuffersOffsetsWithRange() bool

	// optional
	MemoryBarrierWithScopeAfterStagesBeforeStages(scope BarrierScope, after RenderStages, before RenderStages)
	HasMemoryBarrierWithScopeAfterStagesBeforeStages() bool

	// optional
	SetFragmentIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable IntersectionFunctionTableWrapper, bufferIndex uint)
	HasSetFragmentIntersectionFunctionTableAtBufferIndex() bool

	// optional
	DispatchThreadsPerTile(threadsPerTile Size)
	HasDispatchThreadsPerTile() bool

	// optional
	SetFragmentVisibleFunctionTablesWithBufferRange(functionTables VisibleFunctionTableWrapper, range_ foundation.Range)
	HasSetFragmentVisibleFunctionTablesWithBufferRange() bool

	// optional
	SetVertexTexturesWithRange(textures TextureWrapper, range_ foundation.Range)
	HasSetVertexTexturesWithRange() bool

	// optional
	SetFragmentSamplerStateAtIndex(sampler SamplerStateWrapper, index uint)
	HasSetFragmentSamplerStateAtIndex() bool

	// optional
	SetTileSamplerStateAtIndex(sampler SamplerStateWrapper, index uint)
	HasSetTileSamplerStateAtIndex() bool

	// optional
	SetColorStoreActionOptionsAtIndex(storeActionOptions StoreActionOptions, colorAttachmentIndex uint)
	HasSetColorStoreActionOptionsAtIndex() bool

	// optional
	SetVertexTextureAtIndex(texture TextureWrapper, index uint)
	HasSetVertexTextureAtIndex() bool

	// optional
	SetTessellationFactorScale(scale float64)
	HasSetTessellationFactorScale() bool

	// optional
	UseHeapStages(heap HeapWrapper, stages RenderStages)
	HasUseHeapStages() bool

	// optional
	SetColorStoreActionAtIndex(storeAction StoreAction, colorAttachmentIndex uint)
	HasSetColorStoreActionAtIndex() bool

	// optional
	SetStencilStoreActionOptions(storeActionOptions StoreActionOptions)
	HasSetStencilStoreActionOptions() bool

	// optional
	UseResourcesCountUsageStages(resources ResourceWrapper, count uint, usage ResourceUsage, stages RenderStages)
	HasUseResourcesCountUsageStages() bool

	// optional
	SetDepthClipMode(depthClipMode DepthClipMode)
	HasSetDepthClipMode() bool

	// optional
	SetVertexBufferOffsetAtIndex(offset uint, index uint)
	HasSetVertexBufferOffsetAtIndex() bool

	// optional
	UseHeapsCountStages(heaps HeapWrapper, count uint, stages RenderStages)
	HasUseHeapsCountStages() bool

	// optional
	SetTileVisibleFunctionTablesWithBufferRange(functionTables VisibleFunctionTableWrapper, range_ foundation.Range)
	HasSetTileVisibleFunctionTablesWithBufferRange() bool

	// optional
	SetBlendColorRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64)
	HasSetBlendColorRedGreenBlueAlpha() bool

	// optional
	SetVertexVisibleFunctionTableAtBufferIndex(functionTable VisibleFunctionTableWrapper, bufferIndex uint)
	HasSetVertexVisibleFunctionTableAtBufferIndex() bool

	// optional
	UpdateFenceAfterStages(fence FenceWrapper, stages RenderStages)
	HasUpdateFenceAfterStages() bool

	// optional
	SetScissorRectsCount(scissorRects *ScissorRect, count uint)
	HasSetScissorRectsCount() bool

	// optional
	SetDepthStoreActionOptions(storeActionOptions StoreActionOptions)
	HasSetDepthStoreActionOptions() bool

	// optional
	SetTileAccelerationStructureAtBufferIndex(accelerationStructure AccelerationStructureWrapper, bufferIndex uint)
	HasSetTileAccelerationStructureAtBufferIndex() bool

	// optional
	SetFragmentBuffersOffsetsWithRange(buffers BufferWrapper, offsets *uint, range_ foundation.Range)
	HasSetFragmentBuffersOffsetsWithRange() bool

	// optional
	SetVertexSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers SamplerStateWrapper, lodMinClamps *float64, lodMaxClamps *float64, range_ foundation.Range)
	HasSetVertexSamplerStatesLodMinClampsLodMaxClampsWithRange() bool

	// optional
	SetFrontFacingWinding(frontFacingWinding Winding)
	HasSetFrontFacingWinding() bool

	// optional
	DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCount(primitiveType PrimitiveType, indexCount uint, indexType IndexType, indexBuffer BufferWrapper, indexBufferOffset uint, instanceCount uint)
	HasDrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCount() bool

	// optional
	SetStencilReferenceValue(referenceValue uint32)
	HasSetStencilReferenceValue() bool

	// optional
	WaitForFenceBeforeStages(fence FenceWrapper, stages RenderStages)
	HasWaitForFenceBeforeStages() bool

	// optional
	TileWidth() uint
	HasTileWidth() bool

	// optional
	TileHeight() uint
	HasTileHeight() bool
}

// A concrete type wrapper for the [PRenderCommandEncoder] protocol.
type RenderCommandEncoderWrapper struct {
	objc.Object
}

func (r_ RenderCommandEncoderWrapper) HasSetViewportsCount() bool {
	return r_.RespondsToSelector(objc.Sel("setViewports:count:"))
}

// Configures the render pipeline with multiple viewports that apply transformations and clipping rectangles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2869738-setviewports?language=objc
func (r_ RenderCommandEncoderWrapper) SetViewportsCount(viewports *Viewport, count uint) {
	objc.Call[objc.Void](r_, objc.Sel("setViewports:count:"), viewports, count)
}

func (r_ RenderCommandEncoderWrapper) HasSetThreadgroupMemoryLengthOffsetAtIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setThreadgroupMemoryLength:offset:atIndex:"))
}

// Configures the size of a threadgroup memory buffer for an entry in the fragment or tile shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2866165-setthreadgroupmemorylength?language=objc
func (r_ RenderCommandEncoderWrapper) SetThreadgroupMemoryLengthOffsetAtIndex(length uint, offset uint, index uint) {
	objc.Call[objc.Void](r_, objc.Sel("setThreadgroupMemoryLength:offset:atIndex:"), length, offset, index)
}

func (r_ RenderCommandEncoderWrapper) HasSetFragmentVisibleFunctionTableAtBufferIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setFragmentVisibleFunctionTable:atBufferIndex:"))
}

// Assigns a visible function table to an entry in the fragment shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/3750556-setfragmentvisiblefunctiontable?language=objc
func (r_ RenderCommandEncoderWrapper) SetFragmentVisibleFunctionTableAtBufferIndex(functionTable PVisibleFunctionTable, bufferIndex uint) {
	po0 := objc.WrapAsProtocol("MTLVisibleFunctionTable", functionTable)
	objc.Call[objc.Void](r_, objc.Sel("setFragmentVisibleFunctionTable:atBufferIndex:"), po0, bufferIndex)
}

func (r_ RenderCommandEncoderWrapper) HasSetFragmentTexturesWithRange() bool {
	return r_.RespondsToSelector(objc.Sel("setFragmentTextures:withRange:"))
}

// Assigns multiple textures to a range of entries in the fragment shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515878-setfragmenttextures?language=objc
func (r_ RenderCommandEncoderWrapper) SetFragmentTexturesWithRange(textures PTexture, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLTexture", textures)
	objc.Call[objc.Void](r_, objc.Sel("setFragmentTextures:withRange:"), po0, range_)
}

func (r_ RenderCommandEncoderWrapper) HasSetViewport() bool {
	return r_.RespondsToSelector(objc.Sel("setViewport:"))
}

// Configures the render pipeline with a viewport that applies a transformation and a clipping rectangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515527-setviewport?language=objc
func (r_ RenderCommandEncoderWrapper) SetViewport(viewport Viewport) {
	objc.Call[objc.Void](r_, objc.Sel("setViewport:"), viewport)
}

func (r_ RenderCommandEncoderWrapper) HasDrawIndexedPatchesPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetIndirectBufferIndirectBufferOffset() bool {
	return r_.RespondsToSelector(objc.Sel("drawIndexedPatches:patchIndexBuffer:patchIndexBufferOffset:controlPointIndexBuffer:controlPointIndexBufferOffset:indirectBuffer:indirectBufferOffset:"))
}

// Encodes a draw command that renders multiple instances of tessellated patches with a control point index buffer and indirect arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1639949-drawindexedpatches?language=objc
func (r_ RenderCommandEncoderWrapper) DrawIndexedPatchesPatchIndexBufferPatchIndexBufferOffsetControlPointIndexBufferControlPointIndexBufferOffsetIndirectBufferIndirectBufferOffset(numberOfPatchControlPoints uint, patchIndexBuffer PBuffer, patchIndexBufferOffset uint, controlPointIndexBuffer PBuffer, controlPointIndexBufferOffset uint, indirectBuffer PBuffer, indirectBufferOffset uint) {
	po1 := objc.WrapAsProtocol("MTLBuffer", patchIndexBuffer)
	po3 := objc.WrapAsProtocol("MTLBuffer", controlPointIndexBuffer)
	po5 := objc.WrapAsProtocol("MTLBuffer", indirectBuffer)
	objc.Call[objc.Void](r_, objc.Sel("drawIndexedPatches:patchIndexBuffer:patchIndexBufferOffset:controlPointIndexBuffer:controlPointIndexBufferOffset:indirectBuffer:indirectBufferOffset:"), numberOfPatchControlPoints, po1, patchIndexBufferOffset, po3, controlPointIndexBufferOffset, po5, indirectBufferOffset)
}

func (r_ RenderCommandEncoderWrapper) HasSetTileSamplerStatesLodMinClampsLodMaxClampsWithRange() bool {
	return r_.RespondsToSelector(objc.Sel("setTileSamplerStates:lodMinClamps:lodMaxClamps:withRange:"))
}

// Assigns multiple sampler states and clamp values to a range of entries in the tile shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2866170-settilesamplerstates?language=objc
func (r_ RenderCommandEncoderWrapper) SetTileSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers PSamplerState, lodMinClamps *float64, lodMaxClamps *float64, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLSamplerState", samplers)
	objc.Call[objc.Void](r_, objc.Sel("setTileSamplerStates:lodMinClamps:lodMaxClamps:withRange:"), po0, lodMinClamps, lodMaxClamps, range_)
}

func (r_ RenderCommandEncoderWrapper) HasSetFragmentTextureAtIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setFragmentTexture:atIndex:"))
}

// Assigns a texture to an entry in the fragment shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515390-setfragmenttexture?language=objc
func (r_ RenderCommandEncoderWrapper) SetFragmentTextureAtIndex(texture PTexture, index uint) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	objc.Call[objc.Void](r_, objc.Sel("setFragmentTexture:atIndex:"), po0, index)
}

func (r_ RenderCommandEncoderWrapper) HasSetDepthStoreAction() bool {
	return r_.RespondsToSelector(objc.Sel("setDepthStoreAction:"))
}

// Configures the store action for the depth attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1640036-setdepthstoreaction?language=objc
func (r_ RenderCommandEncoderWrapper) SetDepthStoreAction(storeAction StoreAction) {
	objc.Call[objc.Void](r_, objc.Sel("setDepthStoreAction:"), storeAction)
}

func (r_ RenderCommandEncoderWrapper) HasSetFragmentSamplerStatesWithRange() bool {
	return r_.RespondsToSelector(objc.Sel("setFragmentSamplerStates:withRange:"))
}

// Assigns multiple sampler states to a range of entries in the fragment shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515970-setfragmentsamplerstates?language=objc
func (r_ RenderCommandEncoderWrapper) SetFragmentSamplerStatesWithRange(samplers PSamplerState, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLSamplerState", samplers)
	objc.Call[objc.Void](r_, objc.Sel("setFragmentSamplerStates:withRange:"), po0, range_)
}

func (r_ RenderCommandEncoderWrapper) HasSetVisibilityResultModeOffset() bool {
	return r_.RespondsToSelector(objc.Sel("setVisibilityResultMode:offset:"))
}

// Configures which visibility test the GPU runs and the destination for any results it generates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515556-setvisibilityresultmode?language=objc
func (r_ RenderCommandEncoderWrapper) SetVisibilityResultModeOffset(mode VisibilityResultMode, offset uint) {
	objc.Call[objc.Void](r_, objc.Sel("setVisibilityResultMode:offset:"), mode, offset)
}

func (r_ RenderCommandEncoderWrapper) HasSetTileBufferOffsetAtIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setTileBufferOffset:atIndex:"))
}

// Updates an entry in the tile shader argument table with a new location within the entry’s current buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2866159-settilebufferoffset?language=objc
func (r_ RenderCommandEncoderWrapper) SetTileBufferOffsetAtIndex(offset uint, index uint) {
	objc.Call[objc.Void](r_, objc.Sel("setTileBufferOffset:atIndex:"), offset, index)
}

func (r_ RenderCommandEncoderWrapper) HasSetFragmentBufferOffsetAtIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setFragmentBufferOffset:atIndex:"))
}

// Updates an entry in the fragment shader argument table with a new location within the entry’s current buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515917-setfragmentbufferoffset?language=objc
func (r_ RenderCommandEncoderWrapper) SetFragmentBufferOffsetAtIndex(offset uint, index uint) {
	objc.Call[objc.Void](r_, objc.Sel("setFragmentBufferOffset:atIndex:"), offset, index)
}

func (r_ RenderCommandEncoderWrapper) HasSetVertexBufferOffsetAtIndex_() bool {
	return r_.RespondsToSelector(objc.Sel("setVertexBuffer:offset:atIndex:"))
}

// Assigns a buffer to an entry in the vertex shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515829-setvertexbuffer?language=objc
func (r_ RenderCommandEncoderWrapper) SetVertexBufferOffsetAtIndex_(buffer PBuffer, offset uint, index uint) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	objc.Call[objc.Void](r_, objc.Sel("setVertexBuffer:offset:atIndex:"), po0, offset, index)
}

func (r_ RenderCommandEncoderWrapper) HasSetVertexIntersectionFunctionTableAtBufferIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setVertexIntersectionFunctionTable:atBufferIndex:"))
}

// Assigns an intersection function table to an entry in the vertex shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/3750564-setvertexintersectionfunctiontab?language=objc
func (r_ RenderCommandEncoderWrapper) SetVertexIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable PIntersectionFunctionTable, bufferIndex uint) {
	po0 := objc.WrapAsProtocol("MTLIntersectionFunctionTable", intersectionFunctionTable)
	objc.Call[objc.Void](r_, objc.Sel("setVertexIntersectionFunctionTable:atBufferIndex:"), po0, bufferIndex)
}

func (r_ RenderCommandEncoderWrapper) HasSetTileBufferOffsetAtIndex_() bool {
	return r_.RespondsToSelector(objc.Sel("setTileBuffer:offset:atIndex:"))
}

// Assigns a buffer to an entry in the tile shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2866157-settilebuffer?language=objc
func (r_ RenderCommandEncoderWrapper) SetTileBufferOffsetAtIndex_(buffer PBuffer, offset uint, index uint) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	objc.Call[objc.Void](r_, objc.Sel("setTileBuffer:offset:atIndex:"), po0, offset, index)
}

func (r_ RenderCommandEncoderWrapper) HasSetTessellationFactorBufferOffsetInstanceStride() bool {
	return r_.RespondsToSelector(objc.Sel("setTessellationFactorBuffer:offset:instanceStride:"))
}

// Configures the per-patch tessellation factors for any subsequent patch-drawing commands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1640035-settessellationfactorbuffer?language=objc
func (r_ RenderCommandEncoderWrapper) SetTessellationFactorBufferOffsetInstanceStride(buffer PBuffer, offset uint, instanceStride uint) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	objc.Call[objc.Void](r_, objc.Sel("setTessellationFactorBuffer:offset:instanceStride:"), po0, offset, instanceStride)
}

func (r_ RenderCommandEncoderWrapper) HasSetTileVisibleFunctionTableAtBufferIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setTileVisibleFunctionTable:atBufferIndex:"))
}

// Assigns a visible function table to an entry in the tile shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/3750561-settilevisiblefunctiontable?language=objc
func (r_ RenderCommandEncoderWrapper) SetTileVisibleFunctionTableAtBufferIndex(functionTable PVisibleFunctionTable, bufferIndex uint) {
	po0 := objc.WrapAsProtocol("MTLVisibleFunctionTable", functionTable)
	objc.Call[objc.Void](r_, objc.Sel("setTileVisibleFunctionTable:atBufferIndex:"), po0, bufferIndex)
}

func (r_ RenderCommandEncoderWrapper) HasDrawPatchesPatchIndexBufferPatchIndexBufferOffsetIndirectBufferIndirectBufferOffset() bool {
	return r_.RespondsToSelector(objc.Sel("drawPatches:patchIndexBuffer:patchIndexBufferOffset:indirectBuffer:indirectBufferOffset:"))
}

// Encodes a draw command that renders multiple instances of tessellated patches with indirect arguments. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1639895-drawpatches?language=objc
func (r_ RenderCommandEncoderWrapper) DrawPatchesPatchIndexBufferPatchIndexBufferOffsetIndirectBufferIndirectBufferOffset(numberOfPatchControlPoints uint, patchIndexBuffer PBuffer, patchIndexBufferOffset uint, indirectBuffer PBuffer, indirectBufferOffset uint) {
	po1 := objc.WrapAsProtocol("MTLBuffer", patchIndexBuffer)
	po3 := objc.WrapAsProtocol("MTLBuffer", indirectBuffer)
	objc.Call[objc.Void](r_, objc.Sel("drawPatches:patchIndexBuffer:patchIndexBufferOffset:indirectBuffer:indirectBufferOffset:"), numberOfPatchControlPoints, po1, patchIndexBufferOffset, po3, indirectBufferOffset)
}

func (r_ RenderCommandEncoderWrapper) HasUseResourceUsageStages() bool {
	return r_.RespondsToSelector(objc.Sel("useResource:usage:stages:"))
}

// Ensures the shaders in the render pass’s subsequent draw commands have access to a resource. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/3043404-useresource?language=objc
func (r_ RenderCommandEncoderWrapper) UseResourceUsageStages(resource PResource, usage ResourceUsage, stages RenderStages) {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	objc.Call[objc.Void](r_, objc.Sel("useResource:usage:stages:"), po0, usage, stages)
}

func (r_ RenderCommandEncoderWrapper) HasDrawPrimitivesVertexStartVertexCount() bool {
	return r_.RespondsToSelector(objc.Sel("drawPrimitives:vertexStart:vertexCount:"))
}

// Encodes a draw command that renders an instance of a geometric primitive. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1516326-drawprimitives?language=objc
func (r_ RenderCommandEncoderWrapper) DrawPrimitivesVertexStartVertexCount(primitiveType PrimitiveType, vertexStart uint, vertexCount uint) {
	objc.Call[objc.Void](r_, objc.Sel("drawPrimitives:vertexStart:vertexCount:"), primitiveType, vertexStart, vertexCount)
}

func (r_ RenderCommandEncoderWrapper) HasSetVertexBuffersOffsetsWithRange() bool {
	return r_.RespondsToSelector(objc.Sel("setVertexBuffers:offsets:withRange:"))
}

// Assigns multiple buffers to a range of entries in the vertex shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515987-setvertexbuffers?language=objc
func (r_ RenderCommandEncoderWrapper) SetVertexBuffersOffsetsWithRange(buffers PBuffer, offsets *uint, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffers)
	objc.Call[objc.Void](r_, objc.Sel("setVertexBuffers:offsets:withRange:"), po0, offsets, range_)
}

func (r_ RenderCommandEncoderWrapper) HasSetFragmentBufferOffsetAtIndex_() bool {
	return r_.RespondsToSelector(objc.Sel("setFragmentBuffer:offset:atIndex:"))
}

// Assigns a buffer to an entry in the fragment shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515470-setfragmentbuffer?language=objc
func (r_ RenderCommandEncoderWrapper) SetFragmentBufferOffsetAtIndex_(buffer PBuffer, offset uint, index uint) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	objc.Call[objc.Void](r_, objc.Sel("setFragmentBuffer:offset:atIndex:"), po0, offset, index)
}

func (r_ RenderCommandEncoderWrapper) HasSetDepthStencilState() bool {
	return r_.RespondsToSelector(objc.Sel("setDepthStencilState:"))
}

// Configures the combined depth and stencil state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1516119-setdepthstencilstate?language=objc
func (r_ RenderCommandEncoderWrapper) SetDepthStencilState(depthStencilState PDepthStencilState) {
	po0 := objc.WrapAsProtocol("MTLDepthStencilState", depthStencilState)
	objc.Call[objc.Void](r_, objc.Sel("setDepthStencilState:"), po0)
}

func (r_ RenderCommandEncoderWrapper) HasSetVertexAccelerationStructureAtBufferIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setVertexAccelerationStructure:atBufferIndex:"))
}

// Assigns an acceleration structure to an entry in the vertex shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/3750563-setvertexaccelerationstructure?language=objc
func (r_ RenderCommandEncoderWrapper) SetVertexAccelerationStructureAtBufferIndex(accelerationStructure PAccelerationStructure, bufferIndex uint) {
	po0 := objc.WrapAsProtocol("MTLAccelerationStructure", accelerationStructure)
	objc.Call[objc.Void](r_, objc.Sel("setVertexAccelerationStructure:atBufferIndex:"), po0, bufferIndex)
}

func (r_ RenderCommandEncoderWrapper) HasSetVertexBytesLengthAtIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setVertexBytes:length:atIndex:"))
}

// Creates a buffer from bytes and assigns it to an entry in the vertex shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515846-setvertexbytes?language=objc
func (r_ RenderCommandEncoderWrapper) SetVertexBytesLengthAtIndex(bytes unsafe.Pointer, length uint, index uint) {
	objc.Call[objc.Void](r_, objc.Sel("setVertexBytes:length:atIndex:"), bytes, length, index)
}

func (r_ RenderCommandEncoderWrapper) HasSetVertexAmplificationCountViewMappings() bool {
	return r_.RespondsToSelector(objc.Sel("setVertexAmplificationCount:viewMappings:"))
}

// Configures the number of output vertices the render pipeline produces for each input vertex, optionally with render target and viewport offsets. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/3088849-setvertexamplificationcount?language=objc
func (r_ RenderCommandEncoderWrapper) SetVertexAmplificationCountViewMappings(count uint, viewMappings *VertexAmplificationViewMapping) {
	objc.Call[objc.Void](r_, objc.Sel("setVertexAmplificationCount:viewMappings:"), count, viewMappings)
}

func (r_ RenderCommandEncoderWrapper) HasSetScissorRect() bool {
	return r_.RespondsToSelector(objc.Sel("setScissorRect:"))
}

// Configures a rectangle for the fragment scissor test. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515583-setscissorrect?language=objc
func (r_ RenderCommandEncoderWrapper) SetScissorRect(rect ScissorRect) {
	objc.Call[objc.Void](r_, objc.Sel("setScissorRect:"), rect)
}

func (r_ RenderCommandEncoderWrapper) HasSetTileTextureAtIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setTileTexture:atIndex:"))
}

// Assigns a texture to an entry in the tile shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2866173-settiletexture?language=objc
func (r_ RenderCommandEncoderWrapper) SetTileTextureAtIndex(texture PTexture, index uint) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	objc.Call[objc.Void](r_, objc.Sel("setTileTexture:atIndex:"), po0, index)
}

func (r_ RenderCommandEncoderWrapper) HasSetVertexSamplerStateLodMinClampLodMaxClampAtIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setVertexSamplerState:lodMinClamp:lodMaxClamp:atIndex:"))
}

// Assigns a sampler state and clamp values to an entry in the vertex shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515864-setvertexsamplerstate?language=objc
func (r_ RenderCommandEncoderWrapper) SetVertexSamplerStateLodMinClampLodMaxClampAtIndex(sampler PSamplerState, lodMinClamp float64, lodMaxClamp float64, index uint) {
	po0 := objc.WrapAsProtocol("MTLSamplerState", sampler)
	objc.Call[objc.Void](r_, objc.Sel("setVertexSamplerState:lodMinClamp:lodMaxClamp:atIndex:"), po0, lodMinClamp, lodMaxClamp, index)
}

func (r_ RenderCommandEncoderWrapper) HasSetStencilFrontReferenceValueBackReferenceValue() bool {
	return r_.RespondsToSelector(objc.Sel("setStencilFrontReferenceValue:backReferenceValue:"))
}

// Configures different comparison values for front- and back-facing primitives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515522-setstencilfrontreferencevalue?language=objc
func (r_ RenderCommandEncoderWrapper) SetStencilFrontReferenceValueBackReferenceValue(frontReferenceValue uint32, backReferenceValue uint32) {
	objc.Call[objc.Void](r_, objc.Sel("setStencilFrontReferenceValue:backReferenceValue:"), frontReferenceValue, backReferenceValue)
}

func (r_ RenderCommandEncoderWrapper) HasSetVertexVisibleFunctionTablesWithBufferRange() bool {
	return r_.RespondsToSelector(objc.Sel("setVertexVisibleFunctionTables:withBufferRange:"))
}

// Assigns multiple visible function tables to a range of entries in the vertex shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/3750567-setvertexvisiblefunctiontables?language=objc
func (r_ RenderCommandEncoderWrapper) SetVertexVisibleFunctionTablesWithBufferRange(functionTables PVisibleFunctionTable, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLVisibleFunctionTable", functionTables)
	objc.Call[objc.Void](r_, objc.Sel("setVertexVisibleFunctionTables:withBufferRange:"), po0, range_)
}

func (r_ RenderCommandEncoderWrapper) HasSetStencilStoreAction() bool {
	return r_.RespondsToSelector(objc.Sel("setStencilStoreAction:"))
}

// Configures the store action for the stencil attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1639936-setstencilstoreaction?language=objc
func (r_ RenderCommandEncoderWrapper) SetStencilStoreAction(storeAction StoreAction) {
	objc.Call[objc.Void](r_, objc.Sel("setStencilStoreAction:"), storeAction)
}

func (r_ RenderCommandEncoderWrapper) HasSetCullMode() bool {
	return r_.RespondsToSelector(objc.Sel("setCullMode:"))
}

// Configures how the render pipeline determines which primitives to remove. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515975-setcullmode?language=objc
func (r_ RenderCommandEncoderWrapper) SetCullMode(cullMode CullMode) {
	objc.Call[objc.Void](r_, objc.Sel("setCullMode:"), cullMode)
}

func (r_ RenderCommandEncoderWrapper) HasSetTriangleFillMode() bool {
	return r_.RespondsToSelector(objc.Sel("setTriangleFillMode:"))
}

// Configures how subsequent draw commands rasterize triangle and triangle strip primitives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1516029-settrianglefillmode?language=objc
func (r_ RenderCommandEncoderWrapper) SetTriangleFillMode(fillMode TriangleFillMode) {
	objc.Call[objc.Void](r_, objc.Sel("setTriangleFillMode:"), fillMode)
}

func (r_ RenderCommandEncoderWrapper) HasSetTileTexturesWithRange() bool {
	return r_.RespondsToSelector(objc.Sel("setTileTextures:withRange:"))
}

// Assigns multiple textures to a range of entries in the tile shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2866172-settiletextures?language=objc
func (r_ RenderCommandEncoderWrapper) SetTileTexturesWithRange(textures PTexture, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLTexture", textures)
	objc.Call[objc.Void](r_, objc.Sel("setTileTextures:withRange:"), po0, range_)
}

func (r_ RenderCommandEncoderWrapper) HasMemoryBarrierWithResourcesCountAfterStagesBeforeStages() bool {
	return r_.RespondsToSelector(objc.Sel("memoryBarrierWithResources:count:afterStages:beforeStages:"))
}

// Creates a memory barrier that enforces the order of write and read operations for specific resources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2967441-memorybarrierwithresources?language=objc
func (r_ RenderCommandEncoderWrapper) MemoryBarrierWithResourcesCountAfterStagesBeforeStages(resources PResource, count uint, after RenderStages, before RenderStages) {
	po0 := objc.WrapAsProtocol("MTLResource", resources)
	objc.Call[objc.Void](r_, objc.Sel("memoryBarrierWithResources:count:afterStages:beforeStages:"), po0, count, after, before)
}

func (r_ RenderCommandEncoderWrapper) HasSetFragmentAccelerationStructureAtBufferIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setFragmentAccelerationStructure:atBufferIndex:"))
}

// Assigns an acceleration structure to an entry in the fragment shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/3750553-setfragmentaccelerationstructure?language=objc
func (r_ RenderCommandEncoderWrapper) SetFragmentAccelerationStructureAtBufferIndex(accelerationStructure PAccelerationStructure, bufferIndex uint) {
	po0 := objc.WrapAsProtocol("MTLAccelerationStructure", accelerationStructure)
	objc.Call[objc.Void](r_, objc.Sel("setFragmentAccelerationStructure:atBufferIndex:"), po0, bufferIndex)
}

func (r_ RenderCommandEncoderWrapper) HasSetDepthBiasSlopeScaleClamp() bool {
	return r_.RespondsToSelector(objc.Sel("setDepthBias:slopeScale:clamp:"))
}

// Configures the adjustments a render pass applies to depth values from fragment functions by a scaling factor and bias. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1516269-setdepthbias?language=objc
func (r_ RenderCommandEncoderWrapper) SetDepthBiasSlopeScaleClamp(depthBias float64, slopeScale float64, clamp float64) {
	objc.Call[objc.Void](r_, objc.Sel("setDepthBias:slopeScale:clamp:"), depthBias, slopeScale, clamp)
}

func (r_ RenderCommandEncoderWrapper) HasSampleCountersInBufferAtSampleIndexWithBarrier() bool {
	return r_.RespondsToSelector(objc.Sel("sampleCountersInBuffer:atSampleIndex:withBarrier:"))
}

// Encodes a command that samples hardware counters during the render pass and stores the data into a counter sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/3194379-samplecountersinbuffer?language=objc
func (r_ RenderCommandEncoderWrapper) SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer PCounterSampleBuffer, sampleIndex uint, barrier bool) {
	po0 := objc.WrapAsProtocol("MTLCounterSampleBuffer", sampleBuffer)
	objc.Call[objc.Void](r_, objc.Sel("sampleCountersInBuffer:atSampleIndex:withBarrier:"), po0, sampleIndex, barrier)
}

func (r_ RenderCommandEncoderWrapper) HasSetTileIntersectionFunctionTableAtBufferIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setTileIntersectionFunctionTable:atBufferIndex:"))
}

// Assigns an intersection function table to an entry in the tile shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/3750559-settileintersectionfunctiontable?language=objc
func (r_ RenderCommandEncoderWrapper) SetTileIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable PIntersectionFunctionTable, bufferIndex uint) {
	po0 := objc.WrapAsProtocol("MTLIntersectionFunctionTable", intersectionFunctionTable)
	objc.Call[objc.Void](r_, objc.Sel("setTileIntersectionFunctionTable:atBufferIndex:"), po0, bufferIndex)
}

func (r_ RenderCommandEncoderWrapper) HasSetRenderPipelineState() bool {
	return r_.RespondsToSelector(objc.Sel("setRenderPipelineState:"))
}

// Configures the encoder with a render or tile pipeline state instance that applies to your subsequent draw commands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515811-setrenderpipelinestate?language=objc
func (r_ RenderCommandEncoderWrapper) SetRenderPipelineState(pipelineState PRenderPipelineState) {
	po0 := objc.WrapAsProtocol("MTLRenderPipelineState", pipelineState)
	objc.Call[objc.Void](r_, objc.Sel("setRenderPipelineState:"), po0)
}

func (r_ RenderCommandEncoderWrapper) HasSetTileBytesLengthAtIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setTileBytes:length:atIndex:"))
}

// Creates a buffer from bytes and assigns it to an entry in the tile shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2866175-settilebytes?language=objc
func (r_ RenderCommandEncoderWrapper) SetTileBytesLengthAtIndex(bytes unsafe.Pointer, length uint, index uint) {
	objc.Call[objc.Void](r_, objc.Sel("setTileBytes:length:atIndex:"), bytes, length, index)
}

func (r_ RenderCommandEncoderWrapper) HasExecuteCommandsInBufferWithRange() bool {
	return r_.RespondsToSelector(objc.Sel("executeCommandsInBuffer:withRange:"))
}

// Encodes a command that runs a range of commands from an indirect command buffer (ICB). [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2967440-executecommandsinbuffer?language=objc
func (r_ RenderCommandEncoderWrapper) ExecuteCommandsInBufferWithRange(indirectCommandBuffer PIndirectCommandBuffer, executionRange foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLIndirectCommandBuffer", indirectCommandBuffer)
	objc.Call[objc.Void](r_, objc.Sel("executeCommandsInBuffer:withRange:"), po0, executionRange)
}

func (r_ RenderCommandEncoderWrapper) HasSetFragmentBytesLengthAtIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setFragmentBytes:length:atIndex:"))
}

// Creates a buffer from bytes and assigns it to an entry in the fragment shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1516192-setfragmentbytes?language=objc
func (r_ RenderCommandEncoderWrapper) SetFragmentBytesLengthAtIndex(bytes unsafe.Pointer, length uint, index uint) {
	objc.Call[objc.Void](r_, objc.Sel("setFragmentBytes:length:atIndex:"), bytes, length, index)
}

func (r_ RenderCommandEncoderWrapper) HasSetTileBuffersOffsetsWithRange() bool {
	return r_.RespondsToSelector(objc.Sel("setTileBuffers:offsets:withRange:"))
}

// Assigns multiple buffers to a range of entries in the tile shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2866162-settilebuffers?language=objc
func (r_ RenderCommandEncoderWrapper) SetTileBuffersOffsetsWithRange(buffers PBuffer, offsets *uint, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffers)
	objc.Call[objc.Void](r_, objc.Sel("setTileBuffers:offsets:withRange:"), po0, offsets, range_)
}

func (r_ RenderCommandEncoderWrapper) HasMemoryBarrierWithScopeAfterStagesBeforeStages() bool {
	return r_.RespondsToSelector(objc.Sel("memoryBarrierWithScope:afterStages:beforeStages:"))
}

// Creates a memory barrier that enforces the order of write and read operations for specific resource types. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2967442-memorybarrierwithscope?language=objc
func (r_ RenderCommandEncoderWrapper) MemoryBarrierWithScopeAfterStagesBeforeStages(scope BarrierScope, after RenderStages, before RenderStages) {
	objc.Call[objc.Void](r_, objc.Sel("memoryBarrierWithScope:afterStages:beforeStages:"), scope, after, before)
}

func (r_ RenderCommandEncoderWrapper) HasSetFragmentIntersectionFunctionTableAtBufferIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setFragmentIntersectionFunctionTable:atBufferIndex:"))
}

// Assigns an intersection function table to an entry in the fragment shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/3750554-setfragmentintersectionfunctiont?language=objc
func (r_ RenderCommandEncoderWrapper) SetFragmentIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable PIntersectionFunctionTable, bufferIndex uint) {
	po0 := objc.WrapAsProtocol("MTLIntersectionFunctionTable", intersectionFunctionTable)
	objc.Call[objc.Void](r_, objc.Sel("setFragmentIntersectionFunctionTable:atBufferIndex:"), po0, bufferIndex)
}

func (r_ RenderCommandEncoderWrapper) HasDispatchThreadsPerTile() bool {
	return r_.RespondsToSelector(objc.Sel("dispatchThreadsPerTile:"))
}

// Encodes a command that invokes GPU functions from the encoder’s current tile render pipeline state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2866171-dispatchthreadspertile?language=objc
func (r_ RenderCommandEncoderWrapper) DispatchThreadsPerTile(threadsPerTile Size) {
	objc.Call[objc.Void](r_, objc.Sel("dispatchThreadsPerTile:"), threadsPerTile)
}

func (r_ RenderCommandEncoderWrapper) HasSetFragmentVisibleFunctionTablesWithBufferRange() bool {
	return r_.RespondsToSelector(objc.Sel("setFragmentVisibleFunctionTables:withBufferRange:"))
}

// Assigns multiple visible function tables to a range of entries in the fragment shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/3750557-setfragmentvisiblefunctiontables?language=objc
func (r_ RenderCommandEncoderWrapper) SetFragmentVisibleFunctionTablesWithBufferRange(functionTables PVisibleFunctionTable, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLVisibleFunctionTable", functionTables)
	objc.Call[objc.Void](r_, objc.Sel("setFragmentVisibleFunctionTables:withBufferRange:"), po0, range_)
}

func (r_ RenderCommandEncoderWrapper) HasSetVertexTexturesWithRange() bool {
	return r_.RespondsToSelector(objc.Sel("setVertexTextures:withRange:"))
}

// Assigns multiple textures to a range of entries in the vertex shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1516109-setvertextextures?language=objc
func (r_ RenderCommandEncoderWrapper) SetVertexTexturesWithRange(textures PTexture, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLTexture", textures)
	objc.Call[objc.Void](r_, objc.Sel("setVertexTextures:withRange:"), po0, range_)
}

func (r_ RenderCommandEncoderWrapper) HasSetFragmentSamplerStateAtIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setFragmentSamplerState:atIndex:"))
}

// Assigns a sampler state to an entry in the fragment shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515577-setfragmentsamplerstate?language=objc
func (r_ RenderCommandEncoderWrapper) SetFragmentSamplerStateAtIndex(sampler PSamplerState, index uint) {
	po0 := objc.WrapAsProtocol("MTLSamplerState", sampler)
	objc.Call[objc.Void](r_, objc.Sel("setFragmentSamplerState:atIndex:"), po0, index)
}

func (r_ RenderCommandEncoderWrapper) HasSetTileSamplerStateAtIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setTileSamplerState:atIndex:"))
}

// Assigns a sampler state to an entry in the tile shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2866169-settilesamplerstate?language=objc
func (r_ RenderCommandEncoderWrapper) SetTileSamplerStateAtIndex(sampler PSamplerState, index uint) {
	po0 := objc.WrapAsProtocol("MTLSamplerState", sampler)
	objc.Call[objc.Void](r_, objc.Sel("setTileSamplerState:atIndex:"), po0, index)
}

func (r_ RenderCommandEncoderWrapper) HasSetColorStoreActionOptionsAtIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setColorStoreActionOptions:atIndex:"))
}

// Configures the store action options for a color attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2919773-setcolorstoreactionoptions?language=objc
func (r_ RenderCommandEncoderWrapper) SetColorStoreActionOptionsAtIndex(storeActionOptions StoreActionOptions, colorAttachmentIndex uint) {
	objc.Call[objc.Void](r_, objc.Sel("setColorStoreActionOptions:atIndex:"), storeActionOptions, colorAttachmentIndex)
}

func (r_ RenderCommandEncoderWrapper) HasSetVertexTextureAtIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setVertexTexture:atIndex:"))
}

// Assigns a texture to an entry in the vertex shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515842-setvertextexture?language=objc
func (r_ RenderCommandEncoderWrapper) SetVertexTextureAtIndex(texture PTexture, index uint) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	objc.Call[objc.Void](r_, objc.Sel("setVertexTexture:atIndex:"), po0, index)
}

func (r_ RenderCommandEncoderWrapper) HasSetTessellationFactorScale() bool {
	return r_.RespondsToSelector(objc.Sel("setTessellationFactorScale:"))
}

// Configures the scale factor for per-patch tessellation factors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1639992-settessellationfactorscale?language=objc
func (r_ RenderCommandEncoderWrapper) SetTessellationFactorScale(scale float64) {
	objc.Call[objc.Void](r_, objc.Sel("setTessellationFactorScale:"), scale)
}

func (r_ RenderCommandEncoderWrapper) HasUseHeapStages() bool {
	return r_.RespondsToSelector(objc.Sel("useHeap:stages:"))
}

// Ensures the shaders in the render pass’s subsequent draw commands have access to the resources you allocate from a heap. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/3043402-useheap?language=objc
func (r_ RenderCommandEncoderWrapper) UseHeapStages(heap PHeap, stages RenderStages) {
	po0 := objc.WrapAsProtocol("MTLHeap", heap)
	objc.Call[objc.Void](r_, objc.Sel("useHeap:stages:"), po0, stages)
}

func (r_ RenderCommandEncoderWrapper) HasSetColorStoreActionAtIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setColorStoreAction:atIndex:"))
}

// Configures the store action for a color attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1640054-setcolorstoreaction?language=objc
func (r_ RenderCommandEncoderWrapper) SetColorStoreActionAtIndex(storeAction StoreAction, colorAttachmentIndex uint) {
	objc.Call[objc.Void](r_, objc.Sel("setColorStoreAction:atIndex:"), storeAction, colorAttachmentIndex)
}

func (r_ RenderCommandEncoderWrapper) HasSetStencilStoreActionOptions() bool {
	return r_.RespondsToSelector(objc.Sel("setStencilStoreActionOptions:"))
}

// Configures the store action options for the stencil attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2919771-setstencilstoreactionoptions?language=objc
func (r_ RenderCommandEncoderWrapper) SetStencilStoreActionOptions(storeActionOptions StoreActionOptions) {
	objc.Call[objc.Void](r_, objc.Sel("setStencilStoreActionOptions:"), storeActionOptions)
}

func (r_ RenderCommandEncoderWrapper) HasUseResourcesCountUsageStages() bool {
	return r_.RespondsToSelector(objc.Sel("useResources:count:usage:stages:"))
}

// Ensures the shaders in the render pass’s subsequent draw commands have access to multiple resources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/3043405-useresources?language=objc
func (r_ RenderCommandEncoderWrapper) UseResourcesCountUsageStages(resources PResource, count uint, usage ResourceUsage, stages RenderStages) {
	po0 := objc.WrapAsProtocol("MTLResource", resources)
	objc.Call[objc.Void](r_, objc.Sel("useResources:count:usage:stages:"), po0, count, usage, stages)
}

func (r_ RenderCommandEncoderWrapper) HasSetDepthClipMode() bool {
	return r_.RespondsToSelector(objc.Sel("setDepthClipMode:"))
}

// Configures how the render pipeline handles fragments outside the near and far planes of the view frustum. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1516267-setdepthclipmode?language=objc
func (r_ RenderCommandEncoderWrapper) SetDepthClipMode(depthClipMode DepthClipMode) {
	objc.Call[objc.Void](r_, objc.Sel("setDepthClipMode:"), depthClipMode)
}

func (r_ RenderCommandEncoderWrapper) HasSetVertexBufferOffsetAtIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setVertexBufferOffset:atIndex:"))
}

// Updates an entry in the vertex shader argument table with a new location within the entry’s current buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515433-setvertexbufferoffset?language=objc
func (r_ RenderCommandEncoderWrapper) SetVertexBufferOffsetAtIndex(offset uint, index uint) {
	objc.Call[objc.Void](r_, objc.Sel("setVertexBufferOffset:atIndex:"), offset, index)
}

func (r_ RenderCommandEncoderWrapper) HasUseHeapsCountStages() bool {
	return r_.RespondsToSelector(objc.Sel("useHeaps:count:stages:"))
}

// Ensures the shaders in the render pass’s subsequent draw commands have access to the resources you allocate from multiple heaps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/3043403-useheaps?language=objc
func (r_ RenderCommandEncoderWrapper) UseHeapsCountStages(heaps PHeap, count uint, stages RenderStages) {
	po0 := objc.WrapAsProtocol("MTLHeap", heaps)
	objc.Call[objc.Void](r_, objc.Sel("useHeaps:count:stages:"), po0, count, stages)
}

func (r_ RenderCommandEncoderWrapper) HasSetTileVisibleFunctionTablesWithBufferRange() bool {
	return r_.RespondsToSelector(objc.Sel("setTileVisibleFunctionTables:withBufferRange:"))
}

// Assigns multiple visible function tables to a range of entries in the tile shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/3750562-settilevisiblefunctiontables?language=objc
func (r_ RenderCommandEncoderWrapper) SetTileVisibleFunctionTablesWithBufferRange(functionTables PVisibleFunctionTable, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLVisibleFunctionTable", functionTables)
	objc.Call[objc.Void](r_, objc.Sel("setTileVisibleFunctionTables:withBufferRange:"), po0, range_)
}

func (r_ RenderCommandEncoderWrapper) HasSetBlendColorRedGreenBlueAlpha() bool {
	return r_.RespondsToSelector(objc.Sel("setBlendColorRed:green:blue:alpha:"))
}

// Configures each pixel component value, including alpha, for the render pipeline’s constant blend color. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515592-setblendcolorred?language=objc
func (r_ RenderCommandEncoderWrapper) SetBlendColorRedGreenBlueAlpha(red float64, green float64, blue float64, alpha float64) {
	objc.Call[objc.Void](r_, objc.Sel("setBlendColorRed:green:blue:alpha:"), red, green, blue, alpha)
}

func (r_ RenderCommandEncoderWrapper) HasSetVertexVisibleFunctionTableAtBufferIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setVertexVisibleFunctionTable:atBufferIndex:"))
}

// Assigns a visible function table to an entry in the vertex shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/3750566-setvertexvisiblefunctiontable?language=objc
func (r_ RenderCommandEncoderWrapper) SetVertexVisibleFunctionTableAtBufferIndex(functionTable PVisibleFunctionTable, bufferIndex uint) {
	po0 := objc.WrapAsProtocol("MTLVisibleFunctionTable", functionTable)
	objc.Call[objc.Void](r_, objc.Sel("setVertexVisibleFunctionTable:atBufferIndex:"), po0, bufferIndex)
}

func (r_ RenderCommandEncoderWrapper) HasUpdateFenceAfterStages() bool {
	return r_.RespondsToSelector(objc.Sel("updateFence:afterStages:"))
}

// Encodes a command that instructs the GPU to update a fence after one or more stages, which signals passes waiting on the fence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1648377-updatefence?language=objc
func (r_ RenderCommandEncoderWrapper) UpdateFenceAfterStages(fence PFence, stages RenderStages) {
	po0 := objc.WrapAsProtocol("MTLFence", fence)
	objc.Call[objc.Void](r_, objc.Sel("updateFence:afterStages:"), po0, stages)
}

func (r_ RenderCommandEncoderWrapper) HasSetScissorRectsCount() bool {
	return r_.RespondsToSelector(objc.Sel("setScissorRects:count:"))
}

// Configures multiple rectangles for the fragment scissor test. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2869722-setscissorrects?language=objc
func (r_ RenderCommandEncoderWrapper) SetScissorRectsCount(scissorRects *ScissorRect, count uint) {
	objc.Call[objc.Void](r_, objc.Sel("setScissorRects:count:"), scissorRects, count)
}

func (r_ RenderCommandEncoderWrapper) HasSetDepthStoreActionOptions() bool {
	return r_.RespondsToSelector(objc.Sel("setDepthStoreActionOptions:"))
}

// Configures the store action options for the depth attachment. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2919772-setdepthstoreactionoptions?language=objc
func (r_ RenderCommandEncoderWrapper) SetDepthStoreActionOptions(storeActionOptions StoreActionOptions) {
	objc.Call[objc.Void](r_, objc.Sel("setDepthStoreActionOptions:"), storeActionOptions)
}

func (r_ RenderCommandEncoderWrapper) HasSetTileAccelerationStructureAtBufferIndex() bool {
	return r_.RespondsToSelector(objc.Sel("setTileAccelerationStructure:atBufferIndex:"))
}

// Assigns an acceleration structure to an entry in the tile shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/3750558-settileaccelerationstructure?language=objc
func (r_ RenderCommandEncoderWrapper) SetTileAccelerationStructureAtBufferIndex(accelerationStructure PAccelerationStructure, bufferIndex uint) {
	po0 := objc.WrapAsProtocol("MTLAccelerationStructure", accelerationStructure)
	objc.Call[objc.Void](r_, objc.Sel("setTileAccelerationStructure:atBufferIndex:"), po0, bufferIndex)
}

func (r_ RenderCommandEncoderWrapper) HasSetFragmentBuffersOffsetsWithRange() bool {
	return r_.RespondsToSelector(objc.Sel("setFragmentBuffers:offsets:withRange:"))
}

// Assigns multiple buffers to a range of entries in the fragment shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515724-setfragmentbuffers?language=objc
func (r_ RenderCommandEncoderWrapper) SetFragmentBuffersOffsetsWithRange(buffers PBuffer, offsets *uint, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffers)
	objc.Call[objc.Void](r_, objc.Sel("setFragmentBuffers:offsets:withRange:"), po0, offsets, range_)
}

func (r_ RenderCommandEncoderWrapper) HasSetVertexSamplerStatesLodMinClampsLodMaxClampsWithRange() bool {
	return r_.RespondsToSelector(objc.Sel("setVertexSamplerStates:lodMinClamps:lodMaxClamps:withRange:"))
}

// Assigns multiple sampler states and clamp values to a range of entries in the vertex shader argument table. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1516322-setvertexsamplerstates?language=objc
func (r_ RenderCommandEncoderWrapper) SetVertexSamplerStatesLodMinClampsLodMaxClampsWithRange(samplers PSamplerState, lodMinClamps *float64, lodMaxClamps *float64, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLSamplerState", samplers)
	objc.Call[objc.Void](r_, objc.Sel("setVertexSamplerStates:lodMinClamps:lodMaxClamps:withRange:"), po0, lodMinClamps, lodMaxClamps, range_)
}

func (r_ RenderCommandEncoderWrapper) HasSetFrontFacingWinding() bool {
	return r_.RespondsToSelector(objc.Sel("setFrontFacingWinding:"))
}

// Configures which face of a primitive, such as a triangle, is the front. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515499-setfrontfacingwinding?language=objc
func (r_ RenderCommandEncoderWrapper) SetFrontFacingWinding(frontFacingWinding Winding) {
	objc.Call[objc.Void](r_, objc.Sel("setFrontFacingWinding:"), frontFacingWinding)
}

func (r_ RenderCommandEncoderWrapper) HasDrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCount() bool {
	return r_.RespondsToSelector(objc.Sel("drawIndexedPrimitives:indexCount:indexType:indexBuffer:indexBufferOffset:instanceCount:"))
}

// Encodes a draw command that renders multiple instances of a geometric primitive with indexed vertices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515699-drawindexedprimitives?language=objc
func (r_ RenderCommandEncoderWrapper) DrawIndexedPrimitivesIndexCountIndexTypeIndexBufferIndexBufferOffsetInstanceCount(primitiveType PrimitiveType, indexCount uint, indexType IndexType, indexBuffer PBuffer, indexBufferOffset uint, instanceCount uint) {
	po3 := objc.WrapAsProtocol("MTLBuffer", indexBuffer)
	objc.Call[objc.Void](r_, objc.Sel("drawIndexedPrimitives:indexCount:indexType:indexBuffer:indexBufferOffset:instanceCount:"), primitiveType, indexCount, indexType, po3, indexBufferOffset, instanceCount)
}

func (r_ RenderCommandEncoderWrapper) HasSetStencilReferenceValue() bool {
	return r_.RespondsToSelector(objc.Sel("setStencilReferenceValue:"))
}

// Configures the same comparison value for front- and back-facing primitives. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1515697-setstencilreferencevalue?language=objc
func (r_ RenderCommandEncoderWrapper) SetStencilReferenceValue(referenceValue uint32) {
	objc.Call[objc.Void](r_, objc.Sel("setStencilReferenceValue:"), referenceValue)
}

func (r_ RenderCommandEncoderWrapper) HasWaitForFenceBeforeStages() bool {
	return r_.RespondsToSelector(objc.Sel("waitForFence:beforeStages:"))
}

// Encodes a command that instructs the GPU to pause before starting one or more stages of the render pass until a pass updates a fence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/1648378-waitforfence?language=objc
func (r_ RenderCommandEncoderWrapper) WaitForFenceBeforeStages(fence PFence, stages RenderStages) {
	po0 := objc.WrapAsProtocol("MTLFence", fence)
	objc.Call[objc.Void](r_, objc.Sel("waitForFence:beforeStages:"), po0, stages)
}

func (r_ RenderCommandEncoderWrapper) HasTileWidth() bool {
	return r_.RespondsToSelector(objc.Sel("tileWidth"))
}

// The width of the tiles, in pixels, for the render command encoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2866174-tilewidth?language=objc
func (r_ RenderCommandEncoderWrapper) TileWidth() uint {
	rv := objc.Call[uint](r_, objc.Sel("tileWidth"))
	return rv
}

func (r_ RenderCommandEncoderWrapper) HasTileHeight() bool {
	return r_.RespondsToSelector(objc.Sel("tileHeight"))
}

// The height of the tiles, in pixels, for the render command encoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlrendercommandencoder/2866164-tileheight?language=objc
func (r_ RenderCommandEncoderWrapper) TileHeight() uint {
	rv := objc.Call[uint](r_, objc.Sel("tileHeight"))
	return rv
}
