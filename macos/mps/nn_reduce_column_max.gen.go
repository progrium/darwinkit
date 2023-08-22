// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReduceColumnMax] class.
var NNReduceColumnMaxClass = _NNReduceColumnMaxClass{objc.GetClass("MPSNNReduceColumnMax")}

type _NNReduceColumnMaxClass struct {
	objc.Class
}

// An interface definition for the [NNReduceColumnMax] class.
type INNReduceColumnMax interface {
	INNReduceUnary
}

// A reduction filter that returns the maximum value for each column in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnmax?language=objc
type NNReduceColumnMax struct {
	NNReduceUnary
}

func NNReduceColumnMaxFrom(ptr unsafe.Pointer) NNReduceColumnMax {
	return NNReduceColumnMax{
		NNReduceUnary: NNReduceUnaryFrom(ptr),
	}
}

func (n_ NNReduceColumnMax) InitWithDevice(device metal.PDevice) NNReduceColumnMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceColumnMax](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnmax/2942541-initwithdevice?language=objc
func NewNNReduceColumnMaxWithDevice(device metal.PDevice) NNReduceColumnMax {
	instance := NNReduceColumnMaxClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNReduceColumnMaxClass) Alloc() NNReduceColumnMax {
	rv := objc.Call[NNReduceColumnMax](nc, objc.Sel("alloc"))
	return rv
}

func NNReduceColumnMax_Alloc() NNReduceColumnMax {
	return NNReduceColumnMaxClass.Alloc()
}

func (nc _NNReduceColumnMaxClass) New() NNReduceColumnMax {
	rv := objc.Call[NNReduceColumnMax](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReduceColumnMax() NNReduceColumnMax {
	return NNReduceColumnMaxClass.New()
}

func (n_ NNReduceColumnMax) Init() NNReduceColumnMax {
	rv := objc.Call[NNReduceColumnMax](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReduceColumnMax) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceColumnMax {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceColumnMax](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReduceColumnMax_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceColumnMax {
	instance := NNReduceColumnMaxClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
