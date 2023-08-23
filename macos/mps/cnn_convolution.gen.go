// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNConvolution] class.
var CNNConvolutionClass = _CNNConvolutionClass{objc.GetClass("MPSCNNConvolution")}

type _CNNConvolutionClass struct {
	objc.Class
}

// An interface definition for the [CNNConvolution] class.
type ICNNConvolution interface {
	ICNNKernel
	ReloadWeightsAndBiasesFromDataSource()
	ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer metal.PCommandBuffer, state ICNNConvolutionWeightsAndBiasesState)
	ReloadWeightsAndBiasesWithCommandBufferObjectState(commandBufferObject objc.IObject, state ICNNConvolutionWeightsAndBiasesState)
	ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary(commandBuffer metal.PCommandBuffer, resultStateCanBeTemporary bool) CNNConvolutionWeightsAndBiasesState
	ExportWeightsAndBiasesWithCommandBufferObjectResultStateCanBeTemporary(commandBufferObject objc.IObject, resultStateCanBeTemporary bool) CNNConvolutionWeightsAndBiasesState
	DataSource() CNNConvolutionDataSourceWrapper
	ChannelMultiplier() uint
	OutputFeatureChannels() uint
	InputFeatureChannels() uint
	Neuron() CNNNeuron
	Groups() uint
	SubPixelScaleFactor() uint
	AccumulatorPrecisionOption() NNConvolutionAccumulatorPrecisionOption
	SetAccumulatorPrecisionOption(value NNConvolutionAccumulatorPrecisionOption)
	FusedNeuronDescriptor() NNNeuronDescriptor
}

// A convolution kernel that convolves the input image with a set of filters, with each producing one feature map in the output image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution?language=objc
type CNNConvolution struct {
	CNNKernel
}

