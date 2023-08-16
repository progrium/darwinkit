// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a triangle kaleidoscope filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citrianglekaleidoscope?language=objc
type PTriangleKaleidoscope interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetDecay(value float64)
	HasSetDecay() bool

	// optional
	Decay() float64
	HasDecay() bool

	// optional
	SetRotation(value float64)
	HasSetRotation() bool

	// optional
	Rotation() float64
	HasRotation() bool

	// optional
	SetPoint(value coregraphics.Point)
	HasSetPoint() bool

	// optional
	Point() coregraphics.Point
	HasPoint() bool

	// optional
	SetSize(value float64)
	HasSetSize() bool

	// optional
	Size() float64
	HasSize() bool
}

// A concrete type wrapper for the [PTriangleKaleidoscope] protocol.
type TriangleKaleidoscopeWrapper struct {
	objc.Object
}

func (t_ TriangleKaleidoscopeWrapper) HasSetInputImage() bool {
	return t_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citrianglekaleidoscope/3228804-inputimage?language=objc
func (t_ TriangleKaleidoscopeWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](t_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (t_ TriangleKaleidoscopeWrapper) HasInputImage() bool {
	return t_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citrianglekaleidoscope/3228804-inputimage?language=objc
func (t_ TriangleKaleidoscopeWrapper) InputImage() Image {
	rv := objc.Call[Image](t_, objc.Sel("inputImage"))
	return rv
}

func (t_ TriangleKaleidoscopeWrapper) HasSetDecay() bool {
	return t_.RespondsToSelector(objc.Sel("setDecay:"))
}

// A value that determines how fast the color fades from the center triangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citrianglekaleidoscope/3228803-decay?language=objc
func (t_ TriangleKaleidoscopeWrapper) SetDecay(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setDecay:"), value)
}

func (t_ TriangleKaleidoscopeWrapper) HasDecay() bool {
	return t_.RespondsToSelector(objc.Sel("decay"))
}

// A value that determines how fast the color fades from the center triangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citrianglekaleidoscope/3228803-decay?language=objc
func (t_ TriangleKaleidoscopeWrapper) Decay() float64 {
	rv := objc.Call[float64](t_, objc.Sel("decay"))
	return rv
}

func (t_ TriangleKaleidoscopeWrapper) HasSetRotation() bool {
	return t_.RespondsToSelector(objc.Sel("setRotation:"))
}

// The rotation angle of the triangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citrianglekaleidoscope/3228806-rotation?language=objc
func (t_ TriangleKaleidoscopeWrapper) SetRotation(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setRotation:"), value)
}

func (t_ TriangleKaleidoscopeWrapper) HasRotation() bool {
	return t_.RespondsToSelector(objc.Sel("rotation"))
}

// The rotation angle of the triangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citrianglekaleidoscope/3228806-rotation?language=objc
func (t_ TriangleKaleidoscopeWrapper) Rotation() float64 {
	rv := objc.Call[float64](t_, objc.Sel("rotation"))
	return rv
}

func (t_ TriangleKaleidoscopeWrapper) HasSetPoint() bool {
	return t_.RespondsToSelector(objc.Sel("setPoint:"))
}

// The x and y position to use as the center of the triangular area in the input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citrianglekaleidoscope/3228805-point?language=objc
func (t_ TriangleKaleidoscopeWrapper) SetPoint(value coregraphics.Point) {
	objc.Call[objc.Void](t_, objc.Sel("setPoint:"), value)
}

func (t_ TriangleKaleidoscopeWrapper) HasPoint() bool {
	return t_.RespondsToSelector(objc.Sel("point"))
}

// The x and y position to use as the center of the triangular area in the input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citrianglekaleidoscope/3228805-point?language=objc
func (t_ TriangleKaleidoscopeWrapper) Point() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](t_, objc.Sel("point"))
	return rv
}

func (t_ TriangleKaleidoscopeWrapper) HasSetSize() bool {
	return t_.RespondsToSelector(objc.Sel("setSize:"))
}

// The size, in pixels, of the triangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citrianglekaleidoscope/3228807-size?language=objc
func (t_ TriangleKaleidoscopeWrapper) SetSize(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setSize:"), value)
}

func (t_ TriangleKaleidoscopeWrapper) HasSize() bool {
	return t_.RespondsToSelector(objc.Sel("size"))
}

// The size, in pixels, of the triangle. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citrianglekaleidoscope/3228807-size?language=objc
func (t_ TriangleKaleidoscopeWrapper) Size() float64 {
	rv := objc.Call[float64](t_, objc.Sel("size"))
	return rv
}
