// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronPower] class.
var CNNNeuronPowerClass = _CNNNeuronPowerClass{objc.GetClass("MPSCNNNeuronPower")}

type _CNNNeuronPowerClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronPower] class.
type ICNNNeuronPower interface {
	ICNNNeuron
}

// A power neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronpower?language=objc
type CNNNeuronPower struct {
	CNNNeuron
}

func CNNNeuronPowerFrom(ptr unsafe.Pointer) CNNNeuronPower {
	return CNNNeuronPower{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}

func (cc _CNNNeuronPowerClass) Alloc() CNNNeuronPower {
	rv := objc.Call[CNNNeuronPower](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronPower_Alloc() CNNNeuronPower {
	return CNNNeuronPowerClass.Alloc()
}

func (cc _CNNNeuronPowerClass) New() CNNNeuronPower {
	rv := objc.Call[CNNNeuronPower](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronPower() CNNNeuronPower {
	return CNNNeuronPowerClass.New()
}

func (c_ CNNNeuronPower) Init() CNNNeuronPower {
	rv := objc.Call[CNNNeuronPower](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNNeuronPower) InitWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronPower {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronPower](c_, objc.Sel("initWithDevice:neuronDescriptor:"), po0, objc.Ptr(neuronDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942315-initwithdevice?language=objc
func NewCNNNeuronPowerWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronPower {
	instance := CNNNeuronPowerClass.Alloc().InitWithDeviceNeuronDescriptor(device, neuronDescriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronPower) InitWithDevice(device metal.PDevice) CNNNeuronPower {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronPower](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNNeuronPowerWithDevice(device metal.PDevice) CNNNeuronPower {
	instance := CNNNeuronPowerClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronPower) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronPower {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronPower](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNNeuronPower_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronPower {
	instance := CNNNeuronPowerClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
