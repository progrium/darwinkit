// Code generated by DarwinKit. DO NOT EDIT.

package coreimage

import (
	"github.com/progrium/darwinkit/objc"
)

// The properties you use to configure a saliency map filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisaliencymap?language=objc
type PSaliencyMap interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() Image
	HasInputImage() bool
}

// ensure impl type implements protocol interface
var _ PSaliencyMap = (*SaliencyMapObject)(nil)

// A concrete type for the [PSaliencyMap] protocol.
type SaliencyMapObject struct {
	objc.Object
}

func (s_ SaliencyMapObject) HasSetInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisaliencymap/3228700-inputimage?language=objc
func (s_ SaliencyMapObject) SetInputImage(value Image) {
	objc.Call[objc.Void](s_, objc.Sel("setInputImage:"), value)
}

func (s_ SaliencyMapObject) HasInputImage() bool {
	return s_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cisaliencymap/3228700-inputimage?language=objc
func (s_ SaliencyMapObject) InputImage() Image {
	rv := objc.Call[Image](s_, objc.Sel("inputImage"))
	return rv
}
