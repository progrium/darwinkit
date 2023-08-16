// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a palette centroid filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettecentroid?language=objc
type PPaletteCentroid interface {
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

// A concrete type wrapper for the [PPaletteCentroid] protocol.
type PaletteCentroidWrapper struct {
	objc.Object
}

func (p_ PaletteCentroidWrapper) HasSetInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettecentroid/3228632-inputimage?language=objc
func (p_ PaletteCentroidWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](p_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (p_ PaletteCentroidWrapper) HasInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettecentroid/3228632-inputimage?language=objc
func (p_ PaletteCentroidWrapper) InputImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("inputImage"))
	return rv
}

func (p_ PaletteCentroidWrapper) HasSetPerceptual() bool {
	return p_.RespondsToSelector(objc.Sel("setPerceptual:"))
}

// A Boolean value that specifies whether the filter applies the color palette in a perceptual color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettecentroid/3228634-perceptual?language=objc
func (p_ PaletteCentroidWrapper) SetPerceptual(value bool) {
	objc.Call[objc.Void](p_, objc.Sel("setPerceptual:"), value)
}

func (p_ PaletteCentroidWrapper) HasPerceptual() bool {
	return p_.RespondsToSelector(objc.Sel("perceptual"))
}

// A Boolean value that specifies whether the filter applies the color palette in a perceptual color space. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettecentroid/3228634-perceptual?language=objc
func (p_ PaletteCentroidWrapper) Perceptual() bool {
	rv := objc.Call[bool](p_, objc.Sel("perceptual"))
	return rv
}

func (p_ PaletteCentroidWrapper) HasSetPaletteImage() bool {
	return p_.RespondsToSelector(objc.Sel("setPaletteImage:"))
}

// The input color palette, obtained by using a k-means clustering filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettecentroid/3228633-paletteimage?language=objc
func (p_ PaletteCentroidWrapper) SetPaletteImage(value IImage) {
	objc.Call[objc.Void](p_, objc.Sel("setPaletteImage:"), objc.Ptr(value))
}

func (p_ PaletteCentroidWrapper) HasPaletteImage() bool {
	return p_.RespondsToSelector(objc.Sel("paletteImage"))
}

// The input color palette, obtained by using a k-means clustering filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipalettecentroid/3228633-paletteimage?language=objc
func (p_ PaletteCentroidWrapper) PaletteImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("paletteImage"))
	return rv
}
