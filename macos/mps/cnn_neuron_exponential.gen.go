// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronExponential] class.
var CNNNeuronExponentialClass = _CNNNeuronExponentialClass{objc.GetClass("MPSCNNNeuronExponential")}

type _CNNNeuronExponentialClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronExponential] class.
type ICNNNeuronExponential interface {
	ICNNNeuron
}

// An exponential neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronexponential?language=objc
type CNNNeuronExponential struct {
	CNNNeuron
}

func CNNNeuronExponentialFrom(ptr unsafe.Pointer) CNNNeuronExponential {
	return CNNNeuronExponential{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}

func (cc _CNNNeuronExponentialClass) Alloc() CNNNeuronExponential {
	rv := objc.Call[CNNNeuronExponential](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronExponential_Alloc() CNNNeuronExponential {
	return CNNNeuronExponentialClass.Alloc()
}

func (cc _CNNNeuronExponentialClass) New() CNNNeuronExponential {
	rv := objc.Call[CNNNeuronExponential](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronExponential() CNNNeuronExponential {
	return CNNNeuronExponentialClass.New()
}

func (c_ CNNNeuronExponential) Init() CNNNeuronExponential {
	rv := objc.Call[CNNNeuronExponential](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNNeuronExponential) InitWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronExponential {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronExponential](c_, objc.Sel("initWithDevice:neuronDescriptor:"), po0, objc.Ptr(neuronDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942315-initwithdevice?language=objc
func NewCNNNeuronExponentialWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronExponential {
	instance := CNNNeuronExponentialClass.Alloc().InitWithDeviceNeuronDescriptor(device, neuronDescriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronExponential) InitWithDevice(device metal.PDevice) CNNNeuronExponential {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronExponential](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNNeuronExponentialWithDevice(device metal.PDevice) CNNNeuronExponential {
	instance := CNNNeuronExponentialClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronExponential) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronExponential {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronExponential](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNNeuronExponential_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronExponential {
	instance := CNNNeuronExponentialClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
