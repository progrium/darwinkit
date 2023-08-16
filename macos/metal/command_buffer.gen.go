// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/macos/corefoundation"
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/objc"
)

// A container that stores a sequence of GPU commands that you encode into it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer?language=objc
type PCommandBuffer interface {
	// optional
	Enqueue()
	HasEnqueue() bool

	// optional
	EncodeSignalEventValue(event EventWrapper, value uint64)
	HasEncodeSignalEventValue() bool

	// optional
	WaitUntilCompleted()
	HasWaitUntilCompleted() bool

	// optional
	PresentDrawable(drawable DrawableWrapper)
	HasPresentDrawable() bool

	// optional
	PushDebugGroup(string_ string)
	HasPushDebugGroup() bool

	// optional
	BlitCommandEncoderWithDescriptor(blitPassDescriptor BlitPassDescriptor) PBlitCommandEncoder
	HasBlitCommandEncoderWithDescriptor() bool

	// optional
	BlitCommandEncoder() PBlitCommandEncoder
	HasBlitCommandEncoder() bool

	// optional
	ComputeCommandEncoder() PComputeCommandEncoder
	HasComputeCommandEncoder() bool

	// optional
	ParallelRenderCommandEncoderWithDescriptor(renderPassDescriptor RenderPassDescriptor) PParallelRenderCommandEncoder
	HasParallelRenderCommandEncoderWithDescriptor() bool

	// optional
	EncodeWaitForEventValue(event EventWrapper, value uint64)
	HasEncodeWaitForEventValue() bool

	// optional
	Commit()
	HasCommit() bool

	// optional
	ComputeCommandEncoderWithDescriptor(computePassDescriptor ComputePassDescriptor) PComputeCommandEncoder
	HasComputeCommandEncoderWithDescriptor() bool

	// optional
	ResourceStateCommandEncoder() PResourceStateCommandEncoder
	HasResourceStateCommandEncoder() bool

	// optional
	AddCompletedHandler(block CommandBufferHandler)
	HasAddCompletedHandler() bool

	// optional
	AccelerationStructureCommandEncoder() PAccelerationStructureCommandEncoder
	HasAccelerationStructureCommandEncoder() bool

	// optional
	PopDebugGroup()
	HasPopDebugGroup() bool

	// optional
	ResourceStateCommandEncoderWithDescriptor(resourceStatePassDescriptor ResourceStatePassDescriptor) PResourceStateCommandEncoder
	HasResourceStateCommandEncoderWithDescriptor() bool

	// optional
	RenderCommandEncoderWithDescriptor(renderPassDescriptor RenderPassDescriptor) PRenderCommandEncoder
	HasRenderCommandEncoderWithDescriptor() bool

	// optional
	AddScheduledHandler(block CommandBufferHandler)
	HasAddScheduledHandler() bool

	// optional
	ComputeCommandEncoderWithDispatchType(dispatchType DispatchType) PComputeCommandEncoder
	HasComputeCommandEncoderWithDispatchType() bool

	// optional
	WaitUntilScheduled()
	HasWaitUntilScheduled() bool

	// optional
	RetainedReferences() bool
	HasRetainedReferences() bool

	// optional
	GPUStartTime() corefoundation.TimeInterval
	HasGPUStartTime() bool

	// optional
	CommandQueue() PCommandQueue
	HasCommandQueue() bool

	// optional
	Error() foundation.IError
	HasError() bool

	// optional
	GPUEndTime() corefoundation.TimeInterval
	HasGPUEndTime() bool

	// optional
	Logs() PLogContainer
	HasLogs() bool

	// optional
	KernelStartTime() corefoundation.TimeInterval
	HasKernelStartTime() bool

	// optional
	Device() PDevice
	HasDevice() bool

	// optional
	ErrorOptions() CommandBufferErrorOption
	HasErrorOptions() bool

	// optional
	KernelEndTime() corefoundation.TimeInterval
	HasKernelEndTime() bool

	// optional
	SetLabel(value string)
	HasSetLabel() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	Status() CommandBufferStatus
	HasStatus() bool
}

// A concrete type wrapper for the [PCommandBuffer] protocol.
type CommandBufferWrapper struct {
	objc.Object
}

