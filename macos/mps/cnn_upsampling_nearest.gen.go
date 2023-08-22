// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNUpsamplingNearest] class.
var CNNUpsamplingNearestClass = _CNNUpsamplingNearestClass{objc.GetClass("MPSCNNUpsamplingNearest")}

type _CNNUpsamplingNearestClass struct {
	objc.Class
}

// An interface definition for the [CNNUpsamplingNearest] class.
type ICNNUpsamplingNearest interface {
	ICNNUpsampling
}

// A nearest spatial upsampling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearest?language=objc
type CNNUpsamplingNearest struct {
	CNNUpsampling
}

func CNNUpsamplingNearestFrom(ptr unsafe.Pointer) CNNUpsamplingNearest {
	return CNNUpsamplingNearest{
		CNNUpsampling: CNNUpsamplingFrom(ptr),
	}
}

func (c_ CNNUpsamplingNearest) InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.PDevice, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingNearest {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNUpsamplingNearest](c_, objc.Sel("initWithDevice:integerScaleFactorX:integerScaleFactorY:"), po0, integerScaleFactorX, integerScaleFactorY)
	return rv
}

// Initializes a nearest spatial upsampling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsamplingnearest/2875223-initwithdevice?language=objc
func NewCNNUpsamplingNearestWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device metal.PDevice, integerScaleFactorX uint, integerScaleFactorY uint) CNNUpsamplingNearest {
	instance := CNNUpsamplingNearestClass.Alloc().InitWithDeviceIntegerScaleFactorXIntegerScaleFactorY(device, integerScaleFactorX, integerScaleFactorY)
	instance.Autorelease()
	return instance
}

func (cc _CNNUpsamplingNearestClass) Alloc() CNNUpsamplingNearest {
	rv := objc.Call[CNNUpsamplingNearest](cc, objc.Sel("alloc"))
	return rv
}

func CNNUpsamplingNearest_Alloc() CNNUpsamplingNearest {
	return CNNUpsamplingNearestClass.Alloc()
}

func (cc _CNNUpsamplingNearestClass) New() CNNUpsamplingNearest {
	rv := objc.Call[CNNUpsamplingNearest](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNUpsamplingNearest() CNNUpsamplingNearest {
	return CNNUpsamplingNearestClass.New()
}

func (c_ CNNUpsamplingNearest) Init() CNNUpsamplingNearest {
	rv := objc.Call[CNNUpsamplingNearest](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNUpsamplingNearest) InitWithDevice(device metal.PDevice) CNNUpsamplingNearest {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNUpsamplingNearest](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNUpsamplingNearestWithDevice(device metal.PDevice) CNNUpsamplingNearest {
	instance := CNNUpsamplingNearestClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNUpsamplingNearest) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNUpsamplingNearest {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNUpsamplingNearest](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNUpsamplingNearest_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNUpsamplingNearest {
	instance := CNNUpsamplingNearestClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
