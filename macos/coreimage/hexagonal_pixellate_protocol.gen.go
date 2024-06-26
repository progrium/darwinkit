// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/macos/coregraphics"
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a hexagonal pixellate filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihexagonalpixellate?language=objc
type PHexagonalPixellate interface {
	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool

	// optional
	SetScale(value float32)
	HasSetScale() bool

	// optional
	Scale() float32
	HasScale() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool
}

// ensure impl type implements protocol interface
var _ PHexagonalPixellate = (*HexagonalPixellateObject)(nil)

// A concrete type for the [PHexagonalPixellate] protocol.
type HexagonalPixellateObject struct {
	objc.Object
}

func (h_ HexagonalPixellateObject) HasSetCenter() bool {
	return h_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihexagonalpixellate/3228490-center?language=objc
func (h_ HexagonalPixellateObject) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](h_, objc.Sel("setCenter:"), value)
}

func (h_ HexagonalPixellateObject) HasCenter() bool {
	return h_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihexagonalpixellate/3228490-center?language=objc
func (h_ HexagonalPixellateObject) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](h_, objc.Sel("center"))
	return rv
}

func (h_ HexagonalPixellateObject) HasSetScale() bool {
	return h_.RespondsToSelector(objc.Sel("setScale:"))
}

// The size of the hexagons. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihexagonalpixellate/3228492-scale?language=objc
func (h_ HexagonalPixellateObject) SetScale(value float32) {
	objc.Call[objc.Void](h_, objc.Sel("setScale:"), value)
}

func (h_ HexagonalPixellateObject) HasScale() bool {
	return h_.RespondsToSelector(objc.Sel("scale"))
}

// The size of the hexagons. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihexagonalpixellate/3228492-scale?language=objc
func (h_ HexagonalPixellateObject) Scale() float32 {
	rv := objc.Call[float32](h_, objc.Sel("scale"))
	return rv
}

func (h_ HexagonalPixellateObject) HasSetInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihexagonalpixellate/3228491-inputimage?language=objc
func (h_ HexagonalPixellateObject) SetInputImage(value Image) {
	objc.Call[objc.Void](h_, objc.Sel("setInputImage:"), value)
}

func (h_ HexagonalPixellateObject) HasInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihexagonalpixellate/3228491-inputimage?language=objc
func (h_ HexagonalPixellateObject) InputImage() Image {
	rv := objc.Call[Image](h_, objc.Sel("inputImage"))
	return rv
}
