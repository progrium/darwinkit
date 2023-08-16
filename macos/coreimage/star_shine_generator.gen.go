// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a star-shine generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator?language=objc
type PStarShineGenerator interface {
	// optional
	SetCrossOpacity(value float64)
	HasSetCrossOpacity() bool

	// optional
	CrossOpacity() float64
	HasCrossOpacity() bool

	// optional
	SetColor(value Color)
	HasSetColor() bool

	// optional
	Color() IColor
	HasColor() bool

	// optional
	SetCrossAngle(value float64)
	HasSetCrossAngle() bool

	// optional
	CrossAngle() float64
	HasCrossAngle() bool

	// optional
	SetEpsilon(value float64)
	HasSetEpsilon() bool

	// optional
	Epsilon() float64
	HasEpsilon() bool

	// optional
	SetRadius(value float64)
	HasSetRadius() bool

	// optional
	Radius() float64
	HasRadius() bool

	// optional
	SetCrossScale(value float64)
	HasSetCrossScale() bool

	// optional
	CrossScale() float64
	HasCrossScale() bool

	// optional
	SetCrossWidth(value float64)
	HasSetCrossWidth() bool

	// optional
	CrossWidth() float64
	HasCrossWidth() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PStarShineGenerator] protocol.
type StarShineGeneratorWrapper struct {
	objc.Object
}

func (s_ StarShineGeneratorWrapper) HasSetCrossOpacity() bool {
	return s_.RespondsToSelector(objc.Sel("setCrossOpacity:"))
}

// The opacity of the cross pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228752-crossopacity?language=objc
func (s_ StarShineGeneratorWrapper) SetCrossOpacity(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setCrossOpacity:"), value)
}

func (s_ StarShineGeneratorWrapper) HasCrossOpacity() bool {
	return s_.RespondsToSelector(objc.Sel("crossOpacity"))
}

// The opacity of the cross pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228752-crossopacity?language=objc
func (s_ StarShineGeneratorWrapper) CrossOpacity() float64 {
	rv := objc.Call[float64](s_, objc.Sel("crossOpacity"))
	return rv
}

func (s_ StarShineGeneratorWrapper) HasSetColor() bool {
	return s_.RespondsToSelector(objc.Sel("setColor:"))
}

// The color to use for the outer shell of the circular star. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228750-color?language=objc
func (s_ StarShineGeneratorWrapper) SetColor(value IColor) {
	objc.Call[objc.Void](s_, objc.Sel("setColor:"), objc.Ptr(value))
}

func (s_ StarShineGeneratorWrapper) HasColor() bool {
	return s_.RespondsToSelector(objc.Sel("color"))
}

// The color to use for the outer shell of the circular star. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228750-color?language=objc
func (s_ StarShineGeneratorWrapper) Color() Color {
	rv := objc.Call[Color](s_, objc.Sel("color"))
	return rv
}

func (s_ StarShineGeneratorWrapper) HasSetCrossAngle() bool {
	return s_.RespondsToSelector(objc.Sel("setCrossAngle:"))
}

// The angle of the cross pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228751-crossangle?language=objc
func (s_ StarShineGeneratorWrapper) SetCrossAngle(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setCrossAngle:"), value)
}

func (s_ StarShineGeneratorWrapper) HasCrossAngle() bool {
	return s_.RespondsToSelector(objc.Sel("crossAngle"))
}

// The angle of the cross pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228751-crossangle?language=objc
func (s_ StarShineGeneratorWrapper) CrossAngle() float64 {
	rv := objc.Call[float64](s_, objc.Sel("crossAngle"))
	return rv
}

func (s_ StarShineGeneratorWrapper) HasSetEpsilon() bool {
	return s_.RespondsToSelector(objc.Sel("setEpsilon:"))
}

// The length of the cross spikes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228755-epsilon?language=objc
func (s_ StarShineGeneratorWrapper) SetEpsilon(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setEpsilon:"), value)
}

func (s_ StarShineGeneratorWrapper) HasEpsilon() bool {
	return s_.RespondsToSelector(objc.Sel("epsilon"))
}

// The length of the cross spikes. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228755-epsilon?language=objc
func (s_ StarShineGeneratorWrapper) Epsilon() float64 {
	rv := objc.Call[float64](s_, objc.Sel("epsilon"))
	return rv
}

func (s_ StarShineGeneratorWrapper) HasSetRadius() bool {
	return s_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius of the star. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228756-radius?language=objc
func (s_ StarShineGeneratorWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setRadius:"), value)
}

func (s_ StarShineGeneratorWrapper) HasRadius() bool {
	return s_.RespondsToSelector(objc.Sel("radius"))
}

// The radius of the star. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228756-radius?language=objc
func (s_ StarShineGeneratorWrapper) Radius() float64 {
	rv := objc.Call[float64](s_, objc.Sel("radius"))
	return rv
}

func (s_ StarShineGeneratorWrapper) HasSetCrossScale() bool {
	return s_.RespondsToSelector(objc.Sel("setCrossScale:"))
}

// The size of the cross pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228753-crossscale?language=objc
func (s_ StarShineGeneratorWrapper) SetCrossScale(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setCrossScale:"), value)
}

func (s_ StarShineGeneratorWrapper) HasCrossScale() bool {
	return s_.RespondsToSelector(objc.Sel("crossScale"))
}

// The size of the cross pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228753-crossscale?language=objc
func (s_ StarShineGeneratorWrapper) CrossScale() float64 {
	rv := objc.Call[float64](s_, objc.Sel("crossScale"))
	return rv
}

func (s_ StarShineGeneratorWrapper) HasSetCrossWidth() bool {
	return s_.RespondsToSelector(objc.Sel("setCrossWidth:"))
}

// The width of the cross pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228754-crosswidth?language=objc
func (s_ StarShineGeneratorWrapper) SetCrossWidth(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setCrossWidth:"), value)
}

func (s_ StarShineGeneratorWrapper) HasCrossWidth() bool {
	return s_.RespondsToSelector(objc.Sel("crossWidth"))
}

// The width of the cross pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228754-crosswidth?language=objc
func (s_ StarShineGeneratorWrapper) CrossWidth() float64 {
	rv := objc.Call[float64](s_, objc.Sel("crossWidth"))
	return rv
}

func (s_ StarShineGeneratorWrapper) HasSetCenter() bool {
	return s_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the star. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228749-center?language=objc
func (s_ StarShineGeneratorWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](s_, objc.Sel("setCenter:"), value)
}

func (s_ StarShineGeneratorWrapper) HasCenter() bool {
	return s_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the star. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cistarshinegenerator/3228749-center?language=objc
func (s_ StarShineGeneratorWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](s_, objc.Sel("center"))
	return rv
}
