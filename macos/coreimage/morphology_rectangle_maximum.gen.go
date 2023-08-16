// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a morphology rectangle maximum filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectanglemaximum?language=objc
type PMorphologyRectangleMaximum interface {
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

// A concrete type wrapper for the [PMorphologyRectangleMaximum] protocol.
type MorphologyRectangleMaximumWrapper struct {
	objc.Object
}

func (m_ MorphologyRectangleMaximumWrapper) HasSetWidth() bool {
	return m_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width, in pixels, of the morphological structuring element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectanglemaximum/3228585-width?language=objc
func (m_ MorphologyRectangleMaximumWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setWidth:"), value)
}

func (m_ MorphologyRectangleMaximumWrapper) HasWidth() bool {
	return m_.RespondsToSelector(objc.Sel("width"))
}

// The width, in pixels, of the morphological structuring element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectanglemaximum/3228585-width?language=objc
func (m_ MorphologyRectangleMaximumWrapper) Width() float64 {
	rv := objc.Call[float64](m_, objc.Sel("width"))
	return rv
}

func (m_ MorphologyRectangleMaximumWrapper) HasSetInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectanglemaximum/3228584-inputimage?language=objc
func (m_ MorphologyRectangleMaximumWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](m_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (m_ MorphologyRectangleMaximumWrapper) HasInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectanglemaximum/3228584-inputimage?language=objc
func (m_ MorphologyRectangleMaximumWrapper) InputImage() Image {
	rv := objc.Call[Image](m_, objc.Sel("inputImage"))
	return rv
}

func (m_ MorphologyRectangleMaximumWrapper) HasSetHeight() bool {
	return m_.RespondsToSelector(objc.Sel("setHeight:"))
}

// The height, in pixels, of the morphological structuring element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectanglemaximum/3228583-height?language=objc
func (m_ MorphologyRectangleMaximumWrapper) SetHeight(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setHeight:"), value)
}

func (m_ MorphologyRectangleMaximumWrapper) HasHeight() bool {
	return m_.RespondsToSelector(objc.Sel("height"))
}

// The height, in pixels, of the morphological structuring element. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimorphologyrectanglemaximum/3228583-height?language=objc
func (m_ MorphologyRectangleMaximumWrapper) Height() float64 {
	rv := objc.Call[float64](m_, objc.Sel("height"))
	return rv
}
