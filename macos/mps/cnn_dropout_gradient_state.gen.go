// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNDropoutGradientState] class.
var CNNDropoutGradientStateClass = _CNNDropoutGradientStateClass{objc.GetClass("MPSCNNDropoutGradientState")}

type _CNNDropoutGradientStateClass struct {
	objc.Class
}

// An interface definition for the [CNNDropoutGradientState] class.
type ICNNDropoutGradientState interface {
	INNGradientState
	MaskData() []byte
}

// A class that stores the mask used by dropout and gradient dropout filters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientstate?language=objc
type CNNDropoutGradientState struct {
	NNGradientState
}

func CNNDropoutGradientStateFrom(ptr unsafe.Pointer) CNNDropoutGradientState {
	return CNNDropoutGradientState{
		NNGradientState: NNGradientStateFrom(ptr),
	}
}

func (cc _CNNDropoutGradientStateClass) Alloc() CNNDropoutGradientState {
	rv := objc.Call[CNNDropoutGradientState](cc, objc.Sel("alloc"))
	return rv
}

func CNNDropoutGradientState_Alloc() CNNDropoutGradientState {
	return CNNDropoutGradientStateClass.Alloc()
}

func (cc _CNNDropoutGradientStateClass) New() CNNDropoutGradientState {
	rv := objc.Call[CNNDropoutGradientState](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNDropoutGradientState() CNNDropoutGradientState {
	return CNNDropoutGradientStateClass.New()
}

func (c_ CNNDropoutGradientState) Init() CNNDropoutGradientState {
	rv := objc.Call[CNNDropoutGradientState](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNDropoutGradientState) InitWithResources(resources []metal.PResource) CNNDropoutGradientState {
	rv := objc.Call[CNNDropoutGradientState](c_, objc.Sel("initWithResources:"), resources)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947895-initwithresources?language=objc
func NewCNNDropoutGradientStateWithResources(resources []metal.PResource) CNNDropoutGradientState {
	instance := CNNDropoutGradientStateClass.Alloc().InitWithResources(resources)
	instance.Autorelease()
	return instance
}

func (c_ CNNDropoutGradientState) InitWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNDropoutGradientState {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDropoutGradientState](c_, objc.Sel("initWithDevice:bufferSize:"), po0, bufferSize)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942392-initwithdevice?language=objc
func NewCNNDropoutGradientStateWithDeviceBufferSize(device metal.PDevice, bufferSize uint) CNNDropoutGradientState {
	instance := CNNDropoutGradientStateClass.Alloc().InitWithDeviceBufferSize(device, bufferSize)
	instance.Autorelease()
	return instance
}

func (c_ CNNDropoutGradientState) InitWithResource(resource metal.PResource) CNNDropoutGradientState {
	po0 := objc.WrapAsProtocol("MTLResource", resource)
	rv := objc.Call[CNNDropoutGradientState](c_, objc.Sel("initWithResource:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2942390-initwithresource?language=objc
func NewCNNDropoutGradientStateWithResource(resource metal.PResource) CNNDropoutGradientState {
	instance := CNNDropoutGradientStateClass.Alloc().InitWithResource(resource)
	instance.Autorelease()
	return instance
}

func (cc _CNNDropoutGradientStateClass) TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNDropoutGradientState {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[CNNDropoutGradientState](cc, objc.Sel("temporaryStateWithCommandBuffer:resourceList:"), po0, objc.Ptr(resourceList))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsstate/2947915-temporarystatewithcommandbuffer?language=objc
func CNNDropoutGradientState_TemporaryStateWithCommandBufferResourceList(commandBuffer metal.PCommandBuffer, resourceList IStateResourceList) CNNDropoutGradientState {
	return CNNDropoutGradientStateClass.TemporaryStateWithCommandBufferResourceList(commandBuffer, resourceList)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropoutgradientstate/2942527-maskdata?language=objc
func (c_ CNNDropoutGradientState) MaskData() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("maskData"))
	return rv
}
