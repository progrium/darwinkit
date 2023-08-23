// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNPoolingMax] class.
var CNNPoolingMaxClass = _CNNPoolingMaxClass{objc.GetClass("MPSCNNPoolingMax")}

type _CNNPoolingMaxClass struct {
	objc.Class
}

// An interface definition for the [CNNPoolingMax] class.
type ICNNPoolingMax interface {
	ICNNPooling
}

// A max pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingmax?language=objc
type CNNPoolingMax struct {
	CNNPooling
}

func CNNPoolingMaxFrom(ptr unsafe.Pointer) CNNPoolingMax {
	return CNNPoolingMax{
		CNNPooling: CNNPoolingFrom(ptr),
	}
}

func (c_ CNNPoolingMax) InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingMax](c_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:strideInPixelsX:strideInPixelsY:"), po0, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	return rv
}

// Initializes a max pooling filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnpoolingmax/2875151-initwithdevice?language=objc
func NewCNNPoolingMaxWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device metal.PDevice, kernelWidth uint, kernelHeight uint, strideInPixelsX uint, strideInPixelsY uint) CNNPoolingMax {
	instance := CNNPoolingMaxClass.Alloc().InitWithDeviceKernelWidthKernelHeightStrideInPixelsXStrideInPixelsY(device, kernelWidth, kernelHeight, strideInPixelsX, strideInPixelsY)
	instance.Autorelease()
	return instance
}

func (cc _CNNPoolingMaxClass) Alloc() CNNPoolingMax {
	rv := objc.Call[CNNPoolingMax](cc, objc.Sel("alloc"))
	return rv
}

func CNNPoolingMax_Alloc() CNNPoolingMax {
	return CNNPoolingMaxClass.Alloc()
}

func (cc _CNNPoolingMaxClass) New() CNNPoolingMax {
	rv := objc.Call[CNNPoolingMax](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNPoolingMax() CNNPoolingMax {
	return CNNPoolingMaxClass.New()
}

func (c_ CNNPoolingMax) Init() CNNPoolingMax {
	rv := objc.Call[CNNPoolingMax](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNPoolingMax) InitWithDevice(device metal.PDevice) CNNPoolingMax {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingMax](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNPoolingMaxWithDevice(device metal.PDevice) CNNPoolingMax {
	instance := CNNPoolingMaxClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNPoolingMax) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNPoolingMax {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNPoolingMax](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNPoolingMax_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNPoolingMax {
	instance := CNNPoolingMaxClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
