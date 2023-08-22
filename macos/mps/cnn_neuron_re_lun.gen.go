// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuronReLUN] class.
var CNNNeuronReLUNClass = _CNNNeuronReLUNClass{objc.GetClass("MPSCNNNeuronReLUN")}

type _CNNNeuronReLUNClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuronReLUN] class.
type ICNNNeuronReLUN interface {
	ICNNNeuron
}

// A ReLUN neuron filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuronrelun?language=objc
type CNNNeuronReLUN struct {
	CNNNeuron
}

func CNNNeuronReLUNFrom(ptr unsafe.Pointer) CNNNeuronReLUN {
	return CNNNeuronReLUN{
		CNNNeuron: CNNNeuronFrom(ptr),
	}
}

func (cc _CNNNeuronReLUNClass) Alloc() CNNNeuronReLUN {
	rv := objc.Call[CNNNeuronReLUN](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuronReLUN_Alloc() CNNNeuronReLUN {
	return CNNNeuronReLUNClass.Alloc()
}

func (cc _CNNNeuronReLUNClass) New() CNNNeuronReLUN {
	rv := objc.Call[CNNNeuronReLUN](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuronReLUN() CNNNeuronReLUN {
	return CNNNeuronReLUNClass.New()
}

func (c_ CNNNeuronReLUN) Init() CNNNeuronReLUN {
	rv := objc.Call[CNNNeuronReLUN](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNNeuronReLUN) InitWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronReLUN {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronReLUN](c_, objc.Sel("initWithDevice:neuronDescriptor:"), po0, objc.Ptr(neuronDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942315-initwithdevice?language=objc
func NewCNNNeuronReLUNWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuronReLUN {
	instance := CNNNeuronReLUNClass.Alloc().InitWithDeviceNeuronDescriptor(device, neuronDescriptor)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronReLUN) InitWithDevice(device metal.PDevice) CNNNeuronReLUN {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronReLUN](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNNeuronReLUNWithDevice(device metal.PDevice) CNNNeuronReLUN {
	instance := CNNNeuronReLUNClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuronReLUN) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronReLUN {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuronReLUN](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNNeuronReLUN_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuronReLUN {
	instance := CNNNeuronReLUNClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}
