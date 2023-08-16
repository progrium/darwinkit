// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a morphology rectangle minimum filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectangleminimum?language=objc
type PMorphologyRectangleMinimum interface {
	// optional
	SetWidth(value float64)
	HasSetWidth() bool

	// optional
	Width() float64
	HasWidth() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetHeight(value float64)
	HasSetHeight() bool

	// optional
	Height() float64
	HasHeight() bool
}

// A concrete type wrapper for the [PMorphologyRectangleMinimum] protocol.
type MorphologyRectangleMinimumWrapper struct {
	objc.Object
}

func (m_ MorphologyRectangleMinimumWrapper) HasSetWidth() bool {
	return m_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width, in pixels, of the morphological structuring element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectangleminimum/3228589-width?language=objc
func (m_ MorphologyRectangleMinimumWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setWidth:"), value)
}

func (m_ MorphologyRectangleMinimumWrapper) HasWidth() bool {
	return m_.RespondsToSelector(objc.Sel("width"))
}

// The width, in pixels, of the morphological structuring element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectangleminimum/3228589-width?language=objc
func (m_ MorphologyRectangleMinimumWrapper) Width() float64 {
	rv := objc.Call[float64](m_, objc.Sel("width"))
	return rv
}

func (m_ MorphologyRectangleMinimumWrapper) HasSetInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectangleminimum/3228588-inputimage?language=objc
func (m_ MorphologyRectangleMinimumWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](m_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (m_ MorphologyRectangleMinimumWrapper) HasInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectangleminimum/3228588-inputimage?language=objc
func (m_ MorphologyRectangleMinimumWrapper) InputImage() Image {
	rv := objc.Call[Image](m_, objc.Sel("inputImage"))
	return rv
}

func (m_ MorphologyRectangleMinimumWrapper) HasSetHeight() bool {
	return m_.RespondsToSelector(objc.Sel("setHeight:"))
}

// The height, in pixels, of the morphological structuring element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectangleminimum/3228587-height?language=objc
func (m_ MorphologyRectangleMinimumWrapper) SetHeight(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setHeight:"), value)
}

func (m_ MorphologyRectangleMinimumWrapper) HasHeight() bool {
	return m_.RespondsToSelector(objc.Sel("height"))
}

// The height, in pixels, of the morphological structuring element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectangleminimum/3228587-height?language=objc
func (m_ MorphologyRectangleMinimumWrapper) Height() float64 {
	rv := objc.Call[float64](m_, objc.Sel("height"))
	return rv
}
