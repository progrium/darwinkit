// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageAdd] class.
var ImageAddClass = _ImageAddClass{objc.GetClass("MPSImageAdd")}

type _ImageAddClass struct {
	objc.Class
}

// An interface definition for the [ImageAdd] class.
type IImageAdd interface {
	IImageArithmetic
}

// A filter that returns the element-wise sum of its two input images. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageadd?language=objc
type ImageAdd struct {
	ImageArithmetic
}

func ImageAddFrom(ptr unsafe.Pointer) ImageAdd {
	return ImageAdd{
		ImageArithmetic: ImageArithmeticFrom(ptr),
	}
}

func (i_ ImageAdd) InitWithDevice(device metal.PDevice) ImageAdd {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageAdd](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageadd/2866610-initwithdevice?language=objc
func NewImageAddWithDevice(device metal.PDevice) ImageAdd {
	instance := ImageAddClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (ic _ImageAddClass) Alloc() ImageAdd {
	rv := objc.Call[ImageAdd](ic, objc.Sel("alloc"))
	return rv
}

func ImageAdd_Alloc() ImageAdd {
	return ImageAddClass.Alloc()
}

func (ic _ImageAddClass) New() ImageAdd {
	rv := objc.Call[ImageAdd](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageAdd() ImageAdd {
	return ImageAddClass.New()
}

func (i_ ImageAdd) Init() ImageAdd {
	rv := objc.Call[ImageAdd](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageAdd) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageAdd {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageAdd](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageAdd_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageAdd {
	instance := ImageAddClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
