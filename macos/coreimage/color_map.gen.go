// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a color map filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormap?language=objc
type PColorMap interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetGradientImage(value Image)
	HasSetGradientImage() bool

	// optional
	GradientImage() IImage
	HasGradientImage() bool
}

// A concrete type wrapper for the [PColorMap] protocol.
type ColorMapWrapper struct {
	objc.Object
}

func (c_ ColorMapWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormap/3228158-inputimage?language=objc
func (c_ ColorMapWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ ColorMapWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormap/3228158-inputimage?language=objc
func (c_ ColorMapWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ ColorMapWrapper) HasSetGradientImage() bool {
	return c_.RespondsToSelector(objc.Sel("setGradientImage:"))
}

// The image data that transforms the source image values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormap/3228157-gradientimage?language=objc
func (c_ ColorMapWrapper) SetGradientImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setGradientImage:"), objc.Ptr(value))
}

func (c_ ColorMapWrapper) HasGradientImage() bool {
	return c_.RespondsToSelector(objc.Sel("gradientImage"))
}

// The image data that transforms the source image values. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolormap/3228157-gradientimage?language=objc
func (c_ ColorMapWrapper) GradientImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("gradientImage"))
	return rv
}
