// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a sixfold reflected tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisixfoldreflectedtile?language=objc
type PSixfoldReflectedTile interface {
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

// A concrete type wrapper for the [PSixfoldReflectedTile] protocol.
type SixfoldReflectedTileWrapper struct {
	objc.Object
}

func (s_ SixfoldReflectedTileWrapper) HasSetWidth() bool {
	return s_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisixfoldreflectedtile/3228716-width?language=objc
func (s_ SixfoldReflectedTileWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setWidth:"), value)
}

func (s_ SixfoldReflectedTileWrapper) HasWidth() bool {
	return s_.RespondsToSelector(objc.Sel("width"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisixfoldreflectedtile/3228716-width?language=objc
func (s_ SixfoldReflectedTileWrapper) Width() float64 {
	rv := objc.Call[float64](s_, objc.Sel("width"))
	return rv
}

func (s_ SixfoldReflectedTileWrapper) HasSetInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisixfoldreflectedtile/3228715-inputimage?language=objc
func (s_ SixfoldReflectedTileWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](s_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (s_ SixfoldReflectedTileWrapper) HasInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisixfoldreflectedtile/3228715-inputimage?language=objc
func (s_ SixfoldReflectedTileWrapper) InputImage() Image {
	rv := objc.Call[Image](s_, objc.Sel("inputImage"))
	return rv
}

func (s_ SixfoldReflectedTileWrapper) HasSetAngle() bool {
	return s_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisixfoldreflectedtile/3228713-angle?language=objc
func (s_ SixfoldReflectedTileWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](s_, objc.Sel("setAngle:"), value)
}

func (s_ SixfoldReflectedTileWrapper) HasAngle() bool {
	return s_.RespondsToSelector(objc.Sel("angle"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisixfoldreflectedtile/3228713-angle?language=objc
func (s_ SixfoldReflectedTileWrapper) Angle() float64 {
	rv := objc.Call[float64](s_, objc.Sel("angle"))
	return rv
}

func (s_ SixfoldReflectedTileWrapper) HasSetCenter() bool {
	return s_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisixfoldreflectedtile/3228714-center?language=objc
func (s_ SixfoldReflectedTileWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](s_, objc.Sel("setCenter:"), value)
}

func (s_ SixfoldReflectedTileWrapper) HasCenter() bool {
	return s_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisixfoldreflectedtile/3228714-center?language=objc
func (s_ SixfoldReflectedTileWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](s_, objc.Sel("center"))
	return rv
}
