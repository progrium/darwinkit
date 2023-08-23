// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [MatrixLogSoftMax] class.
var MatrixLogSoftMaxClass = _MatrixLogSoftMaxClass{objc.GetClass("MPSMatrixLogSoftMax")}

type _MatrixLogSoftMaxClass struct {
	objc.Class
}

// An interface definition for the [MatrixLogSoftMax] class.
type IMatrixLogSoftMax interface {
	IMatrixSoftMax
}

// A logarithmic softmax kernel that operates on matrices. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixlogsoftmax?language=objc
type MatrixLogSoftMax struct {
	MatrixSoftMax
}

func MatrixLogSoftMaxFrom(ptr unsafe.Pointer) MatrixLogSoftMax {
	return MatrixLogSoftMax{
		MatrixSoftMax: MatrixSoftMaxFrom(ptr),
	}
}

func (mc _MatrixLogSoftMaxClass) Alloc() MatrixLogSoftMax {
	rv := objc.Call[MatrixLogSoftMax](mc, objc.Sel("alloc"))
	return rv
}

func MatrixLogSoftMax_Alloc() MatrixLogSoftMax {
	return MatrixLogSoftMaxClass.Alloc()
}

func (mc _MatrixLogSoftMaxClass) New() MatrixLogSoftMax {
	rv := objc.Call[MatrixLogSoftMax](mc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewMatrixLogSoftMax() MatrixLogSoftMax {
	return MatrixLogSoftMaxClass.New()
}

func (m_ MatrixLogSoftMax) Init() MatrixLogSoftMax {
	rv := objc.Call[MatrixLogSoftMax](m_, objc.Sel("init"))
	return rv
}

func (m_ MatrixLogSoftMax) InitWithDevice(device metal.PDevice) MatrixLogSoftMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixLogSoftMax](m_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935562-initwithdevice?language=objc
func NewMatrixLogSoftMaxWithDevice(device metal.PDevice) MatrixLogSoftMax {
	instance := MatrixLogSoftMaxClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (m_ MatrixLogSoftMax) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixLogSoftMax {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[MatrixLogSoftMax](m_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsmatrixsoftmax/2935566-copywithzone?language=objc
func MatrixLogSoftMax_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) MatrixLogSoftMax {
	instance := MatrixLogSoftMaxClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
