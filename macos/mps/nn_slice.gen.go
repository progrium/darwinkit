// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNSlice] class.
var NNSliceClass = _NNSliceClass{objc.GetClass("MPSNNSlice")}

type _NNSliceClass struct {
	objc.Class
}

// An interface definition for the [NNSlice] class.
type INNSlice interface {
	ICNNKernel
}

// A kernel that extracts a slice from an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnslice?language=objc
type NNSlice struct {
	CNNKernel
}

func NNSliceFrom(ptr unsafe.Pointer) NNSlice {
	return NNSlice{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (n_ NNSlice) InitWithDevice(device metal.PDevice) NNSlice {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNSlice](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnslice/2942401-initwithdevice?language=objc
func NewNNSliceWithDevice(device metal.PDevice) NNSlice {
	instance := NNSliceClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (nc _NNSliceClass) Alloc() NNSlice {
	rv := objc.Call[NNSlice](nc, objc.Sel("alloc"))
	return rv
}

func NNSlice_Alloc() NNSlice {
	return NNSliceClass.Alloc()
}

func (nc _NNSliceClass) New() NNSlice {
	rv := objc.Call[NNSlice](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNSlice() NNSlice {
	return NNSliceClass.New()
}

func (n_ NNSlice) Init() NNSlice {
	rv := objc.Call[NNSlice](n_, objc.Sel("init"))
	return rv
}

func (n_ NNSlice) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNSlice {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNSlice](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNSlice_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNSlice {
	instance := NNSliceClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
