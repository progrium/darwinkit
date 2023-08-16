// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// An interface you can use to encode GPU commands that copy and modify the underlying memory of various Metal resources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitcommandencoder?language=objc
type PBlitCommandEncoder interface {
	// optional
	CopyIndirectCommandBufferSourceRangeDestinationDestinationIndex(source IndirectCommandBufferWrapper, sourceRange foundation.Range, destination IndirectCommandBufferWrapper, destinationIndex uint)
	HasCopyIndirectCommandBufferSourceRangeDestinationDestinationIndex() bool

	// optional
	SynchronizeTextureSliceLevel(texture TextureWrapper, slice uint, level uint)
	HasSynchronizeTextureSliceLevel() bool

	// optional
	ResolveCountersInRangeDestinationBufferDestinationOffset(sampleBuffer CounterSampleBufferWrapper, range_ foundation.Range, destinationBuffer BufferWrapper, destinationOffset uint)
	HasResolveCountersInRangeDestinationBufferDestinationOffset() bool

	// optional
	OptimizeContentsForGPUAccess(texture TextureWrapper)
	HasOptimizeContentsForGPUAccess() bool

	// optional
	SynchronizeResource(resource ResourceWrapper)
	HasSynchronizeResource() bool

	// optional
	SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer CounterSampleBufferWrapper, sampleIndex uint, barrier bool)
	HasSampleCountersInBufferAtSampleIndexWithBarrier() bool

	// optional
	ResetTextureAccessCountersRegionMipLevelSlice(texture TextureWrapper, region Region, mipLevel uint, slice uint)
	HasResetTextureAccessCountersRegionMipLevelSlice() bool

	// optional
	CopyFromBufferSourceOffsetToBufferDestinationOffsetSize(sourceBuffer BufferWrapper, sourceOffset uint, destinationBuffer BufferWrapper, destinationOffset uint, size uint)
	HasCopyFromBufferSourceOffsetToBufferDestinationOffsetSize() bool

	// optional
	FillBufferRangeValue(buffer BufferWrapper, range_ foundation.Range, value uint8)
	HasFillBufferRangeValue() bool

	// optional
	OptimizeContentsForCPUAccess(texture TextureWrapper)
	HasOptimizeContentsForCPUAccess() bool

	// optional
	OptimizeIndirectCommandBufferWithRange(indirectCommandBuffer IndirectCommandBufferWrapper, range_ foundation.Range)
	HasOptimizeIndirectCommandBufferWithRange() bool

	// optional
	ResetCommandsInBufferWithRange(buffer IndirectCommandBufferWrapper, range_ foundation.Range)
	HasResetCommandsInBufferWithRange() bool

	// optional
	GetTextureAccessCountersRegionMipLevelSliceResetCountersCountersBufferCountersBufferOffset(texture TextureWrapper, region Region, mipLevel uint, slice uint, resetCounters bool, countersBuffer BufferWrapper, countersBufferOffset uint)
	HasGetTextureAccessCountersRegionMipLevelSliceResetCountersCountersBufferCountersBufferOffset() bool

	// optional
	GenerateMipmapsForTexture(texture TextureWrapper)
	HasGenerateMipmapsForTexture() bool

	// optional
	UpdateFence(fence FenceWrapper)
	HasUpdateFence() bool

	// optional
	CopyFromTextureToTexture(sourceTexture TextureWrapper, destinationTexture TextureWrapper)
	HasCopyFromTextureToTexture() bool

	// optional
	WaitForFence(fence FenceWrapper)
	HasWaitForFence() bool
}

// A concrete type wrapper for the [PBlitCommandEncoder] protocol.
type BlitCommandEncoderWrapper struct {
	objc.Object
}

func (b_ BlitCommandEncoderWrapper) HasCopyIndirectCommandBufferSourceRangeDestinationDestinationIndex() bool {
	return b_.RespondsToSelector(objc.Sel("copyIndirectCommandBuffer:sourceRange:destination:destinationIndex:"))
}

