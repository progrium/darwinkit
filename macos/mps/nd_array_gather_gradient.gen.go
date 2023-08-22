// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NDArrayGatherGradient] class.
var NDArrayGatherGradientClass = _NDArrayGatherGradientClass{objc.GetClass("MPSNDArrayGatherGradient")}

type _NDArrayGatherGradientClass struct {
	objc.Class
}

// An interface definition for the [NDArrayGatherGradient] class.
type INDArrayGatherGradient interface {
	INDArrayBinaryPrimaryGradientKernel
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraygathergradient?language=objc
type NDArrayGatherGradient struct {
	NDArrayBinaryPrimaryGradientKernel
}

func NDArrayGatherGradientFrom(ptr unsafe.Pointer) NDArrayGatherGradient {
	return NDArrayGatherGradient{
		NDArrayBinaryPrimaryGradientKernel: NDArrayBinaryPrimaryGradientKernelFrom(ptr),
	}
}

func (nc _NDArrayGatherGradientClass) Alloc() NDArrayGatherGradient {
	rv := objc.Call[NDArrayGatherGradient](nc, objc.Sel("alloc"))
	return rv
}

func NDArrayGatherGradient_Alloc() NDArrayGatherGradient {
	return NDArrayGatherGradientClass.Alloc()
}

func (nc _NDArrayGatherGradientClass) New() NDArrayGatherGradient {
	rv := objc.Call[NDArrayGatherGradient](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNDArrayGatherGradient() NDArrayGatherGradient {
	return NDArrayGatherGradientClass.New()
}

func (n_ NDArrayGatherGradient) Init() NDArrayGatherGradient {
	rv := objc.Call[NDArrayGatherGradient](n_, objc.Sel("init"))
	return rv
}

func (n_ NDArrayGatherGradient) InitWithDevice(device metal.PDevice) NDArrayGatherGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayGatherGradient](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraybinaryprimarygradientkernel/3143515-initwithdevice?language=objc
func NewNDArrayGatherGradientWithDevice(device metal.PDevice) NDArrayGatherGradient {
	instance := NDArrayGatherGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayGatherGradient) InitWithDeviceSourceCountSourceGradientIndex(device metal.PDevice, count uint, sourceGradientIndex uint) NDArrayGatherGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayGatherGradient](n_, objc.Sel("initWithDevice:sourceCount:sourceGradientIndex:"), po0, count, sourceGradientIndex)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarygradientkernel/3143524-initwithdevice?language=objc
func NewNDArrayGatherGradientWithDeviceSourceCountSourceGradientIndex(device metal.PDevice, count uint, sourceGradientIndex uint) NDArrayGatherGradient {
	instance := NDArrayGatherGradientClass.Alloc().InitWithDeviceSourceCountSourceGradientIndex(device, count, sourceGradientIndex)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayGatherGradient) InitWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayGatherGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayGatherGradient](n_, objc.Sel("initWithDevice:sourceCount:"), po0, count)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131741-initwithdevice?language=objc
func NewNDArrayGatherGradientWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayGatherGradient {
	instance := NDArrayGatherGradientClass.Alloc().InitWithDeviceSourceCount(device, count)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayGatherGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayGatherGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayGatherGradient](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131734-copywithzone?language=objc
func NDArrayGatherGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayGatherGradient {
	instance := NDArrayGatherGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
