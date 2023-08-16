// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure an affine tile filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaffinetile?language=objc
type PAffineTile interface {
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

// A concrete type wrapper for the [PAffineTile] protocol.
type AffineTileWrapper struct {
	objc.Object
}

func (a_ AffineTileWrapper) HasSetInputImage() bool {
	return a_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaffinetile/3228057-inputimage?language=objc
func (a_ AffineTileWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](a_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (a_ AffineTileWrapper) HasInputImage() bool {
	return a_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaffinetile/3228057-inputimage?language=objc
func (a_ AffineTileWrapper) InputImage() Image {
	rv := objc.Call[Image](a_, objc.Sel("inputImage"))
	return rv
}

func (a_ AffineTileWrapper) HasSetTransform() bool {
	return a_.RespondsToSelector(objc.Sel("setTransform:"))
}

// The transform to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaffinetile/3228058-transform?language=objc
func (a_ AffineTileWrapper) SetTransform(value coregraphics.AffineTransform) {
	objc.Call[objc.Void](a_, objc.Sel("setTransform:"), value)
}

func (a_ AffineTileWrapper) HasTransform() bool {
	return a_.RespondsToSelector(objc.Sel("transform"))
}

// The transform to apply to the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciaffinetile/3228058-transform?language=objc
func (a_ AffineTileWrapper) Transform() coregraphics.AffineTransform {
	rv := objc.Call[coregraphics.AffineTransform](a_, objc.Sel("transform"))
	return rv
}
