// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a hue-saturation-value gradient filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihuesaturationvaluegradient?language=objc
type PHueSaturationValueGradient interface {
	// optional
	SetValue(value float64)
	HasSetValue() bool

	// optional
	Value() float64
	HasValue() bool

	// optional
	SetDither(value float64)
	HasSetDither() bool

	// optional
	Dither() float64
	HasDither() bool

	// optional
	SetColorSpace(value coregraphics.ColorSpaceRef)
	HasSetColorSpace() bool

	// optional
	ColorSpace() coregraphics.ColorSpaceRef
	HasColorSpace() bool

	// optional
	SetRadius(value float64)
	HasSetRadius() bool

	// optional
	Radius() float64
	HasRadius() bool

	// optional
	SetSoftness(value float64)
	HasSetSoftness() bool

	// optional
	Softness() float64
	HasSoftness() bool
}

// A concrete type wrapper for the [PHueSaturationValueGradient] protocol.
type HueSaturationValueGradientWrapper struct {
	objc.Object
}

func (h_ HueSaturationValueGradientWrapper) HasSetValue() bool {
	return h_.RespondsToSelector(objc.Sel("setValue:"))
}

// The lightness of the hue-saturation gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihuesaturationvaluegradient/3228506-value?language=objc
func (h_ HueSaturationValueGradientWrapper) SetValue(value float64) {
	objc.Call[objc.Void](h_, objc.Sel("setValue:"), value)
}

func (h_ HueSaturationValueGradientWrapper) HasValue() bool {
	return h_.RespondsToSelector(objc.Sel("value"))
}

// The lightness of the hue-saturation gradient. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihuesaturationvaluegradient/3228506-value?language=objc
func (h_ HueSaturationValueGradientWrapper) Value() float64 {
	rv := objc.Call[float64](h_, objc.Sel("value"))
	return rv
}

func (h_ HueSaturationValueGradientWrapper) HasSetDither() bool {
	return h_.RespondsToSelector(objc.Sel("setDither:"))
}

// A Boolean value specifying whether the dither the generated output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihuesaturationvaluegradient/3228503-dither?language=objc
func (h_ HueSaturationValueGradientWrapper) SetDither(value float64) {
	objc.Call[objc.Void](h_, objc.Sel("setDither:"), value)
}

func (h_ HueSaturationValueGradientWrapper) HasDither() bool {
	return h_.RespondsToSelector(objc.Sel("dither"))
}

// A Boolean value specifying whether the dither the generated output. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihuesaturationvaluegradient/3228503-dither?language=objc
func (h_ HueSaturationValueGradientWrapper) Dither() float64 {
	rv := objc.Call[float64](h_, objc.Sel("dither"))
	return rv
}

func (h_ HueSaturationValueGradientWrapper) HasSetColorSpace() bool {
	return h_.RespondsToSelector(objc.Sel("setColorSpace:"))
}

// The color space for the generated color wheel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihuesaturationvaluegradient/3228502-colorspace?language=objc
func (h_ HueSaturationValueGradientWrapper) SetColorSpace(value coregraphics.ColorSpaceRef) {
	objc.Call[objc.Void](h_, objc.Sel("setColorSpace:"), value)
}

func (h_ HueSaturationValueGradientWrapper) HasColorSpace() bool {
	return h_.RespondsToSelector(objc.Sel("colorSpace"))
}

// The color space for the generated color wheel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihuesaturationvaluegradient/3228502-colorspace?language=objc
func (h_ HueSaturationValueGradientWrapper) ColorSpace() coregraphics.ColorSpaceRef {
	rv := objc.Call[coregraphics.ColorSpaceRef](h_, objc.Sel("colorSpace"))
	return rv
}

func (h_ HueSaturationValueGradientWrapper) HasSetRadius() bool {
	return h_.RespondsToSelector(objc.Sel("setRadius:"))
}

// The distance from the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihuesaturationvaluegradient/3228504-radius?language=objc
func (h_ HueSaturationValueGradientWrapper) SetRadius(value float64) {
	objc.Call[objc.Void](h_, objc.Sel("setRadius:"), value)
}

func (h_ HueSaturationValueGradientWrapper) HasRadius() bool {
	return h_.RespondsToSelector(objc.Sel("radius"))
}

// The distance from the center of the effect. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihuesaturationvaluegradient/3228504-radius?language=objc
func (h_ HueSaturationValueGradientWrapper) Radius() float64 {
	rv := objc.Call[float64](h_, objc.Sel("radius"))
	return rv
}

func (h_ HueSaturationValueGradientWrapper) HasSetSoftness() bool {
	return h_.RespondsToSelector(objc.Sel("setSoftness:"))
}

// The softness of the generated color wheel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihuesaturationvaluegradient/3228505-softness?language=objc
func (h_ HueSaturationValueGradientWrapper) SetSoftness(value float64) {
	objc.Call[objc.Void](h_, objc.Sel("setSoftness:"), value)
}

func (h_ HueSaturationValueGradientWrapper) HasSoftness() bool {
	return h_.RespondsToSelector(objc.Sel("softness"))
}

// The softness of the generated color wheel. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cihuesaturationvaluegradient/3228505-softness?language=objc
func (h_ HueSaturationValueGradientWrapper) Softness() float64 {
	rv := objc.Call[float64](h_, objc.Sel("softness"))
	return rv
}
