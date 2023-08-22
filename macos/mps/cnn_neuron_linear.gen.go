// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronLinear] class.
var CNNNeuronLinearClass = _CNNNeuronLinearClass{objc.GetClass("MPSCNNNeuronLinear")}

type _CNNNeuronLinearClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronLinear] class.
type ICNNNeuronLinear interface {
	ICNNNeuron
}

// A linear neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlinear?language=objc
type CNNNeuronLinear struct {
	CNNNeuron
}

func CNNNeuronLinearFrom(ptr unsafe.Pointer) CNNNeuronLinear {
	return CNNNeuronLinear{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}

func (cc _CNNNeuronLinearClass) Alloc() CNNNeuronLinear {
	rv := objc.Call[CNNNeuronLinear](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronLinear_Alloc() CNNNeuronLinear {
	return CNNNeuronLinearClass.Alloc()
}

func (cc _CNNNeuronLinearClass) New() CNNNeuronLinear {
	rv := objc.Call[CNNNeuronLinear](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronLinear() CNNNeuronLinear {
	return CNNNeuronLinearClass.New()
}

func (c_ CNNNeuronLinear) Init() CNNNeuronLinear {
	rv := objc.Call[CNNNeuronLinear](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNNeuronLinear) InitWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronLinear {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronLinear](c_, objc.Sel("initWithDevice:neuronDescriptor:"), po0, objc.Ptr(neuronDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942315-initwithdevice?language=objc
func NewCNNNeuronLinearWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronLinear {
	instance := CNNNeuronLinearClass.Alloc().InitWithDeviceNeuronDescriptor(device, neuronDescriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronLinear) InitWithDevice(device metal.PDevice) CNNNeuronLinear {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronLinear](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNNeuronLinearWithDevice(device metal.PDevice) CNNNeuronLinear {
	instance := CNNNeuronLinearClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronLinear) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronLinear {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronLinear](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNNeuronLinear_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronLinear {
	instance := CNNNeuronLinearClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
