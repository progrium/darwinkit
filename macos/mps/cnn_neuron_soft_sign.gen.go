// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronSoftSign] class.
var CNNNeuronSoftSignClass = _CNNNeuronSoftSignClass{objc.GetClass("MPSCNNNeuronSoftSign")}

type _CNNNeuronSoftSignClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronSoftSign] class.
type ICNNNeuronSoftSign interface {
	ICNNNeuron
}

// A softsign neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronsoftsign?language=objc
type CNNNeuronSoftSign struct {
	CNNNeuron
}

func CNNNeuronSoftSignFrom(ptr unsafe.Pointer) CNNNeuronSoftSign {
	return CNNNeuronSoftSign{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}

func (cc _CNNNeuronSoftSignClass) Alloc() CNNNeuronSoftSign {
	rv := objc.Call[CNNNeuronSoftSign](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronSoftSign_Alloc() CNNNeuronSoftSign {
	return CNNNeuronSoftSignClass.Alloc()
}

func (cc _CNNNeuronSoftSignClass) New() CNNNeuronSoftSign {
	rv := objc.Call[CNNNeuronSoftSign](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronSoftSign() CNNNeuronSoftSign {
	return CNNNeuronSoftSignClass.New()
}

func (c_ CNNNeuronSoftSign) Init() CNNNeuronSoftSign {
	rv := objc.Call[CNNNeuronSoftSign](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNNeuronSoftSign) InitWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronSoftSign {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronSoftSign](c_, objc.Sel("initWithDevice:neuronDescriptor:"), po0, objc.Ptr(neuronDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942315-initwithdevice?language=objc
func NewCNNNeuronSoftSignWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronSoftSign {
	instance := CNNNeuronSoftSignClass.Alloc().InitWithDeviceNeuronDescriptor(device, neuronDescriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronSoftSign) InitWithDevice(device metal.PDevice) CNNNeuronSoftSign {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronSoftSign](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNNeuronSoftSignWithDevice(device metal.PDevice) CNNNeuronSoftSign {
	instance := CNNNeuronSoftSignClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronSoftSign) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronSoftSign {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronSoftSign](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNNeuronSoftSign_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronSoftSign {
	instance := CNNNeuronSoftSignClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
