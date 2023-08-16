// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a saliency map filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisaliencymap?language=objc
type PSaliencyMap interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool
}

// A concrete type wrapper for the [PSaliencyMap] protocol.
type SaliencyMapWrapper struct {
	objc.Object
}

func (s_ SaliencyMapWrapper) HasSetInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisaliencymap/3228700-inputimage?language=objc
func (s_ SaliencyMapWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](s_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (s_ SaliencyMapWrapper) HasInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisaliencymap/3228700-inputimage?language=objc
func (s_ SaliencyMapWrapper) InputImage() Image {
	rv := objc.Call[Image](s_, objc.Sel("inputImage"))
	return rv
}
