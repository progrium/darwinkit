// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReduceFeatureChannelsArgumentMax] class.
var NNReduceFeatureChannelsArgumentMaxClass = _NNReduceFeatureChannelsArgumentMaxClass{objc.GetClass("MPSNNReduceFeatureChannelsArgumentMax")}

type _NNReduceFeatureChannelsArgumentMaxClass struct {
	objc.Class
}

// An interface definition for the [NNReduceFeatureChannelsArgumentMax] class.
type INNReduceFeatureChannelsArgumentMax interface {
	INNReduceUnary
}

// A reduction filter that returns the index of the location of the maximum value for each feature channel in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsargumentmax?language=objc
type NNReduceFeatureChannelsArgumentMax struct {
	NNReduceUnary
}

func NNReduceFeatureChannelsArgumentMaxFrom(ptr unsafe.Pointer) NNReduceFeatureChannelsArgumentMax {
	return NNReduceFeatureChannelsArgumentMax{
		NNReduceUnary: NNReduceUnaryFrom(ptr),
	}
}

func (n_ NNReduceFeatureChannelsArgumentMax) InitWithDevice(device metal.PDevice) NNReduceFeatureChannelsArgumentMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceFeatureChannelsArgumentMax](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsargumentmax/2976518-initwithdevice?language=objc
func NewNNReduceFeatureChannelsArgumentMaxWithDevice(device metal.PDevice) NNReduceFeatureChannelsArgumentMax {
	instance := NNReduceFeatureChannelsArgumentMaxClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNReduceFeatureChannelsArgumentMaxClass) Alloc() NNReduceFeatureChannelsArgumentMax {
	rv := objc.Call[NNReduceFeatureChannelsArgumentMax](nc, objc.Sel("alloc"))
	return rv
}

func NNReduceFeatureChannelsArgumentMax_Alloc() NNReduceFeatureChannelsArgumentMax {
	return NNReduceFeatureChannelsArgumentMaxClass.Alloc()
}

func (nc _NNReduceFeatureChannelsArgumentMaxClass) New() NNReduceFeatureChannelsArgumentMax {
	rv := objc.Call[NNReduceFeatureChannelsArgumentMax](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReduceFeatureChannelsArgumentMax() NNReduceFeatureChannelsArgumentMax {
	return NNReduceFeatureChannelsArgumentMaxClass.New()
}

func (n_ NNReduceFeatureChannelsArgumentMax) Init() NNReduceFeatureChannelsArgumentMax {
	rv := objc.Call[NNReduceFeatureChannelsArgumentMax](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReduceFeatureChannelsArgumentMax) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceFeatureChannelsArgumentMax {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceFeatureChannelsArgumentMax](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReduceFeatureChannelsArgumentMax_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceFeatureChannelsArgumentMax {
	instance := NNReduceFeatureChannelsArgumentMaxClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
