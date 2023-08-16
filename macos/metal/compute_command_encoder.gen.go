// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// An object for encoding commands in a compute pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder?language=objc
type PComputeCommandEncoder interface {
	// optional
	SetThreadgroupMemoryLengthAtIndex(length uint, index uint)
	HasSetThreadgroupMemoryLengthAtIndex() bool

	// optional
	SetSamplerStateAtIndex(sampler SamplerStateWrapper, index uint)
	HasSetSamplerStateAtIndex() bool

	// optional
	SetComputePipelineState(state ComputePipelineStateWrapper)
	HasSetComputePipelineState() bool

	// optional
	SetStageInRegionWithIndirectBufferIndirectBufferOffset(indirectBuffer BufferWrapper, indirectBufferOffset uint)
	HasSetStageInRegionWithIndirectBufferIndirectBufferOffset() bool

	// optional
	SetBuffersOffsetsWithRange(buffers BufferWrapper, offsets *uint, range_ foundation.Range)
	HasSetBuffersOffsetsWithRange() bool

	// optional
	UseResourceUsage(resource ResourceWrapper, usage ResourceUsage)
	HasUseResourceUsage() bool

	// optional
	SetVisibleFunctionTableAtBufferIndex(visibleFunctionTable VisibleFunctionTableWrapper, bufferIndex uint)
	HasSetVisibleFunctionTableAtBufferIndex() bool

	// optional
	SetBytesLengthAtIndex(bytes unsafe.Pointer, length uint, index uint)
	HasSetBytesLengthAtIndex() bool

	// optional
	SetSamplerStatesWithRange(samplers SamplerStateWrapper, range_ foundation.Range)
	HasSetSamplerStatesWithRange() bool

	// optional
	MemoryBarrierWithResourcesCount(resources ResourceWrapper, count uint)
	HasMemoryBarrierWithResourcesCount() bool

	// optional
	SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer CounterSampleBufferWrapper, sampleIndex uint, barrier bool)
	HasSampleCountersInBufferAtSampleIndexWithBarrier() bool

	// optional
	SetAccelerationStructureAtBufferIndex(accelerationStructure AccelerationStructureWrapper, bufferIndex uint)
	HasSetAccelerationStructureAtBufferIndex() bool

	// optional
	SetStageInRegion(region Region)
	HasSetStageInRegion() bool

	// optional
	SetBufferOffsetAtIndex(offset uint, index uint)
	HasSetBufferOffsetAtIndex() bool

	// optional
	DispatchThreadgroupsThreadsPerThreadgroup(threadgroupsPerGrid Size, threadsPerThreadgroup Size)
	HasDispatchThreadgroupsThreadsPerThreadgroup() bool

	// optional
	ExecuteCommandsInBufferWithRange(indirectCommandBuffer IndirectCommandBufferWrapper, executionRange foundation.Range)
	HasExecuteCommandsInBufferWithRange() bool

	// optional
	SetVisibleFunctionTablesWithBufferRange(visibleFunctionTables VisibleFunctionTableWrapper, range_ foundation.Range)
	HasSetVisibleFunctionTablesWithBufferRange() bool

	// optional
	SetTexturesWithRange(textures TextureWrapper, range_ foundation.Range)
	HasSetTexturesWithRange() bool

	// optional
	MemoryBarrierWithScope(scope BarrierScope)
	HasMemoryBarrierWithScope() bool

	// optional
	DispatchThreadsThreadsPerThreadgroup(threadsPerGrid Size, threadsPerThreadgroup Size)
	HasDispatchThreadsThreadsPerThreadgroup() bool

	// optional
	SetIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable IntersectionFunctionTableWrapper, bufferIndex uint)
	HasSetIntersectionFunctionTableAtBufferIndex() bool

	// optional
	SetTextureAtIndex(texture TextureWrapper, index uint)
	HasSetTextureAtIndex() bool

	// optional
	SetIntersectionFunctionTablesWithBufferRange(intersectionFunctionTables IntersectionFunctionTableWrapper, range_ foundation.Range)
	HasSetIntersectionFunctionTablesWithBufferRange() bool

	// optional
	DispatchThreadgroupsWithIndirectBufferIndirectBufferOffsetThreadsPerThreadgroup(indirectBuffer BufferWrapper, indirectBufferOffset uint, threadsPerThreadgroup Size)
	HasDispatchThreadgroupsWithIndirectBufferIndirectBufferOffsetThreadsPerThreadgroup() bool

	// optional
	UseHeap(heap HeapWrapper)
	HasUseHeap() bool

	// optional
	SetBufferOffsetAtIndex_(buffer BufferWrapper, offset uint, index uint)
	HasSetBufferOffsetAtIndex_() bool

	// optional
	UseResourcesCountUsage(resources ResourceWrapper, count uint, usage ResourceUsage)
	HasUseResourcesCountUsage() bool

	// optional
	UseHeapsCount(heaps HeapWrapper, count uint)
	HasUseHeapsCount() bool

	// optional
	UpdateFence(fence FenceWrapper)
	HasUpdateFence() bool

	// optional
	SetImageblockWidthHeight(width uint, height uint)
	HasSetImageblockWidthHeight() bool

	// optional
	WaitForFence(fence FenceWrapper)
	HasWaitForFence() bool

	// optional
	DispatchType() DispatchType
	HasDispatchType() bool
}

