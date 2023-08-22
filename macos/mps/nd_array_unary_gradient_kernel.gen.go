// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NDArrayUnaryGradientKernel] class.
var NDArrayUnaryGradientKernelClass = _NDArrayUnaryGradientKernelClass{objc.GetClass("MPSNDArrayUnaryGradientKernel")}

type _NDArrayUnaryGradientKernelClass struct {
	objc.Class
}

// An interface definition for the [NDArrayUnaryGradientKernel] class.
type INDArrayUnaryGradientKernel interface {
	INDArrayMultiaryGradientKernel
	EncodeToCommandBufferSourceArraySourceGradientGradientState(cmdBuf metal.PCommandBuffer, sourceArray INDArray, gradient INDArray, state IState) NDArray
	EncodeToCommandBufferObjectSourceArraySourceGradientGradientState(cmdBufObject objc.IObject, sourceArray INDArray, gradient INDArray, state IState) NDArray
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarygradientkernel?language=objc
type NDArrayUnaryGradientKernel struct {
	NDArrayMultiaryGradientKernel
}

func NDArrayUnaryGradientKernelFrom(ptr unsafe.Pointer) NDArrayUnaryGradientKernel {
	return NDArrayUnaryGradientKernel{
		NDArrayMultiaryGradientKernel: NDArrayMultiaryGradientKernelFrom(ptr),
	}
}

func (n_ NDArrayUnaryGradientKernel) InitWithDevice(device metal.PDevice) NDArrayUnaryGradientKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayUnaryGradientKernel](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarygradientkernel/3143532-initwithdevice?language=objc
func NewNDArrayUnaryGradientKernelWithDevice(device metal.PDevice) NDArrayUnaryGradientKernel {
	instance := NDArrayUnaryGradientKernelClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NDArrayUnaryGradientKernelClass) Alloc() NDArrayUnaryGradientKernel {
	rv := objc.Call[NDArrayUnaryGradientKernel](nc, objc.Sel("alloc"))
	return rv
}

func NDArrayUnaryGradientKernel_Alloc() NDArrayUnaryGradientKernel {
	return NDArrayUnaryGradientKernelClass.Alloc()
}

func (nc _NDArrayUnaryGradientKernelClass) New() NDArrayUnaryGradientKernel {
	rv := objc.Call[NDArrayUnaryGradientKernel](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNDArrayUnaryGradientKernel() NDArrayUnaryGradientKernel {
	return NDArrayUnaryGradientKernelClass.New()
}

func (n_ NDArrayUnaryGradientKernel) Init() NDArrayUnaryGradientKernel {
	rv := objc.Call[NDArrayUnaryGradientKernel](n_, objc.Sel("init"))
	return rv
}

func (n_ NDArrayUnaryGradientKernel) InitWithDeviceSourceCountSourceGradientIndex(device metal.PDevice, count uint, sourceGradientIndex uint) NDArrayUnaryGradientKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayUnaryGradientKernel](n_, objc.Sel("initWithDevice:sourceCount:sourceGradientIndex:"), po0, count, sourceGradientIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel/3143524-initwithdevice?language=objc
func NewNDArrayUnaryGradientKernelWithDeviceSourceCountSourceGradientIndex(device metal.PDevice, count uint, sourceGradientIndex uint) NDArrayUnaryGradientKernel {
	instance := NDArrayUnaryGradientKernelClass.Alloc().InitWithDeviceSourceCountSourceGradientIndex(device, count, sourceGradientIndex)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayUnaryGradientKernel) InitWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayUnaryGradientKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayUnaryGradientKernel](n_, objc.Sel("initWithDevice:sourceCount:"), po0, count)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131741-initwithdevice?language=objc
func NewNDArrayUnaryGradientKernelWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayUnaryGradientKernel {
	instance := NDArrayUnaryGradientKernelClass.Alloc().InitWithDeviceSourceCount(device, count)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayUnaryGradientKernel) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayUnaryGradientKernel {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayUnaryGradientKernel](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131734-copywithzone?language=objc
func NDArrayUnaryGradientKernel_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayUnaryGradientKernel {
	instance := NDArrayUnaryGradientKernelClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarygradientkernel/3143530-encodetocommandbuffer?language=objc
func (n_ NDArrayUnaryGradientKernel) EncodeToCommandBufferSourceArraySourceGradientGradientState(cmdBuf metal.PCommandBuffer, sourceArray INDArray, gradient INDArray, state IState) NDArray {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[NDArray](n_, objc.Sel("encodeToCommandBuffer:sourceArray:sourceGradient:gradientState:"), po0, objc.Ptr(sourceArray), objc.Ptr(gradient), objc.Ptr(state))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarygradientkernel/3143530-encodetocommandbuffer?language=objc
func (n_ NDArrayUnaryGradientKernel) EncodeToCommandBufferObjectSourceArraySourceGradientGradientState(cmdBufObject objc.IObject, sourceArray INDArray, gradient INDArray, state IState) NDArray {
	rv := objc.Call[NDArray](n_, objc.Sel("encodeToCommandBuffer:sourceArray:sourceGradient:gradientState:"), objc.Ptr(cmdBufObject), objc.Ptr(sourceArray), objc.Ptr(gradient), objc.Ptr(state))
	return rv
}
