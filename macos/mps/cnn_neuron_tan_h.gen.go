// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronTanH] class.
var CNNNeuronTanHClass = _CNNNeuronTanHClass{objc.GetClass("MPSCNNNeuronTanH")}

type _CNNNeuronTanHClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronTanH] class.
type ICNNNeuronTanH interface {
	ICNNNeuron
}

// A hyperbolic tangent neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneurontanh?language=objc
type CNNNeuronTanH struct {
	CNNNeuron
}

func CNNNeuronTanHFrom(ptr unsafe.Pointer) CNNNeuronTanH {
	return CNNNeuronTanH{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}

func (cc _CNNNeuronTanHClass) Alloc() CNNNeuronTanH {
	rv := objc.Call[CNNNeuronTanH](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronTanH_Alloc() CNNNeuronTanH {
	return CNNNeuronTanHClass.Alloc()
}

func (cc _CNNNeuronTanHClass) New() CNNNeuronTanH {
	rv := objc.Call[CNNNeuronTanH](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronTanH() CNNNeuronTanH {
	return CNNNeuronTanHClass.New()
}

func (c_ CNNNeuronTanH) Init() CNNNeuronTanH {
	rv := objc.Call[CNNNeuronTanH](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNNeuronTanH) InitWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronTanH {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronTanH](c_, objc.Sel("initWithDevice:neuronDescriptor:"), po0, objc.Ptr(neuronDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942315-initwithdevice?language=objc
func NewCNNNeuronTanHWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronTanH {
	instance := CNNNeuronTanHClass.Alloc().InitWithDeviceNeuronDescriptor(device, neuronDescriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronTanH) InitWithDevice(device metal.PDevice) CNNNeuronTanH {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronTanH](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNNeuronTanHWithDevice(device metal.PDevice) CNNNeuronTanH {
	instance := CNNNeuronTanHClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronTanH) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronTanH {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronTanH](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNNeuronTanH_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronTanH {
	instance := CNNNeuronTanHClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
