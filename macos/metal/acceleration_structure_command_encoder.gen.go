// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// An object for encoding commands that build or refit acceleration structures. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecommandencoder?language=objc
type PAccelerationStructureCommandEncoder interface {
	// optional
	UseResourceUsage(resource ResourceWrapper, usage ResourceUsage)
	HasUseResourceUsage() bool

	// optional
	CopyAndCompactAccelerationStructureToAccelerationStructure(sourceAccelerationStructure AccelerationStructureWrapper, destinationAccelerationStructure AccelerationStructureWrapper)
	HasCopyAndCompactAccelerationStructureToAccelerationStructure() bool

	// optional
	RefitAccelerationStructureDescriptorDestinationScratchBufferScratchBufferOffset(sourceAccelerationStructure AccelerationStructureWrapper, descriptor AccelerationStructureDescriptor, destinationAccelerationStructure AccelerationStructureWrapper, scratchBuffer BufferWrapper, scratchBufferOffset uint)
	HasRefitAccelerationStructureDescriptorDestinationScratchBufferScratchBufferOffset() bool

	// optional
	SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer CounterSampleBufferWrapper, sampleIndex uint, barrier bool)
	HasSampleCountersInBufferAtSampleIndexWithBarrier() bool

	// optional
	CopyAccelerationStructureToAccelerationStructure(sourceAccelerationStructure AccelerationStructureWrapper, destinationAccelerationStructure AccelerationStructureWrapper)
	HasCopyAccelerationStructureToAccelerationStructure() bool

	// optional
	WriteCompactedAccelerationStructureSizeToBufferOffset(accelerationStructure AccelerationStructureWrapper, buffer BufferWrapper, offset uint)
	HasWriteCompactedAccelerationStructureSizeToBufferOffset() bool

	// optional
	UseHeap(heap HeapWrapper)
	HasUseHeap() bool

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
	BuildAccelerationStructureDescriptorScratchBufferScratchBufferOffset(accelerationStructure AccelerationStructureWrapper, descriptor AccelerationStructureDescriptor, scratchBuffer BufferWrapper, scratchBufferOffset uint)
	HasBuildAccelerationStructureDescriptorScratchBufferScratchBufferOffset() bool

	// optional
	WaitForFence(fence FenceWrapper)
	HasWaitForFence() bool
}

// A concrete type wrapper for the [PAccelerationStructureCommandEncoder] protocol.
type AccelerationStructureCommandEncoderWrapper struct {
	objc.Object
}

func (a_ AccelerationStructureCommandEncoderWrapper) HasUseResourceUsage() bool {
	return a_.RespondsToSelector(objc.Sel("useResource:usage:"))
}

// Makes a resource available to the acceleration structure pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecommandencoder/3553904-useresource?language=objc
func (a_ AccelerationStructureCommandEncoderWrapper) UseResourceUsage(resource PResource, usage ResourceUsage) {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	objc.Call[objc.Void](a_, objc.Sel("useResource:usage:"), po0, usage)
}

func (a_ AccelerationStructureCommandEncoderWrapper) HasCopyAndCompactAccelerationStructureToAccelerationStructure() bool {
	return a_.RespondsToSelector(objc.Sel("copyAndCompactAccelerationStructure:toAccelerationStructure:"))
}

// Encodes a command to compact an acceleration structure’s data and copy it into a different acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecommandencoder/3553896-copyandcompactaccelerationstruct?language=objc
func (a_ AccelerationStructureCommandEncoderWrapper) CopyAndCompactAccelerationStructureToAccelerationStructure(sourceAccelerationStructure PAccelerationStructure, destinationAccelerationStructure PAccelerationStructure) {
	po0 := objc.WrapAsProtocol("MTLAccelerationStructure", sourceAccelerationStructure)
	po1 := objc.WrapAsProtocol("MTLAccelerationStructure", destinationAccelerationStructure)
	objc.Call[objc.Void](a_, objc.Sel("copyAndCompactAccelerationStructure:toAccelerationStructure:"), po0, po1)
}

