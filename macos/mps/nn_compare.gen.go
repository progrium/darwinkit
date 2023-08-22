// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNCompare] class.
var NNCompareClass = _NNCompareClass{objc.GetClass("MPSNNCompare")}

type _NNCompareClass struct {
	objc.Class
}

// An interface definition for the [NNCompare] class.
type INNCompare interface {
	ICNNArithmetic
	ComparisonType() NNComparisonType
	SetComparisonType(value NNComparisonType)
	Threshold() float64
	SetThreshold(value float64)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncompare?language=objc
type NNCompare struct {
	CNNArithmetic
}

func NNCompareFrom(ptr unsafe.Pointer) NNCompare {
	return NNCompare{
		CNNArithmetic: CNNArithmeticFrom(ptr),
	}
}

func (n_ NNCompare) InitWithDevice(device metal.PDevice) NNCompare {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNCompare](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncompare/3037375-initwithdevice?language=objc
func NewNNCompareWithDevice(device metal.PDevice) NNCompare {
	instance := NNCompareClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNCompareClass) Alloc() NNCompare {
	rv := objc.Call[NNCompare](nc, objc.Sel("alloc"))
	return rv
}

func NNCompare_Alloc() NNCompare {
	return NNCompareClass.Alloc()
}

func (nc _NNCompareClass) New() NNCompare {
	rv := objc.Call[NNCompare](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNCompare() NNCompare {
	return NNCompareClass.New()
}

func (n_ NNCompare) Init() NNCompare {
	rv := objc.Call[NNCompare](n_, objc.Sel("init"))
	return rv
}

func (n_ NNCompare) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNCompare {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNCompare](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNCompare_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNCompare {
	instance := NNCompareClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncompare/3037374-comparisontype?language=objc
func (n_ NNCompare) ComparisonType() NNComparisonType {
	rv := objc.Call[NNComparisonType](n_, objc.Sel("comparisonType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncompare/3037374-comparisontype?language=objc
func (n_ NNCompare) SetComparisonType(value NNComparisonType) {
	objc.Call[objc.Void](n_, objc.Sel("setComparisonType:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncompare/3037376-threshold?language=objc
func (n_ NNCompare) Threshold() float64 {
	rv := objc.Call[float64](n_, objc.Sel("threshold"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnncompare/3037376-threshold?language=objc
func (n_ NNCompare) SetThreshold(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setThreshold:"), value)
}
