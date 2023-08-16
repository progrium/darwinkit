// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure an unsharp mask filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciunsharpmask?language=objc
type PUnsharpMask interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetIntensity(value float64)
	HasSetIntensity() bool

	// optional
	Intensity() float64
	HasIntensity() bool

	// optional
	SetRadius(value float64)
	HasSetRadius() bool

	// optional
	Radius() float64
	HasRadius() bool
}

// A concrete type wrapper for the [PUnsharpMask] protocol.
type UnsharpMaskWrapper struct {
	objc.Object
}

func (u_ UnsharpMaskWrapper) HasSetInputImage() bool {
	return u_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciunsharpmask/3228819-inputimage?language=objc
func (u_ UnsharpMaskWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](u_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (u_ UnsharpMaskWrapper) HasInputImage() bool {
	return u_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciunsharpmask/3228819-inputimage?language=objc
func (u_ UnsharpMaskWrapper) InputImage() Image {
	rv := objc.Call[Image](u_, objc.Sel("inputImage"))
	return rv
}

func (u_ UnsharpMaskWrapper) HasSetIntensity() bool {
	return u_.RespondsToSelector(objc.Sel("setIntensity:"))
}

// The intensity of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciunsharpmask/3228820-intensity?language=objc
func (u_ UnsharpMaskWrapper) SetIntensity(value float64) {
	objc.Call[objc.Void](u_, objc.Sel("setIntensity:"), value)
}

func (u_ UnsharpMaskWrapper) HasIntensity() bool {
	return u_.RespondsToSelector(objc.Sel("intensity"))
}

// The intensity of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciunsharpmask/3228820-intensity?language=objc
func (u_ UnsharpMaskWrapper) Intensity() float64 {
	rv := objc.Call[float64](u_, objc.Sel("intensity"))
	return rv
}

func (u_ UnsharpMaskWrapper) HasSetRadius() bool {
	return u_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius of the unsharp mask effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciunsharpmask/3228821-radius?language=objc
func (u_ UnsharpMaskWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](u_, objc.Sel("setRadius:"), value)
}

func (u_ UnsharpMaskWrapper) HasRadius() bool {
	return u_.RespondsToSelector(objc.Sel("radius"))
}

// The radius of the unsharp mask effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciunsharpmask/3228821-radius?language=objc
func (u_ UnsharpMaskWrapper) Radius() float64 {
	rv := objc.Call[float64](u_, objc.Sel("radius"))
	return rv
}
