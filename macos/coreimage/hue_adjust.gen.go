// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a hue adjust filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihueadjust?language=objc
type PHueAdjust interface {
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

// A concrete type wrapper for the [PHueAdjust] protocol.
type HueAdjustWrapper struct {
	objc.Object
}

func (h_ HueAdjustWrapper) HasSetInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihueadjust/3228500-inputimage?language=objc
func (h_ HueAdjustWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](h_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (h_ HueAdjustWrapper) HasInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihueadjust/3228500-inputimage?language=objc
func (h_ HueAdjustWrapper) InputImage() Image {
	rv := objc.Call[Image](h_, objc.Sel("inputImage"))
	return rv
}

func (h_ HueAdjustWrapper) HasSetAngle() bool {
	return h_.RespondsToSelector(objc.Sel("setAngle:"))
}

// An angle, in radians, to use to correct the hue of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihueadjust/3228499-angle?language=objc
func (h_ HueAdjustWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](h_, objc.Sel("setAngle:"), value)
}

func (h_ HueAdjustWrapper) HasAngle() bool {
	return h_.RespondsToSelector(objc.Sel("angle"))
}

// An angle, in radians, to use to correct the hue of an image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihueadjust/3228499-angle?language=objc
func (h_ HueAdjustWrapper) Angle() float64 {
	rv := objc.Call[float64](h_, objc.Sel("angle"))
	return rv
}
