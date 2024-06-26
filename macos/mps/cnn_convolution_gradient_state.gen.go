// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [CNNConvolutionGradientState] class.
var CNNConvolutionGradientStateClass = _CNNConvolutionGradientStateClass{objc.GetClass("MPSCNNConvolutionGradientState")}

type _CNNConvolutionGradientStateClass struct {
	objc.Class
}

// An interface definition for the [CNNConvolutionGradientState] class.
type ICNNConvolutionGradientState interface {
	INNGradientState
	GradientForWeightsLayout() CNNConvolutionWeightsLayout
	GradientForBiases() metal.BufferObject
	GradientForWeights() metal.BufferObject
	Convolution() CNNConvolution
}

// An object that exposes a gradient convolution kernel's gradient with respect to weights and biases. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientstate?language=objc
type CNNConvolutionGradientState struct {
	NNGradientState
}

func CNNConvolutionGradientStateFrom(ptr unsafe.Pointer) CNNConvolutionGradientState {
	return CNNConvolutionGradientState{
		NNGradientState: NNGradientStateFrom(ptr),
	}
}

func (cc _CNNConvolutionGradientStateClass) Alloc() CNNConvolutionGradientState {
	rv := objc.Call[CNNConvolutionGradientState](cc, objc.Sel("alloc"))
	return rv
}

func (cc _CNNConvolutionGradientStateClass) New() CNNConvolutionGradientState {
	rv := objc.Call[CNNConvolutionGradientState](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNConvolutionGradientState() CNNConvolutionGradientState {
	return CNNConvolutionGradientStateClass.New()
}

func (c_ CNNConvolutionGradientState) Init() CNNConvolutionGradientState {
	rv := objc.Call[CNNConvolutionGradientState](c_, objc.Sel("init"))
	return rv
}

func (cc _CNNConvolutionGradientStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNConvolutionGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNConvolutionGradientState](cc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, resourceList)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func CNNConvolutionGradientState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNConvolutionGradientState {
	return CNNConvolutionGradientStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}

func (c_ CNNConvolutionGradientState) InitWithDeviceTextureDescriptor(device metal.PDevice, descriptor metal.ITextureDescriptor) CNNConvolutionGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNConvolutionGradientState](c_, objc.Sel("initWithDevice:textureDescriptor:"), po0, descriptor)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942400-initwithdevice?language=objc
func NewCNNConvolutionGradientStateWithDeviceTextureDescriptor(device metal.PDevice, descriptor metal.ITextureDescriptor) CNNConvolutionGradientState {
	instance := CNNConvolutionGradientStateClass.Alloc().InitWithDeviceTextureDescriptor(device, descriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNConvolutionGradientState) InitWithResource(resource metal.PResource) CNNConvolutionGradientState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[CNNConvolutionGradientState](c_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewCNNConvolutionGradientStateWithResource(resource metal.PResource) CNNConvolutionGradientState {
	instance := CNNConvolutionGradientStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (c_ CNNConvolutionGradientState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNConvolutionGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNConvolutionGradientState](c_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewCNNConvolutionGradientStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNConvolutionGradientState {
	instance := CNNConvolutionGradientStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (c_ CNNConvolutionGradientState) InitWithResources(resources []metal.PResource) CNNConvolutionGradientState {
	rv := objc.Call[CNNConvolutionGradientState](c_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewCNNConvolutionGradientStateWithResources(resources []metal.PResource) CNNConvolutionGradientState {
	instance := CNNConvolutionGradientStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (cc _CNNConvolutionGradientStateClass) TemporaryStateWithCommandBufferBufferSize(cmdBuf metal.PCommandBuffer, bufferSize uint) CNNConvolutionGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[CNNConvolutionGradientState](cc, objc.Sel("temporaryStateWithCommandBuffer:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942391-temporarystatewithcommandbuffer?language=objc
func CNNConvolutionGradientState_TemporaryStateWithCommandBufferBufferSize(cmdBuf metal.PCommandBuffer, bufferSize uint) CNNConvolutionGradientState {
	return CNNConvolutionGradientStateClass.TemporaryStateWithCommandBufferBufferSize(cmdBuf, bufferSize)
}

func (cc _CNNConvolutionGradientStateClass) TemporaryStateWithCommandBuffer(cmdBuf metal.PCommandBuffer) CNNConvolutionGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[CNNConvolutionGradientState](cc, objc.Sel("temporaryStateWithCommandBuffer:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942393-temporarystatewithcommandbuffer?language=objc
func CNNConvolutionGradientState_TemporaryStateWithCommandBuffer(cmdBuf metal.PCommandBuffer) CNNConvolutionGradientState {
	return CNNConvolutionGradientStateClass.TemporaryStateWithCommandBuffer(cmdBuf)
}

func (cc _CNNConvolutionGradientStateClass) TemporaryStateWithCommandBufferTextureDescriptor(cmdBuf metal.PCommandBuffer, descriptor metal.ITextureDescriptor) CNNConvolutionGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[CNNConvolutionGradientState](cc, objc.Sel("temporaryStateWithCommandBuffer:textureDescriptor:"), po0, descriptor)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942395-temporarystatewithcommandbuffer?language=objc
func CNNConvolutionGradientState_TemporaryStateWithCommandBufferTextureDescriptor(cmdBuf metal.PCommandBuffer, descriptor metal.ITextureDescriptor) CNNConvolutionGradientState {
	return CNNConvolutionGradientStateClass.TemporaryStateWithCommandBufferTextureDescriptor(cmdBuf, descriptor)
}

func (c_ CNNConvolutionGradientState) InitWithDeviceResourceList(device metal.PDevice, resourceList IStateResourceList) CNNConvolutionGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNConvolutionGradientState](c_, objc.Sel("initWithDevice:resourceList:"), po0, resourceList)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947908-initwithdevice?language=objc
func NewCNNConvolutionGradientStateWithDeviceResourceList(device metal.PDevice, resourceList IStateResourceList) CNNConvolutionGradientState {
	instance := CNNConvolutionGradientStateClass.Alloc().InitWithDeviceResourceList(device, resourceList)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientstate/3325841-gradientforweightslayout?language=objc
func (c_ CNNConvolutionGradientState) GradientForWeightsLayout() CNNConvolutionWeightsLayout {
	rv := objc.Call[CNNConvolutionWeightsLayout](c_, objc.Sel("gradientForWeightsLayout"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientstate/2947887-gradientforbiases?language=objc
func (c_ CNNConvolutionGradientState) GradientForBiases() metal.BufferObject {
	rv := objc.Call[metal.BufferObject](c_, objc.Sel("gradientForBiases"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientstate/2947889-gradientforweights?language=objc
func (c_ CNNConvolutionGradientState) GradientForWeights() metal.BufferObject {
	rv := objc.Call[metal.BufferObject](c_, objc.Sel("gradientForWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnconvolutiongradientstate/2953958-convolution?language=objc
func (c_ CNNConvolutionGradientState) Convolution() CNNConvolution {
	rv := objc.Call[CNNConvolution](c_, objc.Sel("convolution"))
	return rv
}
