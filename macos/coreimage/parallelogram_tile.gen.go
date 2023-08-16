// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a parallelogram tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciparallelogramtile?language=objc
type PParallelogramTile interface {
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
	SetAcuteAngle(value float64)
	HasSetAcuteAngle() bool

	// optional
	AcuteAngle() float64
	HasAcuteAngle() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PParallelogramTile] protocol.
type ParallelogramTileWrapper struct {
	objc.Object
}

func (p_ ParallelogramTileWrapper) HasSetWidth() bool {
	return p_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciparallelogramtile/3228644-width?language=objc
func (p_ ParallelogramTileWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setWidth:"), value)
}

func (p_ ParallelogramTileWrapper) HasWidth() bool {
	return p_.RespondsToSelector(objc.Sel("width"))
}

// The width of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciparallelogramtile/3228644-width?language=objc
func (p_ ParallelogramTileWrapper) Width() float64 {
	rv := objc.Call[float64](p_, objc.Sel("width"))
	return rv
}

func (p_ ParallelogramTileWrapper) HasSetInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciparallelogramtile/3228643-inputimage?language=objc
func (p_ ParallelogramTileWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](p_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (p_ ParallelogramTileWrapper) HasInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciparallelogramtile/3228643-inputimage?language=objc
func (p_ ParallelogramTileWrapper) InputImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("inputImage"))
	return rv
}

func (p_ ParallelogramTileWrapper) HasSetAngle() bool {
	return p_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciparallelogramtile/3228641-angle?language=objc
func (p_ ParallelogramTileWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setAngle:"), value)
}

func (p_ ParallelogramTileWrapper) HasAngle() bool {
	return p_.RespondsToSelector(objc.Sel("angle"))
}

// The angle, in radians, of the tiled pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciparallelogramtile/3228641-angle?language=objc
func (p_ ParallelogramTileWrapper) Angle() float64 {
	rv := objc.Call[float64](p_, objc.Sel("angle"))
	return rv
}

func (p_ ParallelogramTileWrapper) HasSetAcuteAngle() bool {
	return p_.RespondsToSelector(objc.Sel("setAcuteAngle:"))
}

// The primary angle for the repeating parallelogram tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciparallelogramtile/3228640-acuteangle?language=objc
func (p_ ParallelogramTileWrapper) SetAcuteAngle(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setAcuteAngle:"), value)
}

func (p_ ParallelogramTileWrapper) HasAcuteAngle() bool {
	return p_.RespondsToSelector(objc.Sel("acuteAngle"))
}

// The primary angle for the repeating parallelogram tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciparallelogramtile/3228640-acuteangle?language=objc
func (p_ ParallelogramTileWrapper) AcuteAngle() float64 {
	rv := objc.Call[float64](p_, objc.Sel("acuteAngle"))
	return rv
}

func (p_ ParallelogramTileWrapper) HasSetCenter() bool {
	return p_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciparallelogramtile/3228642-center?language=objc
func (p_ ParallelogramTileWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](p_, objc.Sel("setCenter:"), value)
}

func (p_ ParallelogramTileWrapper) HasCenter() bool {
	return p_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciparallelogramtile/3228642-center?language=objc
func (p_ ParallelogramTileWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](p_, objc.Sel("center"))
	return rv
}
