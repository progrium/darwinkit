// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [CNNArithmetic] class.
var CNNArithmeticClass = _CNNArithmeticClass{objc.GetClass("MPSCNNArithmetic")}

type _CNNArithmeticClass struct {
	objc.Class
}

// An interface definition for the [CNNArithmetic] class.
type ICNNArithmetic interface {
	ICNNBinaryKernel
	EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationStatesDestinationImages(commandBuffer metal.PCommandBuffer, primaryImages *foundation.Array, secondaryImages *foundation.Array, destinationStates *foundation.Array, destinationImages *foundation.Array)
	EncodeBatchToCommandBufferObjectPrimaryImagesSecondaryImagesDestinationStatesDestinationImages(commandBufferObject objc.IObject, primaryImages *foundation.Array, secondaryImages *foundation.Array, destinationStates *foundation.Array, destinationImages *foundation.Array)
	EncodeToCommandBufferPrimaryImageSecondaryImageDestinationStateDestinationImage(commandBuffer metal.PCommandBuffer, primaryImage IImage, secondaryImage IImage, destinationState ICNNArithmeticGradientState, destinationImage IImage)
	EncodeToCommandBufferObjectPrimaryImageSecondaryImageDestinationStateDestinationImage(commandBufferObject objc.IObject, primaryImage IImage, secondaryImage IImage, destinationState ICNNArithmeticGradientState, destinationImage IImage)
	SecondaryScale() float64
	SetSecondaryScale(value float64)
	MaximumValue() float64
	SetMaximumValue(value float64)
	PrimaryScale() float64
	SetPrimaryScale(value float64)
	MinimumValue() float64
	SetMinimumValue(value float64)
	Bias() float64
	SetBias(value float64)
	SecondaryStrideInFeatureChannels() uint
	SetSecondaryStrideInFeatureChannels(value uint)
	PrimaryStrideInFeatureChannels() uint
	SetPrimaryStrideInFeatureChannels(value uint)
}

// The base class for arithmetic operators. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic?language=objc
type CNNArithmetic struct {
	CNNBinaryKernel
}

func CNNArithmeticFrom(ptr unsafe.Pointer) CNNArithmetic {
	return CNNArithmetic{
		CNNBinaryKernel: CNNBinaryKernelFrom(ptr),
	}
}

func (cc _CNNArithmeticClass) Alloc() CNNArithmetic {
	rv := objc.Call[CNNArithmetic](cc, objc.Sel("alloc"))
	return rv
}

func CNNArithmetic_Alloc() CNNArithmetic {
	return CNNArithmeticClass.Alloc()
}

