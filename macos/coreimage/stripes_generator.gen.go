// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a stripes generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator?language=objc
type PStripesGenerator interface {
	// optional
	SetWidth(value float64)
	HasSetWidth() bool

	// optional
	Width() float64
	HasWidth() bool

	// optional
	SetSharpness(value float64)
	HasSetSharpness() bool

	// optional
	Sharpness() float64
	HasSharpness() bool

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

// A concrete type wrapper for the [PStripesGenerator] protocol.
type StripesGeneratorWrapper struct {
	objc.Object
}

func (s_ StripesGeneratorWrapper) HasSetWidth() bool {
	return s_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of a stripe. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228765-width?language=objc
func (s_ StripesGeneratorWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setWidth:"), value)
}

func (s_ StripesGeneratorWrapper) HasWidth() bool {
	return s_.RespondsToSelector(objc.Sel("width"))
}

// The width of a stripe. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228765-width?language=objc
func (s_ StripesGeneratorWrapper) Width() float64 {
	rv := objc.Call[float64](s_, objc.Sel("width"))
	return rv
}

func (s_ StripesGeneratorWrapper) HasSetSharpness() bool {
	return s_.RespondsToSelector(objc.Sel("setSharpness:"))
}

// The sharpness of the stripe pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228764-sharpness?language=objc
func (s_ StripesGeneratorWrapper) SetSharpness(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setSharpness:"), value)
}

func (s_ StripesGeneratorWrapper) HasSharpness() bool {
	return s_.RespondsToSelector(objc.Sel("sharpness"))
}

// The sharpness of the stripe pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228764-sharpness?language=objc
func (s_ StripesGeneratorWrapper) Sharpness() float64 {
	rv := objc.Call[float64](s_, objc.Sel("sharpness"))
	return rv
}

func (s_ StripesGeneratorWrapper) HasSetColor0() bool {
	return s_.RespondsToSelector(objc.Sel("setColor0:"))
}

// A color to use for the odd stripes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228762-color0?language=objc
func (s_ StripesGeneratorWrapper) SetColor0(value IColor) {
	objc.Call[objc.Void](s_, objc.Sel("setColor0:"), objc.Ptr(value))
}

func (s_ StripesGeneratorWrapper) HasColor0() bool {
	return s_.RespondsToSelector(objc.Sel("color0"))
}

// A color to use for the odd stripes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228762-color0?language=objc
func (s_ StripesGeneratorWrapper) Color0() Color {
	rv := objc.Call[Color](s_, objc.Sel("color0"))
	return rv
}

func (s_ StripesGeneratorWrapper) HasSetColor1() bool {
	return s_.RespondsToSelector(objc.Sel("setColor1:"))
}

// A color to use for the even stripes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228763-color1?language=objc
func (s_ StripesGeneratorWrapper) SetColor1(value IColor) {
	objc.Call[objc.Void](s_, objc.Sel("setColor1:"), objc.Ptr(value))
}

func (s_ StripesGeneratorWrapper) HasColor1() bool {
	return s_.RespondsToSelector(objc.Sel("color1"))
}

// A color to use for the even stripes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228763-color1?language=objc
func (s_ StripesGeneratorWrapper) Color1() Color {
	rv := objc.Call[Color](s_, objc.Sel("color1"))
	return rv
}

func (s_ StripesGeneratorWrapper) HasSetCenter() bool {
	return s_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the stripe pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228761-center?language=objc
func (s_ StripesGeneratorWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](s_, objc.Sel("setCenter:"), value)
}

func (s_ StripesGeneratorWrapper) HasCenter() bool {
	return s_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the stripe pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistripesgenerator/3228761-center?language=objc
func (s_ StripesGeneratorWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](s_, objc.Sel("center"))
	return rv
}
