// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/kernel"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The protocol that provides convolution filter weights and bias terms. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondatasource?language=objc
type PCNNConvolutionDataSource interface {
	// optional
	RangesForUInt8Kernel() *kernel.Vector_float2
	HasRangesForUInt8Kernel() bool

	// optional
	UpdateWithGradientStateSourceState(gradientState CNNConvolutionGradientState, sourceState CNNConvolutionWeightsAndBiasesState) bool
	HasUpdateWithGradientStateSourceState() bool

	// optional
	WeightsLayout() CNNConvolutionWeightsLayout
	HasWeightsLayout() bool

	// optional
	Weights() unsafe.Pointer
	HasWeights() bool

	// optional
	WeightsQuantizationType() CNNWeightsQuantizationType
	HasWeightsQuantizationType() bool

	// optional
	LookupTableForUInt8Kernel() *float64
	HasLookupTableForUInt8Kernel() bool

	// optional
	BiasTerms() *float64
	HasBiasTerms() bool

	// optional
	Purge()
	HasPurge() bool

	// optional
	DataType() DataType
	HasDataType() bool

	// optional
	Descriptor() ICNNConvolutionDescriptor
	HasDescriptor() bool

	// optional
	CopyWithZoneDevice(zone unsafe.Pointer, device metal.DeviceWrapper) objc.IObject
	HasCopyWithZoneDevice() bool

	// optional
	Load() bool
	HasLoad() bool

	// optional
	Label() string
	HasLabel() bool

	// optional
	UpdateWithCommandBufferGradientStateSourceState(commandBuffer metal.CommandBufferWrapper, gradientState CNNConvolutionGradientState, sourceState CNNConvolutionWeightsAndBiasesState) ICNNConvolutionWeightsAndBiasesState
	HasUpdateWithCommandBufferGradientStateSourceState() bool

	// optional
	KernelWeightsDataType() DataType
	HasKernelWeightsDataType() bool
}

// A concrete type wrapper for the [PCNNConvolutionDataSource] protocol.
type CNNConvolutionDataSourceWrapper struct {
	objc.Object
}

