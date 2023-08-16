// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a Gabor gradients filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaborgradients?language=objc
type PGaborGradients interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool
}

// A concrete type wrapper for the [PGaborGradients] protocol.
type GaborGradientsWrapper struct {
	objc.Object
}

func (g_ GaborGradientsWrapper) HasSetInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaborgradients/3325514-inputimage?language=objc
func (g_ GaborGradientsWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](g_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (g_ GaborGradientsWrapper) HasInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigaborgradients/3325514-inputimage?language=objc
func (g_ GaborGradientsWrapper) InputImage() Image {
	rv := objc.Call[Image](g_, objc.Sel("inputImage"))
	return rv
}
