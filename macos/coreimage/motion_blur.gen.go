// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a motion blur filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimotionblur?language=objc
type PMotionBlur interface {
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
	SetAngle(value float64)
	HasSetAngle() bool

	// optional
	Angle() float64
	HasAngle() bool
}

// A concrete type wrapper for the [PMotionBlur] protocol.
type MotionBlurWrapper struct {
	objc.Object
}

func (m_ MotionBlurWrapper) HasSetInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimotionblur/3228592-inputimage?language=objc
func (m_ MotionBlurWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](m_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (m_ MotionBlurWrapper) HasInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimotionblur/3228592-inputimage?language=objc
func (m_ MotionBlurWrapper) InputImage() Image {
	rv := objc.Call[Image](m_, objc.Sel("inputImage"))
	return rv
}

func (m_ MotionBlurWrapper) HasSetRadius() bool {
	return m_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius of the blur, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimotionblur/3228593-radius?language=objc
func (m_ MotionBlurWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setRadius:"), value)
}

func (m_ MotionBlurWrapper) HasRadius() bool {
	return m_.RespondsToSelector(objc.Sel("radius"))
}

// The radius of the blur, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimotionblur/3228593-radius?language=objc
func (m_ MotionBlurWrapper) Radius() float64 {
	rv := objc.Call[float64](m_, objc.Sel("radius"))
	return rv
}

func (m_ MotionBlurWrapper) HasSetAngle() bool {
	return m_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle of the motion, in radians, that determines which direction the blur smears. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimotionblur/3228591-angle?language=objc
func (m_ MotionBlurWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setAngle:"), value)
}

func (m_ MotionBlurWrapper) HasAngle() bool {
	return m_.RespondsToSelector(objc.Sel("angle"))
}

// The angle of the motion, in radians, that determines which direction the blur smears. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimotionblur/3228591-angle?language=objc
func (m_ MotionBlurWrapper) Angle() float64 {
	rv := objc.Call[float64](m_, objc.Sel("angle"))
	return rv
}
