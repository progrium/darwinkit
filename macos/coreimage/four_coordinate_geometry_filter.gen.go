// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a geometry adjustment filters that requires four coordinates. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter?language=objc
type PFourCoordinateGeometryFilter interface {
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

// A concrete type wrapper for the [PFourCoordinateGeometryFilter] protocol.
type FourCoordinateGeometryFilterWrapper struct {
	objc.Object
}

func (f_ FourCoordinateGeometryFilterWrapper) HasSetBottomRight() bool {
	return f_.RespondsToSelector(objc.Sel("setBottomRight:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338732-bottomright?language=objc
func (f_ FourCoordinateGeometryFilterWrapper) SetBottomRight(value coregraphics.Point) {
	objc.Call[objc.Void](f_, objc.Sel("setBottomRight:"), value)
}

func (f_ FourCoordinateGeometryFilterWrapper) HasBottomRight() bool {
	return f_.RespondsToSelector(objc.Sel("bottomRight"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338732-bottomright?language=objc
func (f_ FourCoordinateGeometryFilterWrapper) BottomRight() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](f_, objc.Sel("bottomRight"))
	return rv
}

func (f_ FourCoordinateGeometryFilterWrapper) HasSetInputImage() bool {
	return f_.RespondsToSelector(objc.Sel("setInputImage:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338733-inputimage?language=objc
func (f_ FourCoordinateGeometryFilterWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](f_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (f_ FourCoordinateGeometryFilterWrapper) HasInputImage() bool {
	return f_.RespondsToSelector(objc.Sel("inputImage"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338733-inputimage?language=objc
func (f_ FourCoordinateGeometryFilterWrapper) InputImage() Image {
	rv := objc.Call[Image](f_, objc.Sel("inputImage"))
	return rv
}

func (f_ FourCoordinateGeometryFilterWrapper) HasSetBottomLeft() bool {
	return f_.RespondsToSelector(objc.Sel("setBottomLeft:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338731-bottomleft?language=objc
func (f_ FourCoordinateGeometryFilterWrapper) SetBottomLeft(value coregraphics.Point) {
	objc.Call[objc.Void](f_, objc.Sel("setBottomLeft:"), value)
}

func (f_ FourCoordinateGeometryFilterWrapper) HasBottomLeft() bool {
	return f_.RespondsToSelector(objc.Sel("bottomLeft"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338731-bottomleft?language=objc
func (f_ FourCoordinateGeometryFilterWrapper) BottomLeft() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](f_, objc.Sel("bottomLeft"))
	return rv
}

func (f_ FourCoordinateGeometryFilterWrapper) HasSetTopRight() bool {
	return f_.RespondsToSelector(objc.Sel("setTopRight:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338735-topright?language=objc
func (f_ FourCoordinateGeometryFilterWrapper) SetTopRight(value coregraphics.Point) {
	objc.Call[objc.Void](f_, objc.Sel("setTopRight:"), value)
}

func (f_ FourCoordinateGeometryFilterWrapper) HasTopRight() bool {
	return f_.RespondsToSelector(objc.Sel("topRight"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338735-topright?language=objc
func (f_ FourCoordinateGeometryFilterWrapper) TopRight() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](f_, objc.Sel("topRight"))
	return rv
}

func (f_ FourCoordinateGeometryFilterWrapper) HasSetTopLeft() bool {
	return f_.RespondsToSelector(objc.Sel("setTopLeft:"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338734-topleft?language=objc
func (f_ FourCoordinateGeometryFilterWrapper) SetTopLeft(value coregraphics.Point) {
	objc.Call[objc.Void](f_, objc.Sel("setTopLeft:"), value)
}

func (f_ FourCoordinateGeometryFilterWrapper) HasTopLeft() bool {
	return f_.RespondsToSelector(objc.Sel("topLeft"))
}

//	[Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifourcoordinategeometryfilter/3338734-topleft?language=objc
func (f_ FourCoordinateGeometryFilterWrapper) TopLeft() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](f_, objc.Sel("topLeft"))
	return rv
}
