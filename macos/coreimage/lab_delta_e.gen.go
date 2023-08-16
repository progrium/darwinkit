// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a Lab Delta E filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilabdeltae?language=objc
type PLabDeltaE interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetImage2(value Image)
	HasSetImage2() bool

	// optional
	Image2() IImage
	HasImage2() bool
}

// A concrete type wrapper for the [PLabDeltaE] protocol.
type LabDeltaEWrapper struct {
	objc.Object
}

func (l_ LabDeltaEWrapper) HasSetInputImage() bool {
	return l_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The first input image for comparison. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilabdeltae/3228514-inputimage?language=objc
func (l_ LabDeltaEWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](l_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (l_ LabDeltaEWrapper) HasInputImage() bool {
	return l_.RespondsToSelector(objc.Sel("inputImage"))
}

// The first input image for comparison. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilabdeltae/3228514-inputimage?language=objc
func (l_ LabDeltaEWrapper) InputImage() Image {
	rv := objc.Call[Image](l_, objc.Sel("inputImage"))
	return rv
}

func (l_ LabDeltaEWrapper) HasSetImage2() bool {
	return l_.RespondsToSelector(objc.Sel("setImage2:"))
}

// The second input image for comparison. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilabdeltae/3228513-image2?language=objc
func (l_ LabDeltaEWrapper) SetImage2(value IImage) {
	objc.Call[objc.Void](l_, objc.Sel("setImage2:"), objc.Ptr(value))
}

func (l_ LabDeltaEWrapper) HasImage2() bool {
	return l_.RespondsToSelector(objc.Sel("image2"))
}

// The second input image for comparison. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cilabdeltae/3228513-image2?language=objc
func (l_ LabDeltaEWrapper) Image2() Image {
	rv := objc.Call[Image](l_, objc.Sel("image2"))
	return rv
}
