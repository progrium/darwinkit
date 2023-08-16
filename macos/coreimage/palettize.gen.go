// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a palettize filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettize?language=objc
type PPalettize interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetPerceptual(value bool)
	HasSetPerceptual() bool

	// optional
	Perceptual() bool
	HasPerceptual() bool

	// optional
	SetPaletteImage(value Image)
	HasSetPaletteImage() bool

	// optional
	PaletteImage() IImage
	HasPaletteImage() bool
}

// A concrete type wrapper for the [PPalettize] protocol.
type PalettizeWrapper struct {
	objc.Object
}

func (p_ PalettizeWrapper) HasSetInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettize/3228636-inputimage?language=objc
func (p_ PalettizeWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](p_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (p_ PalettizeWrapper) HasInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettize/3228636-inputimage?language=objc
func (p_ PalettizeWrapper) InputImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("inputImage"))
	return rv
}

func (p_ PalettizeWrapper) HasSetPerceptual() bool {
	return p_.RespondsToSelector(objc.Sel("setPerceptual:"))
}

// A Boolean value that specifies whether the filter applies the color palette in a perceptual color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettize/3228638-perceptual?language=objc
func (p_ PalettizeWrapper) SetPerceptual(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setPerceptual:"), value)
}

func (p_ PalettizeWrapper) HasPerceptual() bool {
	return p_.RespondsToSelector(objc.Sel("perceptual"))
}

// A Boolean value that specifies whether the filter applies the color palette in a perceptual color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettize/3228638-perceptual?language=objc
func (p_ PalettizeWrapper) Perceptual() bool {
	rv := objc.Call[bool](p_, objc.Sel("perceptual"))
	return rv
}

func (p_ PalettizeWrapper) HasSetPaletteImage() bool {
	return p_.RespondsToSelector(objc.Sel("setPaletteImage:"))
}

// The input color palette, obtained by using a k-means clustering filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettize/3228637-paletteimage?language=objc
func (p_ PalettizeWrapper) SetPaletteImage(value IImage) {
	objc.Call[objc.Void](p_, objc.Sel("setPaletteImage:"), objc.Ptr(value))
}

func (p_ PalettizeWrapper) HasPaletteImage() bool {
	return p_.RespondsToSelector(objc.Sel("paletteImage"))
}

// The input color palette, obtained by using a k-means clustering filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettize/3228637-paletteimage?language=objc
func (p_ PalettizeWrapper) PaletteImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("paletteImage"))
	return rv
}
