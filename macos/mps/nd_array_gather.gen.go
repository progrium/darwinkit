// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NDArrayGather] class.
var NDArrayGatherClass = _NDArrayGatherClass{objc.GetClass("MPSNDArrayGather")}

type _NDArrayGatherClass struct {
	objc.Class
}

// An interface definition for the [NDArrayGather] class.
type INDArrayGather interface {
	INDArrayBinaryKernel
	Axis() uint
	SetAxis(value uint)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraygather?language=objc
type NDArrayGather struct {
	NDArrayBinaryKernel
}

func NDArrayGatherFrom(ptr unsafe.Pointer) NDArrayGather {
	return NDArrayGather{
		NDArrayBinaryKernel: NDArrayBinaryKernelFrom(ptr),
	}
}

func (nc _NDArrayGatherClass) Alloc() NDArrayGather {
	rv := objc.Call[NDArrayGather](nc, objc.Sel("alloc"))
	return rv
}

func NDArrayGather_Alloc() NDArrayGather {
	return NDArrayGatherClass.Alloc()
}

func (nc _NDArrayGatherClass) New() NDArrayGather {
	rv := objc.Call[NDArrayGather](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNDArrayGather() NDArrayGather {
	return NDArrayGatherClass.New()
}

func (n_ NDArrayGather) Init() NDArrayGather {
	rv := objc.Call[NDArrayGather](n_, objc.Sel("init"))
	return rv
}

func (n_ NDArrayGather) InitWithDevice(device metal.PDevice) NDArrayGather {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayGather](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinarykernel/3143501-initwithdevice?language=objc
func NewNDArrayGatherWithDevice(device metal.PDevice) NDArrayGather {
	instance := NDArrayGatherClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayGather) InitWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayGather {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayGather](n_, objc.Sel("initWithDevice:sourceCount:"), po0, count)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3175010-initwithdevice?language=objc
func NewNDArrayGatherWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayGather {
	instance := NDArrayGatherClass.Alloc().InitWithDeviceSourceCount(device, count)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayGather) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayGather {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayGather](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131734-copywithzone?language=objc
func NDArrayGather_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayGather {
	instance := NDArrayGatherClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraygather/3152529-axis?language=objc
func (n_ NDArrayGather) Axis() uint {
	rv := objc.Call[uint](n_, objc.Sel("axis"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraygather/3152529-axis?language=objc
func (n_ NDArrayGather) SetAxis(value uint) {
	objc.Call[objc.Void](n_, objc.Sel("setAxis:"), value)
}
