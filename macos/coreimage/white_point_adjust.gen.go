// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a white-point adjust filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciwhitepointadjust?language=objc
type PWhitePointAdjust interface {
	// optional
	SetColor(value Color)
	HasSetColor() bool

	// optional
	Color() IColor
	HasColor() bool

	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool
}

// A concrete type wrapper for the [PWhitePointAdjust] protocol.
type WhitePointAdjustWrapper struct {
	objc.Object
}

func (w_ WhitePointAdjustWrapper) HasSetColor() bool {
	return w_.RespondsToSelector(objc.Sel("setColor:"))
}

// A color to use as the white point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciwhitepointadjust/3228836-color?language=objc
func (w_ WhitePointAdjustWrapper) SetColor(value IColor) {
	objc.Call[objc.Void](w_, objc.Sel("setColor:"), objc.Ptr(value))
}

func (w_ WhitePointAdjustWrapper) HasColor() bool {
	return w_.RespondsToSelector(objc.Sel("color"))
}

// A color to use as the white point. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciwhitepointadjust/3228836-color?language=objc
func (w_ WhitePointAdjustWrapper) Color() Color {
	rv := objc.Call[Color](w_, objc.Sel("color"))
	return rv
}

func (w_ WhitePointAdjustWrapper) HasSetInputImage() bool {
	return w_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciwhitepointadjust/3228837-inputimage?language=objc
func (w_ WhitePointAdjustWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](w_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (w_ WhitePointAdjustWrapper) HasInputImage() bool {
	return w_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/ciwhitepointadjust/3228837-inputimage?language=objc
func (w_ WhitePointAdjustWrapper) InputImage() Image {
	rv := objc.Call[Image](w_, objc.Sel("inputImage"))
	return rv
}