// Encodes a command that copies commands from one indirect command buffer into another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/2967412-copyindirectcommandbuffer?language=objc
func (b_ BlitCommandEncoderWrapper) CopyIndirectCommandBufferSourceRangeDestinationDestinationIndex(source PIndirectCommandBuffer, sourceRange foundation.Range, destination PIndirectCommandBuffer, destinationIndex uint) {
	po0 := objc.WrapAsProtocol("MTLIndirectCommandBuffer", source)
	po2 := objc.WrapAsProtocol("MTLIndirectCommandBuffer", destination)
	objc.Call[objc.Void](b_, objc.Sel("copyIndirectCommandBuffer:sourceRange:destination:destinationIndex:"), po0, sourceRange, po2, destinationIndex)
}

func (b_ BlitCommandEncoderWrapper) HasSynchronizeTextureSliceLevel() bool {
	return b_.RespondsToSelector(objc.Sel("synchronizeTexture:slice:level:"))
}

// Encodes a command that synchronizes a part of the CPU’s copy of a texture so that it matches the GPU’s copy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/1400757-synchronizetexture?language=objc
func (b_ BlitCommandEncoderWrapper) SynchronizeTextureSliceLevel(texture PTexture, slice uint, level uint) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	objc.Call[objc.Void](b_, objc.Sel("synchronizeTexture:slice:level:"), po0, slice, level)
}

func (b_ BlitCommandEncoderWrapper) HasResolveCountersInRangeDestinationBufferDestinationOffset() bool {
	return b_.RespondsToSelector(objc.Sel("resolveCounters:inRange:destinationBuffer:destinationOffset:"))
}

// Encodes a command that resolves the data from the samples in a sample counter buffer and stores the results into a buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/3194347-resolvecounters?language=objc
func (b_ BlitCommandEncoderWrapper) ResolveCountersInRangeDestinationBufferDestinationOffset(sampleBuffer PCounterSampleBuffer, range_ foundation.Range, destinationBuffer PBuffer, destinationOffset uint) {
	po0 := objc.WrapAsProtocol("MTLCounterSampleBuffer", sampleBuffer)
	po2 := objc.WrapAsProtocol("MTLBuffer", destinationBuffer)
	objc.Call[objc.Void](b_, objc.Sel("resolveCounters:inRange:destinationBuffer:destinationOffset:"), po0, range_, po2, destinationOffset)
}

func (b_ BlitCommandEncoderWrapper) HasOptimizeContentsForGPUAccess() bool {
	return b_.RespondsToSelector(objc.Sel("optimizeContentsForGPUAccess:"))
}

// Encodes a command that improves the performance of the GPU’s accesses to a texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/2966539-optimizecontentsforgpuaccess?language=objc
func (b_ BlitCommandEncoderWrapper) OptimizeContentsForGPUAccess(texture PTexture) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	objc.Call[objc.Void](b_, objc.Sel("optimizeContentsForGPUAccess:"), po0)
}

func (b_ BlitCommandEncoderWrapper) HasSynchronizeResource() bool {
	return b_.RespondsToSelector(objc.Sel("synchronizeResource:"))
}

// Encodes a command that synchronizes the CPU’s copy of a managed resource, such as a buffer or texture, so that it matches the GPU’s copy. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/1400775-synchronizeresource?language=objc
func (b_ BlitCommandEncoderWrapper) SynchronizeResource(resource PResource) {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	objc.Call[objc.Void](b_, objc.Sel("synchronizeResource:"), po0)
}

func (b_ BlitCommandEncoderWrapper) HasSampleCountersInBufferAtSampleIndexWithBarrier() bool {
	return b_.RespondsToSelector(objc.Sel("sampleCountersInBuffer:atSampleIndex:withBarrier:"))
}

// Encodes a command that samples the GPU’s hardware counters during a blit pass and stores the data in a counter sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/3194348-samplecountersinbuffer?language=objc
func (b_ BlitCommandEncoderWrapper) SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer PCounterSampleBuffer, sampleIndex uint, barrier bool) {
	po0 := objc.WrapAsProtocol("MTLCounterSampleBuffer", sampleBuffer)
	objc.Call[objc.Void](b_, objc.Sel("sampleCountersInBuffer:atSampleIndex:withBarrier:"), po0, sampleIndex, barrier)
}

