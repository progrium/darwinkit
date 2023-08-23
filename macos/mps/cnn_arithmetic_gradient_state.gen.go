// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNArithmeticGradientState] class.
var CNNArithmeticGradientStateClass = _CNNArithmeticGradientStateClass{objc.GetClass("MPSCNNArithmeticGradientState")}

type _CNNArithmeticGradientStateClass struct {
	objc.Class
}

// An interface definition for the [CNNArithmeticGradientState] class.
type ICNNArithmeticGradientState interface {
	INNBinaryGradientState
}

// An object that stores the clamp mask used by gradient arithmetic operators. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmeticgradientstate?language=objc
type CNNArithmeticGradientState struct {
	NNBinaryGradientState
}

func CNNArithmeticGradientStateFrom(ptr unsafe.Pointer) CNNArithmeticGradientState {
	return CNNArithmeticGradientState{
		NNBinaryGradientState: NNBinaryGradientStateFrom(ptr),
	}
}

func (cc _CNNArithmeticGradientStateClass) Alloc() CNNArithmeticGradientState {
	rv := objc.Call[CNNArithmeticGradientState](cc, objc.Sel("alloc"))
	return rv
}

func CNNArithmeticGradientState_Alloc() CNNArithmeticGradientState {
	return CNNArithmeticGradientStateClass.Alloc()
}

func (cc _CNNArithmeticGradientStateClass) New() CNNArithmeticGradientState {
	rv := objc.Call[CNNArithmeticGradientState](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNArithmeticGradientState() CNNArithmeticGradientState {
	return CNNArithmeticGradientStateClass.New()
}

func (c_ CNNArithmeticGradientState) Init() CNNArithmeticGradientState {
	rv := objc.Call[CNNArithmeticGradientState](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNArithmeticGradientState) InitWithResources(resources []metal.PResource) CNNArithmeticGradientState {
	rv := objc.Call[CNNArithmeticGradientState](c_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewCNNArithmeticGradientStateWithResources(resources []metal.PResource) CNNArithmeticGradientState {
	instance := CNNArithmeticGradientStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (c_ CNNArithmeticGradientState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNArithmeticGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNArithmeticGradientState](c_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewCNNArithmeticGradientStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNArithmeticGradientState {
	instance := CNNArithmeticGradientStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (c_ CNNArithmeticGradientState) InitWithResource(resource metal.PResource) CNNArithmeticGradientState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[CNNArithmeticGradientState](c_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewCNNArithmeticGradientStateWithResource(resource metal.PResource) CNNArithmeticGradientState {
	instance := CNNArithmeticGradientStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (cc _CNNArithmeticGradientStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNArithmeticGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNArithmeticGradientState](cc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func CNNArithmeticGradientState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNArithmeticGradientState {
	return CNNArithmeticGradientStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}
