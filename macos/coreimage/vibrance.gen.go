// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a vibrance filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civibrance?language=objc
type PVibrance interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetAmount(value float64)
	HasSetAmount() bool

	// optional
	Amount() float64
	HasAmount() bool
}

// A concrete type wrapper for the [PVibrance] protocol.
type VibranceWrapper struct {
	objc.Object
}

func (v_ VibranceWrapper) HasSetInputImage() bool {
	return v_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civibrance/3228824-inputimage?language=objc
func (v_ VibranceWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](v_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (v_ VibranceWrapper) HasInputImage() bool {
	return v_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civibrance/3228824-inputimage?language=objc
func (v_ VibranceWrapper) InputImage() Image {
	rv := objc.Call[Image](v_, objc.Sel("inputImage"))
	return rv
}

func (v_ VibranceWrapper) HasSetAmount() bool {
	return v_.RespondsToSelector(objc.Sel("setAmount:"))
}

// The amount to adjust the saturation by. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civibrance/3228823-amount?language=objc
func (v_ VibranceWrapper) SetAmount(value float64) {
	objc.Call[objc.Void](v_, objc.Sel("setAmount:"), value)
}

func (v_ VibranceWrapper) HasAmount() bool {
	return v_.RespondsToSelector(objc.Sel("amount"))
}

// The amount to adjust the saturation by. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/civibrance/3228823-amount?language=objc
func (v_ VibranceWrapper) Amount() float64 {
	rv := objc.Call[float64](v_, objc.Sel("amount"))
	return rv
}
