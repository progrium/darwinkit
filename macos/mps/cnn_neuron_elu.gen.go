// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronELU] class.
var CNNNeuronELUClass = _CNNNeuronELUClass{objc.GetClass("MPSCNNNeuronELU")}

type _CNNNeuronELUClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronELU] class.
type ICNNNeuronELU interface {
	ICNNNeuron
}

// A parametric ELU neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronelu?language=objc
type CNNNeuronELU struct {
	CNNNeuron
}

func CNNNeuronELUFrom(ptr unsafe.Pointer) CNNNeuronELU {
	return CNNNeuronELU{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}

func (cc _CNNNeuronELUClass) Alloc() CNNNeuronELU {
	rv := objc.Call[CNNNeuronELU](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronELU_Alloc() CNNNeuronELU {
	return CNNNeuronELUClass.Alloc()
}

func (cc _CNNNeuronELUClass) New() CNNNeuronELU {
	rv := objc.Call[CNNNeuronELU](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronELU() CNNNeuronELU {
	return CNNNeuronELUClass.New()
}

func (c_ CNNNeuronELU) Init() CNNNeuronELU {
	rv := objc.Call[CNNNeuronELU](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNNeuronELU) InitWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronELU {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronELU](c_, objc.Sel("initWithDevice:neuronDescriptor:"), po0, objc.Ptr(neuronDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942315-initwithdevice?language=objc
func NewCNNNeuronELUWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronELU {
	instance := CNNNeuronELUClass.Alloc().InitWithDeviceNeuronDescriptor(device, neuronDescriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronELU) InitWithDevice(device metal.PDevice) CNNNeuronELU {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronELU](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNNeuronELUWithDevice(device metal.PDevice) CNNNeuronELU {
	instance := CNNNeuronELUClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronELU) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronELU {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronELU](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNNeuronELU_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronELU {
	instance := CNNNeuronELUClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
