// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a sharpen luminance filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisharpenluminance?language=objc
type PSharpenLuminance interface {
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
	SetSharpness(value float64)
	HasSetSharpness() bool

	// optional
	Sharpness() float64
	HasSharpness() bool
}

// A concrete type wrapper for the [PSharpenLuminance] protocol.
type SharpenLuminanceWrapper struct {
	objc.Object
}

func (s_ SharpenLuminanceWrapper) HasSetInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisharpenluminance/3228709-inputimage?language=objc
func (s_ SharpenLuminanceWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](s_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (s_ SharpenLuminanceWrapper) HasInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisharpenluminance/3228709-inputimage?language=objc
func (s_ SharpenLuminanceWrapper) InputImage() Image {
	rv := objc.Call[Image](s_, objc.Sel("inputImage"))
	return rv
}

func (s_ SharpenLuminanceWrapper) HasSetRadius() bool {
	return s_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The distance from the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisharpenluminance/3228710-radius?language=objc
func (s_ SharpenLuminanceWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setRadius:"), value)
}

func (s_ SharpenLuminanceWrapper) HasRadius() bool {
	return s_.RespondsToSelector(objc.Sel("radius"))
}

// The distance from the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisharpenluminance/3228710-radius?language=objc
func (s_ SharpenLuminanceWrapper) Radius() float64 {
	rv := objc.Call[float64](s_, objc.Sel("radius"))
	return rv
}

func (s_ SharpenLuminanceWrapper) HasSetSharpness() bool {
	return s_.RespondsToSelector(objc.Sel("setSharpness:"))
}

// The amount of sharpening to apply. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisharpenluminance/3228711-sharpness?language=objc
func (s_ SharpenLuminanceWrapper) SetSharpness(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setSharpness:"), value)
}

func (s_ SharpenLuminanceWrapper) HasSharpness() bool {
	return s_.RespondsToSelector(objc.Sel("sharpness"))
}

// The amount of sharpening to apply. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisharpenluminance/3228711-sharpness?language=objc
func (s_ SharpenLuminanceWrapper) Sharpness() float64 {
	rv := objc.Call[float64](s_, objc.Sel("sharpness"))
	return rv
}
