// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNBatchNormalizationState] class.
var CNNBatchNormalizationStateClass = _CNNBatchNormalizationStateClass{objc.GetClass("MPSCNNBatchNormalizationState")}

type _CNNBatchNormalizationStateClass struct {
	objc.Class
}

// An interface definition for the [CNNBatchNormalizationState] class.
type ICNNBatchNormalizationState interface {
	INNGradientState
	GradientForGamma() metal.BufferWrapper
	Beta() metal.BufferWrapper
	Gamma() metal.BufferWrapper
	GradientForBeta() metal.BufferWrapper
	Variance() metal.BufferWrapper
	Reset()
	Mean() metal.BufferWrapper
	BatchNormalization() CNNBatchNormalization
}

// An object that stores data required to execute batch normalization. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate?language=objc
type CNNBatchNormalizationState struct {
	NNGradientState
}

func CNNBatchNormalizationStateFrom(ptr unsafe.Pointer) CNNBatchNormalizationState {
	return CNNBatchNormalizationState{
		NNGradientState: NNGradientStateFrom(ptr),
	}
}

func (cc _CNNBatchNormalizationStateClass) Alloc() CNNBatchNormalizationState {
	rv := objc.Call[CNNBatchNormalizationState](cc, objc.Sel("alloc"))
	return rv
}

func CNNBatchNormalizationState_Alloc() CNNBatchNormalizationState {
	return CNNBatchNormalizationStateClass.Alloc()
}

func (cc _CNNBatchNormalizationStateClass) New() CNNBatchNormalizationState {
	rv := objc.Call[CNNBatchNormalizationState](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNBatchNormalizationState() CNNBatchNormalizationState {
	return CNNBatchNormalizationStateClass.New()
}

func (c_ CNNBatchNormalizationState) Init() CNNBatchNormalizationState {
	rv := objc.Call[CNNBatchNormalizationState](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNBatchNormalizationState) InitWithResources(resources []metal.PResource) CNNBatchNormalizationState {
	rv := objc.Call[CNNBatchNormalizationState](c_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewCNNBatchNormalizationStateWithResources(resources []metal.PResource) CNNBatchNormalizationState {
	instance := CNNBatchNormalizationStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (c_ CNNBatchNormalizationState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNBatchNormalizationState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBatchNormalizationState](c_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewCNNBatchNormalizationStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNBatchNormalizationState {
	instance := CNNBatchNormalizationStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (c_ CNNBatchNormalizationState) InitWithResource(resource metal.PResource) CNNBatchNormalizationState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[CNNBatchNormalizationState](c_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewCNNBatchNormalizationStateWithResource(resource metal.PResource) CNNBatchNormalizationState {
	instance := CNNBatchNormalizationStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (cc _CNNBatchNormalizationStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNBatchNormalizationState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNBatchNormalizationState](cc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func CNNBatchNormalizationState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNBatchNormalizationState {
	return CNNBatchNormalizationStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2951893-gradientforgamma?language=objc
func (c_ CNNBatchNormalizationState) GradientForGamma() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("gradientForGamma"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2951888-beta?language=objc
func (c_ CNNBatchNormalizationState) Beta() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("beta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2951892-gamma?language=objc
func (c_ CNNBatchNormalizationState) Gamma() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("gamma"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2951890-gradientforbeta?language=objc
func (c_ CNNBatchNormalizationState) GradientForBeta() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("gradientForBeta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2942603-variance?language=objc
func (c_ CNNBatchNormalizationState) Variance() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("variance"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2942587-reset?language=objc
func (c_ CNNBatchNormalizationState) Reset() {
	objc.Call[objc.Void](c_, objc.Sel("reset"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2942612-mean?language=objc
func (c_ CNNBatchNormalizationState) Mean() metal.BufferWrapper {
	rv := objc.Call[metal.BufferWrapper](c_, objc.Sel("mean"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbatchnormalizationstate/2953969-batchnormalization?language=objc
func (c_ CNNBatchNormalizationState) BatchNormalization() CNNBatchNormalization {
	rv := objc.Call[CNNBatchNormalization](c_, objc.Sel("batchNormalization"))
	return rv
}
