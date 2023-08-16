// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a zoom blur filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cizoomblur?language=objc
type PZoomBlur interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetAmount(value float64)
	HasSetAmount() bool

	// optional
	Amount() float64
	HasAmount() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PZoomBlur] protocol.
type ZoomBlurWrapper struct {
	objc.Object
}

func (z_ ZoomBlurWrapper) HasSetInputImage() bool {
	return z_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cizoomblur/3228843-inputimage?language=objc
func (z_ ZoomBlurWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](z_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (z_ ZoomBlurWrapper) HasInputImage() bool {
	return z_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cizoomblur/3228843-inputimage?language=objc
func (z_ ZoomBlurWrapper) InputImage() Image {
	rv := objc.Call[Image](z_, objc.Sel("inputImage"))
	return rv
}

func (z_ ZoomBlurWrapper) HasSetAmount() bool {
	return z_.RespondsToSelector(objc.Sel("setAmount:"))
}

// The zoom-in amount. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cizoomblur/3228841-amount?language=objc
func (z_ ZoomBlurWrapper) SetAmount(value float64) {
	objc.Call[objc.Void](z_, objc.Sel("setAmount:"), value)
}

func (z_ ZoomBlurWrapper) HasAmount() bool {
	return z_.RespondsToSelector(objc.Sel("amount"))
}

// The zoom-in amount. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cizoomblur/3228841-amount?language=objc
func (z_ ZoomBlurWrapper) Amount() float64 {
	rv := objc.Call[float64](z_, objc.Sel("amount"))
	return rv
}

func (z_ ZoomBlurWrapper) HasSetCenter() bool {
	return z_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The center of the effect, as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cizoomblur/3228842-center?language=objc
func (z_ ZoomBlurWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](z_, objc.Sel("setCenter:"), value)
}

func (z_ ZoomBlurWrapper) HasCenter() bool {
	return z_.RespondsToSelector(objc.Sel("center"))
}

// The center of the effect, as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cizoomblur/3228842-center?language=objc
func (z_ ZoomBlurWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](z_, objc.Sel("center"))
	return rv
}
