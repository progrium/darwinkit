// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a hatched screen filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen?language=objc
type PHatchedScreen interface {
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

// A concrete type wrapper for the [PHatchedScreen] protocol.
type HatchedScreenWrapper struct {
	objc.Object
}

func (h_ HatchedScreenWrapper) HasSetWidth() bool {
	return h_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The distance between lines in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228485-width?language=objc
func (h_ HatchedScreenWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](h_, objc.Sel("setWidth:"), value)
}

func (h_ HatchedScreenWrapper) HasWidth() bool {
	return h_.RespondsToSelector(objc.Sel("width"))
}

// The distance between lines in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228485-width?language=objc
func (h_ HatchedScreenWrapper) Width() float64 {
	rv := objc.Call[float64](h_, objc.Sel("width"))
	return rv
}

func (h_ HatchedScreenWrapper) HasSetInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228483-inputimage?language=objc
func (h_ HatchedScreenWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](h_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (h_ HatchedScreenWrapper) HasInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228483-inputimage?language=objc
func (h_ HatchedScreenWrapper) InputImage() Image {
	rv := objc.Call[Image](h_, objc.Sel("inputImage"))
	return rv
}

func (h_ HatchedScreenWrapper) HasSetSharpness() bool {
	return h_.RespondsToSelector(objc.Sel("setSharpness:"))
}

// The amount of sharpening to apply. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228484-sharpness?language=objc
func (h_ HatchedScreenWrapper) SetSharpness(value float64) {
	objc.Call[objc.Void](h_, objc.Sel("setSharpness:"), value)
}

func (h_ HatchedScreenWrapper) HasSharpness() bool {
	return h_.RespondsToSelector(objc.Sel("sharpness"))
}

// The amount of sharpening to apply. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228484-sharpness?language=objc
func (h_ HatchedScreenWrapper) Sharpness() float64 {
	rv := objc.Call[float64](h_, objc.Sel("sharpness"))
	return rv
}

func (h_ HatchedScreenWrapper) HasSetAngle() bool {
	return h_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle of the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228481-angle?language=objc
func (h_ HatchedScreenWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](h_, objc.Sel("setAngle:"), value)
}

func (h_ HatchedScreenWrapper) HasAngle() bool {
	return h_.RespondsToSelector(objc.Sel("angle"))
}

// The angle of the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228481-angle?language=objc
func (h_ HatchedScreenWrapper) Angle() float64 {
	rv := objc.Call[float64](h_, objc.Sel("angle"))
	return rv
}

func (h_ HatchedScreenWrapper) HasSetCenter() bool {
	return h_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the hatched screen pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228482-center?language=objc
func (h_ HatchedScreenWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](h_, objc.Sel("setCenter:"), value)
}

func (h_ HatchedScreenWrapper) HasCenter() bool {
	return h_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the hatched screen pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihatchedscreen/3228482-center?language=objc
func (h_ HatchedScreenWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](h_, objc.Sel("center"))
	return rv
}
