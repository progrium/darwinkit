// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a flash transition filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciflashtransition?language=objc
type PFlashTransition interface {
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
	SetStriationStrength(value float64)
	HasSetStriationStrength() bool

	// optional
	StriationStrength() float64
	HasStriationStrength() bool

	// optional
	SetFadeThreshold(value float64)
	HasSetFadeThreshold() bool

	// optional
	FadeThreshold() float64
	HasFadeThreshold() bool

	// optional
	SetExtent(value coregraphics.Rect)
	HasSetExtent() bool

	// optional
	Extent() coregraphics.Rect
	HasExtent() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PFlashTransition] protocol.
type FlashTransitionWrapper struct {
	objc.Object
}

func (f_ FlashTransitionWrapper) HasSetStriationContrast() bool {
	return f_.RespondsToSelector(objc.Sel("setStriationContrast:"))
}

// The contrast of the light rays emanating from the flash. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciflashtransition/3228441-striationcontrast?language=objc
func (f_ FlashTransitionWrapper) SetStriationContrast(value float64) {
	objc.Call[objc.Void](f_, objc.Sel("setStriationContrast:"), value)
}

func (f_ FlashTransitionWrapper) HasStriationContrast() bool {
	return f_.RespondsToSelector(objc.Sel("striationContrast"))
}

// The contrast of the light rays emanating from the flash. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciflashtransition/3228441-striationcontrast?language=objc
func (f_ FlashTransitionWrapper) StriationContrast() float64 {
	rv := objc.Call[float64](f_, objc.Sel("striationContrast"))
	return rv
}

func (f_ FlashTransitionWrapper) HasSetColor() bool {
	return f_.RespondsToSelector(objc.Sel("setColor:"))
}

// The color of the light rays emanating from the flash. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciflashtransition/3228437-color?language=objc
func (f_ FlashTransitionWrapper) SetColor(value IColor) {
	objc.Call[objc.Void](f_, objc.Sel("setColor:"), objc.Ptr(value))
}

func (f_ FlashTransitionWrapper) HasColor() bool {
	return f_.RespondsToSelector(objc.Sel("color"))
}

// The color of the light rays emanating from the flash. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciflashtransition/3228437-color?language=objc
func (f_ FlashTransitionWrapper) Color() Color {
	rv := objc.Call[Color](f_, objc.Sel("color"))
	return rv
}

func (f_ FlashTransitionWrapper) HasSetMaxStriationRadius() bool {
	return f_.RespondsToSelector(objc.Sel("setMaxStriationRadius:"))
}

// The radius of the light rays emanating from the flash. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciflashtransition/3228440-maxstriationradius?language=objc
func (f_ FlashTransitionWrapper) SetMaxStriationRadius(value float64) {
	objc.Call[objc.Void](f_, objc.Sel("setMaxStriationRadius:"), value)
}

func (f_ FlashTransitionWrapper) HasMaxStriationRadius() bool {
	return f_.RespondsToSelector(objc.Sel("maxStriationRadius"))
}

// The radius of the light rays emanating from the flash. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciflashtransition/3228440-maxstriationradius?language=objc
func (f_ FlashTransitionWrapper) MaxStriationRadius() float64 {
	rv := objc.Call[float64](f_, objc.Sel("maxStriationRadius"))
	return rv
}

func (f_ FlashTransitionWrapper) HasSetStriationStrength() bool {
	return f_.RespondsToSelector(objc.Sel("setStriationStrength:"))
}

// The strength of the light rays emanating from the flash. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciflashtransition/3228442-striationstrength?language=objc
func (f_ FlashTransitionWrapper) SetStriationStrength(value float64) {
	objc.Call[objc.Void](f_, objc.Sel("setStriationStrength:"), value)
}

func (f_ FlashTransitionWrapper) HasStriationStrength() bool {
	return f_.RespondsToSelector(objc.Sel("striationStrength"))
}

// The strength of the light rays emanating from the flash. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciflashtransition/3228442-striationstrength?language=objc
func (f_ FlashTransitionWrapper) StriationStrength() float64 {
	rv := objc.Call[float64](f_, objc.Sel("striationStrength"))
	return rv
}

func (f_ FlashTransitionWrapper) HasSetFadeThreshold() bool {
	return f_.RespondsToSelector(objc.Sel("setFadeThreshold:"))
}

// The amount of fade between the flash and the target image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciflashtransition/3228439-fadethreshold?language=objc
func (f_ FlashTransitionWrapper) SetFadeThreshold(value float64) {
	objc.Call[objc.Void](f_, objc.Sel("setFadeThreshold:"), value)
}

func (f_ FlashTransitionWrapper) HasFadeThreshold() bool {
	return f_.RespondsToSelector(objc.Sel("fadeThreshold"))
}

// The amount of fade between the flash and the target image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciflashtransition/3228439-fadethreshold?language=objc
func (f_ FlashTransitionWrapper) FadeThreshold() float64 {
	rv := objc.Call[float64](f_, objc.Sel("fadeThreshold"))
	return rv
}

func (f_ FlashTransitionWrapper) HasSetExtent() bool {
	return f_.RespondsToSelector(objc.Sel("setExtent:"))
}

// The extent of the flash. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciflashtransition/3228438-extent?language=objc
func (f_ FlashTransitionWrapper) SetExtent(value coregraphics.Rect) {
	objc.Call[objc.Void](f_, objc.Sel("setExtent:"), value)
}

func (f_ FlashTransitionWrapper) HasExtent() bool {
	return f_.RespondsToSelector(objc.Sel("extent"))
}

// The extent of the flash. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciflashtransition/3228438-extent?language=objc
func (f_ FlashTransitionWrapper) Extent() coregraphics.Rect {
	rv := objc.Call[coregraphics.Rect](f_, objc.Sel("extent"))
	return rv
}

func (f_ FlashTransitionWrapper) HasSetCenter() bool {
	return f_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciflashtransition/3228436-center?language=objc
func (f_ FlashTransitionWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](f_, objc.Sel("setCenter:"), value)
}

func (f_ FlashTransitionWrapper) HasCenter() bool {
	return f_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciflashtransition/3228436-center?language=objc
func (f_ FlashTransitionWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](f_, objc.Sel("center"))
	return rv
}
