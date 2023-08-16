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

// A container of image data and information for use in a custom image processor. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorinput?language=objc
type PImageProcessorInput interface {
	// optional
	Surface() iosurface.Ref
	HasSurface() bool

	// optional
	MetalTexture() metal.PTexture
	HasMetalTexture() bool

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

// A concrete type wrapper for the [PImageProcessorInput] protocol.
type ImageProcessorInputWrapper struct {
	objc.Object
}

func (i_ ImageProcessorInputWrapper) HasSurface() bool {
	return i_.RespondsToSelector(objc.Sel("surface"))
}

// An IOSurface object containing the image data to be processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorinput/1639657-surface?language=objc
func (i_ ImageProcessorInputWrapper) Surface() iosurface.Ref {
	rv := objc.Call[iosurface.Ref](i_, objc.Sel("surface"))
	return rv
}

func (i_ ImageProcessorInputWrapper) HasMetalTexture() bool {
	return i_.RespondsToSelector(objc.Sel("metalTexture"))
}

// A Metal texture containing the image data to be processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorinput/1639651-metaltexture?language=objc
func (i_ ImageProcessorInputWrapper) MetalTexture() metal.TextureWrapper {
	rv := objc.Call[metal.TextureWrapper](i_, objc.Sel("metalTexture"))
	return rv
}

func (i_ ImageProcessorInputWrapper) HasPixelBuffer() bool {
	return i_.RespondsToSelector(objc.Sel("pixelBuffer"))
}

// A CoreVideo pixel buffer containing the image data to be processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorinput/1639649-pixelbuffer?language=objc
func (i_ ImageProcessorInputWrapper) PixelBuffer() corevideo.PixelBufferRef {
	rv := objc.Call[corevideo.PixelBufferRef](i_, objc.Sel("pixelBuffer"))
	return rv
}

func (i_ ImageProcessorInputWrapper) HasBytesPerRow() bool {
	return i_.RespondsToSelector(objc.Sel("bytesPerRow"))
}

// The number of bytes per row of pixels in the input image data. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorinput/1639655-bytesperrow?language=objc
func (i_ ImageProcessorInputWrapper) BytesPerRow() uint {
	rv := objc.Call[uint](i_, objc.Sel("bytesPerRow"))
	return rv
}

func (i_ ImageProcessorInputWrapper) HasBaseAddress() bool {
	return i_.RespondsToSelector(objc.Sel("baseAddress"))
}

// A pointer to the image data in CPU memory to be processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorinput/1639645-baseaddress?language=objc
func (i_ ImageProcessorInputWrapper) BaseAddress() unsafe.Pointer {
	rv := objc.Call[unsafe.Pointer](i_, objc.Sel("baseAddress"))
	return rv
}

func (i_ ImageProcessorInputWrapper) HasRegion() bool {
	return i_.RespondsToSelector(objc.Sel("region"))
}

// The area within the input image to be processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorinput/1639633-region?language=objc
func (i_ ImageProcessorInputWrapper) Region() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](i_, objc.Sel("region"))
	return rv
}

func (i_ ImageProcessorInputWrapper) HasFormat() bool {
	return i_.RespondsToSelector(objc.Sel("format"))
}

// The per-pixel data format of the image to be processed. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciimageprocessorinput/1639639-format?language=objc
func (i_ ImageProcessorInputWrapper) Format() Format {
	rv := objc.Call[Format](i_, objc.Sel("format"))
	return rv
}
