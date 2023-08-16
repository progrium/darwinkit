// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a lenticular halo generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilenticularhalogenerator?language=objc
type PLenticularHaloGenerator interface {
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
	SetHaloRadius(value float64)
	HasSetHaloRadius() bool

	// optional
	HaloRadius() float64
	HasHaloRadius() bool

	// optional
	SetHaloWidth(value float64)
	HasSetHaloWidth() bool

	// optional
	HaloWidth() float64
	HasHaloWidth() bool

	// optional
	SetTime(value float64)
	HasSetTime() bool

	// optional
	Time() float64
	HasTime() bool

	// optional
	SetHaloOverlap(value float64)
	HasSetHaloOverlap() bool

	// optional
	HaloOverlap() float64
	HasHaloOverlap() bool

	// optional
	SetStriationStrength(value float64)
	HasSetStriationStrength() bool

	// optional
	StriationStrength() float64
	HasStriationStrength() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PLenticularHaloGenerator] protocol.
type LenticularHaloGeneratorWrapper struct {
	objc.Object
}

func (l_ LenticularHaloGeneratorWrapper) HasSetStriationContrast() bool {
	return l_.RespondsToSelector(objc.Sel("setStriationContrast:"))
}

// The contrast of the halo colors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilenticularhalogenerator/3228525-striationcontrast?language=objc
func (l_ LenticularHaloGeneratorWrapper) SetStriationContrast(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setStriationContrast:"), value)
}

func (l_ LenticularHaloGeneratorWrapper) HasStriationContrast() bool {
	return l_.RespondsToSelector(objc.Sel("striationContrast"))
}

// The contrast of the halo colors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilenticularhalogenerator/3228525-striationcontrast?language=objc
func (l_ LenticularHaloGeneratorWrapper) StriationContrast() float64 {
	rv := objc.Call[float64](l_, objc.Sel("striationContrast"))
	return rv
}

func (l_ LenticularHaloGeneratorWrapper) HasSetColor() bool {
	return l_.RespondsToSelector(objc.Sel("setColor:"))
}

// The color of the halo. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilenticularhalogenerator/3228521-color?language=objc
func (l_ LenticularHaloGeneratorWrapper) SetColor(value IColor) {
	objc.Call[objc.Void](l_, objc.Sel("setColor:"), objc.Ptr(value))
}

func (l_ LenticularHaloGeneratorWrapper) HasColor() bool {
	return l_.RespondsToSelector(objc.Sel("color"))
}

// The color of the halo. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilenticularhalogenerator/3228521-color?language=objc
func (l_ LenticularHaloGeneratorWrapper) Color() Color {
	rv := objc.Call[Color](l_, objc.Sel("color"))
	return rv
}

func (l_ LenticularHaloGeneratorWrapper) HasSetHaloRadius() bool {
	return l_.RespondsToSelector(objc.Sel("setHaloRadius:"))
}

// The radius of the halo. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilenticularhalogenerator/3228523-haloradius?language=objc
func (l_ LenticularHaloGeneratorWrapper) SetHaloRadius(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setHaloRadius:"), value)
}

func (l_ LenticularHaloGeneratorWrapper) HasHaloRadius() bool {
	return l_.RespondsToSelector(objc.Sel("haloRadius"))
}

// The radius of the halo. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilenticularhalogenerator/3228523-haloradius?language=objc
func (l_ LenticularHaloGeneratorWrapper) HaloRadius() float64 {
	rv := objc.Call[float64](l_, objc.Sel("haloRadius"))
	return rv
}

func (l_ LenticularHaloGeneratorWrapper) HasSetHaloWidth() bool {
	return l_.RespondsToSelector(objc.Sel("setHaloWidth:"))
}

// The width of the halo, from its inner radius to its outer radius. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilenticularhalogenerator/3228524-halowidth?language=objc
func (l_ LenticularHaloGeneratorWrapper) SetHaloWidth(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setHaloWidth:"), value)
}

func (l_ LenticularHaloGeneratorWrapper) HasHaloWidth() bool {
	return l_.RespondsToSelector(objc.Sel("haloWidth"))
}

// The width of the halo, from its inner radius to its outer radius. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilenticularhalogenerator/3228524-halowidth?language=objc
func (l_ LenticularHaloGeneratorWrapper) HaloWidth() float64 {
	rv := objc.Call[float64](l_, objc.Sel("haloWidth"))
	return rv
}

func (l_ LenticularHaloGeneratorWrapper) HasSetTime() bool {
	return l_.RespondsToSelector(objc.Sel("setTime:"))
}

// The current time of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilenticularhalogenerator/3228527-time?language=objc
func (l_ LenticularHaloGeneratorWrapper) SetTime(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setTime:"), value)
}

func (l_ LenticularHaloGeneratorWrapper) HasTime() bool {
	return l_.RespondsToSelector(objc.Sel("time"))
}

// The current time of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilenticularhalogenerator/3228527-time?language=objc
func (l_ LenticularHaloGeneratorWrapper) Time() float64 {
	rv := objc.Call[float64](l_, objc.Sel("time"))
	return rv
}

func (l_ LenticularHaloGeneratorWrapper) HasSetHaloOverlap() bool {
	return l_.RespondsToSelector(objc.Sel("setHaloOverlap:"))
}

// The separation of colors in the halo. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilenticularhalogenerator/3228522-halooverlap?language=objc
func (l_ LenticularHaloGeneratorWrapper) SetHaloOverlap(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setHaloOverlap:"), value)
}

func (l_ LenticularHaloGeneratorWrapper) HasHaloOverlap() bool {
	return l_.RespondsToSelector(objc.Sel("haloOverlap"))
}

// The separation of colors in the halo. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilenticularhalogenerator/3228522-halooverlap?language=objc
func (l_ LenticularHaloGeneratorWrapper) HaloOverlap() float64 {
	rv := objc.Call[float64](l_, objc.Sel("haloOverlap"))
	return rv
}

func (l_ LenticularHaloGeneratorWrapper) HasSetStriationStrength() bool {
	return l_.RespondsToSelector(objc.Sel("setStriationStrength:"))
}

// The intensity of the halo colors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilenticularhalogenerator/3228526-striationstrength?language=objc
func (l_ LenticularHaloGeneratorWrapper) SetStriationStrength(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setStriationStrength:"), value)
}

func (l_ LenticularHaloGeneratorWrapper) HasStriationStrength() bool {
	return l_.RespondsToSelector(objc.Sel("striationStrength"))
}

// The intensity of the halo colors. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilenticularhalogenerator/3228526-striationstrength?language=objc
func (l_ LenticularHaloGeneratorWrapper) StriationStrength() float64 {
	rv := objc.Call[float64](l_, objc.Sel("striationStrength"))
	return rv
}

func (l_ LenticularHaloGeneratorWrapper) HasSetCenter() bool {
	return l_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the halo. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilenticularhalogenerator/3228520-center?language=objc
func (l_ LenticularHaloGeneratorWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](l_, objc.Sel("setCenter:"), value)
}

func (l_ LenticularHaloGeneratorWrapper) HasCenter() bool {
	return l_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the halo. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilenticularhalogenerator/3228520-center?language=objc
func (l_ LenticularHaloGeneratorWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](l_, objc.Sel("center"))
	return rv
}
