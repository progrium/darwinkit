// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNInstanceNormalizationGradientState] class.
var CNNInstanceNormalizationGradientStateClass = _CNNInstanceNormalizationGradientStateClass{objc.GetClass("MPSCNNInstanceNormalizationGradientState")}

type _CNNInstanceNormalizationGradientStateClass struct {
	objc.Class
}

// An interface definition for the [CNNInstanceNormalizationGradientState] class.
type ICNNInstanceNormalizationGradientState interface {
	INNGradientState
	InstanceNormalization() CNNInstanceNormalization
	GradientForGamma() metal.BufferWrapper
	Beta() metal.BufferWrapper
	Gamma() metal.BufferWrapper
	GradientForBeta() metal.BufferWrapper
}

// An object that stores information required to execute a gradient pass for instance normalization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate?language=objc
type CNNInstanceNormalizationGradientState struct {
	NNGradientState
}

func CNNInstanceNormalizationGradientStateFrom(ptr unsafe.Pointer) CNNInstanceNormalizationGradientState {
	return CNNInstanceNormalizationGradientState{
		NNGradientState: NNGradientStateFrom(ptr),
	}
}

func (cc _CNNInstanceNormalizationGradientStateClass) Alloc() CNNInstanceNormalizationGradientState {
	rv := objc.Call[CNNInstanceNormalizationGradientState](cc, objc.Sel("alloc"))
	return rv
}

func CNNInstanceNormalizationGradientState_Alloc() CNNInstanceNormalizationGradientState {
	return CNNInstanceNormalizationGradientStateClass.Alloc()
}

func (cc _CNNInstanceNormalizationGradientStateClass) New() CNNInstanceNormalizationGradientState {
	rv := objc.Call[CNNInstanceNormalizationGradientState](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNInstanceNormalizationGradientState() CNNInstanceNormalizationGradientState {
	return CNNInstanceNormalizationGradientStateClass.New()
}

func (c_ CNNInstanceNormalizationGradientState) Init() CNNInstanceNormalizationGradientState {
	rv := objc.Call[CNNInstanceNormalizationGradientState](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNInstanceNormalizationGradientState) InitWithResources(resources []metal.PResource) CNNInstanceNormalizationGradientState {
	rv := objc.Call[CNNInstanceNormalizationGradientState](c_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewCNNInstanceNormalizationGradientStateWithResources(resources []metal.PResource) CNNInstanceNormalizationGradientState {
	instance := CNNInstanceNormalizationGradientStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (c_ CNNInstanceNormalizationGradientState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNInstanceNormalizationGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNInstanceNormalizationGradientState](c_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewCNNInstanceNormalizationGradientStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNInstanceNormalizationGradientState {
	instance := CNNInstanceNormalizationGradientStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (c_ CNNInstanceNormalizationGradientState) InitWithResource(resource metal.PResource) CNNInstanceNormalizationGradientState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[CNNInstanceNormalizationGradientState](c_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewCNNInstanceNormalizationGradientStateWithResource(resource metal.PResource) CNNInstanceNormalizationGradientState {
	instance := CNNInstanceNormalizationGradientStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (cc _CNNInstanceNormalizationGradientStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNInstanceNormalizationGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNInstanceNormalizationGradientState](cc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func CNNInstanceNormalizationGradientState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNInstanceNormalizationGradientState {
	return CNNInstanceNormalizationGradientStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2953924-instancenormalization?language=objc
func (c_ CNNInstanceNormalizationGradientState) InstanceNormalization() CNNInstanceNormalization {
	rv := objc.Call[CNNInstanceNormalization](c_, objc.Sel("instanceNormalization"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2953928-gradientforgamma?language=objc
func (c_ CNNInstanceNormalizationGradientState) GradientForGamma() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("gradientForGamma"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2956162-beta?language=objc
func (c_ CNNInstanceNormalizationGradientState) Beta() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("beta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2956161-gamma?language=objc
func (c_ CNNInstanceNormalizationGradientState) Gamma() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("gamma"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnninstancenormalizationgradientstate/2953930-gradientforbeta?language=objc
func (c_ CNNInstanceNormalizationGradientState) GradientForBeta() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("gradientForBeta"))
	return rv
}
