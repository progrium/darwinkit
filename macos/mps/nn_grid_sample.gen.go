// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNGridSample] class.
var NNGridSampleClass = _NNGridSampleClass{objc.GetClass("MPSNNGridSample")}

type _NNGridSampleClass struct {
	objc.Class
}

// An interface definition for the [NNGridSample] class.
type INNGridSample interface {
	ICNNBinaryKernel
	UseGridValueAsInputCoordinate() bool
	SetUseGridValueAsInputCoordinate(value bool)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngridsample?language=objc
type NNGridSample struct {
	CNNBinaryKernel
}

func NNGridSampleFrom(ptr unsafe.Pointer) NNGridSample {
	return NNGridSample{
		CNNBinaryKernel: CNNBinaryKernelFrom(ptr),
	}
}

func (n_ NNGridSample) InitWithDevice(device metal.PDevice) NNGridSample {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNGridSample](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngridsample/3131871-initwithdevice?language=objc
func NewNNGridSampleWithDevice(device metal.PDevice) NNGridSample {
	instance := NNGridSampleClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNGridSampleClass) Alloc() NNGridSample {
	rv := objc.Call[NNGridSample](nc, objc.Sel("alloc"))
	return rv
}

func NNGridSample_Alloc() NNGridSample {
	return NNGridSampleClass.Alloc()
}

func (nc _NNGridSampleClass) New() NNGridSample {
	rv := objc.Call[NNGridSample](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNGridSample() NNGridSample {
	return NNGridSampleClass.New()
}

func (n_ NNGridSample) Init() NNGridSample {
	rv := objc.Call[NNGridSample](n_, objc.Sel("init"))
	return rv
}

func (n_ NNGridSample) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNGridSample {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNGridSample](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNGridSample_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNGridSample {
	instance := NNGridSampleClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngridsample/3131872-usegridvalueasinputcoordinate?language=objc
func (n_ NNGridSample) UseGridValueAsInputCoordinate() bool {
	rv := objc.Call[bool](n_, objc.Sel("useGridValueAsInputCoordinate"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnngridsample/3131872-usegridvalueasinputcoordinate?language=objc
func (n_ NNGridSample) SetUseGridValueAsInputCoordinate(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setUseGridValueAsInputCoordinate:"), value)
}
