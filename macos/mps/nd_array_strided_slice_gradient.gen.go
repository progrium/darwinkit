// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NDArrayStridedSliceGradient] class.
var NDArrayStridedSliceGradientClass = _NDArrayStridedSliceGradientClass{objc.GetClass("MPSNDArrayStridedSliceGradient")}

type _NDArrayStridedSliceGradientClass struct {
	objc.Class
}

// An interface definition for the [NDArrayStridedSliceGradient] class.
type INDArrayStridedSliceGradient interface {
	INDArrayUnaryGradientKernel
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraystridedslicegradient?language=objc
type NDArrayStridedSliceGradient struct {
	NDArrayUnaryGradientKernel
}

func NDArrayStridedSliceGradientFrom(ptr unsafe.Pointer) NDArrayStridedSliceGradient {
	return NDArrayStridedSliceGradient{
		NDArrayUnaryGradientKernel: NDArrayUnaryGradientKernelFrom(ptr),
	}
}

func (nc _NDArrayStridedSliceGradientClass) Alloc() NDArrayStridedSliceGradient {
	rv := objc.Call[NDArrayStridedSliceGradient](nc, objc.Sel("alloc"))
	return rv
}

func NDArrayStridedSliceGradient_Alloc() NDArrayStridedSliceGradient {
	return NDArrayStridedSliceGradientClass.Alloc()
}

func (nc _NDArrayStridedSliceGradientClass) New() NDArrayStridedSliceGradient {
	rv := objc.Call[NDArrayStridedSliceGradient](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNDArrayStridedSliceGradient() NDArrayStridedSliceGradient {
	return NDArrayStridedSliceGradientClass.New()
}

func (n_ NDArrayStridedSliceGradient) Init() NDArrayStridedSliceGradient {
	rv := objc.Call[NDArrayStridedSliceGradient](n_, objc.Sel("init"))
	return rv
}

func (n_ NDArrayStridedSliceGradient) InitWithDevice(device metal.PDevice) NDArrayStridedSliceGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayStridedSliceGradient](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarrayunarygradientkernel/3143532-initwithdevice?language=objc
func NewNDArrayStridedSliceGradientWithDevice(device metal.PDevice) NDArrayStridedSliceGradient {
	instance := NDArrayStridedSliceGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayStridedSliceGradient) InitWithDeviceSourceCountSourceGradientIndex(device metal.PDevice, count uint, sourceGradientIndex uint) NDArrayStridedSliceGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayStridedSliceGradient](n_, objc.Sel("initWithDevice:sourceCount:sourceGradientIndex:"), po0, count, sourceGradientIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel/3143524-initwithdevice?language=objc
func NewNDArrayStridedSliceGradientWithDeviceSourceCountSourceGradientIndex(device metal.PDevice, count uint, sourceGradientIndex uint) NDArrayStridedSliceGradient {
	instance := NDArrayStridedSliceGradientClass.Alloc().InitWithDeviceSourceCountSourceGradientIndex(device, count, sourceGradientIndex)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayStridedSliceGradient) InitWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayStridedSliceGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayStridedSliceGradient](n_, objc.Sel("initWithDevice:sourceCount:"), po0, count)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131741-initwithdevice?language=objc
func NewNDArrayStridedSliceGradientWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayStridedSliceGradient {
	instance := NDArrayStridedSliceGradientClass.Alloc().InitWithDeviceSourceCount(device, count)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayStridedSliceGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayStridedSliceGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayStridedSliceGradient](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131734-copywithzone?language=objc
func NDArrayStridedSliceGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayStridedSliceGradient {
	instance := NDArrayStridedSliceGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
