// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a box blur filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciboxblur?language=objc
type PBoxBlur interface {
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

// A concrete type wrapper for the [PBoxBlur] protocol.
type BoxBlurWrapper struct {
	objc.Object
}

func (b_ BoxBlurWrapper) HasSetInputImage() bool {
	return b_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciboxblur/3228094-inputimage?language=objc
func (b_ BoxBlurWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](b_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (b_ BoxBlurWrapper) HasInputImage() bool {
	return b_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciboxblur/3228094-inputimage?language=objc
func (b_ BoxBlurWrapper) InputImage() Image {
	rv := objc.Call[Image](b_, objc.Sel("inputImage"))
	return rv
}

func (b_ BoxBlurWrapper) HasSetRadius() bool {
	return b_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius of the blur, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciboxblur/3228095-radius?language=objc
func (b_ BoxBlurWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setRadius:"), value)
}

func (b_ BoxBlurWrapper) HasRadius() bool {
	return b_.RespondsToSelector(objc.Sel("radius"))
}

// The radius of the blur, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciboxblur/3228095-radius?language=objc
func (b_ BoxBlurWrapper) Radius() float64 {
	rv := objc.Call[float64](b_, objc.Sel("radius"))
	return rv
}
