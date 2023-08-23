// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNPoolingAverage] class.
var CNNPoolingAverageClass = _CNNPoolingAverageClass{objc.GetClass("MPSCNNPoolingAverage")}

type _CNNPoolingAverageClass struct {
	objc.Class
}

// An interface definition for the [CNNPoolingAverage] class.
type ICNNPoolingAverage interface {
	ICNNPooling
	ZeroPadSizeX() uint
	SetZeroPadSizeX(value uint)
	ZeroPadSizeY() uint
	SetZeroPadSizeY(value uint)
}

// An average pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaverage?language=objc
type CNNPoolingAverage struct {
	CNNPooling
}

func CNNPoolingAverageFrom(ptr unsafe.Pointer) CNNPoolingAverage {
	return CNNPoolingAverage{
		CNNPooling: CNNPoolingFrom(ptr),
	}
}

func (c_ CNNPoolingAverage) InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingAverage {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingAverage](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), po0, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return rv
}

// Initializes an average pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaverage/2875216-initwithdevice?language=objc
func NewCNNPoolingAverageWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingAverage {
	instance := CNNPoolingAverageClass.Alloc().InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	instance.Autorelease()
	return instance
}

func (cc _CNNPoolingAverageClass) Alloc() CNNPoolingAverage {
	rv := objc.Call[CNNPoolingAverage](cc, objc.Sel("alloc"))
	return rv
}

func CNNPoolingAverage_Alloc() CNNPoolingAverage {
	return CNNPoolingAverageClass.Alloc()
}

func (cc _CNNPoolingAverageClass) New() CNNPoolingAverage {
	rv := objc.Call[CNNPoolingAverage](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNPoolingAverage() CNNPoolingAverage {
	return CNNPoolingAverageClass.New()
}

func (c_ CNNPoolingAverage) Init() CNNPoolingAverage {
	rv := objc.Call[CNNPoolingAverage](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNPoolingAverage) InitWithDevice(device metal.PDevice) CNNPoolingAverage {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingAverage](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNPoolingAverageWithDevice(device metal.PDevice) CNNPoolingAverage {
	instance := CNNPoolingAverageClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNPoolingAverage) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNPoolingAverage {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingAverage](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNPoolingAverage_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNPoolingAverage {
	instance := CNNPoolingAverageClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaverage/2875207-zeropadsizex?language=objc
func (c_ CNNPoolingAverage) ZeroPadSizeX() uint {
	rv := objc.Call[uint](c_, objc.Sel("zeroPadSizeX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaverage/2875207-zeropadsizex?language=objc
func (c_ CNNPoolingAverage) SetZeroPadSizeX(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setZeroPadSizeX:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaverage/2875221-zeropadsizey?language=objc
func (c_ CNNPoolingAverage) ZeroPadSizeY() uint {
	rv := objc.Call[uint](c_, objc.Sel("zeroPadSizeY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaverage/2875221-zeropadsizey?language=objc
func (c_ CNNPoolingAverage) SetZeroPadSizeY(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setZeroPadSizeY:"), value)
}
