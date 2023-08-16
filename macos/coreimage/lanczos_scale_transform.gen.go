// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a Lanczos scale transform filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilanczosscaletransform?language=objc
type PLanczosScaleTransform interface {
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
	SetAspectRatio(value float64)
	HasSetAspectRatio() bool

	// optional
	AspectRatio() float64
	HasAspectRatio() bool
}

// A concrete type wrapper for the [PLanczosScaleTransform] protocol.
type LanczosScaleTransformWrapper struct {
	objc.Object
}

func (l_ LanczosScaleTransformWrapper) HasSetScale() bool {
	return l_.RespondsToSelector(objc.Sel("setScale:"))
}

// The scaling factor to use on the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilanczosscaletransform/3228518-scale?language=objc
func (l_ LanczosScaleTransformWrapper) SetScale(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setScale:"), value)
}

func (l_ LanczosScaleTransformWrapper) HasScale() bool {
	return l_.RespondsToSelector(objc.Sel("scale"))
}

// The scaling factor to use on the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilanczosscaletransform/3228518-scale?language=objc
func (l_ LanczosScaleTransformWrapper) Scale() float64 {
	rv := objc.Call[float64](l_, objc.Sel("scale"))
	return rv
}

func (l_ LanczosScaleTransformWrapper) HasSetInputImage() bool {
	return l_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilanczosscaletransform/3228517-inputimage?language=objc
func (l_ LanczosScaleTransformWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](l_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (l_ LanczosScaleTransformWrapper) HasInputImage() bool {
	return l_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilanczosscaletransform/3228517-inputimage?language=objc
func (l_ LanczosScaleTransformWrapper) InputImage() Image {
	rv := objc.Call[Image](l_, objc.Sel("inputImage"))
	return rv
}

func (l_ LanczosScaleTransformWrapper) HasSetAspectRatio() bool {
	return l_.RespondsToSelector(objc.Sel("setAspectRatio:"))
}

// The additional horizontal scaling factor to use on the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilanczosscaletransform/3228516-aspectratio?language=objc
func (l_ LanczosScaleTransformWrapper) SetAspectRatio(value float64) {
	objc.Call[objc.Void](l_, objc.Sel("setAspectRatio:"), value)
}

func (l_ LanczosScaleTransformWrapper) HasAspectRatio() bool {
	return l_.RespondsToSelector(objc.Sel("aspectRatio"))
}

// The additional horizontal scaling factor to use on the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilanczosscaletransform/3228516-aspectratio?language=objc
func (l_ LanczosScaleTransformWrapper) AspectRatio() float64 {
	rv := objc.Call[float64](l_, objc.Sel("aspectRatio"))
	return rv
}
