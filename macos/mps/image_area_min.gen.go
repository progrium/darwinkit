// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageAreaMin] class.
var ImageAreaMinClass = _ImageAreaMinClass{objc.GetClass("MPSImageAreaMin")}

type _ImageAreaMinClass struct {
	objc.Class
}

// An interface definition for the [ImageAreaMin] class.
type IImageAreaMin interface {
	IImageAreaMax
}

// A filter that finds the minimum pixel value in a rectangular region centered around each pixel in the source image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageareamin?language=objc
type ImageAreaMin struct {
	ImageAreaMax
}

func ImageAreaMinFrom(ptr unsafe.Pointer) ImageAreaMin {
	return ImageAreaMin{
		ImageAreaMax: ImageAreaMaxFrom(ptr),
	}
}

func (ic _ImageAreaMinClass) Alloc() ImageAreaMin {
	rv := objc.Call[ImageAreaMin](ic, objc.Sel("alloc"))
	return rv
}

func ImageAreaMin_Alloc() ImageAreaMin {
	return ImageAreaMinClass.Alloc()
}

func (ic _ImageAreaMinClass) New() ImageAreaMin {
	rv := objc.Call[ImageAreaMin](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageAreaMin() ImageAreaMin {
	return ImageAreaMinClass.New()
}

func (i_ ImageAreaMin) Init() ImageAreaMin {
	rv := objc.Call[ImageAreaMin](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageAreaMin) InitWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) ImageAreaMin {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageAreaMin](i_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), po0, kernelWidth, kernelHeight)
	return rv
}

// Initializes the kernel with a specified width and height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageareamax/1618281-initwithdevice?language=objc
func NewImageAreaMinWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) ImageAreaMin {
	instance := ImageAreaMinClass.Alloc().InitWithDeviceKernelWidthKernelHeight(device, kernelWidth, kernelHeight)
	instance.Autorelease()
	return instance
}

func (i_ ImageAreaMin) InitWithDevice(device metal.PDevice) ImageAreaMin {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageAreaMin](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageAreaMinWithDevice(device metal.PDevice) ImageAreaMin {
	instance := ImageAreaMinClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageAreaMin) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageAreaMin {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageAreaMin](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageAreaMin_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageAreaMin {
	instance := ImageAreaMinClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