func (cc _CNNArithmeticClass) New() CNNArithmetic {
	rv := objc.Call[CNNArithmetic](cc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewCNNArithmetic() CNNArithmetic {
	return CNNArithmeticClass.New()
}

func (c_ CNNArithmetic) Init() CNNArithmetic {
	rv := objc.Call[CNNArithmetic](c_, objc.Sel("init"))
	return rv
}

func (c_ CNNArithmetic) InitWithDevice(device metal.PDevice) CNNArithmetic {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNArithmetic](c_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnbinarykernel/2865629-initwithdevice?language=objc
func NewCNNArithmeticWithDevice(device metal.PDevice) CNNArithmetic {
	instance := CNNArithmeticClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (c_ CNNArithmetic) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNArithmetic {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[CNNArithmetic](c_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

// Makes a copy of this kernel object for a new device. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpskernel/1618912-copywithzone?language=objc
func CNNArithmetic_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) CNNArithmetic {
	instance := CNNArithmeticClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2954877-encodebatchtocommandbuffer?language=objc
func (c_ CNNArithmetic) EncodeBatchToCommandBufferPrimaryImagesSecondaryImagesDestinationStatesDestinationImages(commandBuffer metal.PCommandBuffer, primaryImages *foundation.Array, secondaryImages *foundation.Array, destinationStates *foundation.Array, destinationImages *foundation.Array) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](c_, objc.Sel("encodeBatchToCommandBuffer:primaryImages:secondaryImages:destinationStates:destinationImages:"), po0, primaryImages, secondaryImages, destinationStates, destinationImages)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2954877-encodebatchtocommandbuffer?language=objc
func (c_ CNNArithmetic) EncodeBatchToCommandBufferObjectPrimaryImagesSecondaryImagesDestinationStatesDestinationImages(commandBufferObject objc.IObject, primaryImages *foundation.Array, secondaryImages *foundation.Array, destinationStates *foundation.Array, destinationImages *foundation.Array) {
	objc.Call[objc.Void](c_, objc.Sel("encodeBatchToCommandBuffer:primaryImages:secondaryImages:destinationStates:destinationImages:"), objc.Ptr(commandBufferObject), primaryImages, secondaryImages, destinationStates, destinationImages)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2954876-encodetocommandbuffer?language=objc
func (c_ CNNArithmetic) EncodeToCommandBufferPrimaryImageSecondaryImageDestinationStateDestinationImage(commandBuffer metal.PCommandBuffer, primaryImage IImage, secondaryImage IImage, destinationState ICNNArithmeticGradientState, destinationImage IImage) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	objc.Call[objc.Void](c_, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:destinationState:destinationImage:"), po0, objc.Ptr(primaryImage), objc.Ptr(secondaryImage), objc.Ptr(destinationState), objc.Ptr(destinationImage))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2954876-encodetocommandbuffer?language=objc
func (c_ CNNArithmetic) EncodeToCommandBufferObjectPrimaryImageSecondaryImageDestinationStateDestinationImage(commandBufferObject objc.IObject, primaryImage IImage, secondaryImage IImage, destinationState ICNNArithmeticGradientState, destinationImage IImage) {
	objc.Call[objc.Void](c_, objc.Sel("encodeToCommandBuffer:primaryImage:secondaryImage:destinationState:destinationImage:"), objc.Ptr(commandBufferObject), objc.Ptr(primaryImage), objc.Ptr(secondaryImage), objc.Ptr(destinationState), objc.Ptr(destinationImage))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942497-secondaryscale?language=objc
func (c_ CNNArithmetic) SecondaryScale() float64 {
	rv := objc.Call[float64](c_, objc.Sel("secondaryScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942497-secondaryscale?language=objc
func (c_ CNNArithmetic) SetSecondaryScale(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setSecondaryScale:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942498-maximumvalue?language=objc
func (c_ CNNArithmetic) MaximumValue() float64 {
	rv := objc.Call[float64](c_, objc.Sel("maximumValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942498-maximumvalue?language=objc
func (c_ CNNArithmetic) SetMaximumValue(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setMaximumValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942509-primaryscale?language=objc
func (c_ CNNArithmetic) PrimaryScale() float64 {
	rv := objc.Call[float64](c_, objc.Sel("primaryScale"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942509-primaryscale?language=objc
func (c_ CNNArithmetic) SetPrimaryScale(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setPrimaryScale:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942502-minimumvalue?language=objc
func (c_ CNNArithmetic) MinimumValue() float64 {
	rv := objc.Call[float64](c_, objc.Sel("minimumValue"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942502-minimumvalue?language=objc
func (c_ CNNArithmetic) SetMinimumValue(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setMinimumValue:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942499-bias?language=objc
func (c_ CNNArithmetic) Bias() float64 {
	rv := objc.Call[float64](c_, objc.Sel("bias"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2942499-bias?language=objc
func (c_ CNNArithmetic) SetBias(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setBias:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2947964-secondarystrideinfeaturechannels?language=objc
func (c_ CNNArithmetic) SecondaryStrideInFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("secondaryStrideInFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2947964-secondarystrideinfeaturechannels?language=objc
func (c_ CNNArithmetic) SetSecondaryStrideInFeatureChannels(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setSecondaryStrideInFeatureChannels:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2947963-primarystrideinfeaturechannels?language=objc
func (c_ CNNArithmetic) PrimaryStrideInFeatureChannels() uint {
	rv := objc.Call[uint](c_, objc.Sel("primaryStrideInFeatureChannels"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpscnnarithmetic/2947963-primarystrideinfeaturechannels?language=objc
func (c_ CNNArithmetic) SetPrimaryStrideInFeatureChannels(value uint) {
	objc.Call[objc.Void](c_, objc.Sel("setPrimaryStrideInFeatureChannels:"), value)
}
