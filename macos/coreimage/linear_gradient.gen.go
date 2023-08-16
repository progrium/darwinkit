// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a linear gradient filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineargradient?language=objc
type PLinearGradient interface {
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

// A concrete type wrapper for the [PLinearGradient] protocol.
type LinearGradientWrapper struct {
	objc.Object
}

func (l_ LinearGradientWrapper) HasSetPoint1() bool {
	return l_.RespondsToSelector(objc.Sel("setPoint1:"))
}

// The ending position of the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineargradient/3228545-point1?language=objc
func (l_ LinearGradientWrapper) SetPoint1(value coregraphics.Point) {
	objc.Call[objc.Void](l_, objc.Sel("setPoint1:"), value)
}

func (l_ LinearGradientWrapper) HasPoint1() bool {
	return l_.RespondsToSelector(objc.Sel("point1"))
}

// The ending position of the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineargradient/3228545-point1?language=objc
func (l_ LinearGradientWrapper) Point1() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](l_, objc.Sel("point1"))
	return rv
}

func (l_ LinearGradientWrapper) HasSetPoint0() bool {
	return l_.RespondsToSelector(objc.Sel("setPoint0:"))
}

// The starting position of the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineargradient/3228544-point0?language=objc
func (l_ LinearGradientWrapper) SetPoint0(value coregraphics.Point) {
	objc.Call[objc.Void](l_, objc.Sel("setPoint0:"), value)
}

func (l_ LinearGradientWrapper) HasPoint0() bool {
	return l_.RespondsToSelector(objc.Sel("point0"))
}

// The starting position of the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineargradient/3228544-point0?language=objc
func (l_ LinearGradientWrapper) Point0() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](l_, objc.Sel("point0"))
	return rv
}

func (l_ LinearGradientWrapper) HasSetColor0() bool {
	return l_.RespondsToSelector(objc.Sel("setColor0:"))
}

// The first color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineargradient/3228542-color0?language=objc
func (l_ LinearGradientWrapper) SetColor0(value IColor) {
	objc.Call[objc.Void](l_, objc.Sel("setColor0:"), objc.Ptr(value))
}

func (l_ LinearGradientWrapper) HasColor0() bool {
	return l_.RespondsToSelector(objc.Sel("color0"))
}

// The first color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineargradient/3228542-color0?language=objc
func (l_ LinearGradientWrapper) Color0() Color {
	rv := objc.Call[Color](l_, objc.Sel("color0"))
	return rv
}

func (l_ LinearGradientWrapper) HasSetColor1() bool {
	return l_.RespondsToSelector(objc.Sel("setColor1:"))
}

// The second color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineargradient/3228543-color1?language=objc
func (l_ LinearGradientWrapper) SetColor1(value IColor) {
	objc.Call[objc.Void](l_, objc.Sel("setColor1:"), objc.Ptr(value))
}

func (l_ LinearGradientWrapper) HasColor1() bool {
	return l_.RespondsToSelector(objc.Sel("color1"))
}

// The second color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilineargradient/3228543-color1?language=objc
func (l_ LinearGradientWrapper) Color1() Color {
	rv := objc.Call[Color](l_, objc.Sel("color1"))
	return rv
}
