// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a dither filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidither?language=objc
type PDither interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetIntensity(value float64)
	HasSetIntensity() bool

	// optional
	Intensity() float64
	HasIntensity() bool
}

// A concrete type wrapper for the [PDither] protocol.
type DitherWrapper struct {
	objc.Object
}

func (d_ DitherWrapper) HasSetInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidither/3228225-inputimage?language=objc
func (d_ DitherWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](d_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (d_ DitherWrapper) HasInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidither/3228225-inputimage?language=objc
func (d_ DitherWrapper) InputImage() Image {
	rv := objc.Call[Image](d_, objc.Sel("inputImage"))
	return rv
}

func (d_ DitherWrapper) HasSetIntensity() bool {
	return d_.RespondsToSelector(objc.Sel("setIntensity:"))
}

// The intensity of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidither/3228226-intensity?language=objc
func (d_ DitherWrapper) SetIntensity(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setIntensity:"), value)
}

func (d_ DitherWrapper) HasIntensity() bool {
	return d_.RespondsToSelector(objc.Sel("intensity"))
}

// The intensity of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidither/3228226-intensity?language=objc
func (d_ DitherWrapper) Intensity() float64 {
	rv := objc.Call[float64](d_, objc.Sel("intensity"))
	return rv
}
