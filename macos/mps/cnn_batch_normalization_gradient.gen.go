// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNBatchNormalizationGradient] class.
var CNNBatchNormalizationGradientClass = _CNNBatchNormalizationGradientClass{objc.GetClass("MPSCNNBatchNormalizationGradient")}

type _CNNBatchNormalizationGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNBatchNormalizationGradient] class.
type ICNNBatchNormalizationGradient interface {
	ICNNGradientKernel
	EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationState(commandBuffer metal.PCommandBuffer, sourceGradients *foundation.Array, sourceImages *foundation.Array, batchNormalizationState ICNNBatchNormalizationState) *foundation.Array
	EncodeBatchToCommandBufferObjectSourceGradientsSourceImagesBatchNormalizationState(commandBufferObject objc.IObject, sourceGradients *foundation.Array, sourceImages *foundation.Array, batchNormalizationState ICNNBatchNormalizationState) *foundation.Array
	EncodeToCommandBufferSourceGradientSourceImageBatchNormalizationState(commandBuffer metal.PCommandBuffer, sourceGradient IImage, sourceImage IImage, batchNormalizationState ICNNBatchNormalizationState) Image
	EncodeToCommandBufferObjectSourceGradientSourceImageBatchNormalizationState(commandBufferObject objc.IObject, sourceGradient IImage, sourceImage IImage, batchNormalizationState ICNNBatchNormalizationState) Image
}

// A gradient batch normalization kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradient?language=objc
type CNNBatchNormalizationGradient struct {
	CNNGradientKernel
}

func CNNBatchNormalizationGradientFrom(ptr unsafe.Pointer) CNNBatchNormalizationGradient {
	return CNNBatchNormalizationGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (c_ CNNBatchNormalizationGradient) InitWithDeviceFusedNeuronDescriptor(device metal.PDevice, fusedNeuronDescriptor INNNeuronDescriptor) CNNBatchNormalizationGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBatchNormalizationGradient](c_, objc.Sel("initWithDevice:fusedNeuronDescriptor:"), po0, objc.Ptr(fusedNeuronDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradient/3019332-initwithdevice?language=objc
func NewCNNBatchNormalizationGradientWithDeviceFusedNeuronDescriptor(device metal.PDevice, fusedNeuronDescriptor INNNeuronDescriptor) CNNBatchNormalizationGradient {
	instance := CNNBatchNormalizationGradientClass.Alloc().InitWithDeviceFusedNeuronDescriptor(device, fusedNeuronDescriptor)
	instance.Autorelease()
	return instance
}

func (cc _CNNBatchNormalizationGradientClass) Alloc() CNNBatchNormalizationGradient {
	rv := objc.Call[CNNBatchNormalizationGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNBatchNormalizationGradient_Alloc() CNNBatchNormalizationGradient {
	return CNNBatchNormalizationGradientClass.Alloc()
}

func (cc _CNNBatchNormalizationGradientClass) New() CNNBatchNormalizationGradient {
	rv := objc.Call[CNNBatchNormalizationGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNBatchNormalizationGradient() CNNBatchNormalizationGradient {
	return CNNBatchNormalizationGradientClass.New()
}

func (c_ CNNBatchNormalizationGradient) Init() CNNBatchNormalizationGradient {
	rv := objc.Call[CNNBatchNormalizationGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNBatchNormalizationGradient) InitWithDevice(device metal.PDevice) CNNBatchNormalizationGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBatchNormalizationGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNBatchNormalizationGradientWithDevice(device metal.PDevice) CNNBatchNormalizationGradient {
	instance := CNNBatchNormalizationGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNBatchNormalizationGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNBatchNormalizationGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBatchNormalizationGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNBatchNormalizationGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNBatchNormalizationGradient {
	instance := CNNBatchNormalizationGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradient/2942590-encodebatchtocommandbuffer?language=objc
func (c_ CNNBatchNormalizationGradient) EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationState(commandBuffer metal.PCommandBuffer, sourceGradients *foundation.Array, sourceImages *foundation.Array, batchNormalizationState ICNNBatchNormalizationState) *foundation.Array {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[*foundation.Array](c_, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:batchNormalizationState:"), po0, sourceGradients, sourceImages, objc.Ptr(batchNormalizationState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradient/2942590-encodebatchtocommandbuffer?language=objc
func (c_ CNNBatchNormalizationGradient) EncodeBatchToCommandBufferObjectSourceGradientsSourceImagesBatchNormalizationState(commandBufferObject objc.IObject, sourceGradients *foundation.Array, sourceImages *foundation.Array, batchNormalizationState ICNNBatchNormalizationState) *foundation.Array {
	rv := objc.Call[*foundation.Array](c_, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:batchNormalizationState:"), objc.Ptr(commandBufferObject), sourceGradients, sourceImages, objc.Ptr(batchNormalizationState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradient/2951885-encodetocommandbuffer?language=objc
func (c_ CNNBatchNormalizationGradient) EncodeToCommandBufferSourceGradientSourceImageBatchNormalizationState(commandBuffer metal.PCommandBuffer, sourceGradient IImage, sourceImage IImage, batchNormalizationState ICNNBatchNormalizationState) Image {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[Image](c_, objc.Sel("encodeToCommandBuffer:sourceGradient:sourceImage:batchNormalizationState:"), po0, objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(batchNormalizationState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationgradient/2951885-encodetocommandbuffer?language=objc
func (c_ CNNBatchNormalizationGradient) EncodeToCommandBufferObjectSourceGradientSourceImageBatchNormalizationState(commandBufferObject objc.IObject, sourceGradient IImage, sourceImage IImage, batchNormalizationState ICNNBatchNormalizationState) Image {
	rv := objc.Call[Image](c_, objc.Sel("encodeToCommandBuffer:sourceGradient:sourceImage:batchNormalizationState:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceGradient), objc.Ptr(sourceImage), objc.Ptr(batchNormalizationState))
	return rv
}