// A concrete type wrapper for the [PComputeCommandEncoder] protocol.
type ComputeCommandEncoderWrapper struct {
	objc.Object
}

func (c_ ComputeCommandEncoderWrapper) HasSetThreadgroupMemoryLengthAtIndex() bool {
	return c_.RespondsToSelector(objc.Sel("setThreadgroupMemoryLength:atIndex:"))
}

// Sets the size of a block of threadgroup memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/1443142-setthreadgroupmemorylength?language=objc
func (c_ ComputeCommandEncoderWrapper) SetThreadgroupMemoryLengthAtIndex(length uint, index uint) {
	objc.Call[objc.Void](c_, objc.Sel("setThreadgroupMemoryLength:atIndex:"), length, index)
}

func (c_ ComputeCommandEncoderWrapper) HasSetSamplerStateAtIndex() bool {
	return c_.RespondsToSelector(objc.Sel("setSamplerState:atIndex:"))
}

// Sets a sampler for the compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/1443144-setsamplerstate?language=objc
func (c_ ComputeCommandEncoderWrapper) SetSamplerStateAtIndex(sampler PSamplerState, index uint) {
	po0 := objc.WrapAsProtocol("MTLSamplerState", sampler)
	objc.Call[objc.Void](c_, objc.Sel("setSamplerState:atIndex:"), po0, index)
}

func (c_ ComputeCommandEncoderWrapper) HasSetComputePipelineState() bool {
	return c_.RespondsToSelector(objc.Sel("setComputePipelineState:"))
}

// Sets the current compute pipeline state object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/1443140-setcomputepipelinestate?language=objc
func (c_ ComputeCommandEncoderWrapper) SetComputePipelineState(state PComputePipelineState) {
	po0 := objc.WrapAsProtocol("MTLComputePipelineState", state)
	objc.Call[objc.Void](c_, objc.Sel("setComputePipelineState:"), po0)
}

func (c_ ComputeCommandEncoderWrapper) HasSetStageInRegionWithIndirectBufferIndirectBufferOffset() bool {
	return c_.RespondsToSelector(objc.Sel("setStageInRegionWithIndirectBuffer:indirectBufferOffset:"))
}

// Sets the region of the stage-in attributes to apply to the compute kernel using an indirect buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/2966554-setstageinregionwithindirectbuff?language=objc
func (c_ ComputeCommandEncoderWrapper) SetStageInRegionWithIndirectBufferIndirectBufferOffset(indirectBuffer PBuffer, indirectBufferOffset uint) {
	po0 := objc.WrapAsProtocol("MTLBuffer", indirectBuffer)
	objc.Call[objc.Void](c_, objc.Sel("setStageInRegionWithIndirectBuffer:indirectBufferOffset:"), po0, indirectBufferOffset)
}

func (c_ ComputeCommandEncoderWrapper) HasSetBuffersOffsetsWithRange() bool {
	return c_.RespondsToSelector(objc.Sel("setBuffers:offsets:withRange:"))
}

// Sets an array of buffers for the compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/1443134-setbuffers?language=objc
func (c_ ComputeCommandEncoderWrapper) SetBuffersOffsetsWithRange(buffers PBuffer, offsets *uint, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffers)
	objc.Call[objc.Void](c_, objc.Sel("setBuffers:offsets:withRange:"), po0, offsets, range_)
}

