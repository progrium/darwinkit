// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a vignette filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignette?language=objc
type PVignette interface {
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

	// optional
	SetRadius(value float64)
	HasSetRadius() bool

	// optional
	Radius() float64
	HasRadius() bool
}

// A concrete type wrapper for the [PVignette] protocol.
type VignetteWrapper struct {
	objc.Object
}

func (v_ VignetteWrapper) HasSetInputImage() bool {
	return v_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignette/3228826-inputimage?language=objc
func (v_ VignetteWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](v_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (v_ VignetteWrapper) HasInputImage() bool {
	return v_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignette/3228826-inputimage?language=objc
func (v_ VignetteWrapper) InputImage() Image {
	rv := objc.Call[Image](v_, objc.Sel("inputImage"))
	return rv
}

func (v_ VignetteWrapper) HasSetIntensity() bool {
	return v_.RespondsToSelector(objc.Sel("setIntensity:"))
}

// The intensity of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignette/3228827-intensity?language=objc
func (v_ VignetteWrapper) SetIntensity(value float64) {
	objc.Call[objc.Void](v_, objc.Sel("setIntensity:"), value)
}

func (v_ VignetteWrapper) HasIntensity() bool {
	return v_.RespondsToSelector(objc.Sel("intensity"))
}

// The intensity of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignette/3228827-intensity?language=objc
func (v_ VignetteWrapper) Intensity() float64 {
	rv := objc.Call[float64](v_, objc.Sel("intensity"))
	return rv
}

func (v_ VignetteWrapper) HasSetRadius() bool {
	return v_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The distance from the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignette/3228828-radius?language=objc
func (v_ VignetteWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](v_, objc.Sel("setRadius:"), value)
}

func (v_ VignetteWrapper) HasRadius() bool {
	return v_.RespondsToSelector(objc.Sel("radius"))
}

// The distance from the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignette/3228828-radius?language=objc
func (v_ VignetteWrapper) Radius() float64 {
	rv := objc.Call[float64](v_, objc.Sel("radius"))
	return rv
}
