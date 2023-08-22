// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNGroupNormalizationGradientState] class.
var CNNGroupNormalizationGradientStateClass = _CNNGroupNormalizationGradientStateClass{objc.GetClass("MPSCNNGroupNormalizationGradientState")}

type _CNNGroupNormalizationGradientStateClass struct {
	objc.Class
}

// An interface definition for the [CNNGroupNormalizationGradientState] class.
type ICNNGroupNormalizationGradientState interface {
	INNGradientState
	GroupNormalization() CNNGroupNormalization
	GradientForGamma() metal.BufferWrapper
	Beta() metal.BufferWrapper
	Gamma() metal.BufferWrapper
	GradientForBeta() metal.BufferWrapper
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate?language=objc
type CNNGroupNormalizationGradientState struct {
	NNGradientState
}

func CNNGroupNormalizationGradientStateFrom(ptr unsafe.Pointer) CNNGroupNormalizationGradientState {
	return CNNGroupNormalizationGradientState{
		NNGradientState: NNGradientStateFrom(ptr),
	}
}

func (cc _CNNGroupNormalizationGradientStateClass) Alloc() CNNGroupNormalizationGradientState {
	rv := objc.Call[CNNGroupNormalizationGradientState](cc, objc.Sel("alloc"))
	return rv
}

func CNNGroupNormalizationGradientState_Alloc() CNNGroupNormalizationGradientState {
	return CNNGroupNormalizationGradientStateClass.Alloc()
}

func (cc _CNNGroupNormalizationGradientStateClass) New() CNNGroupNormalizationGradientState {
	rv := objc.Call[CNNGroupNormalizationGradientState](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNGroupNormalizationGradientState() CNNGroupNormalizationGradientState {
	return CNNGroupNormalizationGradientStateClass.New()
}

func (c_ CNNGroupNormalizationGradientState) Init() CNNGroupNormalizationGradientState {
	rv := objc.Call[CNNGroupNormalizationGradientState](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNGroupNormalizationGradientState) InitWithResources(resources []metal.PResource) CNNGroupNormalizationGradientState {
	rv := objc.Call[CNNGroupNormalizationGradientState](c_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewCNNGroupNormalizationGradientStateWithResources(resources []metal.PResource) CNNGroupNormalizationGradientState {
	instance := CNNGroupNormalizationGradientStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (c_ CNNGroupNormalizationGradientState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNGroupNormalizationGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNGroupNormalizationGradientState](c_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewCNNGroupNormalizationGradientStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNGroupNormalizationGradientState {
	instance := CNNGroupNormalizationGradientStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (c_ CNNGroupNormalizationGradientState) InitWithResource(resource metal.PResource) CNNGroupNormalizationGradientState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[CNNGroupNormalizationGradientState](c_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewCNNGroupNormalizationGradientStateWithResource(resource metal.PResource) CNNGroupNormalizationGradientState {
	instance := CNNGroupNormalizationGradientStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (cc _CNNGroupNormalizationGradientStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNGroupNormalizationGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNGroupNormalizationGradientState](cc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func CNNGroupNormalizationGradientState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNGroupNormalizationGradientState {
	return CNNGroupNormalizationGradientStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152561-groupnormalization?language=objc
func (c_ CNNGroupNormalizationGradientState) GroupNormalization() CNNGroupNormalization {
	rv := objc.Call[CNNGroupNormalization](c_, objc.Sel("groupNormalization"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152560-gradientforgamma?language=objc
func (c_ CNNGroupNormalizationGradientState) GradientForGamma() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("gradientForGamma"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152557-beta?language=objc
func (c_ CNNGroupNormalizationGradientState) Beta() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("beta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152558-gamma?language=objc
func (c_ CNNGroupNormalizationGradientState) Gamma() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("gamma"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngroupnormalizationgradientstate/3152559-gradientforbeta?language=objc
func (c_ CNNGroupNormalizationGradientState) GradientForBeta() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("gradientForBeta"))
	return rv
}