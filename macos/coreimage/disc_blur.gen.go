// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a disc blur filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidiscblur?language=objc
type PDiscBlur interface {
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

// A concrete type wrapper for the [PDiscBlur] protocol.
type DiscBlurWrapper struct {
	objc.Object
}

func (d_ DiscBlurWrapper) HasSetInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidiscblur/3228214-inputimage?language=objc
func (d_ DiscBlurWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](d_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (d_ DiscBlurWrapper) HasInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidiscblur/3228214-inputimage?language=objc
func (d_ DiscBlurWrapper) InputImage() Image {
	rv := objc.Call[Image](d_, objc.Sel("inputImage"))
	return rv
}

func (d_ DiscBlurWrapper) HasSetRadius() bool {
	return d_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius of the blur, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidiscblur/3228215-radius?language=objc
func (d_ DiscBlurWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setRadius:"), value)
}

func (d_ DiscBlurWrapper) HasRadius() bool {
	return d_.RespondsToSelector(objc.Sel("radius"))
}

// The radius of the blur, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidiscblur/3228215-radius?language=objc
func (d_ DiscBlurWrapper) Radius() float64 {
	rv := objc.Call[float64](d_, objc.Sel("radius"))
	return rv
}
