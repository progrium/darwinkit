// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a disparity-to-depth filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisparitytodepth?language=objc
type PDisparityToDepth interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool
}

// A concrete type wrapper for the [PDisparityToDepth] protocol.
type DisparityToDepthWrapper struct {
	objc.Object
}

func (d_ DisparityToDepthWrapper) HasSetInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisparitytodepth/3228222-inputimage?language=objc
func (d_ DisparityToDepthWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](d_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (d_ DisparityToDepthWrapper) HasInputImage() bool {
	return d_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cidisparitytodepth/3228222-inputimage?language=objc
func (d_ DisparityToDepthWrapper) InputImage() Image {
	rv := objc.Call[Image](d_, objc.Sel("inputImage"))
	return rv
}
