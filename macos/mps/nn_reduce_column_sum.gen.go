// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReduceColumnSum] class.
var NNReduceColumnSumClass = _NNReduceColumnSumClass{objc.GetClass("MPSNNReduceColumnSum")}

type _NNReduceColumnSumClass struct {
	objc.Class
}

// An interface definition for the [NNReduceColumnSum] class.
type INNReduceColumnSum interface {
	INNReduceUnary
}

// A reduction filter that returns the sum of all values for each column in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnsum?language=objc
type NNReduceColumnSum struct {
	NNReduceUnary
}

func NNReduceColumnSumFrom(ptr unsafe.Pointer) NNReduceColumnSum {
	return NNReduceColumnSum{
		NNReduceUnary: NNReduceUnaryFrom(ptr),
	}
}

func (n_ NNReduceColumnSum) InitWithDevice(device metal.PDevice) NNReduceColumnSum {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceColumnSum](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnsum/2942540-initwithdevice?language=objc
func NewNNReduceColumnSumWithDevice(device metal.PDevice) NNReduceColumnSum {
	instance := NNReduceColumnSumClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNReduceColumnSumClass) Alloc() NNReduceColumnSum {
	rv := objc.Call[NNReduceColumnSum](nc, objc.Sel("alloc"))
	return rv
}

func NNReduceColumnSum_Alloc() NNReduceColumnSum {
	return NNReduceColumnSumClass.Alloc()
}

func (nc _NNReduceColumnSumClass) New() NNReduceColumnSum {
	rv := objc.Call[NNReduceColumnSum](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReduceColumnSum() NNReduceColumnSum {
	return NNReduceColumnSumClass.New()
}

func (n_ NNReduceColumnSum) Init() NNReduceColumnSum {
	rv := objc.Call[NNReduceColumnSum](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReduceColumnSum) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceColumnSum {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceColumnSum](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReduceColumnSum_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceColumnSum {
	instance := NNReduceColumnSumClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
