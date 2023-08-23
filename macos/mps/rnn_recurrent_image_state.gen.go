// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RNNRecurrentImageState] class.
var RNNRecurrentImageStateClass = _RNNRecurrentImageStateClass{objc.GetClass("MPSRNNRecurrentImageState")}

type _RNNRecurrentImageStateClass struct {
	objc.Class
}

// An interface definition for the [RNNRecurrentImageState] class.
type IRNNRecurrentImageState interface {
	IState
	GetRecurrentOutputImageForLayerIndex(layerIndex uint) Image
	GetMemoryCellImageForLayerIndex(layerIndex uint) Image
}

// A class that holds all the data that's passed from one sequence iteration of the image-based recurrent neural network layer (stack) to the next. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnrecurrentimagestate?language=objc
type RNNRecurrentImageState struct {
	State
}

func RNNRecurrentImageStateFrom(ptr unsafe.Pointer) RNNRecurrentImageState {
	return RNNRecurrentImageState{
		State: StateFrom(ptr),
	}
}

func (rc _RNNRecurrentImageStateClass) Alloc() RNNRecurrentImageState {
	rv := objc.Call[RNNRecurrentImageState](rc, objc.Sel("alloc"))
	return rv
}

func RNNRecurrentImageState_Alloc() RNNRecurrentImageState {
	return RNNRecurrentImageStateClass.Alloc()
}

func (rc _RNNRecurrentImageStateClass) New() RNNRecurrentImageState {
	rv := objc.Call[RNNRecurrentImageState](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRNNRecurrentImageState() RNNRecurrentImageState {
	return RNNRecurrentImageStateClass.New()
}

func (r_ RNNRecurrentImageState) Init() RNNRecurrentImageState {
	rv := objc.Call[RNNRecurrentImageState](r_, objc.Sel("init"))
	return rv
}

func (r_ RNNRecurrentImageState) InitWithResources(resources []metal.PResource) RNNRecurrentImageState {
	rv := objc.Call[RNNRecurrentImageState](r_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewRNNRecurrentImageStateWithResources(resources []metal.PResource) RNNRecurrentImageState {
	instance := RNNRecurrentImageStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (r_ RNNRecurrentImageState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) RNNRecurrentImageState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[RNNRecurrentImageState](r_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewRNNRecurrentImageStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) RNNRecurrentImageState {
	instance := RNNRecurrentImageStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (r_ RNNRecurrentImageState) InitWithResource(resource metal.PResource) RNNRecurrentImageState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[RNNRecurrentImageState](r_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewRNNRecurrentImageStateWithResource(resource metal.PResource) RNNRecurrentImageState {
	instance := RNNRecurrentImageStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (rc _RNNRecurrentImageStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) RNNRecurrentImageState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[RNNRecurrentImageState](rc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func RNNRecurrentImageState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) RNNRecurrentImageState {
	return RNNRecurrentImageStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnrecurrentimagestate/2865742-getrecurrentoutputimageforlayeri?language=objc
func (r_ RNNRecurrentImageState) GetRecurrentOutputImageForLayerIndex(layerIndex uint) Image {
	rv := objc.Call[Image](r_, objc.Sel("getRecurrentOutputImageForLayerIndex:"), layerIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnrecurrentimagestate/2865740-getmemorycellimageforlayerindex?language=objc
func (r_ RNNRecurrentImageState) GetMemoryCellImageForLayerIndex(layerIndex uint) Image {
	rv := objc.Call[Image](r_, objc.Sel("getMemoryCellImageForLayerIndex:"), layerIndex)
	return rv
}
