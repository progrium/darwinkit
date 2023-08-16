// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a dot screen filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidotscreen?language=objc
type PDotScreen interface {
	// optional
	SetWidth(value float64)
	HasSetWidth() bool

	// optional
	Width() float64
	HasWidth() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetSharpness(value float64)
	HasSetSharpness() bool

	// optional
	Sharpness() float64
	HasSharpness() bool

	// optional
	SetAngle(value float64)
	HasSetAngle() bool

	// optional
	Angle() float64
	HasAngle() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PDotScreen] protocol.
type DotScreenWrapper struct {
	objc.Object
}

func (d_ DotScreenWrapper) HasSetWidth() bool {
	return d_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The distance between dots in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidotscreen/3228235-width?language=objc
func (d_ DotScreenWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setWidth:"), value)
}

func (d_ DotScreenWrapper) HasWidth() bool {
	return d_.RespondsToSelector(objc.Sel("width"))
}

// The distance between dots in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidotscreen/3228235-width?language=objc
func (d_ DotScreenWrapper) Width() float64 {
	rv := objc.Call[float64](d_, objc.Sel("width"))
	return rv
}

func (d_ DotScreenWrapper) HasSetInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidotscreen/3228233-inputimage?language=objc
func (d_ DotScreenWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](d_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (d_ DotScreenWrapper) HasInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidotscreen/3228233-inputimage?language=objc
func (d_ DotScreenWrapper) InputImage() Image {
	rv := objc.Call[Image](d_, objc.Sel("inputImage"))
	return rv
}

func (d_ DotScreenWrapper) HasSetSharpness() bool {
	return d_.RespondsToSelector(objc.Sel("setSharpness:"))
}

// The sharpness of the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidotscreen/3228234-sharpness?language=objc
func (d_ DotScreenWrapper) SetSharpness(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setSharpness:"), value)
}

func (d_ DotScreenWrapper) HasSharpness() bool {
	return d_.RespondsToSelector(objc.Sel("sharpness"))
}

// The sharpness of the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidotscreen/3228234-sharpness?language=objc
func (d_ DotScreenWrapper) Sharpness() float64 {
	rv := objc.Call[float64](d_, objc.Sel("sharpness"))
	return rv
}

func (d_ DotScreenWrapper) HasSetAngle() bool {
	return d_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle of the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidotscreen/3228231-angle?language=objc
func (d_ DotScreenWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setAngle:"), value)
}

func (d_ DotScreenWrapper) HasAngle() bool {
	return d_.RespondsToSelector(objc.Sel("angle"))
}

// The angle of the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidotscreen/3228231-angle?language=objc
func (d_ DotScreenWrapper) Angle() float64 {
	rv := objc.Call[float64](d_, objc.Sel("angle"))
	return rv
}

func (d_ DotScreenWrapper) HasSetCenter() bool {
	return d_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the dot screen pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidotscreen/3228232-center?language=objc
func (d_ DotScreenWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](d_, objc.Sel("setCenter:"), value)
}

func (d_ DotScreenWrapper) HasCenter() bool {
	return d_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the dot screen pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidotscreen/3228232-center?language=objc
func (d_ DotScreenWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](d_, objc.Sel("center"))
	return rv
}
