// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixRandomPhilox] class.
var MatrixRandomPhiloxClass = _MatrixRandomPhiloxClass{objc.GetClass("MPSMatrixRandomPhilox")}

type _MatrixRandomPhiloxClass struct {
	objc.Class
}

// An interface definition for the [MatrixRandomPhilox] class.
type IMatrixRandomPhilox interface {
	IMatrixRandom
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomphilox?language=objc
type MatrixRandomPhilox struct {
	MatrixRandom
}

func MatrixRandomPhiloxFrom(ptr unsafe.Pointer) MatrixRandomPhilox {
	return MatrixRandomPhilox{
		MatrixRandom: MatrixRandomFrom(ptr),
	}
}

func (m_ MatrixRandomPhilox) InitWithDevice(device metal.PDevice) MatrixRandomPhilox {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixRandomPhilox](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixrandomphilox/3242871-initwithdevice?language=objc
func NewMatrixRandomPhiloxWithDevice(device metal.PDevice) MatrixRandomPhilox {
	instance := MatrixRandomPhiloxClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (mc _MatrixRandomPhiloxClass) Alloc() MatrixRandomPhilox {
	rv := objc.Call[MatrixRandomPhilox](mc, objc.Sel("alloc"))
	return rv
}

func MatrixRandomPhilox_Alloc() MatrixRandomPhilox {
	return MatrixRandomPhiloxClass.Alloc()
}

func (mc _MatrixRandomPhiloxClass) New() MatrixRandomPhilox {
	rv := objc.Call[MatrixRandomPhilox](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixRandomPhilox() MatrixRandomPhilox {
	return MatrixRandomPhiloxClass.New()
}

func (m_ MatrixRandomPhilox) Init() MatrixRandomPhilox {
	rv := objc.Call[MatrixRandomPhilox](m_, objc.Sel("init"))
	return rv
}

func (m_ MatrixRandomPhilox) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixRandomPhilox {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixRandomPhilox](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func MatrixRandomPhilox_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixRandomPhilox {
	instance := MatrixRandomPhiloxClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
