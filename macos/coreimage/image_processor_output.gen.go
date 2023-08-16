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

// A container for writing image data and information produced by a custom image processor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessoroutput?language=objc
type PImageProcessorOutput interface {
	// optional
	Surface() iosurface.Ref
	HasSurface() bool

	// optional
	MetalTexture() metal.PTexture
	HasMetalTexture() bool

	// optional
	MetalCommandBuffer() metal.PCommandBuffer
	HasMetalCommandBuffer() bool

	// optional
	PixelBuffer() corevideo.PixelBufferRef
	HasPixelBuffer() bool

	// optional
	BytesPerRow() uint
	HasBytesPerRow() bool

	// optional
	BaseAddress() unsafe.Pointer
	HasBaseAddress() bool

	// optional
	Region() coregraphics.Rect
	HasRegion() bool

	// optional
	Format() Format
	HasFormat() bool
}

// A concrete type wrapper for the [PImageProcessorOutput] protocol.
type ImageProcessorOutputWrapper struct {
	objc.Object
}

func (i_ ImageProcessorOutputWrapper) HasSurface() bool {
	return i_.RespondsToSelector(objc.Sel("surface"))
}

// An IOSurface object to which you can write output pixel data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessoroutput/1639627-surface?language=objc
func (i_ ImageProcessorOutputWrapper) Surface() iosurface.Ref {
	rv := objc.Call[iosurface.Ref](i_, objc.Sel("surface"))
	return rv
}

func (i_ ImageProcessorOutputWrapper) HasMetalTexture() bool {
	return i_.RespondsToSelector(objc.Sel("metalTexture"))
}

// A Metal texture to which you can write output pixel data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessoroutput/1639631-metaltexture?language=objc
func (i_ ImageProcessorOutputWrapper) MetalTexture() metal.TextureWrapper {
	rv := objc.Call[metal.TextureWrapper](i_, objc.Sel("metalTexture"))
	return rv
}

func (i_ ImageProcessorOutputWrapper) HasMetalCommandBuffer() bool {
	return i_.RespondsToSelector(objc.Sel("metalCommandBuffer"))
}

// A command buffer to use for image processing using Metal. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessoroutput/1639641-metalcommandbuffer?language=objc
func (i_ ImageProcessorOutputWrapper) MetalCommandBuffer() metal.CommandBufferWrapper {
	rv := objc.Call[metal.CommandBufferWrapper](i_, objc.Sel("metalCommandBuffer"))
	return rv
}

func (i_ ImageProcessorOutputWrapper) HasPixelBuffer() bool {
	return i_.RespondsToSelector(objc.Sel("pixelBuffer"))
}

// A CoreVideo pixel buffer to which you can write output pixel data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessoroutput/1639647-pixelbuffer?language=objc
func (i_ ImageProcessorOutputWrapper) PixelBuffer() corevideo.PixelBufferRef {
	rv := objc.Call[corevideo.PixelBufferRef](i_, objc.Sel("pixelBuffer"))
	return rv
}

func (i_ ImageProcessorOutputWrapper) HasBytesPerRow() bool {
	return i_.RespondsToSelector(objc.Sel("bytesPerRow"))
}

// The number of bytes per row of pixels for the output image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessoroutput/1639635-bytesperrow?language=objc
func (i_ ImageProcessorOutputWrapper) BytesPerRow() uint {
	rv := objc.Call[uint](i_, objc.Sel("bytesPerRow"))
	return rv
}

func (i_ ImageProcessorOutputWrapper) HasBaseAddress() bool {
	return i_.RespondsToSelector(objc.Sel("baseAddress"))
}

// A pointer to CPU memory at which to write output pixel data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessoroutput/1639626-baseaddress?language=objc
func (i_ ImageProcessorOutputWrapper) BaseAddress() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](i_, objc.Sel("baseAddress"))
	return rv
}

func (i_ ImageProcessorOutputWrapper) HasRegion() bool {
	return i_.RespondsToSelector(objc.Sel("region"))
}

// The rectangular region of the output image that your processor must provide. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessoroutput/1639629-region?language=objc
func (i_ ImageProcessorOutputWrapper) Region() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](i_, objc.Sel("region"))
	return rv
}

func (i_ ImageProcessorOutputWrapper) HasFormat() bool {
	return i_.RespondsToSelector(objc.Sel("format"))
}

// The per-pixel data format expected of the output image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessoroutput/1639628-format?language=objc
func (i_ ImageProcessorOutputWrapper) Format() Format {
	rv := objc.Call[Format](i_, objc.Sel("format"))
	return rv
}
