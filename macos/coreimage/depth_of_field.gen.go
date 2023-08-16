// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a depth-of-field filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidepthoffield?language=objc
type PDepthOfField interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetSaturation(value float64)
	HasSetSaturation() bool

	// optional
	Saturation() float64
	HasSaturation() bool

	// optional
	SetUnsharpMaskIntensity(value float64)
	HasSetUnsharpMaskIntensity() bool

	// optional
	UnsharpMaskIntensity() float64
	HasUnsharpMaskIntensity() bool

	// optional
	SetRadius(value float64)
	HasSetRadius() bool

	// optional
	Radius() float64
	HasRadius() bool

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
	SetUnsharpMaskRadius(value float64)
	HasSetUnsharpMaskRadius() bool

	// optional
	UnsharpMaskRadius() float64
	HasUnsharpMaskRadius() bool
}

// A concrete type wrapper for the [PDepthOfField] protocol.
type DepthOfFieldWrapper struct {
	objc.Object
}

func (d_ DepthOfFieldWrapper) HasSetInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidepthoffield/3228204-inputimage?language=objc
func (d_ DepthOfFieldWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](d_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (d_ DepthOfFieldWrapper) HasInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidepthoffield/3228204-inputimage?language=objc
func (d_ DepthOfFieldWrapper) InputImage() Image {
	rv := objc.Call[Image](d_, objc.Sel("inputImage"))
	return rv
}

func (d_ DepthOfFieldWrapper) HasSetSaturation() bool {
	return d_.RespondsToSelector(objc.Sel("setSaturation:"))
}

// The amount to adjust the saturation by. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidepthoffield/3228208-saturation?language=objc
func (d_ DepthOfFieldWrapper) SetSaturation(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setSaturation:"), value)
}

func (d_ DepthOfFieldWrapper) HasSaturation() bool {
	return d_.RespondsToSelector(objc.Sel("saturation"))
}

// The amount to adjust the saturation by. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidepthoffield/3228208-saturation?language=objc
func (d_ DepthOfFieldWrapper) Saturation() float64 {
	rv := objc.Call[float64](d_, objc.Sel("saturation"))
	return rv
}

func (d_ DepthOfFieldWrapper) HasSetUnsharpMaskIntensity() bool {
	return d_.RespondsToSelector(objc.Sel("setUnsharpMaskIntensity:"))
}

// The intensity of the unsharp mask effect applied to the in-focus area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidepthoffield/3228209-unsharpmaskintensity?language=objc
func (d_ DepthOfFieldWrapper) SetUnsharpMaskIntensity(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setUnsharpMaskIntensity:"), value)
}

func (d_ DepthOfFieldWrapper) HasUnsharpMaskIntensity() bool {
	return d_.RespondsToSelector(objc.Sel("unsharpMaskIntensity"))
}

// The intensity of the unsharp mask effect applied to the in-focus area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidepthoffield/3228209-unsharpmaskintensity?language=objc
func (d_ DepthOfFieldWrapper) UnsharpMaskIntensity() float64 {
	rv := objc.Call[float64](d_, objc.Sel("unsharpMaskIntensity"))
	return rv
}

func (d_ DepthOfFieldWrapper) HasSetRadius() bool {
	return d_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The distance from the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidepthoffield/3228207-radius?language=objc
func (d_ DepthOfFieldWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setRadius:"), value)
}

func (d_ DepthOfFieldWrapper) HasRadius() bool {
	return d_.RespondsToSelector(objc.Sel("radius"))
}

// The distance from the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidepthoffield/3228207-radius?language=objc
func (d_ DepthOfFieldWrapper) Radius() float64 {
	rv := objc.Call[float64](d_, objc.Sel("radius"))
	return rv
}

func (d_ DepthOfFieldWrapper) HasSetPoint1() bool {
	return d_.RespondsToSelector(objc.Sel("setPoint1:"))
}

// The second point in the focused region of the output image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidepthoffield/3228206-point1?language=objc
func (d_ DepthOfFieldWrapper) SetPoint1(value coregraphics.Point) {
	objc.Call[objc.Void](d_, objc.Sel("setPoint1:"), value)
}

func (d_ DepthOfFieldWrapper) HasPoint1() bool {
	return d_.RespondsToSelector(objc.Sel("point1"))
}

// The second point in the focused region of the output image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidepthoffield/3228206-point1?language=objc
func (d_ DepthOfFieldWrapper) Point1() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](d_, objc.Sel("point1"))
	return rv
}

func (d_ DepthOfFieldWrapper) HasSetPoint0() bool {
	return d_.RespondsToSelector(objc.Sel("setPoint0:"))
}

// The first point in the focused region of the output image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidepthoffield/3228205-point0?language=objc
func (d_ DepthOfFieldWrapper) SetPoint0(value coregraphics.Point) {
	objc.Call[objc.Void](d_, objc.Sel("setPoint0:"), value)
}

func (d_ DepthOfFieldWrapper) HasPoint0() bool {
	return d_.RespondsToSelector(objc.Sel("point0"))
}

// The first point in the focused region of the output image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidepthoffield/3228205-point0?language=objc
func (d_ DepthOfFieldWrapper) Point0() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](d_, objc.Sel("point0"))
	return rv
}

func (d_ DepthOfFieldWrapper) HasSetUnsharpMaskRadius() bool {
	return d_.RespondsToSelector(objc.Sel("setUnsharpMaskRadius:"))
}

// The radius of the unsharp mask effect applied to the in-focus area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidepthoffield/3228210-unsharpmaskradius?language=objc
func (d_ DepthOfFieldWrapper) SetUnsharpMaskRadius(value float64) {
	objc.Call[objc.Void](d_, objc.Sel("setUnsharpMaskRadius:"), value)
}

func (d_ DepthOfFieldWrapper) HasUnsharpMaskRadius() bool {
	return d_.RespondsToSelector(objc.Sel("unsharpMaskRadius"))
}

// The radius of the unsharp mask effect applied to the in-focus area. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidepthoffield/3228210-unsharpmaskradius?language=objc
func (d_ DepthOfFieldWrapper) UnsharpMaskRadius() float64 {
	rv := objc.Call[float64](d_, objc.Sel("unsharpMaskRadius"))
	return rv
}
