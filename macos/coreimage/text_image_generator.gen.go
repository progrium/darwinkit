// AUTO-GENERATED CODE, DO NOT MODIFY

package coreimage

import (
	"github.com/progrium/macdriver/objc"
)

// The properties you use to configure a text image generator filter. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citextimagegenerator?language=objc
type PTextImageGenerator interface {
	// optional
	SetScaleFactor(value float64)
	HasSetScaleFactor() bool

	// optional
	ScaleFactor() float64
	HasScaleFactor() bool

	// optional
	SetFontSize(value float64)
	HasSetFontSize() bool

	// optional
	FontSize() float64
	HasFontSize() bool

	// optional
	SetFontName(value string)
	HasSetFontName() bool

	// optional
	FontName() string
	HasFontName() bool

	// optional
	SetText(value string)
	HasSetText() bool

	// optional
	Text() string
	HasText() bool
}

// A concrete type wrapper for the [PTextImageGenerator] protocol.
type TextImageGeneratorWrapper struct {
	objc.Object
}

func (t_ TextImageGeneratorWrapper) HasSetScaleFactor() bool {
	return t_.RespondsToSelector(objc.Sel("setScaleFactor:"))
}

// The scale of the font to use for the generated text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citextimagegenerator/3228787-scalefactor?language=objc
func (t_ TextImageGeneratorWrapper) SetScaleFactor(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setScaleFactor:"), value)
}

func (t_ TextImageGeneratorWrapper) HasScaleFactor() bool {
	return t_.RespondsToSelector(objc.Sel("scaleFactor"))
}

// The scale of the font to use for the generated text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citextimagegenerator/3228787-scalefactor?language=objc
func (t_ TextImageGeneratorWrapper) ScaleFactor() float64 {
	rv := objc.Call[float64](t_, objc.Sel("scaleFactor"))
	return rv
}

func (t_ TextImageGeneratorWrapper) HasSetFontSize() bool {
	return t_.RespondsToSelector(objc.Sel("setFontSize:"))
}

// The size of the font to use for the generated text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citextimagegenerator/3228786-fontsize?language=objc
func (t_ TextImageGeneratorWrapper) SetFontSize(value float64) {
	objc.Call[objc.Void](t_, objc.Sel("setFontSize:"), value)
}

func (t_ TextImageGeneratorWrapper) HasFontSize() bool {
	return t_.RespondsToSelector(objc.Sel("fontSize"))
}

// The size of the font to use for the generated text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citextimagegenerator/3228786-fontsize?language=objc
func (t_ TextImageGeneratorWrapper) FontSize() float64 {
	rv := objc.Call[float64](t_, objc.Sel("fontSize"))
	return rv
}

func (t_ TextImageGeneratorWrapper) HasSetFontName() bool {
	return t_.RespondsToSelector(objc.Sel("setFontName:"))
}

// The name of the font to use for the generated text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citextimagegenerator/3228785-fontname?language=objc
func (t_ TextImageGeneratorWrapper) SetFontName(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setFontName:"), value)
}

func (t_ TextImageGeneratorWrapper) HasFontName() bool {
	return t_.RespondsToSelector(objc.Sel("fontName"))
}

// The name of the font to use for the generated text. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citextimagegenerator/3228785-fontname?language=objc
func (t_ TextImageGeneratorWrapper) FontName() string {
	rv := objc.Call[string](t_, objc.Sel("fontName"))
	return rv
}

func (t_ TextImageGeneratorWrapper) HasSetText() bool {
	return t_.RespondsToSelector(objc.Sel("setText:"))
}

// The text to render. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citextimagegenerator/3228788-text?language=objc
func (t_ TextImageGeneratorWrapper) SetText(value string) {
	objc.Call[objc.Void](t_, objc.Sel("setText:"), value)
}

func (t_ TextImageGeneratorWrapper) HasText() bool {
	return t_.RespondsToSelector(objc.Sel("text"))
}

// The text to render. [Full Topic]
//
// [Full Topic]: https://developer.apple.com/documentation/coreimage/citextimagegenerator/3228788-text?language=objc
func (t_ TextImageGeneratorWrapper) Text() string {
	rv := objc.Call[string](t_, objc.Sel("text"))
	return rv
}
