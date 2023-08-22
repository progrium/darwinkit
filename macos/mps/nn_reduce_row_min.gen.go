// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReduceRowMin] class.
var NNReduceRowMinClass = _NNReduceRowMinClass{objc.GetClass("MPSNNReduceRowMin")}

type _NNReduceRowMinClass struct {
	objc.Class
}

// An interface definition for the [NNReduceRowMin] class.
type INNReduceRowMin interface {
	INNReduceUnary
}

// A reduction filter that returns the minimum value for each row in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowmin?language=objc
type NNReduceRowMin struct {
	NNReduceUnary
}

func NNReduceRowMinFrom(ptr unsafe.Pointer) NNReduceRowMin {
	return NNReduceRowMin{
		NNReduceUnary: NNReduceUnaryFrom(ptr),
	}
}

func (n_ NNReduceRowMin) InitWithDevice(device metal.PDevice) NNReduceRowMin {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceRowMin](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducerowmin/2942555-initwithdevice?language=objc
func NewNNReduceRowMinWithDevice(device metal.PDevice) NNReduceRowMin {
	instance := NNReduceRowMinClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNReduceRowMinClass) Alloc() NNReduceRowMin {
	rv := objc.Call[NNReduceRowMin](nc, objc.Sel("alloc"))
	return rv
}

func NNReduceRowMin_Alloc() NNReduceRowMin {
	return NNReduceRowMinClass.Alloc()
}

func (nc _NNReduceRowMinClass) New() NNReduceRowMin {
	rv := objc.Call[NNReduceRowMin](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReduceRowMin() NNReduceRowMin {
	return NNReduceRowMinClass.New()
}

func (n_ NNReduceRowMin) Init() NNReduceRowMin {
	rv := objc.Call[NNReduceRowMin](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReduceRowMin) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceRowMin {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceRowMin](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReduceRowMin_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceRowMin {
	instance := NNReduceRowMinClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
