// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReduceFeatureChannelsMin] class.
var NNReduceFeatureChannelsMinClass = _NNReduceFeatureChannelsMinClass{objc.GetClass("MPSNNReduceFeatureChannelsMin")}

type _NNReduceFeatureChannelsMinClass struct {
	objc.Class
}

// An interface definition for the [NNReduceFeatureChannelsMin] class.
type INNReduceFeatureChannelsMin interface {
	INNReduceUnary
}

// A reduction filter that returns the minimum value for each feature channel in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsmin?language=objc
type NNReduceFeatureChannelsMin struct {
	NNReduceUnary
}

func NNReduceFeatureChannelsMinFrom(ptr unsafe.Pointer) NNReduceFeatureChannelsMin {
	return NNReduceFeatureChannelsMin{
		NNReduceUnary: NNReduceUnaryFrom(ptr),
	}
}

func (n_ NNReduceFeatureChannelsMin) InitWithDevice(device metal.PDevice) NNReduceFeatureChannelsMin {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceFeatureChannelsMin](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsmin/2942565-initwithdevice?language=objc
func NewNNReduceFeatureChannelsMinWithDevice(device metal.PDevice) NNReduceFeatureChannelsMin {
	instance := NNReduceFeatureChannelsMinClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNReduceFeatureChannelsMinClass) Alloc() NNReduceFeatureChannelsMin {
	rv := objc.Call[NNReduceFeatureChannelsMin](nc, objc.Sel("alloc"))
	return rv
}

func NNReduceFeatureChannelsMin_Alloc() NNReduceFeatureChannelsMin {
	return NNReduceFeatureChannelsMinClass.Alloc()
}

func (nc _NNReduceFeatureChannelsMinClass) New() NNReduceFeatureChannelsMin {
	rv := objc.Call[NNReduceFeatureChannelsMin](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReduceFeatureChannelsMin() NNReduceFeatureChannelsMin {
	return NNReduceFeatureChannelsMinClass.New()
}

func (n_ NNReduceFeatureChannelsMin) Init() NNReduceFeatureChannelsMin {
	rv := objc.Call[NNReduceFeatureChannelsMin](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReduceFeatureChannelsMin) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceFeatureChannelsMin {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceFeatureChannelsMin](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReduceFeatureChannelsMin_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceFeatureChannelsMin {
	instance := NNReduceFeatureChannelsMinClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
