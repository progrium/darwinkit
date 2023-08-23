// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageallocator?language=objc
type PImageAllocator interface {
	// optional
	ImageForCommandBufferImageDescriptorKernel(cmdBuf metal.CommandBufferWrapper, descriptor ImageDescriptor, kernel Kernel) IImage
	HasImageForCommandBufferImageDescriptorKernel() bool

	// optional
	ImageBatchForCommandBufferImageDescriptorKernelCount(cmdBuf metal.CommandBufferWrapper, descriptor ImageDescriptor, kernel Kernel, count uint) *foundation.Array
	HasImageBatchForCommandBufferImageDescriptorKernelCount() bool
}

// A concrete type wrapper for the [PImageAllocator] protocol.
type ImageAllocatorWrapper struct {
	objc.Object
}

func (i_ ImageAllocatorWrapper) HasImageForCommandBufferImageDescriptorKernel() bool {
	return i_.RespondsToSelector(objc.Sel("imageForCommandBuffer:imageDescriptor:kernel:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageallocator/2866966-imageforcommandbuffer?language=objc
func (i_ ImageAllocatorWrapper) ImageForCommandBufferImageDescriptorKernel(cmdBuf metal.PCommandBuffer, descriptor IImageDescriptor, kernel IKernel) Image {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[Image](i_, objc.Sel("imageForCommandBuffer:imageDescriptor:kernel:"), po0, objc.Ptr(descriptor), objc.Ptr(kernel))
	return rv
}

func (i_ ImageAllocatorWrapper) HasImageBatchForCommandBufferImageDescriptorKernelCount() bool {
	return i_.RespondsToSelector(objc.Sel("imageBatchForCommandBuffer:imageDescriptor:kernel:count:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageallocator/3020685-imagebatchforcommandbuffer?language=objc
func (i_ ImageAllocatorWrapper) ImageBatchForCommandBufferImageDescriptorKernelCount(cmdBuf metal.PCommandBuffer, descriptor IImageDescriptor, kernel IKernel, count uint) *foundation.Array {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[*foundation.Array](i_, objc.Sel("imageBatchForCommandBuffer:imageDescriptor:kernel:count:"), po0, objc.Ptr(descriptor), objc.Ptr(kernel), count)
	return rv
}