func (c_ CommandBufferWrapper) HasEnqueue() bool {
	return c_.RespondsToSelector(objc.Sel("enqueue"))
}

// Reserves the next available place for the command buffer in its command queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443019-enqueue?language=objc
func (c_ CommandBufferWrapper) Enqueue() {
	objc.Call[objc.Void](c_, objc.Sel("enqueue"))
}

func (c_ CommandBufferWrapper) HasEncodeSignalEventValue() bool {
	return c_.RespondsToSelector(objc.Sel("encodeSignalEvent:value:"))
}

// Encodes a command into the command buffer that pauses the GPU from running subsequent passes until the event equals or exceeds a value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/2966542-encodesignalevent?language=objc
func (c_ CommandBufferWrapper) EncodeSignalEventValue(event PEvent, value uint64) {
	po0 := objc.WrapAsProtocol("MTLEvent", event)
	objc.Call[objc.Void](c_, objc.Sel("encodeSignalEvent:value:"), po0, value)
}

func (c_ CommandBufferWrapper) HasWaitUntilCompleted() bool {
	return c_.RespondsToSelector(objc.Sel("waitUntilCompleted"))
}

// Blocks the current thread until the GPU finishes executing the command buffer and all of its completion handlers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443039-waituntilcompleted?language=objc
func (c_ CommandBufferWrapper) WaitUntilCompleted() {
	objc.Call[objc.Void](c_, objc.Sel("waitUntilCompleted"))
}

func (c_ CommandBufferWrapper) HasPresentDrawable() bool {
	return c_.RespondsToSelector(objc.Sel("presentDrawable:"))
}

// Presents a drawable as early as possible. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443029-presentdrawable?language=objc
func (c_ CommandBufferWrapper) PresentDrawable(drawable PDrawable) {
	po0 := objc.WrapAsProtocol("MTLDrawable", drawable)
	objc.Call[objc.Void](c_, objc.Sel("presentDrawable:"), po0)
}

func (c_ CommandBufferWrapper) HasPushDebugGroup() bool {
	return c_.RespondsToSelector(objc.Sel("pushDebugGroup:"))
}

// Marks the beginning of a debug group and gives it an identifying label, which temporarily replaces the previous group, if applicable. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/2869550-pushdebuggroup?language=objc
func (c_ CommandBufferWrapper) PushDebugGroup(string_ string) {
	objc.Call[objc.Void](c_, objc.Sel("pushDebugGroup:"), string_)
}

func (c_ CommandBufferWrapper) HasBlitCommandEncoderWithDescriptor() bool {
	return c_.RespondsToSelector(objc.Sel("blitCommandEncoderWithDescriptor:"))
}

// Creates a block information transfer (blit) encoder from a descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/3564431-blitcommandencoderwithdescriptor?language=objc
func (c_ CommandBufferWrapper) BlitCommandEncoderWithDescriptor(blitPassDescriptor IBlitPassDescriptor) BlitCommandEncoderWrapper {
	rv := objc.Call[BlitCommandEncoderWrapper](c_, objc.Sel("blitCommandEncoderWithDescriptor:"), objc.Ptr(blitPassDescriptor))
	return rv
}

func (c_ CommandBufferWrapper) HasBlitCommandEncoder() bool {
	return c_.RespondsToSelector(objc.Sel("blitCommandEncoder"))
}

// Creates a block information transfer (blit) encoder. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443001-blitcommandencoder?language=objc
func (c_ CommandBufferWrapper) BlitCommandEncoder() BlitCommandEncoderWrapper {
	rv := objc.Call[BlitCommandEncoderWrapper](c_, objc.Sel("blitCommandEncoder"))
	return rv
}

func (c_ CommandBufferWrapper) HasComputeCommandEncoder() bool {
	return c_.RespondsToSelector(objc.Sel("computeCommandEncoder"))
}

// Creates a compute command encoder that uses default settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443044-computecommandencoder?language=objc
func (c_ CommandBufferWrapper) ComputeCommandEncoder() ComputeCommandEncoderWrapper {
	rv := objc.Call[ComputeCommandEncoderWrapper](c_, objc.Sel("computeCommandEncoder"))
	return rv
}