func (a_ AccelerationStructureCommandEncoderWrapper) HasRefitAccelerationStructureDescriptorDestinationScratchBufferScratchBufferOffset() bool {
	return a_.RespondsToSelector(objc.Sel("refitAccelerationStructure:descriptor:destination:scratchBuffer:scratchBufferOffset:"))
}

// Updates an acceleration structure with new geometry or instance data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecommandencoder/3553898-refitaccelerationstructure?language=objc
func (a_ AccelerationStructureCommandEncoderWrapper) RefitAccelerationStructureDescriptorDestinationScratchBufferScratchBufferOffset(sourceAccelerationStructure PAccelerationStructure, descriptor IAccelerationStructureDescriptor, destinationAccelerationStructure PAccelerationStructure, scratchBuffer PBuffer, scratchBufferOffset uint) {
	po0 := objc.WrapAsProtocol("MTLAccelerationStructure", sourceAccelerationStructure)
	po2 := objc.WrapAsProtocol("MTLAccelerationStructure", destinationAccelerationStructure)
	po3 := objc.WrapAsProtocol("MTLBuffer", scratchBuffer)
	objc.Call[objc.Void](a_, objc.Sel("refitAccelerationStructure:descriptor:destination:scratchBuffer:scratchBufferOffset:"), po0, objc.Ptr(descriptor), po2, po3, scratchBufferOffset)
}

func (a_ AccelerationStructureCommandEncoderWrapper) HasSampleCountersInBufferAtSampleIndexWithBarrier() bool {
	return a_.RespondsToSelector(objc.Sel("sampleCountersInBuffer:atSampleIndex:withBarrier:"))
}

// Encodes a command to sample hardware counters at this point in the acceleration structure pass and store the samples into a counter sample buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecommandencoder/3553899-samplecountersinbuffer?language=objc
func (a_ AccelerationStructureCommandEncoderWrapper) SampleCountersInBufferAtSampleIndexWithBarrier(sampleBuffer PCounterSampleBuffer, sampleIndex uint, barrier bool) {
	po0 := objc.WrapAsProtocol("MTLCounterSampleBuffer", sampleBuffer)
	objc.Call[objc.Void](a_, objc.Sel("sampleCountersInBuffer:atSampleIndex:withBarrier:"), po0, sampleIndex, barrier)
}

func (a_ AccelerationStructureCommandEncoderWrapper) HasCopyAccelerationStructureToAccelerationStructure() bool {
	return a_.RespondsToSelector(objc.Sel("copyAccelerationStructure:toAccelerationStructure:"))
}

// Encodes a command to copy the data from one acceleration structure to another. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecommandencoder/3553895-copyaccelerationstructure?language=objc
func (a_ AccelerationStructureCommandEncoderWrapper) CopyAccelerationStructureToAccelerationStructure(sourceAccelerationStructure PAccelerationStructure, destinationAccelerationStructure PAccelerationStructure) {
	po0 := objc.WrapAsProtocol("MTLAccelerationStructure", sourceAccelerationStructure)
	po1 := objc.WrapAsProtocol("MTLAccelerationStructure", destinationAccelerationStructure)
	objc.Call[objc.Void](a_, objc.Sel("copyAccelerationStructure:toAccelerationStructure:"), po0, po1)
}

func (a_ AccelerationStructureCommandEncoderWrapper) HasWriteCompactedAccelerationStructureSizeToBufferOffset() bool {
	return a_.RespondsToSelector(objc.Sel("writeCompactedAccelerationStructureSize:toBuffer:offset:"))
}

// Encodes a command to calculate the compacted size of an acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecommandencoder/3553907-writecompactedaccelerationstruct?language=objc
func (a_ AccelerationStructureCommandEncoderWrapper) WriteCompactedAccelerationStructureSizeToBufferOffset(accelerationStructure PAccelerationStructure, buffer PBuffer, offset uint) {
	po0 := objc.WrapAsProtocol("MTLAccelerationStructure", accelerationStructure)
	po1 := objc.WrapAsProtocol("MTLBuffer", buffer)
	objc.Call[objc.Void](a_, objc.Sel("writeCompactedAccelerationStructureSize:toBuffer:offset:"), po0, po1, offset)
}

