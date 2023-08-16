// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a straighten filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistraighten?language=objc
type PStraighten interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetAngle(value float64)
	HasSetAngle() bool

	// optional
	Angle() float64
	HasAngle() bool
}

// A concrete type wrapper for the [PStraighten] protocol.
type StraightenWrapper struct {
	objc.Object
}

func (s_ StraightenWrapper) HasSetInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistraighten/3228759-inputimage?language=objc
func (s_ StraightenWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](s_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (s_ StraightenWrapper) HasInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistraighten/3228759-inputimage?language=objc
func (s_ StraightenWrapper) InputImage() Image {
	rv := objc.Call[Image](s_, objc.Sel("inputImage"))
	return rv
}

func (s_ StraightenWrapper) HasSetAngle() bool {
	return s_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The rotation angle, in radians. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistraighten/3228758-angle?language=objc
func (s_ StraightenWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setAngle:"), value)
}

func (s_ StraightenWrapper) HasAngle() bool {
	return s_.RespondsToSelector(objc.Sel("angle"))
}

// The rotation angle, in radians. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistraighten/3228758-angle?language=objc
func (s_ StraightenWrapper) Angle() float64 {
	rv := objc.Call[float64](s_, objc.Sel("angle"))
	return rv
}
