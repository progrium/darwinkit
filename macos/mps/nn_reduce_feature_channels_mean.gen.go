// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReduceFeatureChannelsMean] class.
var NNReduceFeatureChannelsMeanClass = _NNReduceFeatureChannelsMeanClass{objc.GetClass("MPSNNReduceFeatureChannelsMean")}

type _NNReduceFeatureChannelsMeanClass struct {
	objc.Class
}

// An interface definition for the [NNReduceFeatureChannelsMean] class.
type INNReduceFeatureChannelsMean interface {
	INNReduceUnary
}

// A reduction filter that returns the mean value for each feature channel in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsmean?language=objc
type NNReduceFeatureChannelsMean struct {
	NNReduceUnary
}

func NNReduceFeatureChannelsMeanFrom(ptr unsafe.Pointer) NNReduceFeatureChannelsMean {
	return NNReduceFeatureChannelsMean{
		NNReduceUnary: NNReduceUnaryFrom(ptr),
	}
}

func (n_ NNReduceFeatureChannelsMean) InitWithDevice(device metal.PDevice) NNReduceFeatureChannelsMean {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceFeatureChannelsMean](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsmean/2942557-initwithdevice?language=objc
func NewNNReduceFeatureChannelsMeanWithDevice(device metal.PDevice) NNReduceFeatureChannelsMean {
	instance := NNReduceFeatureChannelsMeanClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNReduceFeatureChannelsMeanClass) Alloc() NNReduceFeatureChannelsMean {
	rv := objc.Call[NNReduceFeatureChannelsMean](nc, objc.Sel("alloc"))
	return rv
}

func NNReduceFeatureChannelsMean_Alloc() NNReduceFeatureChannelsMean {
	return NNReduceFeatureChannelsMeanClass.Alloc()
}

func (nc _NNReduceFeatureChannelsMeanClass) New() NNReduceFeatureChannelsMean {
	rv := objc.Call[NNReduceFeatureChannelsMean](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReduceFeatureChannelsMean() NNReduceFeatureChannelsMean {
	return NNReduceFeatureChannelsMeanClass.New()
}

func (n_ NNReduceFeatureChannelsMean) Init() NNReduceFeatureChannelsMean {
	rv := objc.Call[NNReduceFeatureChannelsMean](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReduceFeatureChannelsMean) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceFeatureChannelsMean {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceFeatureChannelsMean](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReduceFeatureChannelsMean_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceFeatureChannelsMean {
	instance := NNReduceFeatureChannelsMeanClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