func (a_ AccelerationStructureCommandEncoderWrapper) HasUseHeap() bool {
	return a_.RespondsToSelector(objc.Sel("useHeap:"))
}

// Makes the resources contained in the specified heap available to the acceleration structure pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecommandencoder/3553902-useheap?language=objc
func (a_ AccelerationStructureCommandEncoderWrapper) UseHeap(heap PHeap) {
	po0 := objc.WrapAsProtocol("MTLHeap", heap)
	objc.Call[objc.Void](a_, objc.Sel("useHeap:"), po0)
}

func (a_ AccelerationStructureCommandEncoderWrapper) HasUseResourcesCountUsage() bool {
	return a_.RespondsToSelector(objc.Sel("useResources:count:usage:"))
}

// Specifies that an array of resources in an argument buffer can be safely used by the acceleration structure pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecommandencoder/3553905-useresources?language=objc
func (a_ AccelerationStructureCommandEncoderWrapper) UseResourcesCountUsage(resources PResource, count uint, usage ResourceUsage) {
	po0 := objc.WrapAsProtocol("MTLResource", resources)
	objc.Call[objc.Void](a_, objc.Sel("useResources:count:usage:"), po0, count, usage)
}

func (a_ AccelerationStructureCommandEncoderWrapper) HasUseHeapsCount() bool {
	return a_.RespondsToSelector(objc.Sel("useHeaps:count:"))
}

// Specifies that an array of heaps containing resources in an argument buffer can be safely used by the acceleration structure pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecommandencoder/3553903-useheaps?language=objc
func (a_ AccelerationStructureCommandEncoderWrapper) UseHeapsCount(heaps PHeap, count uint) {
	po0 := objc.WrapAsProtocol("MTLHeap", heaps)
	objc.Call[objc.Void](a_, objc.Sel("useHeaps:count:"), po0, count)
}

func (a_ AccelerationStructureCommandEncoderWrapper) HasUpdateFence() bool {
	return a_.RespondsToSelector(objc.Sel("updateFence:"))
}

// Tells the GPU to update the fence after all commands encoded by the encoder have finished executing. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecommandencoder/3553901-updatefence?language=objc
func (a_ AccelerationStructureCommandEncoderWrapper) UpdateFence(fence PFence) {
	po0 := objc.WrapAsProtocol("MTLFence", fence)
	objc.Call[objc.Void](a_, objc.Sel("updateFence:"), po0)
}

func (a_ AccelerationStructureCommandEncoderWrapper) HasBuildAccelerationStructureDescriptorScratchBufferScratchBufferOffset() bool {
	return a_.RespondsToSelector(objc.Sel("buildAccelerationStructure:descriptor:scratchBuffer:scratchBufferOffset:"))
}

// Encodes a command to build a new acceleration structure. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecommandencoder/3553894-buildaccelerationstructure?language=objc
func (a_ AccelerationStructureCommandEncoderWrapper) BuildAccelerationStructureDescriptorScratchBufferScratchBufferOffset(accelerationStructure PAccelerationStructure, descriptor IAccelerationStructureDescriptor, scratchBuffer PBuffer, scratchBufferOffset uint) {
	po0 := objc.WrapAsProtocol("MTLAccelerationStructure", accelerationStructure)
	po2 := objc.WrapAsProtocol("MTLBuffer", scratchBuffer)
	objc.Call[objc.Void](a_, objc.Sel("buildAccelerationStructure:descriptor:scratchBuffer:scratchBufferOffset:"), po0, objc.Ptr(descriptor), po2, scratchBufferOffset)
}

func (a_ AccelerationStructureCommandEncoderWrapper) HasWaitForFence() bool {
	return a_.RespondsToSelector(objc.Sel("waitForFence:"))
}

// Tells the GPU to wait until the fence is updated before executing any commands encoded by the command encoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlaccelerationstructurecommandencoder/3553906-waitforfence?language=objc
func (a_ AccelerationStructureCommandEncoderWrapper) WaitForFence(fence PFence) {
	po0 := objc.WrapAsProtocol("MTLFence", fence)
	objc.Call[objc.Void](a_, objc.Sel("waitForFence:"), po0)
}
