// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageTranspose] class.
var ImageTransposeClass = _ImageTransposeClass{objc.GetClass("MPSImageTranspose")}

type _ImageTransposeClass struct {
	objc.Class
}

// An interface definition for the [ImageTranspose] class.
type IImageTranspose interface {
	IUnaryImageKernel
}

// A filter that transposes an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagetranspose?language=objc
type ImageTranspose struct {
	UnaryImageKernel
}

func ImageTransposeFrom(ptr unsafe.Pointer) ImageTranspose {
	return ImageTranspose{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (ic _ImageTransposeClass) Alloc() ImageTranspose {
	rv := objc.Call[ImageTranspose](ic, objc.Sel("alloc"))
	return rv
}

func ImageTranspose_Alloc() ImageTranspose {
	return ImageTransposeClass.Alloc()
}

func (ic _ImageTransposeClass) New() ImageTranspose {
	rv := objc.Call[ImageTranspose](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageTranspose() ImageTranspose {
	return ImageTransposeClass.New()
}

func (i_ ImageTranspose) Init() ImageTranspose {
	rv := objc.Call[ImageTranspose](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageTranspose) InitWithDevice(device metal.PDevice) ImageTranspose {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageTranspose](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageTransposeWithDevice(device metal.PDevice) ImageTranspose {
	instance := ImageTransposeClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageTranspose) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageTranspose {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageTranspose](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageTranspose_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageTranspose {
	instance := ImageTransposeClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
