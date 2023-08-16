// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a vignette-effect filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignetteeffect?language=objc
type PVignetteEffect interface {
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
	SetFalloff(value float64)
	HasSetFalloff() bool

	// optional
	Falloff() float64
	HasFalloff() bool

	// optional
	SetRadius(value float64)
	HasSetRadius() bool

	// optional
	Radius() float64
	HasRadius() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PVignetteEffect] protocol.
type VignetteEffectWrapper struct {
	objc.Object
}

func (v_ VignetteEffectWrapper) HasSetInputImage() bool {
	return v_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignetteeffect/3228832-inputimage?language=objc
func (v_ VignetteEffectWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](v_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (v_ VignetteEffectWrapper) HasInputImage() bool {
	return v_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignetteeffect/3228832-inputimage?language=objc
func (v_ VignetteEffectWrapper) InputImage() Image {
	rv := objc.Call[Image](v_, objc.Sel("inputImage"))
	return rv
}

func (v_ VignetteEffectWrapper) HasSetIntensity() bool {
	return v_.RespondsToSelector(objc.Sel("setIntensity:"))
}

// The intensity of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignetteeffect/3228833-intensity?language=objc
func (v_ VignetteEffectWrapper) SetIntensity(value float64) {
	objc.Call[objc.Void](v_, objc.Sel("setIntensity:"), value)
}

func (v_ VignetteEffectWrapper) HasIntensity() bool {
	return v_.RespondsToSelector(objc.Sel("intensity"))
}

// The intensity of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignetteeffect/3228833-intensity?language=objc
func (v_ VignetteEffectWrapper) Intensity() float64 {
	rv := objc.Call[float64](v_, objc.Sel("intensity"))
	return rv
}

func (v_ VignetteEffectWrapper) HasSetFalloff() bool {
	return v_.RespondsToSelector(objc.Sel("setFalloff:"))
}

// The falloff of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignetteeffect/3228831-falloff?language=objc
func (v_ VignetteEffectWrapper) SetFalloff(value float64) {
	objc.Call[objc.Void](v_, objc.Sel("setFalloff:"), value)
}

func (v_ VignetteEffectWrapper) HasFalloff() bool {
	return v_.RespondsToSelector(objc.Sel("falloff"))
}

// The falloff of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignetteeffect/3228831-falloff?language=objc
func (v_ VignetteEffectWrapper) Falloff() float64 {
	rv := objc.Call[float64](v_, objc.Sel("falloff"))
	return rv
}

func (v_ VignetteEffectWrapper) HasSetRadius() bool {
	return v_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The distance from the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignetteeffect/3228834-radius?language=objc
func (v_ VignetteEffectWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](v_, objc.Sel("setRadius:"), value)
}

func (v_ VignetteEffectWrapper) HasRadius() bool {
	return v_.RespondsToSelector(objc.Sel("radius"))
}

// The distance from the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignetteeffect/3228834-radius?language=objc
func (v_ VignetteEffectWrapper) Radius() float64 {
	rv := objc.Call[float64](v_, objc.Sel("radius"))
	return rv
}

func (v_ VignetteEffectWrapper) HasSetCenter() bool {
	return v_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The center of the effect as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignetteeffect/3228830-center?language=objc
func (v_ VignetteEffectWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](v_, objc.Sel("setCenter:"), value)
}

func (v_ VignetteEffectWrapper) HasCenter() bool {
	return v_.RespondsToSelector(objc.Sel("center"))
}

// The center of the effect as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civignetteeffect/3228830-center?language=objc
func (v_ VignetteEffectWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](v_, objc.Sel("center"))
	return rv
}
