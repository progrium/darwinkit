// AUTO-GENERATED CODE, DO NOT MODIFY

package mps

import (
	"github.com/progrium/macdriver/macos/metal"
	"github.com/progrium/macdriver/objc"
)

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgftextureallocator?language=objc
type PSVGFTextureAllocator interface {
	// optional
	ReturnTexture(texture metal.TextureWrapper)
	HasReturnTexture() bool

	// optional
	TextureWithPixelFormatWidthHeight(pixelFormat metal.PixelFormat, width uint, height uint) metal.PTexture
	HasTextureWithPixelFormatWidthHeight() bool
}

// A concrete type wrapper for the [PSVGFTextureAllocator] protocol.
type SVGFTextureAllocatorWrapper struct {
	objc.Object
}

func (s_ SVGFTextureAllocatorWrapper) HasReturnTexture() bool {
	return s_.RespondsToSelector(objc.Sel("returnTexture:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgftextureallocator/3242917-returntexture?language=objc
func (s_ SVGFTextureAllocatorWrapper) ReturnTexture(texture metal.PTexture) {
	po0 := objc.WrapAsProtocol("MTLTexture", texture)
	objc.Call[objc.Void](s_, objc.Sel("returnTexture:"), po0)
}

func (s_ SVGFTextureAllocatorWrapper) HasTextureWithPixelFormatWidthHeight() bool {
	return s_.RespondsToSelector(objc.Sel("textureWithPixelFormat:width:height:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/metalperformanceshaders/mpssvgftextureallocator/3242918-texturewithpixelformat?language=objc
func (s_ SVGFTextureAllocatorWrapper) TextureWithPixelFormatWidthHeight(pixelFormat metal.PixelFormat, width uint, height uint) metal.TextureWrapper {
	rv := objc.Call[metal.TextureWrapper](s_, objc.Sel("textureWithPixelFormat:width:height:"), pixelFormat, width, height)
	return rv
}
