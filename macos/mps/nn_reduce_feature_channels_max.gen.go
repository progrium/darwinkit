// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReduceFeatureChannelsMax] class.
var NNReduceFeatureChannelsMaxClass = _NNReduceFeatureChannelsMaxClass{objc.GetClass("MPSNNReduceFeatureChannelsMax")}

type _NNReduceFeatureChannelsMaxClass struct {
	objc.Class
}

// An interface definition for the [NNReduceFeatureChannelsMax] class.
type INNReduceFeatureChannelsMax interface {
	INNReduceUnary
}

// A reduction filter that returns the maximum value for each feature channel in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsmax?language=objc
type NNReduceFeatureChannelsMax struct {
	NNReduceUnary
}

func NNReduceFeatureChannelsMaxFrom(ptr unsafe.Pointer) NNReduceFeatureChannelsMax {
	return NNReduceFeatureChannelsMax{
		NNReduceUnary: NNReduceUnaryFrom(ptr),
	}
}

func (n_ NNReduceFeatureChannelsMax) InitWithDevice(device metal.PDevice) NNReduceFeatureChannelsMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceFeatureChannelsMax](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsmax/2942532-initwithdevice?language=objc
func NewNNReduceFeatureChannelsMaxWithDevice(device metal.PDevice) NNReduceFeatureChannelsMax {
	instance := NNReduceFeatureChannelsMaxClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNReduceFeatureChannelsMaxClass) Alloc() NNReduceFeatureChannelsMax {
	rv := objc.Call[NNReduceFeatureChannelsMax](nc, objc.Sel("alloc"))
	return rv
}

func NNReduceFeatureChannelsMax_Alloc() NNReduceFeatureChannelsMax {
	return NNReduceFeatureChannelsMaxClass.Alloc()
}

func (nc _NNReduceFeatureChannelsMaxClass) New() NNReduceFeatureChannelsMax {
	rv := objc.Call[NNReduceFeatureChannelsMax](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReduceFeatureChannelsMax() NNReduceFeatureChannelsMax {
	return NNReduceFeatureChannelsMaxClass.New()
}

func (n_ NNReduceFeatureChannelsMax) Init() NNReduceFeatureChannelsMax {
	rv := objc.Call[NNReduceFeatureChannelsMax](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReduceFeatureChannelsMax) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceFeatureChannelsMax {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceFeatureChannelsMax](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReduceFeatureChannelsMax_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceFeatureChannelsMax {
	instance := NNReduceFeatureChannelsMaxClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
