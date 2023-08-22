// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNBatchNormalizationStatistics] class.
var CNNBatchNormalizationStatisticsClass = _CNNBatchNormalizationStatisticsClass{objc.GetClass("MPSCNNBatchNormalizationStatistics")}

type _CNNBatchNormalizationStatisticsClass struct {
	objc.Class
}

// An interface definition for the [CNNBatchNormalizationStatistics] class.
type ICNNBatchNormalizationStatistics interface {
	ICNNKernel
	EncodeBatchToCommandBufferSourceImagesBatchNormalizationState(commandBuffer metal.PCommandBuffer, sourceImages *foundation.Array, batchNormalizationState ICNNBatchNormalizationState)
	EncodeBatchToCommandBufferObjectSourceImagesBatchNormalizationState(commandBufferObject objc.IObject, sourceImages *foundation.Array, batchNormalizationState ICNNBatchNormalizationState)
}

// An object that stores statistics required to execute batch normalization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstatistics?language=objc
type CNNBatchNormalizationStatistics struct {
	CNNKernel
}

func CNNBatchNormalizationStatisticsFrom(ptr unsafe.Pointer) CNNBatchNormalizationStatistics {
	return CNNBatchNormalizationStatistics{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (c_ CNNBatchNormalizationStatistics) InitWithDevice(device metal.PDevice) CNNBatchNormalizationStatistics {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBatchNormalizationStatistics](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstatistics/2953968-initwithdevice?language=objc
func NewCNNBatchNormalizationStatisticsWithDevice(device metal.PDevice) CNNBatchNormalizationStatistics {
	instance := CNNBatchNormalizationStatisticsClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (cc _CNNBatchNormalizationStatisticsClass) Alloc() CNNBatchNormalizationStatistics {
	rv := objc.Call[CNNBatchNormalizationStatistics](cc, objc.Sel("alloc"))
	return rv
}

func CNNBatchNormalizationStatistics_Alloc() CNNBatchNormalizationStatistics {
	return CNNBatchNormalizationStatisticsClass.Alloc()
}

func (cc _CNNBatchNormalizationStatisticsClass) New() CNNBatchNormalizationStatistics {
	rv := objc.Call[CNNBatchNormalizationStatistics](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNBatchNormalizationStatistics() CNNBatchNormalizationStatistics {
	return CNNBatchNormalizationStatisticsClass.New()
}

func (c_ CNNBatchNormalizationStatistics) Init() CNNBatchNormalizationStatistics {
	rv := objc.Call[CNNBatchNormalizationStatistics](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNBatchNormalizationStatistics) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNBatchNormalizationStatistics {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBatchNormalizationStatistics](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNBatchNormalizationStatistics_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNBatchNormalizationStatistics {
	instance := CNNBatchNormalizationStatisticsClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstatistics/2942584-encodebatchtocommandbuffer?language=objc
func (c_ CNNBatchNormalizationStatistics) EncodeBatchToCommandBufferSourceImagesBatchNormalizationState(commandBuffer metal.PCommandBuffer, sourceImages *foundation.Array, batchNormalizationState ICNNBatchNormalizationState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](c_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:batchNormalizationState:"), po0, sourceImages, objc.Ptr(batchNormalizationState))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstatistics/2942584-encodebatchtocommandbuffer?language=objc
func (c_ CNNBatchNormalizationStatistics) EncodeBatchToCommandBufferObjectSourceImagesBatchNormalizationState(commandBufferObject objc.IObject, sourceImages *foundation.Array, batchNormalizationState ICNNBatchNormalizationState) {
	objc.Call[objc.Void](c_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:batchNormalizationState:"), objc.Ptr(commandBufferObject), sourceImages, objc.Ptr(batchNormalizationState))
}
