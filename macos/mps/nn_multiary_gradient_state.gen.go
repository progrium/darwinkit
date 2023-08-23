// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNMultiaryGradientState] class.
var NNMultiaryGradientStateClass = _NNMultiaryGradientStateClass{objc.GetClass("MPSNNMultiaryGradientState")}

type _NNMultiaryGradientStateClass struct {
	objc.Class
}

// An interface definition for the [NNMultiaryGradientState] class.
type INNMultiaryGradientState interface {
	IState
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnmultiarygradientstate?language=objc
type NNMultiaryGradientState struct {
	State
}

func NNMultiaryGradientStateFrom(ptr unsafe.Pointer) NNMultiaryGradientState {
	return NNMultiaryGradientState{
		State: StateFrom(ptr),
	}
}

func (nc _NNMultiaryGradientStateClass) Alloc() NNMultiaryGradientState {
	rv := objc.Call[NNMultiaryGradientState](nc, objc.Sel("alloc"))
	return rv
}

func NNMultiaryGradientState_Alloc() NNMultiaryGradientState {
	return NNMultiaryGradientStateClass.Alloc()
}

func (nc _NNMultiaryGradientStateClass) New() NNMultiaryGradientState {
	rv := objc.Call[NNMultiaryGradientState](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNMultiaryGradientState() NNMultiaryGradientState {
	return NNMultiaryGradientStateClass.New()
}

func (n_ NNMultiaryGradientState) Init() NNMultiaryGradientState {
	rv := objc.Call[NNMultiaryGradientState](n_, objc.Sel("init"))
	return rv
}

func (n_ NNMultiaryGradientState) InitWithResources(resources []metal.PResource) NNMultiaryGradientState {
	rv := objc.Call[NNMultiaryGradientState](n_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewNNMultiaryGradientStateWithResources(resources []metal.PResource) NNMultiaryGradientState {
	instance := NNMultiaryGradientStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (n_ NNMultiaryGradientState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) NNMultiaryGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNMultiaryGradientState](n_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewNNMultiaryGradientStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) NNMultiaryGradientState {
	instance := NNMultiaryGradientStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (n_ NNMultiaryGradientState) InitWithResource(resource metal.PResource) NNMultiaryGradientState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[NNMultiaryGradientState](n_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewNNMultiaryGradientStateWithResource(resource metal.PResource) NNMultiaryGradientState {
	instance := NNMultiaryGradientStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (nc _NNMultiaryGradientStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) NNMultiaryGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[NNMultiaryGradientState](nc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func NNMultiaryGradientState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) NNMultiaryGradientState {
	return NNMultiaryGradientStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}
