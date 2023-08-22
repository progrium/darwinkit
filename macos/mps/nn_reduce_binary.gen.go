// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNReduceBinary] class.
var NNReduceBinaryClass = _NNReduceBinaryClass{objc.GetClass("MPSNNReduceBinary")}

type _NNReduceBinaryClass struct {
	objc.Class
}

// An interface definition for the [NNReduceBinary] class.
type INNReduceBinary interface {
	ICNNBinaryKernel
	PrimarySourceClipRect() metal.Region
	SetPrimarySourceClipRect(value metal.Region)
	SecondarySourceClipRect() metal.Region
	SetSecondarySourceClipRect(value metal.Region)
}

// The base class for binary reduction filters. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary?language=objc
type NNReduceBinary struct {
	CNNBinaryKernel
}

func NNReduceBinaryFrom(ptr unsafe.Pointer) NNReduceBinary {
	return NNReduceBinary{
		CNNBinaryKernel: CNNBinaryKernelFrom(ptr),
	}
}

func (nc _NNReduceBinaryClass) Alloc() NNReduceBinary {
	rv := objc.Call[NNReduceBinary](nc, objc.Sel("alloc"))
	return rv
}

func NNReduceBinary_Alloc() NNReduceBinary {
	return NNReduceBinaryClass.Alloc()
}

func (nc _NNReduceBinaryClass) New() NNReduceBinary {
	rv := objc.Call[NNReduceBinary](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNReduceBinary() NNReduceBinary {
	return NNReduceBinaryClass.New()
}

func (n_ NNReduceBinary) Init() NNReduceBinary {
	rv := objc.Call[NNReduceBinary](n_, objc.Sel("init"))
	return rv
}

func (n_ NNReduceBinary) InitWithDevice(device metal.PDevice) NNReduceBinary {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceBinary](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865629-initwithdevice?language=objc
func NewNNReduceBinaryWithDevice(device metal.PDevice) NNReduceBinary {
	instance := NNReduceBinaryClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (n_ NNReduceBinary) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceBinary {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNReduceBinary](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNReduceBinary_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNReduceBinary {
	instance := NNReduceBinaryClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/2942561-primarysourcecliprect?language=objc
func (n_ NNReduceBinary) PrimarySourceClipRect() metal.Region {
	rv := objc.Call[metal.Region](n_, objc.Sel("primarySourceClipRect"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/2942561-primarysourcecliprect?language=objc
func (n_ NNReduceBinary) SetPrimarySourceClipRect(value metal.Region) {
	objc.Call[objc.Void](n_, objc.Sel("setPrimarySourceClipRect:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/2942556-secondarysourcecliprect?language=objc
func (n_ NNReduceBinary) SecondarySourceClipRect() metal.Region {
	rv := objc.Call[metal.Region](n_, objc.Sel("secondarySourceClipRect"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnreducebinary/2942556-secondarysourcecliprect?language=objc
func (n_ NNReduceBinary) SetSecondarySourceClipRect(value metal.Region) {
	objc.Call[objc.Void](n_, objc.Sel("setSecondarySourceClipRect:"), value)
}
