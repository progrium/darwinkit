// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure an sRGB-to-linear filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisrgbtonecurvetolinear?language=objc
type PSRGBToneCurveToLinear interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool
}

// A concrete type wrapper for the [PSRGBToneCurveToLinear] protocol.
type SRGBToneCurveToLinearWrapper struct {
	objc.Object
}

func (s_ SRGBToneCurveToLinearWrapper) HasSetInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisrgbtonecurvetolinear/3228698-inputimage?language=objc
func (s_ SRGBToneCurveToLinearWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](s_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (s_ SRGBToneCurveToLinearWrapper) HasInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisrgbtonecurvetolinear/3228698-inputimage?language=objc
func (s_ SRGBToneCurveToLinearWrapper) InputImage() Image {
	rv := objc.Call[Image](s_, objc.Sel("inputImage"))
	return rv
}
