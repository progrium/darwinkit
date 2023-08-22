// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNConvolutionTransposeGradient] class.
var CNNConvolutionTransposeGradientClass = _CNNConvolutionTransposeGradientClass{objc.GetClass("MPSCNNConvolutionTransposeGradient")}

type _CNNConvolutionTransposeGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNConvolutionTransposeGradient] class.
type ICNNConvolutionTransposeGradient interface {
	ICNNGradientKernel
	ReloadWeightsAndBiasesFromDataSource()
	ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer metal.PCommandBuffer, state ICNNConvolutionWeightsAndBiasesState)
	ReloadWeightsAndBiasesWithCommandBufferObjectState(commandBufferObject objc.IObject, state ICNNConvolutionWeightsAndBiasesState)
	DataSource() CNNConvolutionDataSourceWrapper
	SourceImageFeatureChannels() uint
	GradientOption() CNNConvolutionGradientOption
	SetGradientOption(value CNNConvolutionGradientOption)
	SourceGradientFeatureChannels() uint
	Groups() uint
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient?language=objc
type CNNConvolutionTransposeGradient struct {
	CNNGradientKernel
}

func CNNConvolutionTransposeGradientFrom(ptr unsafe.Pointer) CNNConvolutionTransposeGradient {
	return CNNConvolutionTransposeGradient{
		CNNGradientKernel: CNNGradientKernelFrom(ptr),
	}
}

func (c_ CNNConvolutionTransposeGradient) InitWithDeviceWeights(device metal.PDevice, weights PCNNConvolutionDataSource) CNNConvolutionTransposeGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNConvolutionTransposeGradient](c_, objc.Sel("initWithDevice:weights:"), po0, po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131784-initwithdevice?language=objc
func NewCNNConvolutionTransposeGradientWithDeviceWeights(device metal.PDevice, weights PCNNConvolutionDataSource) CNNConvolutionTransposeGradient {
	instance := CNNConvolutionTransposeGradientClass.Alloc().InitWithDeviceWeights(device, weights)
	instance.Autorelease()
	return instance
}

func (cc _CNNConvolutionTransposeGradientClass) Alloc() CNNConvolutionTransposeGradient {
	rv := objc.Call[CNNConvolutionTransposeGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNConvolutionTransposeGradient_Alloc() CNNConvolutionTransposeGradient {
	return CNNConvolutionTransposeGradientClass.Alloc()
}

func (cc _CNNConvolutionTransposeGradientClass) New() CNNConvolutionTransposeGradient {
	rv := objc.Call[CNNConvolutionTransposeGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNConvolutionTransposeGradient() CNNConvolutionTransposeGradient {
	return CNNConvolutionTransposeGradientClass.New()
}

func (c_ CNNConvolutionTransposeGradient) Init() CNNConvolutionTransposeGradient {
	rv := objc.Call[CNNConvolutionTransposeGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNConvolutionTransposeGradient) InitWithDevice(device metal.PDevice) CNNConvolutionTransposeGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNConvolutionTransposeGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNConvolutionTransposeGradientWithDevice(device metal.PDevice) CNNConvolutionTransposeGradient {
	instance := CNNConvolutionTransposeGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNConvolutionTransposeGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNConvolutionTransposeGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNConvolutionTransposeGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNConvolutionTransposeGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNConvolutionTransposeGradient {
	instance := CNNConvolutionTransposeGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131785-reloadweightsandbiasesfromdataso?language=objc
func (c_ CNNConvolutionTransposeGradient) ReloadWeightsAndBiasesFromDataSource() {
	objc.Call[objc.Void](c_, objc.Sel("reloadWeightsAndBiasesFromDataSource"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131786-reloadweightsandbiaseswithcomman?language=objc
func (c_ CNNConvolutionTransposeGradient) ReloadWeightsAndBiasesWithCommandBufferState(commandBuffer metal.PCommandBuffer, state ICNNConvolutionWeightsAndBiasesState) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](c_, objc.Sel("reloadWeightsAndBiasesWithCommandBuffer:state:"), po0, objc.Ptr(state))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131786-reloadweightsandbiaseswithcomman?language=objc
func (c_ CNNConvolutionTransposeGradient) ReloadWeightsAndBiasesWithCommandBufferObjectState(commandBufferObject objc.IObject, state ICNNConvolutionWeightsAndBiasesState) {
	objc.Call[objc.Void](c_, objc.Sel("reloadWeightsAndBiasesWithCommandBuffer:state:"), objc.Ptr(commandBufferObject), objc.Ptr(state))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131780-datasource?language=objc
func (c_ CNNConvolutionTransposeGradient) DataSource() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](c_, objc.Sel("dataSource"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131788-sourceimagefeaturechannels?language=objc
func (c_ CNNConvolutionTransposeGradient) SourceImageFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("sourceImageFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131781-gradientoption?language=objc
func (c_ CNNConvolutionTransposeGradient) GradientOption() CNNConvolutionGradientOption {
	rv := objc.Call[CNNConvolutionGradientOption](c_, objc.Sel("gradientOption"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131781-gradientoption?language=objc
func (c_ CNNConvolutionTransposeGradient) SetGradientOption(value CNNConvolutionGradientOption) {
	objc.Call[objc.Void](c_, objc.Sel("setGradientOption:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131787-sourcegradientfeaturechannels?language=objc
func (c_ CNNConvolutionTransposeGradient) SourceGradientFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("sourceGradientFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiontransposegradient/3131782-groups?language=objc
func (c_ CNNConvolutionTransposeGradient) Groups() uint {
	rv := objc.Call[uint](c_, objc.Sel("groups"))
	return rv
}
