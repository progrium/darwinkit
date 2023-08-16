// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a masked variable blur filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimaskedvariableblur?language=objc
type PMaskedVariableBlur interface {
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

	// optional
	SetMask(value Image)
	HasSetMask() bool

	// optional
	Mask() IImage
	HasMask() bool
}

// A concrete type wrapper for the [PMaskedVariableBlur] protocol.
type MaskedVariableBlurWrapper struct {
	objc.Object
}

func (m_ MaskedVariableBlurWrapper) HasSetInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimaskedvariableblur/3228551-inputimage?language=objc
func (m_ MaskedVariableBlurWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](m_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (m_ MaskedVariableBlurWrapper) HasInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimaskedvariableblur/3228551-inputimage?language=objc
func (m_ MaskedVariableBlurWrapper) InputImage() Image {
	rv := objc.Call[Image](m_, objc.Sel("inputImage"))
	return rv
}

func (m_ MaskedVariableBlurWrapper) HasSetRadius() bool {
	return m_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The distance from the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimaskedvariableblur/3228553-radius?language=objc
func (m_ MaskedVariableBlurWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setRadius:"), value)
}

func (m_ MaskedVariableBlurWrapper) HasRadius() bool {
	return m_.RespondsToSelector(objc.Sel("radius"))
}

// The distance from the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimaskedvariableblur/3228553-radius?language=objc
func (m_ MaskedVariableBlurWrapper) Radius() float64 {
	rv := objc.Call[float64](m_, objc.Sel("radius"))
	return rv
}

func (m_ MaskedVariableBlurWrapper) HasSetMask() bool {
	return m_.RespondsToSelector(objc.Sel("setMask:"))
}

// A grayscale mask that defines the blur amount. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimaskedvariableblur/3228552-mask?language=objc
func (m_ MaskedVariableBlurWrapper) SetMask(value IImage) {
	objc.Call[objc.Void](m_, objc.Sel("setMask:"), objc.Ptr(value))
}

func (m_ MaskedVariableBlurWrapper) HasMask() bool {
	return m_.RespondsToSelector(objc.Sel("mask"))
}

// A grayscale mask that defines the blur amount. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimaskedvariableblur/3228552-mask?language=objc
func (m_ MaskedVariableBlurWrapper) Mask() Image {
	rv := objc.Call[Image](m_, objc.Sel("mask"))
	return rv
}
