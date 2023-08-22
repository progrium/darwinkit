// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageErode] class.
var ImageErodeClass = _ImageErodeClass{objc.GetClass("MPSImageErode")}

type _ImageErodeClass struct {
	objc.Class
}

// An interface definition for the [ImageErode] class.
type IImageErode interface {
	IImageDilate
}

// A filter that finds the minimum pixel value in a rectangular region by applying an erosion function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageerode?language=objc
type ImageErode struct {
	ImageDilate
}

func ImageErodeFrom(ptr unsafe.Pointer) ImageErode {
	return ImageErode{
		ImageDilate: ImageDilateFrom(ptr),
	}
}

func (ic _ImageErodeClass) Alloc() ImageErode {
	rv := objc.Call[ImageErode](ic, objc.Sel("alloc"))
	return rv
}

func ImageErode_Alloc() ImageErode {
	return ImageErodeClass.Alloc()
}

func (ic _ImageErodeClass) New() ImageErode {
	rv := objc.Call[ImageErode](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageErode() ImageErode {
	return ImageErodeClass.New()
}

func (i_ ImageErode) Init() ImageErode {
	rv := objc.Call[ImageErode](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageErode) InitWithDeviceKernelWidthKernelHeightValues(device metal.PDevice, kernelWidth uint, kernelHeight uint, values *float64) ImageErode {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageErode](i_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:values:"), po0, kernelWidth, kernelHeight, values)
	return rv
}

// Initializes the kernel with a specified width, height, and weight values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedilate/1618285-initwithdevice?language=objc
func NewImageErodeWithDeviceKernelWidthKernelHeightValues(device metal.PDevice, kernelWidth uint, kernelHeight uint, values *float64) ImageErode {
	instance := ImageErodeClass.Alloc().InitWithDeviceKernelWidthKernelHeightValues(device, kernelWidth, kernelHeight, values)
	instance.Autorelease()
	return instance
}

func (i_ ImageErode) InitWithDevice(device metal.PDevice) ImageErode {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageErode](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageErodeWithDevice(device metal.PDevice) ImageErode {
	instance := ImageErodeClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageErode) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageErode {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageErode](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageErode_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageErode {
	instance := ImageErodeClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
