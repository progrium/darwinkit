// Code generated by DarwinKit. DO NOT EDIT.

package mps

import (
	"unsafe"

	"github.com/progrium/darwinkit/macos/metal"
	"github.com/progrium/darwinkit/objc"
)

// The class instance for the [SVGFDenoiser] class.
var SVGFDenoiserClass = _SVGFDenoiserClass{objc.GetClass("MPSSVGFDenoiser")}

type _SVGFDenoiserClass struct {
	objc.Class
}

// An interface definition for the [SVGFDenoiser] class.
type ISVGFDenoiser interface {
	objc.IObject
	EncodeToCommandBufferSourceTextureMotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer metal.PCommandBuffer, sourceTexture metal.PTexture, motionVectorTexture metal.PTexture, depthNormalTexture metal.PTexture, previousDepthNormalTexture metal.PTexture) metal.TextureObject
	EncodeToCommandBufferObjectSourceTextureObjectMotionVectorTextureObjectDepthNormalTextureObjectPreviousDepthNormalTextureObject(commandBufferObject objc.IObject, sourceTextureObject objc.IObject, motionVectorTextureObject objc.IObject, depthNormalTextureObject objc.IObject, previousDepthNormalTextureObject objc.IObject) metal.TextureObject
	ClearTemporalHistory()
	ReleaseTemporaryTextures()
	EncodeToCommandBufferSourceTextureDestinationTextureSourceTexture2DestinationTexture2MotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer metal.PCommandBuffer, sourceTexture metal.PTexture, destinationTexture unsafe.Pointer, sourceTexture2 metal.PTexture, destinationTexture2 unsafe.Pointer, motionVectorTexture metal.PTexture, depthNormalTexture metal.PTexture, previousDepthNormalTexture metal.PTexture)
	EncodeToCommandBufferObjectSourceTextureObjectDestinationTextureObjectSourceTexture2ObjectDestinationTexture2ObjectMotionVectorTextureObjectDepthNormalTextureObjectPreviousDepthNormalTextureObject(commandBufferObject objc.IObject, sourceTextureObject objc.IObject, destinationTextureObject objc.IObject, sourceTexture2Object objc.IObject, destinationTexture2Object objc.IObject, motionVectorTextureObject objc.IObject, depthNormalTextureObject objc.IObject, previousDepthNormalTextureObject objc.IObject)
	Svgf() SVGF
	BilateralFilterIterations() uint
	SetBilateralFilterIterations(value uint)
	TextureAllocator() SVGFTextureAllocatorObject
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser?language=objc
type SVGFDenoiser struct {
	objc.Object
}

func SVGFDenoiserFrom(ptr unsafe.Pointer) SVGFDenoiser {
	return SVGFDenoiser{
		Object: objc.ObjectFrom(ptr),
	}
}

