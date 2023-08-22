// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReduceFeatureChannelsAndWeightsSum] class.
var NNReduceFeatureChannelsAndWeightsSumClass = _NNReduceFeatureChannelsAndWeightsSumClass{objc.GetClass("MPSNNReduceFeatureChannelsAndWeightsSum")}

type _NNReduceFeatureChannelsAndWeightsSumClass struct {
	objc.Class
}

// An interface definition for the [NNReduceFeatureChannelsAndWeightsSum] class.
type INNReduceFeatureChannelsAndWeightsSum interface {
	INNReduceBinary
	DoWeightedSumByNonZeroWeights() bool
}

// A reduction filter that returns the weighted sum of all values for each feature channel in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightssum?language=objc
type NNReduceFeatureChannelsAndWeightsSum struct {
	NNReduceBinary
}

func NNReduceFeatureChannelsAndWeightsSumFrom(ptr unsafe.Pointer) NNReduceFeatureChannelsAndWeightsSum {
	return NNReduceFeatureChannelsAndWeightsSum{
		NNReduceBinary: NNReduceBinaryFrom(ptr),
	}
}

func (n_ NNReduceFeatureChannelsAndWeightsSum) InitWithDevice(device metal.PDevice) NNReduceFeatureChannelsAndWeightsSum {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceFeatureChannelsAndWeightsSum](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightssum/2942562-initwithdevice?language=objc
func NewNNReduceFeatureChannelsAndWeightsSumWithDevice(device metal.PDevice) NNReduceFeatureChannelsAndWeightsSum {
	instance := NNReduceFeatureChannelsAndWeightsSumClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNReduceFeatureChannelsAndWeightsSumClass) Alloc() NNReduceFeatureChannelsAndWeightsSum {
	rv := objc.Call[NNReduceFeatureChannelsAndWeightsSum](nc, objc.Sel("alloc"))
	return rv
}

func NNReduceFeatureChannelsAndWeightsSum_Alloc() NNReduceFeatureChannelsAndWeightsSum {
	return NNReduceFeatureChannelsAndWeightsSumClass.Alloc()
}

func (nc _NNReduceFeatureChannelsAndWeightsSumClass) New() NNReduceFeatureChannelsAndWeightsSum {
	rv := objc.Call[NNReduceFeatureChannelsAndWeightsSum](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReduceFeatureChannelsAndWeightsSum() NNReduceFeatureChannelsAndWeightsSum {
	return NNReduceFeatureChannelsAndWeightsSumClass.New()
}

func (n_ NNReduceFeatureChannelsAndWeightsSum) Init() NNReduceFeatureChannelsAndWeightsSum {
	rv := objc.Call[NNReduceFeatureChannelsAndWeightsSum](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReduceFeatureChannelsAndWeightsSum) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceFeatureChannelsAndWeightsSum {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceFeatureChannelsAndWeightsSum](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReduceFeatureChannelsAndWeightsSum_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceFeatureChannelsAndWeightsSum {
	instance := NNReduceFeatureChannelsAndWeightsSumClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducefeaturechannelsandweightssum/2942543-doweightedsumbynonzeroweights?language=objc
func (n_ NNReduceFeatureChannelsAndWeightsSum) DoWeightedSumByNonZeroWeights() bool {
	rv := objc.Call[bool](n_, objc.Sel("doWeightedSumByNonZeroWeights"))
	return rv
}
