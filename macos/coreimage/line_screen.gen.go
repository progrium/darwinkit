// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a line screen filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilinescreen?language=objc
type PLineScreen interface {
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

// A concrete type wrapper for the [PLineScreen] protocol.
type LineScreenWrapper struct {
	objc.Object
}

func (l_ LineScreenWrapper) HasSetWidth() bool {
	return l_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The distance between lines in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilinescreen/3228540-width?language=objc
func (l_ LineScreenWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setWidth:"), value)
}

func (l_ LineScreenWrapper) HasWidth() bool {
	return l_.RespondsToSelector(objc.Sel("width"))
}

// The distance between lines in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilinescreen/3228540-width?language=objc
func (l_ LineScreenWrapper) Width() float64 {
	rv := objc.Call[float64](l_, objc.Sel("width"))
	return rv
}

func (l_ LineScreenWrapper) HasSetInputImage() bool {
	return l_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilinescreen/3228538-inputimage?language=objc
func (l_ LineScreenWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](l_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (l_ LineScreenWrapper) HasInputImage() bool {
	return l_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilinescreen/3228538-inputimage?language=objc
func (l_ LineScreenWrapper) InputImage() Image {
	rv := objc.Call[Image](l_, objc.Sel("inputImage"))
	return rv
}

func (l_ LineScreenWrapper) HasSetSharpness() bool {
	return l_.RespondsToSelector(objc.Sel("setSharpness:"))
}

// The sharpness of the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilinescreen/3228539-sharpness?language=objc
func (l_ LineScreenWrapper) SetSharpness(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setSharpness:"), value)
}

func (l_ LineScreenWrapper) HasSharpness() bool {
	return l_.RespondsToSelector(objc.Sel("sharpness"))
}

// The sharpness of the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilinescreen/3228539-sharpness?language=objc
func (l_ LineScreenWrapper) Sharpness() float64 {
	rv := objc.Call[float64](l_, objc.Sel("sharpness"))
	return rv
}

func (l_ LineScreenWrapper) HasSetAngle() bool {
	return l_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle of the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilinescreen/3228536-angle?language=objc
func (l_ LineScreenWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setAngle:"), value)
}

func (l_ LineScreenWrapper) HasAngle() bool {
	return l_.RespondsToSelector(objc.Sel("angle"))
}

// The angle of the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilinescreen/3228536-angle?language=objc
func (l_ LineScreenWrapper) Angle() float64 {
	rv := objc.Call[float64](l_, objc.Sel("angle"))
	return rv
}

func (l_ LineScreenWrapper) HasSetCenter() bool {
	return l_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the line screen pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilinescreen/3228537-center?language=objc
func (l_ LineScreenWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](l_, objc.Sel("setCenter:"), value)
}

func (l_ LineScreenWrapper) HasCenter() bool {
	return l_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the line screen pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilinescreen/3228537-center?language=objc
func (l_ LineScreenWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](l_, objc.Sel("center"))
	return rv
}
