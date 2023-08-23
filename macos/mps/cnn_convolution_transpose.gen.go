// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNConvolutionTranspose] class.
var CNNConvolutionTransposeClass = _CNNConvolutionTransposeClass{objc.GetClass("MPSCNNConvolutionTranspose")}

type _CNNConvolutionTransposeClass struct {
	objc.Class
}

// An interface definition for the [CNNConvolutionTranspose] class.
type ICNNConvolutionTranspose interface {
	ICNNKernel
	ReloadWeightsAndBiasesFromDataSource()
	EncodeBatchToCommandBufferSourceImagesConvolutionGradientStates(commandBuffer metal.PCommandBuffer, sourceImage *foundation.Array, convolutionGradientState *foundation.Array) *foundation.Array
	EncodeBatchToCommandBufferObjectSourceImagesConvolutionGradientStates(commandBufferObject objc.IObject, sourceImage *foundation.Array, convolutionGradientState *foundation.Array) *foundation.Array
	ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer metal.PCommandBuffer, state ICNNConvolutionWeightsAndBiasesState)
	ReloadWeightsAndBiasesWithCommandBufferObjectState(commandBufferObject objc.IObject, state ICNNConvolutionWeightsAndBiasesState)
	ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary(commandBuffer metal.PCommandBuffer, resultStateCanBeTemporary bool) CNNConvolutionWeightsAndBiasesState
	ExportWeightsAndBiasesWithCommandBufferObjectResultStateCanBeTemporary(commandBufferObject objc.IObject, resultStateCanBeTemporary bool) CNNConvolutionWeightsAndBiasesState
	EncodeToCommandBufferSourceImageConvolutionGradientState(commandBuffer metal.PCommandBuffer, sourceImage IImage, convolutionGradientState ICNNConvolutionGradientState) Image
	EncodeToCommandBufferObjectSourceImageConvolutionGradientState(commandBufferObject objc.IObject, sourceImage IImage, convolutionGradientState ICNNConvolutionGradientState) Image
	DataSource() CNNConvolutionDataSourceWrapper
	OutputFeatureChannels() uint
	InputFeatureChannels() uint
	KernelOffsetX() int
	SetKernelOffsetX(value int)
	Groups() uint
	KernelOffsetY() int
	SetKernelOffsetY(value int)
	AccumulatorPrecisionOption() NNConvolutionAccumulatorPrecisionOption
	SetAccumulatorPrecisionOption(value NNConvolutionAccumulatorPrecisionOption)
}

// A transposed convolution kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose?language=objc
type CNNConvolutionTranspose struct {
	CNNKernel
}

