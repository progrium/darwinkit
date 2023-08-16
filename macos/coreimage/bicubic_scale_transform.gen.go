// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a bicubic scale transform filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibicubicscaletransform?language=objc
type PBicubicScaleTransform interface {
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
	SetParameterC(value float64)
	HasSetParameterC() bool

	// optional
	ParameterC() float64
	HasParameterC() bool

	// optional
	SetAspectRatio(value float64)
	HasSetAspectRatio() bool

	// optional
	AspectRatio() float64
	HasAspectRatio() bool

	// optional
	SetParameterB(value float64)
	HasSetParameterB() bool

	// optional
	ParameterB() float64
	HasParameterB() bool
}

// A concrete type wrapper for the [PBicubicScaleTransform] protocol.
type BicubicScaleTransformWrapper struct {
	objc.Object
}

func (b_ BicubicScaleTransformWrapper) HasSetScale() bool {
	return b_.RespondsToSelector(objc.Sel("setScale:"))
}

// The scaling factor to use on the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibicubicscaletransform/3228078-scale?language=objc
func (b_ BicubicScaleTransformWrapper) SetScale(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setScale:"), value)
}

func (b_ BicubicScaleTransformWrapper) HasScale() bool {
	return b_.RespondsToSelector(objc.Sel("scale"))
}

// The scaling factor to use on the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibicubicscaletransform/3228078-scale?language=objc
func (b_ BicubicScaleTransformWrapper) Scale() float64 {
	rv := objc.Call[float64](b_, objc.Sel("scale"))
	return rv
}

func (b_ BicubicScaleTransformWrapper) HasSetInputImage() bool {
	return b_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibicubicscaletransform/3228075-inputimage?language=objc
func (b_ BicubicScaleTransformWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](b_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (b_ BicubicScaleTransformWrapper) HasInputImage() bool {
	return b_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibicubicscaletransform/3228075-inputimage?language=objc
func (b_ BicubicScaleTransformWrapper) InputImage() Image {
	rv := objc.Call[Image](b_, objc.Sel("inputImage"))
	return rv
}

func (b_ BicubicScaleTransformWrapper) HasSetParameterC() bool {
	return b_.RespondsToSelector(objc.Sel("setParameterC:"))
}

// The value of C to use for the cubic resampling function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibicubicscaletransform/3228077-parameterc?language=objc
func (b_ BicubicScaleTransformWrapper) SetParameterC(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setParameterC:"), value)
}

func (b_ BicubicScaleTransformWrapper) HasParameterC() bool {
	return b_.RespondsToSelector(objc.Sel("parameterC"))
}

// The value of C to use for the cubic resampling function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibicubicscaletransform/3228077-parameterc?language=objc
func (b_ BicubicScaleTransformWrapper) ParameterC() float64 {
	rv := objc.Call[float64](b_, objc.Sel("parameterC"))
	return rv
}

func (b_ BicubicScaleTransformWrapper) HasSetAspectRatio() bool {
	return b_.RespondsToSelector(objc.Sel("setAspectRatio:"))
}

// The additional horizontal scaling factor to use on the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibicubicscaletransform/3228074-aspectratio?language=objc
func (b_ BicubicScaleTransformWrapper) SetAspectRatio(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setAspectRatio:"), value)
}

func (b_ BicubicScaleTransformWrapper) HasAspectRatio() bool {
	return b_.RespondsToSelector(objc.Sel("aspectRatio"))
}

// The additional horizontal scaling factor to use on the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibicubicscaletransform/3228074-aspectratio?language=objc
func (b_ BicubicScaleTransformWrapper) AspectRatio() float64 {
	rv := objc.Call[float64](b_, objc.Sel("aspectRatio"))
	return rv
}

func (b_ BicubicScaleTransformWrapper) HasSetParameterB() bool {
	return b_.RespondsToSelector(objc.Sel("setParameterB:"))
}

// The value of B to use for the cubic resampling function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibicubicscaletransform/3228076-parameterb?language=objc
func (b_ BicubicScaleTransformWrapper) SetParameterB(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setParameterB:"), value)
}

func (b_ BicubicScaleTransformWrapper) HasParameterB() bool {
	return b_.RespondsToSelector(objc.Sel("parameterB"))
}

// The value of B to use for the cubic resampling function. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibicubicscaletransform/3228076-parameterb?language=objc
func (b_ BicubicScaleTransformWrapper) ParameterB() float64 {
	rv := objc.Call[float64](b_, objc.Sel("parameterB"))
	return rv
}
