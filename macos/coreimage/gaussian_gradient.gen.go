// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a Gaussian gradient filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussiangradient?language=objc
type PGaussianGradient interface {
	// optional
	SetRadius(value float64)
	HasSetRadius() bool

	// optional
	Radius() float64
	HasRadius() bool

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

// A concrete type wrapper for the [PGaussianGradient] protocol.
type GaussianGradientWrapper struct {
	objc.Object
}

func (g_ GaussianGradientWrapper) HasSetRadius() bool {
	return g_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius of the Gaussian distribution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussiangradient/3228470-radius?language=objc
func (g_ GaussianGradientWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setRadius:"), value)
}

func (g_ GaussianGradientWrapper) HasRadius() bool {
	return g_.RespondsToSelector(objc.Sel("radius"))
}

// The radius of the Gaussian distribution. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussiangradient/3228470-radius?language=objc
func (g_ GaussianGradientWrapper) Radius() float64 {
	rv := objc.Call[float64](g_, objc.Sel("radius"))
	return rv
}

func (g_ GaussianGradientWrapper) HasSetColor0() bool {
	return g_.RespondsToSelector(objc.Sel("setColor0:"))
}

// The first color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussiangradient/3228468-color0?language=objc
func (g_ GaussianGradientWrapper) SetColor0(value IColor) {
	objc.Call[objc.Void](g_, objc.Sel("setColor0:"), objc.Ptr(value))
}

func (g_ GaussianGradientWrapper) HasColor0() bool {
	return g_.RespondsToSelector(objc.Sel("color0"))
}

// The first color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussiangradient/3228468-color0?language=objc
func (g_ GaussianGradientWrapper) Color0() Color {
	rv := objc.Call[Color](g_, objc.Sel("color0"))
	return rv
}

func (g_ GaussianGradientWrapper) HasSetColor1() bool {
	return g_.RespondsToSelector(objc.Sel("setColor1:"))
}

// The second color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussiangradient/3228469-color1?language=objc
func (g_ GaussianGradientWrapper) SetColor1(value IColor) {
	objc.Call[objc.Void](g_, objc.Sel("setColor1:"), objc.Ptr(value))
}

func (g_ GaussianGradientWrapper) HasColor1() bool {
	return g_.RespondsToSelector(objc.Sel("color1"))
}

// The second color to use in the gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussiangradient/3228469-color1?language=objc
func (g_ GaussianGradientWrapper) Color1() Color {
	rv := objc.Call[Color](g_, objc.Sel("color1"))
	return rv
}

func (g_ GaussianGradientWrapper) HasSetCenter() bool {
	return g_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The center of the effect as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussiangradient/3228467-center?language=objc
func (g_ GaussianGradientWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](g_, objc.Sel("setCenter:"), value)
}

func (g_ GaussianGradientWrapper) HasCenter() bool {
	return g_.RespondsToSelector(objc.Sel("center"))
}

// The center of the effect as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaussiangradient/3228467-center?language=objc
func (g_ GaussianGradientWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](g_, objc.Sel("center"))
	return rv
}
