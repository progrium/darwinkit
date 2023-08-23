// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixLogSoftMaxGradient] class.
var MatrixLogSoftMaxGradientClass = _MatrixLogSoftMaxGradientClass{objc.GetClass("MPSMatrixLogSoftMaxGradient")}

type _MatrixLogSoftMaxGradientClass struct {
	objc.Class
}

// An interface definition for the [MatrixLogSoftMaxGradient] class.
type IMatrixLogSoftMaxGradient interface {
	IMatrixSoftMaxGradient
}

// A logarithmic gradient softmax kernel that operates on matrices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixlogsoftmaxgradient?language=objc
type MatrixLogSoftMaxGradient struct {
	MatrixSoftMaxGradient
}

func MatrixLogSoftMaxGradientFrom(ptr unsafe.Pointer) MatrixLogSoftMaxGradient {
	return MatrixLogSoftMaxGradient{
		MatrixSoftMaxGradient: MatrixSoftMaxGradientFrom(ptr),
	}
}

func (mc _MatrixLogSoftMaxGradientClass) Alloc() MatrixLogSoftMaxGradient {
	rv := objc.Call[MatrixLogSoftMaxGradient](mc, objc.Sel("alloc"))
	return rv
}

func MatrixLogSoftMaxGradient_Alloc() MatrixLogSoftMaxGradient {
	return MatrixLogSoftMaxGradientClass.Alloc()
}

func (mc _MatrixLogSoftMaxGradientClass) New() MatrixLogSoftMaxGradient {
	rv := objc.Call[MatrixLogSoftMaxGradient](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixLogSoftMaxGradient() MatrixLogSoftMaxGradient {
	return MatrixLogSoftMaxGradientClass.New()
}

func (m_ MatrixLogSoftMaxGradient) Init() MatrixLogSoftMaxGradient {
	rv := objc.Call[MatrixLogSoftMaxGradient](m_, objc.Sel("init"))
	return rv
}

func (m_ MatrixLogSoftMaxGradient) InitWithDevice(device metal.PDevice) MatrixLogSoftMaxGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixLogSoftMaxGradient](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966654-initwithdevice?language=objc
func NewMatrixLogSoftMaxGradientWithDevice(device metal.PDevice) MatrixLogSoftMaxGradient {
	instance := MatrixLogSoftMaxGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixLogSoftMaxGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixLogSoftMaxGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixLogSoftMaxGradient](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmaxgradient/2966651-copywithzone?language=objc
func MatrixLogSoftMaxGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixLogSoftMaxGradient {
	instance := MatrixLogSoftMaxGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
