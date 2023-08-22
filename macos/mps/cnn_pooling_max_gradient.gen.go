// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNPoolingMaxGradient] class.
var CNNPoolingMaxGradientClass = _CNNPoolingMaxGradientClass{objc.GetClass("MPSCNNPoolingMaxGradient")}

type _CNNPoolingMaxGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNPoolingMaxGradient] class.
type ICNNPoolingMaxGradient interface {
	ICNNPoolingGradient
}

// A gradient max pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingmaxgradient?language=objc
type CNNPoolingMaxGradient struct {
	CNNPoolingGradient
}

func CNNPoolingMaxGradientFrom(ptr unsafe.Pointer) CNNPoolingMaxGradient {
	return CNNPoolingMaxGradient{
		CNNPoolingGradient: CNNPoolingGradientFrom(ptr),
	}
}

func (c_ CNNPoolingMaxGradient) InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingMaxGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingMaxGradient](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), po0, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingmaxgradient/2942348-initwithdevice?language=objc
func NewCNNPoolingMaxGradientWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingMaxGradient {
	instance := CNNPoolingMaxGradientClass.Alloc().InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	instance.Autorelease()
	return instance
}

func (cc _CNNPoolingMaxGradientClass) Alloc() CNNPoolingMaxGradient {
	rv := objc.Call[CNNPoolingMaxGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNPoolingMaxGradient_Alloc() CNNPoolingMaxGradient {
	return CNNPoolingMaxGradientClass.Alloc()
}

func (cc _CNNPoolingMaxGradientClass) New() CNNPoolingMaxGradient {
	rv := objc.Call[CNNPoolingMaxGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNPoolingMaxGradient() CNNPoolingMaxGradient {
	return CNNPoolingMaxGradientClass.New()
}

func (c_ CNNPoolingMaxGradient) Init() CNNPoolingMaxGradient {
	rv := objc.Call[CNNPoolingMaxGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNPoolingMaxGradient) InitWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNPoolingMaxGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingMaxGradient](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), po0, kernelWidth, kernelHeight)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradient/2942337-initwithdevice?language=objc
func NewCNNPoolingMaxGradientWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNPoolingMaxGradient {
	instance := CNNPoolingMaxGradientClass.Alloc().InitWithDeviceKernelWidthKernelHeight(device, kernelWidth, kernelHeight)
	instance.Autorelease()
	return instance
}

func (c_ CNNPoolingMaxGradient) InitWithDevice(device metal.PDevice) CNNPoolingMaxGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingMaxGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNPoolingMaxGradientWithDevice(device metal.PDevice) CNNPoolingMaxGradient {
	instance := CNNPoolingMaxGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNPoolingMaxGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNPoolingMaxGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingMaxGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNPoolingMaxGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNPoolingMaxGradient {
	instance := CNNPoolingMaxGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
