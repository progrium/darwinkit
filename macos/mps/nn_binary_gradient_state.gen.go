// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNBinaryGradientState] class.
var NNBinaryGradientStateClass = _NNBinaryGradientStateClass{objc.GetClass("MPSNNBinaryGradientState")}

type _NNBinaryGradientStateClass struct {
	objc.Class
}

// An interface definition for the [NNBinaryGradientState] class.
type INNBinaryGradientState interface {
	IState
}

// A class representing the state of a gradient binary kernel when it was encoded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnbinarygradientstate?language=objc
type NNBinaryGradientState struct {
	State
}

func NNBinaryGradientStateFrom(ptr unsafe.Pointer) NNBinaryGradientState {
	return NNBinaryGradientState{
		State: StateFrom(ptr),
	}
}

func (nc _NNBinaryGradientStateClass) Alloc() NNBinaryGradientState {
	rv := objc.Call[NNBinaryGradientState](nc, objc.Sel("alloc"))
	return rv
}

func NNBinaryGradientState_Alloc() NNBinaryGradientState {
	return NNBinaryGradientStateClass.Alloc()
}

func (nc _NNBinaryGradientStateClass) New() NNBinaryGradientState {
	rv := objc.Call[NNBinaryGradientState](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNBinaryGradientState() NNBinaryGradientState {
	return NNBinaryGradientStateClass.New()
}

func (n_ NNBinaryGradientState) Init() NNBinaryGradientState {
	rv := objc.Call[NNBinaryGradientState](n_, objc.Sel("init"))
	return rv
}

func (n_ NNBinaryGradientState) InitWithResources(resources []metal.PResource) NNBinaryGradientState {
	rv := objc.Call[NNBinaryGradientState](n_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewNNBinaryGradientStateWithResources(resources []metal.PResource) NNBinaryGradientState {
	instance := NNBinaryGradientStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (n_ NNBinaryGradientState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) NNBinaryGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNBinaryGradientState](n_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewNNBinaryGradientStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) NNBinaryGradientState {
	instance := NNBinaryGradientStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (n_ NNBinaryGradientState) InitWithResource(resource metal.PResource) NNBinaryGradientState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[NNBinaryGradientState](n_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewNNBinaryGradientStateWithResource(resource metal.PResource) NNBinaryGradientState {
	instance := NNBinaryGradientStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (nc _NNBinaryGradientStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) NNBinaryGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[NNBinaryGradientState](nc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func NNBinaryGradientState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) NNBinaryGradientState {
	return NNBinaryGradientStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}
