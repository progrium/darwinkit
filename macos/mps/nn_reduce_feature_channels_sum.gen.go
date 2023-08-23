// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReduceFeatureChannelsSum] class.
var NNReduceFeatureChannelsSumClass = _NNReduceFeatureChannelsSumClass{objc.GetClass("MPSNNReduceFeatureChannelsSum")}

type _NNReduceFeatureChannelsSumClass struct {
	objc.Class
}

// An interface definition for the [NNReduceFeatureChannelsSum] class.
type INNReduceFeatureChannelsSum interface {
	INNReduceUnary
	Weight() float64
	SetWeight(value float64)
}

// A reduction filter that returns the sum of all values for each feature channel in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelssum?language=objc
type NNReduceFeatureChannelsSum struct {
	NNReduceUnary
}

func NNReduceFeatureChannelsSumFrom(ptr unsafe.Pointer) NNReduceFeatureChannelsSum {
	return NNReduceFeatureChannelsSum{
		NNReduceUnary: NNReduceUnaryFrom(ptr),
	}
}

func (n_ NNReduceFeatureChannelsSum) InitWithDevice(device metal.PDevice) NNReduceFeatureChannelsSum {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceFeatureChannelsSum](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelssum/2942538-initwithdevice?language=objc
func NewNNReduceFeatureChannelsSumWithDevice(device metal.PDevice) NNReduceFeatureChannelsSum {
	instance := NNReduceFeatureChannelsSumClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNReduceFeatureChannelsSumClass) Alloc() NNReduceFeatureChannelsSum {
	rv := objc.Call[NNReduceFeatureChannelsSum](nc, objc.Sel("alloc"))
	return rv
}

func NNReduceFeatureChannelsSum_Alloc() NNReduceFeatureChannelsSum {
	return NNReduceFeatureChannelsSumClass.Alloc()
}

func (nc _NNReduceFeatureChannelsSumClass) New() NNReduceFeatureChannelsSum {
	rv := objc.Call[NNReduceFeatureChannelsSum](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReduceFeatureChannelsSum() NNReduceFeatureChannelsSum {
	return NNReduceFeatureChannelsSumClass.New()
}

func (n_ NNReduceFeatureChannelsSum) Init() NNReduceFeatureChannelsSum {
	rv := objc.Call[NNReduceFeatureChannelsSum](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReduceFeatureChannelsSum) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceFeatureChannelsSum {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceFeatureChannelsSum](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReduceFeatureChannelsSum_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceFeatureChannelsSum {
	instance := NNReduceFeatureChannelsSumClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelssum/2942545-weight?language=objc
func (n_ NNReduceFeatureChannelsSum) Weight() float64 {
	rv := objc.Call[float64](n_, objc.Sel("weight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelssum/2942545-weight?language=objc
func (n_ NNReduceFeatureChannelsSum) SetWeight(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setWeight:"), value)
}
