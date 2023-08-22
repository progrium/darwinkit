// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNUpsampling] class.
var CNNUpsamplingClass = _CNNUpsamplingClass{objc.GetClass("MPSCNNUpsampling")}

type _CNNUpsamplingClass struct {
	objc.Class
}

// An interface definition for the [CNNUpsampling] class.
type ICNNUpsampling interface {
	ICNNKernel
	AlignCorners() bool
	ScaleFactorY() float64
	ScaleFactorX() float64
}

// A filter that resamples an existing MPS image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsampling?language=objc
type CNNUpsampling struct {
	CNNKernel
}

func CNNUpsamplingFrom(ptr unsafe.Pointer) CNNUpsampling {
	return CNNUpsampling{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (cc _CNNUpsamplingClass) Alloc() CNNUpsampling {
	rv := objc.Call[CNNUpsampling](cc, objc.Sel("alloc"))
	return rv
}

func CNNUpsampling_Alloc() CNNUpsampling {
	return CNNUpsamplingClass.Alloc()
}

func (cc _CNNUpsamplingClass) New() CNNUpsampling {
	rv := objc.Call[CNNUpsampling](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNUpsampling() CNNUpsampling {
	return CNNUpsamplingClass.New()
}

func (c_ CNNUpsampling) Init() CNNUpsampling {
	rv := objc.Call[CNNUpsampling](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNUpsampling) InitWithDevice(device metal.PDevice) CNNUpsampling {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNUpsampling](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNUpsamplingWithDevice(device metal.PDevice) CNNUpsampling {
	instance := CNNUpsamplingClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNUpsampling) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNUpsampling {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNUpsampling](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNUpsampling_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNUpsampling {
	instance := CNNUpsamplingClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsampling/2966660-aligncorners?language=objc
func (c_ CNNUpsampling) AlignCorners() bool {
	rv := objc.Call[bool](c_, objc.Sel("alignCorners"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsampling/2875154-scalefactory?language=objc
func (c_ CNNUpsampling) ScaleFactorY() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleFactorY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnupsampling/2875206-scalefactorx?language=objc
func (c_ CNNUpsampling) ScaleFactorX() float64 {
	rv := objc.Call[float64](c_, objc.Sel("scaleFactorX"))
	return rv
}
