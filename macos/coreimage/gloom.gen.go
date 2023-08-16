// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a gloom filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigloom?language=objc
type PGloom interface {
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

// A concrete type wrapper for the [PGloom] protocol.
type GloomWrapper struct {
	objc.Object
}

func (g_ GloomWrapper) HasSetInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigloom/3228477-inputimage?language=objc
func (g_ GloomWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](g_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (g_ GloomWrapper) HasInputImage() bool {
	return g_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigloom/3228477-inputimage?language=objc
func (g_ GloomWrapper) InputImage() Image {
	rv := objc.Call[Image](g_, objc.Sel("inputImage"))
	return rv
}

func (g_ GloomWrapper) HasSetIntensity() bool {
	return g_.RespondsToSelector(objc.Sel("setIntensity:"))
}

// The intensity of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigloom/3228478-intensity?language=objc
func (g_ GloomWrapper) SetIntensity(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setIntensity:"), value)
}

func (g_ GloomWrapper) HasIntensity() bool {
	return g_.RespondsToSelector(objc.Sel("intensity"))
}

// The intensity of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigloom/3228478-intensity?language=objc
func (g_ GloomWrapper) Intensity() float64 {
	rv := objc.Call[float64](g_, objc.Sel("intensity"))
	return rv
}

func (g_ GloomWrapper) HasSetRadius() bool {
	return g_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius, in pixels, of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigloom/3228479-radius?language=objc
func (g_ GloomWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](g_, objc.Sel("setRadius:"), value)
}

func (g_ GloomWrapper) HasRadius() bool {
	return g_.RespondsToSelector(objc.Sel("radius"))
}

// The radius, in pixels, of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cigloom/3228479-radius?language=objc
func (g_ GloomWrapper) Radius() float64 {
	rv := objc.Call[float64](g_, objc.Sel("radius"))
	return rv
}
