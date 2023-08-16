// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a circular screen filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicircularscreen?language=objc
type PCircularScreen interface {
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
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PCircularScreen] protocol.
type CircularScreenWrapper struct {
	objc.Object
}

func (c_ CircularScreenWrapper) HasSetWidth() bool {
	return c_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The distance between each circle in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicircularscreen/3228114-width?language=objc
func (c_ CircularScreenWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setWidth:"), value)
}

func (c_ CircularScreenWrapper) HasWidth() bool {
	return c_.RespondsToSelector(objc.Sel("width"))
}

// The distance between each circle in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicircularscreen/3228114-width?language=objc
func (c_ CircularScreenWrapper) Width() float64 {
	rv := objc.Call[float64](c_, objc.Sel("width"))
	return rv
}

func (c_ CircularScreenWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicircularscreen/3228112-inputimage?language=objc
func (c_ CircularScreenWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ CircularScreenWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicircularscreen/3228112-inputimage?language=objc
func (c_ CircularScreenWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ CircularScreenWrapper) HasSetSharpness() bool {
	return c_.RespondsToSelector(objc.Sel("setSharpness:"))
}

// The sharpness of the circles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicircularscreen/3228113-sharpness?language=objc
func (c_ CircularScreenWrapper) SetSharpness(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setSharpness:"), value)
}

func (c_ CircularScreenWrapper) HasSharpness() bool {
	return c_.RespondsToSelector(objc.Sel("sharpness"))
}

// The sharpness of the circles. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicircularscreen/3228113-sharpness?language=objc
func (c_ CircularScreenWrapper) Sharpness() float64 {
	rv := objc.Call[float64](c_, objc.Sel("sharpness"))
	return rv
}

func (c_ CircularScreenWrapper) HasSetCenter() bool {
	return c_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the circular screen pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicircularscreen/3228111-center?language=objc
func (c_ CircularScreenWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](c_, objc.Sel("setCenter:"), value)
}

func (c_ CircularScreenWrapper) HasCenter() bool {
	return c_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the circular screen pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicircularscreen/3228111-center?language=objc
func (c_ CircularScreenWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](c_, objc.Sel("center"))
	return rv
}
