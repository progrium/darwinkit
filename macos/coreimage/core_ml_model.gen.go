// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coreml"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a Core ML model filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicoremlmodel?language=objc
type PCoreMLModel interface {
	// optional
	SetModel(value coreml.Model)
	HasSetModel() bool

	// optional
	Model() coreml.IModel
	HasModel() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetHeadIndex(value float64)
	HasSetHeadIndex() bool

	// optional
	HeadIndex() float64
	HasHeadIndex() bool

	// optional
	SetSoftmaxNormalization(value bool)
	HasSetSoftmaxNormalization() bool

	// optional
	SoftmaxNormalization() bool
	HasSoftmaxNormalization() bool
}

// A concrete type wrapper for the [PCoreMLModel] protocol.
type CoreMLModelWrapper struct {
	objc.Object
}

func (c_ CoreMLModelWrapper) HasSetModel() bool {
	return c_.RespondsToSelector(objc.Sel("setModel:"))
}

// The Core ML model used to apply the effect on the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicoremlmodel/3228197-model?language=objc
func (c_ CoreMLModelWrapper) SetModel(value coreml.IModel) {
	objc.Call[objc.Void](c_, objc.Sel("setModel:"), objc.Ptr(value))
}

func (c_ CoreMLModelWrapper) HasModel() bool {
	return c_.RespondsToSelector(objc.Sel("model"))
}

// The Core ML model used to apply the effect on the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicoremlmodel/3228197-model?language=objc
func (c_ CoreMLModelWrapper) Model() coreml.Model {
	rv := objc.Call[coreml.Model](c_, objc.Sel("model"))
	return rv
}

func (c_ CoreMLModelWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicoremlmodel/3228196-inputimage?language=objc
func (c_ CoreMLModelWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ CoreMLModelWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicoremlmodel/3228196-inputimage?language=objc
func (c_ CoreMLModelWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ CoreMLModelWrapper) HasSetHeadIndex() bool {
	return c_.RespondsToSelector(objc.Sel("setHeadIndex:"))
}

// A number that specifies which output of a multihead Core ML model applies the effect on the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicoremlmodel/3228195-headindex?language=objc
func (c_ CoreMLModelWrapper) SetHeadIndex(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setHeadIndex:"), value)
}

func (c_ CoreMLModelWrapper) HasHeadIndex() bool {
	return c_.RespondsToSelector(objc.Sel("headIndex"))
}

// A number that specifies which output of a multihead Core ML model applies the effect on the image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicoremlmodel/3228195-headindex?language=objc
func (c_ CoreMLModelWrapper) HeadIndex() float64 {
	rv := objc.Call[float64](c_, objc.Sel("headIndex"))
	return rv
}

func (c_ CoreMLModelWrapper) HasSetSoftmaxNormalization() bool {
	return c_.RespondsToSelector(objc.Sel("setSoftmaxNormalization:"))
}

// A Boolean value that specifies whether to apply Softmax normalization to the output of the model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicoremlmodel/3228198-softmaxnormalization?language=objc
func (c_ CoreMLModelWrapper) SetSoftmaxNormalization(value bool) {
	objc.Call[objc.Void](c_, objc.Sel("setSoftmaxNormalization:"), value)
}

func (c_ CoreMLModelWrapper) HasSoftmaxNormalization() bool {
	return c_.RespondsToSelector(objc.Sel("softmaxNormalization"))
}

// A Boolean value that specifies whether to apply Softmax normalization to the output of the model. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicoremlmodel/3228198-softmaxnormalization?language=objc
func (c_ CoreMLModelWrapper) SoftmaxNormalization() bool {
	rv := objc.Call[bool](c_, objc.Sel("softmaxNormalization"))
	return rv
}
