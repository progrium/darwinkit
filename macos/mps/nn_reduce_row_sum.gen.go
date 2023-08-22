// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReduceRowSum] class.
var NNReduceRowSumClass = _NNReduceRowSumClass{objc.GetClass("MPSNNReduceRowSum")}

type _NNReduceRowSumClass struct {
	objc.Class
}

// An interface definition for the [NNReduceRowSum] class.
type INNReduceRowSum interface {
	INNReduceUnary
}

// A reduction filter that returns the sum of all values for each row in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowsum?language=objc
type NNReduceRowSum struct {
	NNReduceUnary
}

func NNReduceRowSumFrom(ptr unsafe.Pointer) NNReduceRowSum {
	return NNReduceRowSum{
		NNReduceUnary: NNReduceUnaryFrom(ptr),
	}
}

func (n_ NNReduceRowSum) InitWithDevice(device metal.PDevice) NNReduceRowSum {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceRowSum](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowsum/2942536-initwithdevice?language=objc
func NewNNReduceRowSumWithDevice(device metal.PDevice) NNReduceRowSum {
	instance := NNReduceRowSumClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNReduceRowSumClass) Alloc() NNReduceRowSum {
	rv := objc.Call[NNReduceRowSum](nc, objc.Sel("alloc"))
	return rv
}

func NNReduceRowSum_Alloc() NNReduceRowSum {
	return NNReduceRowSumClass.Alloc()
}

func (nc _NNReduceRowSumClass) New() NNReduceRowSum {
	rv := objc.Call[NNReduceRowSum](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReduceRowSum() NNReduceRowSum {
	return NNReduceRowSumClass.New()
}

func (n_ NNReduceRowSum) Init() NNReduceRowSum {
	rv := objc.Call[NNReduceRowSum](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReduceRowSum) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceRowSum {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceRowSum](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReduceRowSum_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceRowSum {
	instance := NNReduceRowSumClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
