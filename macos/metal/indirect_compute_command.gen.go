// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// A compute command in an indirect command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcomputecommand?language=objc
type PIndirectComputeCommand interface {
	// optional
	SetThreadgroupMemoryLengthAtIndex(length uint, index uint)
	HasSetThreadgroupMemoryLengthAtIndex() bool

	// optional
	SetComputePipelineState(pipelineState ComputePipelineStateWrapper)
	HasSetComputePipelineState() bool

	// optional
	ConcurrentDispatchThreadgroupsThreadsPerThreadgroup(threadgroupsPerGrid Size, threadsPerThreadgroup Size)
	HasConcurrentDispatchThreadgroupsThreadsPerThreadgroup() bool

	// optional
	SetKernelBufferOffsetAtIndex(buffer BufferWrapper, offset uint, index uint)
	HasSetKernelBufferOffsetAtIndex() bool

	// optional
	SetStageInRegion(region Region)
	HasSetStageInRegion() bool

	// optional
	ConcurrentDispatchThreadsThreadsPerThreadgroup(threadsPerGrid Size, threadsPerThreadgroup Size)
	HasConcurrentDispatchThreadsThreadsPerThreadgroup() bool

	// optional
	ClearBarrier()
	HasClearBarrier() bool

	// optional
	SetBarrier()
	HasSetBarrier() bool

	// optional
	SetImageblockWidthHeight(width uint, height uint)
	HasSetImageblockWidthHeight() bool

	// optional
	Reset()
	HasReset() bool
}

// A concrete type wrapper for the [PIndirectComputeCommand] protocol.
type IndirectComputeCommandWrapper struct {
	objc.Object
}

func (i_ IndirectComputeCommandWrapper) HasSetThreadgroupMemoryLengthAtIndex() bool {
	return i_.RespondsToSelector(objc.Sel("setThreadgroupMemoryLength:atIndex:"))
}

// Sets the size of a block of threadgroup memory. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcomputecommand/2966612-setthreadgroupmemorylength?language=objc
func (i_ IndirectComputeCommandWrapper) SetThreadgroupMemoryLengthAtIndex(length uint, index uint) {
	objc.Call[objc.Void](i_, objc.Sel("setThreadgroupMemoryLength:atIndex:"), length, index)
}

func (i_ IndirectComputeCommandWrapper) HasSetComputePipelineState() bool {
	return i_.RespondsToSelector(objc.Sel("setComputePipelineState:"))
}

// Sets the command’s compute pipeline state object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcomputecommand/2966609-setcomputepipelinestate?language=objc
func (i_ IndirectComputeCommandWrapper) SetComputePipelineState(pipelineState PComputePipelineState) {
	po0 := objc.WrapAsProtocol("MTLComputePipelineState", pipelineState)
	objc.Call[objc.Void](i_, objc.Sel("setComputePipelineState:"), po0)
}

func (i_ IndirectComputeCommandWrapper) HasConcurrentDispatchThreadgroupsThreadsPerThreadgroup() bool {
	return i_.RespondsToSelector(objc.Sel("concurrentDispatchThreadgroups:threadsPerThreadgroup:"))
}

// Encodes a compute command using a grid aligned to threadgroup boundaries. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcomputecommand/2966604-concurrentdispatchthreadgroups?language=objc
func (i_ IndirectComputeCommandWrapper) ConcurrentDispatchThreadgroupsThreadsPerThreadgroup(threadgroupsPerGrid Size, threadsPerThreadgroup Size) {
	objc.Call[objc.Void](i_, objc.Sel("concurrentDispatchThreadgroups:threadsPerThreadgroup:"), threadgroupsPerGrid, threadsPerThreadgroup)
}

func (i_ IndirectComputeCommandWrapper) HasSetKernelBufferOffsetAtIndex() bool {
	return i_.RespondsToSelector(objc.Sel("setKernelBuffer:offset:atIndex:"))
}

// Sets a buffer for the compute function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcomputecommand/2976452-setkernelbuffer?language=objc
func (i_ IndirectComputeCommandWrapper) SetKernelBufferOffsetAtIndex(buffer PBuffer, offset uint, index uint) {
	po0 := objc.WrapAsProtocol("MTLBuffer", buffer)
	objc.Call[objc.Void](i_, objc.Sel("setKernelBuffer:offset:atIndex:"), po0, offset, index)
}

func (i_ IndirectComputeCommandWrapper) HasSetStageInRegion() bool {
	return i_.RespondsToSelector(objc.Sel("setStageInRegion:"))
}

// Sets the region of the stage-in attributes to apply to the compute kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcomputecommand/2966611-setstageinregion?language=objc
func (i_ IndirectComputeCommandWrapper) SetStageInRegion(region Region) {
	objc.Call[objc.Void](i_, objc.Sel("setStageInRegion:"), region)
}

func (i_ IndirectComputeCommandWrapper) HasConcurrentDispatchThreadsThreadsPerThreadgroup() bool {
	return i_.RespondsToSelector(objc.Sel("concurrentDispatchThreads:threadsPerThreadgroup:"))
}

// Encodes a compute command using an arbitrarily sized grid. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcomputecommand/3043398-concurrentdispatchthreads?language=objc
func (i_ IndirectComputeCommandWrapper) ConcurrentDispatchThreadsThreadsPerThreadgroup(threadsPerGrid Size, threadsPerThreadgroup Size) {
	objc.Call[objc.Void](i_, objc.Sel("concurrentDispatchThreads:threadsPerThreadgroup:"), threadsPerGrid, threadsPerThreadgroup)
}

func (i_ IndirectComputeCommandWrapper) HasClearBarrier() bool {
	return i_.RespondsToSelector(objc.Sel("clearBarrier"))
}

// Removes any barrier set on the command. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcomputecommand/3043397-clearbarrier?language=objc
func (i_ IndirectComputeCommandWrapper) ClearBarrier() {
	objc.Call[objc.Void](i_, objc.Sel("clearBarrier"))
}

func (i_ IndirectComputeCommandWrapper) HasSetBarrier() bool {
	return i_.RespondsToSelector(objc.Sel("setBarrier"))
}

// Adds a barrier to ensure that commands executed prior to this command are complete before this command executes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcomputecommand/3043399-setbarrier?language=objc
func (i_ IndirectComputeCommandWrapper) SetBarrier() {
	objc.Call[objc.Void](i_, objc.Sel("setBarrier"))
}

func (i_ IndirectComputeCommandWrapper) HasSetImageblockWidthHeight() bool {
	return i_.RespondsToSelector(objc.Sel("setImageblockWidth:height:"))
}

// Sets the size, in pixels, of the imageblock. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcomputecommand/3172768-setimageblockwidth?language=objc
func (i_ IndirectComputeCommandWrapper) SetImageblockWidthHeight(width uint, height uint) {
	objc.Call[objc.Void](i_, objc.Sel("setImageblockWidth:height:"), width, height)
}

func (i_ IndirectComputeCommandWrapper) HasReset() bool {
	return i_.RespondsToSelector(objc.Sel("reset"))
}

// Resets the command to its default state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlindirectcomputecommand/2981029-reset?language=objc
func (i_ IndirectComputeCommandWrapper) Reset() {
	objc.Call[objc.Void](i_, objc.Sel("reset"))
}
