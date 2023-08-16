// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a color monochrome filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormonochrome?language=objc
type PColorMonochrome interface {
	// optional
	SetColor(value Color)
	HasSetColor() bool

	// optional
	Color() IColor
	HasColor() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetIntensity(value float64)
	HasSetIntensity() bool

	// optional
	Intensity() float64
	HasIntensity() bool
}

// A concrete type wrapper for the [PColorMonochrome] protocol.
type ColorMonochromeWrapper struct {
	objc.Object
}

func (c_ ColorMonochromeWrapper) HasSetColor() bool {
	return c_.RespondsToSelector(objc.Sel("setColor:"))
}

// The monochrome color to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormonochrome/3228167-color?language=objc
func (c_ ColorMonochromeWrapper) SetColor(value IColor) {
	objc.Call[objc.Void](c_, objc.Sel("setColor:"), objc.Ptr(value))
}

func (c_ ColorMonochromeWrapper) HasColor() bool {
	return c_.RespondsToSelector(objc.Sel("color"))
}

// The monochrome color to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormonochrome/3228167-color?language=objc
func (c_ ColorMonochromeWrapper) Color() Color {
	rv := objc.Call[Color](c_, objc.Sel("color"))
	return rv
}

func (c_ ColorMonochromeWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormonochrome/3228168-inputimage?language=objc
func (c_ ColorMonochromeWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ ColorMonochromeWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormonochrome/3228168-inputimage?language=objc
func (c_ ColorMonochromeWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ ColorMonochromeWrapper) HasSetIntensity() bool {
	return c_.RespondsToSelector(objc.Sel("setIntensity:"))
}

// The intensity of the monochrome effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormonochrome/3228169-intensity?language=objc
func (c_ ColorMonochromeWrapper) SetIntensity(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setIntensity:"), value)
}

func (c_ ColorMonochromeWrapper) HasIntensity() bool {
	return c_.RespondsToSelector(objc.Sel("intensity"))
}

// The intensity of the monochrome effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormonochrome/3228169-intensity?language=objc
func (c_ ColorMonochromeWrapper) Intensity() float64 {
	rv := objc.Call[float64](c_, objc.Sel("intensity"))
	return rv
}
