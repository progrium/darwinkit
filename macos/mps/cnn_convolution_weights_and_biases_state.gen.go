// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNConvolutionWeightsAndBiasesState] class.
var CNNConvolutionWeightsAndBiasesStateClass = _CNNConvolutionWeightsAndBiasesStateClass{objc.GetClass("MPSCNNConvolutionWeightsAndBiasesState")}

type _CNNConvolutionWeightsAndBiasesStateClass struct {
	objc.Class
}

// An interface definition for the [CNNConvolutionWeightsAndBiasesState] class.
type ICNNConvolutionWeightsAndBiasesState interface {
	IState
	WeightsOffset() uint
	Weights() metal.BufferWrapper
	BiasesOffset() uint
	Biases() metal.BufferWrapper
}

// A class that stores weights and biases. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate?language=objc
type CNNConvolutionWeightsAndBiasesState struct {
	State
}

func CNNConvolutionWeightsAndBiasesStateFrom(ptr unsafe.Pointer) CNNConvolutionWeightsAndBiasesState {
	return CNNConvolutionWeightsAndBiasesState{
		State: StateFrom(ptr),
	}
}

func (c_ CNNConvolutionWeightsAndBiasesState) InitWithDeviceCnnConvolutionDescriptor(device metal.PDevice, descriptor ICNNConvolutionDescriptor) CNNConvolutionWeightsAndBiasesState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNConvolutionWeightsAndBiasesState](c_, objc.Sel("initWithDevice:cnnConvolutionDescriptor:"), po0, objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/2953004-initwithdevice?language=objc
func NewCNNConvolutionWeightsAndBiasesStateWithDeviceCnnConvolutionDescriptor(device metal.PDevice, descriptor ICNNConvolutionDescriptor) CNNConvolutionWeightsAndBiasesState {
	instance := CNNConvolutionWeightsAndBiasesStateClass.Alloc().InitWithDeviceCnnConvolutionDescriptor(device, descriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNConvolutionWeightsAndBiasesState) InitWithWeightsBiases(weights metal.PBuffer, biases metal.PBuffer) CNNConvolutionWeightsAndBiasesState {
	po0 := objc.WrapAsProtocol("MTLBuffer", weights)
	po1 := objc.WrapAsProtocol("MTLBuffer", biases)
	rv := objc.Call[CNNConvolutionWeightsAndBiasesState](c_, objc.Sel("initWithWeights:biases:"), po0, po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/2953008-initwithweights?language=objc
func NewCNNConvolutionWeightsAndBiasesStateWithWeightsBiases(weights metal.PBuffer, biases metal.PBuffer) CNNConvolutionWeightsAndBiasesState {
	instance := CNNConvolutionWeightsAndBiasesStateClass.Alloc().InitWithWeightsBiases(weights, biases)
	instance.Autorelease()
	return instance
}

func (cc _CNNConvolutionWeightsAndBiasesStateClass) TemporaryCNNConvolutionWeightsAndBiasesStateWithCommandBufferCnnConvolutionDescriptor(commandBuffer metal.PCommandBuffer, descriptor ICNNConvolutionDescriptor) CNNConvolutionWeightsAndBiasesState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNConvolutionWeightsAndBiasesState](cc, objc.Sel("temporaryCNNConvolutionWeightsAndBiasesStateWithCommandBuffer:cnnConvolutionDescriptor:"), po0, objc.Ptr(descriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/2953005-temporarycnnconvolutionweightsan?language=objc
func CNNConvolutionWeightsAndBiasesState_TemporaryCNNConvolutionWeightsAndBiasesStateWithCommandBufferCnnConvolutionDescriptor(commandBuffer metal.PCommandBuffer, descriptor ICNNConvolutionDescriptor) CNNConvolutionWeightsAndBiasesState {
	return CNNConvolutionWeightsAndBiasesStateClass.TemporaryCNNConvolutionWeightsAndBiasesStateWithCommandBufferCnnConvolutionDescriptor(commandBuffer, descriptor)
}

func (cc _CNNConvolutionWeightsAndBiasesStateClass) Alloc() CNNConvolutionWeightsAndBiasesState {
	rv := objc.Call[CNNConvolutionWeightsAndBiasesState](cc, objc.Sel("alloc"))
	return rv
}

func CNNConvolutionWeightsAndBiasesState_Alloc() CNNConvolutionWeightsAndBiasesState {
	return CNNConvolutionWeightsAndBiasesStateClass.Alloc()
}

func (cc _CNNConvolutionWeightsAndBiasesStateClass) New() CNNConvolutionWeightsAndBiasesState {
	rv := objc.Call[CNNConvolutionWeightsAndBiasesState](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNConvolutionWeightsAndBiasesState() CNNConvolutionWeightsAndBiasesState {
	return CNNConvolutionWeightsAndBiasesStateClass.New()
}

func (c_ CNNConvolutionWeightsAndBiasesState) Init() CNNConvolutionWeightsAndBiasesState {
	rv := objc.Call[CNNConvolutionWeightsAndBiasesState](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNConvolutionWeightsAndBiasesState) InitWithResources(resources []metal.PResource) CNNConvolutionWeightsAndBiasesState {
	rv := objc.Call[CNNConvolutionWeightsAndBiasesState](c_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewCNNConvolutionWeightsAndBiasesStateWithResources(resources []metal.PResource) CNNConvolutionWeightsAndBiasesState {
	instance := CNNConvolutionWeightsAndBiasesStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (c_ CNNConvolutionWeightsAndBiasesState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNConvolutionWeightsAndBiasesState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNConvolutionWeightsAndBiasesState](c_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewCNNConvolutionWeightsAndBiasesStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNConvolutionWeightsAndBiasesState {
	instance := CNNConvolutionWeightsAndBiasesStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (c_ CNNConvolutionWeightsAndBiasesState) InitWithResource(resource metal.PResource) CNNConvolutionWeightsAndBiasesState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[CNNConvolutionWeightsAndBiasesState](c_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewCNNConvolutionWeightsAndBiasesStateWithResource(resource metal.PResource) CNNConvolutionWeightsAndBiasesState {
	instance := CNNConvolutionWeightsAndBiasesStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (cc _CNNConvolutionWeightsAndBiasesStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNConvolutionWeightsAndBiasesState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNConvolutionWeightsAndBiasesState](cc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func CNNConvolutionWeightsAndBiasesState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNConvolutionWeightsAndBiasesState {
	return CNNConvolutionWeightsAndBiasesStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/3325844-weightsoffset?language=objc
func (c_ CNNConvolutionWeightsAndBiasesState) WeightsOffset() uint {
	rv := objc.Call[uint](c_, objc.Sel("weightsOffset"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/2953006-weights?language=objc
func (c_ CNNConvolutionWeightsAndBiasesState) Weights() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("weights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/3325842-biasesoffset?language=objc
func (c_ CNNConvolutionWeightsAndBiasesState) BiasesOffset() uint {
	rv := objc.Call[uint](c_, objc.Sel("biasesOffset"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutionweightsandbiasesstate/2953002-biases?language=objc
func (c_ CNNConvolutionWeightsAndBiasesState) Biases() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("biases"))
	return rv
}
