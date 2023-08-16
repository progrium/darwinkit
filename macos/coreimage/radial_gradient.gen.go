// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a radial gradient filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient?language=objc
type PRadialGradient interface {
	// optional
	SetRadius1(value float64)
	HasSetRadius1() bool

	// optional
	Radius1() float64
	HasRadius1() bool

	// optional
	SetRadius0(value float64)
	HasSetRadius0() bool

	// optional
	Radius0() float64
	HasRadius0() bool

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

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PRadialGradient] protocol.
type RadialGradientWrapper struct {
	objc.Object
}

func (r_ RadialGradientWrapper) HasSetRadius1() bool {
	return r_.RespondsToSelector(objc.Sel("setRadius1:"))
}

// The radius of the ending circle to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228689-radius1?language=objc
func (r_ RadialGradientWrapper) SetRadius1(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setRadius1:"), value)
}

func (r_ RadialGradientWrapper) HasRadius1() bool {
	return r_.RespondsToSelector(objc.Sel("radius1"))
}

// The radius of the ending circle to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228689-radius1?language=objc
func (r_ RadialGradientWrapper) Radius1() float64 {
	rv := objc.Call[float64](r_, objc.Sel("radius1"))
	return rv
}

func (r_ RadialGradientWrapper) HasSetRadius0() bool {
	return r_.RespondsToSelector(objc.Sel("setRadius0:"))
}

// The radius of the starting circle to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228688-radius0?language=objc
func (r_ RadialGradientWrapper) SetRadius0(value float64) {
	objc.Call[objc.Void](r_, objc.Sel("setRadius0:"), value)
}

func (r_ RadialGradientWrapper) HasRadius0() bool {
	return r_.RespondsToSelector(objc.Sel("radius0"))
}

// The radius of the starting circle to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228688-radius0?language=objc
func (r_ RadialGradientWrapper) Radius0() float64 {
	rv := objc.Call[float64](r_, objc.Sel("radius0"))
	return rv
}

func (r_ RadialGradientWrapper) HasSetColor0() bool {
	return r_.RespondsToSelector(objc.Sel("setColor0:"))
}

// The first color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228686-color0?language=objc
func (r_ RadialGradientWrapper) SetColor0(value IColor) {
	objc.Call[objc.Void](r_, objc.Sel("setColor0:"), objc.Ptr(value))
}

func (r_ RadialGradientWrapper) HasColor0() bool {
	return r_.RespondsToSelector(objc.Sel("color0"))
}

// The first color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228686-color0?language=objc
func (r_ RadialGradientWrapper) Color0() Color {
	rv := objc.Call[Color](r_, objc.Sel("color0"))
	return rv
}

func (r_ RadialGradientWrapper) HasSetColor1() bool {
	return r_.RespondsToSelector(objc.Sel("setColor1:"))
}

// The second color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228687-color1?language=objc
func (r_ RadialGradientWrapper) SetColor1(value IColor) {
	objc.Call[objc.Void](r_, objc.Sel("setColor1:"), objc.Ptr(value))
}

func (r_ RadialGradientWrapper) HasColor1() bool {
	return r_.RespondsToSelector(objc.Sel("color1"))
}

// The second color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228687-color1?language=objc
func (r_ RadialGradientWrapper) Color1() Color {
	rv := objc.Call[Color](r_, objc.Sel("color1"))
	return rv
}

func (r_ RadialGradientWrapper) HasSetCenter() bool {
	return r_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The center of the effect as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228685-center?language=objc
func (r_ RadialGradientWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](r_, objc.Sel("setCenter:"), value)
}

func (r_ RadialGradientWrapper) HasCenter() bool {
	return r_.RespondsToSelector(objc.Sel("center"))
}

// The center of the effect as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciradialgradient/3228685-center?language=objc
func (r_ RadialGradientWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](r_, objc.Sel("center"))
	return rv
}
