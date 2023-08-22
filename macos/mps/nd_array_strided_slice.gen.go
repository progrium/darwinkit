// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NDArrayStridedSlice] class.
var NDArrayStridedSliceClass = _NDArrayStridedSliceClass{objc.GetClass("MPSNDArrayStridedSlice")}

type _NDArrayStridedSliceClass struct {
	objc.Class
}

// An interface definition for the [NDArrayStridedSlice] class.
type INDArrayStridedSlice interface {
	INDArrayUnaryKernel
	Strides() NDArrayOffsets
	SetStrides(value NDArrayOffsets)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraystridedslice?language=objc
type NDArrayStridedSlice struct {
	NDArrayUnaryKernel
}

func NDArrayStridedSliceFrom(ptr unsafe.Pointer) NDArrayStridedSlice {
	return NDArrayStridedSlice{
		NDArrayUnaryKernel: NDArrayUnaryKernelFrom(ptr),
	}
}

func (nc _NDArrayStridedSliceClass) Alloc() NDArrayStridedSlice {
	rv := objc.Call[NDArrayStridedSlice](nc, objc.Sel("alloc"))
	return rv
}

func NDArrayStridedSlice_Alloc() NDArrayStridedSlice {
	return NDArrayStridedSliceClass.Alloc()
}

func (nc _NDArrayStridedSliceClass) New() NDArrayStridedSlice {
	rv := objc.Call[NDArrayStridedSlice](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNDArrayStridedSlice() NDArrayStridedSlice {
	return NDArrayStridedSliceClass.New()
}

func (n_ NDArrayStridedSlice) Init() NDArrayStridedSlice {
	rv := objc.Call[NDArrayStridedSlice](n_, objc.Sel("init"))
	return rv
}

func (n_ NDArrayStridedSlice) InitWithDevice(device metal.PDevice) NDArrayStridedSlice {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayStridedSlice](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarykernel/3143540-initwithdevice?language=objc
func NewNDArrayStridedSliceWithDevice(device metal.PDevice) NDArrayStridedSlice {
	instance := NDArrayStridedSliceClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayStridedSlice) InitWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayStridedSlice {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayStridedSlice](n_, objc.Sel("initWithDevice:sourceCount:"), po0, count)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3175010-initwithdevice?language=objc
func NewNDArrayStridedSliceWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayStridedSlice {
	instance := NDArrayStridedSliceClass.Alloc().InitWithDeviceSourceCount(device, count)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayStridedSlice) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayStridedSlice {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayStridedSlice](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131734-copywithzone?language=objc
func NDArrayStridedSlice_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayStridedSlice {
	instance := NDArrayStridedSliceClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraystridedslice/3143546-strides?language=objc
func (n_ NDArrayStridedSlice) Strides() NDArrayOffsets {
	rv := objc.Call[NDArrayOffsets](n_, objc.Sel("strides"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraystridedslice/3143546-strides?language=objc
func (n_ NDArrayStridedSlice) SetStrides(value NDArrayOffsets) {
	objc.Call[objc.Void](n_, objc.Sel("setStrides:"), value)
}