func (c_ CNNConvolutionDataSourceWrapper) HasRangesForUInt8Kernel() bool {
	return c_.RespondsToSelector(objc.Sel("rangesForUInt8Kernel"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondatasource/2867145-rangesforuint8kernel?language=objc
func (c_ CNNConvolutionDataSourceWrapper) RangesForUInt8Kernel() *kernel.Vector_float2 {
	rv := objc.Call[*kernel.Vector_float2](c_, objc.Sel("rangesForUInt8Kernel"))
	return rv
}

func (c_ CNNConvolutionDataSourceWrapper) HasUpdateWithGradientStateSourceState() bool {
	return c_.RespondsToSelector(objc.Sel("updateWithGradientState:sourceState:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondatasource/2953009-updatewithgradientstate?language=objc
func (c_ CNNConvolutionDataSourceWrapper) UpdateWithGradientStateSourceState(gradientState ICNNConvolutionGradientState, sourceState ICNNConvolutionWeightsAndBiasesState) bool {
	rv := objc.Call[bool](c_, objc.Sel("updateWithGradientState:sourceState:"), objc.Ptr(gradientState), objc.Ptr(sourceState))
	return rv
}

func (c_ CNNConvolutionDataSourceWrapper) HasWeightsLayout() bool {
	return c_.RespondsToSelector(objc.Sel("weightsLayout"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondatasource/3325840-weightslayout?language=objc
func (c_ CNNConvolutionDataSourceWrapper) WeightsLayout() CNNConvolutionWeightsLayout {
	rv := objc.Call[CNNConvolutionWeightsLayout](c_, objc.Sel("weightsLayout"))
	return rv
}

func (c_ CNNConvolutionDataSourceWrapper) HasWeights() bool {
	return c_.RespondsToSelector(objc.Sel("weights"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondatasource/2867187-weights?language=objc
func (c_ CNNConvolutionDataSourceWrapper) Weights() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](c_, objc.Sel("weights"))
	return rv
}

func (c_ CNNConvolutionDataSourceWrapper) HasWeightsQuantizationType() bool {
	return c_.RespondsToSelector(objc.Sel("weightsQuantizationType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondatasource/2976466-weightsquantizationtype?language=objc
func (c_ CNNConvolutionDataSourceWrapper) WeightsQuantizationType() CNNWeightsQuantizationType {
	rv := objc.Call[CNNWeightsQuantizationType](c_, objc.Sel("weightsQuantizationType"))
	return rv
}

func (c_ CNNConvolutionDataSourceWrapper) HasLookupTableForUInt8Kernel() bool {
	return c_.RespondsToSelector(objc.Sel("lookupTableForUInt8Kernel"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondatasource/2867186-lookuptableforuint8kernel?language=objc
func (c_ CNNConvolutionDataSourceWrapper) LookupTableForUInt8Kernel() *float64 {
	rv := objc.Call[*float64](c_, objc.Sel("lookupTableForUInt8Kernel"))
	return rv
}

func (c_ CNNConvolutionDataSourceWrapper) HasBiasTerms() bool {
	return c_.RespondsToSelector(objc.Sel("biasTerms"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondatasource/2867023-biasterms?language=objc
func (c_ CNNConvolutionDataSourceWrapper) BiasTerms() *float64 {
	rv := objc.Call[*float64](c_, objc.Sel("biasTerms"))
	return rv
}

func (c_ CNNConvolutionDataSourceWrapper) HasPurge() bool {
	return c_.RespondsToSelector(objc.Sel("purge"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondatasource/2867134-purge?language=objc
func (c_ CNNConvolutionDataSourceWrapper) Purge() {
	objc.Call[objc.Void](c_, objc.Sel("purge"))
}

func (c_ CNNConvolutionDataSourceWrapper) HasDataType() bool {
	return c_.RespondsToSelector(objc.Sel("dataType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondatasource/2867139-datatype?language=objc
func (c_ CNNConvolutionDataSourceWrapper) DataType() DataType {
	rv := objc.Call[DataType](c_, objc.Sel("dataType"))
	return rv
}

func (c_ CNNConvolutionDataSourceWrapper) HasDescriptor() bool {
	return c_.RespondsToSelector(objc.Sel("descriptor"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondatasource/2867050-descriptor?language=objc
func (c_ CNNConvolutionDataSourceWrapper) Descriptor() CNNConvolutionDescriptor {
	rv := objc.Call[CNNConvolutionDescriptor](c_, objc.Sel("descriptor"))
	return rv
}

func (c_ CNNConvolutionDataSourceWrapper) HasCopyWithZoneDevice() bool {
	return c_.RespondsToSelector(objc.Sel("copyWithZone:device:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondatasource/3013778-copywithzone?language=objc
func (c_ CNNConvolutionDataSourceWrapper) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) objc.Object {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[objc.Object](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

func (c_ CNNConvolutionDataSourceWrapper) HasLoad() bool {
	return c_.RespondsToSelector(objc.Sel("load"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondatasource/2867049-load?language=objc
func (c_ CNNConvolutionDataSourceWrapper) Load() bool {
	rv := objc.Call[bool](c_, objc.Sel("load"))
	return rv
}

func (c_ CNNConvolutionDataSourceWrapper) HasLabel() bool {
	return c_.RespondsToSelector(objc.Sel("label"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondatasource/2881197-label?language=objc
func (c_ CNNConvolutionDataSourceWrapper) Label() string {
	rv := objc.Call[string](c_, objc.Sel("label"))
	return rv
}

func (c_ CNNConvolutionDataSourceWrapper) HasUpdateWithCommandBufferGradientStateSourceState() bool {
	return c_.RespondsToSelector(objc.Sel("updateWithCommandBuffer:gradientState:sourceState:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondatasource/2953007-updatewithcommandbuffer?language=objc
func (c_ CNNConvolutionDataSourceWrapper) UpdateWithCommandBufferGradientStateSourceState(commandBuffer metal.PCommandBuffer, gradientState ICNNConvolutionGradientState, sourceState ICNNConvolutionWeightsAndBiasesState) CNNConvolutionWeightsAndBiasesState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNConvolutionWeightsAndBiasesState](c_, objc.Sel("updateWithCommandBuffer:gradientState:sourceState:"), po0, objc.Ptr(gradientState), objc.Ptr(sourceState))
	return rv
}

func (c_ CNNConvolutionDataSourceWrapper) HasKernelWeightsDataType() bool {
	return c_.RespondsToSelector(objc.Sel("kernelWeightsDataType"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiondatasource/3564466-kernelweightsdatatype?language=objc
func (c_ CNNConvolutionDataSourceWrapper) KernelWeightsDataType() DataType {
	rv := objc.Call[DataType](c_, objc.Sel("kernelWeightsDataType"))
	return rv
}
