// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a pixellate filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipixellate?language=objc
type PPixellate interface {
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

// A concrete type wrapper for the [PPixellate] protocol.
type PixellateWrapper struct {
	objc.Object
}

func (p_ PixellateWrapper) HasSetScale() bool {
	return p_.RespondsToSelector(objc.Sel("setScale:"))
}

// A value that determines the size of the squares. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipixellate/3228676-scale?language=objc
func (p_ PixellateWrapper) SetScale(value float64) {
	objc.Call[objc.Void](p_, objc.Sel("setScale:"), value)
}

func (p_ PixellateWrapper) HasScale() bool {
	return p_.RespondsToSelector(objc.Sel("scale"))
}

// A value that determines the size of the squares. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipixellate/3228676-scale?language=objc
func (p_ PixellateWrapper) Scale() float64 {
	rv := objc.Call[float64](p_, objc.Sel("scale"))
	return rv
}

func (p_ PixellateWrapper) HasSetInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipixellate/3228675-inputimage?language=objc
func (p_ PixellateWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](p_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (p_ PixellateWrapper) HasInputImage() bool {
	return p_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipixellate/3228675-inputimage?language=objc
func (p_ PixellateWrapper) InputImage() Image {
	rv := objc.Call[Image](p_, objc.Sel("inputImage"))
	return rv
}

func (p_ PixellateWrapper) HasSetCenter() bool {
	return p_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipixellate/3228674-center?language=objc
func (p_ PixellateWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](p_, objc.Sel("setCenter:"), value)
}

func (p_ PixellateWrapper) HasCenter() bool {
	return p_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cipixellate/3228674-center?language=objc
func (p_ PixellateWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](p_, objc.Sel("center"))
	return rv
}
