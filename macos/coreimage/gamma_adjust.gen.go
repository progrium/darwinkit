// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a gamma adjust filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigammaadjust?language=objc
type PGammaAdjust interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetPower(value float64)
	HasSetPower() bool

	// optional
	Power() float64
	HasPower() bool
}

// A concrete type wrapper for the [PGammaAdjust] protocol.
type GammaAdjustWrapper struct {
	objc.Object
}

func (g_ GammaAdjustWrapper) HasSetInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigammaadjust/3228461-inputimage?language=objc
func (g_ GammaAdjustWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](g_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (g_ GammaAdjustWrapper) HasInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigammaadjust/3228461-inputimage?language=objc
func (g_ GammaAdjustWrapper) InputImage() Image {
	rv := objc.Call[Image](g_, objc.Sel("inputImage"))
	return rv
}

func (g_ GammaAdjustWrapper) HasSetPower() bool {
	return g_.RespondsToSelector(objc.Sel("setPower:"))
}

// A gamma value to use to correct image brightness. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigammaadjust/3228462-power?language=objc
func (g_ GammaAdjustWrapper) SetPower(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setPower:"), value)
}

func (g_ GammaAdjustWrapper) HasPower() bool {
	return g_.RespondsToSelector(objc.Sel("power"))
}

// A gamma value to use to correct image brightness. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigammaadjust/3228462-power?language=objc
func (g_ GammaAdjustWrapper) Power() float64 {
	rv := objc.Call[float64](g_, objc.Sel("power"))
	return rv
}