func (b_ BlitCommandEncoderWrapper) HasResetTextureAccessCountersRegionMipLevelSlice() bool {
	return b_.RespondsToSelector(objc.Sel("resetTextureAccessCounters:region:mipLevel:slice:"))
}

// Encodes a command that resets a sparse texture’s access data for a specific region, mipmap level, and slice. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/3043911-resettextureaccesscounters?language=objc
func (b_ BlitCommandEncoderWrapper) ResetTextureAccessCountersRegionMipLevelSlice(texture PTexture, region Region, mipLevel uint, slice uint) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	objc.Call[objc.Void](b_, objc.Sel("resetTextureAccessCounters:region:mipLevel:slice:"), po0, region, mipLevel, slice)
}

func (b_ BlitCommandEncoderWrapper) HasCopyFromBufferSourceOffsetToBufferDestinationOffsetSize() bool {
	return b_.RespondsToSelector(objc.Sel("copyFromBuffer:sourceOffset:toBuffer:destinationOffset:size:"))
}

// Encodes a command that copies data from one buffer into another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/1400767-copyfrombuffer?language=objc
func (b_ BlitCommandEncoderWrapper) CopyFromBufferSourceOffsetToBufferDestinationOffsetSize(sourceBuffer PBuffer, sourceOffset uint, destinationBuffer PBuffer, destinationOffset uint, size uint) {
	po0 := objc.WrapAsProtocol("MTLBuffer", sourceBuffer)
	po2 := objc.WrapAsProtocol("MTLBuffer", destinationBuffer)
	objc.Call[objc.Void](b_, objc.Sel("copyFromBuffer:sourceOffset:toBuffer:destinationOffset:size:"), po0, sourceOffset, po2, destinationOffset, size)
}

func (b_ BlitCommandEncoderWrapper) HasFillBufferRangeValue() bool {
	return b_.RespondsToSelector(objc.Sel("fillBuffer:range:value:"))
}

// Encodes a command that fills a buffer with a constant value for each byte. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/1400761-fillbuffer?language=objc
func (b_ BlitCommandEncoderWrapper) FillBufferRangeValue(buffer PBuffer, range_ foundation.Range, value uint8) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	objc.Call[objc.Void](b_, objc.Sel("fillBuffer:range:value:"), po0, range_, value)
}

func (b_ BlitCommandEncoderWrapper) HasOptimizeContentsForCPUAccess() bool {
	return b_.RespondsToSelector(objc.Sel("optimizeContentsForCPUAccess:"))
}

// Encodes a command that improves the performance of the CPU’s accesses to a texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/2966537-optimizecontentsforcpuaccess?language=objc
func (b_ BlitCommandEncoderWrapper) OptimizeContentsForCPUAccess(texture PTexture) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	objc.Call[objc.Void](b_, objc.Sel("optimizeContentsForCPUAccess:"), po0)
}

func (b_ BlitCommandEncoderWrapper) HasOptimizeIndirectCommandBufferWithRange() bool {
	return b_.RespondsToSelector(objc.Sel("optimizeIndirectCommandBuffer:withRange:"))
}

// Encodes a command that can improve the performance of a range of commands within an indirect command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/2967413-optimizeindirectcommandbuffer?language=objc
func (b_ BlitCommandEncoderWrapper) OptimizeIndirectCommandBufferWithRange(indirectCommandBuffer PIndirectCommandBuffer, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLIndirectCommandBuffer", indirectCommandBuffer)
	objc.Call[objc.Void](b_, objc.Sel("optimizeIndirectCommandBuffer:withRange:"), po0, range_)
}

func (b_ BlitCommandEncoderWrapper) HasResetCommandsInBufferWithRange() bool {
	return b_.RespondsToSelector(objc.Sel("resetCommandsInBuffer:withRange:"))
}

