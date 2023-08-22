// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageLaplacian] class.
var ImageLaplacianClass = _ImageLaplacianClass{objc.GetClass("MPSImageLaplacian")}

type _ImageLaplacianClass struct {
	objc.Class
}

// An interface definition for the [ImageLaplacian] class.
type IImageLaplacian interface {
	IUnaryImageKernel
	Bias() float64
	SetBias(value float64)
}

// An optimized Laplacian filter, provided for ease of use. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacian?language=objc
type ImageLaplacian struct {
	UnaryImageKernel
}

func ImageLaplacianFrom(ptr unsafe.Pointer) ImageLaplacian {
	return ImageLaplacian{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (ic _ImageLaplacianClass) Alloc() ImageLaplacian {
	rv := objc.Call[ImageLaplacian](ic, objc.Sel("alloc"))
	return rv
}

func ImageLaplacian_Alloc() ImageLaplacian {
	return ImageLaplacianClass.Alloc()
}

func (ic _ImageLaplacianClass) New() ImageLaplacian {
	rv := objc.Call[ImageLaplacian](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageLaplacian() ImageLaplacian {
	return ImageLaplacianClass.New()
}

func (i_ ImageLaplacian) Init() ImageLaplacian {
	rv := objc.Call[ImageLaplacian](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageLaplacian) InitWithDevice(device metal.PDevice) ImageLaplacian {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageLaplacian](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageLaplacianWithDevice(device metal.PDevice) ImageLaplacian {
	instance := ImageLaplacianClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageLaplacian) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageLaplacian {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageLaplacian](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageLaplacian_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageLaplacian {
	instance := ImageLaplacianClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// The value added to a convolved pixel before it is converted back to its intended storage format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacian/1648929-bias?language=objc
func (i_ ImageLaplacian) Bias() float64 {
	rv := objc.Call[float64](i_, objc.Sel("bias"))
	return rv
}

// The value added to a convolved pixel before it is converted back to its intended storage format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimagelaplacian/1648929-bias?language=objc
func (i_ ImageLaplacian) SetBias(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setBias:"), value)
}
