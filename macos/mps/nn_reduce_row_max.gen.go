// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReduceRowMax] class.
var NNReduceRowMaxClass = _NNReduceRowMaxClass{objc.GetClass("MPSNNReduceRowMax")}

type _NNReduceRowMaxClass struct {
	objc.Class
}

// An interface definition for the [NNReduceRowMax] class.
type INNReduceRowMax interface {
	INNReduceUnary
}

// A reduction filter that returns the maximum value for each row in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowmax?language=objc
type NNReduceRowMax struct {
	NNReduceUnary
}

func NNReduceRowMaxFrom(ptr unsafe.Pointer) NNReduceRowMax {
	return NNReduceRowMax{
		NNReduceUnary: NNReduceUnaryFrom(ptr),
	}
}

func (n_ NNReduceRowMax) InitWithDevice(device metal.PDevice) NNReduceRowMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceRowMax](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowmax/2942559-initwithdevice?language=objc
func NewNNReduceRowMaxWithDevice(device metal.PDevice) NNReduceRowMax {
	instance := NNReduceRowMaxClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNReduceRowMaxClass) Alloc() NNReduceRowMax {
	rv := objc.Call[NNReduceRowMax](nc, objc.Sel("alloc"))
	return rv
}

func NNReduceRowMax_Alloc() NNReduceRowMax {
	return NNReduceRowMaxClass.Alloc()
}

func (nc _NNReduceRowMaxClass) New() NNReduceRowMax {
	rv := objc.Call[NNReduceRowMax](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReduceRowMax() NNReduceRowMax {
	return NNReduceRowMaxClass.New()
}

func (n_ NNReduceRowMax) Init() NNReduceRowMax {
	rv := objc.Call[NNReduceRowMax](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReduceRowMax) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceRowMax {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceRowMax](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReduceRowMax_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceRowMax {
	instance := NNReduceRowMaxClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
