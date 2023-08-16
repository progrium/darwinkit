// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure an exposure adjust filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciexposureadjust?language=objc
type PExposureAdjust interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetEV(value float64)
	HasSetEV() bool

	// optional
	EV() float64
	HasEV() bool
}

// A concrete type wrapper for the [PExposureAdjust] protocol.
type ExposureAdjustWrapper struct {
	objc.Object
}

func (e_ ExposureAdjustWrapper) HasSetInputImage() bool {
	return e_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciexposureadjust/3228254-inputimage?language=objc
func (e_ ExposureAdjustWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](e_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (e_ ExposureAdjustWrapper) HasInputImage() bool {
	return e_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciexposureadjust/3228254-inputimage?language=objc
func (e_ ExposureAdjustWrapper) InputImage() Image {
	rv := objc.Call[Image](e_, objc.Sel("inputImage"))
	return rv
}

func (e_ ExposureAdjustWrapper) HasSetEV() bool {
	return e_.RespondsToSelector(objc.Sel("setEV:"))
}

// The amount to adjust the exposure of the image by. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciexposureadjust/3228253-ev?language=objc
func (e_ ExposureAdjustWrapper) SetEV(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setEV:"), value)
}

func (e_ ExposureAdjustWrapper) HasEV() bool {
	return e_.RespondsToSelector(objc.Sel("EV"))
}

// The amount to adjust the exposure of the image by. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciexposureadjust/3228253-ev?language=objc
func (e_ ExposureAdjustWrapper) EV() float64 {
	rv := objc.Call[float64](e_, objc.Sel("EV"))
	return rv
}
