// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a spotlight filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight?language=objc
type PSpotLight interface {
	// optional
	SetColor(value Color)
	HasSetColor() bool

	// optional
	Color() IColor
	HasColor() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetLightPointsAt(value Vector)
	HasSetLightPointsAt() bool

	// optional
	LightPointsAt() IVector
	HasLightPointsAt() bool

	// optional
	SetConcentration(value float64)
	HasSetConcentration() bool

	// optional
	Concentration() float64
	HasConcentration() bool

	// optional
	SetLightPosition(value Vector)
	HasSetLightPosition() bool

	// optional
	LightPosition() IVector
	HasLightPosition() bool

	// optional
	SetBrightness(value float64)
	HasSetBrightness() bool

	// optional
	Brightness() float64
	HasBrightness() bool
}

// A concrete type wrapper for the [PSpotLight] protocol.
type SpotLightWrapper struct {
	objc.Object
}

func (s_ SpotLightWrapper) HasSetColor() bool {
	return s_.RespondsToSelector(objc.Sel("setColor:"))
}

// The color of the spotlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228743-color?language=objc
func (s_ SpotLightWrapper) SetColor(value IColor) {
	objc.Call[objc.Void](s_, objc.Sel("setColor:"), objc.Ptr(value))
}

func (s_ SpotLightWrapper) HasColor() bool {
	return s_.RespondsToSelector(objc.Sel("color"))
}

// The color of the spotlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228743-color?language=objc
func (s_ SpotLightWrapper) Color() Color {
	rv := objc.Call[Color](s_, objc.Sel("color"))
	return rv
}

func (s_ SpotLightWrapper) HasSetInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228745-inputimage?language=objc
func (s_ SpotLightWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](s_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (s_ SpotLightWrapper) HasInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228745-inputimage?language=objc
func (s_ SpotLightWrapper) InputImage() Image {
	rv := objc.Call[Image](s_, objc.Sel("inputImage"))
	return rv
}

func (s_ SpotLightWrapper) HasSetLightPointsAt() bool {
	return s_.RespondsToSelector(objc.Sel("setLightPointsAt:"))
}

// The x and y position that the spotlight points at. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228746-lightpointsat?language=objc
func (s_ SpotLightWrapper) SetLightPointsAt(value IVector) {
	objc.Call[objc.Void](s_, objc.Sel("setLightPointsAt:"), objc.Ptr(value))
}

func (s_ SpotLightWrapper) HasLightPointsAt() bool {
	return s_.RespondsToSelector(objc.Sel("lightPointsAt"))
}

// The x and y position that the spotlight points at. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228746-lightpointsat?language=objc
func (s_ SpotLightWrapper) LightPointsAt() Vector {
	rv := objc.Call[Vector](s_, objc.Sel("lightPointsAt"))
	return rv
}

func (s_ SpotLightWrapper) HasSetConcentration() bool {
	return s_.RespondsToSelector(objc.Sel("setConcentration:"))
}

// The size of the spotlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228744-concentration?language=objc
func (s_ SpotLightWrapper) SetConcentration(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setConcentration:"), value)
}

func (s_ SpotLightWrapper) HasConcentration() bool {
	return s_.RespondsToSelector(objc.Sel("concentration"))
}

// The size of the spotlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228744-concentration?language=objc
func (s_ SpotLightWrapper) Concentration() float64 {
	rv := objc.Call[float64](s_, objc.Sel("concentration"))
	return rv
}

func (s_ SpotLightWrapper) HasSetLightPosition() bool {
	return s_.RespondsToSelector(objc.Sel("setLightPosition:"))
}

// The x and y position of the spotlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228747-lightposition?language=objc
func (s_ SpotLightWrapper) SetLightPosition(value IVector) {
	objc.Call[objc.Void](s_, objc.Sel("setLightPosition:"), objc.Ptr(value))
}

func (s_ SpotLightWrapper) HasLightPosition() bool {
	return s_.RespondsToSelector(objc.Sel("lightPosition"))
}

// The x and y position of the spotlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228747-lightposition?language=objc
func (s_ SpotLightWrapper) LightPosition() Vector {
	rv := objc.Call[Vector](s_, objc.Sel("lightPosition"))
	return rv
}

func (s_ SpotLightWrapper) HasSetBrightness() bool {
	return s_.RespondsToSelector(objc.Sel("setBrightness:"))
}

// The brightness of the spotlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228742-brightness?language=objc
func (s_ SpotLightWrapper) SetBrightness(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setBrightness:"), value)
}

func (s_ SpotLightWrapper) HasBrightness() bool {
	return s_.RespondsToSelector(objc.Sel("brightness"))
}

// The brightness of the spotlight. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cispotlight/3228742-brightness?language=objc
func (s_ SpotLightWrapper) Brightness() float64 {
	rv := objc.Call[float64](s_, objc.Sel("brightness"))
	return rv
}
