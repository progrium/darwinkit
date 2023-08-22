// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronPReLU] class.
var CNNNeuronPReLUClass = _CNNNeuronPReLUClass{objc.GetClass("MPSCNNNeuronPReLU")}

type _CNNNeuronPReLUClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronPReLU] class.
type ICNNNeuronPReLU interface {
	ICNNNeuron
}

// A parametric ReLU (Rectified Linear Unit) neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronprelu?language=objc
type CNNNeuronPReLU struct {
	CNNNeuron
}

func CNNNeuronPReLUFrom(ptr unsafe.Pointer) CNNNeuronPReLU {
	return CNNNeuronPReLU{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}

func (cc _CNNNeuronPReLUClass) Alloc() CNNNeuronPReLU {
	rv := objc.Call[CNNNeuronPReLU](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronPReLU_Alloc() CNNNeuronPReLU {
	return CNNNeuronPReLUClass.Alloc()
}

func (cc _CNNNeuronPReLUClass) New() CNNNeuronPReLU {
	rv := objc.Call[CNNNeuronPReLU](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronPReLU() CNNNeuronPReLU {
	return CNNNeuronPReLUClass.New()
}

func (c_ CNNNeuronPReLU) Init() CNNNeuronPReLU {
	rv := objc.Call[CNNNeuronPReLU](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNNeuronPReLU) InitWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronPReLU {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronPReLU](c_, objc.Sel("initWithDevice:neuronDescriptor:"), po0, objc.Ptr(neuronDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942315-initwithdevice?language=objc
func NewCNNNeuronPReLUWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronPReLU {
	instance := CNNNeuronPReLUClass.Alloc().InitWithDeviceNeuronDescriptor(device, neuronDescriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronPReLU) InitWithDevice(device metal.PDevice) CNNNeuronPReLU {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronPReLU](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNNeuronPReLUWithDevice(device metal.PDevice) CNNNeuronPReLU {
	instance := CNNNeuronPReLUClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronPReLU) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronPReLU {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronPReLU](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNNeuronPReLU_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronPReLU {
	instance := CNNNeuronPReLUClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
