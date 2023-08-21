// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"unsafe"

	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/macos/corevideo"
	"github.com/progrium/macdriver/macos/iosurface"
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

// The class instance for the [RenderDestination] class.
var RenderDestinationClass = _RenderDestinationClass{objc.GetClass("CIRenderDestination")}

type _RenderDestinationClass struct {
	objc.Class
}

// An interface definition for the [RenderDestination] class.
type IRenderDestination interface {
	objc.IObject
	Width() uint
	BlendKernel() BlendKernel
	SetBlendKernel(value IBlendKernel)
	IsClamped() bool
	SetClamped(value bool)
	AlphaMode() RenderDestinationAlphaMode
	SetAlphaMode(value RenderDestinationAlphaMode)
	Height() uint
	ColorSpace() coregraphics.ColorSpaceRef
	SetColorSpace(value coregraphics.ColorSpaceRef)
	BlendsInDestinationColorSpace() bool
	SetBlendsInDestinationColorSpace(value bool)
	IsFlipped() bool
	SetFlipped(value bool)
	IsDithered() bool
	SetDithered(value bool)
}

// A specification for configuring all attributes of a render task's destination and issuing asynchronous render tasks. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination?language=objc
type RenderDestination struct {
	objc.Object
}

func RenderDestinationFrom(ptr unsafe.Pointer) RenderDestination {
	return RenderDestination{
		Object: objc.ObjectFrom(ptr),
	}
}

func (r_ RenderDestination) InitWithGLTextureTargetWidthHeight(texture int, target int, width uint, height uint) RenderDestination {
	rv := objc.Call[RenderDestination](r_, objc.Sel("initWithGLTexture:target:width:height:"), texture, target, width, height)
	return rv
}

// Creates a render destination based on an OpenGL texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875438-initwithgltexture?language=objc
func NewRenderDestinationWithGLTextureTargetWidthHeight(texture int, target int, width uint, height uint) RenderDestination {
	instance := RenderDestinationClass.Alloc().InitWithGLTextureTargetWidthHeight(texture, target, width, height)
	instance.Autorelease()
	return instance
}

func (r_ RenderDestination) InitWithIOSurface(surface iosurface.IIOSurface) RenderDestination {
	rv := objc.Call[RenderDestination](r_, objc.Sel("initWithIOSurface:"), objc.Ptr(surface))
	return rv
}

// Creates a render destination based on an IOSurface object. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2876044-initwithiosurface?language=objc
func NewRenderDestinationWithIOSurface(surface iosurface.IIOSurface) RenderDestination {
	instance := RenderDestinationClass.Alloc().InitWithIOSurface(surface)
	instance.Autorelease()
	return instance
}

func (r_ RenderDestination) InitWithBitmapDataWidthHeightBytesPerRowFormat(data unsafe.Pointer, width uint, height uint, bytesPerRow uint, format Format) RenderDestination {
	rv := objc.Call[RenderDestination](r_, objc.Sel("initWithBitmapData:width:height:bytesPerRow:format:"), data, width, height, bytesPerRow, format)
	return rv
}

// Creates a render destination based on a client-managed buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875427-initwithbitmapdata?language=objc
func NewRenderDestinationWithBitmapDataWidthHeightBytesPerRowFormat(data unsafe.Pointer, width uint, height uint, bytesPerRow uint, format Format) RenderDestination {
	instance := RenderDestinationClass.Alloc().InitWithBitmapDataWidthHeightBytesPerRowFormat(data, width, height, bytesPerRow, format)
	instance.Autorelease()
	return instance
}

func (r_ RenderDestination) InitWithWidthHeightPixelFormatCommandBufferMtlTextureProvider(width uint, height uint, pixelFormat metal.PixelFormat, commandBuffer metal.PCommandBuffer, block func() metal.TextureWrapper) RenderDestination {
	po3 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[RenderDestination](r_, objc.Sel("initWithWidth:height:pixelFormat:commandBuffer:mtlTextureProvider:"), width, height, pixelFormat, po3, block)
	return rv
}

// Creates a render destination based on a Metal texture with specified pixel format. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2880274-initwithwidth?language=objc
func NewRenderDestinationWithWidthHeightPixelFormatCommandBufferMtlTextureProvider(width uint, height uint, pixelFormat metal.PixelFormat, commandBuffer metal.PCommandBuffer, block func() metal.TextureWrapper) RenderDestination {
	instance := RenderDestinationClass.Alloc().InitWithWidthHeightPixelFormatCommandBufferMtlTextureProvider(width, height, pixelFormat, commandBuffer, block)
	instance.Autorelease()
	return instance
}

func (r_ RenderDestination) InitWithMTLTextureCommandBuffer(texture metal.PTexture, commandBuffer metal.PCommandBuffer) RenderDestination {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	po1 := objc.WrapAsProtocol("MTLCommandBuffer", commandBuffer)
	rv := objc.Call[RenderDestination](r_, objc.Sel("initWithMTLTexture:commandBuffer:"), po0, po1)
	return rv
}

