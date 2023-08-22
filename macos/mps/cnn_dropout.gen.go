// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNDropout] class.
var CNNDropoutClass = _CNNDropoutClass{objc.GetClass("MPSCNNDropout")}

type _CNNDropoutClass struct {
	objc.Class
}

// An interface definition for the [CNNDropout] class.
type ICNNDropout interface {
	ICNNKernel
	KeepProbability() float64
	MaskStrideInPixels() metal.Size
	Seed() uint
}

// A dropout filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout?language=objc
type CNNDropout struct {
	CNNKernel
}

func CNNDropoutFrom(ptr unsafe.Pointer) CNNDropout {
	return CNNDropout{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (c_ CNNDropout) InitWithDeviceKeepProbabilitySeedMaskStrideInPixels(device metal.PDevice, keepProbability float64, seed uint, maskStrideInPixels metal.Size) CNNDropout {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDropout](c_, objc.Sel("initWithDevice:keepProbability:seed:maskStrideInPixels:"), po0, keepProbability, seed, maskStrideInPixels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/2942522-initwithdevice?language=objc
func NewCNNDropoutWithDeviceKeepProbabilitySeedMaskStrideInPixels(device metal.PDevice, keepProbability float64, seed uint, maskStrideInPixels metal.Size) CNNDropout {
	instance := CNNDropoutClass.Alloc().InitWithDeviceKeepProbabilitySeedMaskStrideInPixels(device, keepProbability, seed, maskStrideInPixels)
	instance.Autorelease()
	return instance
}

func (cc _CNNDropoutClass) Alloc() CNNDropout {
	rv := objc.Call[CNNDropout](cc, objc.Sel("alloc"))
	return rv
}

func CNNDropout_Alloc() CNNDropout {
	return CNNDropoutClass.Alloc()
}

func (cc _CNNDropoutClass) New() CNNDropout {
	rv := objc.Call[CNNDropout](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNDropout() CNNDropout {
	return CNNDropoutClass.New()
}

func (c_ CNNDropout) Init() CNNDropout {
	rv := objc.Call[CNNDropout](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNDropout) InitWithDevice(device metal.PDevice) CNNDropout {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDropout](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNDropoutWithDevice(device metal.PDevice) CNNDropout {
	instance := CNNDropoutClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNDropout) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNDropout {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNDropout](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNDropout_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNDropout {
	instance := CNNDropoutClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/2942524-keepprobability?language=objc
func (c_ CNNDropout) KeepProbability() float64 {
	rv := objc.Call[float64](c_, objc.Sel("keepProbability"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/2942519-maskstrideinpixels?language=objc
func (c_ CNNDropout) MaskStrideInPixels() metal.Size {
	rv := objc.Call[metal.Size](c_, objc.Sel("maskStrideInPixels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnndropout/2942517-seed?language=objc
func (c_ CNNDropout) Seed() uint {
	rv := objc.Call[uint](c_, objc.Sel("seed"))
	return rv
}
