// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NDArrayMultiaryGradientKernel] class.
var NDArrayMultiaryGradientKernelClass = _NDArrayMultiaryGradientKernelClass{objc.GetClass("MPSNDArrayMultiaryGradientKernel")}

type _NDArrayMultiaryGradientKernelClass struct {
	objc.Class
}

// An interface definition for the [NDArrayMultiaryGradientKernel] class.
type INDArrayMultiaryGradientKernel interface {
	INDArrayMultiaryBase
	EncodeToCommandBufferSourceArraysSourceGradientGradientState(cmdBuf metal.PCommandBuffer, sources []INDArray, gradient INDArray, state IState) NDArray
	EncodeToCommandBufferObjectSourceArraysSourceGradientGradientState(cmdBufObject objc.IObject, sources []INDArray, gradient INDArray, state IState) NDArray
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel?language=objc
type NDArrayMultiaryGradientKernel struct {
	NDArrayMultiaryBase
}

func NDArrayMultiaryGradientKernelFrom(ptr unsafe.Pointer) NDArrayMultiaryGradientKernel {
	return NDArrayMultiaryGradientKernel{
		NDArrayMultiaryBase: NDArrayMultiaryBaseFrom(ptr),
	}
}

func (n_ NDArrayMultiaryGradientKernel) InitWithDeviceSourceCountSourceGradientIndex(device metal.PDevice, count uint, sourceGradientIndex uint) NDArrayMultiaryGradientKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayMultiaryGradientKernel](n_, objc.Sel("initWithDevice:sourceCount:sourceGradientIndex:"), po0, count, sourceGradientIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel/3143524-initwithdevice?language=objc
func NewNDArrayMultiaryGradientKernelWithDeviceSourceCountSourceGradientIndex(device metal.PDevice, count uint, sourceGradientIndex uint) NDArrayMultiaryGradientKernel {
	instance := NDArrayMultiaryGradientKernelClass.Alloc().InitWithDeviceSourceCountSourceGradientIndex(device, count, sourceGradientIndex)
	instance.Autorelease()
	return instance
}

func (nc _NDArrayMultiaryGradientKernelClass) Alloc() NDArrayMultiaryGradientKernel {
	rv := objc.Call[NDArrayMultiaryGradientKernel](nc, objc.Sel("alloc"))
	return rv
}

func NDArrayMultiaryGradientKernel_Alloc() NDArrayMultiaryGradientKernel {
	return NDArrayMultiaryGradientKernelClass.Alloc()
}

func (nc _NDArrayMultiaryGradientKernelClass) New() NDArrayMultiaryGradientKernel {
	rv := objc.Call[NDArrayMultiaryGradientKernel](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNDArrayMultiaryGradientKernel() NDArrayMultiaryGradientKernel {
	return NDArrayMultiaryGradientKernelClass.New()
}

func (n_ NDArrayMultiaryGradientKernel) Init() NDArrayMultiaryGradientKernel {
	rv := objc.Call[NDArrayMultiaryGradientKernel](n_, objc.Sel("init"))
	return rv
}

func (n_ NDArrayMultiaryGradientKernel) InitWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayMultiaryGradientKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayMultiaryGradientKernel](n_, objc.Sel("initWithDevice:sourceCount:"), po0, count)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131741-initwithdevice?language=objc
func NewNDArrayMultiaryGradientKernelWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayMultiaryGradientKernel {
	instance := NDArrayMultiaryGradientKernelClass.Alloc().InitWithDeviceSourceCount(device, count)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayMultiaryGradientKernel) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayMultiaryGradientKernel {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayMultiaryGradientKernel](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131734-copywithzone?language=objc
func NDArrayMultiaryGradientKernel_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayMultiaryGradientKernel {
	instance := NDArrayMultiaryGradientKernelClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayMultiaryGradientKernel) InitWithDevice(device metal.PDevice) NDArrayMultiaryGradientKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayMultiaryGradientKernel](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewNDArrayMultiaryGradientKernelWithDevice(device metal.PDevice) NDArrayMultiaryGradientKernel {
	instance := NDArrayMultiaryGradientKernelClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel/3143522-encodetocommandbuffer?language=objc
func (n_ NDArrayMultiaryGradientKernel) EncodeToCommandBufferSourceArraysSourceGradientGradientState(cmdBuf metal.PCommandBuffer, sources []INDArray, gradient INDArray, state IState) NDArray {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[NDArray](n_, objc.Sel("encodeToCommandBuffer:sourceArrays:sourceGradient:gradientState:"), po0, sources, objc.Ptr(gradient), objc.Ptr(state))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel/3143522-encodetocommandbuffer?language=objc
func (n_ NDArrayMultiaryGradientKernel) EncodeToCommandBufferObjectSourceArraysSourceGradientGradientState(cmdBufObject objc.IObject, sources []INDArray, gradient INDArray, state IState) NDArray {
	rv := objc.Call[NDArray](n_, objc.Sel("encodeToCommandBuffer:sourceArrays:sourceGradient:gradientState:"), objc.Ptr(cmdBufObject), sources, objc.Ptr(gradient), objc.Ptr(state))
	return rv
}