func CNNConvolutionTransposeFrom(ptr unsafe.Pointer) CNNConvolutionTranspose {
	return CNNConvolutionTranspose{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (c_ CNNConvolutionTranspose) InitWithDeviceWeights(device metal.PDevice, weights PCNNConvolutionDataSource) CNNConvolutionTranspose {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNConvolutionTranspose](c_, objc.Sel("initWithDevice:weights:"), po0, po1)
	return rv
}

// Initializes a transposed convolution kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867157-initwithdevice?language=objc
func NewCNNConvolutionTransposeWithDeviceWeights(device metal.PDevice, weights PCNNConvolutionDataSource) CNNConvolutionTranspose {
	instance := CNNConvolutionTransposeClass.Alloc().InitWithDeviceWeights(device, weights)
	instance.Autorelease()
	return instance
}

func (cc _CNNConvolutionTransposeClass) Alloc() CNNConvolutionTranspose {
	rv := objc.Call[CNNConvolutionTranspose](cc, objc.Sel("alloc"))
	return rv
}

func CNNConvolutionTranspose_Alloc() CNNConvolutionTranspose {
	return CNNConvolutionTransposeClass.Alloc()
}

func (cc _CNNConvolutionTransposeClass) New() CNNConvolutionTranspose {
	rv := objc.Call[CNNConvolutionTranspose](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNConvolutionTranspose() CNNConvolutionTranspose {
	return CNNConvolutionTransposeClass.New()
}

func (c_ CNNConvolutionTranspose) Init() CNNConvolutionTranspose {
	rv := objc.Call[CNNConvolutionTranspose](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNConvolutionTranspose) InitWithDevice(device metal.PDevice) CNNConvolutionTranspose {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNConvolutionTranspose](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNConvolutionTransposeWithDevice(device metal.PDevice) CNNConvolutionTranspose {
	instance := CNNConvolutionTransposeClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNConvolutionTranspose) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNConvolutionTranspose {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNConvolutionTranspose](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNConvolutionTranspose_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNConvolutionTranspose {
	instance := CNNConvolutionTransposeClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131773-reloadweightsandbiasesfromdataso?language=objc
func (c_ CNNConvolutionTranspose) ReloadWeightsAndBiasesFromDataSource() {
	objc.Call[objc.Void](c_, objc.Sel("reloadWeightsAndBiasesFromDataSource"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2942406-encodebatchtocommandbuffer?language=objc
func (c_ CNNConvolutionTranspose) EncodeBatchToCommandBufferSourceImagesConvolutionGradientStates(commandBuffer metal.PCommandBuffer, sourceImage *foundation.Array, convolutionGradientState *foundation.Array) *foundation.Array {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[*foundation.Array](c_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:convolutionGradientStates:"), po0, sourceImage, convolutionGradientState)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2942406-encodebatchtocommandbuffer?language=objc
func (c_ CNNConvolutionTranspose) EncodeBatchToCommandBufferObjectSourceImagesConvolutionGradientStates(commandBufferObject objc.IObject, sourceImage *foundation.Array, convolutionGradientState *foundation.Array) *foundation.Array {
	rv := objc.Call[*foundation.Array](c_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:convolutionGradientStates:"), objc.Ptr(commandBufferObject), sourceImage, convolutionGradientState)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131774-reloadweightsandbiaseswithcomman?language=objc
func (c_ CNNConvolutionTranspose) ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer metal.PCommandBuffer, state ICNNConvolutionWeightsAndBiasesState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](c_, objc.Sel("reloadWeightsAndBiasesWithCommandBuffer:state:"), po0, objc.Ptr(state))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131774-reloadweightsandbiaseswithcomman?language=objc
func (c_ CNNConvolutionTranspose) ReloadWeightsAndBiasesWithCommandBufferObjectState(commandBufferObject objc.IObject, state ICNNConvolutionWeightsAndBiasesState) {
	objc.Call[objc.Void](c_, objc.Sel("reloadWeightsAndBiasesWithCommandBuffer:state:"), objc.Ptr(commandBufferObject), objc.Ptr(state))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131772-exportweightsandbiaseswithcomman?language=objc
func (c_ CNNConvolutionTranspose) ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary(commandBuffer metal.PCommandBuffer, resultStateCanBeTemporary bool) CNNConvolutionWeightsAndBiasesState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNConvolutionWeightsAndBiasesState](c_, objc.Sel("exportWeightsAndBiasesWithCommandBuffer:resultStateCanBeTemporary:"), po0, resultStateCanBeTemporary)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131772-exportweightsandbiaseswithcomman?language=objc
func (c_ CNNConvolutionTranspose) ExportWeightsAndBiasesWithCommandBufferObjectResultStateCanBeTemporary(commandBufferObject objc.IObject, resultStateCanBeTemporary bool) CNNConvolutionWeightsAndBiasesState {
	rv := objc.Call[CNNConvolutionWeightsAndBiasesState](c_, objc.Sel("exportWeightsAndBiasesWithCommandBuffer:resultStateCanBeTemporary:"), objc.Ptr(commandBufferObject), resultStateCanBeTemporary)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2942409-encodetocommandbuffer?language=objc
func (c_ CNNConvolutionTranspose) EncodeToCommandBufferSourceImageConvolutionGradientState(commandBuffer metal.PCommandBuffer, sourceImage IImage, convolutionGradientState ICNNConvolutionGradientState) Image {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[Image](c_, objc.Sel("encodeToCommandBuffer:sourceImage:convolutionGradientState:"), po0, objc.Ptr(sourceImage), objc.Ptr(convolutionGradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2942409-encodetocommandbuffer?language=objc
func (c_ CNNConvolutionTranspose) EncodeToCommandBufferObjectSourceImageConvolutionGradientState(commandBufferObject objc.IObject, sourceImage IImage, convolutionGradientState ICNNConvolutionGradientState) Image {
	rv := objc.Call[Image](c_, objc.Sel("encodeToCommandBuffer:sourceImage:convolutionGradientState:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceImage), objc.Ptr(convolutionGradientState))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/3131769-datasource?language=objc
func (c_ CNNConvolutionTranspose) DataSource() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](c_, objc.Sel("dataSource"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867016-outputfeaturechannels?language=objc
func (c_ CNNConvolutionTranspose) OutputFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("outputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867174-inputfeaturechannels?language=objc
func (c_ CNNConvolutionTranspose) InputFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("inputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867176-kerneloffsetx?language=objc
func (c_ CNNConvolutionTranspose) KernelOffsetX() int {
	rv := objc.Call[int](c_, objc.Sel("kernelOffsetX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867176-kerneloffsetx?language=objc
func (c_ CNNConvolutionTranspose) SetKernelOffsetX(value int) {
	objc.Call[objc.Void](c_, objc.Sel("setKernelOffsetX:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867099-groups?language=objc
func (c_ CNNConvolutionTranspose) Groups() uint {
	rv := objc.Call[uint](c_, objc.Sel("groups"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867086-kerneloffsety?language=objc
func (c_ CNNConvolutionTranspose) KernelOffsetY() int {
	rv := objc.Call[int](c_, objc.Sel("kernelOffsetY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2867086-kerneloffsety?language=objc
func (c_ CNNConvolutionTranspose) SetKernelOffsetY(value int) {
	objc.Call[objc.Void](c_, objc.Sel("setKernelOffsetY:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2951924-accumulatorprecisionoption?language=objc
func (c_ CNNConvolutionTranspose) AccumulatorPrecisionOption() NNConvolutionAccumulatorPrecisionOption {
	rv := objc.Call[NNConvolutionAccumulatorPrecisionOption](c_, objc.Sel("accumulatorPrecisionOption"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontranspose/2951924-accumulatorprecisionoption?language=objc
func (c_ CNNConvolutionTranspose) SetAccumulatorPrecisionOption(value NNConvolutionAccumulatorPrecisionOption) {
	objc.Call[objc.Void](c_, objc.Sel("setAccumulatorPrecisionOption:"), value)
}
