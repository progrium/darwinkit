// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNNeuron] class.
var CNNNeuronClass = _CNNNeuronClass{objc.GetClass("MPSCNNNeuron")}

type _CNNNeuronClass struct {
	objc.Class
}

// An interface definition for the [CNNNeuron] class.
type ICNNNeuron interface {
	ICNNKernel
	A() float64
	Data() []byte
	NeuronType() CNNNeuronType
	C() float64
	B() float64
}

// A filter that applies a neuron activation function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron?language=objc
type CNNNeuron struct {
	CNNKernel
}

func CNNNeuronFrom(ptr unsafe.Pointer) CNNNeuron {
	return CNNNeuron{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (c_ CNNNeuron) InitWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuron {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuron](c_, objc.Sel("initWithDevice:neuronDescriptor:"), po0, objc.Ptr(neuronDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942315-initwithdevice?language=objc
func NewCNNNeuronWithDeviceNeuronDescriptor(device metal.PDevice, neuronDescriptor INNNeuronDescriptor) CNNNeuron {
	instance := CNNNeuronClass.Alloc().InitWithDeviceNeuronDescriptor(device, neuronDescriptor)
	instance.Autorelease()
	return instance
}

func (cc _CNNNeuronClass) Alloc() CNNNeuron {
	rv := objc.Call[CNNNeuron](cc, objc.Sel("alloc"))
	return rv
}

func CNNNeuron_Alloc() CNNNeuron {
	return CNNNeuronClass.Alloc()
}

func (cc _CNNNeuronClass) New() CNNNeuron {
	rv := objc.Call[CNNNeuron](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNNeuron() CNNNeuron {
	return CNNNeuronClass.New()
}

func (c_ CNNNeuron) Init() CNNNeuron {
	rv := objc.Call[CNNNeuron](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNNeuron) InitWithDevice(device metal.PDevice) CNNNeuron {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuron](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewCNNNeuronWithDevice(device metal.PDevice) CNNNeuron {
	instance := CNNNeuronClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNNeuron) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuron {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNNeuron](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNNeuron_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNNeuron {
	instance := CNNNeuronClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942297-a?language=objc
func (c_ CNNNeuron) A() float64 {
	rv := objc.Call[float64](c_, objc.Sel("a"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942308-data?language=objc
func (c_ CNNNeuron) Data() []byte {
	rv := objc.Call[[]byte](c_, objc.Sel("data"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942309-neurontype?language=objc
func (c_ CNNNeuron) NeuronType() CNNNeuronType {
	rv := objc.Call[CNNNeuronType](c_, objc.Sel("neuronType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942303-c?language=objc
func (c_ CNNNeuron) C() float64 {
	rv := objc.Call[float64](c_, objc.Sel("c"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnneuron/2942306-b?language=objc
func (c_ CNNNeuron) B() float64 {
	rv := objc.Call[float64](c_, objc.Sel("b"))
	return rv
}
