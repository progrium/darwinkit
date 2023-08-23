// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/foundation"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [TemporalAA] class.
var TemporalAAClass = _TemporalAAClass{objc.GetClass("MPSTemporalAA")}

type _TemporalAAClass struct {
	objc.Class
}

// An interface definition for the [TemporalAA] class.
type ITemporalAA interface {
	IKernel
	EncodeWithCoder(coder foundation.ICoder)
	EncodeToCommandBufferSourceTexturePreviousTextureDestinationTextureMotionVectorTextureDepthTexture(commandBuffer metal.PCommandBuffer, sourceTexture metal.PTexture, previousTexture metal.PTexture, destinationTexture metal.PTexture, motionVectorTexture metal.PTexture, depthTexture metal.PTexture)
	EncodeToCommandBufferObjectSourceTextureObjectPreviousTextureObjectDestinationTextureObjectMotionVectorTextureObjectDepthTextureObject(commandBufferObject objc.IObject, sourceTextureObject objc.IObject, previousTextureObject objc.IObject, destinationTextureObject objc.IObject, motionVectorTextureObject objc.IObject, depthTextureObject objc.IObject)
	BlendFactor() float64
	SetBlendFactor(value float64)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa?language=objc
type TemporalAA struct {
	Kernel
}

func TemporalAAFrom(ptr unsafe.Pointer) TemporalAA {
	return TemporalAA{
		Kernel: KernelFrom(ptr),
	}
}

func (t_ TemporalAA) InitWithDevice(device metal.PDevice) TemporalAA {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[TemporalAA](t_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143587-initwithdevice?language=objc
func NewTemporalAAWithDevice(device metal.PDevice) TemporalAA {
	instance := TemporalAAClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (t_ TemporalAA) CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) TemporalAA {
	po1 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[TemporalAA](t_, objc.Sel("copyWithZone:device:"), zone, po1)
	rv.Autorelease()
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143583-copywithzone?language=objc
func TemporalAA_CopyWithZoneDevice(zone unsafe.Pointer, device metal.PDevice) TemporalAA {
	instance := TemporalAAClass.Alloc().CopyWithZoneDevice(zone, device)
	instance.Autorelease()
	return instance
}

func (tc _TemporalAAClass) Alloc() TemporalAA {
	rv := objc.Call[TemporalAA](tc, objc.Sel("alloc"))
	return rv
}

func TemporalAA_Alloc() TemporalAA {
	return TemporalAAClass.Alloc()
}

func (tc _TemporalAAClass) New() TemporalAA {
	rv := objc.Call[TemporalAA](tc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewTemporalAA() TemporalAA {
	return TemporalAAClass.New()
}

func (t_ TemporalAA) Init() TemporalAA {
	rv := objc.Call[TemporalAA](t_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143585-encodewithcoder?language=objc
func (t_ TemporalAA) EncodeWithCoder(coder foundation.ICoder) {
	objc.Call[objc.Void](t_, objc.Sel("encodeWithCoder:"), objc.Ptr(coder))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143584-encodetocommandbuffer?language=objc
func (t_ TemporalAA) EncodeToCommandBufferSourceTexturePreviousTextureDestinationTextureMotionVectorTextureDepthTexture(commandBuffer metal.PCommandBuffer, sourceTexture metal.PTexture, previousTexture metal.PTexture, destinationTexture metal.PTexture, motionVectorTexture metal.PTexture, depthTexture metal.PTexture) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	po1 := objc.WrapAsProtocol("MTLTexture", sourceTexture)
	po2 := objc.WrapAsProtocol("MTLTexture", previousTexture)
	po3 := objc.WrapAsProtocol("MTLTexture", destinationTexture)
	po4 := objc.WrapAsProtocol("MTLTexture", motionVectorTexture)
	po5 := objc.WrapAsProtocol("MTLTexture", depthTexture)
	objc.Call[objc.Void](t_, objc.Sel("encodeToCommandBuffer:sourceTexture:previousTexture:destinationTexture:motionVectorTexture:depthTexture:"), po0, po1, po2, po3, po4, po5)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143584-encodetocommandbuffer?language=objc
func (t_ TemporalAA) EncodeToCommandBufferObjectSourceTextureObjectPreviousTextureObjectDestinationTextureObjectMotionVectorTextureObjectDepthTextureObject(commandBufferObject objc.IObject, sourceTextureObject objc.IObject, previousTextureObject objc.IObject, destinationTextureObject objc.IObject, motionVectorTextureObject objc.IObject, depthTextureObject objc.IObject) {
	objc.Call[objc.Void](t_, objc.Sel("encodeToCommandBuffer:sourceTexture:previousTexture:destinationTexture:motionVectorTexture:depthTexture:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceTextureObject), objc.Ptr(previousTextureObject), objc.Ptr(destinationTextureObject), objc.Ptr(motionVectorTextureObject), objc.Ptr(depthTextureObject))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143582-blendfactor?language=objc
func (t_ TemporalAA) BlendFactor() float64 {
	rv := objc.Call[float64](t_, objc.Sel("blendFactor"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpstemporalaa/3143582-blendfactor?language=objc
func (t_ TemporalAA) SetBlendFactor(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setBlendFactor:"), value)
}
