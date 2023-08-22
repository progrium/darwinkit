// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RNNMatrixTrainingState] class.
var RNNMatrixTrainingStateClass = _RNNMatrixTrainingStateClass{objc.GetClass("MPSRNNMatrixTrainingState")}

type _RNNMatrixTrainingStateClass struct {
	objc.Class
}

// An interface definition for the [RNNMatrixTrainingState] class.
type IRNNMatrixTrainingState interface {
	IState
}

// A class that holds data from a forward pass to be used in a backward pass. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnmatrixtrainingstate?language=objc
type RNNMatrixTrainingState struct {
	State
}

func RNNMatrixTrainingStateFrom(ptr unsafe.Pointer) RNNMatrixTrainingState {
	return RNNMatrixTrainingState{
		State: StateFrom(ptr),
	}
}

func (rc _RNNMatrixTrainingStateClass) Alloc() RNNMatrixTrainingState {
	rv := objc.Call[RNNMatrixTrainingState](rc, objc.Sel("alloc"))
	return rv
}

func RNNMatrixTrainingState_Alloc() RNNMatrixTrainingState {
	return RNNMatrixTrainingStateClass.Alloc()
}

func (rc _RNNMatrixTrainingStateClass) New() RNNMatrixTrainingState {
	rv := objc.Call[RNNMatrixTrainingState](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRNNMatrixTrainingState() RNNMatrixTrainingState {
	return RNNMatrixTrainingStateClass.New()
}

func (r_ RNNMatrixTrainingState) Init() RNNMatrixTrainingState {
	rv := objc.Call[RNNMatrixTrainingState](r_, objc.Sel("init"))
	return rv
}

func (r_ RNNMatrixTrainingState) InitWithResources(resources []metal.PResource) RNNMatrixTrainingState {
	rv := objc.Call[RNNMatrixTrainingState](r_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewRNNMatrixTrainingStateWithResources(resources []metal.PResource) RNNMatrixTrainingState {
	instance := RNNMatrixTrainingStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (r_ RNNMatrixTrainingState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) RNNMatrixTrainingState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[RNNMatrixTrainingState](r_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewRNNMatrixTrainingStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) RNNMatrixTrainingState {
	instance := RNNMatrixTrainingStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (r_ RNNMatrixTrainingState) InitWithResource(resource metal.PResource) RNNMatrixTrainingState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[RNNMatrixTrainingState](r_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewRNNMatrixTrainingStateWithResource(resource metal.PResource) RNNMatrixTrainingState {
	instance := RNNMatrixTrainingStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (rc _RNNMatrixTrainingStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) RNNMatrixTrainingState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[RNNMatrixTrainingState](rc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func RNNMatrixTrainingState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) RNNMatrixTrainingState {
	return RNNMatrixTrainingStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}
