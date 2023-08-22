// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNBatchNormalization] class.
var CNNBatchNormalizationClass = _CNNBatchNormalizationClass{objc.GetClass("MPSCNNBatchNormalization")}

type _CNNBatchNormalizationClass struct {
	objc.Class
}

// An interface definition for the [CNNBatchNormalization] class.
type ICNNBatchNormalization interface {
	ICNNKernel
	ReloadMeanAndVarianceFromDataSource()
	ReloadGammaAndBetaFromDataSource()
	EncodeBatchToCommandBufferSourceImagesBatchNormalizationStateDestinationImages(commandBuffer metal.PCommandBuffer, sourceImages *foundation.Array, batchNormalizationState ICNNBatchNormalizationState, destinationImages *foundation.Array)
	EncodeBatchToCommandBufferObjectSourceImagesBatchNormalizationStateDestinationImages(commandBufferObject objc.IObject, sourceImages *foundation.Array, batchNormalizationState ICNNBatchNormalizationState, destinationImages *foundation.Array)
	ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer metal.PCommandBuffer, gammaAndBetaState ICNNNormalizationGammaAndBetaState)
	ReloadGammaAndBetaWithCommandBufferObjectGammaAndBetaState(commandBufferObject objc.IObject, gammaAndBetaState ICNNNormalizationGammaAndBetaState)
	ReloadMeanAndVarianceWithCommandBufferMeanAndVarianceState(commandBuffer metal.PCommandBuffer, meanAndVarianceState ICNNNormalizationMeanAndVarianceState)
	ReloadMeanAndVarianceWithCommandBufferObjectMeanAndVarianceState(commandBufferObject objc.IObject, meanAndVarianceState ICNNNormalizationMeanAndVarianceState)
	EncodeToCommandBufferSourceImageBatchNormalizationStateDestinationImage(commandBuffer metal.PCommandBuffer, sourceImage IImage, batchNormalizationState ICNNBatchNormalizationState, destinationImage IImage)
	EncodeToCommandBufferObjectSourceImageBatchNormalizationStateDestinationImage(commandBufferObject objc.IObject, sourceImage IImage, batchNormalizationState ICNNBatchNormalizationState, destinationImage IImage)
	DataSource() CNNBatchNormalizationDataSourceWrapper
	Epsilon() float64
	SetEpsilon(value float64)
	NumberOfFeatureChannels() uint
}

// A batch normalization kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization?language=objc
type CNNBatchNormalization struct {
	CNNKernel
}

