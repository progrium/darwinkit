// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronLogarithm] class.
var CNNNeuronLogarithmClass = _CNNNeuronLogarithmClass{objc.GetClass("MPSCNNNeuronLogarithm")}

type _CNNNeuronLogarithmClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronLogarithm] class.
type ICNNNeuronLogarithm interface {
	ICNNNeuron
}

// A logarithm neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronlogarithm?language=objc
type CNNNeuronLogarithm struct {
	CNNNeuron
}

func CNNNeuronLogarithmFrom(ptr unsafe.Pointer) CNNNeuronLogarithm {
	return CNNNeuronLogarithm{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}

func (cc _CNNNeuronLogarithmClass) Alloc() CNNNeuronLogarithm {
	rv := objc.Call[CNNNeuronLogarithm](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronLogarithm_Alloc() CNNNeuronLogarithm {
	return CNNNeuronLogarithmClass.Alloc()
}

func (cc _CNNNeuronLogarithmClass) New() CNNNeuronLogarithm {
	rv := objc.Call[CNNNeuronLogarithm](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronLogarithm() CNNNeuronLogarithm {
	return CNNNeuronLogarithmClass.New()
}

func (c_ CNNNeuronLogarithm) Init() CNNNeuronLogarithm {
	rv := objc.Call[CNNNeuronLogarithm](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNNeuronLogarithm) InitWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronLogarithm {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronLogarithm](c_, objc.Sel("initWithDevice:neuronDescriptor:"), po0, objc.Ptr(neuronDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942315-initwithdevice?language=objc
func NewCNNNeuronLogarithmWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronLogarithm {
	instance := CNNNeuronLogarithmClass.Alloc().InitWithDeviceNeuronDescriptor(device, neuronDescriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronLogarithm) InitWithDevice(device metal.PDevice) CNNNeuronLogarithm {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronLogarithm](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNNeuronLogarithmWithDevice(device metal.PDevice) CNNNeuronLogarithm {
	instance := CNNNeuronLogarithmClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronLogarithm) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronLogarithm {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronLogarithm](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNNeuronLogarithm_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronLogarithm {
	instance := CNNNeuronLogarithmClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
