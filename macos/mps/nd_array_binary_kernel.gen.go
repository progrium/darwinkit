// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NDArrayBinaryKernel] class.
var NDArrayBinaryKernelClass = _NDArrayBinaryKernelClass{objc.GetClass("MPSNDArrayBinaryKernel")}

type _NDArrayBinaryKernelClass struct {
	objc.Class
}

// An interface definition for the [NDArrayBinaryKernel] class.
type INDArrayBinaryKernel interface {
	INDArrayMultiaryKernel
	EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayResultStateDestinationArray(cmdBuf metal.PCommandBuffer, primarySourceArray INDArray, secondarySourceArray INDArray, outGradientState IState, destination INDArray)
	EncodeToCommandBufferObjectPrimarySourceArraySecondarySourceArrayResultStateDestinationArray(cmdBufObject objc.IObject, primarySourceArray INDArray, secondarySourceArray INDArray, outGradientState IState, destination INDArray)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel?language=objc
type NDArrayBinaryKernel struct {
	NDArrayMultiaryKernel
}

func NDArrayBinaryKernelFrom(ptr unsafe.Pointer) NDArrayBinaryKernel {
	return NDArrayBinaryKernel{
		NDArrayMultiaryKernel: NDArrayMultiaryKernelFrom(ptr),
	}
}

func (n_ NDArrayBinaryKernel) InitWithDevice(device metal.PDevice) NDArrayBinaryKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayBinaryKernel](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143501-initwithdevice?language=objc
func NewNDArrayBinaryKernelWithDevice(device metal.PDevice) NDArrayBinaryKernel {
	instance := NDArrayBinaryKernelClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NDArrayBinaryKernelClass) Alloc() NDArrayBinaryKernel {
	rv := objc.Call[NDArrayBinaryKernel](nc, objc.Sel("alloc"))
	return rv
}

func NDArrayBinaryKernel_Alloc() NDArrayBinaryKernel {
	return NDArrayBinaryKernelClass.Alloc()
}

func (nc _NDArrayBinaryKernelClass) New() NDArrayBinaryKernel {
	rv := objc.Call[NDArrayBinaryKernel](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNDArrayBinaryKernel() NDArrayBinaryKernel {
	return NDArrayBinaryKernelClass.New()
}

func (n_ NDArrayBinaryKernel) Init() NDArrayBinaryKernel {
	rv := objc.Call[NDArrayBinaryKernel](n_, objc.Sel("init"))
	return rv
}

func (n_ NDArrayBinaryKernel) InitWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayBinaryKernel {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayBinaryKernel](n_, objc.Sel("initWithDevice:sourceCount:"), po0, count)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3175010-initwithdevice?language=objc
func NewNDArrayBinaryKernelWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayBinaryKernel {
	instance := NDArrayBinaryKernelClass.Alloc().InitWithDeviceSourceCount(device, count)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayBinaryKernel) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayBinaryKernel {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayBinaryKernel](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131734-copywithzone?language=objc
func NDArrayBinaryKernel_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayBinaryKernel {
	instance := NDArrayBinaryKernelClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143499-encodetocommandbuffer?language=objc
func (n_ NDArrayBinaryKernel) EncodeToCommandBufferPrimarySourceArraySecondarySourceArrayResultStateDestinationArray(cmdBuf metal.PCommandBuffer, primarySourceArray INDArray, secondarySourceArray INDArray, outGradientState IState, destination INDArray) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", cmdBuf)
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:resultState:destinationArray:"), po0, objc.Ptr(primarySourceArray), objc.Ptr(secondarySourceArray), objc.Ptr(outGradientState), objc.Ptr(destination))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143499-encodetocommandbuffer?language=objc
func (n_ NDArrayBinaryKernel) EncodeToCommandBufferObjectPrimarySourceArraySecondarySourceArrayResultStateDestinationArray(cmdBufObject objc.IObject, primarySourceArray INDArray, secondarySourceArray INDArray, outGradientState IState, destination INDArray) {
	objc.Call[objc.Void](n_, objc.Sel("encodeToCommandBuffer:primarySourceArray:secondarySourceArray:resultState:destinationArray:"), objc.Ptr(cmdBufObject), objc.Ptr(primarySourceArray), objc.Ptr(secondarySourceArray), objc.Ptr(outGradientState), objc.Ptr(destination))
}
