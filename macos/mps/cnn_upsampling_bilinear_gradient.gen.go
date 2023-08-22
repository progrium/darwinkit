// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNUpsamplingBilinearGradient] class.
var CNNUpsamplingBilinearGradientClass = _CNNUpsamplingBilinearGradientClass{objc.GetClass("MPSCNNUpsamplingBilinearGradient")}

type _CNNUpsamplingBilinearGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNUpsamplingBilinearGradient] class.
type ICNNUpsamplingBilinearGradient interface {
	ICNNUpsamplingGradient
}

// A gradient bilinear spatial upsampling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilineargradient?language=objc
type CNNUpsamplingBilinearGradient struct {
	CNNUpsamplingGradient
}

func CNNUpsamplingBilinearGradientFrom(ptr unsafe.Pointer) CNNUpsamplingBilinearGradient {
	return CNNUpsamplingBilinearGradient{
		CNNUpsamplingGradient: CNNUpsamplingGradientFrom(ptr),
	}
}

func (c_ CNNUpsamplingBilinearGradient) InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.PDevice, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingBilinearGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNUpsamplingBilinearGradient](c_, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:"), po0, integerScaleFactorX, integerScaleFactorY)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingbilineargradient/2947918-initwithdevice?language=objc
func NewCNNUpsamplingBilinearGradientWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.PDevice, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingBilinearGradient {
	instance := CNNUpsamplingBilinearGradientClass.Alloc().InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device, integerScaleFactorX, integerScaleFactorY)
	instance.Autorelease()
	return instance
}

func (cc _CNNUpsamplingBilinearGradientClass) Alloc() CNNUpsamplingBilinearGradient {
	rv := objc.Call[CNNUpsamplingBilinearGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNUpsamplingBilinearGradient_Alloc() CNNUpsamplingBilinearGradient {
	return CNNUpsamplingBilinearGradientClass.Alloc()
}

func (cc _CNNUpsamplingBilinearGradientClass) New() CNNUpsamplingBilinearGradient {
	rv := objc.Call[CNNUpsamplingBilinearGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNUpsamplingBilinearGradient() CNNUpsamplingBilinearGradient {
	return CNNUpsamplingBilinearGradientClass.New()
}

func (c_ CNNUpsamplingBilinearGradient) Init() CNNUpsamplingBilinearGradient {
	rv := objc.Call[CNNUpsamplingBilinearGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNUpsamplingBilinearGradient) InitWithDevice(device metal.PDevice) CNNUpsamplingBilinearGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNUpsamplingBilinearGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNUpsamplingBilinearGradientWithDevice(device metal.PDevice) CNNUpsamplingBilinearGradient {
	instance := CNNUpsamplingBilinearGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNUpsamplingBilinearGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNUpsamplingBilinearGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNUpsamplingBilinearGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNUpsamplingBilinearGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNUpsamplingBilinearGradient {
	instance := CNNUpsamplingBilinearGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
