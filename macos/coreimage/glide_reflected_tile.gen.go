// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a glide reflected tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglidereflectedtile?language=objc
type PGlideReflectedTile interface {
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

// A concrete type wrapper for the [PGlideReflectedTile] protocol.
type GlideReflectedTileWrapper struct {
	objc.Object
}

func (g_ GlideReflectedTileWrapper) HasSetWidth() bool {
	return g_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglidereflectedtile/3228475-width?language=objc
func (g_ GlideReflectedTileWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setWidth:"), value)
}

func (g_ GlideReflectedTileWrapper) HasWidth() bool {
	return g_.RespondsToSelector(objc.Sel("width"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglidereflectedtile/3228475-width?language=objc
func (g_ GlideReflectedTileWrapper) Width() float64 {
	rv := objc.Call[float64](g_, objc.Sel("width"))
	return rv
}

func (g_ GlideReflectedTileWrapper) HasSetInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglidereflectedtile/3228474-inputimage?language=objc
func (g_ GlideReflectedTileWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](g_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (g_ GlideReflectedTileWrapper) HasInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglidereflectedtile/3228474-inputimage?language=objc
func (g_ GlideReflectedTileWrapper) InputImage() Image {
	rv := objc.Call[Image](g_, objc.Sel("inputImage"))
	return rv
}

func (g_ GlideReflectedTileWrapper) HasSetAngle() bool {
	return g_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglidereflectedtile/3228472-angle?language=objc
func (g_ GlideReflectedTileWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setAngle:"), value)
}

func (g_ GlideReflectedTileWrapper) HasAngle() bool {
	return g_.RespondsToSelector(objc.Sel("angle"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglidereflectedtile/3228472-angle?language=objc
func (g_ GlideReflectedTileWrapper) Angle() float64 {
	rv := objc.Call[float64](g_, objc.Sel("angle"))
	return rv
}

func (g_ GlideReflectedTileWrapper) HasSetCenter() bool {
	return g_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglidereflectedtile/3228473-center?language=objc
func (g_ GlideReflectedTileWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](g_, objc.Sel("setCenter:"), value)
}

func (g_ GlideReflectedTileWrapper) HasCenter() bool {
	return g_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciglidereflectedtile/3228473-center?language=objc
func (g_ GlideReflectedTileWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](g_, objc.Sel("center"))
	return rv
}
