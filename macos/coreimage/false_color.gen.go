// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a false color filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifalsecolor?language=objc
type PFalseColor interface {
	// optional
	SetInputImage(value Image)
	HasSetInputImage() bool

	// optional
	InputImage() IImage
	HasInputImage() bool

	// optional
	SetColor0(value Color)
	HasSetColor0() bool

	// optional
	Color0() IColor
	HasColor0() bool

	// optional
	SetColor1(value Color)
	HasSetColor1() bool

	// optional
	Color1() IColor
	HasColor1() bool
}

// A concrete type wrapper for the [PFalseColor] protocol.
type FalseColorWrapper struct {
	objc.Object
}

func (f_ FalseColorWrapper) HasSetInputImage() bool {
	return f_.RespondsToSelector(objc.Sel("setInputImage:"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifalsecolor/3228258-inputimage?language=objc
func (f_ FalseColorWrapper) SetInputImage(value IImage) {
	objc.Call[objc.Void](f_, objc.Sel("setInputImage:"), objc.Ptr(value))
}

func (f_ FalseColorWrapper) HasInputImage() bool {
	return f_.RespondsToSelector(objc.Sel("inputImage"))
}

// The image to use as an input image. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifalsecolor/3228258-inputimage?language=objc
func (f_ FalseColorWrapper) InputImage() Image {
	rv := objc.Call[Image](f_, objc.Sel("inputImage"))
	return rv
}

func (f_ FalseColorWrapper) HasSetColor0() bool {
	return f_.RespondsToSelector(objc.Sel("setColor0:"))
}

// The first color to use for the color ramp. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifalsecolor/3228256-color0?language=objc
func (f_ FalseColorWrapper) SetColor0(value IColor) {
	objc.Call[objc.Void](f_, objc.Sel("setColor0:"), objc.Ptr(value))
}

func (f_ FalseColorWrapper) HasColor0() bool {
	return f_.RespondsToSelector(objc.Sel("color0"))
}

// The first color to use for the color ramp. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifalsecolor/3228256-color0?language=objc
func (f_ FalseColorWrapper) Color0() Color {
	rv := objc.Call[Color](f_, objc.Sel("color0"))
	return rv
}

func (f_ FalseColorWrapper) HasSetColor1() bool {
	return f_.RespondsToSelector(objc.Sel("setColor1:"))
}

// The second color to use for the color ramp. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifalsecolor/3228257-color1?language=objc
func (f_ FalseColorWrapper) SetColor1(value IColor) {
	objc.Call[objc.Void](f_, objc.Sel("setColor1:"), objc.Ptr(value))
}

func (f_ FalseColorWrapper) HasColor1() bool {
	return f_.RespondsToSelector(objc.Sel("color1"))
}

// The second color to use for the color ramp. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/cifalsecolor/3228257-color1?language=objc
func (f_ FalseColorWrapper) Color1() Color {
	rv := objc.Call[Color](f_, objc.Sel("color1"))
	return rv
}
