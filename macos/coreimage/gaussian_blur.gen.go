// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a Gaussian blur filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussianblur?language=objc
type PGaussianBlur interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetRadius(value float64)
	HasSetRadius() bool

	// optional
	Radius() float64
	HasRadius() bool
}

// A concrete type wrapper for the [PGaussianBlur] protocol.
type GaussianBlurWrapper struct {
	objc.Object
}

func (g_ GaussianBlurWrapper) HasSetInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussianblur/3228464-inputimage?language=objc
func (g_ GaussianBlurWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](g_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (g_ GaussianBlurWrapper) HasInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussianblur/3228464-inputimage?language=objc
func (g_ GaussianBlurWrapper) InputImage() Image {
	rv := objc.Call[Image](g_, objc.Sel("inputImage"))
	return rv
}

func (g_ GaussianBlurWrapper) HasSetRadius() bool {
	return g_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius of the blur, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussianblur/3228465-radius?language=objc
func (g_ GaussianBlurWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setRadius:"), value)
}

func (g_ GaussianBlurWrapper) HasRadius() bool {
	return g_.RespondsToSelector(objc.Sel("radius"))
}

// The radius of the blur, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussianblur/3228465-radius?language=objc
func (g_ GaussianBlurWrapper) Radius() float64 {
	rv := objc.Call[float64](g_, objc.Sel("radius"))
	return rv
}