func (c_ CommandBufferWrapper) HasParallelRenderCommandEncoderWithDescriptor() bool {
	return c_.RespondsToSelector(objc.Sel("parallelRenderCommandEncoderWithDescriptor:"))
}

// Creates a parallel render command encoder from a descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443009-parallelrendercommandencoderwith?language=objc
func (c_ CommandBufferWrapper) ParallelRenderCommandEncoderWithDescriptor(renderPassDescriptor IRenderPassDescriptor) ParallelRenderCommandEncoderWrapper {
	rv := objc.Call[ParallelRenderCommandEncoderWrapper](c_, objc.Sel("parallelRenderCommandEncoderWithDescriptor:"), objc.Ptr(renderPassDescriptor))
	return rv
}

func (c_ CommandBufferWrapper) HasEncodeWaitForEventValue() bool {
	return c_.RespondsToSelector(objc.Sel("encodeWaitForEvent:value:"))
}

// Encodes a command into the command buffer that pauses the GPU from running subsequent passes until the event equals or exceeds a value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/2966543-encodewaitforevent?language=objc
func (c_ CommandBufferWrapper) EncodeWaitForEventValue(event PEvent, value uint64) {
	po0 := objc.WrapAsProtocol("MTLEvent", event)
	objc.Call[objc.Void](c_, objc.Sel("encodeWaitForEvent:value:"), po0, value)
}

func (c_ CommandBufferWrapper) HasCommit() bool {
	return c_.RespondsToSelector(objc.Sel("commit"))
}

// Submits the command buffer to run on the GPU. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443003-commit?language=objc
func (c_ CommandBufferWrapper) Commit() {
	objc.Call[objc.Void](c_, objc.Sel("commit"))
}

func (c_ CommandBufferWrapper) HasComputeCommandEncoderWithDescriptor() bool {
	return c_.RespondsToSelector(objc.Sel("computeCommandEncoderWithDescriptor:"))
}

// Creates a compute command encoder from a descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/3564432-computecommandencoderwithdescrip?language=objc
func (c_ CommandBufferWrapper) ComputeCommandEncoderWithDescriptor(computePassDescriptor IComputePassDescriptor) ComputeCommandEncoderWrapper {
	rv := objc.Call[ComputeCommandEncoderWrapper](c_, objc.Sel("computeCommandEncoderWithDescriptor:"), objc.Ptr(computePassDescriptor))
	return rv
}

func (c_ CommandBufferWrapper) HasResourceStateCommandEncoder() bool {
	return c_.RespondsToSelector(objc.Sel("resourceStateCommandEncoder"))
}

// Creates a resource state command encoder that uses default settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/3043915-resourcestatecommandencoder?language=objc
func (c_ CommandBufferWrapper) ResourceStateCommandEncoder() ResourceStateCommandEncoderWrapper {
	rv := objc.Call[ResourceStateCommandEncoderWrapper](c_, objc.Sel("resourceStateCommandEncoder"))
	return rv
}

func (c_ CommandBufferWrapper) HasAddCompletedHandler() bool {
	return c_.RespondsToSelector(objc.Sel("addCompletedHandler:"))
}

// Registers a completion handler the GPU device calls immediately after the GPU finishes running the commands in the command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1442997-addcompletedhandler?language=objc
func (c_ CommandBufferWrapper) AddCompletedHandler(block CommandBufferHandler) {
	objc.Call[objc.Void](c_, objc.Sel("addCompletedHandler:"), block)
}

func (c_ CommandBufferWrapper) HasAccelerationStructureCommandEncoder() bool {
	return c_.RespondsToSelector(objc.Sel("accelerationStructureCommandEncoder"))
}

// Creates a ray-tracing acceleration structure command encoder that uses default settings. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/3553937-accelerationstructurecommandenco?language=objc
func (c_ CommandBufferWrapper) AccelerationStructureCommandEncoder() AccelerationStructureCommandEncoderWrapper {
	rv := objc.Call[AccelerationStructureCommandEncoderWrapper](c_, objc.Sel("accelerationStructureCommandEncoder"))
	return rv
}

func (c_ CommandBufferWrapper) HasPopDebugGroup() bool {
	return c_.RespondsToSelector(objc.Sel("popDebugGroup"))
}

