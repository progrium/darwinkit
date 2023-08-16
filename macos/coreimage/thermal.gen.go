// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a thermal filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cithermal?language=objc
type PThermal interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool
}

// A concrete type wrapper for the [PThermal] protocol.
type ThermalWrapper struct {
	objc.Object
}

func (t_ ThermalWrapper) HasSetInputImage() bool {
	return t_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cithermal/3228790-inputimage?language=objc
func (t_ ThermalWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](t_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (t_ ThermalWrapper) HasInputImage() bool {
	return t_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cithermal/3228790-inputimage?language=objc
func (t_ ThermalWrapper) InputImage() Image {
	rv := objc.Call[Image](t_, objc.Sel("inputImage"))
	return rv
}