func (c_ ComputeCommandEncoderWrapper) HasUseResourceUsage() bool {
	return c_.RespondsToSelector(objc.Sel("useResource:usage:"))
}

// Specifies that a resource in an argument buffer can be safely used by a compute pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/2866548-useresource?language=objc
func (c_ ComputeCommandEncoderWrapper) UseResourceUsage(resource PResource, usage ResourceUsage) {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	objc.Call[objc.Void](c_, objc.Sel("useResource:usage:"), po0, usage)
}

func (c_ ComputeCommandEncoderWrapper) HasSetVisibleFunctionTableAtBufferIndex() bool {
	return c_.RespondsToSelector(objc.Sel("setVisibleFunctionTable:atBufferIndex:"))
}

// Sets a visible function table for the compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/3566539-setvisiblefunctiontable?language=objc
func (c_ ComputeCommandEncoderWrapper) SetVisibleFunctionTableAtBufferIndex(visibleFunctionTable PVisibleFunctionTable, bufferIndex uint) {
	po0 := objc.WrapAsProtocol("MTLVisibleFunctionTable", visibleFunctionTable)
	objc.Call[objc.Void](c_, objc.Sel("setVisibleFunctionTable:atBufferIndex:"), po0, bufferIndex)
}

func (c_ ComputeCommandEncoderWrapper) HasSetBytesLengthAtIndex() bool {
	return c_.RespondsToSelector(objc.Sel("setBytes:length:atIndex:"))
}

// Sets a block of data for the compute shader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/1443159-setbytes?language=objc
func (c_ ComputeCommandEncoderWrapper) SetBytesLengthAtIndex(bytes unsafe.Pointer, length uint, index uint) {
	objc.Call[objc.Void](c_, objc.Sel("setBytes:length:atIndex:"), bytes, length, index)
}

func (c_ ComputeCommandEncoderWrapper) HasSetSamplerStatesWithRange() bool {
	return c_.RespondsToSelector(objc.Sel("setSamplerStates:withRange:"))
}

// Sets multiple samplers for the compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/1443155-setsamplerstates?language=objc
func (c_ ComputeCommandEncoderWrapper) SetSamplerStatesWithRange(samplers PSamplerState, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLSamplerState", samplers)
	objc.Call[objc.Void](c_, objc.Sel("setSamplerStates:withRange:"), po0, range_)
}

func (c_ ComputeCommandEncoderWrapper) HasMemoryBarrierWithResourcesCount() bool {
	return c_.RespondsToSelector(objc.Sel("memoryBarrierWithResources:count:"))
}

// Encodes a barrier so that changes to a set of resources made by commands encoded before the barrier are completed before further commands are executed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/2966552-memorybarrierwithresources?language=objc
func (c_ ComputeCommandEncoderWrapper) MemoryBarrierWithResourcesCount(resources PResource, count uint) {
	po0 := objc.WrapAsProtocol("MTLResource", resources)
	objc.Call[objc.Void](c_, objc.Sel("memoryBarrierWithResources:count:"), po0, count)
}

func (c_ ComputeCommandEncoderWrapper) HasSampleCountersInBufferAtSampleIndexWithBarrier() bool {
	return c_.RespondsToSelector(objc.Sel("sampleCountersInBuffer:atSampleIndex:withBarrier:"))
}

// Encodes a command to sample hardware counters at this point in the compute pass and store the samples into a counter sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/3194349-samplecountersinbuffer?language=objc
func (c_ ComputeCommandEncoderWrapper) SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer PCounterSampleBuffer, sampleIndex uint, barrier bool) {
	po0 := objc.WrapAsProtocol("MTLCounterSampleBuffer", sampleBuffer)
	objc.Call[objc.Void](c_, objc.Sel("sampleCountersInBuffer:atSampleIndex:withBarrier:"), po0, sampleIndex, barrier)
}

func (c_ ComputeCommandEncoderWrapper) HasSetAccelerationStructureAtBufferIndex() bool {
	return c_.RespondsToSelector(objc.Sel("setAccelerationStructure:atBufferIndex:"))
}

