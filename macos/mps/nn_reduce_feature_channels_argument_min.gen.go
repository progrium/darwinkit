// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReduceFeatureChannelsArgumentMin] class.
var NNReduceFeatureChannelsArgumentMinClass = _NNReduceFeatureChannelsArgumentMinClass{objc.GetClass("MPSNNReduceFeatureChannelsArgumentMin")}

type _NNReduceFeatureChannelsArgumentMinClass struct {
	objc.Class
}

// An interface definition for the [NNReduceFeatureChannelsArgumentMin] class.
type INNReduceFeatureChannelsArgumentMin interface {
	INNReduceUnary
}

// A reduction filter that returns the index of the location of the minimum value for each feature channel in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsargumentmin?language=objc
type NNReduceFeatureChannelsArgumentMin struct {
	NNReduceUnary
}

func NNReduceFeatureChannelsArgumentMinFrom(ptr unsafe.Pointer) NNReduceFeatureChannelsArgumentMin {
	return NNReduceFeatureChannelsArgumentMin{
		NNReduceUnary: NNReduceUnaryFrom(ptr),
	}
}

func (n_ NNReduceFeatureChannelsArgumentMin) InitWithDevice(device metal.PDevice) NNReduceFeatureChannelsArgumentMin {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceFeatureChannelsArgumentMin](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsargumentmin/2976520-initwithdevice?language=objc
func NewNNReduceFeatureChannelsArgumentMinWithDevice(device metal.PDevice) NNReduceFeatureChannelsArgumentMin {
	instance := NNReduceFeatureChannelsArgumentMinClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNReduceFeatureChannelsArgumentMinClass) Alloc() NNReduceFeatureChannelsArgumentMin {
	rv := objc.Call[NNReduceFeatureChannelsArgumentMin](nc, objc.Sel("alloc"))
	return rv
}

func NNReduceFeatureChannelsArgumentMin_Alloc() NNReduceFeatureChannelsArgumentMin {
	return NNReduceFeatureChannelsArgumentMinClass.Alloc()
}

func (nc _NNReduceFeatureChannelsArgumentMinClass) New() NNReduceFeatureChannelsArgumentMin {
	rv := objc.Call[NNReduceFeatureChannelsArgumentMin](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReduceFeatureChannelsArgumentMin() NNReduceFeatureChannelsArgumentMin {
	return NNReduceFeatureChannelsArgumentMinClass.New()
}

func (n_ NNReduceFeatureChannelsArgumentMin) Init() NNReduceFeatureChannelsArgumentMin {
	rv := objc.Call[NNReduceFeatureChannelsArgumentMin](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReduceFeatureChannelsArgumentMin) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceFeatureChannelsArgumentMin {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceFeatureChannelsArgumentMin](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReduceFeatureChannelsArgumentMin_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceFeatureChannelsArgumentMin {
	instance := NNReduceFeatureChannelsArgumentMinClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
