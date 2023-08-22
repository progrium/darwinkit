// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageTent] class.
var ImageTentClass = _ImageTentClass{objc.GetClass("MPSImageTent")}

type _ImageTentClass struct {
	objc.Class
}

// An interface definition for the [ImageTent] class.
type IImageTent interface {
	IImageBox
}

// A filter that convolves an image with a tent filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagetent?language=objc
type ImageTent struct {
	ImageBox
}

func ImageTentFrom(ptr unsafe.Pointer) ImageTent {
	return ImageTent{
		ImageBox: ImageBoxFrom(ptr),
	}
}

func (ic _ImageTentClass) Alloc() ImageTent {
	rv := objc.Call[ImageTent](ic, objc.Sel("alloc"))
	return rv
}

func ImageTent_Alloc() ImageTent {
	return ImageTentClass.Alloc()
}

func (ic _ImageTentClass) New() ImageTent {
	rv := objc.Call[ImageTent](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageTent() ImageTent {
	return ImageTentClass.New()
}

func (i_ ImageTent) Init() ImageTent {
	rv := objc.Call[ImageTent](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageTent) InitWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) ImageTent {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageTent](i_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:"), po0, kernelWidth, kernelHeight)
	return rv
}

// Initializes a box filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagebox/1618789-initwithdevice?language=objc
func NewImageTentWithDeviceKernelWidthKernelHeight(device metal.PDevice, kernelWidth uint, kernelHeight uint) ImageTent {
	instance := ImageTentClass.Alloc().InitWithDeviceKernelWidthKernelHeight(device, kernelWidth, kernelHeight)
	instance.Autorelease()
	return instance
}

func (i_ ImageTent) InitWithDevice(device metal.PDevice) ImageTent {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageTent](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageTentWithDevice(device metal.PDevice) ImageTent {
	instance := ImageTentClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageTent) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageTent {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageTent](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageTent_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageTent {
	instance := ImageTentClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