func CNNBatchNormalizationFrom(ptr unsafe.Pointer) CNNBatchNormalization {
	return CNNBatchNormalization{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (c_ CNNBatchNormalization) InitWithDeviceDataSource(device metal.PDevice, dataSource PCNNBatchNormalizationDataSource) CNNBatchNormalization {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	po1 := objc.WrapAsProtocol("MPSCNNBatchNormalizationDataSource", dataSource)
	rv := objc.Call[CNNBatchNormalization](c_, objc.Sel("initWithDevice:dataSource:"), po0, po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2942600-initwithdevice?language=objc
func NewCNNBatchNormalizationWithDeviceDataSource(device metal.PDevice, dataSource PCNNBatchNormalizationDataSource) CNNBatchNormalization {
	instance := CNNBatchNormalizationClass.Alloc().InitWithDeviceDataSource(device, dataSource)
	instance.Autorelease()
	return instance
}

func (cc _CNNBatchNormalizationClass) Alloc() CNNBatchNormalization {
	rv := objc.Call[CNNBatchNormalization](cc, objc.Sel("alloc"))
	return rv
}

func CNNBatchNormalization_Alloc() CNNBatchNormalization {
	return CNNBatchNormalizationClass.Alloc()
}

func (cc _CNNBatchNormalizationClass) New() CNNBatchNormalization {
	rv := objc.Call[CNNBatchNormalization](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNBatchNormalization() CNNBatchNormalization {
	return CNNBatchNormalizationClass.New()
}

func (c_ CNNBatchNormalization) Init() CNNBatchNormalization {
	rv := objc.Call[CNNBatchNormalization](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNBatchNormalization) InitWithDevice(device metal.PDevice) CNNBatchNormalization {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBatchNormalization](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNBatchNormalizationWithDevice(device metal.PDevice) CNNBatchNormalization {
	instance := CNNBatchNormalizationClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNBatchNormalization) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNBatchNormalization {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBatchNormalization](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNBatchNormalization_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNBatchNormalization {
	instance := CNNBatchNormalizationClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/3002358-reloadmeanandvariancefromdatasou?language=objc
func (c_ CNNBatchNormalization) ReloadMeanAndVarianceFromDataSource() {
	objc.Call[objc.Void](c_, objc.Sel("reloadMeanAndVarianceFromDataSource"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2976464-reloadgammaandbetafromdatasource?language=objc
func (c_ CNNBatchNormalization) ReloadGammaAndBetaFromDataSource() {
	objc.Call[objc.Void](c_, objc.Sel("reloadGammaAndBetaFromDataSource"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2942610-encodebatchtocommandbuffer?language=objc
func (c_ CNNBatchNormalization) EncodeBatchToCommandBufferSourceImagesBatchNormalizationStateDestinationImages(commandBuffer metal.PCommandBuffer, sourceImages *foundation.Array, batchNormalizationState ICNNBatchNormalizationState, destinationImages *foundation.Array) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](c_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:batchNormalizationState:destinationImages:"), po0, sourceImages, objc.Ptr(batchNormalizationState), destinationImages)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2942610-encodebatchtocommandbuffer?language=objc
func (c_ CNNBatchNormalization) EncodeBatchToCommandBufferObjectSourceImagesBatchNormalizationStateDestinationImages(commandBufferObject objc.IObject, sourceImages *foundation.Array, batchNormalizationState ICNNBatchNormalizationState, destinationImages *foundation.Array) {
	objc.Call[objc.Void](c_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:batchNormalizationState:destinationImages:"), objc.Ptr(commandBufferObject), sourceImages, objc.Ptr(batchNormalizationState), destinationImages)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2953965-reloadgammaandbetawithcommandbuf?language=objc
func (c_ CNNBatchNormalization) ReloadGammaAndBetaWithCommandBufferGammaAndBetaState(commandBuffer metal.PCommandBuffer, gammaAndBetaState ICNNNormalizationGammaAndBetaState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](c_, objc.Sel("reloadGammaAndBetaWithCommandBuffer:gammaAndBetaState:"), po0, objc.Ptr(gammaAndBetaState))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2953965-reloadgammaandbetawithcommandbuf?language=objc
func (c_ CNNBatchNormalization) ReloadGammaAndBetaWithCommandBufferObjectGammaAndBetaState(commandBufferObject objc.IObject, gammaAndBetaState ICNNNormalizationGammaAndBetaState) {
	objc.Call[objc.Void](c_, objc.Sel("reloadGammaAndBetaWithCommandBuffer:gammaAndBetaState:"), objc.Ptr(commandBufferObject), objc.Ptr(gammaAndBetaState))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/3002359-reloadmeanandvariancewithcommand?language=objc
func (c_ CNNBatchNormalization) ReloadMeanAndVarianceWithCommandBufferMeanAndVarianceState(commandBuffer metal.PCommandBuffer, meanAndVarianceState ICNNNormalizationMeanAndVarianceState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](c_, objc.Sel("reloadMeanAndVarianceWithCommandBuffer:meanAndVarianceState:"), po0, objc.Ptr(meanAndVarianceState))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/3002359-reloadmeanandvariancewithcommand?language=objc
func (c_ CNNBatchNormalization) ReloadMeanAndVarianceWithCommandBufferObjectMeanAndVarianceState(commandBufferObject objc.IObject, meanAndVarianceState ICNNNormalizationMeanAndVarianceState) {
	objc.Call[objc.Void](c_, objc.Sel("reloadMeanAndVarianceWithCommandBuffer:meanAndVarianceState:"), objc.Ptr(commandBufferObject), objc.Ptr(meanAndVarianceState))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2942591-encodetocommandbuffer?language=objc
func (c_ CNNBatchNormalization) EncodeToCommandBufferSourceImageBatchNormalizationStateDestinationImage(commandBuffer metal.PCommandBuffer, sourceImage IImage, batchNormalizationState ICNNBatchNormalizationState, destinationImage IImage) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](c_, objc.Sel("encodeToCommandBuffer:sourceImage:batchNormalizationState:destinationImage:"), po0, objc.Ptr(sourceImage), objc.Ptr(batchNormalizationState), objc.Ptr(destinationImage))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2942591-encodetocommandbuffer?language=objc
func (c_ CNNBatchNormalization) EncodeToCommandBufferObjectSourceImageBatchNormalizationStateDestinationImage(commandBufferObject objc.IObject, sourceImage IImage, batchNormalizationState ICNNBatchNormalizationState, destinationImage IImage) {
	objc.Call[objc.Void](c_, objc.Sel("encodeToCommandBuffer:sourceImage:batchNormalizationState:destinationImage:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceImage), objc.Ptr(batchNormalizationState), objc.Ptr(destinationImage))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2953967-datasource?language=objc
func (c_ CNNBatchNormalization) DataSource() CNNBatchNormalizationDataSourceWrapper {
	rv := objc.Call[CNNBatchNormalizationDataSourceWrapper](c_, objc.Sel("dataSource"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2942599-epsilon?language=objc
func (c_ CNNBatchNormalization) Epsilon() float64 {
	rv := objc.Call[float64](c_, objc.Sel("epsilon"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2942599-epsilon?language=objc
func (c_ CNNBatchNormalization) SetEpsilon(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setEpsilon:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalization/2942604-numberoffeaturechannels?language=objc
func (c_ CNNBatchNormalization) NumberOfFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("numberOfFeatureChannels"))
	return rv
}
