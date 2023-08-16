// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a mix filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimix?language=objc
type PMix interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetBackgroundImage(value Image)
	HasSetBackgroundImage() bool

	// optional
	BackgroundImage() IImage
	HasBackgroundImage() bool

	// optional
	SetAmount(value float64)
	HasSetAmount() bool

	// optional
	Amount() float64
	HasAmount() bool
}

// A concrete type wrapper for the [PMix] protocol.
type MixWrapper struct {
	objc.Object
}

func (m_ MixWrapper) HasSetInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as a foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimix/3228567-inputimage?language=objc
func (m_ MixWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](m_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (m_ MixWrapper) HasInputImage() bool {
	return m_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as a foreground image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimix/3228567-inputimage?language=objc
func (m_ MixWrapper) InputImage() Image {
	rv := objc.Call[Image](m_, objc.Sel("inputImage"))
	return rv
}

func (m_ MixWrapper) HasSetBackgroundImage() bool {
	return m_.RespondsToSelector(objc.Sel("setBackgroundImage:"))
}

// The image to use as a background image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimix/3228566-backgroundimage?language=objc
func (m_ MixWrapper) SetBackgroundImage(value IImage) {
	objc.Call[objc.Void](m_, objc.Sel("setBackgroundImage:"), objc.Ptr(value))
}

func (m_ MixWrapper) HasBackgroundImage() bool {
	return m_.RespondsToSelector(objc.Sel("backgroundImage"))
}

// The image to use as a background image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimix/3228566-backgroundimage?language=objc
func (m_ MixWrapper) BackgroundImage() Image {
	rv := objc.Call[Image](m_, objc.Sel("backgroundImage"))
	return rv
}

func (m_ MixWrapper) HasSetAmount() bool {
	return m_.RespondsToSelector(objc.Sel("setAmount:"))
}

// The amount of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimix/3228565-amount?language=objc
func (m_ MixWrapper) SetAmount(value float64) {
	objc.Call[objc.Void](m_, objc.Sel("setAmount:"), value)
}

func (m_ MixWrapper) HasAmount() bool {
	return m_.RespondsToSelector(objc.Sel("amount"))
}

// The amount of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cimix/3228565-amount?language=objc
func (m_ MixWrapper) Amount() float64 {
	rv := objc.Call[float64](m_, objc.Sel("amount"))
	return rv
}