// Marks the end of a debug group and, if applicable, restores the previous group from a stack. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/2869549-popdebuggroup?language=objc
func (c_ CommandBufferWrapper) PopDebugGroup() {
	objc.Call[objc.Void](c_, objc.Sel("popDebugGroup"))
}

func (c_ CommandBufferWrapper) HasResourceStateCommandEncoderWithDescriptor() bool {
	return c_.RespondsToSelector(objc.Sel("resourceStateCommandEncoderWithDescriptor:"))
}

// Creates a resource state command encoder from a descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/3566536-resourcestatecommandencoderwithd?language=objc
func (c_ CommandBufferWrapper) ResourceStateCommandEncoderWithDescriptor(resourceStatePassDescriptor IResourceStatePassDescriptor) ResourceStateCommandEncoderWrapper {
	rv := objc.Call[ResourceStateCommandEncoderWrapper](c_, objc.Sel("resourceStateCommandEncoderWithDescriptor:"), objc.Ptr(resourceStatePassDescriptor))
	return rv
}

func (c_ CommandBufferWrapper) HasRenderCommandEncoderWithDescriptor() bool {
	return c_.RespondsToSelector(objc.Sel("renderCommandEncoderWithDescriptor:"))
}

// Creates a render command encoder from a descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1442999-rendercommandencoderwithdescript?language=objc
func (c_ CommandBufferWrapper) RenderCommandEncoderWithDescriptor(renderPassDescriptor IRenderPassDescriptor) RenderCommandEncoderWrapper {
	rv := objc.Call[RenderCommandEncoderWrapper](c_, objc.Sel("renderCommandEncoderWithDescriptor:"), objc.Ptr(renderPassDescriptor))
	return rv
}

func (c_ CommandBufferWrapper) HasAddScheduledHandler() bool {
	return c_.RespondsToSelector(objc.Sel("addScheduledHandler:"))
}

// Registers a completion handler the GPU device calls immediately after it schedules the command buffer to run on the GPU. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1442991-addscheduledhandler?language=objc
func (c_ CommandBufferWrapper) AddScheduledHandler(block CommandBufferHandler) {
	objc.Call[objc.Void](c_, objc.Sel("addScheduledHandler:"), block)
}

func (c_ CommandBufferWrapper) HasComputeCommandEncoderWithDispatchType() bool {
	return c_.RespondsToSelector(objc.Sel("computeCommandEncoderWithDispatchType:"))
}

// Creates a compute command encoder with a dispatch type. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/2966541-computecommandencoderwithdispatc?language=objc
func (c_ CommandBufferWrapper) ComputeCommandEncoderWithDispatchType(dispatchType DispatchType) ComputeCommandEncoderWrapper {
	rv := objc.Call[ComputeCommandEncoderWrapper](c_, objc.Sel("computeCommandEncoderWithDispatchType:"), dispatchType)
	return rv
}

func (c_ CommandBufferWrapper) HasWaitUntilScheduled() bool {
	return c_.RespondsToSelector(objc.Sel("waitUntilScheduled"))
}

// Blocks the current thread until the command queue schedules the buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443036-waituntilscheduled?language=objc
func (c_ CommandBufferWrapper) WaitUntilScheduled() {
	objc.Call[objc.Void](c_, objc.Sel("waitUntilScheduled"))
}

func (c_ CommandBufferWrapper) HasRetainedReferences() bool {
	return c_.RespondsToSelector(objc.Sel("retainedReferences"))
}

// A Boolean value that indicates whether the command buffer maintains strong references to the resources it uses. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443011-retainedreferences?language=objc
func (c_ CommandBufferWrapper) RetainedReferences() bool {
	rv := objc.Call[bool](c_, objc.Sel("retainedReferences"))
	return rv
}

func (c_ CommandBufferWrapper) HasGPUStartTime() bool {
	return c_.RespondsToSelector(objc.Sel("GPUStartTime"))
}

// The host time, in seconds, when the GPU starts command buffer execution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1639924-gpustarttime?language=objc
func (c_ CommandBufferWrapper) GPUStartTime() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](c_, objc.Sel("GPUStartTime"))
	return rv
}

func (c_ CommandBufferWrapper) HasCommandQueue() bool {
	return c_.RespondsToSelector(objc.Sel("commandQueue"))
}

