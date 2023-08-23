// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReduceUnary] class.
var NNReduceUnaryClass = _NNReduceUnaryClass{objc.GetClass("MPSNNReduceUnary")}

type _NNReduceUnaryClass struct {
	objc.Class
}

// An interface definition for the [NNReduceUnary] class.
type INNReduceUnary interface {
	ICNNKernel
	ClipRectSource() metal.Region
	SetClipRectSource(value metal.Region)
}

// The base class for unary reduction filters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreduceunary?language=objc
type NNReduceUnary struct {
	CNNKernel
}

func NNReduceUnaryFrom(ptr unsafe.Pointer) NNReduceUnary {
	return NNReduceUnary{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (nc _NNReduceUnaryClass) Alloc() NNReduceUnary {
	rv := objc.Call[NNReduceUnary](nc, objc.Sel("alloc"))
	return rv
}

func NNReduceUnary_Alloc() NNReduceUnary {
	return NNReduceUnaryClass.Alloc()
}

func (nc _NNReduceUnaryClass) New() NNReduceUnary {
	rv := objc.Call[NNReduceUnary](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReduceUnary() NNReduceUnary {
	return NNReduceUnaryClass.New()
}

func (n_ NNReduceUnary) Init() NNReduceUnary {
	rv := objc.Call[NNReduceUnary](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReduceUnary) InitWithDevice(device metal.PDevice) NNReduceUnary {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceUnary](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewNNReduceUnaryWithDevice(device metal.PDevice) NNReduceUnary {
	instance := NNReduceUnaryClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (n_ NNReduceUnary) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceUnary {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceUnary](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReduceUnary_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceUnary {
	instance := NNReduceUnaryClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreduceunary/2942547-cliprectsource?language=objc
func (n_ NNReduceUnary) ClipRectSource() metal.Region {
	rv := objc.Call[metal.Region](n_, objc.Sel("clipRectSource"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreduceunary/2942547-cliprectsource?language=objc
func (n_ NNReduceUnary) SetClipRectSource(value metal.Region) {
	objc.Call[objc.Void](n_, objc.Sel("setClipRectSource:"), value)
}
