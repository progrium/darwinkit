// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNGradientState] class.
var NNGradientStateClass = _NNGradientStateClass{objc.GetClass("MPSNNGradientState")}

type _NNGradientStateClass struct {
	objc.Class
}

// An interface definition for the [NNGradientState] class.
type INNGradientState interface {
	IState
}

// A class representing the state of a gradient kernel when it was encoded. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngradientstate?language=objc
type NNGradientState struct {
	State
}

func NNGradientStateFrom(ptr unsafe.Pointer) NNGradientState {
	return NNGradientState{
		State: StateFrom(ptr),
	}
}

func (nc _NNGradientStateClass) Alloc() NNGradientState {
	rv := objc.Call[NNGradientState](nc, objc.Sel("alloc"))
	return rv
}

func NNGradientState_Alloc() NNGradientState {
	return NNGradientStateClass.Alloc()
}

func (nc _NNGradientStateClass) New() NNGradientState {
	rv := objc.Call[NNGradientState](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNGradientState() NNGradientState {
	return NNGradientStateClass.New()
}

func (n_ NNGradientState) Init() NNGradientState {
	rv := objc.Call[NNGradientState](n_, objc.Sel("init"))
	return rv
}

func (n_ NNGradientState) InitWithResources(resources []metal.PResource) NNGradientState {
	rv := objc.Call[NNGradientState](n_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewNNGradientStateWithResources(resources []metal.PResource) NNGradientState {
	instance := NNGradientStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (n_ NNGradientState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) NNGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNGradientState](n_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewNNGradientStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) NNGradientState {
	instance := NNGradientStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (n_ NNGradientState) InitWithResource(resource metal.PResource) NNGradientState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[NNGradientState](n_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewNNGradientStateWithResource(resource metal.PResource) NNGradientState {
	instance := NNGradientStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (nc _NNGradientStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) NNGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[NNGradientState](nc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func NNGradientState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) NNGradientState {
	return NNGradientStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}
