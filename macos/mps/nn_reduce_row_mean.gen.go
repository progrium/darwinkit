// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReduceRowMean] class.
var NNReduceRowMeanClass = _NNReduceRowMeanClass{objc.GetClass("MPSNNReduceRowMean")}

type _NNReduceRowMeanClass struct {
	objc.Class
}

// An interface definition for the [NNReduceRowMean] class.
type INNReduceRowMean interface {
	INNReduceUnary
}

// A reduction filter that returns the mean value for each row in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowmean?language=objc
type NNReduceRowMean struct {
	NNReduceUnary
}

func NNReduceRowMeanFrom(ptr unsafe.Pointer) NNReduceRowMean {
	return NNReduceRowMean{
		NNReduceUnary: NNReduceUnaryFrom(ptr),
	}
}

func (n_ NNReduceRowMean) InitWithDevice(device metal.PDevice) NNReduceRowMean {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceRowMean](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowmean/2942548-initwithdevice?language=objc
func NewNNReduceRowMeanWithDevice(device metal.PDevice) NNReduceRowMean {
	instance := NNReduceRowMeanClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNReduceRowMeanClass) Alloc() NNReduceRowMean {
	rv := objc.Call[NNReduceRowMean](nc, objc.Sel("alloc"))
	return rv
}

func NNReduceRowMean_Alloc() NNReduceRowMean {
	return NNReduceRowMeanClass.Alloc()
}

func (nc _NNReduceRowMeanClass) New() NNReduceRowMean {
	rv := objc.Call[NNReduceRowMean](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReduceRowMean() NNReduceRowMean {
	return NNReduceRowMeanClass.New()
}

func (n_ NNReduceRowMean) Init() NNReduceRowMean {
	rv := objc.Call[NNReduceRowMean](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReduceRowMean) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceRowMean {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceRowMean](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReduceRowMean_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceRowMean {
	instance := NNReduceRowMeanClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
