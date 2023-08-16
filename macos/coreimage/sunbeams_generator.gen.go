// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a sunbeams generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator?language=objc
type PSunbeamsGenerator interface {
	// optional
	SetStriationContrast(value float64)
	HasSetStriationContrast() bool

	// optional
	StriationContrast() float64
	HasStriationContrast() bool

	// optional
	SetColor(value Color)
	HasSetColor() bool

	// optional
	Color() IColor
	HasColor() bool

	// optional
	SetMaxStriationRadius(value float64)
	HasSetMaxStriationRadius() bool

	// optional
	MaxStriationRadius() float64
	HasMaxStriationRadius() bool

	// optional
	SetTime(value float64)
	HasSetTime() bool

	// optional
	Time() float64
	HasTime() bool

	// optional
	SetStriationStrength(value float64)
	HasSetStriationStrength() bool

	// optional
	StriationStrength() float64
	HasStriationStrength() bool

	// optional
	SetSunRadius(value float64)
	HasSetSunRadius() bool

	// optional
	SunRadius() float64
	HasSunRadius() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PSunbeamsGenerator] protocol.
type SunbeamsGeneratorWrapper struct {
	objc.Object
}

func (s_ SunbeamsGeneratorWrapper) HasSetStriationContrast() bool {
	return s_.RespondsToSelector(objc.Sel("setStriationContrast:"))
}

// The contrast of the sunbeams. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228770-striationcontrast?language=objc
func (s_ SunbeamsGeneratorWrapper) SetStriationContrast(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setStriationContrast:"), value)
}

func (s_ SunbeamsGeneratorWrapper) HasStriationContrast() bool {
	return s_.RespondsToSelector(objc.Sel("striationContrast"))
}

// The contrast of the sunbeams. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228770-striationcontrast?language=objc
func (s_ SunbeamsGeneratorWrapper) StriationContrast() float64 {
	rv := objc.Call[float64](s_, objc.Sel("striationContrast"))
	return rv
}

func (s_ SunbeamsGeneratorWrapper) HasSetColor() bool {
	return s_.RespondsToSelector(objc.Sel("setColor:"))
}

// The color of the sun. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228768-color?language=objc
func (s_ SunbeamsGeneratorWrapper) SetColor(value IColor) {
	objc.Call[objc.Void](s_, objc.Sel("setColor:"), objc.Ptr(value))
}

func (s_ SunbeamsGeneratorWrapper) HasColor() bool {
	return s_.RespondsToSelector(objc.Sel("color"))
}

// The color of the sun. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228768-color?language=objc
func (s_ SunbeamsGeneratorWrapper) Color() Color {
	rv := objc.Call[Color](s_, objc.Sel("color"))
	return rv
}

func (s_ SunbeamsGeneratorWrapper) HasSetMaxStriationRadius() bool {
	return s_.RespondsToSelector(objc.Sel("setMaxStriationRadius:"))
}

// The radius of the sunbeams. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228769-maxstriationradius?language=objc
func (s_ SunbeamsGeneratorWrapper) SetMaxStriationRadius(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setMaxStriationRadius:"), value)
}

func (s_ SunbeamsGeneratorWrapper) HasMaxStriationRadius() bool {
	return s_.RespondsToSelector(objc.Sel("maxStriationRadius"))
}

// The radius of the sunbeams. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228769-maxstriationradius?language=objc
func (s_ SunbeamsGeneratorWrapper) MaxStriationRadius() float64 {
	rv := objc.Call[float64](s_, objc.Sel("maxStriationRadius"))
	return rv
}

func (s_ SunbeamsGeneratorWrapper) HasSetTime() bool {
	return s_.RespondsToSelector(objc.Sel("setTime:"))
}

// The duration of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228773-time?language=objc
func (s_ SunbeamsGeneratorWrapper) SetTime(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setTime:"), value)
}

func (s_ SunbeamsGeneratorWrapper) HasTime() bool {
	return s_.RespondsToSelector(objc.Sel("time"))
}

// The duration of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228773-time?language=objc
func (s_ SunbeamsGeneratorWrapper) Time() float64 {
	rv := objc.Call[float64](s_, objc.Sel("time"))
	return rv
}

func (s_ SunbeamsGeneratorWrapper) HasSetStriationStrength() bool {
	return s_.RespondsToSelector(objc.Sel("setStriationStrength:"))
}

// The intensity of the sunbeams. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228771-striationstrength?language=objc
func (s_ SunbeamsGeneratorWrapper) SetStriationStrength(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setStriationStrength:"), value)
}

func (s_ SunbeamsGeneratorWrapper) HasStriationStrength() bool {
	return s_.RespondsToSelector(objc.Sel("striationStrength"))
}

// The intensity of the sunbeams. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228771-striationstrength?language=objc
func (s_ SunbeamsGeneratorWrapper) StriationStrength() float64 {
	rv := objc.Call[float64](s_, objc.Sel("striationStrength"))
	return rv
}

func (s_ SunbeamsGeneratorWrapper) HasSetSunRadius() bool {
	return s_.RespondsToSelector(objc.Sel("setSunRadius:"))
}

// The radius of the sun. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228772-sunradius?language=objc
func (s_ SunbeamsGeneratorWrapper) SetSunRadius(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setSunRadius:"), value)
}

func (s_ SunbeamsGeneratorWrapper) HasSunRadius() bool {
	return s_.RespondsToSelector(objc.Sel("sunRadius"))
}

// The radius of the sun. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228772-sunradius?language=objc
func (s_ SunbeamsGeneratorWrapper) SunRadius() float64 {
	rv := objc.Call[float64](s_, objc.Sel("sunRadius"))
	return rv
}

func (s_ SunbeamsGeneratorWrapper) HasSetCenter() bool {
	return s_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the sunbeam pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228767-center?language=objc
func (s_ SunbeamsGeneratorWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](s_, objc.Sel("setCenter:"), value)
}

func (s_ SunbeamsGeneratorWrapper) HasCenter() bool {
	return s_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the sunbeam pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisunbeamsgenerator/3228767-center?language=objc
func (s_ SunbeamsGeneratorWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](s_, objc.Sel("center"))
	return rv
}