func (s_ SVGFDenoiser) InitWithSVGFTextureAllocator(svgf ISVGF, textureAllocator PSVGFTextureAllocator) SVGFDenoiser {
	po1 := objc.WrapAsProtocol("MPSSVGFTextureAllocator", textureAllocator)
	rv := objc.Call[SVGFDenoiser](s_, objc.Sel("initWithSVGF:textureAllocator:"), svgf, po1)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3242908-initwithsvgf?language=objc
func NewSVGFDenoiserWithSVGFTextureAllocator(svgf ISVGF, textureAllocator PSVGFTextureAllocator) SVGFDenoiser {
	instance := SVGFDenoiserClass.Alloc().InitWithSVGFTextureAllocator(svgf, textureAllocator)
	instance.Autorelease()
	return instance
}

func (s_ SVGFDenoiser) InitWithDevice(device metal.PDevice) SVGFDenoiser {
	po0 := objc.WrapAsProtocol("MTLDevice", device)
	rv := objc.Call[SVGFDenoiser](s_, objc.Sel("initWithDevice:"), po0)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3353094-initwithdevice?language=objc
func NewSVGFDenoiserWithDevice(device metal.PDevice) SVGFDenoiser {
	instance := SVGFDenoiserClass.Alloc().InitWithDevice(device)
	instance.Autorelease()
	return instance
}

func (sc _SVGFDenoiserClass) Alloc() SVGFDenoiser {
	rv := objc.Call[SVGFDenoiser](sc, objc.Sel("alloc"))
	return rv
}

func (sc _SVGFDenoiserClass) New() SVGFDenoiser {
	rv := objc.Call[SVGFDenoiser](sc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewSVGFDenoiser() SVGFDenoiser {
	return SVGFDenoiserClass.New()
}

func (s_ SVGFDenoiser) Init() SVGFDenoiser {
	rv := objc.Call[SVGFDenoiser](s_, objc.Sel("init"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3353093-encodetocommandbuffer?language=objc
func (s_ SVGFDenoiser) EncodeToCommandBufferSourceTextureMotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer metal.PCommandBuffer, sourceTexture metal.PTexture, motionVectorTexture metal.PTexture, depthNormalTexture metal.PTexture, previousDepthNormalTexture metal.PTexture) metal.TextureObject {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	po1 := objc.WrapAsProtocol("MTLTexture", sourceTexture)
	po2 := objc.WrapAsProtocol("MTLTexture", motionVectorTexture)
	po3 := objc.WrapAsProtocol("MTLTexture", depthNormalTexture)
	po4 := objc.WrapAsProtocol("MTLTexture", previousDepthNormalTexture)
	rv := objc.Call[metal.TextureObject](s_, objc.Sel("encodeToCommandBuffer:sourceTexture:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:"), po0, po1, po2, po3, po4)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3353093-encodetocommandbuffer?language=objc
func (s_ SVGFDenoiser) EncodeToCommandBufferObjectSourceTextureObjectMotionVectorTextureObjectDepthNormalTextureObjectPreviousDepthNormalTextureObject(commandBufferObject objc.IObject, sourceTextureObject objc.IObject, motionVectorTextureObject objc.IObject, depthNormalTextureObject objc.IObject, previousDepthNormalTextureObject objc.IObject) metal.TextureObject {
	rv := objc.Call[metal.TextureObject](s_, objc.Sel("encodeToCommandBuffer:sourceTexture:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:"), commandBufferObject, sourceTextureObject, motionVectorTextureObject, depthNormalTextureObject, previousDepthNormalTextureObject)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3242903-cleartemporalhistory?language=objc
func (s_ SVGFDenoiser) ClearTemporalHistory() {
	objc.Call[objc.Void](s_, objc.Sel("clearTemporalHistory"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3242911-releasetemporarytextures?language=objc
func (s_ SVGFDenoiser) ReleaseTemporaryTextures() {
	objc.Call[objc.Void](s_, objc.Sel("releaseTemporaryTextures"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3353092-encodetocommandbuffer?language=objc
func (s_ SVGFDenoiser) EncodeToCommandBufferSourceTextureDestinationTextureSourceTexture2DestinationTexture2MotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer metal.PCommandBuffer, sourceTexture metal.PTexture, destinationTexture unsafe.Pointer, sourceTexture2 metal.PTexture, destinationTexture2 unsafe.Pointer, motionVectorTexture metal.PTexture, depthNormalTexture metal.PTexture, previousDepthNormalTexture metal.PTexture) {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	po1 := objc.WrapAsProtocol("MTLTexture", sourceTexture)
	po2 := objc.WrapAsProtocol("MTLTexture", destinationTexture)
	po3 := objc.WrapAsProtocol("MTLTexture", sourceTexture2)
	po4 := objc.WrapAsProtocol("MTLTexture", destinationTexture2)
	po5 := objc.WrapAsProtocol("MTLTexture", motionVectorTexture)
	po6 := objc.WrapAsProtocol("MTLTexture", depthNormalTexture)
	po7 := objc.WrapAsProtocol("MTLTexture", previousDepthNormalTexture)
	objc.Call[objc.Void](s_, objc.Sel("encodeToCommandBuffer:sourceTexture:destinationTexture:sourceTexture2:destinationTexture2:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:"), po0, po1, po2, po3, po4, po5, po6, po7)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3353092-encodetocommandbuffer?language=objc
func (s_ SVGFDenoiser) EncodeToCommandBufferObjectSourceTextureObjectDestinationTextureObjectSourceTexture2ObjectDestinationTexture2ObjectMotionVectorTextureObjectDepthNormalTextureObjectPreviousDepthNormalTextureObject(commandBufferObject objc.IObject, sourceTextureObject objc.IObject, destinationTextureObject objc.IObject, sourceTexture2Object objc.IObject, destinationTexture2Object objc.IObject, motionVectorTextureObject objc.IObject, depthNormalTextureObject objc.IObject, previousDepthNormalTextureObject objc.IObject) {
	objc.Call[objc.Void](s_, objc.Sel("encodeToCommandBuffer:sourceTexture:destinationTexture:sourceTexture2:destinationTexture2:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:"), commandBufferObject, sourceTextureObject, destinationTextureObject, sourceTexture2Object, destinationTexture2Object, motionVectorTextureObject, depthNormalTextureObject, previousDepthNormalTextureObject)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3242914-svgf?language=objc
func (s_ SVGFDenoiser) Svgf() SVGF {
	rv := objc.Call[SVGF](s_, objc.Sel("svgf"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3242902-bilateralfilteriterations?language=objc
func (s_ SVGFDenoiser) BilateralFilterIterations() uint {
	rv := objc.Call[uint](s_, objc.Sel("bilateralFilterIterations"))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3242902-bilateralfilteriterations?language=objc
func (s_ SVGFDenoiser) SetBilateralFilterIterations(value uint) {
	objc.Call[objc.Void](s_, objc.Sel("setBilateralFilterIterations:"), value)
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3242915-textureallocator?language=objc
func (s_ SVGFDenoiser) TextureAllocator() SVGFTextureAllocatorObject {
	rv := objc.Call[SVGFTextureAllocatorObject](s_, objc.Sel("textureAllocator"))
	return rv
}
