// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNNeuronDescriptor] class.
var NNNeuronDescriptorClass = _NNNeuronDescriptorClass{objc.GetClass("MPSNNNeuronDescriptor")}

type _NNNeuronDescriptorClass struct {
	objc.Class
}

// An interface definition for the [NNNeuronDescriptor] class.
type INNNeuronDescriptor interface {
	objc.IObject
	A() float64
	SetA(value float64)
	Data() []byte
	SetData(value []byte)
	NeuronType() CNNNeuronType
	SetNeuronType(value CNNNeuronType)
	C() float64
	SetC(value float64)
	B() float64
	SetB(value float64)
}

// An object that specifies properties used by a neuron kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor?language=objc
type NNNeuronDescriptor struct {
	objc.Object
}

func NNNeuronDescriptorFrom(ptr unsafe.Pointer) NNNeuronDescriptor {
	return NNNeuronDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (nc _NNNeuronDescriptorClass) Alloc() NNNeuronDescriptor {
	rv := objc.Call[NNNeuronDescriptor](nc, objc.Sel("alloc"))
	return rv
}

func NNNeuronDescriptor_Alloc() NNNeuronDescriptor {
	return NNNeuronDescriptorClass.Alloc()
}

func (nc _NNNeuronDescriptorClass) New() NNNeuronDescriptor {
	rv := objc.Call[NNNeuronDescriptor](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNNeuronDescriptor() NNNeuronDescriptor {
	return NNNeuronDescriptorClass.New()
}

func (n_ NNNeuronDescriptor) Init() NNNeuronDescriptor {
	rv := objc.Call[NNNeuronDescriptor](n_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942307-cnnneurondescriptorwithtype?language=objc
func (nc _NNNeuronDescriptorClass) CnnNeuronDescriptorWithType(neuronType CNNNeuronType) NNNeuronDescriptor {
	rv := objc.Call[NNNeuronDescriptor](nc, objc.Sel("cnnNeuronDescriptorWithType:"), neuronType)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942307-cnnneurondescriptorwithtype?language=objc
func NNNeuronDescriptor_CnnNeuronDescriptorWithType(neuronType CNNNeuronType) NNNeuronDescriptor {
	return NNNeuronDescriptorClass.CnnNeuronDescriptorWithType(neuronType)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942317-cnnneuronpreludescriptorwithdata?language=objc
func (nc _NNNeuronDescriptorClass) CnnNeuronPReLUDescriptorWithDataNoCopy(data []byte, noCopy bool) NNNeuronDescriptor {
	rv := objc.Call[NNNeuronDescriptor](nc, objc.Sel("cnnNeuronPReLUDescriptorWithData:noCopy:"), data, noCopy)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942317-cnnneuronpreludescriptorwithdata?language=objc
func NNNeuronDescriptor_CnnNeuronPReLUDescriptorWithDataNoCopy(data []byte, noCopy bool) NNNeuronDescriptor {
	return NNNeuronDescriptorClass.CnnNeuronPReLUDescriptorWithDataNoCopy(data, noCopy)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942316-a?language=objc
func (n_ NNNeuronDescriptor) A() float64 {
	rv := objc.Call[float64](n_, objc.Sel("a"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942316-a?language=objc
func (n_ NNNeuronDescriptor) SetA(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setA:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942299-data?language=objc
func (n_ NNNeuronDescriptor) Data() []byte {
	rv := objc.Call[[]byte](n_, objc.Sel("data"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942299-data?language=objc
func (n_ NNNeuronDescriptor) SetData(value []byte) {
	objc.Call[objc.Void](n_, objc.Sel("setData:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942292-neurontype?language=objc
func (n_ NNNeuronDescriptor) NeuronType() CNNNeuronType {
	rv := objc.Call[CNNNeuronType](n_, objc.Sel("neuronType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942292-neurontype?language=objc
func (n_ NNNeuronDescriptor) SetNeuronType(value CNNNeuronType) {
	objc.Call[objc.Void](n_, objc.Sel("setNeuronType:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942305-c?language=objc
func (n_ NNNeuronDescriptor) C() float64 {
	rv := objc.Call[float64](n_, objc.Sel("c"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942305-c?language=objc
func (n_ NNNeuronDescriptor) SetC(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setC:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942302-b?language=objc
func (n_ NNNeuronDescriptor) B() float64 {
	rv := objc.Call[float64](n_, objc.Sel("b"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnneurondescriptor/2942302-b?language=objc
func (n_ NNNeuronDescriptor) SetB(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setB:"), value)
}
