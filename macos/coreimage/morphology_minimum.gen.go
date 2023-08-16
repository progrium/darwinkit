// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a morphology minimum filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyminimum?language=objc
type PMorphologyMinimum interface {
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

// A concrete type wrapper for the [PMorphologyMinimum] protocol.
type MorphologyMinimumWrapper struct {
	objc.Object
}

func (m_ MorphologyMinimumWrapper) HasSetInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyminimum/3228580-inputimage?language=objc
func (m_ MorphologyMinimumWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](m_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (m_ MorphologyMinimumWrapper) HasInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyminimum/3228580-inputimage?language=objc
func (m_ MorphologyMinimumWrapper) InputImage() Image {
	rv := objc.Call[Image](m_, objc.Sel("inputImage"))
	return rv
}

func (m_ MorphologyMinimumWrapper) HasSetRadius() bool {
	return m_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius of the circular morphological structuring element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyminimum/3228581-radius?language=objc
func (m_ MorphologyMinimumWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setRadius:"), value)
}

func (m_ MorphologyMinimumWrapper) HasRadius() bool {
	return m_.RespondsToSelector(objc.Sel("radius"))
}

// The radius of the circular morphological structuring element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyminimum/3228581-radius?language=objc
func (m_ MorphologyMinimumWrapper) Radius() float64 {
	rv := objc.Call[float64](m_, objc.Sel("radius"))
	return rv
}
