// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a perspective tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivetile?language=objc
type PPerspectiveTile interface {
	// optional
	SetBottomRight(value coregraphics.Point)
	HasSetBottomRight() bool

	// optional
	BottomRight() coregraphics.Point
	HasBottomRight() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetBottomLeft(value coregraphics.Point)
	HasSetBottomLeft() bool

	// optional
	BottomLeft() coregraphics.Point
	HasBottomLeft() bool

	// optional
	SetTopRight(value coregraphics.Point)
	HasSetTopRight() bool

	// optional
	TopRight() coregraphics.Point
	HasTopRight() bool

	// optional
	SetTopLeft(value coregraphics.Point)
	HasSetTopLeft() bool

	// optional
	TopLeft() coregraphics.Point
	HasTopLeft() bool
}

// A concrete type wrapper for the [PPerspectiveTile] protocol.
type PerspectiveTileWrapper struct {
	objc.Object
}

func (p_ PerspectiveTileWrapper) HasSetBottomRight() bool {
	return p_.RespondsToSelector(objc.Sel("setBottomRight:"))
}

// The bottom-right coordinate of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivetile/3228654-bottomright?language=objc
func (p_ PerspectiveTileWrapper) SetBottomRight(value coregraphics.Point) {
	objc.Call[objc.Void](p_, objc.Sel("setBottomRight:"), value)
}

func (p_ PerspectiveTileWrapper) HasBottomRight() bool {
	return p_.RespondsToSelector(objc.Sel("bottomRight"))
}

// The bottom-right coordinate of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivetile/3228654-bottomright?language=objc
func (p_ PerspectiveTileWrapper) BottomRight() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](p_, objc.Sel("bottomRight"))
	return rv
}

func (p_ PerspectiveTileWrapper) HasSetInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivetile/3228655-inputimage?language=objc
func (p_ PerspectiveTileWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](p_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (p_ PerspectiveTileWrapper) HasInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivetile/3228655-inputimage?language=objc
func (p_ PerspectiveTileWrapper) InputImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("inputImage"))
	return rv
}

func (p_ PerspectiveTileWrapper) HasSetBottomLeft() bool {
	return p_.RespondsToSelector(objc.Sel("setBottomLeft:"))
}

// The bottom-left coordinate of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivetile/3228653-bottomleft?language=objc
func (p_ PerspectiveTileWrapper) SetBottomLeft(value coregraphics.Point) {
	objc.Call[objc.Void](p_, objc.Sel("setBottomLeft:"), value)
}

func (p_ PerspectiveTileWrapper) HasBottomLeft() bool {
	return p_.RespondsToSelector(objc.Sel("bottomLeft"))
}

// The bottom-left coordinate of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivetile/3228653-bottomleft?language=objc
func (p_ PerspectiveTileWrapper) BottomLeft() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](p_, objc.Sel("bottomLeft"))
	return rv
}

func (p_ PerspectiveTileWrapper) HasSetTopRight() bool {
	return p_.RespondsToSelector(objc.Sel("setTopRight:"))
}

// The top-right coordinate of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivetile/3228657-topright?language=objc
func (p_ PerspectiveTileWrapper) SetTopRight(value coregraphics.Point) {
	objc.Call[objc.Void](p_, objc.Sel("setTopRight:"), value)
}

func (p_ PerspectiveTileWrapper) HasTopRight() bool {
	return p_.RespondsToSelector(objc.Sel("topRight"))
}

// The top-right coordinate of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivetile/3228657-topright?language=objc
func (p_ PerspectiveTileWrapper) TopRight() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](p_, objc.Sel("topRight"))
	return rv
}

func (p_ PerspectiveTileWrapper) HasSetTopLeft() bool {
	return p_.RespondsToSelector(objc.Sel("setTopLeft:"))
}

// The top-left coordinate of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivetile/3228656-topleft?language=objc
func (p_ PerspectiveTileWrapper) SetTopLeft(value coregraphics.Point) {
	objc.Call[objc.Void](p_, objc.Sel("setTopLeft:"), value)
}

func (p_ PerspectiveTileWrapper) HasTopLeft() bool {
	return p_.RespondsToSelector(objc.Sel("topLeft"))
}

// The top-left coordinate of a tile. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciperspectivetile/3228656-topleft?language=objc
func (p_ PerspectiveTileWrapper) TopLeft() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](p_, objc.Sel("topLeft"))
	return rv
}
