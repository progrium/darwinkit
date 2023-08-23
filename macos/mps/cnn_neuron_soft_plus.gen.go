// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronSoftPlus] class.
var CNNNeuronSoftPlusClass = _CNNNeuronSoftPlusClass{objc.GetClass("MPSCNNNeuronSoftPlus")}

type _CNNNeuronSoftPlusClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronSoftPlus] class.
type ICNNNeuronSoftPlus interface {
	ICNNNeuron
}

// A parametric softplus neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftplus?language=objc
type CNNNeuronSoftPlus struct {
	CNNNeuron
}

func CNNNeuronSoftPlusFrom(ptr unsafe.Pointer) CNNNeuronSoftPlus {
	return CNNNeuronSoftPlus{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}

func (cc _CNNNeuronSoftPlusClass) Alloc() CNNNeuronSoftPlus {
	rv := objc.Call[CNNNeuronSoftPlus](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronSoftPlus_Alloc() CNNNeuronSoftPlus {
	return CNNNeuronSoftPlusClass.Alloc()
}

func (cc _CNNNeuronSoftPlusClass) New() CNNNeuronSoftPlus {
	rv := objc.Call[CNNNeuronSoftPlus](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronSoftPlus() CNNNeuronSoftPlus {
	return CNNNeuronSoftPlusClass.New()
}

func (c_ CNNNeuronSoftPlus) Init() CNNNeuronSoftPlus {
	rv := objc.Call[CNNNeuronSoftPlus](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNNeuronSoftPlus) InitWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronSoftPlus {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronSoftPlus](c_, objc.Sel("initWithDevice:neuronDescriptor:"), po0, objc.Ptr(neuronDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942315-initwithdevice?language=objc
func NewCNNNeuronSoftPlusWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronSoftPlus {
	instance := CNNNeuronSoftPlusClass.Alloc().InitWithDeviceNeuronDescriptor(device, neuronDescriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronSoftPlus) InitWithDevice(device metal.PDevice) CNNNeuronSoftPlus {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronSoftPlus](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNNeuronSoftPlusWithDevice(device metal.PDevice) CNNNeuronSoftPlus {
	instance := CNNNeuronSoftPlusClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronSoftPlus) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronSoftPlus {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronSoftPlus](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNNeuronSoftPlus_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronSoftPlus {
	instance := CNNNeuronSoftPlusClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