// Creates a render destination based on a Metal texture. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2880273-initwithmtltexture?language=objc
func NewRenderDestinationWithMTLTextureCommandBuffer(texture metal.PTexture, commandBuffer metal.PCommandBuffer) RenderDestination {
	instance := RenderDestinationClass.Alloc().InitWithMTLTextureCommandBuffer(texture, commandBuffer)
	instance.Autorelease()
	return instance
}

func (r_ RenderDestination) InitWithPixelBuffer(pixelBuffer corevideo.PixelBufferRef) RenderDestination {
	rv := objc.Call[RenderDestination](r_, objc.Sel("initWithPixelBuffer:"), pixelBuffer)
	return rv
}

// Creates a render destination based on a Core Video pixel buffer. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875436-initwithpixelbuffer?language=objc
func NewRenderDestinationWithPixelBuffer(pixelBuffer corevideo.PixelBufferRef) RenderDestination {
	instance := RenderDestinationClass.Alloc().InitWithPixelBuffer(pixelBuffer)
	instance.Autorelease()
	return instance
}

func (rc _RenderDestinationClass) Alloc() RenderDestination {
	rv := objc.Call[RenderDestination](rc, objc.Sel("alloc"))
	return rv
}

func RenderDestination_Alloc() RenderDestination {
	return RenderDestinationClass.Alloc()
}

func (rc _RenderDestinationClass) New() RenderDestination {
	rv := objc.Call[RenderDestination](rc, objc.Sel("new"))
	rv.Autorelease()
	return rv
}

func NewRenderDestination() RenderDestination {
	return RenderDestinationClass.New()
}

func (r_ RenderDestination) Init() RenderDestination {
	rv := objc.Call[RenderDestination](r_, objc.Sel("init"))
	return rv
}

// The render destination's row width. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875434-width?language=objc
func (r_ RenderDestination) Width() uint {
	rv := objc.Call[uint](r_, objc.Sel("width"))
	return rv
}

// The destination's blend kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875452-blendkernel?language=objc
func (r_ RenderDestination) BlendKernel() BlendKernel {
	rv := objc.Call[BlendKernel](r_, objc.Sel("blendKernel"))
	return rv
}

// The destination's blend kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875452-blendkernel?language=objc
func (r_ RenderDestination) SetBlendKernel(value IBlendKernel) {
	objc.Call[objc.Void](r_, objc.Sel("setBlendKernel:"), objc.Ptr(value))
}

// Indicator of whether or not the destination clamps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875451-clamped?language=objc
func (r_ RenderDestination) IsClamped() bool {
	rv := objc.Call[bool](r_, objc.Sel("isClamped"))
	return rv
}

// Indicator of whether or not the destination clamps. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875451-clamped?language=objc
func (r_ RenderDestination) SetClamped(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setClamped:"), value)
}

// The render destination's representation of alpha (transparency) values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875443-alphamode?language=objc
func (r_ RenderDestination) AlphaMode() RenderDestinationAlphaMode {
	rv := objc.Call[RenderDestinationAlphaMode](r_, objc.Sel("alphaMode"))
	return rv
}

// The render destination's representation of alpha (transparency) values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875443-alphamode?language=objc
func (r_ RenderDestination) SetAlphaMode(value RenderDestinationAlphaMode) {
	objc.Call[objc.Void](r_, objc.Sel("setAlphaMode:"), value)
}

// The render destination's buffer height. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875433-height?language=objc
func (r_ RenderDestination) Height() uint {
	rv := objc.Call[uint](r_, objc.Sel("height"))
	return rv
}

// The destination's color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875439-colorspace?language=objc
func (r_ RenderDestination) ColorSpace() coregraphics.ColorSpaceRef {
	rv := objc.Call[coregraphics.ColorSpaceRef](r_, objc.Sel("colorSpace"))
	return rv
}

// The destination's color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875439-colorspace?language=objc
func (r_ RenderDestination) SetColorSpace(value coregraphics.ColorSpaceRef) {
	objc.Call[objc.Void](r_, objc.Sel("setColorSpace:"), value)
}

// Indicator of whether to blend in the destination's color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875437-blendsindestinationcolorspace?language=objc
func (r_ RenderDestination) BlendsInDestinationColorSpace() bool {
	rv := objc.Call[bool](r_, objc.Sel("blendsInDestinationColorSpace"))
	return rv
}

// Indicator of whether to blend in the destination's color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875437-blendsindestinationcolorspace?language=objc
func (r_ RenderDestination) SetBlendsInDestinationColorSpace(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setBlendsInDestinationColorSpace:"), value)
}

// Indicator of whether the destination is flipped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875442-flipped?language=objc
func (r_ RenderDestination) IsFlipped() bool {
	rv := objc.Call[bool](r_, objc.Sel("isFlipped"))
	return rv
}

// Indicator of whether the destination is flipped. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875442-flipped?language=objc
func (r_ RenderDestination) SetFlipped(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setFlipped:"), value)
}

// Indicator of whether or not the destination dithers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875441-dithered?language=objc
func (r_ RenderDestination) IsDithered() bool {
	rv := objc.Call[bool](r_, objc.Sel("isDithered"))
	return rv
}

// Indicator of whether or not the destination dithers. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cirenderdestination/2875441-dithered?language=objc
func (r_ RenderDestination) SetDithered(value bool) {
	objc.Call[objc.Void](r_, objc.Sel("setDithered:"), value)
}
