// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a checkerboard generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator?language=objc
type PCheckerboardGenerator interface {
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

// A concrete type wrapper for the [PCheckerboardGenerator] protocol.
type CheckerboardGeneratorWrapper struct {
	objc.Object
}

func (c_ CheckerboardGeneratorWrapper) HasSetWidth() bool {
	return c_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of the squares in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228109-width?language=objc
func (c_ CheckerboardGeneratorWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setWidth:"), value)
}

func (c_ CheckerboardGeneratorWrapper) HasWidth() bool {
	return c_.RespondsToSelector(objc.Sel("width"))
}

// The width of the squares in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228109-width?language=objc
func (c_ CheckerboardGeneratorWrapper) Width() float64 {
	rv := objc.Call[float64](c_, objc.Sel("width"))
	return rv
}

func (c_ CheckerboardGeneratorWrapper) HasSetSharpness() bool {
	return c_.RespondsToSelector(objc.Sel("setSharpness:"))
}

// The sharpness of the edges in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228108-sharpness?language=objc
func (c_ CheckerboardGeneratorWrapper) SetSharpness(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setSharpness:"), value)
}

func (c_ CheckerboardGeneratorWrapper) HasSharpness() bool {
	return c_.RespondsToSelector(objc.Sel("sharpness"))
}

// The sharpness of the edges in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228108-sharpness?language=objc
func (c_ CheckerboardGeneratorWrapper) Sharpness() float64 {
	rv := objc.Call[float64](c_, objc.Sel("sharpness"))
	return rv
}

func (c_ CheckerboardGeneratorWrapper) HasSetColor0() bool {
	return c_.RespondsToSelector(objc.Sel("setColor0:"))
}

// A color to use for the first set of squares. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228106-color0?language=objc
func (c_ CheckerboardGeneratorWrapper) SetColor0(value IColor) {
	objc.Call[objc.Void](c_, objc.Sel("setColor0:"), objc.Ptr(value))
}

func (c_ CheckerboardGeneratorWrapper) HasColor0() bool {
	return c_.RespondsToSelector(objc.Sel("color0"))
}

// A color to use for the first set of squares. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228106-color0?language=objc
func (c_ CheckerboardGeneratorWrapper) Color0() Color {
	rv := objc.Call[Color](c_, objc.Sel("color0"))
	return rv
}

func (c_ CheckerboardGeneratorWrapper) HasSetColor1() bool {
	return c_.RespondsToSelector(objc.Sel("setColor1:"))
}

// A color to use for the second set of squares. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228107-color1?language=objc
func (c_ CheckerboardGeneratorWrapper) SetColor1(value IColor) {
	objc.Call[objc.Void](c_, objc.Sel("setColor1:"), objc.Ptr(value))
}

func (c_ CheckerboardGeneratorWrapper) HasColor1() bool {
	return c_.RespondsToSelector(objc.Sel("color1"))
}

// A color to use for the second set of squares. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228107-color1?language=objc
func (c_ CheckerboardGeneratorWrapper) Color1() Color {
	rv := objc.Call[Color](c_, objc.Sel("color1"))
	return rv
}

func (c_ CheckerboardGeneratorWrapper) HasSetCenter() bool {
	return c_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The center of the effect as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228105-center?language=objc
func (c_ CheckerboardGeneratorWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](c_, objc.Sel("setCenter:"), value)
}

func (c_ CheckerboardGeneratorWrapper) HasCenter() bool {
	return c_.RespondsToSelector(objc.Sel("center"))
}

// The center of the effect as x and y coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicheckerboardgenerator/3228105-center?language=objc
func (c_ CheckerboardGeneratorWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](c_, objc.Sel("center"))
	return rv
}
