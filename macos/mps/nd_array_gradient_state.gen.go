// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NDArrayGradientState] class.
var NDArrayGradientStateClass = _NDArrayGradientStateClass{objc.GetClass("MPSNDArrayGradientState")}

type _NDArrayGradientStateClass struct {
	objc.Class
}

// An interface definition for the [NDArrayGradientState] class.
type INDArrayGradientState interface {
	IState
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraygradientstate?language=objc
type NDArrayGradientState struct {
	State
}

func NDArrayGradientStateFrom(ptr unsafe.Pointer) NDArrayGradientState {
	return NDArrayGradientState{
		State: StateFrom(ptr),
	}
}

func (nc _NDArrayGradientStateClass) Alloc() NDArrayGradientState {
	rv := objc.Call[NDArrayGradientState](nc, objc.Sel("alloc"))
	return rv
}

func NDArrayGradientState_Alloc() NDArrayGradientState {
	return NDArrayGradientStateClass.Alloc()
}

func (nc _NDArrayGradientStateClass) New() NDArrayGradientState {
	rv := objc.Call[NDArrayGradientState](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNDArrayGradientState() NDArrayGradientState {
	return NDArrayGradientStateClass.New()
}

func (n_ NDArrayGradientState) Init() NDArrayGradientState {
	rv := objc.Call[NDArrayGradientState](n_, objc.Sel("init"))
	return rv
}

func (n_ NDArrayGradientState) InitWithResources(resources []metal.PResource) NDArrayGradientState {
	rv := objc.Call[NDArrayGradientState](n_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewNDArrayGradientStateWithResources(resources []metal.PResource) NDArrayGradientState {
	instance := NDArrayGradientStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayGradientState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) NDArrayGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayGradientState](n_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewNDArrayGradientStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) NDArrayGradientState {
	instance := NDArrayGradientStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayGradientState) InitWithResource(resource metal.PResource) NDArrayGradientState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[NDArrayGradientState](n_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewNDArrayGradientStateWithResource(resource metal.PResource) NDArrayGradientState {
	instance := NDArrayGradientStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (nc _NDArrayGradientStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) NDArrayGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[NDArrayGradientState](nc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func NDArrayGradientState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) NDArrayGradientState {
	return NDArrayGradientStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}
