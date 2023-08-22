// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNPoolingAverageGradient] class.
var CNNPoolingAverageGradientClass = _CNNPoolingAverageGradientClass{objc.GetClass("MPSCNNPoolingAverageGradient")}

type _CNNPoolingAverageGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNPoolingAverageGradient] class.
type ICNNPoolingAverageGradient interface {
	ICNNPoolingGradient
	ZeroPadSizeX() uint
	SetZeroPadSizeX(value uint)
	ZeroPadSizeY() uint
	SetZeroPadSizeY(value uint)
}

// A gradient average pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradient?language=objc
type CNNPoolingAverageGradient struct {
	CNNPoolingGradient
}

func CNNPoolingAverageGradientFrom(ptr unsafe.Pointer) CNNPoolingAverageGradient {
	return CNNPoolingAverageGradient{
		CNNPoolingGradient: CNNPoolingGradientFrom(ptr),
	}
}

func (c_ CNNPoolingAverageGradient) InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingAverageGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingAverageGradient](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), po0, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradient/2942339-initwithdevice?language=objc
func NewCNNPoolingAverageGradientWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingAverageGradient {
	instance := CNNPoolingAverageGradientClass.Alloc().InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	instance.Autorelease()
	return instance
}

func (cc _CNNPoolingAverageGradientClass) Alloc() CNNPoolingAverageGradient {
	rv := objc.Call[CNNPoolingAverageGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNPoolingAverageGradient_Alloc() CNNPoolingAverageGradient {
	return CNNPoolingAverageGradientClass.Alloc()
}

func (cc _CNNPoolingAverageGradientClass) New() CNNPoolingAverageGradient {
	rv := objc.Call[CNNPoolingAverageGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNPoolingAverageGradient() CNNPoolingAverageGradient {
	return CNNPoolingAverageGradientClass.New()
}

func (c_ CNNPoolingAverageGradient) Init() CNNPoolingAverageGradient {
	rv := objc.Call[CNNPoolingAverageGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNPoolingAverageGradient) InitWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNPoolingAverageGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingAverageGradient](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), po0, kernelWidth, kernelHeight)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradient/2942337-initwithdevice?language=objc
func NewCNNPoolingAverageGradientWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNPoolingAverageGradient {
	instance := CNNPoolingAverageGradientClass.Alloc().InitWithDeviceKernelWidthKernelHeight(device, kernelWidth, kernelHeight)
	instance.Autorelease()
	return instance
}

func (c_ CNNPoolingAverageGradient) InitWithDevice(device metal.PDevice) CNNPoolingAverageGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingAverageGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNPoolingAverageGradientWithDevice(device metal.PDevice) CNNPoolingAverageGradient {
	instance := CNNPoolingAverageGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNPoolingAverageGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNPoolingAverageGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingAverageGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNPoolingAverageGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNPoolingAverageGradient {
	instance := CNNPoolingAverageGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradient/2942341-zeropadsizex?language=objc
func (c_ CNNPoolingAverageGradient) ZeroPadSizeX() uint {
	rv := objc.Call[uint](c_, objc.Sel("zeroPadSizeX"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradient/2942341-zeropadsizex?language=objc
func (c_ CNNPoolingAverageGradient) SetZeroPadSizeX(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setZeroPadSizeX:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradient/2942354-zeropadsizey?language=objc
func (c_ CNNPoolingAverageGradient) ZeroPadSizeY() uint {
	rv := objc.Call[uint](c_, objc.Sel("zeroPadSizeY"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingaveragegradient/2942354-zeropadsizey?language=objc
func (c_ CNNPoolingAverageGradient) SetZeroPadSizeY(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setZeroPadSizeY:"), value)
}
