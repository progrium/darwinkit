// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a linear-to-sRGB filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineartosrgbtonecurve?language=objc
type PLinearToSRGBToneCurve interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool
}

// A concrete type wrapper for the [PLinearToSRGBToneCurve] protocol.
type LinearToSRGBToneCurveWrapper struct {
	objc.Object
}

func (l_ LinearToSRGBToneCurveWrapper) HasSetInputImage() bool {
	return l_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineartosrgbtonecurve/3228547-inputimage?language=objc
func (l_ LinearToSRGBToneCurveWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](l_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (l_ LinearToSRGBToneCurveWrapper) HasInputImage() bool {
	return l_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineartosrgbtonecurve/3228547-inputimage?language=objc
func (l_ LinearToSRGBToneCurveWrapper) InputImage() Image {
	rv := objc.Call[Image](l_, objc.Sel("inputImage"))
	return rv
}
