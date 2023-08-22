// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReduceFeatureChannelsAndWeightsMean] class.
var NNReduceFeatureChannelsAndWeightsMeanClass = _NNReduceFeatureChannelsAndWeightsMeanClass{objc.GetClass("MPSNNReduceFeatureChannelsAndWeightsMean")}

type _NNReduceFeatureChannelsAndWeightsMeanClass struct {
	objc.Class
}

// An interface definition for the [NNReduceFeatureChannelsAndWeightsMean] class.
type INNReduceFeatureChannelsAndWeightsMean interface {
	INNReduceBinary
}

// A reduction filter that returns the weighted sum for each feature channel in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightsmean?language=objc
type NNReduceFeatureChannelsAndWeightsMean struct {
	NNReduceBinary
}

func NNReduceFeatureChannelsAndWeightsMeanFrom(ptr unsafe.Pointer) NNReduceFeatureChannelsAndWeightsMean {
	return NNReduceFeatureChannelsAndWeightsMean{
		NNReduceBinary: NNReduceBinaryFrom(ptr),
	}
}

func (n_ NNReduceFeatureChannelsAndWeightsMean) InitWithDevice(device metal.PDevice) NNReduceFeatureChannelsAndWeightsMean {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceFeatureChannelsAndWeightsMean](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightsmean/2942537-initwithdevice?language=objc
func NewNNReduceFeatureChannelsAndWeightsMeanWithDevice(device metal.PDevice) NNReduceFeatureChannelsAndWeightsMean {
	instance := NNReduceFeatureChannelsAndWeightsMeanClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNReduceFeatureChannelsAndWeightsMeanClass) Alloc() NNReduceFeatureChannelsAndWeightsMean {
	rv := objc.Call[NNReduceFeatureChannelsAndWeightsMean](nc, objc.Sel("alloc"))
	return rv
}

func NNReduceFeatureChannelsAndWeightsMean_Alloc() NNReduceFeatureChannelsAndWeightsMean {
	return NNReduceFeatureChannelsAndWeightsMeanClass.Alloc()
}

func (nc _NNReduceFeatureChannelsAndWeightsMeanClass) New() NNReduceFeatureChannelsAndWeightsMean {
	rv := objc.Call[NNReduceFeatureChannelsAndWeightsMean](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReduceFeatureChannelsAndWeightsMean() NNReduceFeatureChannelsAndWeightsMean {
	return NNReduceFeatureChannelsAndWeightsMeanClass.New()
}

func (n_ NNReduceFeatureChannelsAndWeightsMean) Init() NNReduceFeatureChannelsAndWeightsMean {
	rv := objc.Call[NNReduceFeatureChannelsAndWeightsMean](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReduceFeatureChannelsAndWeightsMean) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceFeatureChannelsAndWeightsMean {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceFeatureChannelsAndWeightsMean](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReduceFeatureChannelsAndWeightsMean_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceFeatureChannelsAndWeightsMean {
	instance := NNReduceFeatureChannelsAndWeightsMeanClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
