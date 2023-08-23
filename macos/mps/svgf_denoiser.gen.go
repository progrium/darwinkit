// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [SVGFDenoiser] class.
var SVGFDenoiserClass = _SVGFDenoiserClass{objc.GetClass("MPSSVGFDenoiser")}

type _SVGFDenoiserClass struct {
	objc.Class
}

// An interface definition for the [SVGFDenoiser] class.
type ISVGFDenoiser interface {
	objc.IObject
	ClearTemporalHistory()
	ReleaseTemporaryTextures()
	EncodeToCommandBufferSourceTextureMotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer metal.PCommandBuffer, sourceTexture metal.PTexture, motionVectorTexture metal.PTexture, depthNormalTexture metal.PTexture, previousDepthNormalTexture metal.PTexture) metal.TextureWrapper
	EncodeToCommandBufferObjectSourceTextureObjectMotionVectorTextureObjectDepthNormalTextureObjectPreviousDepthNormalTextureObject(commandBufferObject objc.IObject, sourceTextureObject objc.IObject, motionVectorTextureObject objc.IObject, depthNormalTextureObject objc.IObject, previousDepthNormalTextureObject objc.IObject) metal.TextureWrapper
	TextureAllocator() SVGFTextureAllocatorWrapper
	Svgf() SVGF
	BilateralFilterIterations() uint
	SetBilateralFilterIterations(value uint)
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
	rv := objc.Call[SVGFDenoiser](s_, objc.Sel("initWithSVGF:textureAllocator:"), objc.Ptr(svgf), po1)
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

func SVGFDenoiser_Alloc() SVGFDenoiser {
	return SVGFDenoiserClass.Alloc()
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
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3353093-encodetocommandbuffer?language=objc
func (s_ SVGFDenoiser) EncodeToCommandBufferSourceTextureMotionVectorTextureDepthNormalTexturePreviousDepthNormalTexture(commandBuffer metal.PCommandBuffer, sourceTexture metal.PTexture, motionVectorTexture metal.PTexture, depthNormalTexture metal.PTexture, previousDepthNormalTexture metal.PTexture) metal.TextureWrapper {
	po0 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	po1 := objc.WrapAsProtocol("MTLTexture", sourceTexture)
	po2 := objc.WrapAsProtocol("MTLTexture", motionVectorTexture)
	po3 := objc.WrapAsProtocol("MTLTexture", depthNormalTexture)
	po4 := objc.WrapAsProtocol("MTLTexture", previousDepthNormalTexture)
	rv := objc.Call[metal.TextureWrapper](s_, objc.Sel("encodeToCommandBuffer:sourceTexture:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:"), po0, po1, po2, po3, po4)
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3353093-encodetocommandbuffer?language=objc
func (s_ SVGFDenoiser) EncodeToCommandBufferObjectSourceTextureObjectMotionVectorTextureObjectDepthNormalTextureObjectPreviousDepthNormalTextureObject(commandBufferObject objc.IObject, sourceTextureObject objc.IObject, motionVectorTextureObject objc.IObject, depthNormalTextureObject objc.IObject, previousDepthNormalTextureObject objc.IObject) metal.TextureWrapper {
	rv := objc.Call[metal.TextureWrapper](s_, objc.Sel("encodeToCommandBuffer:sourceTexture:motionVectorTexture:depthNormalTexture:previousDepthNormalTexture:"), objc.Ptr(commandBufferObject), objc.Ptr(sourceTextureObject), objc.Ptr(motionVectorTextureObject), objc.Ptr(depthNormalTextureObject), objc.Ptr(previousDepthNormalTextureObject))
	return rv
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgfdenoiser/3242915-textureallocator?language=objc
func (s_ SVGFDenoiser) TextureAllocator() SVGFTextureAllocatorWrapper {
	rv := objc.Call[SVGFTextureAllocatorWrapper](s_, objc.Sel("textureAllocator"))
	return rv
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
