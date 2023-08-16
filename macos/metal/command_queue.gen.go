// AUTO-GENERATED CODE, DO NOT MODIFY

package metal

import (
	"github.com/progrium/macdriver/objc"
)

// An instance you use to create, submit, and schedule command buffers to a specific GPU device to run the commands within those buffers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandqueue?language=objc
type PCommandQueue interface {
	// optional
	CommandBuffer() PCommandBuffer
	HasCommandBuffer() bool

	// optional
	CommandBufferWithDescriptor(descriptor CommandBufferDescriptor) PCommandBuffer
	HasCommandBufferWithDescriptor() bool

	// optional
	CommandBufferWithUnretainedReferences() PCommandBuffer
	HasCommandBufferWithUnretainedReferences() bool

	// optional
	Device() PDevice
	HasDevice() bool

	// optional
	SetLabel(value string)
	HasSetLabel() bool

	// optional
	Label() string
	HasLabel() bool
}

// A concrete type wrapper for the [PCommandQueue] protocol.
type CommandQueueWrapper struct {
	objc.Object
}

func (c_ CommandQueueWrapper) HasCommandBuffer() bool {
	return c_.RespondsToSelector(objc.Sel("commandBuffer"))
}

// Returns a command buffer from the command queue that maintains strong references to resources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandqueue/1508686-commandbuffer?language=objc
func (c_ CommandQueueWrapper) CommandBuffer() CommandBufferWrapper {
	rv := objc.Call[CommandBufferWrapper](c_, objc.Sel("commandBuffer"))
	return rv
}

func (c_ CommandQueueWrapper) HasCommandBufferWithDescriptor() bool {
	return c_.RespondsToSelector(objc.Sel("commandBufferWithDescriptor:"))
}

// Returns a command buffer from the command queue that you configure with a descriptor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandqueue/3553957-commandbufferwithdescriptor?language=objc
func (c_ CommandQueueWrapper) CommandBufferWithDescriptor(descriptor ICommandBufferDescriptor) CommandBufferWrapper {
	rv := objc.Call[CommandBufferWrapper](c_, objc.Sel("commandBufferWithDescriptor:"), objc.Ptr(descriptor))
	return rv
}

func (c_ CommandQueueWrapper) HasCommandBufferWithUnretainedReferences() bool {
	return c_.RespondsToSelector(objc.Sel("commandBufferWithUnretainedReferences"))
}

// Returns a command buffer from the command queue that doesn’t maintain strong references to resources. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandqueue/1508684-commandbufferwithunretainedrefer?language=objc
func (c_ CommandQueueWrapper) CommandBufferWithUnretainedReferences() CommandBufferWrapper {
	rv := objc.Call[CommandBufferWrapper](c_, objc.Sel("commandBufferWithUnretainedReferences"))
	return rv
}

func (c_ CommandQueueWrapper) HasDevice() bool {
	return c_.RespondsToSelector(objc.Sel("device"))
}

// The GPU device that creates the command queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandqueue/1508687-device?language=objc
func (c_ CommandQueueWrapper) Device() DeviceWrapper {
	rv := objc.Call[DeviceWrapper](c_, objc.Sel("device"))
	return rv
}

func (c_ CommandQueueWrapper) HasSetLabel() bool {
	return c_.RespondsToSelector(objc.Sel("setLabel:"))
}

// An optional name that can help you identify the command queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandqueue/1508690-label?language=objc
func (c_ CommandQueueWrapper) SetLabel(value string) {
	objc.Call[objc.Void](c_, objc.Sel("setLabel:"), value)
}

func (c_ CommandQueueWrapper) HasLabel() bool {
	return c_.RespondsToSelector(objc.Sel("label"))
}

// An optional name that can help you identify the command queue. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metal/mtlcommandqueue/1508690-label?language=objc
func (c_ CommandQueueWrapper) Label() string {
	rv := objc.Call[string](c_, objc.Sel("label"))
	return rv
}
