// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNormalizationGammaAndBetaState] class.
var CNNNormalizationGammaAndBetaStateClass = _CNNNormalizationGammaAndBetaStateClass{objc.GetClass("MPSCNNNormalizationGammaAndBetaState")}

type _CNNNormalizationGammaAndBetaStateClass struct {
	objc.Class
}

// An interface definition for the [CNNNormalizationGammaAndBetaState] class.
type ICNNNormalizationGammaAndBetaState interface {
	IState
	Beta() metal.BufferWrapper
	Gamma() metal.BufferWrapper
}

// An object that stores gamma and beta terms used to apply a scale and bias in instance- or batch-normalization operations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate?language=objc
type CNNNormalizationGammaAndBetaState struct {
	State
}

func CNNNormalizationGammaAndBetaStateFrom(ptr unsafe.Pointer) CNNNormalizationGammaAndBetaState {
	return CNNNormalizationGammaAndBetaState{
		State: StateFrom(ptr),
	}
}

func (c_ CNNNormalizationGammaAndBetaState) InitWithGammaBeta(gamma metal.PBuffer, beta metal.PBuffer) CNNNormalizationGammaAndBetaState {
	po0 := objc.WrapAsProtocol("MTLBuffer", gamma)
	po1 := objc.WrapAsProtocol("MTLBuffer", beta)
	rv := objc.Call[CNNNormalizationGammaAndBetaState](c_, objc.Sel("initWithGamma:beta:"), po0, po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate/2953936-initwithgamma?language=objc
func NewCNNNormalizationGammaAndBetaStateWithGammaBeta(gamma metal.PBuffer, beta metal.PBuffer) CNNNormalizationGammaAndBetaState {
	instance := CNNNormalizationGammaAndBetaStateClass.Alloc().InitWithGammaBeta(gamma, beta)
	instance.Autorelease()
	return instance
}

func (cc _CNNNormalizationGammaAndBetaStateClass) TemporaryStateWithCommandBufferNumberOfFeatureChannels(commandBuffer metal.PCommandBuffer, numberOfFeatureChannels uint) CNNNormalizationGammaAndBetaState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNNormalizationGammaAndBetaState](cc, objc.Sel("temporaryStateWithCommandBuffer:numberOfFeatureChannels:"), po0, numberOfFeatureChannels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate/2953937-temporarystatewithcommandbuffer?language=objc
func CNNNormalizationGammaAndBetaState_TemporaryStateWithCommandBufferNumberOfFeatureChannels(commandBuffer metal.PCommandBuffer, numberOfFeatureChannels uint) CNNNormalizationGammaAndBetaState {
	return CNNNormalizationGammaAndBetaStateClass.TemporaryStateWithCommandBufferNumberOfFeatureChannels(commandBuffer, numberOfFeatureChannels)
}

func (cc _CNNNormalizationGammaAndBetaStateClass) Alloc() CNNNormalizationGammaAndBetaState {
	rv := objc.Call[CNNNormalizationGammaAndBetaState](cc, objc.Sel("alloc"))
	return rv
}

func CNNNormalizationGammaAndBetaState_Alloc() CNNNormalizationGammaAndBetaState {
	return CNNNormalizationGammaAndBetaStateClass.Alloc()
}

func (cc _CNNNormalizationGammaAndBetaStateClass) New() CNNNormalizationGammaAndBetaState {
	rv := objc.Call[CNNNormalizationGammaAndBetaState](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNormalizationGammaAndBetaState() CNNNormalizationGammaAndBetaState {
	return CNNNormalizationGammaAndBetaStateClass.New()
}

func (c_ CNNNormalizationGammaAndBetaState) Init() CNNNormalizationGammaAndBetaState {
	rv := objc.Call[CNNNormalizationGammaAndBetaState](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNNormalizationGammaAndBetaState) InitWithResources(resources []metal.PResource) CNNNormalizationGammaAndBetaState {
	rv := objc.Call[CNNNormalizationGammaAndBetaState](c_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewCNNNormalizationGammaAndBetaStateWithResources(resources []metal.PResource) CNNNormalizationGammaAndBetaState {
	instance := CNNNormalizationGammaAndBetaStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (c_ CNNNormalizationGammaAndBetaState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNNormalizationGammaAndBetaState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNormalizationGammaAndBetaState](c_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewCNNNormalizationGammaAndBetaStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNNormalizationGammaAndBetaState {
	instance := CNNNormalizationGammaAndBetaStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (c_ CNNNormalizationGammaAndBetaState) InitWithResource(resource metal.PResource) CNNNormalizationGammaAndBetaState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[CNNNormalizationGammaAndBetaState](c_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewCNNNormalizationGammaAndBetaStateWithResource(resource metal.PResource) CNNNormalizationGammaAndBetaState {
	instance := CNNNormalizationGammaAndBetaStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (cc _CNNNormalizationGammaAndBetaStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNNormalizationGammaAndBetaState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNNormalizationGammaAndBetaState](cc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func CNNNormalizationGammaAndBetaState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNNormalizationGammaAndBetaState {
	return CNNNormalizationGammaAndBetaStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate/2953938-beta?language=objc
func (c_ CNNNormalizationGammaAndBetaState) Beta() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("beta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnnormalizationgammaandbetastate/2953934-gamma?language=objc
func (c_ CNNNormalizationGammaAndBetaState) Gamma() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("gamma"))
	return rv
}