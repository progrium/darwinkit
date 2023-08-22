// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageDilate] class.
var ImageDilateClass = _ImageDilateClass{objc.GetClass("MPSImageDilate")}

type _ImageDilateClass struct {
	objc.Class
}

// An interface definition for the [ImageDilate] class.
type IImageDilate interface {
	IUnaryImageKernel
	KernelHeight() uint
	KernelWidth() uint
}

// A filter that finds the maximum pixel value in a rectangular region by applying a dilation function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedilate?language=objc
type ImageDilate struct {
	UnaryImageKernel
}

func ImageDilateFrom(ptr unsafe.Pointer) ImageDilate {
	return ImageDilate{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageDilate) InitWithDeviceKernelWidthKernelHeightValues(device metal.PDevice, kernelWidth uint, kernelHeight uint, values *float64) ImageDilate {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageDilate](i_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:values:"), po0, kernelWidth, kernelHeight, values)
	return rv
}

// Initializes the kernel with a specified width, height, and weight values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedilate/1618285-initwithdevice?language=objc
func NewImageDilateWithDeviceKernelWidthKernelHeightValues(device metal.PDevice, kernelWidth uint, kernelHeight uint, values *float64) ImageDilate {
	instance := ImageDilateClass.Alloc().InitWithDeviceKernelWidthKernelHeightValues(device, kernelWidth, kernelHeight, values)
	instance.Autorelease()
	return instance
}

func (ic _ImageDilateClass) Alloc() ImageDilate {
	rv := objc.Call[ImageDilate](ic, objc.Sel("alloc"))
	return rv
}

func ImageDilate_Alloc() ImageDilate {
	return ImageDilateClass.Alloc()
}

func (ic _ImageDilateClass) New() ImageDilate {
	rv := objc.Call[ImageDilate](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageDilate() ImageDilate {
	return ImageDilateClass.New()
}

func (i_ ImageDilate) Init() ImageDilate {
	rv := objc.Call[ImageDilate](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageDilate) InitWithDevice(device metal.PDevice) ImageDilate {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageDilate](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageDilateWithDevice(device metal.PDevice) ImageDilate {
	instance := ImageDilateClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageDilate) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageDilate {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageDilate](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageDilate_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageDilate {
	instance := ImageDilateClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// The height of the filter window. which must be an odd number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedilate/1618280-kernelheight?language=objc
func (i_ ImageDilate) KernelHeight() uint {
	rv := objc.Call[uint](i_, objc.Sel("kernelHeight"))
	return rv
}

// The width of the filter window which must be an odd number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagedilate/1618279-kernelwidth?language=objc
func (i_ ImageDilate) KernelWidth() uint {
	rv := objc.Call[uint](i_, objc.Sel("kernelWidth"))
	return rv
}
