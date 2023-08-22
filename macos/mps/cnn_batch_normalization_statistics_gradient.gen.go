// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNBatchNormalizationStatisticsGradient] class.
var CNNBatchNormalizationStatisticsGradientClass = _CNNBatchNormalizationStatisticsGradientClass{objc.GetClass("MPSCNNBatchNormalizationStatisticsGradient")}

type _CNNBatchNormalizationStatisticsGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNBatchNormalizationStatisticsGradient] class.
type ICNNBatchNormalizationStatisticsGradient interface {
	ICNNGradientKernel
	EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationState(commandBuffer metal.PCommandBuffer, sourceGradients *foundation.Array, sourceImages *foundation.Array, batchNormalizationState ICNNBatchNormalizationState)
	EncodeBatchToCommandBufferObjectSourceGradientsSourceImagesBatchNormalizationState(commandBufferObject objc.IObject, sourceGradients *foundation.Array, sourceImages *foundation.Array, batchNormalizationState ICNNBatchNormalizationState)
}

// An object that stores the gradient of the loss function with respect to the batch statistics and batch normalization weights. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstatisticsgradient?language=objc
type CNNBatchNormalizationStatisticsGradient struct {
	CNNGradientKernel
}

func CNNBatchNormalizationStatisticsGradientFrom(ptr unsafe.Pointer) CNNBatchNormalizationStatisticsGradient {
	return CNNBatchNormalizationStatisticsGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (c_ CNNBatchNormalizationStatisticsGradient) InitWithDeviceFusedNeuronDescriptor(device metal.PDevice, fusedNeuronDescriptor INNNeuronDescriptor) CNNBatchNormalizationStatisticsGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBatchNormalizationStatisticsGradient](c_, objc.Sel("initWithDevice:fusedNeuronDescriptor:"), po0, objc.Ptr(fusedNeuronDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstatisticsgradient/3013775-initwithdevice?language=objc
func NewCNNBatchNormalizationStatisticsGradientWithDeviceFusedNeuronDescriptor(device metal.PDevice, fusedNeuronDescriptor INNNeuronDescriptor) CNNBatchNormalizationStatisticsGradient {
	instance := CNNBatchNormalizationStatisticsGradientClass.Alloc().InitWithDeviceFusedNeuronDescriptor(device, fusedNeuronDescriptor)
	instance.Autorelease()
	return instance
}

func (cc _CNNBatchNormalizationStatisticsGradientClass) Alloc() CNNBatchNormalizationStatisticsGradient {
	rv := objc.Call[CNNBatchNormalizationStatisticsGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNBatchNormalizationStatisticsGradient_Alloc() CNNBatchNormalizationStatisticsGradient {
	return CNNBatchNormalizationStatisticsGradientClass.Alloc()
}

func (cc _CNNBatchNormalizationStatisticsGradientClass) New() CNNBatchNormalizationStatisticsGradient {
	rv := objc.Call[CNNBatchNormalizationStatisticsGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNBatchNormalizationStatisticsGradient() CNNBatchNormalizationStatisticsGradient {
	return CNNBatchNormalizationStatisticsGradientClass.New()
}

func (c_ CNNBatchNormalizationStatisticsGradient) Init() CNNBatchNormalizationStatisticsGradient {
	rv := objc.Call[CNNBatchNormalizationStatisticsGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNBatchNormalizationStatisticsGradient) InitWithDevice(device metal.PDevice) CNNBatchNormalizationStatisticsGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBatchNormalizationStatisticsGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNBatchNormalizationStatisticsGradientWithDevice(device metal.PDevice) CNNBatchNormalizationStatisticsGradient {
	instance := CNNBatchNormalizationStatisticsGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNBatchNormalizationStatisticsGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNBatchNormalizationStatisticsGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBatchNormalizationStatisticsGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNBatchNormalizationStatisticsGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNBatchNormalizationStatisticsGradient {
	instance := CNNBatchNormalizationStatisticsGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstatisticsgradient/2953964-encodebatchtocommandbuffer?language=objc
func (c_ CNNBatchNormalizationStatisticsGradient) EncodeBatchToCommandBufferSourceGradientsSourceImagesBatchNormalizationState(commandBuffer metal.PCommandBuffer, sourceGradients *foundation.Array, sourceImages *foundation.Array, batchNormalizationState ICNNBatchNormalizationState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](c_, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:batchNormalizationState:"), po0, sourceGradients, sourceImages, objc.Ptr(batchNormalizationState))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstatisticsgradient/2953964-encodebatchtocommandbuffer?language=objc
func (c_ CNNBatchNormalizationStatisticsGradient) EncodeBatchToCommandBufferObjectSourceGradientsSourceImagesBatchNormalizationState(commandBufferObject objc.IObject, sourceGradients *foundation.Array, sourceImages *foundation.Array, batchNormalizationState ICNNBatchNormalizationState) {
	objc.Call[objc.Void](c_, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:batchNormalizationState:"), objc.Ptr(commandBufferObject), sourceGradients, sourceImages, objc.Ptr(batchNormalizationState))
}
