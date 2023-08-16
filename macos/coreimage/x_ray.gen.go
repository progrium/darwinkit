// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure an X-ray filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cixray?language=objc
type PXRay interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool
}

// A concrete type wrapper for the [PXRay] protocol.
type XRayWrapper struct {
	objc.Object
}

func (x_ XRayWrapper) HasSetInputImage() bool {
	return x_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cixray/3228839-inputimage?language=objc
func (x_ XRayWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](x_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (x_ XRayWrapper) HasInputImage() bool {
	return x_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cixray/3228839-inputimage?language=objc
func (x_ XRayWrapper) InputImage() Image {
	rv := objc.Call[Image](x_, objc.Sel("inputImage"))
	return rv
}
