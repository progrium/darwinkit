// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a depth-to-disparity filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidepthtodisparity?language=objc
type PDepthToDisparity interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool
}

// A concrete type wrapper for the [PDepthToDisparity] protocol.
type DepthToDisparityWrapper struct {
	objc.Object
}

func (d_ DepthToDisparityWrapper) HasSetInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidepthtodisparity/3228212-inputimage?language=objc
func (d_ DepthToDisparityWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](d_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (d_ DepthToDisparityWrapper) HasInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidepthtodisparity/3228212-inputimage?language=objc
func (d_ DepthToDisparityWrapper) InputImage() Image {
	rv := objc.Call[Image](d_, objc.Sel("inputImage"))
	return rv
}
