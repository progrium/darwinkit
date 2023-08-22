// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronAbsolute] class.
var CNNNeuronAbsoluteClass = _CNNNeuronAbsoluteClass{objc.GetClass("MPSCNNNeuronAbsolute")}

type _CNNNeuronAbsoluteClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronAbsolute] class.
type ICNNNeuronAbsolute interface {
	ICNNNeuron
}

// An absolute neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronabsolute?language=objc
type CNNNeuronAbsolute struct {
	CNNNeuron
}

func CNNNeuronAbsoluteFrom(ptr unsafe.Pointer) CNNNeuronAbsolute {
	return CNNNeuronAbsolute{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}

func (cc _CNNNeuronAbsoluteClass) Alloc() CNNNeuronAbsolute {
	rv := objc.Call[CNNNeuronAbsolute](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronAbsolute_Alloc() CNNNeuronAbsolute {
	return CNNNeuronAbsoluteClass.Alloc()
}

func (cc _CNNNeuronAbsoluteClass) New() CNNNeuronAbsolute {
	rv := objc.Call[CNNNeuronAbsolute](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronAbsolute() CNNNeuronAbsolute {
	return CNNNeuronAbsoluteClass.New()
}

func (c_ CNNNeuronAbsolute) Init() CNNNeuronAbsolute {
	rv := objc.Call[CNNNeuronAbsolute](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNNeuronAbsolute) InitWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronAbsolute {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronAbsolute](c_, objc.Sel("initWithDevice:neuronDescriptor:"), po0, objc.Ptr(neuronDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942315-initwithdevice?language=objc
func NewCNNNeuronAbsoluteWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronAbsolute {
	instance := CNNNeuronAbsoluteClass.Alloc().InitWithDeviceNeuronDescriptor(device, neuronDescriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronAbsolute) InitWithDevice(device metal.PDevice) CNNNeuronAbsolute {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronAbsolute](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNNeuronAbsoluteWithDevice(device metal.PDevice) CNNNeuronAbsolute {
	instance := CNNNeuronAbsoluteClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronAbsolute) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronAbsolute {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronAbsolute](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNNeuronAbsolute_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronAbsolute {
	instance := CNNNeuronAbsoluteClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