// Encodes a command that resets a range of commands in an indirect command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/2967414-resetcommandsinbuffer?language=objc
func (b_ BlitCommandEncoderWrapper) ResetCommandsInBufferWithRange(buffer PIndirectCommandBuffer, range_ foundation.Range) {
	po0 := objc.WrapAsProtocol("MTLIndirectCommandBuffer", buffer)
	objc.Call[objc.Void](b_, objc.Sel("resetCommandsInBuffer:withRange:"), po0, range_)
}

func (b_ BlitCommandEncoderWrapper) HasGetTextureAccessCountersRegionMipLevelSliceResetCountersCountersBufferCountersBufferOffset() bool {
	return b_.RespondsToSelector(objc.Sel("getTextureAccessCounters:region:mipLevel:slice:resetCounters:countersBuffer:countersBufferOffset:"))
}

// Encodes a command that retrieves a sparse texture’s access data for a specific region, mipmap level, and slice. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/3043910-gettextureaccesscounters?language=objc
func (b_ BlitCommandEncoderWrapper) GetTextureAccessCountersRegionMipLevelSliceResetCountersCountersBufferCountersBufferOffset(texture PTexture, region Region, mipLevel uint, slice uint, resetCounters bool, countersBuffer PBuffer, countersBufferOffset uint) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	po5 := objc.WrapAsProtocol("MTLBuffer", countersBuffer)
	objc.Call[objc.Void](b_, objc.Sel("getTextureAccessCounters:region:mipLevel:slice:resetCounters:countersBuffer:countersBufferOffset:"), po0, region, mipLevel, slice, resetCounters, po5, countersBufferOffset)
}

func (b_ BlitCommandEncoderWrapper) HasGenerateMipmapsForTexture() bool {
	return b_.RespondsToSelector(objc.Sel("generateMipmapsForTexture:"))
}

// Encodes a command that generates mipmaps for a texture from the base mipmap level up to the highest mipmap level. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/1400748-generatemipmapsfortexture?language=objc
func (b_ BlitCommandEncoderWrapper) GenerateMipmapsForTexture(texture PTexture) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	objc.Call[objc.Void](b_, objc.Sel("generateMipmapsForTexture:"), po0)
}

func (b_ BlitCommandEncoderWrapper) HasUpdateFence() bool {
	return b_.RespondsToSelector(objc.Sel("updateFence:"))
}

// Encodes a command that instructs the GPU to update a fence, which can signal a pass that’s waiting for it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/1649359-updatefence?language=objc
func (b_ BlitCommandEncoderWrapper) UpdateFence(fence PFence) {
	po0 := objc.WrapAsProtocol("MTLFence", fence)
	objc.Call[objc.Void](b_, objc.Sel("updateFence:"), po0)
}

func (b_ BlitCommandEncoderWrapper) HasCopyFromTextureToTexture() bool {
	return b_.RespondsToSelector(objc.Sel("copyFromTexture:toTexture:"))
}

// Encodes a command that copies data from one texture to another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/3075383-copyfromtexture?language=objc
func (b_ BlitCommandEncoderWrapper) CopyFromTextureToTexture(sourceTexture PTexture, destinationTexture PTexture) {
	po0 := objc.WrapAsProtocol("MTLTexture", sourceTexture)
	po1 := objc.WrapAsProtocol("MTLTexture", destinationTexture)
	objc.Call[objc.Void](b_, objc.Sel("copyFromTexture:toTexture:"), po0, po1)
}

func (b_ BlitCommandEncoderWrapper) HasWaitForFence() bool {
	return b_.RespondsToSelector(objc.Sel("waitForFence:"))
}

// Encodes a command that instructs the GPU to wait until a pass updates a fence. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlblitcommandencoder/1649358-waitforfence?language=objc
func (b_ BlitCommandEncoderWrapper) WaitForFence(fence PFence) {
	po0 := objc.WrapAsProtocol("MTLFence", fence)
	objc.Call[objc.Void](b_, objc.Sel("waitForFence:"), po0)
}
