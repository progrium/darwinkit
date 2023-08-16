// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a temperature and tint filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citemperatureandtint?language=objc
type PTemperatureAndTint interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetTargetNeutral(value Vector)
	HasSetTargetNeutral() bool

	// optional
	TargetNeutral() IVector
	HasTargetNeutral() bool

	// optional
	SetNeutral(value Vector)
	HasSetNeutral() bool

	// optional
	Neutral() IVector
	HasNeutral() bool
}

// A concrete type wrapper for the [PTemperatureAndTint] protocol.
type TemperatureAndTintWrapper struct {
	objc.Object
}

func (t_ TemperatureAndTintWrapper) HasSetInputImage() bool {
	return t_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citemperatureandtint/3228781-inputimage?language=objc
func (t_ TemperatureAndTintWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](t_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (t_ TemperatureAndTintWrapper) HasInputImage() bool {
	return t_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citemperatureandtint/3228781-inputimage?language=objc
func (t_ TemperatureAndTintWrapper) InputImage() Image {
	rv := objc.Call[Image](t_, objc.Sel("inputImage"))
	return rv
}

func (t_ TemperatureAndTintWrapper) HasSetTargetNeutral() bool {
	return t_.RespondsToSelector(objc.Sel("setTargetNeutral:"))
}

// A vector containing the desired white point defined by color temperature and tint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citemperatureandtint/3228783-targetneutral?language=objc
func (t_ TemperatureAndTintWrapper) SetTargetNeutral(value IVector) {
	objc.Call[objc.Void](t_, objc.Sel("setTargetNeutral:"), objc.Ptr(value))
}

func (t_ TemperatureAndTintWrapper) HasTargetNeutral() bool {
	return t_.RespondsToSelector(objc.Sel("targetNeutral"))
}

// A vector containing the desired white point defined by color temperature and tint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citemperatureandtint/3228783-targetneutral?language=objc
func (t_ TemperatureAndTintWrapper) TargetNeutral() Vector {
	rv := objc.Call[Vector](t_, objc.Sel("targetNeutral"))
	return rv
}

func (t_ TemperatureAndTintWrapper) HasSetNeutral() bool {
	return t_.RespondsToSelector(objc.Sel("setNeutral:"))
}

// A vector containing the source white point defined by color temperature and tint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citemperatureandtint/3228782-neutral?language=objc
func (t_ TemperatureAndTintWrapper) SetNeutral(value IVector) {
	objc.Call[objc.Void](t_, objc.Sel("setNeutral:"), objc.Ptr(value))
}

func (t_ TemperatureAndTintWrapper) HasNeutral() bool {
	return t_.RespondsToSelector(objc.Sel("neutral"))
}

// A vector containing the source white point defined by color temperature and tint. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citemperatureandtint/3228782-neutral?language=objc
func (t_ TemperatureAndTintWrapper) Neutral() Vector {
	rv := objc.Call[Vector](t_, objc.Sel("neutral"))
	return rv
}
