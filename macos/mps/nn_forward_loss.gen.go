// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNForwardLoss] class.
var NNForwardLossClass = _NNForwardLossClass{objc.GetClass("MPSNNForwardLoss")}

type _NNForwardLossClass struct {
	objc.Class
}

// An interface definition for the [NNForwardLoss] class.
type INNForwardLoss interface {
	ICNNKernel
	EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationStateIsTemporary(commandBuffer metal.PCommandBuffer, sourceImages *foundation.Array, labels *foundation.Array, weights *foundation.Array, outStates *foundation.Array, isTemporary bool) *foundation.Array
	EncodeBatchToCommandBufferObjectSourceImagesLabelsWeightsDestinationStatesDestinationStateIsTemporary(commandBufferObject objc.IObject, sourceImages *foundation.Array, labels *foundation.Array, weights *foundation.Array, outStates *foundation.Array, isTemporary bool) *foundation.Array
	NumberOfClasses() uint
	Weight() float64
	SetWeight(value float64)
	Epsilon() float64
	SetEpsilon(value float64)
	Delta() float64
	SetDelta(value float64)
	ReductionType() CNNReductionType
	ReduceAcrossBatch() bool
	LossType() CNNLossType
	LabelSmoothing() float64
	SetLabelSmoothing(value float64)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss?language=objc
type NNForwardLoss struct {
	CNNKernel
}

func NNForwardLossFrom(ptr unsafe.Pointer) NNForwardLoss {
	return NNForwardLoss{
		CNNKernel: CNNKernelFrom(ptr),
	}
}

func (n_ NNForwardLoss) InitWithDeviceLossDescriptor(device metal.PDevice, lossDescriptor ICNNLossDescriptor) NNForwardLoss {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNForwardLoss](n_, objc.Sel("initWithDevice:lossDescriptor:"), po0, objc.Ptr(lossDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131802-initwithdevice?language=objc
func NewNNForwardLossWithDeviceLossDescriptor(device metal.PDevice, lossDescriptor ICNNLossDescriptor) NNForwardLoss {
	instance := NNForwardLossClass.Alloc().InitWithDeviceLossDescriptor(device, lossDescriptor)
	instance.Autorelease()
	return instance
}

func (nc _NNForwardLossClass) Alloc() NNForwardLoss {
	rv := objc.Call[NNForwardLoss](nc, objc.Sel("alloc"))
	return rv
}

func NNForwardLoss_Alloc() NNForwardLoss {
	return NNForwardLossClass.Alloc()
}

func (nc _NNForwardLossClass) New() NNForwardLoss {
	rv := objc.Call[NNForwardLoss](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNForwardLoss() NNForwardLoss {
	return NNForwardLossClass.New()
}

func (n_ NNForwardLoss) Init() NNForwardLoss {
	rv := objc.Call[NNForwardLoss](n_, objc.Sel("init"))
	return rv
}

func (n_ NNForwardLoss) InitWithDevice(device metal.PDevice) NNForwardLoss {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNForwardLoss](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnkernel/2865653-initwithdevice?language=objc
func NewNNForwardLossWithDevice(device metal.PDevice) NNForwardLoss {
	instance := NNForwardLossClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (n_ NNForwardLoss) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNForwardLoss {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNForwardLoss](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNForwardLoss_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNForwardLoss {
	instance := NNForwardLossClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131799-encodebatchtocommandbuffer?language=objc
func (n_ NNForwardLoss) EncodeBatchToCommandBufferSourceImagesLabelsWeightsDestinationStatesDestinationStateIsTemporary(commandBuffer metal.PCommandBuffer, sourceImages *foundation.Array, labels *foundation.Array, weights *foundation.Array, outStates *foundation.Array, isTemporary bool) *foundation.Array {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[*foundation.Array](n_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:weights:destinationStates:destinationStateIsTemporary:"), po0, sourceImages, labels, weights, outStates, isTemporary)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131799-encodebatchtocommandbuffer?language=objc
func (n_ NNForwardLoss) EncodeBatchToCommandBufferObjectSourceImagesLabelsWeightsDestinationStatesDestinationStateIsTemporary(commandBufferObject objc.IObject, sourceImages *foundation.Array, labels *foundation.Array, weights *foundation.Array, outStates *foundation.Array, isTemporary bool) *foundation.Array {
	rv := objc.Call[*foundation.Array](n_, objc.Sel("encodeBatchToCommandBuffer:sourceImages:labels:weights:destinationStates:destinationStateIsTemporary:"), objc.Ptr(commandBufferObject), sourceImages, labels, weights, outStates, isTemporary)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131805-numberofclasses?language=objc
func (n_ NNForwardLoss) NumberOfClasses() uint {
	rv := objc.Call[uint](n_, objc.Sel("numberOfClasses"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131807-weight?language=objc
func (n_ NNForwardLoss) Weight() float64 {
	rv := objc.Call[float64](n_, objc.Sel("weight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131807-weight?language=objc
func (n_ NNForwardLoss) SetWeight(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setWeight:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131800-epsilon?language=objc
func (n_ NNForwardLoss) Epsilon() float64 {
	rv := objc.Call[float64](n_, objc.Sel("epsilon"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131800-epsilon?language=objc
func (n_ NNForwardLoss) SetEpsilon(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setEpsilon:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131797-delta?language=objc
func (n_ NNForwardLoss) Delta() float64 {
	rv := objc.Call[float64](n_, objc.Sel("delta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131797-delta?language=objc
func (n_ NNForwardLoss) SetDelta(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setDelta:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131806-reductiontype?language=objc
func (n_ NNForwardLoss) ReductionType() CNNReductionType {
	rv := objc.Call[CNNReductionType](n_, objc.Sel("reductionType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3547985-reduceacrossbatch?language=objc
func (n_ NNForwardLoss) ReduceAcrossBatch() bool {
	rv := objc.Call[bool](n_, objc.Sel("reduceAcrossBatch"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131804-losstype?language=objc
func (n_ NNForwardLoss) LossType() CNNLossType {
	rv := objc.Call[CNNLossType](n_, objc.Sel("lossType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131803-labelsmoothing?language=objc
func (n_ NNForwardLoss) LabelSmoothing() float64 {
	rv := objc.Call[float64](n_, objc.Sel("labelSmoothing"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnforwardloss/3131803-labelsmoothing?language=objc
func (n_ NNForwardLoss) SetLabelSmoothing(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setLabelSmoothing:"), value)
}
