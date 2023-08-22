// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NDArrayMultiaryBase] class.
var NDArrayMultiaryBaseClass = _NDArrayMultiaryBaseClass{objc.GetClass("MPSNDArrayMultiaryBase")}

type _NDArrayMultiaryBaseClass struct {
	objc.Class
}

// An interface definition for the [NDArrayMultiaryBase] class.
type INDArrayMultiaryBase interface {
	IKernel
	EncodeWithCoder(coder foundation.ICoder)
	ResultStateForSourceArraysSourceStatesDestinationArray(sourceArrays []INDArray, sourceStates []IState, destinationArray INDArray) State
	DestinationArrayDescriptorForSourceArraysSourceState(sources []INDArray, state IState) NDArrayDescriptor
	DestinationArrayAllocator() NDArrayAllocatorWrapper
	SetDestinationArrayAllocator(value PNDArrayAllocator)
	SetDestinationArrayAllocatorObject(valueObject objc.IObject)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase?language=objc
type NDArrayMultiaryBase struct {
	Kernel
}

func NDArrayMultiaryBaseFrom(ptr unsafe.Pointer) NDArrayMultiaryBase {
	return NDArrayMultiaryBase{
		Kernel: KernelFrom(ptr),
	}
}

func (n_ NDArrayMultiaryBase) InitWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayMultiaryBase {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayMultiaryBase](n_, objc.Sel("initWithDevice:sourceCount:"), po0, count)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131741-initwithdevice?language=objc
func NewNDArrayMultiaryBaseWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayMultiaryBase {
	instance := NDArrayMultiaryBaseClass.Alloc().InitWithDeviceSourceCount(device, count)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayMultiaryBase) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayMultiaryBase {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayMultiaryBase](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131734-copywithzone?language=objc
func NDArrayMultiaryBase_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayMultiaryBase {
	instance := NDArrayMultiaryBaseClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (nc _NDArrayMultiaryBaseClass) Alloc() NDArrayMultiaryBase {
	rv := objc.Call[NDArrayMultiaryBase](nc, objc.Sel("alloc"))
	return rv
}

func NDArrayMultiaryBase_Alloc() NDArrayMultiaryBase {
	return NDArrayMultiaryBaseClass.Alloc()
}

func (nc _NDArrayMultiaryBaseClass) New() NDArrayMultiaryBase {
	rv := objc.Call[NDArrayMultiaryBase](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNDArrayMultiaryBase() NDArrayMultiaryBase {
	return NDArrayMultiaryBaseClass.New()
}

func (n_ NDArrayMultiaryBase) Init() NDArrayMultiaryBase {
	rv := objc.Call[NDArrayMultiaryBase](n_, objc.Sel("init"))
	return rv
}

func (n_ NDArrayMultiaryBase) InitWithDevice(device metal.PDevice) NDArrayMultiaryBase {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayMultiaryBase](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewNDArrayMultiaryBaseWithDevice(device metal.PDevice) NDArrayMultiaryBase {
	instance := NDArrayMultiaryBaseClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131739-encodewithcoder?language=objc
func (n_ NDArrayMultiaryBase) EncodeWithCoder(coder foundation.ICoder) {
	objc.Call[objc.Void](n_, objc.Sel("encodeWithCoder:"), objc.Ptr(coder))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3143521-resultstateforsourcearrays?language=objc
func (n_ NDArrayMultiaryBase) ResultStateForSourceArraysSourceStatesDestinationArray(sourceArrays []INDArray, sourceStates []IState, destinationArray INDArray) State {
	rv := objc.Call[State](n_, objc.Sel("resultStateForSourceArrays:sourceStates:destinationArray:"), sourceArrays, sourceStates, objc.Ptr(destinationArray))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131736-destinationarraydescriptorforsou?language=objc
func (n_ NDArrayMultiaryBase) DestinationArrayDescriptorForSourceArraysSourceState(sources []INDArray, state IState) NDArrayDescriptor {
	rv := objc.Call[NDArrayDescriptor](n_, objc.Sel("destinationArrayDescriptorForSourceArrays:sourceState:"), sources, objc.Ptr(state))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131735-destinationarrayallocator?language=objc
func (n_ NDArrayMultiaryBase) DestinationArrayAllocator() NDArrayAllocatorWrapper {
	rv := objc.Call[NDArrayAllocatorWrapper](n_, objc.Sel("destinationArrayAllocator"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131735-destinationarrayallocator?language=objc
func (n_ NDArrayMultiaryBase) SetDestinationArrayAllocator(value PNDArrayAllocator) {
	po0 := objc.WrapAsProtocol("MPSNDArrayAllocator", value)
	objc.Call[objc.Void](n_, objc.Sel("setDestinationArrayAllocator:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131735-destinationarrayallocator?language=objc
func (n_ NDArrayMultiaryBase) SetDestinationArrayAllocatorObject(valueObject objc.IObject) {
	objc.Call[objc.Void](n_, objc.Sel("setDestinationArrayAllocator:"), objc.Ptr(valueObject))
}