func CNNConvolutionFrom(ptr unsafe.Pointer) CNNConvolution {
	return CNNConvolution{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (c_ CNNConvolution) InitWithDeviceWeights(device metal.PDevice, weights PCNNConvolutionDataSource) CNNConvolution {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNConvolution](c_, objc.Sel("initWithDevice:weights:"), po0, po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2867092-initwithdevice?language=objc
func NewCNNConvolutionWithDeviceWeights(device metal.PDevice, weights PCNNConvolutionDataSource) CNNConvolution {
	instance := CNNConvolutionClass.Alloc().InitWithDeviceWeights(device, weights)
	instance.Autorelease()
	return instance
}

func (cc _CNNConvolutionClass) Alloc() CNNConvolution {
	rv := objc.Call[CNNConvolution](cc, objc.Sel("alloc"))
	return rv
}

func CNNConvolution_Alloc() CNNConvolution {
	return CNNConvolutionClass.Alloc()
}

func (cc _CNNConvolutionClass) New() CNNConvolution {
	rv := objc.Call[CNNConvolution](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNConvolution() CNNConvolution {
	return CNNConvolutionClass.New()
}

func (c_ CNNConvolution) Init() CNNConvolution {
	rv := objc.Call[CNNConvolution](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNConvolution) InitWithDevice(device metal.PDevice) CNNConvolution {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNConvolution](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNConvolutionWithDevice(device metal.PDevice) CNNConvolution {
	instance := CNNConvolutionClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNConvolution) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNConvolution {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNConvolution](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNConvolution_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNConvolution {
	instance := CNNConvolutionClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2966657-reloadweightsandbiasesfromdataso?language=objc
func (c_ CNNConvolution) ReloadWeightsAndBiasesFromDataSource() {
	objc.Call[objc.Void](c_, objc.Sel("reloadWeightsAndBiasesFromDataSource"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2953962-reloadweightsandbiaseswithcomman?language=objc
func (c_ CNNConvolution) ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer metal.PCommandBuffer, state ICNNConvolutionWeightsAndBiasesState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](c_, objc.Sel("reloadWeightsAndBiasesWithCommandBuffer:state:"), po0, objc.Ptr(state))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2953962-reloadweightsandbiaseswithcomman?language=objc
func (c_ CNNConvolution) ReloadWeightsAndBiasesWithCommandBufferObjectState(commandBufferObject objc.IObject, state ICNNConvolutionWeightsAndBiasesState) {
	objc.Call[objc.Void](c_, objc.Sel("reloadWeightsAndBiasesWithCommandBuffer:state:"), objc.Ptr(commandBufferObject), objc.Ptr(state))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2953001-exportweightsandbiaseswithcomman?language=objc
func (c_ CNNConvolution) ExportWeightsAndBiasesWithCommandBufferResultStateCanBeTemporary(commandBuffer metal.PCommandBuffer, resultStateCanBeTemporary bool) CNNConvolutionWeightsAndBiasesState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNConvolutionWeightsAndBiasesState](c_, objc.Sel("exportWeightsAndBiasesWithCommandBuffer:resultStateCanBeTemporary:"), po0, resultStateCanBeTemporary)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2953001-exportweightsandbiaseswithcomman?language=objc
func (c_ CNNConvolution) ExportWeightsAndBiasesWithCommandBufferObjectResultStateCanBeTemporary(commandBufferObject objc.IObject, resultStateCanBeTemporary bool) CNNConvolutionWeightsAndBiasesState {
	rv := objc.Call[CNNConvolutionWeightsAndBiasesState](c_, objc.Sel("exportWeightsAndBiasesWithCommandBuffer:resultStateCanBeTemporary:"), objc.Ptr(commandBufferObject), resultStateCanBeTemporary)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2953961-datasource?language=objc
func (c_ CNNConvolution) DataSource() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](c_, objc.Sel("dataSource"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2919729-channelmultiplier?language=objc
func (c_ CNNConvolution) ChannelMultiplier() uint {
	rv := objc.Call[uint](c_, objc.Sel("channelMultiplier"))
	return rv
}

// The number of feature channels per pixel in the output image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/1845271-outputfeaturechannels?language=objc
func (c_ CNNConvolution) OutputFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("outputFeatureChannels"))
	return rv
}

// The number of feature channels per pixel in the input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/1845268-inputfeaturechannels?language=objc
func (c_ CNNConvolution) InputFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("inputFeatureChannels"))
	return rv
}

// The neuron filter to be applied as part of the convolution operation. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/1845274-neuron?language=objc
func (c_ CNNConvolution) Neuron() CNNNeuron {
	rv := objc.Call[CNNNeuron](c_, objc.Sel("neuron"))
	return rv
}

// The number of groups that the input and output channels are divided into. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/1845269-groups?language=objc
func (c_ CNNConvolution) Groups() uint {
	rv := objc.Call[uint](c_, objc.Sel("groups"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2873341-subpixelscalefactor?language=objc
func (c_ CNNConvolution) SubPixelScaleFactor() uint {
	rv := objc.Call[uint](c_, objc.Sel("subPixelScaleFactor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2942410-accumulatorprecisionoption?language=objc
func (c_ CNNConvolution) AccumulatorPrecisionOption() NNConvolutionAccumulatorPrecisionOption {
	rv := objc.Call[NNConvolutionAccumulatorPrecisionOption](c_, objc.Sel("accumulatorPrecisionOption"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/2942410-accumulatorprecisionoption?language=objc
func (c_ CNNConvolution) SetAccumulatorPrecisionOption(value NNConvolutionAccumulatorPrecisionOption) {
	objc.Call[objc.Void](c_, objc.Sel("setAccumulatorPrecisionOption:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolution/3013776-fusedneurondescriptor?language=objc
func (c_ CNNConvolution) FusedNeuronDescriptor() NNNeuronDescriptor {
	rv := objc.Call[NNNeuronDescriptor](c_, objc.Sel("fusedNeuronDescriptor"))
	return rv
}
