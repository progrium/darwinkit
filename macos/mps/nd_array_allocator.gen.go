// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayallocator?language=objc
type PNDArrayAllocator interface {
	// optional
	ArrayForCommandBufferArrayDescriptorKernel(cmdBuf metal.CommandBufferWrapper, descriptor NDArrayDescriptor, kernel Kernel) INDArray
	HasArrayForCommandBufferArrayDescriptorKernel() bool
}

// A concrete type wrapper for the [PNDArrayAllocator] protocol.
type NDArrayAllocatorWrapper struct {
	objc.Object
}

func (n_ NDArrayAllocatorWrapper) HasArrayForCommandBufferArrayDescriptorKernel() bool {
	return n_.RespondsToSelector(objc.Sel("arrayForCommandBuffer:arrayDescriptor:kernel:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayallocator/3143490-arrayforcommandbuffer?language=objc
func (n_ NDArrayAllocatorWrapper) ArrayForCommandBufferArrayDescriptorKernel(cmdBuf metal.PCommandBuffer, descriptor INDArrayDescriptor, kernel IKernel) NDArray {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[NDArray](n_, objc.Sel("arrayForCommandBuffer:arrayDescriptor:kernel:"), po0, objc.Ptr(descriptor), objc.Ptr(kernel))
	return rv
}
