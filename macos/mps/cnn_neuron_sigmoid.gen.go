// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronSigmoid] class.
var CNNNeuronSigmoidClass = _CNNNeuronSigmoidClass{objc.GetClass("MPSCNNNeuronSigmoid")}

type _CNNNeuronSigmoidClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronSigmoid] class.
type ICNNNeuronSigmoid interface {
	ICNNNeuron
}

// A sigmoid neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsigmoid?language=objc
type CNNNeuronSigmoid struct {
	CNNNeuron
}

func CNNNeuronSigmoidFrom(ptr unsafe.Pointer) CNNNeuronSigmoid {
	return CNNNeuronSigmoid{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}

func (cc _CNNNeuronSigmoidClass) Alloc() CNNNeuronSigmoid {
	rv := objc.Call[CNNNeuronSigmoid](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronSigmoid_Alloc() CNNNeuronSigmoid {
	return CNNNeuronSigmoidClass.Alloc()
}

func (cc _CNNNeuronSigmoidClass) New() CNNNeuronSigmoid {
	rv := objc.Call[CNNNeuronSigmoid](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronSigmoid() CNNNeuronSigmoid {
	return CNNNeuronSigmoidClass.New()
}

func (c_ CNNNeuronSigmoid) Init() CNNNeuronSigmoid {
	rv := objc.Call[CNNNeuronSigmoid](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNNeuronSigmoid) InitWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronSigmoid {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronSigmoid](c_, objc.Sel("initWithDevice:neuronDescriptor:"), po0, objc.Ptr(neuronDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942315-initwithdevice?language=objc
func NewCNNNeuronSigmoidWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronSigmoid {
	instance := CNNNeuronSigmoidClass.Alloc().InitWithDeviceNeuronDescriptor(device, neuronDescriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronSigmoid) InitWithDevice(device metal.PDevice) CNNNeuronSigmoid {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronSigmoid](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNNeuronSigmoidWithDevice(device metal.PDevice) CNNNeuronSigmoid {
	instance := CNNNeuronSigmoidClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronSigmoid) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronSigmoid {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronSigmoid](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNNeuronSigmoid_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronSigmoid {
	instance := CNNNeuronSigmoidClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
