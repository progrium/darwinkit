// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronHardSigmoid] class.
var CNNNeuronHardSigmoidClass = _CNNNeuronHardSigmoidClass{objc.GetClass("MPSCNNNeuronHardSigmoid")}

type _CNNNeuronHardSigmoidClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronHardSigmoid] class.
type ICNNNeuronHardSigmoid interface {
	ICNNNeuron
}

// A hard sigmoid neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronhardsigmoid?language=objc
type CNNNeuronHardSigmoid struct {
	CNNNeuron
}

func CNNNeuronHardSigmoidFrom(ptr unsafe.Pointer) CNNNeuronHardSigmoid {
	return CNNNeuronHardSigmoid{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}

func (cc _CNNNeuronHardSigmoidClass) Alloc() CNNNeuronHardSigmoid {
	rv := objc.Call[CNNNeuronHardSigmoid](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronHardSigmoid_Alloc() CNNNeuronHardSigmoid {
	return CNNNeuronHardSigmoidClass.Alloc()
}

func (cc _CNNNeuronHardSigmoidClass) New() CNNNeuronHardSigmoid {
	rv := objc.Call[CNNNeuronHardSigmoid](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronHardSigmoid() CNNNeuronHardSigmoid {
	return CNNNeuronHardSigmoidClass.New()
}

func (c_ CNNNeuronHardSigmoid) Init() CNNNeuronHardSigmoid {
	rv := objc.Call[CNNNeuronHardSigmoid](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNNeuronHardSigmoid) InitWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronHardSigmoid {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronHardSigmoid](c_, objc.Sel("initWithDevice:neuronDescriptor:"), po0, objc.Ptr(neuronDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942315-initwithdevice?language=objc
func NewCNNNeuronHardSigmoidWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronHardSigmoid {
	instance := CNNNeuronHardSigmoidClass.Alloc().InitWithDeviceNeuronDescriptor(device, neuronDescriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronHardSigmoid) InitWithDevice(device metal.PDevice) CNNNeuronHardSigmoid {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronHardSigmoid](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNNeuronHardSigmoidWithDevice(device metal.PDevice) CNNNeuronHardSigmoid {
	instance := CNNNeuronHardSigmoidClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronHardSigmoid) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronHardSigmoid {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronHardSigmoid](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNNeuronHardSigmoid_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronHardSigmoid {
	instance := CNNNeuronHardSigmoidClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
