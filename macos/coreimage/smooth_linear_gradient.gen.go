// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a smooth linear gradient filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cismoothlineargradient?language=objc
type PSmoothLinearGradient interface {
	// optional
	SetPoint1(value coregraphics.Point)
	HasSetPoint1() bool

	// optional
	Point1() coregraphics.Point
	HasPoint1() bool

	// optional
	SetPoint0(value coregraphics.Point)
	HasSetPoint0() bool

	// optional
	Point0() coregraphics.Point
	HasPoint0() bool

	// optional
	SetColor0(value Color)
	HasSetColor0() bool

	// optional
	Color0() IColor
	HasColor0() bool

	// optional
	SetColor1(value Color)
	HasSetColor1() bool

	// optional
	Color1() IColor
	HasColor1() bool
}

// A concrete type wrapper for the [PSmoothLinearGradient] protocol.
type SmoothLinearGradientWrapper struct {
	objc.Object
}

func (s_ SmoothLinearGradientWrapper) HasSetPoint1() bool {
	return s_.RespondsToSelector(objc.Sel("setPoint1:"))
}

// The ending position of the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cismoothlineargradient/3228726-point1?language=objc
func (s_ SmoothLinearGradientWrapper) SetPoint1(value coregraphics.Point) {
	objc.Call[objc.Void](s_, objc.Sel("setPoint1:"), value)
}

func (s_ SmoothLinearGradientWrapper) HasPoint1() bool {
	return s_.RespondsToSelector(objc.Sel("point1"))
}

// The ending position of the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cismoothlineargradient/3228726-point1?language=objc
func (s_ SmoothLinearGradientWrapper) Point1() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](s_, objc.Sel("point1"))
	return rv
}

func (s_ SmoothLinearGradientWrapper) HasSetPoint0() bool {
	return s_.RespondsToSelector(objc.Sel("setPoint0:"))
}

// The starting position of the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cismoothlineargradient/3228725-point0?language=objc
func (s_ SmoothLinearGradientWrapper) SetPoint0(value coregraphics.Point) {
	objc.Call[objc.Void](s_, objc.Sel("setPoint0:"), value)
}

func (s_ SmoothLinearGradientWrapper) HasPoint0() bool {
	return s_.RespondsToSelector(objc.Sel("point0"))
}

// The starting position of the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cismoothlineargradient/3228725-point0?language=objc
func (s_ SmoothLinearGradientWrapper) Point0() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](s_, objc.Sel("point0"))
	return rv
}

func (s_ SmoothLinearGradientWrapper) HasSetColor0() bool {
	return s_.RespondsToSelector(objc.Sel("setColor0:"))
}

// The first color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cismoothlineargradient/3228723-color0?language=objc
func (s_ SmoothLinearGradientWrapper) SetColor0(value IColor) {
	objc.Call[objc.Void](s_, objc.Sel("setColor0:"), objc.Ptr(value))
}

func (s_ SmoothLinearGradientWrapper) HasColor0() bool {
	return s_.RespondsToSelector(objc.Sel("color0"))
}

// The first color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cismoothlineargradient/3228723-color0?language=objc
func (s_ SmoothLinearGradientWrapper) Color0() Color {
	rv := objc.Call[Color](s_, objc.Sel("color0"))
	return rv
}

func (s_ SmoothLinearGradientWrapper) HasSetColor1() bool {
	return s_.RespondsToSelector(objc.Sel("setColor1:"))
}

// The second color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cismoothlineargradient/3228724-color1?language=objc
func (s_ SmoothLinearGradientWrapper) SetColor1(value IColor) {
	objc.Call[objc.Void](s_, objc.Sel("setColor1:"), objc.Ptr(value))
}

func (s_ SmoothLinearGradientWrapper) HasColor1() bool {
	return s_.RespondsToSelector(objc.Sel("color1"))
}

// The second color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cismoothlineargradient/3228724-color1?language=objc
func (s_ SmoothLinearGradientWrapper) Color1() Color {
	rv := objc.Call[Color](s_, objc.Sel("color1"))
	return rv
}
