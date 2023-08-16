// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a sepia-tone filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisepiatone?language=objc
type PSepiaTone interface {
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

// A concrete type wrapper for the [PSepiaTone] protocol.
type SepiaToneWrapper struct {
	objc.Object
}

func (s_ SepiaToneWrapper) HasSetInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisepiatone/3228702-inputimage?language=objc
func (s_ SepiaToneWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](s_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (s_ SepiaToneWrapper) HasInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisepiatone/3228702-inputimage?language=objc
func (s_ SepiaToneWrapper) InputImage() Image {
	rv := objc.Call[Image](s_, objc.Sel("inputImage"))
	return rv
}

func (s_ SepiaToneWrapper) HasSetIntensity() bool {
	return s_.RespondsToSelector(objc.Sel("setIntensity:"))
}

// The intensity of the sepia effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisepiatone/3228703-intensity?language=objc
func (s_ SepiaToneWrapper) SetIntensity(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setIntensity:"), value)
}

func (s_ SepiaToneWrapper) HasIntensity() bool {
	return s_.RespondsToSelector(objc.Sel("intensity"))
}

// The intensity of the sepia effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisepiatone/3228703-intensity?language=objc
func (s_ SepiaToneWrapper) Intensity() float64 {
	rv := objc.Call[float64](s_, objc.Sel("intensity"))
	return rv
}
