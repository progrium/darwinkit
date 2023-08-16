// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a convolution filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciconvolution?language=objc
type PConvolution interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetWeights(value Vector)
	HasSetWeights() bool

	// optional
	Weights() IVector
	HasWeights() bool

	// optional
	SetBias(value float64)
	HasSetBias() bool

	// optional
	Bias() float64
	HasBias() bool
}

// A concrete type wrapper for the [PConvolution] protocol.
type ConvolutionWrapper struct {
	objc.Object
}

func (c_ ConvolutionWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciconvolution/3228186-inputimage?language=objc
func (c_ ConvolutionWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ ConvolutionWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciconvolution/3228186-inputimage?language=objc
func (c_ ConvolutionWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ ConvolutionWrapper) HasSetWeights() bool {
	return c_.RespondsToSelector(objc.Sel("setWeights:"))
}

// The convolution kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciconvolution/3228187-weights?language=objc
func (c_ ConvolutionWrapper) SetWeights(value IVector) {
	objc.Call[objc.Void](c_, objc.Sel("setWeights:"), objc.Ptr(value))
}

func (c_ ConvolutionWrapper) HasWeights() bool {
	return c_.RespondsToSelector(objc.Sel("weights"))
}

// The convolution kernel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciconvolution/3228187-weights?language=objc
func (c_ ConvolutionWrapper) Weights() Vector {
	rv := objc.Call[Vector](c_, objc.Sel("weights"))
	return rv
}

func (c_ ConvolutionWrapper) HasSetBias() bool {
	return c_.RespondsToSelector(objc.Sel("setBias:"))
}

// A value that’s added to each output pixel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciconvolution/3228185-bias?language=objc
func (c_ ConvolutionWrapper) SetBias(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setBias:"), value)
}

func (c_ ConvolutionWrapper) HasBias() bool {
	return c_.RespondsToSelector(objc.Sel("bias"))
}

// A value that’s added to each output pixel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciconvolution/3228185-bias?language=objc
func (c_ ConvolutionWrapper) Bias() float64 {
	rv := objc.Call[float64](c_, objc.Sel("bias"))
	return rv
}
