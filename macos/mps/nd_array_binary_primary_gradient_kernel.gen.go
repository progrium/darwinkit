// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NDArrayBinaryPrimaryGradientKernel] class.
var NDArrayBinaryPrimaryGradientKernelClass = _NDArrayBinaryPrimaryGradientKernelClass{objc.GetClass("MPSNDArrayBinaryPrimaryGradientKernel")}

type _NDArrayBinaryPrimaryGradientKernelClass struct {
	objc.Class
}

// An interface definition for the [NDArrayBinaryPrimaryGradientKernel] class.
type INDArrayBinaryPrimaryGradientKernel interface {
	INDArrayMultiaryGradientKernel
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState(cmdBuf metal.PCommandBuffer, primarySourceArray INDArray, secondarySourceArray INDArray, gradient INDArray, state IState) NDArray
	EncodeToCommandBufferObjectPrimarySourceArraySecondarySourceArraySourceGradientGradientState(cmdBufObject objc.IObject, primarySourceArray INDArray, secondarySourceArray INDArray, gradient INDArray, state IState) NDArray
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinaryprimarygradientkernel?language=objc
type NDArrayBinaryPrimaryGradientKernel struct {
	NDArrayMultiaryGradientKernel
}

func NDArrayBinaryPrimaryGradientKernelFrom(ptr unsafe.Pointer) NDArrayBinaryPrimaryGradientKernel {
	return NDArrayBinaryPrimaryGradientKernel{
		NDArrayMultiaryGradientKernel: NDArrayMultiaryGradientKernelFrom(ptr),
	}
}

func (n_ NDArrayBinaryPrimaryGradientKernel) InitWithDevice(device metal.PDevice) NDArrayBinaryPrimaryGradientKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayBinaryPrimaryGradientKernel](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinaryprimarygradientkernel/3143515-initwithdevice?language=objc
func NewNDArrayBinaryPrimaryGradientKernelWithDevice(device metal.PDevice) NDArrayBinaryPrimaryGradientKernel {
	instance := NDArrayBinaryPrimaryGradientKernelClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NDArrayBinaryPrimaryGradientKernelClass) Alloc() NDArrayBinaryPrimaryGradientKernel {
	rv := objc.Call[NDArrayBinaryPrimaryGradientKernel](nc, objc.Sel("alloc"))
	return rv
}

func NDArrayBinaryPrimaryGradientKernel_Alloc() NDArrayBinaryPrimaryGradientKernel {
	return NDArrayBinaryPrimaryGradientKernelClass.Alloc()
}

func (nc _NDArrayBinaryPrimaryGradientKernelClass) New() NDArrayBinaryPrimaryGradientKernel {
	rv := objc.Call[NDArrayBinaryPrimaryGradientKernel](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNDArrayBinaryPrimaryGradientKernel() NDArrayBinaryPrimaryGradientKernel {
	return NDArrayBinaryPrimaryGradientKernelClass.New()
}

func (n_ NDArrayBinaryPrimaryGradientKernel) Init() NDArrayBinaryPrimaryGradientKernel {
	rv := objc.Call[NDArrayBinaryPrimaryGradientKernel](n_, objc.Sel("init"))
	return rv
}

func (n_ NDArrayBinaryPrimaryGradientKernel) InitWithDeviceSourceCountSourceGradientIndex(device metal.PDevice, count uint, sourceGradientIndex uint) NDArrayBinaryPrimaryGradientKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayBinaryPrimaryGradientKernel](n_, objc.Sel("initWithDevice:sourceCount:sourceGradientIndex:"), po0, count, sourceGradientIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel/3143524-initwithdevice?language=objc
func NewNDArrayBinaryPrimaryGradientKernelWithDeviceSourceCountSourceGradientIndex(device metal.PDevice, count uint, sourceGradientIndex uint) NDArrayBinaryPrimaryGradientKernel {
	instance := NDArrayBinaryPrimaryGradientKernelClass.Alloc().InitWithDeviceSourceCountSourceGradientIndex(device, count, sourceGradientIndex)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayBinaryPrimaryGradientKernel) InitWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayBinaryPrimaryGradientKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayBinaryPrimaryGradientKernel](n_, objc.Sel("initWithDevice:sourceCount:"), po0, count)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131741-initwithdevice?language=objc
func NewNDArrayBinaryPrimaryGradientKernelWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayBinaryPrimaryGradientKernel {
	instance := NDArrayBinaryPrimaryGradientKernelClass.Alloc().InitWithDeviceSourceCount(device, count)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayBinaryPrimaryGradientKernel) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayBinaryPrimaryGradientKernel {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayBinaryPrimaryGradientKernel](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131734-copywithzone?language=objc
func NDArrayBinaryPrimaryGradientKernel_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayBinaryPrimaryGradientKernel {
	instance := NDArrayBinaryPrimaryGradientKernelClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinaryprimarygradientkernel/3143513-encodetocommandbuffer?language=objc
func (n_ NDArrayBinaryPrimaryGradientKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArraySourceGradientGradientState(cmdBuf metal.PCommandBuffer, primarySourceArray INDArray, secondarySourceArray INDArray, gradient INDArray, state IState) NDArray {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	rv := objc.Call[NDArray](n_, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:"), po0, objc.Ptr(primarySourceArray), objc.Ptr(secondarySourceArray), objc.Ptr(gradient), objc.Ptr(state))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinaryprimarygradientkernel/3143513-encodetocommandbuffer?language=objc
func (n_ NDArrayBinaryPrimaryGradientKernel) EncodeToCommandBufferObjectPrimarySourceArraySecondarySourceArraySourceGradientGradientState(cmdBufObject objc.IObject, primarySourceArray INDArray, secondarySourceArray INDArray, gradient INDArray, state IState) NDArray {
	rv := objc.Call[NDArray](n_, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:sourceGradient:gradientState:"), objc.Ptr(cmdBufObject), objc.Ptr(primarySourceArray), objc.Ptr(secondarySourceArray), objc.Ptr(gradient), objc.Ptr(state))
	return rv
}
