// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronReLU] class.
var CNNNeuronReLUClass = _CNNNeuronReLUClass{objc.GetClass("MPSCNNNeuronReLU")}

type _CNNNeuronReLUClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronReLU] class.
type ICNNNeuronReLU interface {
	ICNNNeuron
}

// A ReLU (Rectified Linear Unit) neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelu?language=objc
type CNNNeuronReLU struct {
	CNNNeuron
}

func CNNNeuronReLUFrom(ptr unsafe.Pointer) CNNNeuronReLU {
	return CNNNeuronReLU{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}

func (cc _CNNNeuronReLUClass) Alloc() CNNNeuronReLU {
	rv := objc.Call[CNNNeuronReLU](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronReLU_Alloc() CNNNeuronReLU {
	return CNNNeuronReLUClass.Alloc()
}

func (cc _CNNNeuronReLUClass) New() CNNNeuronReLU {
	rv := objc.Call[CNNNeuronReLU](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronReLU() CNNNeuronReLU {
	return CNNNeuronReLUClass.New()
}

func (c_ CNNNeuronReLU) Init() CNNNeuronReLU {
	rv := objc.Call[CNNNeuronReLU](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNNeuronReLU) InitWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronReLU {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronReLU](c_, objc.Sel("initWithDevice:neuronDescriptor:"), po0, objc.Ptr(neuronDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942315-initwithdevice?language=objc
func NewCNNNeuronReLUWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronReLU {
	instance := CNNNeuronReLUClass.Alloc().InitWithDeviceNeuronDescriptor(device, neuronDescriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronReLU) InitWithDevice(device metal.PDevice) CNNNeuronReLU {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronReLU](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNNeuronReLUWithDevice(device metal.PDevice) CNNNeuronReLU {
	instance := CNNNeuronReLUClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronReLU) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronReLU {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronReLU](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNNeuronReLU_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronReLU {
	instance := CNNNeuronReLUClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
