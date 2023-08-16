// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a color controls filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcontrols?language=objc
type PColorControls interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetSaturation(value float64)
	HasSetSaturation() bool

	// optional
	Saturation() float64
	HasSaturation() bool

	// optional
	SetContrast(value float64)
	HasSetContrast() bool

	// optional
	Contrast() float64
	HasContrast() bool

	// optional
	SetBrightness(value float64)
	HasSetBrightness() bool

	// optional
	Brightness() float64
	HasBrightness() bool
}

// A concrete type wrapper for the [PColorControls] protocol.
type ColorControlsWrapper struct {
	objc.Object
}

func (c_ ColorControlsWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcontrols/3228126-inputimage?language=objc
func (c_ ColorControlsWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ ColorControlsWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcontrols/3228126-inputimage?language=objc
func (c_ ColorControlsWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ ColorControlsWrapper) HasSetSaturation() bool {
	return c_.RespondsToSelector(objc.Sel("setSaturation:"))
}

// The amount of saturation to apply. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcontrols/3228127-saturation?language=objc
func (c_ ColorControlsWrapper) SetSaturation(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setSaturation:"), value)
}

func (c_ ColorControlsWrapper) HasSaturation() bool {
	return c_.RespondsToSelector(objc.Sel("saturation"))
}

// The amount of saturation to apply. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcontrols/3228127-saturation?language=objc
func (c_ ColorControlsWrapper) Saturation() float64 {
	rv := objc.Call[float64](c_, objc.Sel("saturation"))
	return rv
}

func (c_ ColorControlsWrapper) HasSetContrast() bool {
	return c_.RespondsToSelector(objc.Sel("setContrast:"))
}

// The amount of contrast to apply. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcontrols/3228125-contrast?language=objc
func (c_ ColorControlsWrapper) SetContrast(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setContrast:"), value)
}

func (c_ ColorControlsWrapper) HasContrast() bool {
	return c_.RespondsToSelector(objc.Sel("contrast"))
}

// The amount of contrast to apply. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcontrols/3228125-contrast?language=objc
func (c_ ColorControlsWrapper) Contrast() float64 {
	rv := objc.Call[float64](c_, objc.Sel("contrast"))
	return rv
}

func (c_ ColorControlsWrapper) HasSetBrightness() bool {
	return c_.RespondsToSelector(objc.Sel("setBrightness:"))
}

// The amount of brightness to apply. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcontrols/3228124-brightness?language=objc
func (c_ ColorControlsWrapper) SetBrightness(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setBrightness:"), value)
}

func (c_ ColorControlsWrapper) HasBrightness() bool {
	return c_.RespondsToSelector(objc.Sel("brightness"))
}

// The amount of brightness to apply. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorcontrols/3228124-brightness?language=objc
func (c_ ColorControlsWrapper) Brightness() float64 {
	rv := objc.Call[float64](c_, objc.Sel("brightness"))
	return rv
}
