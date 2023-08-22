// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNBinaryFullyConnected] class.
var CNNBinaryFullyConnectedClass = _CNNBinaryFullyConnectedClass{objc.GetClass("MPSCNNBinaryFullyConnected")}

type _CNNBinaryFullyConnectedClass struct {
	objc.Class
}

// An interface definition for the [CNNBinaryFullyConnected] class.
type ICNNBinaryFullyConnected interface {
	ICNNBinaryConvolution
}

// A fully connected convolution layer with binary weights and optionally binarized input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryfullyconnected?language=objc
type CNNBinaryFullyConnected struct {
	CNNBinaryConvolution
}

func CNNBinaryFullyConnectedFrom(ptr unsafe.Pointer) CNNBinaryFullyConnected {
	return CNNBinaryFullyConnected{
		CNNBinaryConvolution: CNNBinaryConvolutionFrom(ptr),
	}
}

func (c_ CNNBinaryFullyConnected) InitWithDeviceConvolutionDataScaleValueTypeFlags(device metal.PDevice, convolutionData PCNNConvolutionDataSource, scaleValue float64, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryFullyConnected {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", convolutionData)
	rv := objc.Call[CNNBinaryFullyConnected](c_, objc.Sel("initWithDevice:convolutionData:scaleValue:type:flags:"), po0, po1, scaleValue, type_, flags)
	return rv
}

// Initializes a fully connected convolution layer with binary weights. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinaryfullyconnected/2867196-initwithdevice?language=objc
func NewCNNBinaryFullyConnectedWithDeviceConvolutionDataScaleValueTypeFlags(device metal.PDevice, convolutionData PCNNConvolutionDataSource, scaleValue float64, type_ CNNBinaryConvolutionType, flags CNNBinaryConvolutionFlags) CNNBinaryFullyConnected {
	instance := CNNBinaryFullyConnectedClass.Alloc().InitWithDeviceConvolutionDataScaleValueTypeFlags(device, convolutionData, scaleValue, type_, flags)
	instance.Autorelease()
	return instance
}

func (cc _CNNBinaryFullyConnectedClass) Alloc() CNNBinaryFullyConnected {
	rv := objc.Call[CNNBinaryFullyConnected](cc, objc.Sel("alloc"))
	return rv
}

func CNNBinaryFullyConnected_Alloc() CNNBinaryFullyConnected {
	return CNNBinaryFullyConnectedClass.Alloc()
}

func (cc _CNNBinaryFullyConnectedClass) New() CNNBinaryFullyConnected {
	rv := objc.Call[CNNBinaryFullyConnected](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNBinaryFullyConnected() CNNBinaryFullyConnected {
	return CNNBinaryFullyConnectedClass.New()
}

func (c_ CNNBinaryFullyConnected) Init() CNNBinaryFullyConnected {
	rv := objc.Call[CNNBinaryFullyConnected](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNBinaryFullyConnected) InitWithDevice(device metal.PDevice) CNNBinaryFullyConnected {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBinaryFullyConnected](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNBinaryFullyConnectedWithDevice(device metal.PDevice) CNNBinaryFullyConnected {
	instance := CNNBinaryFullyConnectedClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNBinaryFullyConnected) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNBinaryFullyConnected {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNBinaryFullyConnected](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNBinaryFullyConnected_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNBinaryFullyConnected {
	instance := CNNBinaryFullyConnectedClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
