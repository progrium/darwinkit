// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a bokeh blur filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibokehblur?language=objc
type PBokehBlur interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetRingSize(value float64)
	HasSetRingSize() bool

	// optional
	RingSize() float64
	HasRingSize() bool

	// optional
	SetRadius(value float64)
	HasSetRadius() bool

	// optional
	Radius() float64
	HasRadius() bool

	// optional
	SetRingAmount(value float64)
	HasSetRingAmount() bool

	// optional
	RingAmount() float64
	HasRingAmount() bool

	// optional
	SetSoftness(value float64)
	HasSetSoftness() bool

	// optional
	Softness() float64
	HasSoftness() bool
}

// A concrete type wrapper for the [PBokehBlur] protocol.
type BokehBlurWrapper struct {
	objc.Object
}

func (b_ BokehBlurWrapper) HasSetInputImage() bool {
	return b_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibokehblur/3228088-inputimage?language=objc
func (b_ BokehBlurWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](b_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (b_ BokehBlurWrapper) HasInputImage() bool {
	return b_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibokehblur/3228088-inputimage?language=objc
func (b_ BokehBlurWrapper) InputImage() Image {
	rv := objc.Call[Image](b_, objc.Sel("inputImage"))
	return rv
}

func (b_ BokehBlurWrapper) HasSetRingSize() bool {
	return b_.RespondsToSelector(objc.Sel("setRingSize:"))
}

// The radius of the extra emphasis at the ring of the bokeh. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibokehblur/3228091-ringsize?language=objc
func (b_ BokehBlurWrapper) SetRingSize(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setRingSize:"), value)
}

func (b_ BokehBlurWrapper) HasRingSize() bool {
	return b_.RespondsToSelector(objc.Sel("ringSize"))
}

// The radius of the extra emphasis at the ring of the bokeh. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibokehblur/3228091-ringsize?language=objc
func (b_ BokehBlurWrapper) RingSize() float64 {
	rv := objc.Call[float64](b_, objc.Sel("ringSize"))
	return rv
}

func (b_ BokehBlurWrapper) HasSetRadius() bool {
	return b_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The radius of the blur, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibokehblur/3228089-radius?language=objc
func (b_ BokehBlurWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setRadius:"), value)
}

func (b_ BokehBlurWrapper) HasRadius() bool {
	return b_.RespondsToSelector(objc.Sel("radius"))
}

// The radius of the blur, in pixels. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibokehblur/3228089-radius?language=objc
func (b_ BokehBlurWrapper) Radius() float64 {
	rv := objc.Call[float64](b_, objc.Sel("radius"))
	return rv
}

func (b_ BokehBlurWrapper) HasSetRingAmount() bool {
	return b_.RespondsToSelector(objc.Sel("setRingAmount:"))
}

// The amount of extra emphasis at the ring of the bokeh. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibokehblur/3228090-ringamount?language=objc
func (b_ BokehBlurWrapper) SetRingAmount(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setRingAmount:"), value)
}

func (b_ BokehBlurWrapper) HasRingAmount() bool {
	return b_.RespondsToSelector(objc.Sel("ringAmount"))
}

// The amount of extra emphasis at the ring of the bokeh. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibokehblur/3228090-ringamount?language=objc
func (b_ BokehBlurWrapper) RingAmount() float64 {
	rv := objc.Call[float64](b_, objc.Sel("ringAmount"))
	return rv
}

func (b_ BokehBlurWrapper) HasSetSoftness() bool {
	return b_.RespondsToSelector(objc.Sel("setSoftness:"))
}

// The softness of the bokeh effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibokehblur/3228092-softness?language=objc
func (b_ BokehBlurWrapper) SetSoftness(value float64) {
	objc.Call[objc.Void](b_, objc.Sel("setSoftness:"), value)
}

func (b_ BokehBlurWrapper) HasSoftness() bool {
	return b_.RespondsToSelector(objc.Sel("softness"))
}

// The softness of the bokeh effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cibokehblur/3228092-softness?language=objc
func (b_ BokehBlurWrapper) Softness() float64 {
	rv := objc.Call[float64](b_, objc.Sel("softness"))
	return rv
}
