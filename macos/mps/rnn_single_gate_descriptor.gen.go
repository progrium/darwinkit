// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RNNSingleGateDescriptor] class.
var RNNSingleGateDescriptorClass = _RNNSingleGateDescriptorClass{objc.GetClass("MPSRNNSingleGateDescriptor")}

type _RNNSingleGateDescriptorClass struct {
	objc.Class
}

// An interface definition for the [RNNSingleGateDescriptor] class.
type IRNNSingleGateDescriptor interface {
	IRNNDescriptor
	RecurrentWeights() CNNConvolutionDataSourceWrapper
	SetRecurrentWeights(value PCNNConvolutionDataSource)
	SetRecurrentWeightsObject(valueObject objc.IObject)
	InputWeights() CNNConvolutionDataSourceWrapper
	SetInputWeights(value PCNNConvolutionDataSource)
	SetInputWeightsObject(valueObject objc.IObject)
}

// A description of a simple recurrent block or layer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor?language=objc
type RNNSingleGateDescriptor struct {
	RNNDescriptor
}

func RNNSingleGateDescriptorFrom(ptr unsafe.Pointer) RNNSingleGateDescriptor {
	return RNNSingleGateDescriptor{
		RNNDescriptor: RNNDescriptorFrom(ptr),
	}
}

func (rc _RNNSingleGateDescriptorClass) CreateRNNSingleGateDescriptorWithInputFeatureChannelsOutputFeatureChannels(inputFeatureChannels uint, outputFeatureChannels uint) RNNSingleGateDescriptor {
	rv := objc.Call[RNNSingleGateDescriptor](rc, objc.Sel("createRNNSingleGateDescriptorWithInputFeatureChannels:outputFeatureChannels:"), inputFeatureChannels, outputFeatureChannels)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865703-creaternnsinglegatedescriptorwit?language=objc
func RNNSingleGateDescriptor_CreateRNNSingleGateDescriptorWithInputFeatureChannelsOutputFeatureChannels(inputFeatureChannels uint, outputFeatureChannels uint) RNNSingleGateDescriptor {
	return RNNSingleGateDescriptorClass.CreateRNNSingleGateDescriptorWithInputFeatureChannelsOutputFeatureChannels(inputFeatureChannels, outputFeatureChannels)
}

func (rc _RNNSingleGateDescriptorClass) Alloc() RNNSingleGateDescriptor {
	rv := objc.Call[RNNSingleGateDescriptor](rc, objc.Sel("alloc"))
	return rv
}

func RNNSingleGateDescriptor_Alloc() RNNSingleGateDescriptor {
	return RNNSingleGateDescriptorClass.Alloc()
}

func (rc _RNNSingleGateDescriptorClass) New() RNNSingleGateDescriptor {
	rv := objc.Call[RNNSingleGateDescriptor](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRNNSingleGateDescriptor() RNNSingleGateDescriptor {
	return RNNSingleGateDescriptorClass.New()
}

func (r_ RNNSingleGateDescriptor) Init() RNNSingleGateDescriptor {
	rv := objc.Call[RNNSingleGateDescriptor](r_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865686-recurrentweights?language=objc
func (r_ RNNSingleGateDescriptor) RecurrentWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](r_, objc.Sel("recurrentWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865686-recurrentweights?language=objc
func (r_ RNNSingleGateDescriptor) SetRecurrentWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](r_, objc.Sel("setRecurrentWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865686-recurrentweights?language=objc
func (r_ RNNSingleGateDescriptor) SetRecurrentWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("setRecurrentWeights:"), objc.Ptr(valueObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865723-inputweights?language=objc
func (r_ RNNSingleGateDescriptor) InputWeights() CNNConvolutionDataSourceWrapper {
	rv := objc.Call[CNNConvolutionDataSourceWrapper](r_, objc.Sel("inputWeights"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865723-inputweights?language=objc
func (r_ RNNSingleGateDescriptor) SetInputWeights(value PCNNConvolutionDataSource) {
	po0 := objc.WrapAsProtocol("MPSCNNConvolutionDataSource", value)
	objc.Call[objc.Void](r_, objc.Sel("setInputWeights:"), po0)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsrnnsinglegatedescriptor/2865723-inputweights?language=objc
func (r_ RNNSingleGateDescriptor) SetInputWeightsObject(valueObject objc.IObject) {
	objc.Call[objc.Void](r_, objc.Sel("setInputWeights:"), objc.Ptr(valueObject))
}
