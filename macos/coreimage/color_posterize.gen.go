// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a color posterize filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorposterize?language=objc
type PColorPosterize interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetLevels(value float64)
	HasSetLevels() bool

	// optional
	Levels() float64
	HasLevels() bool
}

// A concrete type wrapper for the [PColorPosterize] protocol.
type ColorPosterizeWrapper struct {
	objc.Object
}

func (c_ ColorPosterizeWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorposterize/3228177-inputimage?language=objc
func (c_ ColorPosterizeWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ ColorPosterizeWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorposterize/3228177-inputimage?language=objc
func (c_ ColorPosterizeWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ ColorPosterizeWrapper) HasSetLevels() bool {
	return c_.RespondsToSelector(objc.Sel("setLevels:"))
}

// The number of brightness levels to use for each color component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorposterize/3228178-levels?language=objc
func (c_ ColorPosterizeWrapper) SetLevels(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setLevels:"), value)
}

func (c_ ColorPosterizeWrapper) HasLevels() bool {
	return c_.RespondsToSelector(objc.Sel("levels"))
}

// The number of brightness levels to use for each color component. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicolorposterize/3228178-levels?language=objc
func (c_ ColorPosterizeWrapper) Levels() float64 {
	rv := objc.Call[float64](c_, objc.Sel("levels"))
	return rv
}
