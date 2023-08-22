// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNUpsamplingNearestGradient] class.
var CNNUpsamplingNearestGradientClass = _CNNUpsamplingNearestGradientClass{objc.GetClass("MPSCNNUpsamplingNearestGradient")}

type _CNNUpsamplingNearestGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNUpsamplingNearestGradient] class.
type ICNNUpsamplingNearestGradient interface {
	ICNNUpsamplingGradient
}

// A gradient upsampling filter that samples the pixel nearest to the source when upsampling to the destination pixel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradient?language=objc
type CNNUpsamplingNearestGradient struct {
	CNNUpsamplingGradient
}

func CNNUpsamplingNearestGradientFrom(ptr unsafe.Pointer) CNNUpsamplingNearestGradient {
	return CNNUpsamplingNearestGradient{
		CNNUpsamplingGradient: CNNUpsamplingGradientFrom(ptr),
	}
}

func (c_ CNNUpsamplingNearestGradient) InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.PDevice, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingNearestGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNUpsamplingNearestGradient](c_, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:"), po0, integerScaleFactorX, integerScaleFactorY)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearestgradient/2947920-initwithdevice?language=objc
func NewCNNUpsamplingNearestGradientWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.PDevice, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingNearestGradient {
	instance := CNNUpsamplingNearestGradientClass.Alloc().InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device, integerScaleFactorX, integerScaleFactorY)
	instance.Autorelease()
	return instance
}

func (cc _CNNUpsamplingNearestGradientClass) Alloc() CNNUpsamplingNearestGradient {
	rv := objc.Call[CNNUpsamplingNearestGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNUpsamplingNearestGradient_Alloc() CNNUpsamplingNearestGradient {
	return CNNUpsamplingNearestGradientClass.Alloc()
}

func (cc _CNNUpsamplingNearestGradientClass) New() CNNUpsamplingNearestGradient {
	rv := objc.Call[CNNUpsamplingNearestGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNUpsamplingNearestGradient() CNNUpsamplingNearestGradient {
	return CNNUpsamplingNearestGradientClass.New()
}

func (c_ CNNUpsamplingNearestGradient) Init() CNNUpsamplingNearestGradient {
	rv := objc.Call[CNNUpsamplingNearestGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNUpsamplingNearestGradient) InitWithDevice(device metal.PDevice) CNNUpsamplingNearestGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNUpsamplingNearestGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNUpsamplingNearestGradientWithDevice(device metal.PDevice) CNNUpsamplingNearestGradient {
	instance := CNNUpsamplingNearestGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNUpsamplingNearestGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNUpsamplingNearestGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNUpsamplingNearestGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNUpsamplingNearestGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNUpsamplingNearestGradient {
	instance := CNNUpsamplingNearestGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