// Sets an acceleration structure for the compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/3553958-setaccelerationstructure?language=objc
func (c_ ComputeCommandEncoderWrapper) SetAccelerationStructureAtBufferIndex(accelerationStructure PAccelerationStructure, bufferIndex uint) {
	po0 := objc.WrapAsProtocol("MTLAccelerationStructure", accelerationStructure)
	objc.Call[objc.Void](c_, objc.Sel("setAccelerationStructure:atBufferIndex:"), po0, bufferIndex)
}

func (c_ ComputeCommandEncoderWrapper) HasSetStageInRegion() bool {
	return c_.RespondsToSelector(objc.Sel("setStageInRegion:"))
}

// Sets the region of the stage-in attributes to apply to the compute kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/2097047-setstageinregion?language=objc
func (c_ ComputeCommandEncoderWrapper) SetStageInRegion(region Region) {
	objc.Call[objc.Void](c_, objc.Sel("setStageInRegion:"), region)
}

func (c_ ComputeCommandEncoderWrapper) HasSetBufferOffsetAtIndex() bool {
	return c_.RespondsToSelector(objc.Sel("setBufferOffset:atIndex:"))
}

// Sets where the data begins in a buffer already bound to the compute shader. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/1443146-setbufferoffset?language=objc
func (c_ ComputeCommandEncoderWrapper) SetBufferOffsetAtIndex(offset uint, index uint) {
	objc.Call[objc.Void](c_, objc.Sel("setBufferOffset:atIndex:"), offset, index)
}

func (c_ ComputeCommandEncoderWrapper) HasDispatchThreadgroupsThreadsPerThreadgroup() bool {
	return c_.RespondsToSelector(objc.Sel("dispatchThreadgroups:threadsPerThreadgroup:"))
}

// Encodes a compute command using a grid aligned to threadgroup boundaries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/1443138-dispatchthreadgroups?language=objc
func (c_ ComputeCommandEncoderWrapper) DispatchThreadgroupsThreadsPerThreadgroup(threadgroupsPerGrid Size, threadsPerThreadgroup Size) {
	objc.Call[objc.Void](c_, objc.Sel("dispatchThreadgroups:threadsPerThreadgroup:"), threadgroupsPerGrid, threadsPerThreadgroup)
}

func (c_ ComputeCommandEncoderWrapper) HasExecuteCommandsInBufferWithRange() bool {
	return c_.RespondsToSelector(objc.Sel("executeCommandsInBuffer:withRange:"))
}

// Encodes a command to execute commands in an indirect command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/2967419-executecommandsinbuffer?language=objc
func (c_ ComputeCommandEncoderWrapper) ExecuteCommandsInBufferWithRange(indirectCommandBuffer PIndirectCommandBuffer, executionRange foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLIndirectCommandBuffer", indirectCommandBuffer)
	objc.Call[objc.Void](c_, objc.Sel("executeCommandsInBuffer:withRange:"), po0, executionRange)
}

func (c_ ComputeCommandEncoderWrapper) HasSetVisibleFunctionTablesWithBufferRange() bool {
	return c_.RespondsToSelector(objc.Sel("setVisibleFunctionTables:withBufferRange:"))
}

// Sets an array of visible function tables for the compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/3566540-setvisiblefunctiontables?language=objc
func (c_ ComputeCommandEncoderWrapper) SetVisibleFunctionTablesWithBufferRange(visibleFunctionTables PVisibleFunctionTable, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLVisibleFunctionTable", visibleFunctionTables)
	objc.Call[objc.Void](c_, objc.Sel("setVisibleFunctionTables:withBufferRange:"), po0, range_)
}

func (c_ ComputeCommandEncoderWrapper) HasSetTexturesWithRange() bool {
	return c_.RespondsToSelector(objc.Sel("setTextures:withRange:"))
}

// Sets an array of textures for the compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/1443148-settextures?language=objc
func (c_ ComputeCommandEncoderWrapper) SetTexturesWithRange(textures PTexture, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLTexture", textures)
	objc.Call[objc.Void](c_, objc.Sel("setTextures:withRange:"), po0, range_)
}

func (c_ ComputeCommandEncoderWrapper) HasMemoryBarrierWithScope() bool {
	return c_.RespondsToSelector(objc.Sel("memoryBarrierWithScope:"))
}

