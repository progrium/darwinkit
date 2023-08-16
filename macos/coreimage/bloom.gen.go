// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a bloom filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibloom?language=objc
type PBloom interface {
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

// A concrete type wrapper for the [PBloom] protocol.
type BloomWrapper struct {
	objc.Object
}

func (b_ BloomWrapper) HasSetInputImage() bool {
	return b_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibloom/3228084-inputimage?language=objc
func (b_ BloomWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](b_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (b_ BloomWrapper) HasInputImage() bool {
	return b_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibloom/3228084-inputimage?language=objc
func (b_ BloomWrapper) InputImage() Image {
	rv := objc.Call[Image](b_, objc.Sel("inputImage"))
	return rv
}

func (b_ BloomWrapper) HasSetIntensity() bool {
	return b_.RespondsToSelector(objc.Sel("setIntensity:"))
}

// The intensity of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibloom/3228085-intensity?language=objc
func (b_ BloomWrapper) SetIntensity(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setIntensity:"), value)
}

func (b_ BloomWrapper) HasIntensity() bool {
	return b_.RespondsToSelector(objc.Sel("intensity"))
}

// The intensity of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibloom/3228085-intensity?language=objc
func (b_ BloomWrapper) Intensity() float64 {
	rv := objc.Call[float64](b_, objc.Sel("intensity"))
	return rv
}

func (b_ BloomWrapper) HasSetRadius() bool {
	return b_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius, in pixels, of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibloom/3228086-radius?language=objc
func (b_ BloomWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setRadius:"), value)
}

func (b_ BloomWrapper) HasRadius() bool {
	return b_.RespondsToSelector(objc.Sel("radius"))
}

// The radius, in pixels, of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibloom/3228086-radius?language=objc
func (b_ BloomWrapper) Radius() float64 {
	rv := objc.Call[float64](b_, objc.Sel("radius"))
	return rv
}
