// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNPooling] class.
var CNNPoolingClass = _CNNPoolingClass{objc.GetClass("MPSCNNPooling")}

type _CNNPoolingClass struct {
	objc.Class
}

// An interface definition for the [CNNPooling] class.
type ICNNPooling interface {
	ICNNKernel
}

// A pooling kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpooling?language=objc
type CNNPooling struct {
	CNNKernel
}

func CNNPoolingFrom(ptr unsafe.Pointer) CNNPooling {
	return CNNPooling{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (c_ CNNPooling) InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPooling {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPooling](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), po0, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return rv
}

// Initializes a pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpooling/1648902-initwithdevice?language=objc
func NewCNNPoolingWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPooling {
	instance := CNNPoolingClass.Alloc().InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	instance.Autorelease()
	return instance
}

func (cc _CNNPoolingClass) Alloc() CNNPooling {
	rv := objc.Call[CNNPooling](cc, objc.Sel("alloc"))
	return rv
}

func CNNPooling_Alloc() CNNPooling {
	return CNNPoolingClass.Alloc()
}

func (cc _CNNPoolingClass) New() CNNPooling {
	rv := objc.Call[CNNPooling](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNPooling() CNNPooling {
	return CNNPoolingClass.New()
}

func (c_ CNNPooling) Init() CNNPooling {
	rv := objc.Call[CNNPooling](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNPooling) InitWithDevice(device metal.PDevice) CNNPooling {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPooling](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNPoolingWithDevice(device metal.PDevice) CNNPooling {
	instance := CNNPoolingClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNPooling) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNPooling {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPooling](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNPooling_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNPooling {
	instance := CNNPoolingClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
