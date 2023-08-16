// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a hexagonal pixellate filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihexagonalpixellate?language=objc
type PHexagonalPixellate interface {
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
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PHexagonalPixellate] protocol.
type HexagonalPixellateWrapper struct {
	objc.Object
}

func (h_ HexagonalPixellateWrapper) HasSetScale() bool {
	return h_.RespondsToSelector(objc.Sel("setScale:"))
}

// The size of the hexagons. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihexagonalpixellate/3228492-scale?language=objc
func (h_ HexagonalPixellateWrapper) SetScale(value float64) {
	objc.Call[objc.Void](h_, objc.Sel("setScale:"), value)
}

func (h_ HexagonalPixellateWrapper) HasScale() bool {
	return h_.RespondsToSelector(objc.Sel("scale"))
}

// The size of the hexagons. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihexagonalpixellate/3228492-scale?language=objc
func (h_ HexagonalPixellateWrapper) Scale() float64 {
	rv := objc.Call[float64](h_, objc.Sel("scale"))
	return rv
}

func (h_ HexagonalPixellateWrapper) HasSetInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihexagonalpixellate/3228491-inputimage?language=objc
func (h_ HexagonalPixellateWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](h_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (h_ HexagonalPixellateWrapper) HasInputImage() bool {
	return h_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihexagonalpixellate/3228491-inputimage?language=objc
func (h_ HexagonalPixellateWrapper) InputImage() Image {
	rv := objc.Call[Image](h_, objc.Sel("inputImage"))
	return rv
}

func (h_ HexagonalPixellateWrapper) HasSetCenter() bool {
	return h_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihexagonalpixellate/3228490-center?language=objc
func (h_ HexagonalPixellateWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](h_, objc.Sel("setCenter:"), value)
}

func (h_ HexagonalPixellateWrapper) HasCenter() bool {
	return h_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihexagonalpixellate/3228490-center?language=objc
func (h_ HexagonalPixellateWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](h_, objc.Sel("center"))
	return rv
}
