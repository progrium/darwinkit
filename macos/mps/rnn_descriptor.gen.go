// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RNNDescriptor] class.
var RNNDescriptorClass = _RNNDescriptorClass{objc.GetClass("MPSRNNDescriptor")}

type _RNNDescriptorClass struct {
	objc.Class
}

// An interface definition for the [RNNDescriptor] class.
type IRNNDescriptor interface {
	objc.IObject
	LayerSequenceDirection() RNNSequenceDirection
	SetLayerSequenceDirection(value RNNSequenceDirection)
	OutputFeatureChannels() uint
	SetOutputFeatureChannels(value uint)
	UseLayerInputUnitTransformMode() bool
	SetUseLayerInputUnitTransformMode(value bool)
	InputFeatureChannels() uint
	SetInputFeatureChannels(value uint)
	UseFloat32Weights() bool
	SetUseFloat32Weights(value bool)
}

// A description of a recursive neural network block or layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor?language=objc
type RNNDescriptor struct {
	objc.Object
}

func RNNDescriptorFrom(ptr unsafe.Pointer) RNNDescriptor {
	return RNNDescriptor{
		Object: objc.ObjectFrom(ptr),
	}
}

func (rc _RNNDescriptorClass) Alloc() RNNDescriptor {
	rv := objc.Call[RNNDescriptor](rc, objc.Sel("alloc"))
	return rv
}

func RNNDescriptor_Alloc() RNNDescriptor {
	return RNNDescriptorClass.Alloc()
}

func (rc _RNNDescriptorClass) New() RNNDescriptor {
	rv := objc.Call[RNNDescriptor](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRNNDescriptor() RNNDescriptor {
	return RNNDescriptorClass.New()
}

func (r_ RNNDescriptor) Init() RNNDescriptor {
	rv := objc.Call[RNNDescriptor](r_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2865730-layersequencedirection?language=objc
func (r_ RNNDescriptor) LayerSequenceDirection() RNNSequenceDirection {
	rv := objc.Call[RNNSequenceDirection](r_, objc.Sel("layerSequenceDirection"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2865730-layersequencedirection?language=objc
func (r_ RNNDescriptor) SetLayerSequenceDirection(value RNNSequenceDirection) {
	objc.Call[objc.Void](r_, objc.Sel("setLayerSequenceDirection:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2865702-outputfeaturechannels?language=objc
func (r_ RNNDescriptor) OutputFeatureChannels() uint {
	rv := objc.Call[uint](r_, objc.Sel("outputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2865702-outputfeaturechannels?language=objc
func (r_ RNNDescriptor) SetOutputFeatureChannels(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setOutputFeatureChannels:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2865687-uselayerinputunittransformmode?language=objc
func (r_ RNNDescriptor) UseLayerInputUnitTransformMode() bool {
	rv := objc.Call[bool](r_, objc.Sel("useLayerInputUnitTransformMode"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2865687-uselayerinputunittransformmode?language=objc
func (r_ RNNDescriptor) SetUseLayerInputUnitTransformMode(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setUseLayerInputUnitTransformMode:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2865707-inputfeaturechannels?language=objc
func (r_ RNNDescriptor) InputFeatureChannels() uint {
	rv := objc.Call[uint](r_, objc.Sel("inputFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2865707-inputfeaturechannels?language=objc
func (r_ RNNDescriptor) SetInputFeatureChannels(value uint) {
	objc.Call[objc.Void](r_, objc.Sel("setInputFeatureChannels:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2881202-usefloat32weights?language=objc
func (r_ RNNDescriptor) UseFloat32Weights() bool {
	rv := objc.Call[bool](r_, objc.Sel("useFloat32Weights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnndescriptor/2881202-usefloat32weights?language=objc
func (r_ RNNDescriptor) SetUseFloat32Weights(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setUseFloat32Weights:"), value)
}
