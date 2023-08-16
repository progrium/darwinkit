// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a color invert filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorinvert?language=objc
type PColorInvert interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool
}

// A concrete type wrapper for the [PColorInvert] protocol.
type ColorInvertWrapper struct {
	objc.Object
}

func (c_ ColorInvertWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorinvert/3228155-inputimage?language=objc
func (c_ ColorInvertWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ ColorInvertWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorinvert/3228155-inputimage?language=objc
func (c_ ColorInvertWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}
