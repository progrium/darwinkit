// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a affine clamp filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaffineclamp?language=objc
type PAffineClamp interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetTransform(value coregraphics.AffineTransform)
	HasSetTransform() bool

	// optional
	Transform() coregraphics.AffineTransform
	HasTransform() bool
}

// A concrete type wrapper for the [PAffineClamp] protocol.
type AffineClampWrapper struct {
	objc.Object
}

func (a_ AffineClampWrapper) HasSetInputImage() bool {
	return a_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaffineclamp/3228054-inputimage?language=objc
func (a_ AffineClampWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](a_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (a_ AffineClampWrapper) HasInputImage() bool {
	return a_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaffineclamp/3228054-inputimage?language=objc
func (a_ AffineClampWrapper) InputImage() Image {
	rv := objc.Call[Image](a_, objc.Sel("inputImage"))
	return rv
}

func (a_ AffineClampWrapper) HasSetTransform() bool {
	return a_.RespondsToSelector(objc.Sel("setTransform:"))
}

// The transform to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaffineclamp/3228055-transform?language=objc
func (a_ AffineClampWrapper) SetTransform(value coregraphics.AffineTransform) {
	objc.Call[objc.Void](a_, objc.Sel("setTransform:"), value)
}

func (a_ AffineClampWrapper) HasTransform() bool {
	return a_.RespondsToSelector(objc.Sel("transform"))
}

// The transform to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaffineclamp/3228055-transform?language=objc
func (a_ AffineClampWrapper) Transform() coregraphics.AffineTransform {
	rv := objc.Call[coregraphics.AffineTransform](a_, objc.Sel("transform"))
	return rv
}
