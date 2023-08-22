// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReduceColumnMean] class.
var NNReduceColumnMeanClass = _NNReduceColumnMeanClass{objc.GetClass("MPSNNReduceColumnMean")}

type _NNReduceColumnMeanClass struct {
	objc.Class
}

// An interface definition for the [NNReduceColumnMean] class.
type INNReduceColumnMean interface {
	INNReduceUnary
}

// A reduction filter that returns the mean value for each column in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnmean?language=objc
type NNReduceColumnMean struct {
	NNReduceUnary
}

func NNReduceColumnMeanFrom(ptr unsafe.Pointer) NNReduceColumnMean {
	return NNReduceColumnMean{
		NNReduceUnary: NNReduceUnaryFrom(ptr),
	}
}

func (n_ NNReduceColumnMean) InitWithDevice(device metal.PDevice) NNReduceColumnMean {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceColumnMean](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnmean/2942546-initwithdevice?language=objc
func NewNNReduceColumnMeanWithDevice(device metal.PDevice) NNReduceColumnMean {
	instance := NNReduceColumnMeanClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNReduceColumnMeanClass) Alloc() NNReduceColumnMean {
	rv := objc.Call[NNReduceColumnMean](nc, objc.Sel("alloc"))
	return rv
}

func NNReduceColumnMean_Alloc() NNReduceColumnMean {
	return NNReduceColumnMeanClass.Alloc()
}

func (nc _NNReduceColumnMeanClass) New() NNReduceColumnMean {
	rv := objc.Call[NNReduceColumnMean](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReduceColumnMean() NNReduceColumnMean {
	return NNReduceColumnMeanClass.New()
}

func (n_ NNReduceColumnMean) Init() NNReduceColumnMean {
	rv := objc.Call[NNReduceColumnMean](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReduceColumnMean) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceColumnMean {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceColumnMean](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReduceColumnMean_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceColumnMean {
	instance := NNReduceColumnMeanClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
