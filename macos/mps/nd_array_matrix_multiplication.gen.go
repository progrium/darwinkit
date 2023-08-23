// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NDArrayMatrixMultiplication] class.
var NDArrayMatrixMultiplicationClass = _NDArrayMatrixMultiplicationClass{objc.GetClass("MPSNDArrayMatrixMultiplication")}

type _NDArrayMatrixMultiplicationClass struct {
	objc.Class
}

// An interface definition for the [NDArrayMatrixMultiplication] class.
type INDArrayMatrixMultiplication interface {
	INDArrayMultiaryKernel
	Beta() float64
	SetBeta(value float64)
	Alpha() float64
	SetAlpha(value float64)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymatrixmultiplication?language=objc
type NDArrayMatrixMultiplication struct {
	NDArrayMultiaryKernel
}

func NDArrayMatrixMultiplicationFrom(ptr unsafe.Pointer) NDArrayMatrixMultiplication {
	return NDArrayMatrixMultiplication{
		NDArrayMultiaryKernel: NDArrayMultiaryKernelFrom(ptr),
	}
}

func (nc _NDArrayMatrixMultiplicationClass) Alloc() NDArrayMatrixMultiplication {
	rv := objc.Call[NDArrayMatrixMultiplication](nc, objc.Sel("alloc"))
	return rv
}

func NDArrayMatrixMultiplication_Alloc() NDArrayMatrixMultiplication {
	return NDArrayMatrixMultiplicationClass.Alloc()
}

func (nc _NDArrayMatrixMultiplicationClass) New() NDArrayMatrixMultiplication {
	rv := objc.Call[NDArrayMatrixMultiplication](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNDArrayMatrixMultiplication() NDArrayMatrixMultiplication {
	return NDArrayMatrixMultiplicationClass.New()
}

func (n_ NDArrayMatrixMultiplication) Init() NDArrayMatrixMultiplication {
	rv := objc.Call[NDArrayMatrixMultiplication](n_, objc.Sel("init"))
	return rv
}

func (n_ NDArrayMatrixMultiplication) InitWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayMatrixMultiplication {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayMatrixMultiplication](n_, objc.Sel("initWithDevice:sourceCount:"), po0, count)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarykernel/3175010-initwithdevice?language=objc
func NewNDArrayMatrixMultiplicationWithDeviceSourceCount(device metal.PDevice, count uint) NDArrayMatrixMultiplication {
	instance := NDArrayMatrixMultiplicationClass.Alloc().InitWithDeviceSourceCount(device, count)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayMatrixMultiplication) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayMatrixMultiplication {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayMatrixMultiplication](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymultiarybase/3131734-copywithzone?language=objc
func NDArrayMatrixMultiplication_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NDArrayMatrixMultiplication {
	instance := NDArrayMatrixMultiplicationClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (n_ NDArrayMatrixMultiplication) InitWithDevice(device metal.PDevice) NDArrayMatrixMultiplication {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NDArrayMatrixMultiplication](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

// Initializes a new kernel object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618763-initwithdevice?language=objc
func NewNDArrayMatrixMultiplicationWithDevice(device metal.PDevice) NDArrayMatrixMultiplication {
	instance := NDArrayMatrixMultiplicationClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymatrixmultiplication/3131761-beta?language=objc
func (n_ NDArrayMatrixMultiplication) Beta() float64 {
	rv := objc.Call[float64](n_, objc.Sel("beta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymatrixmultiplication/3131761-beta?language=objc
func (n_ NDArrayMatrixMultiplication) SetBeta(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setBeta:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymatrixmultiplication/3131760-alpha?language=objc
func (n_ NDArrayMatrixMultiplication) Alpha() float64 {
	rv := objc.Call[float64](n_, objc.Sel("alpha"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsndarraymatrixmultiplication/3131760-alpha?language=objc
func (n_ NDArrayMatrixMultiplication) SetAlpha(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setAlpha:"), value)
}
