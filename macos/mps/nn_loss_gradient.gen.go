// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [NNLossGradient] class.
var NNLossGradientClass = _NNLossGradientClass{objc.GetClass("MPSNNLossGradient")}

type _NNLossGradientClass struct {
	objc.Class
}

// An interface definition for the [NNLossGradient] class.
type INNLossGradient interface {
	ICNNBinaryKernel
	EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients(commandBuffer metal.PCommandBuffer, sourceGradients *foundation.Array, sourceImages *foundation.Array, labels *foundation.Array, weights *foundation.Array, sourceStates *foundation.Array, destinationGradients *foundation.Array)
	EncodeBatchToCommandBufferObjectSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients(commandBufferObject objc.IObject, sourceGradients *foundation.Array, sourceImages *foundation.Array, labels *foundation.Array, weights *foundation.Array, sourceStates *foundation.Array, destinationGradients *foundation.Array)
	NumberOfClasses() uint
	Weight() float64
	SetWeight(value float64)
	Epsilon() float64
	SetEpsilon(value float64)
	Delta() float64
	SetDelta(value float64)
	ComputeLabelGradients() bool
	SetComputeLabelGradients(value bool)
	ReductionType() CNNReductionType
	ReduceAcrossBatch() bool
	LossType() CNNLossType
	LabelSmoothing() float64
	SetLabelSmoothing(value float64)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient?language=objc
type NNLossGradient struct {
	CNNBinaryKernel
}

func NNLossGradientFrom(ptr unsafe.Pointer) NNLossGradient {
	return NNLossGradient{
		CNNBinaryKernel: CNNBinaryKernelFrom(ptr),
	}
}

func (n_ NNLossGradient) InitWithDeviceLossDescriptor(device metal.PDevice, lossDescriptor ICNNLossDescriptor) NNLossGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNLossGradient](n_, objc.Sel("initWithDevice:lossDescriptor:"), po0, objc.Ptr(lossDescriptor))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131817-initwithdevice?language=objc
func NewNNLossGradientWithDeviceLossDescriptor(device metal.PDevice, lossDescriptor ICNNLossDescriptor) NNLossGradient {
	instance := NNLossGradientClass.Alloc().InitWithDeviceLossDescriptor(device, lossDescriptor)
	instance.Autorelease()
	return instance
}

func (nc _NNLossGradientClass) Alloc() NNLossGradient {
	rv := objc.Call[NNLossGradient](nc, objc.Sel("alloc"))
	return rv
}

func NNLossGradient_Alloc() NNLossGradient {
	return NNLossGradientClass.Alloc()
}

func (nc _NNLossGradientClass) New() NNLossGradient {
	rv := objc.Call[NNLossGradient](nc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewNNLossGradient() NNLossGradient {
	return NNLossGradientClass.New()
}

func (n_ NNLossGradient) Init() NNLossGradient {
	rv := objc.Call[NNLossGradient](n_, objc.Sel("init"))
	return rv
}

func (n_ NNLossGradient) InitWithDevice(device metal.PDevice) NNLossGradient {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNLossGradient](n_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865629-initwithdevice?language=objc
func NewNNLossGradientWithDevice(device metal.PDevice) NNLossGradient {
	instance := NNLossGradientClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (n_ NNLossGradient) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNLossGradient {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[NNLossGradient](n_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func NNLossGradient_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) NNLossGradient {
	instance := NNLossGradientClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131814-encodebatchtocommandbuffer?language=objc
func (n_ NNLossGradient) EncodeBatchToCommandBufferSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients(commandBuffer metal.PCommandBuffer, sourceGradients *foundation.Array, sourceImages *foundation.Array, labels *foundation.Array, weights *foundation.Array, sourceStates *foundation.Array, destinationGradients *foundation.Array) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](n_, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:labels:weights:sourceStates:destinationGradients:"), po0, sourceGradients, sourceImages, labels, weights, sourceStates, destinationGradients)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131814-encodebatchtocommandbuffer?language=objc
func (n_ NNLossGradient) EncodeBatchToCommandBufferObjectSourceGradientsSourceImagesLabelsWeightsSourceStatesDestinationGradients(commandBufferObject objc.IObject, sourceGradients *foundation.Array, sourceImages *foundation.Array, labels *foundation.Array, weights *foundation.Array, sourceStates *foundation.Array, destinationGradients *foundation.Array) {
	objc.Call[objc.Void](n_, objc.Sel("encodeBatchToCommandBuffer:sourceGradients:sourceImages:labels:weights:sourceStates:destinationGradients:"), objc.Ptr(commandBufferObject), sourceGradients, sourceImages, labels, weights, sourceStates, destinationGradients)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131820-numberofclasses?language=objc
func (n_ NNLossGradient) NumberOfClasses() uint {
	rv := objc.Call[uint](n_, objc.Sel("numberOfClasses"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131822-weight?language=objc
func (n_ NNLossGradient) Weight() float64 {
	rv := objc.Call[float64](n_, objc.Sel("weight"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131822-weight?language=objc
func (n_ NNLossGradient) SetWeight(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setWeight:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131815-epsilon?language=objc
func (n_ NNLossGradient) Epsilon() float64 {
	rv := objc.Call[float64](n_, objc.Sel("epsilon"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131815-epsilon?language=objc
func (n_ NNLossGradient) SetEpsilon(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setEpsilon:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131812-delta?language=objc
func (n_ NNLossGradient) Delta() float64 {
	rv := objc.Call[float64](n_, objc.Sel("delta"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131812-delta?language=objc
func (n_ NNLossGradient) SetDelta(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setDelta:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131811-computelabelgradients?language=objc
func (n_ NNLossGradient) ComputeLabelGradients() bool {
	rv := objc.Call[bool](n_, objc.Sel("computeLabelGradients"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131811-computelabelgradients?language=objc
func (n_ NNLossGradient) SetComputeLabelGradients(value bool) {
	objc.Call[objc.Void](n_, objc.Sel("setComputeLabelGradients:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131821-reductiontype?language=objc
func (n_ NNLossGradient) ReductionType() CNNReductionType {
	rv := objc.Call[CNNReductionType](n_, objc.Sel("reductionType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3547986-reduceacrossbatch?language=objc
func (n_ NNLossGradient) ReduceAcrossBatch() bool {
	rv := objc.Call[bool](n_, objc.Sel("reduceAcrossBatch"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131819-losstype?language=objc
func (n_ NNLossGradient) LossType() CNNLossType {
	rv := objc.Call[CNNLossType](n_, objc.Sel("lossType"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131818-labelsmoothing?language=objc
func (n_ NNLossGradient) LabelSmoothing() float64 {
	rv := objc.Call[float64](n_, objc.Sel("labelSmoothing"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpsnnlossgradient/3131818-labelsmoothing?language=objc
func (n_ NNLossGradient) SetLabelSmoothing(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setLabelSmoothing:"), value)
}