// The command queue that creates the command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443015-commandqueue?language=objc
func (c_ CommandBufferWrapper) CommandQueue() CommandQueueWrapper {
	rv := objc.Call[CommandQueueWrapper](c_, objc.Sel("commandQueue"))
	return rv
}

func (c_ CommandBufferWrapper) HasError() bool {
	return c_.RespondsToSelector(objc.Sel("error"))
}

// A description of an error when the GPU encounters an issue as it runs the command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443040-error?language=objc
func (c_ CommandBufferWrapper) Error() foundation.Error {
	rv := objc.Call[foundation.Error](c_, objc.Sel("error"))
	return rv
}

func (c_ CommandBufferWrapper) HasGPUEndTime() bool {
	return c_.RespondsToSelector(objc.Sel("GPUEndTime"))
}

// The host time, in seconds, when the GPU finishes execution of the command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1639926-gpuendtime?language=objc
func (c_ CommandBufferWrapper) GPUEndTime() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](c_, objc.Sel("GPUEndTime"))
	return rv
}

func (c_ CommandBufferWrapper) HasLogs() bool {
	return c_.RespondsToSelector(objc.Sel("logs"))
}

// The messages the command buffer records as the GPU runs its commands. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/3553939-logs?language=objc
func (c_ CommandBufferWrapper) Logs() LogContainerWrapper {
	rv := objc.Call[LogContainerWrapper](c_, objc.Sel("logs"))
	return rv
}

func (c_ CommandBufferWrapper) HasKernelStartTime() bool {
	return c_.RespondsToSelector(objc.Sel("kernelStartTime"))
}

// The host time, in seconds, when the CPU begins to schedule the command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1639925-kernelstarttime?language=objc
func (c_ CommandBufferWrapper) KernelStartTime() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](c_, objc.Sel("kernelStartTime"))
	return rv
}

func (c_ CommandBufferWrapper) HasDevice() bool {
	return c_.RespondsToSelector(objc.Sel("device"))
}

// The GPU device that indirectly owns the command buffer because you create it from a command queue the device also owns. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1442995-device?language=objc
func (c_ CommandBufferWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](c_, objc.Sel("device"))
	return rv
}

func (c_ CommandBufferWrapper) HasErrorOptions() bool {
	return c_.RespondsToSelector(objc.Sel("errorOptions"))
}

// Settings that determine which information the command buffer records about execution errors, and how it does it. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/3553938-erroroptions?language=objc
func (c_ CommandBufferWrapper) ErrorOptions() CommandBufferErrorOption {
	rv := objc.Call[CommandBufferErrorOption](c_, objc.Sel("errorOptions"))
	return rv
}

func (c_ CommandBufferWrapper) HasKernelEndTime() bool {
	return c_.RespondsToSelector(objc.Sel("kernelEndTime"))
}

// The host time, in seconds, when the CPU finishes scheduling the command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1640027-kernelendtime?language=objc
func (c_ CommandBufferWrapper) KernelEndTime() corefoundation.TimeInterval {
	rv := objc.Call[corefoundation.TimeInterval](c_, objc.Sel("kernelEndTime"))
	return rv
}

func (c_ CommandBufferWrapper) HasSetLabel() bool {
	return c_.RespondsToSelector(objc.Sel("setLabel:"))
}

// An optional name that can help you identify the command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443031-label?language=objc
func (c_ CommandBufferWrapper) SetLabel(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setLabel:"), value)
}

func (c_ CommandBufferWrapper) HasLabel() bool {
	return c_.RespondsToSelector(objc.Sel("label"))
}

// An optional name that can help you identify the command buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443031-label?language=objc
func (c_ CommandBufferWrapper) Label() string {
	rv := objc.Call[string](c_, objc.Sel("label"))
	return rv
}

func (c_ CommandBufferWrapper) HasStatus() bool {
	return c_.RespondsToSelector(objc.Sel("status"))
}

// The command buffer’s current state. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandbuffer/1443048-status?language=objc
func (c_ CommandBufferWrapper) Status() CommandBufferStatus {
	rv := objc.Call[CommandBufferStatus](c_, objc.Sel("status"))
	return rv
}
