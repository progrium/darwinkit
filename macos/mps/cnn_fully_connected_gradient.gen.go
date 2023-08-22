// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNFullyConnectedGradient] class.
var CNNFullyConnectedGradientClass = _CNNFullyConnectedGradientClass{objc.GetClass("MPSCNNFullyConnectedGradient")}

type _CNNFullyConnectedGradientClass struct {
	objc.Class
}

// An interface definition for the [CNNFullyConnectedGradient] class.
type ICNNFullyConnectedGradient interface {
	ICNNConvolutionGradient
}

// A gradient fully connected convolution layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnectedgradient?language=objc
type CNNFullyConnectedGradient struct {
	CNNConvolutionGradient
}

func CNNFullyConnectedGradientFrom(ptr unsafe.Pointer) CNNFullyConnectedGradient {
	return CNNFullyConnectedGradient{
		CNNConvolutionGradient: CNNConvolutionGradientFrom(ptr),
	}
}

func (c_ CNNFullyConnectedGradient) InitWithDeviceWeights(device metal.PDevice, weights PCNNConvolutionDataSource) CNNFullyConnectedGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	po1 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", weights)
	rv := objc.Call[CNNFullyConnectedGradient](c_, objc.Sel("initWithDevice:weights:"), po0, po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnfullyconnectedgradient/2951921-initwithdevice?language=objc
func NewCNNFullyConnectedGradientWithDeviceWeights(device metal.PDevice, weights PCNNConvolutionDataSource) CNNFullyConnectedGradient {
	instance := CNNFullyConnectedGradientClass.Alloc().InitWithDeviceWeights(device, weights)
	instance.Autorelease()
	return instance
}

func (cc _CNNFullyConnectedGradientClass) Alloc() CNNFullyConnectedGradient {
	rv := objc.Call[CNNFullyConnectedGradient](cc, objc.Sel("alloc"))
	return rv
}

func CNNFullyConnectedGradient_Alloc() CNNFullyConnectedGradient {
	return CNNFullyConnectedGradientClass.Alloc()
}

func (cc _CNNFullyConnectedGradientClass) New() CNNFullyConnectedGradient {
	rv := objc.Call[CNNFullyConnectedGradient](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNFullyConnectedGradient() CNNFullyConnectedGradient {
	return CNNFullyConnectedGradientClass.New()
}

func (c_ CNNFullyConnectedGradient) Init() CNNFullyConnectedGradient {
	rv := objc.Call[CNNFullyConnectedGradient](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNFullyConnectedGradient) InitWithDevice(device metal.PDevice) CNNFullyConnectedGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNFullyConnectedGradient](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnngradientkernel/2942657-initwithdevice?language=objc
func NewCNNFullyConnectedGradientWithDevice(device metal.PDevice) CNNFullyConnectedGradient {
	instance := CNNFullyConnectedGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNFullyConnectedGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNFullyConnectedGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNFullyConnectedGradient](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNFullyConnectedGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNFullyConnectedGradient {
	instance := CNNFullyConnectedGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
