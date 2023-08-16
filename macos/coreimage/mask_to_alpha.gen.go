// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a mask-to-alpha filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimasktoalpha?language=objc
type PMaskToAlpha interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool
}

// A concrete type wrapper for the [PMaskToAlpha] protocol.
type MaskToAlphaWrapper struct {
	objc.Object
}

func (m_ MaskToAlphaWrapper) HasSetInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimasktoalpha/3228549-inputimage?language=objc
func (m_ MaskToAlphaWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](m_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (m_ MaskToAlphaWrapper) HasInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimasktoalpha/3228549-inputimage?language=objc
func (m_ MaskToAlphaWrapper) InputImage() Image {
	rv := objc.Call[Image](m_, objc.Sel("inputImage"))
	return rv
}
