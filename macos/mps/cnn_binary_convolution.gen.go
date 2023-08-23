// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNBinaryConvolution] class.
var CNNBinaryConvolutionClass = _CNNBinaryConvolutionClass{objc.GetClass("MPSCNNBinaryConvolution")}

type _CNNBinaryConvolutionClass struct {
	objc.Class
}

// An interface definition for the [CNNBinaryConvolution] class.
type ICNNBinaryConvolution interface {
	ICNNKernel
	OutputFeatureChannels() uint
	InputFeatureChannels() uint
}

// A convolution kernel with binary weights and an input image using binary approximations. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolution?language=objc
type CNNBinaryConvolution struct {
	CNNKernel
}

func CNNBinaryConvolutionFrom(ptr unsafe.Pointer) CNNBinaryConvolution {
	return CNNBinaryConvolution{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (c_ CNNBinaryConvolution) InitWithDeviceConvolutionDataScaleValueTypeFlags(device metal.PDevice, convolutionData PCNNConvolutionDataSource, scaleValue float64, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryConvolution {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", convolutionData)
	rv := objc.Call[CNNBinaryConvolution](c_, objc.Sel("initWithDevice:convolutionData:scaleValue:type:flags:"), po0, po1, scaleValue, type_, flags)
	return rv
}

// Initializes a binary convolution kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolution/2866981-initwithdevice?language=objc
func NewCNNBinaryConvolutionWithDeviceConvolutionDataScaleValueTypeFlags(device metal.PDevice, convolutionData PCNNConvolutionDataSource, scaleValue float64, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryConvolution {
	instance := CNNBinaryConvolutionClass.Alloc().InitWithDeviceConvolutionDataScaleValueTypeFlags(device, convolutionData, scaleValue, type_, flags)
	instance.Autorelease()
	return instance
}

func (cc _CNNBinaryConvolutionClass) Alloc() CNNBinaryConvolution {
	rv := objc.Call[CNNBinaryConvolution](cc, objc.Sel("alloc"))
	return rv
}

func CNNBinaryConvolution_Alloc() CNNBinaryConvolution {
	return CNNBinaryConvolutionClass.Alloc()
}

func (cc _CNNBinaryConvolutionClass) New() CNNBinaryConvolution {
	rv := objc.Call[CNNBinaryConvolution](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNBinaryConvolution() CNNBinaryConvolution {
	return CNNBinaryConvolutionClass.New()
}

func (c_ CNNBinaryConvolution) Init() CNNBinaryConvolution {
	rv := objc.Call[CNNBinaryConvolution](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNBinaryConvolution) InitWithDevice(device metal.PDevice) CNNBinaryConvolution {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBinaryConvolution](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNBinaryConvolutionWithDevice(device metal.PDevice) CNNBinaryConvolution {
	instance := CNNBinaryConvolutionClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNBinaryConvolution) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNBinaryConvolution {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBinaryConvolution](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNBinaryConvolution_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNBinaryConvolution {
	instance := CNNBinaryConvolutionClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolution/2866959-outputfeaturechannels?language=objc
func (c_ CNNBinaryConvolution) OutputFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("outputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryconvolution/2867126-inputfeaturechannels?language=objc
func (c_ CNNBinaryConvolution) InputFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("inputFeatureChannels"))
	return rv
}
