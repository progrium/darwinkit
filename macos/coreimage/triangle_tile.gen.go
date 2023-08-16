// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a triangle tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citriangletile?language=objc
type PTriangleTile interface {
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

// A concrete type wrapper for the [PTriangleTile] protocol.
type TriangleTileWrapper struct {
	objc.Object
}

func (t_ TriangleTileWrapper) HasSetWidth() bool {
	return t_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citriangletile/3228812-width?language=objc
func (t_ TriangleTileWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setWidth:"), value)
}

func (t_ TriangleTileWrapper) HasWidth() bool {
	return t_.RespondsToSelector(objc.Sel("width"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citriangletile/3228812-width?language=objc
func (t_ TriangleTileWrapper) Width() float64 {
	rv := objc.Call[float64](t_, objc.Sel("width"))
	return rv
}

func (t_ TriangleTileWrapper) HasSetInputImage() bool {
	return t_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citriangletile/3228811-inputimage?language=objc
func (t_ TriangleTileWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](t_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (t_ TriangleTileWrapper) HasInputImage() bool {
	return t_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citriangletile/3228811-inputimage?language=objc
func (t_ TriangleTileWrapper) InputImage() Image {
	rv := objc.Call[Image](t_, objc.Sel("inputImage"))
	return rv
}

func (t_ TriangleTileWrapper) HasSetAngle() bool {
	return t_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citriangletile/3228809-angle?language=objc
func (t_ TriangleTileWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setAngle:"), value)
}

func (t_ TriangleTileWrapper) HasAngle() bool {
	return t_.RespondsToSelector(objc.Sel("angle"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citriangletile/3228809-angle?language=objc
func (t_ TriangleTileWrapper) Angle() float64 {
	rv := objc.Call[float64](t_, objc.Sel("angle"))
	return rv
}

func (t_ TriangleTileWrapper) HasSetCenter() bool {
	return t_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citriangletile/3228810-center?language=objc
func (t_ TriangleTileWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](t_, objc.Sel("setCenter:"), value)
}

func (t_ TriangleTileWrapper) HasCenter() bool {
	return t_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citriangletile/3228810-center?language=objc
func (t_ TriangleTileWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](t_, objc.Sel("center"))
	return rv
}
