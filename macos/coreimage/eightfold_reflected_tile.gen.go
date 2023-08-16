// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure an eightfold reflected tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cieightfoldreflectedtile?language=objc
type PEightfoldReflectedTile interface {
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

// A concrete type wrapper for the [PEightfoldReflectedTile] protocol.
type EightfoldReflectedTileWrapper struct {
	objc.Object
}

func (e_ EightfoldReflectedTileWrapper) HasSetWidth() bool {
	return e_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cieightfoldreflectedtile/3228251-width?language=objc
func (e_ EightfoldReflectedTileWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setWidth:"), value)
}

func (e_ EightfoldReflectedTileWrapper) HasWidth() bool {
	return e_.RespondsToSelector(objc.Sel("width"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cieightfoldreflectedtile/3228251-width?language=objc
func (e_ EightfoldReflectedTileWrapper) Width() float64 {
	rv := objc.Call[float64](e_, objc.Sel("width"))
	return rv
}

func (e_ EightfoldReflectedTileWrapper) HasSetInputImage() bool {
	return e_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cieightfoldreflectedtile/3228250-inputimage?language=objc
func (e_ EightfoldReflectedTileWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](e_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (e_ EightfoldReflectedTileWrapper) HasInputImage() bool {
	return e_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cieightfoldreflectedtile/3228250-inputimage?language=objc
func (e_ EightfoldReflectedTileWrapper) InputImage() Image {
	rv := objc.Call[Image](e_, objc.Sel("inputImage"))
	return rv
}

func (e_ EightfoldReflectedTileWrapper) HasSetAngle() bool {
	return e_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cieightfoldreflectedtile/3228248-angle?language=objc
func (e_ EightfoldReflectedTileWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](e_, objc.Sel("setAngle:"), value)
}

func (e_ EightfoldReflectedTileWrapper) HasAngle() bool {
	return e_.RespondsToSelector(objc.Sel("angle"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cieightfoldreflectedtile/3228248-angle?language=objc
func (e_ EightfoldReflectedTileWrapper) Angle() float64 {
	rv := objc.Call[float64](e_, objc.Sel("angle"))
	return rv
}

func (e_ EightfoldReflectedTileWrapper) HasSetCenter() bool {
	return e_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cieightfoldreflectedtile/3228249-center?language=objc
func (e_ EightfoldReflectedTileWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](e_, objc.Sel("setCenter:"), value)
}

func (e_ EightfoldReflectedTileWrapper) HasCenter() bool {
	return e_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cieightfoldreflectedtile/3228249-center?language=objc
func (e_ EightfoldReflectedTileWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](e_, objc.Sel("center"))
	return rv
}
