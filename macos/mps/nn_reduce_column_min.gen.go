// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReduceColumnMin] class.
var NNReduceColumnMinClass = _NNReduceColumnMinClass{objc.GetClass("MPSNNReduceColumnMin")}

type _NNReduceColumnMinClass struct {
	objc.Class
}

// An interface definition for the [NNReduceColumnMin] class.
type INNReduceColumnMin interface {
	INNReduceUnary
}

// A reduction filter that returns the minimum value for each column in an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnmin?language=objc
type NNReduceColumnMin struct {
	NNReduceUnary
}

func NNReduceColumnMinFrom(ptr unsafe.Pointer) NNReduceColumnMin {
	return NNReduceColumnMin{
		NNReduceUnary: NNReduceUnaryFrom(ptr),
	}
}

func (n_ NNReduceColumnMin) InitWithDevice(device metal.PDevice) NNReduceColumnMin {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceColumnMin](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducecolumnmin/2942542-initwithdevice?language=objc
func NewNNReduceColumnMinWithDevice(device metal.PDevice) NNReduceColumnMin {
	instance := NNReduceColumnMinClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNReduceColumnMinClass) Alloc() NNReduceColumnMin {
	rv := objc.Call[NNReduceColumnMin](nc, objc.Sel("alloc"))
	return rv
}

func NNReduceColumnMin_Alloc() NNReduceColumnMin {
	return NNReduceColumnMinClass.Alloc()
}

func (nc _NNReduceColumnMinClass) New() NNReduceColumnMin {
	rv := objc.Call[NNReduceColumnMin](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReduceColumnMin() NNReduceColumnMin {
	return NNReduceColumnMinClass.New()
}

func (n_ NNReduceColumnMin) Init() NNReduceColumnMin {
	rv := objc.Call[NNReduceColumnMin](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReduceColumnMin) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceColumnMin {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceColumnMin](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReduceColumnMin_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceColumnMin {
	instance := NNReduceColumnMinClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