// Encodes a barrier so that changes to a set of resource types made by commands encoded before the barrier are completed before further commands are executed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/2966553-memorybarrierwithscope?language=objc
func (c_ ComputeCommandEncoderWrapper) MemoryBarrierWithScope(scope BarrierScope) {
	objc.Call[objc.Void](c_, objc.Sel("memoryBarrierWithScope:"), scope)
}

func (c_ ComputeCommandEncoderWrapper) HasDispatchThreadsThreadsPerThreadgroup() bool {
	return c_.RespondsToSelector(objc.Sel("dispatchThreads:threadsPerThreadgroup:"))
}

// Encodes a compute command using an arbitrarily sized grid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/2866532-dispatchthreads?language=objc
func (c_ ComputeCommandEncoderWrapper) DispatchThreadsThreadsPerThreadgroup(threadsPerGrid Size, threadsPerThreadgroup Size) {
	objc.Call[objc.Void](c_, objc.Sel("dispatchThreads:threadsPerThreadgroup:"), threadsPerGrid, threadsPerThreadgroup)
}

func (c_ ComputeCommandEncoderWrapper) HasSetIntersectionFunctionTableAtBufferIndex() bool {
	return c_.RespondsToSelector(objc.Sel("setIntersectionFunctionTable:atBufferIndex:"))
}

// Sets a intersection function table for the compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/3566537-setintersectionfunctiontable?language=objc
func (c_ ComputeCommandEncoderWrapper) SetIntersectionFunctionTableAtBufferIndex(intersectionFunctionTable PIntersectionFunctionTable, bufferIndex uint) {
	po0 := objc.WrapAsProtocol("MTLIntersectionFunctionTable", intersectionFunctionTable)
	objc.Call[objc.Void](c_, objc.Sel("setIntersectionFunctionTable:atBufferIndex:"), po0, bufferIndex)
}

func (c_ ComputeCommandEncoderWrapper) HasSetTextureAtIndex() bool {
	return c_.RespondsToSelector(objc.Sel("setTexture:atIndex:"))
}

// Sets a texture for the compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/1443130-settexture?language=objc
func (c_ ComputeCommandEncoderWrapper) SetTextureAtIndex(texture PTexture, index uint) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	objc.Call[objc.Void](c_, objc.Sel("setTexture:atIndex:"), po0, index)
}

func (c_ ComputeCommandEncoderWrapper) HasSetIntersectionFunctionTablesWithBufferRange() bool {
	return c_.RespondsToSelector(objc.Sel("setIntersectionFunctionTables:withBufferRange:"))
}

// Sets an array of intersection function tables for the compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/3566538-setintersectionfunctiontables?language=objc
func (c_ ComputeCommandEncoderWrapper) SetIntersectionFunctionTablesWithBufferRange(intersectionFunctionTables PIntersectionFunctionTable, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLIntersectionFunctionTable", intersectionFunctionTables)
	objc.Call[objc.Void](c_, objc.Sel("setIntersectionFunctionTables:withBufferRange:"), po0, range_)
}

func (c_ ComputeCommandEncoderWrapper) HasDispatchThreadgroupsWithIndirectBufferIndirectBufferOffsetThreadsPerThreadgroup() bool {
	return c_.RespondsToSelector(objc.Sel("dispatchThreadgroupsWithIndirectBuffer:indirectBufferOffset:threadsPerThreadgroup:"))
}

// Encodes a dispatch call for a compute pass, using an indirect buffer that defines the size of a grid aligned to threadgroup boundaries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/1443157-dispatchthreadgroupswithindirect?language=objc
func (c_ ComputeCommandEncoderWrapper) DispatchThreadgroupsWithIndirectBufferIndirectBufferOffsetThreadsPerThreadgroup(indirectBuffer PBuffer, indirectBufferOffset uint, threadsPerThreadgroup Size) {
	po0 := objc.WrapAsProtocol("MTLBuffer", indirectBuffer)
	objc.Call[objc.Void](c_, objc.Sel("dispatchThreadgroupsWithIndirectBuffer:indirectBufferOffset:threadsPerThreadgroup:"), po0, indirectBufferOffset, threadsPerThreadgroup)
}

func (c_ ComputeCommandEncoderWrapper) HasUseHeap() bool {
	return c_.RespondsToSelector(objc.Sel("useHeap:"))
}

