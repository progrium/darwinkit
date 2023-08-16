// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a noise reduction filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinoisereduction?language=objc
type PNoiseReduction interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetSharpness(value float64)
	HasSetSharpness() bool

	// optional
	Sharpness() float64
	HasSharpness() bool

	// optional
	SetNoiseLevel(value float64)
	HasSetNoiseLevel() bool

	// optional
	NoiseLevel() float64
	HasNoiseLevel() bool
}

// A concrete type wrapper for the [PNoiseReduction] protocol.
type NoiseReductionWrapper struct {
	objc.Object
}

func (n_ NoiseReductionWrapper) HasSetInputImage() bool {
	return n_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinoisereduction/3228595-inputimage?language=objc
func (n_ NoiseReductionWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](n_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (n_ NoiseReductionWrapper) HasInputImage() bool {
	return n_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinoisereduction/3228595-inputimage?language=objc
func (n_ NoiseReductionWrapper) InputImage() Image {
	rv := objc.Call[Image](n_, objc.Sel("inputImage"))
	return rv
}

func (n_ NoiseReductionWrapper) HasSetSharpness() bool {
	return n_.RespondsToSelector(objc.Sel("setSharpness:"))
}

// The sharpness of the final image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinoisereduction/3228597-sharpness?language=objc
func (n_ NoiseReductionWrapper) SetSharpness(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setSharpness:"), value)
}

func (n_ NoiseReductionWrapper) HasSharpness() bool {
	return n_.RespondsToSelector(objc.Sel("sharpness"))
}

// The sharpness of the final image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinoisereduction/3228597-sharpness?language=objc
func (n_ NoiseReductionWrapper) Sharpness() float64 {
	rv := objc.Call[float64](n_, objc.Sel("sharpness"))
	return rv
}

func (n_ NoiseReductionWrapper) HasSetNoiseLevel() bool {
	return n_.RespondsToSelector(objc.Sel("setNoiseLevel:"))
}

// The amount of noise reduction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinoisereduction/3228596-noiselevel?language=objc
func (n_ NoiseReductionWrapper) SetNoiseLevel(value float64) {
	objc.Call[objc.Void](n_, objc.Sel("setNoiseLevel:"), value)
}

func (n_ NoiseReductionWrapper) HasNoiseLevel() bool {
	return n_.RespondsToSelector(objc.Sel("noiseLevel"))
}

// The amount of noise reduction. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cinoisereduction/3228596-noiselevel?language=objc
func (n_ NoiseReductionWrapper) NoiseLevel() float64 {
	rv := objc.Call[float64](n_, objc.Sel("noiseLevel"))
	return rv
}
