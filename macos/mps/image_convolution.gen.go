// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [ImageConvolution] class.
var ImageConvolutionClass = _ImageConvolutionClass{objc.GetClass("MPSImageConvolution")}

type _ImageConvolutionClass struct {
	objc.Class
}

// An interface definition for the [ImageConvolution] class.
type IImageConvolution interface {
	IUnaryImageKernel
	KernelHeight() uint
	Bias() float64
	SetBias(value float64)
	KernelWidth() uint
}

// A filter that convolves an image with a given kernel of odd width and height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageconvolution?language=objc
type ImageConvolution struct {
	UnaryImageKernel
}

func ImageConvolutionFrom(ptr unsafe.Pointer) ImageConvolution {
	return ImageConvolution{
		UnaryImageKernel: UnaryImageKernelFrom(ptr),
	}
}

func (i_ ImageConvolution) InitWithDeviceKernelWidthKernelHeightWeights(device metal.PDevice, kernelWidth uint, kernelHeight uint, kernelWeights *float64) ImageConvolution {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageConvolution](i_, objc.Sel("initWithDevice:kernelWidth:kernelHeight:weights:"), po0, kernelWidth, kernelHeight, kernelWeights)
	return rv
}

// Initializes a convolution filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageconvolution/1618902-initwithdevice?language=objc
func NewImageConvolutionWithDeviceKernelWidthKernelHeightWeights(device metal.PDevice, kernelWidth uint, kernelHeight uint, kernelWeights *float64) ImageConvolution {
	instance := ImageConvolutionClass.Alloc().InitWithDeviceKernelWidthKernelHeightWeights(device, kernelWidth, kernelHeight, kernelWeights)
	instance.Autorelease()
	return instance
}

func (ic _ImageConvolutionClass) Alloc() ImageConvolution {
	rv := objc.Call[ImageConvolution](ic, objc.Sel("alloc"))
	return rv
}

func ImageConvolution_Alloc() ImageConvolution {
	return ImageConvolutionClass.Alloc()
}

func (ic _ImageConvolutionClass) New() ImageConvolution {
	rv := objc.Call[ImageConvolution](ic, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewImageConvolution() ImageConvolution {
	return ImageConvolutionClass.New()
}

func (i_ ImageConvolution) Init() ImageConvolution {
	rv := objc.Call[ImageConvolution](i_, objc.Sel("init"))
	return rv
}

func (i_ ImageConvolution) InitWithDevice(device metal.PDevice) ImageConvolution {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageConvolution](i_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsunaryimagekernel/2866332-initwithdevice?language=objc
func NewImageConvolutionWithDevice(device metal.PDevice) ImageConvolution {
	instance := ImageConvolutionClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (i_ ImageConvolution) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageConvolution {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[ImageConvolution](i_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func ImageConvolution_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) ImageConvolution {
	instance := ImageConvolutionClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

// The height of the filter window. Must be an odd number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageconvolution/1618842-kernelheight?language=objc
func (i_ ImageConvolution) KernelHeight() uint {
	rv := objc.Call[uint](i_, objc.Sel("kernelHeight"))
	return rv
}

// The value added to a convolved pixel before it is converted back to its intended storage format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageconvolution/1618841-bias?language=objc
func (i_ ImageConvolution) Bias() float64 {
	rv := objc.Call[float64](i_, objc.Sel("bias"))
	return rv
}

// The value added to a convolved pixel before it is converted back to its intended storage format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageconvolution/1618841-bias?language=objc
func (i_ ImageConvolution) SetBias(value float64) {
	objc.Call[objc.Void](i_, objc.Sel("setBias:"), value)
}

// The width of the filter window. Must be an odd number. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsimageconvolution/1618868-kernelwidth?language=objc
func (i_ ImageConvolution) KernelWidth() uint {
	rv := objc.Call[uint](i_, objc.Sel("kernelWidth"))
	return rv
}