// Specifies that a heap containing resources in an argument buffer can be safely used by a compute pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/2866530-useheap?language=objc
func (c_ ComputeCommandEncoderWrapper) UseHeap(heap PHeap) {
	po0 := objc.WrapAsProtocol("MTLHeap", heap)
	objc.Call[objc.Void](c_, objc.Sel("useHeap:"), po0)
}

func (c_ ComputeCommandEncoderWrapper) HasSetBufferOffsetAtIndex_() bool {
	return c_.RespondsToSelector(objc.Sel("setBuffer:offset:atIndex:"))
}

// Sets a buffer for the compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/1443126-setbuffer?language=objc
func (c_ ComputeCommandEncoderWrapper) SetBufferOffsetAtIndex_(buffer PBuffer, offset uint, index uint) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	objc.Call[objc.Void](c_, objc.Sel("setBuffer:offset:atIndex:"), po0, offset, index)
}

func (c_ ComputeCommandEncoderWrapper) HasUseResourcesCountUsage() bool {
	return c_.RespondsToSelector(objc.Sel("useResources:count:usage:"))
}

// Specifies that an array of resources in an argument buffer can be safely used by a compute pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/2866561-useresources?language=objc
func (c_ ComputeCommandEncoderWrapper) UseResourcesCountUsage(resources PResource, count uint, usage ResourceUsage) {
	po0 := objc.WrapAsProtocol("MTLResource", resources)
	objc.Call[objc.Void](c_, objc.Sel("useResources:count:usage:"), po0, count, usage)
}

func (c_ ComputeCommandEncoderWrapper) HasUseHeapsCount() bool {
	return c_.RespondsToSelector(objc.Sel("useHeaps:count:"))
}

// Specifies that an array of heaps containing resources in an argument buffer can be safely used by a compute pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/2866538-useheaps?language=objc
func (c_ ComputeCommandEncoderWrapper) UseHeapsCount(heaps PHeap, count uint) {
	po0 := objc.WrapAsProtocol("MTLHeap", heaps)
	objc.Call[objc.Void](c_, objc.Sel("useHeaps:count:"), po0, count)
}

func (c_ ComputeCommandEncoderWrapper) HasUpdateFence() bool {
	return c_.RespondsToSelector(objc.Sel("updateFence:"))
}

// Tells the GPU to update the fence after all commands encoded by the compute command encoder have finished executing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/1649789-updatefence?language=objc
func (c_ ComputeCommandEncoderWrapper) UpdateFence(fence PFence) {
	po0 := objc.WrapAsProtocol("MTLFence", fence)
	objc.Call[objc.Void](c_, objc.Sel("updateFence:"), po0)
}

func (c_ ComputeCommandEncoderWrapper) HasSetImageblockWidthHeight() bool {
	return c_.RespondsToSelector(objc.Sel("setImageblockWidth:height:"))
}

// Sets the size, in pixels, of the imageblock. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/2928282-setimageblockwidth?language=objc
func (c_ ComputeCommandEncoderWrapper) SetImageblockWidthHeight(width uint, height uint) {
	objc.Call[objc.Void](c_, objc.Sel("setImageblockWidth:height:"), width, height)
}

func (c_ ComputeCommandEncoderWrapper) HasWaitForFence() bool {
	return c_.RespondsToSelector(objc.Sel("waitForFence:"))
}

// Tells the GPU to wait until the fence is updated before executing any commands encoded by the compute command encoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/1649790-waitforfence?language=objc
func (c_ ComputeCommandEncoderWrapper) WaitForFence(fence PFence) {
	po0 := objc.WrapAsProtocol("MTLFence", fence)
	objc.Call[objc.Void](c_, objc.Sel("waitForFence:"), po0)
}

func (c_ ComputeCommandEncoderWrapper) HasDispatchType() bool {
	return c_.RespondsToSelector(objc.Sel("dispatchType"))
}

// The strategy to use when dispatching commands encoded by the compute command encoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcomputecommandencoder/2966550-dispatchtype?language=objc
func (c_ ComputeCommandEncoderWrapper) DispatchType() DispatchType {
	rv := objc.Call[DispatchType](c_, objc.Sel("dispatchType"))
	return rv
}
