// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNPoolingL2NormGradient] class.
var CNNPoolingL2NormGradientClass = _CNNPoolingL2NormGradientClass{objc.GetClass("MPSCNNPoolingL2NormGradient")}

type _CNNPoolingL2NormGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNPoolingL2NormGradient] class.
type ICNNPoolingL2NormGradient interface {
	ICNNPoolingGradient
}

// A gradient L2-norm pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingl2normgradient?language=objc
type CNNPoolingL2NormGradient struct {
	CNNPoolingGradient
}

func CNNPoolingL2NormGradientFrom(ptr unsafe.Pointer) CNNPoolingL2NormGradient {
	return CNNPoolingL2NormGradient{
		CNNPoolingGradient: CNNPoolingGradientFrom(ptr),
	}
}

func (c_ CNNPoolingL2NormGradient) InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingL2NormGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingL2NormGradient](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), po0, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingl2normgradient/2942355-initwithdevice?language=objc
func NewCNNPoolingL2NormGradientWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingL2NormGradient {
	instance := CNNPoolingL2NormGradientClass.Alloc().InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	instance.Autorelease()
	return instance
}

func (cc _CNNPoolingL2NormGradientClass) Alloc() CNNPoolingL2NormGradient {
	rv := objc.Call[CNNPoolingL2NormGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNPoolingL2NormGradient_Alloc() CNNPoolingL2NormGradient {
	return CNNPoolingL2NormGradientClass.Alloc()
}

func (cc _CNNPoolingL2NormGradientClass) New() CNNPoolingL2NormGradient {
	rv := objc.Call[CNNPoolingL2NormGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNPoolingL2NormGradient() CNNPoolingL2NormGradient {
	return CNNPoolingL2NormGradientClass.New()
}

func (c_ CNNPoolingL2NormGradient) Init() CNNPoolingL2NormGradient {
	rv := objc.Call[CNNPoolingL2NormGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNPoolingL2NormGradient) InitWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNPoolingL2NormGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingL2NormGradient](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), po0, kernelWidth, kernelHeight)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolinggradient/2942337-initwithdevice?language=objc
func NewCNNPoolingL2NormGradientWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) CNNPoolingL2NormGradient {
	instance := CNNPoolingL2NormGradientClass.Alloc().InitWithDeviceKernelWidthKernelHeight(device, kernelWidth, kernelHeight)
	instance.Autorelease()
	return instance
}

func (c_ CNNPoolingL2NormGradient) InitWithDevice(device metal.PDevice) CNNPoolingL2NormGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingL2NormGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNPoolingL2NormGradientWithDevice(device metal.PDevice) CNNPoolingL2NormGradient {
	instance := CNNPoolingL2NormGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNPoolingL2NormGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNPoolingL2NormGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingL2NormGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNPoolingL2NormGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNPoolingL2NormGradient {
	instance := CNNPoolingL2NormGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
