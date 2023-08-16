// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure an optical tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cioptile?language=objc
type POpTile interface {
	// optional
	SetWidth(value float64)
	HasSetWidth() bool

	// optional
	Width() float64
	HasWidth() bool

	// optional
	SetScale(value float64)
	HasSetScale() bool

	// optional
	Scale() float64
	HasScale() bool

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

// A concrete type wrapper for the [POpTile] protocol.
type OpTileWrapper struct {
	objc.Object
}

func (o_ OpTileWrapper) HasSetWidth() bool {
	return o_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cioptile/3228603-width?language=objc
func (o_ OpTileWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](o_, objc.Sel("setWidth:"), value)
}

func (o_ OpTileWrapper) HasWidth() bool {
	return o_.RespondsToSelector(objc.Sel("width"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cioptile/3228603-width?language=objc
func (o_ OpTileWrapper) Width() float64 {
	rv := objc.Call[float64](o_, objc.Sel("width"))
	return rv
}

func (o_ OpTileWrapper) HasSetScale() bool {
	return o_.RespondsToSelector(objc.Sel("setScale:"))
}

// A value that determines the number of tiles in the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cioptile/3228602-scale?language=objc
func (o_ OpTileWrapper) SetScale(value float64) {
	objc.Call[objc.Void](o_, objc.Sel("setScale:"), value)
}

func (o_ OpTileWrapper) HasScale() bool {
	return o_.RespondsToSelector(objc.Sel("scale"))
}

// A value that determines the number of tiles in the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cioptile/3228602-scale?language=objc
func (o_ OpTileWrapper) Scale() float64 {
	rv := objc.Call[float64](o_, objc.Sel("scale"))
	return rv
}

func (o_ OpTileWrapper) HasSetInputImage() bool {
	return o_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cioptile/3228601-inputimage?language=objc
func (o_ OpTileWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](o_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (o_ OpTileWrapper) HasInputImage() bool {
	return o_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cioptile/3228601-inputimage?language=objc
func (o_ OpTileWrapper) InputImage() Image {
	rv := objc.Call[Image](o_, objc.Sel("inputImage"))
	return rv
}

func (o_ OpTileWrapper) HasSetAngle() bool {
	return o_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cioptile/3228599-angle?language=objc
func (o_ OpTileWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](o_, objc.Sel("setAngle:"), value)
}

func (o_ OpTileWrapper) HasAngle() bool {
	return o_.RespondsToSelector(objc.Sel("angle"))
}

// The angle of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cioptile/3228599-angle?language=objc
func (o_ OpTileWrapper) Angle() float64 {
	rv := objc.Call[float64](o_, objc.Sel("angle"))
	return rv
}

func (o_ OpTileWrapper) HasSetCenter() bool {
	return o_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cioptile/3228600-center?language=objc
func (o_ OpTileWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](o_, objc.Sel("setCenter:"), value)
}

func (o_ OpTileWrapper) HasCenter() bool {
	return o_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cioptile/3228600-center?language=objc
func (o_ OpTileWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](o_, objc.Sel("center"))
	return rv
}
