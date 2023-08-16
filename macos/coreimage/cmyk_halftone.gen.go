// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/macos/coregraphics"
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a CMYK halftone filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicmykhalftone?language=objc
type PCMYKHalftone interface {
	// optional
	SetWidth(value float64)
	HasSetWidth() bool

	// optional
	Width() float64
	HasWidth() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetGrayComponentReplacement(value float64)
	HasSetGrayComponentReplacement() bool

	// optional
	GrayComponentReplacement() float64
	HasGrayComponentReplacement() bool

	// optional
	SetSharpness(value float64)
	HasSetSharpness() bool

	// optional
	Sharpness() float64
	HasSharpness() bool

	// optional
	SetAngle(value float64)
	HasSetAngle() bool

	// optional
	Angle() float64
	HasAngle() bool

	// optional
	SetUnderColorRemoval(value float64)
	HasSetUnderColorRemoval() bool

	// optional
	UnderColorRemoval() float64
	HasUnderColorRemoval() bool

	// optional
	SetCenter(value coregraphics.Point)
	HasSetCenter() bool

	// optional
	Center() coregraphics.Point
	HasCenter() bool
}

// A concrete type wrapper for the [PCMYKHalftone] protocol.
type CMYKHalftoneWrapper struct {
	objc.Object
}

func (c_ CMYKHalftoneWrapper) HasSetWidth() bool {
	return c_.RespondsToSelector(objc.Sel("setWidth:"))
}

// The distance between dots in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicmykhalftone/3228103-width?language=objc
func (c_ CMYKHalftoneWrapper) SetWidth(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setWidth:"), value)
}

func (c_ CMYKHalftoneWrapper) HasWidth() bool {
	return c_.RespondsToSelector(objc.Sel("width"))
}

// The distance between dots in the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicmykhalftone/3228103-width?language=objc
func (c_ CMYKHalftoneWrapper) Width() float64 {
	rv := objc.Call[float64](c_, objc.Sel("width"))
	return rv
}

func (c_ CMYKHalftoneWrapper) HasSetInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicmykhalftone/3228100-inputimage?language=objc
func (c_ CMYKHalftoneWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](c_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (c_ CMYKHalftoneWrapper) HasInputImage() bool {
	return c_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicmykhalftone/3228100-inputimage?language=objc
func (c_ CMYKHalftoneWrapper) InputImage() Image {
	rv := objc.Call[Image](c_, objc.Sel("inputImage"))
	return rv
}

func (c_ CMYKHalftoneWrapper) HasSetGrayComponentReplacement() bool {
	return c_.RespondsToSelector(objc.Sel("setGrayComponentReplacement:"))
}

// The gray component replacement value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicmykhalftone/3228099-graycomponentreplacement?language=objc
func (c_ CMYKHalftoneWrapper) SetGrayComponentReplacement(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setGrayComponentReplacement:"), value)
}

func (c_ CMYKHalftoneWrapper) HasGrayComponentReplacement() bool {
	return c_.RespondsToSelector(objc.Sel("grayComponentReplacement"))
}

// The gray component replacement value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicmykhalftone/3228099-graycomponentreplacement?language=objc
func (c_ CMYKHalftoneWrapper) GrayComponentReplacement() float64 {
	rv := objc.Call[float64](c_, objc.Sel("grayComponentReplacement"))
	return rv
}

func (c_ CMYKHalftoneWrapper) HasSetSharpness() bool {
	return c_.RespondsToSelector(objc.Sel("setSharpness:"))
}

// The sharpness of the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicmykhalftone/3228101-sharpness?language=objc
func (c_ CMYKHalftoneWrapper) SetSharpness(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setSharpness:"), value)
}

func (c_ CMYKHalftoneWrapper) HasSharpness() bool {
	return c_.RespondsToSelector(objc.Sel("sharpness"))
}

// The sharpness of the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicmykhalftone/3228101-sharpness?language=objc
func (c_ CMYKHalftoneWrapper) Sharpness() float64 {
	rv := objc.Call[float64](c_, objc.Sel("sharpness"))
	return rv
}

func (c_ CMYKHalftoneWrapper) HasSetAngle() bool {
	return c_.RespondsToSelector(objc.Sel("setAngle:"))
}

// The angle of the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicmykhalftone/3228097-angle?language=objc
func (c_ CMYKHalftoneWrapper) SetAngle(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setAngle:"), value)
}

func (c_ CMYKHalftoneWrapper) HasAngle() bool {
	return c_.RespondsToSelector(objc.Sel("angle"))
}

// The angle of the pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicmykhalftone/3228097-angle?language=objc
func (c_ CMYKHalftoneWrapper) Angle() float64 {
	rv := objc.Call[float64](c_, objc.Sel("angle"))
	return rv
}

func (c_ CMYKHalftoneWrapper) HasSetUnderColorRemoval() bool {
	return c_.RespondsToSelector(objc.Sel("setUnderColorRemoval:"))
}

// The under color removal value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicmykhalftone/3228102-undercolorremoval?language=objc
func (c_ CMYKHalftoneWrapper) SetUnderColorRemoval(value float64) {
	objc.Call[objc.Void](c_, objc.Sel("setUnderColorRemoval:"), value)
}

func (c_ CMYKHalftoneWrapper) HasUnderColorRemoval() bool {
	return c_.RespondsToSelector(objc.Sel("underColorRemoval"))
}

// The under color removal value. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicmykhalftone/3228102-undercolorremoval?language=objc
func (c_ CMYKHalftoneWrapper) UnderColorRemoval() float64 {
	rv := objc.Call[float64](c_, objc.Sel("underColorRemoval"))
	return rv
}

func (c_ CMYKHalftoneWrapper) HasSetCenter() bool {
	return c_.RespondsToSelector(objc.Sel("setCenter:"))
}

// The x and y position to use as the center of the halftone pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicmykhalftone/3228098-center?language=objc
func (c_ CMYKHalftoneWrapper) SetCenter(value coregraphics.Point) {
	objc.Call[objc.Void](c_, objc.Sel("setCenter:"), value)
}

func (c_ CMYKHalftoneWrapper) HasCenter() bool {
	return c_.RespondsToSelector(objc.Sel("center"))
}

// The x and y position to use as the center of the halftone pattern. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cicmykhalftone/3228098-center?language=objc
func (c_ CMYKHalftoneWrapper) Center() coregraphics.Point {
	rv := objc.Call[coregraphics.Point](c_, objc.Sel("center"))
	return rv
}
